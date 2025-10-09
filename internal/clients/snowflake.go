/*
Copyright 2021 Upbound Inc.
*/

package clients

import (
	"context"
	"encoding/json"
	"os"
	"strings"

	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/crossplane/upjet/pkg/terraform"

	"github.com/valkiriaaquatica/provider-snowflake/apis/v1beta1"
)

const (
	// error messages
	errNoProviderConfig     = "no providerConfigRef provided"
	errGetProviderConfig    = "cannot get referenced ProviderConfig"
	errTrackUsage           = "cannot track ProviderConfig usage"
	errExtractCredentials   = "cannot extract credentials"
	errUnmarshalCredentials = "cannot unmarshal snowflake credentials as JSON"
)

const (
	keyOrganizationName     = "organization_name"
	keyAccountName          = "account_name"
	keyUser                 = "user"
	keyPassword             = "password"
	keyHost                 = "host"
	keyRole                 = "role"
	keyAuthenticator        = "authenticator"
	keyToken                = "token"
	keyPrivateKey           = "private_key"
	keyPrivateKeyPassphrase = "private_key_passphrase"

	envPreviewFeatures = "SNOWFLAKE_PREVIEW_FEATURES_ENABLED"
)

// TerraformSetupBuilder builds a terraform.SetupFn function that returns
// Terraform provider setup configuration
func TerraformSetupBuilder(version, providerSource, providerVersion string) terraform.SetupFn {
	return func(ctx context.Context, c client.Client, mg resource.Managed) (terraform.Setup, error) {
		ps := terraform.Setup{
			Version: version,
			Requirement: terraform.ProviderRequirement{
				Source:  providerSource,
				Version: providerVersion,
			},
		}

		if err := populateProviderConfig(ctx, c, mg, &ps); err != nil {
			return ps, err
		}

		// --- Add preview features automatically ---
		// Default list of preview features recommended for Snowflake 2.8.0+
		features := []string{
			"snowflake_database_datasource",
			"snowflake_storage_integration_resource",
			"snowflake_stage_resource",
			"snowflake_pipe_resource",
			"snowflake_table_resource",
			"snowflake_file_format_resource",
		}

		// Allow override or extension via environment variable
		if val := os.Getenv(envPreviewFeatures); val != "" {
			custom := strings.Split(val, ",")
			for _, f := range custom {
				trimmed := strings.TrimSpace(f)
				if trimmed != "" {
					features = append(features, trimmed)
				}
			}
		}

		// Deduplicate features to keep config clean
		unique := map[string]bool{}
		final := []string{}
		for _, f := range features {
			if !unique[f] {
				unique[f] = true
				final = append(final, f)
			}
		}

		// Inject into provider configuration
		if ps.Configuration == nil {
			ps.Configuration = map[string]any{}
		}
		ps.Configuration["preview_features_enabled"] = final
		// --- End preview feature setup ---

		return ps, nil
	}
}

func populateProviderConfig(ctx context.Context, c client.Client, mg resource.Managed, ps *terraform.Setup) error {
	ref := mg.GetProviderConfigReference()
	if ref == nil {
		return errors.New(errNoProviderConfig)
	}

	pc := &v1beta1.ProviderConfig{}
	if err := c.Get(ctx, types.NamespacedName{Name: ref.Name}, pc); err != nil {
		return errors.Wrap(err, errGetProviderConfig)
	}

	if err := trackProviderUsage(ctx, c, mg); err != nil {
		return err
	}

	creds, err := loadCredentials(ctx, c, pc)
	if err != nil {
		return err
	}

	ps.Configuration = buildConfiguration(creds)
	return nil
}

func trackProviderUsage(ctx context.Context, c client.Client, mg resource.Managed) error {
	tracker := resource.NewProviderConfigUsageTracker(c, &v1beta1.ProviderConfigUsage{})
	return errors.Wrap(tracker.Track(ctx, mg), errTrackUsage)
}

func loadCredentials(ctx context.Context, c client.Client, pc *v1beta1.ProviderConfig) (map[string]string, error) {
	data, err := resource.CommonCredentialExtractor(ctx, pc.Spec.Credentials.Source, c, pc.Spec.Credentials.CommonCredentialSelectors)
	if err != nil {
		return nil, errors.Wrap(err, errExtractCredentials)
	}
	creds := map[string]string{}
	if err := json.Unmarshal(data, &creds); err != nil {
		return nil, errors.Wrap(err, errUnmarshalCredentials)
	}
	return creds, nil
}

func buildConfiguration(creds map[string]string) map[string]any {
	cfg := map[string]any{}

	add := func(k string) {
		if v := creds[k]; v != "" {
			cfg[k] = v
		}
	}

	for _, k := range []string{
		keyOrganizationName, keyAccountName, keyUser, keyPassword,
	} {
		add(k)
	}

	for _, k := range []string{
		keyHost, keyRole, keyAuthenticator, keyToken,
		keyPrivateKey, keyPrivateKeyPassphrase,
	} {
		add(k)
	}

	if _, ok := cfg[keyAuthenticator]; !ok {
		if _, ok := cfg[keyToken]; ok {
			cfg[keyAuthenticator] = "PROGRAMMATIC_ACCESS_TOKEN"
		}
		if _, ok := cfg[keyPrivateKey]; ok {
			cfg[keyAuthenticator] = "SNOWFLAKE_JWT"
		}
	}

	return cfg
}

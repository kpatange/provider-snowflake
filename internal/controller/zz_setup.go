// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	authenticationpolicyattachment "github.com/valkiriaaquatica/provider-snowflake/internal/controller/account/authenticationpolicyattachment"
	passwordpolicyattachment "github.com/valkiriaaquatica/provider-snowflake/internal/controller/account/passwordpolicyattachment"
	integration "github.com/valkiriaaquatica/provider-snowflake/internal/controller/api/integration"
	pool "github.com/valkiriaaquatica/provider-snowflake/internal/controller/compute/pool"
	searchservice "github.com/valkiriaaquatica/provider-snowflake/internal/controller/cortex/searchservice"
	account "github.com/valkiriaaquatica/provider-snowflake/internal/controller/current/account"
	organizationaccount "github.com/valkiriaaquatica/provider-snowflake/internal/controller/current/organizationaccount"
	table "github.com/valkiriaaquatica/provider-snowflake/internal/controller/dynamic/table"
	notificationintegration "github.com/valkiriaaquatica/provider-snowflake/internal/controller/email/notificationintegration"
	function "github.com/valkiriaaquatica/provider-snowflake/internal/controller/external/function"
	tableexternal "github.com/valkiriaaquatica/provider-snowflake/internal/controller/external/table"
	volume "github.com/valkiriaaquatica/provider-snowflake/internal/controller/external/volume"
	group "github.com/valkiriaaquatica/provider-snowflake/internal/controller/failover/group"
	java "github.com/valkiriaaquatica/provider-snowflake/internal/controller/function/java"
	javascript "github.com/valkiriaaquatica/provider-snowflake/internal/controller/function/javascript"
	python "github.com/valkiriaaquatica/provider-snowflake/internal/controller/function/python"
	scala "github.com/valkiriaaquatica/provider-snowflake/internal/controller/function/scala"
	sql "github.com/valkiriaaquatica/provider-snowflake/internal/controller/function/sql"
	repository "github.com/valkiriaaquatica/provider-snowflake/internal/controller/git/repository"
	databaserole "github.com/valkiriaaquatica/provider-snowflake/internal/controller/grant/databaserole"
	repositoryimage "github.com/valkiriaaquatica/provider-snowflake/internal/controller/image/repository"
	service "github.com/valkiriaaquatica/provider-snowflake/internal/controller/job/service"
	accountmanaged "github.com/valkiriaaquatica/provider-snowflake/internal/controller/managed/account"
	view "github.com/valkiriaaquatica/provider-snowflake/internal/controller/materialized/view"
	policyattachment "github.com/valkiriaaquatica/provider-snowflake/internal/controller/network/policyattachment"
	rule "github.com/valkiriaaquatica/provider-snowflake/internal/controller/network/rule"
	integrationnotification "github.com/valkiriaaquatica/provider-snowflake/internal/controller/notification/integration"
	parameter "github.com/valkiriaaquatica/provider-snowflake/internal/controller/object/parameter"
	policy "github.com/valkiriaaquatica/provider-snowflake/internal/controller/password/policy"
	javaprocedure "github.com/valkiriaaquatica/provider-snowflake/internal/controller/procedure/java"
	javascriptprocedure "github.com/valkiriaaquatica/provider-snowflake/internal/controller/procedure/javascript"
	pythonprocedure "github.com/valkiriaaquatica/provider-snowflake/internal/controller/procedure/python"
	scalaprocedure "github.com/valkiriaaquatica/provider-snowflake/internal/controller/procedure/scala"
	sqlprocedure "github.com/valkiriaaquatica/provider-snowflake/internal/controller/procedure/sql"
	providerconfig "github.com/valkiriaaquatica/provider-snowflake/internal/controller/providerconfig"
	alert "github.com/valkiriaaquatica/provider-snowflake/internal/controller/snowflake/alert"
	listing "github.com/valkiriaaquatica/provider-snowflake/internal/controller/snowflake/listing"
	pipe "github.com/valkiriaaquatica/provider-snowflake/internal/controller/snowflake/pipe"
	sequence "github.com/valkiriaaquatica/provider-snowflake/internal/controller/snowflake/sequence"
	servicesnowflake "github.com/valkiriaaquatica/provider-snowflake/internal/controller/snowflake/service"
	share "github.com/valkiriaaquatica/provider-snowflake/internal/controller/snowflake/share"
	stage "github.com/valkiriaaquatica/provider-snowflake/internal/controller/snowflake/stage"
	tablesnowflake "github.com/valkiriaaquatica/provider-snowflake/internal/controller/snowflake/table"
	accountsnowflakeaccount "github.com/valkiriaaquatica/provider-snowflake/internal/controller/snowflakeaccount/account"
	parametersnowflakeaccountparameter "github.com/valkiriaaquatica/provider-snowflake/internal/controller/snowflakeaccountparameter/parameter"
	role "github.com/valkiriaaquatica/provider-snowflake/internal/controller/snowflakeaccountrole/role"
	authenticationintegrationwithauthorizationcodegrant "github.com/valkiriaaquatica/provider-snowflake/internal/controller/snowflakeapiauthenticationintegrationwithauthorizationcodegrant/authenticationintegrationwithauthorizationcodegrant"
	authenticationintegrationwithclientcredentials "github.com/valkiriaaquatica/provider-snowflake/internal/controller/snowflakeapiauthenticationintegrationwithclientcredentials/authenticationintegrationwithclientcredentials"
	authenticationintegrationwithjwtbearer "github.com/valkiriaaquatica/provider-snowflake/internal/controller/snowflakeapiauthenticationintegrationwithjwtbearer/authenticationintegrationwithjwtbearer"
	policysnowflakeauthenticationpolicy "github.com/valkiriaaquatica/provider-snowflake/internal/controller/snowflakeauthenticationpolicy/policy"
	database "github.com/valkiriaaquatica/provider-snowflake/internal/controller/snowflakedatabase/database"
	rolesnowflakedatabaserole "github.com/valkiriaaquatica/provider-snowflake/internal/controller/snowflakedatabaserole/role"
	execute "github.com/valkiriaaquatica/provider-snowflake/internal/controller/snowflakeexecute/execute"
	oauthintegration "github.com/valkiriaaquatica/provider-snowflake/internal/controller/snowflakeexternaloauthintegration/oauthintegration"
	format "github.com/valkiriaaquatica/provider-snowflake/internal/controller/snowflakefileformat/format"
	accountrole "github.com/valkiriaaquatica/provider-snowflake/internal/controller/snowflakegrantaccountrole/accountrole"
	applicationrole "github.com/valkiriaaquatica/provider-snowflake/internal/controller/snowflakegrantapplicationrole/applicationrole"
	ownership "github.com/valkiriaaquatica/provider-snowflake/internal/controller/snowflakegrantownership/ownership"
	privilegestoaccountrole "github.com/valkiriaaquatica/provider-snowflake/internal/controller/snowflakegrantprivilegestoaccountrole/privilegestoaccountrole"
	privilegestodatabaserole "github.com/valkiriaaquatica/provider-snowflake/internal/controller/snowflakegrantprivilegestodatabaserole/privilegestodatabaserole"
	privilegestoshare "github.com/valkiriaaquatica/provider-snowflake/internal/controller/snowflakegrantprivilegestoshare/privilegestoshare"
	serviceuser "github.com/valkiriaaquatica/provider-snowflake/internal/controller/snowflakelegacyserviceuser/serviceuser"
	policysnowflakemaskingpolicy "github.com/valkiriaaquatica/provider-snowflake/internal/controller/snowflakemaskingpolicy/policy"
	policysnowflakenetworkpolicy "github.com/valkiriaaquatica/provider-snowflake/internal/controller/snowflakenetworkpolicy/policy"
	integrationforcustomclients "github.com/valkiriaaquatica/provider-snowflake/internal/controller/snowflakeoauthintegrationforcustomclients/integrationforcustomclients"
	integrationforpartnerapplications "github.com/valkiriaaquatica/provider-snowflake/internal/controller/snowflakeoauthintegrationforpartnerapplications/integrationforpartnerapplications"
	connection "github.com/valkiriaaquatica/provider-snowflake/internal/controller/snowflakeprimaryconnection/connection"
	monitor "github.com/valkiriaaquatica/provider-snowflake/internal/controller/snowflakeresourcemonitor/monitor"
	accesspolicy "github.com/valkiriaaquatica/provider-snowflake/internal/controller/snowflakerowaccesspolicy/accesspolicy"
	integrationsnowflakesaml2integration "github.com/valkiriaaquatica/provider-snowflake/internal/controller/snowflakesaml2integration/integration"
	schema "github.com/valkiriaaquatica/provider-snowflake/internal/controller/snowflakeschema/schema"
	integrationsnowflakescimintegration "github.com/valkiriaaquatica/provider-snowflake/internal/controller/snowflakescimintegration/integration"
	connectionsnowflakesecondaryconnection "github.com/valkiriaaquatica/provider-snowflake/internal/controller/snowflakesecondaryconnection/connection"
	databasesnowflakesecondarydatabase "github.com/valkiriaaquatica/provider-snowflake/internal/controller/snowflakesecondarydatabase/database"
	withauthorizationcodegrant "github.com/valkiriaaquatica/provider-snowflake/internal/controller/snowflakesecretwithauthorizationcodegrant/withauthorizationcodegrant"
	withbasicauthentication "github.com/valkiriaaquatica/provider-snowflake/internal/controller/snowflakesecretwithbasicauthentication/withbasicauthentication"
	withclientcredentials "github.com/valkiriaaquatica/provider-snowflake/internal/controller/snowflakesecretwithclientcredentials/withclientcredentials"
	withgenericstring "github.com/valkiriaaquatica/provider-snowflake/internal/controller/snowflakesecretwithgenericstring/withgenericstring"
	user "github.com/valkiriaaquatica/provider-snowflake/internal/controller/snowflakeserviceuser/user"
	databasesnowflakeshareddatabase "github.com/valkiriaaquatica/provider-snowflake/internal/controller/snowflakeshareddatabase/database"
	streamlit "github.com/valkiriaaquatica/provider-snowflake/internal/controller/snowflakestreamlit/streamlit"
	ondirectorytable "github.com/valkiriaaquatica/provider-snowflake/internal/controller/snowflakestreamondirectorytable/ondirectorytable"
	onexternaltable "github.com/valkiriaaquatica/provider-snowflake/internal/controller/snowflakestreamonexternaltable/onexternaltable"
	ontable "github.com/valkiriaaquatica/provider-snowflake/internal/controller/snowflakestreamontable/ontable"
	onview "github.com/valkiriaaquatica/provider-snowflake/internal/controller/snowflakestreamonview/onview"
	tag "github.com/valkiriaaquatica/provider-snowflake/internal/controller/snowflaketag/tag"
	association "github.com/valkiriaaquatica/provider-snowflake/internal/controller/snowflaketagassociation/association"
	task "github.com/valkiriaaquatica/provider-snowflake/internal/controller/snowflaketask/task"
	usersnowflakeuser "github.com/valkiriaaquatica/provider-snowflake/internal/controller/snowflakeuser/user"
	viewsnowflakeview "github.com/valkiriaaquatica/provider-snowflake/internal/controller/snowflakeview/view"
	warehouse "github.com/valkiriaaquatica/provider-snowflake/internal/controller/snowflakewarehouse/warehouse"
	integrationstorage "github.com/valkiriaaquatica/provider-snowflake/internal/controller/storage/integration"
	columnmaskingpolicyapplication "github.com/valkiriaaquatica/provider-snowflake/internal/controller/table/columnmaskingpolicyapplication"
	constraint "github.com/valkiriaaquatica/provider-snowflake/internal/controller/table/constraint"
	authenticationpolicyattachmentuser "github.com/valkiriaaquatica/provider-snowflake/internal/controller/user/authenticationpolicyattachment"
	passwordpolicyattachmentuser "github.com/valkiriaaquatica/provider-snowflake/internal/controller/user/passwordpolicyattachment"
	programmaticaccesstoken "github.com/valkiriaaquatica/provider-snowflake/internal/controller/user/programmaticaccesstoken"
	publickeys "github.com/valkiriaaquatica/provider-snowflake/internal/controller/user/publickeys"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		authenticationpolicyattachment.Setup,
		passwordpolicyattachment.Setup,
		integration.Setup,
		pool.Setup,
		searchservice.Setup,
		account.Setup,
		organizationaccount.Setup,
		table.Setup,
		notificationintegration.Setup,
		function.Setup,
		tableexternal.Setup,
		volume.Setup,
		group.Setup,
		java.Setup,
		javascript.Setup,
		python.Setup,
		scala.Setup,
		sql.Setup,
		repository.Setup,
		databaserole.Setup,
		repositoryimage.Setup,
		service.Setup,
		accountmanaged.Setup,
		view.Setup,
		policyattachment.Setup,
		rule.Setup,
		integrationnotification.Setup,
		parameter.Setup,
		policy.Setup,
		javaprocedure.Setup,
		javascriptprocedure.Setup,
		pythonprocedure.Setup,
		scalaprocedure.Setup,
		sqlprocedure.Setup,
		providerconfig.Setup,
		alert.Setup,
		listing.Setup,
		pipe.Setup,
		sequence.Setup,
		servicesnowflake.Setup,
		share.Setup,
		stage.Setup,
		tablesnowflake.Setup,
		accountsnowflakeaccount.Setup,
		parametersnowflakeaccountparameter.Setup,
		role.Setup,
		authenticationintegrationwithauthorizationcodegrant.Setup,
		authenticationintegrationwithclientcredentials.Setup,
		authenticationintegrationwithjwtbearer.Setup,
		policysnowflakeauthenticationpolicy.Setup,
		database.Setup,
		rolesnowflakedatabaserole.Setup,
		execute.Setup,
		oauthintegration.Setup,
		format.Setup,
		accountrole.Setup,
		applicationrole.Setup,
		ownership.Setup,
		privilegestoaccountrole.Setup,
		privilegestodatabaserole.Setup,
		privilegestoshare.Setup,
		serviceuser.Setup,
		policysnowflakemaskingpolicy.Setup,
		policysnowflakenetworkpolicy.Setup,
		integrationforcustomclients.Setup,
		integrationforpartnerapplications.Setup,
		connection.Setup,
		monitor.Setup,
		accesspolicy.Setup,
		integrationsnowflakesaml2integration.Setup,
		schema.Setup,
		integrationsnowflakescimintegration.Setup,
		connectionsnowflakesecondaryconnection.Setup,
		databasesnowflakesecondarydatabase.Setup,
		withauthorizationcodegrant.Setup,
		withbasicauthentication.Setup,
		withclientcredentials.Setup,
		withgenericstring.Setup,
		user.Setup,
		databasesnowflakeshareddatabase.Setup,
		streamlit.Setup,
		ondirectorytable.Setup,
		onexternaltable.Setup,
		ontable.Setup,
		onview.Setup,
		tag.Setup,
		association.Setup,
		task.Setup,
		usersnowflakeuser.Setup,
		viewsnowflakeview.Setup,
		warehouse.Setup,
		integrationstorage.Setup,
		columnmaskingpolicyapplication.Setup,
		constraint.Setup,
		authenticationpolicyattachmentuser.Setup,
		passwordpolicyattachmentuser.Setup,
		programmaticaccesstoken.Setup,
		publickeys.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

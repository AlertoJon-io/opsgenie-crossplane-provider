/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/terrajet/pkg/controller"

	policy "github.com/crossplane-contrib/provider-jet-opsgenie-provider/internal/controller/alert/policy"
	integration "github.com/crossplane-contrib/provider-jet-opsgenie-provider/internal/controller/api/integration"
	role "github.com/crossplane-contrib/provider-jet-opsgenie-provider/internal/controller/custom/role"
	integrationemail "github.com/crossplane-contrib/provider-jet-opsgenie-provider/internal/controller/email/integration"
	template "github.com/crossplane-contrib/provider-jet-opsgenie-provider/internal/controller/incident/template"
	action "github.com/crossplane-contrib/provider-jet-opsgenie-provider/internal/controller/integration/action"
	policynotification "github.com/crossplane-contrib/provider-jet-opsgenie-provider/internal/controller/notification/policy"
	rule "github.com/crossplane-contrib/provider-jet-opsgenie-provider/internal/controller/notification/rule"
	escalation "github.com/crossplane-contrib/provider-jet-opsgenie-provider/internal/controller/opsgenie/escalation"
	heartbeat "github.com/crossplane-contrib/provider-jet-opsgenie-provider/internal/controller/opsgenie/heartbeat"
	maintenance "github.com/crossplane-contrib/provider-jet-opsgenie-provider/internal/controller/opsgenie/maintenance"
	schedule "github.com/crossplane-contrib/provider-jet-opsgenie-provider/internal/controller/opsgenie/schedule"
	service "github.com/crossplane-contrib/provider-jet-opsgenie-provider/internal/controller/opsgenie/service"
	team "github.com/crossplane-contrib/provider-jet-opsgenie-provider/internal/controller/opsgenie/team"
	user "github.com/crossplane-contrib/provider-jet-opsgenie-provider/internal/controller/opsgenie/user"
	providerconfig "github.com/crossplane-contrib/provider-jet-opsgenie-provider/internal/controller/providerconfig"
	rotation "github.com/crossplane-contrib/provider-jet-opsgenie-provider/internal/controller/schedule/rotation"
	incidentrule "github.com/crossplane-contrib/provider-jet-opsgenie-provider/internal/controller/service/incidentrule"
	routingrule "github.com/crossplane-contrib/provider-jet-opsgenie-provider/internal/controller/team/routingrule"
	contact "github.com/crossplane-contrib/provider-jet-opsgenie-provider/internal/controller/user/contact"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		policy.Setup,
		integration.Setup,
		role.Setup,
		integrationemail.Setup,
		template.Setup,
		action.Setup,
		policynotification.Setup,
		rule.Setup,
		escalation.Setup,
		heartbeat.Setup,
		maintenance.Setup,
		schedule.Setup,
		service.Setup,
		team.Setup,
		user.Setup,
		providerconfig.Setup,
		rotation.Setup,
		incidentrule.Setup,
		routingrule.Setup,
		contact.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

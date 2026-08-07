package config

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

// ExternalNameConfigs contains all external name configurations for this
// provider.
//
// Every authentik resource is identified by a server-assigned primary key (a
// UUID, or a slug the API allocates), so IdentifierFromProvider is the correct
// external-name strategy across the board. Resources are listed here to opt
// them into generation: ExternalNameConfigured feeds the provider's include
// list.
var ExternalNameConfigs = map[string]config.ExternalName{
	// Core objects.
	"authentik_application":             config.IdentifierFromProvider,
	"authentik_application_entitlement": config.IdentifierFromProvider,
	"authentik_blueprint":               config.IdentifierFromProvider,
	"authentik_brand":                   config.IdentifierFromProvider,
	"authentik_certificate_key_pair":    config.IdentifierFromProvider,
	"authentik_enterprise_license":      config.IdentifierFromProvider,
	"authentik_group":                   config.IdentifierFromProvider,
	"authentik_system_settings":         config.IdentifierFromProvider,
	"authentik_task_schedule":           config.IdentifierFromProvider,
	"authentik_token":                   config.IdentifierFromProvider,
	"authentik_user":                    config.IdentifierFromProvider,

	// Flows and stage bindings.
	"authentik_flow":               config.IdentifierFromProvider,
	"authentik_flow_stage_binding": config.IdentifierFromProvider,

	// Stages.
	"authentik_stage_account_lockdown":            config.IdentifierFromProvider,
	"authentik_stage_authenticator_duo":           config.IdentifierFromProvider,
	"authentik_stage_authenticator_email":         config.IdentifierFromProvider,
	"authentik_stage_authenticator_endpoint_gdtc": config.IdentifierFromProvider,
	"authentik_stage_authenticator_sms":           config.IdentifierFromProvider,
	"authentik_stage_authenticator_static":        config.IdentifierFromProvider,
	"authentik_stage_authenticator_totp":          config.IdentifierFromProvider,
	"authentik_stage_authenticator_validate":      config.IdentifierFromProvider,
	"authentik_stage_authenticator_webauthn":      config.IdentifierFromProvider,
	"authentik_stage_captcha":                     config.IdentifierFromProvider,
	"authentik_stage_consent":                     config.IdentifierFromProvider,
	"authentik_stage_deny":                        config.IdentifierFromProvider,
	"authentik_stage_dummy":                       config.IdentifierFromProvider,
	"authentik_stage_email":                       config.IdentifierFromProvider,
	"authentik_stage_endpoints":                   config.IdentifierFromProvider,
	"authentik_stage_identification":              config.IdentifierFromProvider,
	"authentik_stage_invitation":                  config.IdentifierFromProvider,
	"authentik_stage_mutual_tls":                  config.IdentifierFromProvider,
	"authentik_stage_password":                    config.IdentifierFromProvider,
	"authentik_stage_prompt":                      config.IdentifierFromProvider,
	"authentik_stage_prompt_field":                config.IdentifierFromProvider,
	"authentik_stage_redirect":                    config.IdentifierFromProvider,
	"authentik_stage_source":                      config.IdentifierFromProvider,
	"authentik_stage_user_delete":                 config.IdentifierFromProvider,
	"authentik_stage_user_login":                  config.IdentifierFromProvider,
	"authentik_stage_user_logout":                 config.IdentifierFromProvider,
	"authentik_stage_user_write":                  config.IdentifierFromProvider,

	// Policies.
	"authentik_policy_binding":         config.IdentifierFromProvider,
	"authentik_policy_dummy":           config.IdentifierFromProvider,
	"authentik_policy_event_matcher":   config.IdentifierFromProvider,
	"authentik_policy_expiry":          config.IdentifierFromProvider,
	"authentik_policy_expression":      config.IdentifierFromProvider,
	"authentik_policy_geoip":           config.IdentifierFromProvider,
	"authentik_policy_password":        config.IdentifierFromProvider,
	"authentik_policy_reputation":      config.IdentifierFromProvider,
	"authentik_policy_unique_password": config.IdentifierFromProvider,

	// Property mappings.
	"authentik_property_mapping_notification":              config.IdentifierFromProvider,
	"authentik_property_mapping_provider_google_workspace": config.IdentifierFromProvider,
	"authentik_property_mapping_provider_microsoft_entra":  config.IdentifierFromProvider,
	"authentik_property_mapping_provider_rac":              config.IdentifierFromProvider,
	"authentik_property_mapping_provider_radius":           config.IdentifierFromProvider,
	"authentik_property_mapping_provider_saml":             config.IdentifierFromProvider,
	"authentik_property_mapping_provider_scim":             config.IdentifierFromProvider,
	"authentik_property_mapping_provider_scope":            config.IdentifierFromProvider,
	"authentik_property_mapping_source_kerberos":           config.IdentifierFromProvider,
	"authentik_property_mapping_source_ldap":               config.IdentifierFromProvider,
	"authentik_property_mapping_source_oauth":              config.IdentifierFromProvider,
	"authentik_property_mapping_source_plex":               config.IdentifierFromProvider,
	"authentik_property_mapping_source_saml":               config.IdentifierFromProvider,
	"authentik_property_mapping_source_scim":               config.IdentifierFromProvider,

	// Providers.
	"authentik_provider_google_workspace": config.IdentifierFromProvider,
	"authentik_provider_ldap":             config.IdentifierFromProvider,
	"authentik_provider_microsoft_entra":  config.IdentifierFromProvider,
	"authentik_provider_oauth2":           config.IdentifierFromProvider,
	"authentik_provider_proxy":            config.IdentifierFromProvider,
	"authentik_provider_rac":              config.IdentifierFromProvider,
	"authentik_provider_radius":           config.IdentifierFromProvider,
	"authentik_provider_saml":             config.IdentifierFromProvider,
	"authentik_provider_scim":             config.IdentifierFromProvider,
	"authentik_provider_ssf":              config.IdentifierFromProvider,
	"authentik_provider_ws_federation":    config.IdentifierFromProvider,

	// Sources.
	"authentik_source_kerberos": config.IdentifierFromProvider,
	"authentik_source_ldap":     config.IdentifierFromProvider,
	"authentik_source_oauth":    config.IdentifierFromProvider,
	"authentik_source_plex":     config.IdentifierFromProvider,
	"authentik_source_saml":     config.IdentifierFromProvider,
	"authentik_source_scim":     config.IdentifierFromProvider,
	"authentik_source_telegram": config.IdentifierFromProvider,

	// Outposts and service connections.
	"authentik_outpost":                       config.IdentifierFromProvider,
	"authentik_outpost_provider_attachment":   config.IdentifierFromProvider,
	"authentik_service_connection_docker":     config.IdentifierFromProvider,
	"authentik_service_connection_kubernetes": config.IdentifierFromProvider,

	// RAC (Remote Access Control).
	"authentik_rac_endpoint": config.IdentifierFromProvider,

	// RBAC.
	"authentik_rbac_initial_permissions": config.IdentifierFromProvider,
	"authentik_rbac_permission_role":     config.IdentifierFromProvider,
	"authentik_rbac_permission_user":     config.IdentifierFromProvider,
	"authentik_rbac_role":                config.IdentifierFromProvider,

	// Events.
	"authentik_event_rule":      config.IdentifierFromProvider,
	"authentik_event_transport": config.IdentifierFromProvider,

	// Endpoints.
	"authentik_endpoints_connector_agent":                  config.IdentifierFromProvider,
	"authentik_endpoints_connector_agent_enrollment_token": config.IdentifierFromProvider,
	"authentik_endpoints_device_access_group":              config.IdentifierFromProvider,
	"authentik_endpoints_google_chrome_connector":          config.IdentifierFromProvider,
}

func idWithStub() config.ExternalName {
	e := config.IdentifierFromProvider
	e.GetExternalNameFn = func(tfstate map[string]any) (string, error) {
		en, _ := config.IDAsExternalName(tfstate)
		return en, nil
	}
	return e
}

// ExternalNameConfigurations applies all external name configs listed in the
// table ExternalNameConfigs and sets the version of those resources to v1beta1
// assuming they will be tested.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := ExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
			r.Version = "v1alpha1"
		}
	}
}

// ExternalNameConfigured returns the list of all resources whose external name
// is configured manually.
func ExternalNameConfigured() []string {
	l := make([]string, len(ExternalNameConfigs))
	i := 0
	for name := range ExternalNameConfigs {
		// $ is added to match the exact string since the format is regex.
		l[i] = name + "$"
		i++
	}
	return l
}

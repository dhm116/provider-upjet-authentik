package config

import (
	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
)

// Upjet derives an API group and kind for every resource from its Terraform
// name: the group is the second underscore-separated word and the kind is the
// camel-cased remainder, so authentik_provider_oauth2 becomes Oauth2 in the
// "provider" group. That default is right for the great majority of the
// authentik provider's resources, so only the exceptions are listed below.

// shortGroupOverrides maps a Terraform resource name to the API group it
// should land in, where the derived default would be misleading.
var shortGroupOverrides = map[string]string{
	// Two-word names default to the "authentik" root group, which would split
	// these away from the sub-resources they belong with.
	"authentik_application": "application", // joins Entitlement
	"authentik_flow":        "flow",        // joins StageBinding
	"authentik_outpost":     "outpost",     // joins ProviderAttachment

	// "service" is meaningless on its own; these are service connections.
	"authentik_service_connection_docker":     "serviceconnection",
	"authentik_service_connection_kubernetes": "serviceconnection",

	// Breaks a Go import cycle between the generated API packages. User has a
	// roles field referencing Role, and PermissionUser/PermissionRole in the
	// rbac group reference User, so authentik and rbac would import each
	// other. Role has no outgoing references, so moving it in beside User
	// leaves rbac -> authentik as the only edge between the two.
	"authentik_rbac_role": "authentik",
}

// kindOverrides maps a Terraform resource name to the kind it should be
// generated as, where the derived default reads poorly.
var kindOverrides = map[string]string{
	// Without these the kinds would be ConnectionDocker/ConnectionKubernetes,
	// stuttering against the serviceconnection group.
	"authentik_service_connection_docker":     "Docker",
	"authentik_service_connection_kubernetes": "Kubernetes",
}

// configureResources applies the group and kind overrides above. Resources
// needing more than that — references, sensitive fields, custom external
// names — should get their own configurator registered here.
func configureResources(pc *ujconfig.Provider) {
	for name, group := range shortGroupOverrides {
		pc.AddResourceConfigurator(name, func(r *ujconfig.Resource) {
			r.ShortGroup = group
		})
	}
	for name, kind := range kindOverrides {
		pc.AddResourceConfigurator(name, func(r *ujconfig.Resource) {
			r.Kind = kind
		})
	}
}

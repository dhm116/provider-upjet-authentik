package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
	"github.com/crossplane/upjet/v2/pkg/registry/reference"
)

// GetProviderNamespaced returns provider configuration for namespaced resources
func GetProviderNamespaced() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("authentik.m.goauthentik.io"),
		ujconfig.WithShortName("authentik"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		// Generates cross-resource references from the interpolations found in
		// the HCL examples scraped into provider-metadata.yaml.
		ujconfig.WithReferenceInjectors([]ujconfig.ReferenceInjector{reference.NewInjector(modulePath + "/apis/namespaced")}),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	configureResources(pc)

	pc.ConfigureResources()
	return pc
}

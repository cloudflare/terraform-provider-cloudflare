package resources

import (
	"log"

	"github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate/poc/interfaces"
	"github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate/poc/registry"
	"github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate/poc/resources/cloudflare_record"
	"github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate/poc/resources/cloudflare_zone_settings_override"
)

// ResourceFactory is a function that creates a new resource transformer
type ResourceFactory func() interfaces.ResourceTransformer

// ResourceFactories maps resource names to their factory functions
var ResourceFactories = map[string]ResourceFactory{
	"dns_record":             cloudflare_record.NewDNSRecord,
	"zone_settings_override": cloudflare_zone_settings_override.NewZoneSettingsOverride,
}

// RegisterFromFactories registers resource transformers with the registry
func RegisterFromFactories(reg *registry.StrategyRegistry, names ...string) {
	if len(names) == 0 {
		log.Println("Registering all available resources")
		for _, factory := range ResourceFactories {
			reg.Register(factory())
		}
		return
	}
	// Register only specified resources
	log.Println("Registering resources", names)
	for _, name := range names {
		if factory, ok := ResourceFactories[name]; ok {
			reg.Register(factory())
		}
	}
}

// RegisterAll is a convenience function that registers all available resources
func RegisterAll(reg *registry.StrategyRegistry) {
	RegisterFromFactories(reg)
}

// GetAvailableResources returns a list of all available resource names
func GetAvailableResources() []string {
	names := make([]string, 0, len(ResourceFactories))
	for name := range ResourceFactories {
		names = append(names, name)
	}
	return names
}

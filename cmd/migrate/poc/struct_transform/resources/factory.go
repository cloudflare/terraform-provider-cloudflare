package resources

import (
	"github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate/poc/interfaces"
	"github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate/poc/registry"
	"github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate/poc/struct_transform/resources/cloudflare_record"
)

// RegisterAllStructTransformers registers all struct-based resource transformers
// This is the single place to add new struct-based transformers
func RegisterAllStructTransformers(reg *registry.StrategyRegistry) {
	// DNS Record transformers
	reg.Register(cloudflare_record.NewDNSRecordStructTransformer())    // v4 -> v5
	reg.Register(cloudflare_record.NewDNSRecordV5StructTransformer())  // v5 passthrough

	// Add more struct-based transformers here as they are created
	// Example:
	// reg.Register(cloudflare_zone_settings.NewZoneSettingsStructTransformer())
	// reg.Register(cloudflare_load_balancer_pool.NewLoadBalancerPoolStructTransformer())
}

// CreateStructTransformer creates a struct-based transformer for a specific resource type
// This can be used for testing or when you need a specific transformer
func CreateStructTransformer(resourceType string) interfaces.ResourceTransformer {
	switch resourceType {
	case "cloudflare_record":
		return cloudflare_record.NewDNSRecordStructTransformer()
	case "cloudflare_dns_record":
		return cloudflare_record.NewDNSRecordV5StructTransformer()
	default:
		return nil
	}
}
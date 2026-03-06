package v500

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// SourceCloudflareFallbackDomainSchema returns the v4 schema for cloudflare_zero_trust_local_fallback_domain
// (or cloudflare_fallback_domain) from the SDKv2 provider.
//
// This schema is used by the Plugin Framework provider to correctly parse v4 state files
// during migration. It represents the state structure, not the full resource schema.
//
// Key differences from full v4 schema:
// - No Computed, Optional, Required flags (handled by state parsing)
// - No descriptions, validators, or plan modifiers
// - Only structure and types matter for state parsing
//
// This is the source schema for resources WITHOUT policy_id (default profile path).
func SourceCloudflareFallbackDomainSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{},
			"account_id": schema.StringAttribute{},
			"policy_id": schema.StringAttribute{}, // Was Optional in v4, must be null/absent for default profile
			"domains": schema.SetNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"suffix": schema.StringAttribute{},
						"description": schema.StringAttribute{},
						"dns_server": schema.ListAttribute{
							ElementType: types.StringType,
						},
					},
				},
			},
		},
	}
}

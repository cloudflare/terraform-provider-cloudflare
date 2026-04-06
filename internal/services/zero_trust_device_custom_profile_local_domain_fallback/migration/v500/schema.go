package v500

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// SourceCloudflareFallbackDomainSchema returns the source schema for legacy fallback domain resources.
// Schema version: 0 (SDKv2 default)
// Resource types:
//   - cloudflare_zero_trust_local_fallback_domain
//   - cloudflare_fallback_domain (deprecated alias)
//
// This minimal schema is used only for reading v4 state during migration to the custom profile.
// It includes only the properties needed for state parsing (Required, Optional, Computed, ElementType).
// Validators, PlanModifiers, and Descriptions are intentionally omitted.
//
// Note: This schema represents resources WITH policy_id (custom profile path).
func SourceCloudflareFallbackDomainSchema() schema.Schema {
	return schema.Schema{
		Version: 0, // SDKv2 default schema version
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true, // Format: "account_id/policy_id" for custom profile
			},
			"account_id": schema.StringAttribute{
				Required: true,
			},
			"domains": schema.SetNestedAttribute{
				Required: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"suffix": schema.StringAttribute{
							Optional: true, // Optional in v4, Required in v5
						},
						"description": schema.StringAttribute{
							Optional: true,
						},
						"dns_server": schema.ListAttribute{
							ElementType: types.StringType,
							Optional:    true,
						},
					},
				},
			},
			// policy_id must be present for custom profile migration path
			"policy_id": schema.StringAttribute{
				Optional: true, // Optional in v4, Required in v5
			},
		},
	}
}

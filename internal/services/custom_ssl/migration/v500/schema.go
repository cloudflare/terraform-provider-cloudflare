package v500

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// SourceCloudflareCustomSSLSchema returns the legacy cloudflare_custom_ssl schema (schema_version=1).
// This is used by UpgradeFromV4 to parse state from the legacy SDKv2 provider.
//
// Reference: cloudflare-terraform-v4/internal/sdkv2provider/schema_cloudflare_custom_ssl.go
// SchemaVersion: 1 (after the v0→v1 internal upgrader that converted custom_ssl_options from TypeMap to TypeList).
//
// Only Required, Optional, Computed, and ElementType are set. Validators, PlanModifiers,
// and Descriptions are intentionally omitted — this schema is only used for reading state.
func SourceCloudflareCustomSSLSchema() schema.Schema {
	return schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"zone_id": schema.StringAttribute{
				Required: true,
			},
			// custom_ssl_priority: TypeList (write-only reprioritization, not in v5)
			// Modeled as ListNestedAttribute so the state can be parsed and then dropped.
			"custom_ssl_priority": schema.ListNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Optional: true,
						},
						"priority": schema.Int64Attribute{
							Optional: true,
						},
					},
				},
			},
			// Computed top-level fields.
			"hosts": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
			},
			"issuer": schema.StringAttribute{
				Computed: true,
			},
			"signature": schema.StringAttribute{
				Computed: true,
			},
			"status": schema.StringAttribute{
				Computed: true,
			},
			"uploaded_on": schema.StringAttribute{
				Computed: true,
			},
			"modified_on": schema.StringAttribute{
				Computed: true,
			},
			"expires_on": schema.StringAttribute{
				Computed: true,
			},
			// priority: TypeInt in v4, becomes Float64 in v5.
			"priority": schema.Int64Attribute{
				Computed: true,
			},
		},
		Blocks: map[string]schema.Block{
			// custom_ssl_options: TypeList MaxItems:1 block in v4.
			// In v5, its sub-fields are flat top-level attributes.
			"custom_ssl_options": schema.ListNestedBlock{
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"certificate": schema.StringAttribute{
							Optional: true,
						},
						"private_key": schema.StringAttribute{
							Optional:  true,
							Sensitive: true,
						},
						"bundle_method": schema.StringAttribute{
							Optional: true,
						},
						// geo_restrictions: TypeString in v4, becomes SingleNestedAttribute in v5.
						"geo_restrictions": schema.StringAttribute{
							Optional: true,
						},
						"type": schema.StringAttribute{
							Optional: true,
						},
					},
				},
			},
		},
	}
}

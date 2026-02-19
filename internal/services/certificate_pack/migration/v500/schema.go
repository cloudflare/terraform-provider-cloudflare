package v500

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// SourceCloudflareCertificatePackSchema returns the legacy cloudflare_certificate_pack schema (schema_version=0).
//
// This minimal schema includes only the properties needed for state parsing:
// - Required, Optional, Computed flags
// - ElementType for collections
// - Nested structure definitions
//
// Validators, PlanModifiers, and Descriptions are intentionally omitted as they're not needed for reading existing state.
func SourceCloudflareCertificatePackSchema() schema.Schema {
	return schema.Schema{
		// Note: No Version field needed - the version is specified in the UpgradeState handler registration
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"zone_id": schema.StringAttribute{
				Required: true,
			},
			"type": schema.StringAttribute{
				Required: true,
			},
			"hosts": schema.SetAttribute{
				ElementType: types.StringType,
				Required:    true,
			},
			"validation_method": schema.StringAttribute{
				Required: true,
			},
			"validity_days": schema.Int64Attribute{
				Required: true,
			},
			"certificate_authority": schema.StringAttribute{
				Required: true,
			},
			"cloudflare_branding": schema.BoolAttribute{
				Optional: true,
			},
			// wait_for_active_status: removed in v5, but present in v4 state
			"wait_for_active_status": schema.BoolAttribute{
				Optional: true,
			},
		},
		Blocks: map[string]schema.Block{
			// validation_records: Becomes NestedObjectList in v5
			"validation_records": schema.ListNestedBlock{
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"emails": schema.ListAttribute{
							ElementType: types.StringType,
							Computed:    true,
						},
						"http_body": schema.StringAttribute{
							Computed: true,
						},
						"http_url": schema.StringAttribute{
							Computed: true,
						},
						"txt_name": schema.StringAttribute{
							Computed: true,
						},
						"txt_value": schema.StringAttribute{
							Computed: true,
						},
						// cname_target and cname_name: removed in v5, but present in v4 state
						"cname_target": schema.StringAttribute{
							Computed: true,
						},
						"cname_name": schema.StringAttribute{
							Computed: true,
						},
					},
				},
			},
			// validation_errors: Becomes NestedObjectList in v5
			"validation_errors": schema.ListNestedBlock{
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"message": schema.StringAttribute{
							Computed: true,
						},
					},
				},
			},
		},
	}
}

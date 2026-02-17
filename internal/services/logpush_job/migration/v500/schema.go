package v500

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// SourceCloudflareLogpushJobSchema returns the legacy cloudflare_logpush_job schema (schema_version=0).
// This is used by UpgradeFromLegacyV0 to parse state from the legacy SDKv2 provider.
// Reference: cloudflare-terraform-v4/internal/sdkv2provider/schema_cloudflare_logpush_job.go
//
// This minimal schema includes only the properties needed for state parsing:
// - Required, Optional, Computed flags
// - ElementType for collections
// - Nested structure definitions
// Validators, PlanModifiers, and Descriptions are intentionally omitted.
func SourceCloudflareLogpushJobSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"account_id": schema.StringAttribute{
				Optional: true,
			},
			"zone_id": schema.StringAttribute{
				Optional: true,
			},
			"dataset": schema.StringAttribute{
				Required: true,
			},
			"destination_conf": schema.StringAttribute{
				Required: true,
			},
			"enabled": schema.BoolAttribute{
				Optional: true,
			},
			"filter": schema.StringAttribute{
				Optional: true,
			},
			"frequency": schema.StringAttribute{
				Optional: true,
			},
			"kind": schema.StringAttribute{
				Optional: true,
			},
			"logpull_options": schema.StringAttribute{
				Optional: true,
			},
			"max_upload_bytes": schema.Int64Attribute{
				Optional: true,
			},
			"max_upload_interval_seconds": schema.Int64Attribute{
				Optional: true,
			},
			"max_upload_records": schema.Int64Attribute{
				Optional: true,
			},
			"name": schema.StringAttribute{
				Optional: true,
			},
			"ownership_challenge": schema.StringAttribute{
				Optional: true,
			},
			// output_options: v4 SDKv2 TypeList MaxItems:1 is stored as a JSON array [{...}]
			// TransformState is a no-op for this resource, so we receive the raw v4 array format.
			// Using ListNestedAttribute to correctly parse the array format from state.
			"output_options": schema.ListNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"batch_prefix": schema.StringAttribute{
							Optional: true,
						},
						"batch_suffix": schema.StringAttribute{
							Optional: true,
						},
						// Note: v4 field name is cve20214428, not cve_2021_44228
						"cve20214428": schema.BoolAttribute{
							Optional: true,
						},
						"field_delimiter": schema.StringAttribute{
							Optional: true,
						},
						"field_names": schema.ListAttribute{
							ElementType: types.StringType,
							Optional:    true,
						},
						"output_type": schema.StringAttribute{
							Optional: true,
						},
						"record_delimiter": schema.StringAttribute{
							Optional: true,
						},
						"record_prefix": schema.StringAttribute{
							Optional: true,
						},
						"record_suffix": schema.StringAttribute{
							Optional: true,
						},
						"record_template": schema.StringAttribute{
							Optional: true,
						},
						"sample_rate": schema.Float64Attribute{
							Optional: true,
						},
						"timestamp_format": schema.StringAttribute{
							Optional: true,
						},
					},
				},
			},
		},
	}
}

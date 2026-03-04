package v500

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// SourceCloudflareDLPProfileSchema returns the source schema for the legacy DLP profile resource.
// Schema version: 0 (SDKv2 default — neither cloudflare_dlp_profile nor cloudflare_zero_trust_dlp_profile set an explicit SchemaVersion)
// Resource types: cloudflare_dlp_profile, cloudflare_zero_trust_dlp_profile
//
// This minimal schema is used only for reading v4 state during migration.
// Validators, PlanModifiers, and Descriptions are intentionally omitted.
func SourceCloudflareDLPProfileSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"account_id": schema.StringAttribute{
				Required: true,
			},
			"name": schema.StringAttribute{
				Required: true,
			},
			"description": schema.StringAttribute{
				Optional: true,
			},
			"type": schema.StringAttribute{
				Required: true,
			},
			"entry": schema.ListNestedAttribute{
				Required: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Optional: true,
							Computed: true,
						},
						"name": schema.StringAttribute{
							Required: true,
						},
						"enabled": schema.BoolAttribute{
							Optional: true,
						},
						"pattern": schema.ListNestedAttribute{
							Optional: true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"regex": schema.StringAttribute{
										Required: true,
									},
									"validation": schema.StringAttribute{
										Optional: true,
									},
								},
							},
						},
					},
				},
			},
			"allowed_match_count": schema.Int64Attribute{
				Required: true,
			},
			"context_awareness": schema.ListNestedAttribute{
				Optional: true,
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"enabled": schema.BoolAttribute{
							Required: true,
						},
						"skip": schema.ListNestedAttribute{
							Required: true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"files": schema.BoolAttribute{
										Required: true,
									},
								},
							},
						},
					},
				},
			},
			"ocr_enabled": schema.BoolAttribute{
				Optional: true,
			},
		},
	}
}

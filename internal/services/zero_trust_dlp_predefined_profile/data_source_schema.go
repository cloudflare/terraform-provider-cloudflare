// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_dlp_predefined_profile

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*ZeroTrustDLPPredefinedProfileDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Required: true,
			},
			"profile_id": schema.StringAttribute{
				Required: true,
			},
			"ai_context_enabled": schema.BoolAttribute{
				Computed: true,
			},
			"allowed_match_count": schema.Int64Attribute{
				Description: "Related DLP policies will trigger when the match count exceeds the number set.",
				Computed:    true,
			},
			"confidence_threshold": schema.StringAttribute{
				Description: `Available values: "low", "medium", "high", "very_high".`,
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"low",
						"medium",
						"high",
						"very_high",
					),
				},
			},
			"created_at": schema.StringAttribute{
				Description: "When the profile was created",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"description": schema.StringAttribute{
				Description: "The description of the profile",
				Computed:    true,
			},
			"id": schema.StringAttribute{
				Description: "The id of the profile (uuid)",
				Computed:    true,
			},
			"name": schema.StringAttribute{
				Description: "The name of the profile",
				Computed:    true,
			},
			"ocr_enabled": schema.BoolAttribute{
				Computed: true,
			},
			"open_access": schema.BoolAttribute{
				Description: "Whether this profile can be accessed by anyone",
				Computed:    true,
			},
			"type": schema.StringAttribute{
				Description: `Available values: "custom".`,
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"custom",
						"predefined",
						"integration",
					),
				},
			},
			"updated_at": schema.StringAttribute{
				Description: "When the profile was lasted updated",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"context_awareness": schema.SingleNestedAttribute{
				Description: "Scan the context of predefined entries to only return matches surrounded by keywords.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[ZeroTrustDLPPredefinedProfileContextAwarenessDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"enabled": schema.BoolAttribute{
						Description: "If true, scan the context of predefined entries to only return matches surrounded by keywords.",
						Computed:    true,
					},
					"skip": schema.SingleNestedAttribute{
						Description: "Content types to exclude from context analysis and return all matches.",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectType[ZeroTrustDLPPredefinedProfileContextAwarenessSkipDataSourceModel](ctx),
						Attributes: map[string]schema.Attribute{
							"files": schema.BoolAttribute{
								Description: "If the content type is a file, skip context analysis and return all matches.",
								Computed:    true,
							},
						},
					},
				},
			},
			"entries": schema.ListNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectListType[ZeroTrustDLPPredefinedProfileEntriesDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Computed: true,
						},
						"created_at": schema.StringAttribute{
							Computed:   true,
							CustomType: timetypes.RFC3339Type{},
						},
						"enabled": schema.BoolAttribute{
							Computed: true,
						},
						"name": schema.StringAttribute{
							Computed: true,
						},
						"pattern": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustDLPPredefinedProfileEntriesPatternDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"regex": schema.StringAttribute{
									Computed: true,
								},
								"validation": schema.StringAttribute{
									Description: `Available values: "luhn".`,
									Computed:    true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive("luhn"),
									},
								},
							},
						},
						"type": schema.StringAttribute{
							Description: `Available values: "custom".`,
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive(
									"custom",
									"predefined",
									"integration",
									"exact_data",
									"word_list",
								),
							},
						},
						"updated_at": schema.StringAttribute{
							Computed:   true,
							CustomType: timetypes.RFC3339Type{},
						},
						"profile_id": schema.StringAttribute{
							Computed: true,
						},
						"confidence": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustDLPPredefinedProfileEntriesConfidenceDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"ai_context_available": schema.BoolAttribute{
									Description: "Indicates whether this entry has AI remote service validation",
									Computed:    true,
								},
								"available": schema.BoolAttribute{
									Description: "Indicates whether this entry has any form of validation that is not an AI remote service",
									Computed:    true,
								},
							},
						},
						"secret": schema.BoolAttribute{
							Computed: true,
						},
						"word_list": schema.StringAttribute{
							Computed:   true,
							CustomType: jsontypes.NormalizedType{},
						},
					},
				},
			},
		},
	}
}

func (d *ZeroTrustDLPPredefinedProfileDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *ZeroTrustDLPPredefinedProfileDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}

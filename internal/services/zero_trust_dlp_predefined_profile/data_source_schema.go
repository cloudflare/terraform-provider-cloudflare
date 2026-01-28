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
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = (*ZeroTrustDLPPredefinedProfileDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"profile_id": schema.StringAttribute{
				Required: true,
			},
			"account_id": schema.StringAttribute{
				Required: true,
			},
			"ai_context_enabled": schema.BoolAttribute{
				Computed: true,
			},
			"allowed_match_count": schema.Int64Attribute{
				Computed: true,
			},
			"confidence_threshold": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Description: "The name of the predefined profile.",
				Computed:    true,
			},
			"ocr_enabled": schema.BoolAttribute{
				Computed: true,
			},
			"open_access": schema.BoolAttribute{
				Description: "Whether this profile can be accessed by anyone.",
				Computed:    true,
			},
			"enabled_entries": schema.ListAttribute{
				Description: "Entries to enable for this predefined profile. Any entries not provided will be disabled.",
				Computed:    true,
				CustomType:  customfield.NewListType[types.String](ctx),
				ElementType: types.StringType,
			},
			"entries": schema.ListNestedAttribute{
				Description:        "This field has been deprecated for `enabled_entries`.",
				Computed:           true,
				DeprecationMessage: "This attribute is deprecated.",
				CustomType:         customfield.NewNestedObjectListType[ZeroTrustDLPPredefinedProfileEntriesDataSourceModel](ctx),
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
									Description:        `Available values: "luhn".`,
									Computed:           true,
									DeprecationMessage: "This attribute is deprecated.",
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive("luhn"),
									},
								},
							},
						},
						"type": schema.StringAttribute{
							Description: `Available values: "custom", "predefined", "integration", "exact_data", "document_fingerprint", "word_list".`,
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive(
									"custom",
									"predefined",
									"integration",
									"exact_data",
									"document_fingerprint",
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
									Description: "Indicates whether this entry has AI remote service validation.",
									Computed:    true,
								},
								"available": schema.BoolAttribute{
									Description: "Indicates whether this entry has any form of validation that is not an AI remote service.",
									Computed:    true,
								},
							},
						},
						"variant": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustDLPPredefinedProfileEntriesVariantDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"topic_type": schema.StringAttribute{
									Description: `Available values: "Intent", "Content".`,
									Computed:    true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive("Intent", "Content"),
									},
								},
								"type": schema.StringAttribute{
									Description: `Available values: "PromptTopic".`,
									Computed:    true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive("PromptTopic"),
									},
								},
								"description": schema.StringAttribute{
									Computed: true,
								},
							},
						},
						"case_sensitive": schema.BoolAttribute{
							Description: "Only applies to custom word lists.\nDetermines if the words should be matched in a case-sensitive manner\nCannot be set to false if secret is true",
							Computed:    true,
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

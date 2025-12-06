// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_dlp_integration_entry

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ resource.ResourceWithConfigValidators = (*ZeroTrustDLPIntegrationEntryResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"account_id": schema.StringAttribute{
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"entry_id": schema.StringAttribute{
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"profile_id": schema.StringAttribute{
				Description:   "This field is not used as the owning profile.\nFor predefined entries it is already set to a predefined profile.",
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"enabled": schema.BoolAttribute{
				Required: true,
			},
			"case_sensitive": schema.BoolAttribute{
				Description: "Only applies to custom word lists.\nDetermines if the words should be matched in a case-sensitive manner\nCannot be set to false if secret is true",
				Computed:    true,
			},
			"created_at": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"name": schema.StringAttribute{
				Computed: true,
			},
			"secret": schema.BoolAttribute{
				Computed: true,
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
			"confidence": schema.SingleNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectType[ZeroTrustDLPIntegrationEntryConfidenceModel](ctx),
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
			"pattern": schema.SingleNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectType[ZeroTrustDLPIntegrationEntryPatternModel](ctx),
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
			"profiles": schema.ListNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectListType[ZeroTrustDLPIntegrationEntryProfilesModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Computed: true,
						},
						"name": schema.StringAttribute{
							Computed: true,
						},
					},
				},
			},
			"variant": schema.SingleNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectType[ZeroTrustDLPIntegrationEntryVariantModel](ctx),
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
			"word_list": schema.StringAttribute{
				Computed:   true,
				CustomType: jsontypes.NormalizedType{},
			},
		},
	}
}

func (r *ZeroTrustDLPIntegrationEntryResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *ZeroTrustDLPIntegrationEntryResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}

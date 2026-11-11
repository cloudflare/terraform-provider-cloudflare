// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_dlp_proximity_validator

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/schemata"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ resource.ResourceWithConfigValidators = (*ZeroTrustDLPProximityValidatorResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		MarkdownDescription: schemata.Description{
			Scopes: []string{
				"Zero Trust Read",
				"Zero Trust Write",
			},
		}.String(),
		Version: 500,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseNonNullStateForUnknown()},
			},
			"account_id": schema.StringAttribute{
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"name": schema.StringAttribute{
				Description: "A name for this validator, at most 50 characters.",
				Required:    true,
			},
			"words": schema.SetNestedAttribute{
				Description: "The words that raise confidence when found near a match, at most 100 per validator. Matching\nignores case, so the list may not contain the same word twice with different capitalisation.\nWords are stored as sent.\n\nFor example, a validator for medical record numbers might use\n`[{\"word\": \"patient\", \"weight\": \"strong\"}, {\"word\": \"diagnosis\", \"weight\": \"moderate\"},\n{\"word\": \"mrn\", \"weight\": \"weak\"}]`.",
				Required:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"weight": schema.StringAttribute{
							Description: "How much a word contributes to a match's confidence when it is found nearby.\n\nSeparate from a profile's high/medium/low confidence threshold: this describes how telling the\nword is, not how confident the resulting match is.\nAvailable values: \"strong\", \"moderate\", \"weak\".",
							Required:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive(
									"strong",
									"moderate",
									"weak",
								),
							},
						},
						"word": schema.StringAttribute{
							Description: "The word or phrase to look for. Must contain 1–128 characters and is matched\ncase-insensitively.",
							Required:    true,
						},
					},
				},
			},
			"created_at": schema.StringAttribute{
				Computed:      true,
				CustomType:    timetypes.RFC3339Type{},
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseNonNullStateForUnknown()},
			},
			"updated_at": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"referenced_by": schema.SingleNestedAttribute{
				Description: "What currently uses a validator. Deleting a validator that is still in use returns `409`, so\ncheck this first.\n\nRead-only. A validator is attached or detached by updating the detection entry that uses it, not\nby writing this back.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[ZeroTrustDLPProximityValidatorReferencedByModel](ctx),
				Attributes: map[string]schema.Attribute{
					"entries": schema.ListNestedAttribute{
						Description: "The detection entries using this validator, ordered by name.",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectListType[ZeroTrustDLPProximityValidatorReferencedByEntriesModel](ctx),
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									Computed: true,
								},
								"name": schema.StringAttribute{
									Computed: true,
								},
								"type": schema.StringAttribute{
									Description: "The type of a detection entry.\nAvailable values: \"custom\", \"custom_prompt_topic\", \"predefined\", \"integration\", \"exact_data\", \"document_fingerprint\", \"word_list\".",
									Computed:    true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive(
											"custom",
											"custom_prompt_topic",
											"predefined",
											"integration",
											"exact_data",
											"document_fingerprint",
											"word_list",
										),
									},
								},
							},
						},
					},
					"entry_count": schema.Int64Attribute{
						Description: "Number of detection entries this validator is attached to.",
						Computed:    true,
						Validators: []validator.Int64{
							int64validator.AtLeast(0),
						},
					},
				},
			},
		},
	}
}

func (r *ZeroTrustDLPProximityValidatorResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *ZeroTrustDLPProximityValidatorResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}

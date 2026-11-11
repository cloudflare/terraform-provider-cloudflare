// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_casb_finding_type

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/schemata"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithConfigValidators = (*ZeroTrustCasbFindingTypeResource)(nil)

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
				Description:   "The unique identifier of the finding.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseNonNullStateForUnknown()},
			},
			"account_id": schema.StringAttribute{
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"asset_class": schema.StringAttribute{
				Description:   "Asset class this finding type applies to (PascalCase, e.g. GoogleWorkspaceFile)",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"asset_vendor": schema.StringAttribute{
				Description: "Canonical vendor enum for the platform whose assets this finding type evaluates.\nAvailable values: \"ANTHROPIC\", \"AWS\", \"BITBUCKET\", \"BOX\", \"CONFLUENCE\", \"DROPBOX\", \"GITHUB\", \"GOOGLE_CLOUD_PLATFORM\", \"GOOGLE_WORKSPACE\", \"JIRA\", \"MICROSOFT\", \"OPENAI\", \"SALESFORCE\", \"SERVICENOW\", \"SLACK\", \"ZOOM\".",
				Required:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"ANTHROPIC",
						"AWS",
						"BITBUCKET",
						"BOX",
						"CONFLUENCE",
						"DROPBOX",
						"GITHUB",
						"GOOGLE_CLOUD_PLATFORM",
						"GOOGLE_WORKSPACE",
						"JIRA",
						"MICROSOFT",
						"OPENAI",
						"SALESFORCE",
						"SERVICENOW",
						"SLACK",
						"ZOOM",
					),
				},
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"expression": schema.StringAttribute{
				Description:   "Rego expression body for the policy. `package cloudflare.casb.custom` and `import rego.v1` are added automatically. Do not include package or import declarations. This field cannot be changed after creation. Maximum 4096 characters.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"applies_to_all_integrations": schema.BoolAttribute{
				Description: "When true, this finding type applies to every current and future integration of the vendor.",
				Required:    true,
			},
			"description": schema.StringAttribute{
				Description: "Human-readable description of the finding type. Cannot be empty. Maximum 4096 characters.",
				Required:    true,
			},
			"name": schema.StringAttribute{
				Description: "Human-readable name for the finding type. Maximum 256 characters.",
				Required:    true,
			},
			"severity": schema.StringAttribute{
				Description: "Default severity used when no integration-specific severity override exists.\nAvailable values: \"Critical\", \"High\", \"Medium\", \"Low\".",
				Required:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"Critical",
						"High",
						"Medium",
						"Low",
					),
				},
			},
			"integration_ids": schema.ListAttribute{
				Description:   "Required and non-empty when applies_to_all_integrations is false; omit when it is true.",
				Computed:      true,
				Optional:      true,
				CustomType:    customfield.NewListType[types.String](ctx),
				ElementType:   types.StringType,
				PlanModifiers: []planmodifier.List{listplanmodifier.UseNonNullStateForUnknown()},
			},
			"created_at": schema.StringAttribute{
				Description:   "Timestamp when the finding type was created.",
				Computed:      true,
				CustomType:    timetypes.RFC3339Type{},
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseNonNullStateForUnknown()},
			},
			"match_count": schema.Int64Attribute{
				Description: "Number of currently active finding instances across the account's integrations. Zero when none have been detected.",
				Computed:    true,
			},
			"origin": schema.StringAttribute{
				Description: "Whether the finding type is system-managed or customer-created.\nAvailable values: \"system\", \"custom\".",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("system", "custom"),
				},
			},
			"quarantined": schema.BoolAttribute{
				Description: "Whether evaluation of this finding type is suspended by an operator. Present only when true.",
				Computed:    true,
			},
			"updated_at": schema.StringAttribute{
				Description: "Timestamp when the finding type was last updated.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"vendor": schema.StringAttribute{
				Description: "The SaaS/Cloud vendor of the platform with which the finding is associated.",
				Computed:    true,
			},
			"version": schema.Int64Attribute{
				Description: "Current version of the finding type. Incremented on each update.",
				Computed:    true,
			},
			"category": schema.SingleNestedAttribute{
				Description: "Category information for a finding.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[ZeroTrustCasbFindingTypeCategoryModel](ctx),
				Attributes: map[string]schema.Attribute{
					"observation": schema.StringAttribute{
						Description: "The type of the observation.\nAvailable values: \"Issue\", \"Insight\", \"Activity\".",
						Computed:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive(
								"Issue",
								"Insight",
								"Activity",
							),
						},
					},
					"product": schema.StringAttribute{
						Description: "The product category.\nAvailable values: \"SaaS\", \"Cloud\".",
						Computed:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("SaaS", "Cloud"),
						},
					},
					"type": schema.StringAttribute{
						Description: "The type of the finding category.\nAvailable values: \"Content\", \"Posture\".",
						Computed:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("Content", "Posture"),
						},
					},
				},
			},
		},
	}
}

func (r *ZeroTrustCasbFindingTypeResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *ZeroTrustCasbFindingTypeResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}

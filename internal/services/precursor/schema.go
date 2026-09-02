// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package precursor

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ resource.ResourceWithConfigValidators = (*PrecursorResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Version: 500,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "Identifier.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseNonNullStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"zone_id": schema.StringAttribute{
				Description:   "Identifier.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseNonNullStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"default_mode": schema.StringAttribute{
				Description: "The zone-level Precursor enforcement mode applied to requests that do\nnot match a more specific enforcement rule.\nAvailable values: \"off\", \"min-friction\", \"max-security\".",
				Computed:    true,
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"off",
						"min-friction",
						"max-security",
					),
				},
				Default: stringdefault.StaticString("off"),
			},
			"enforcement_rules": schema.ListNestedAttribute{
				Description: "The ordered list of enforcement rules for the zone.",
				Computed:    true,
				Optional:    true,
				CustomType:  customfield.NewNestedObjectListType[PrecursorEnforcementRulesModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"expression": schema.StringAttribute{
							Description: "The filter expression that determines which requests the rule matches.",
							Required:    true,
						},
						"mode": schema.StringAttribute{
							Description: "The override mode Precursor applies to requests matching an enforcement\nrule. Unlike `default_mode`, this cannot be `off`.\nAvailable values: \"min-friction\", \"max-security\".",
							Required:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive("min-friction", "max-security"),
							},
						},
						"id": schema.StringAttribute{
							Description: "The read-only identifier that Cloudflare assigns to the rule.",
							Computed:    true,
						},
						"description": schema.StringAttribute{
							Description: "An informative description of the rule.",
							Computed:    true,
							Optional:    true,
							Default:     stringdefault.StaticString(""),
						},
						"enabled": schema.BoolAttribute{
							Description: "Whether the rule is active.",
							Computed:    true,
							Optional:    true,
							Default:     booldefault.StaticBool(true),
						},
					},
				},
			},
		},
	}
}

func (r *PrecursorResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *PrecursorResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}

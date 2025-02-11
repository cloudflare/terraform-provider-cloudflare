// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_routing_rule

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/float64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithConfigValidators = (*EmailRoutingRuleResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"zone_id": schema.StringAttribute{
				Description:   "Identifier",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"rule_identifier": schema.StringAttribute{
				Description:   "Routing rule identifier.",
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"actions": schema.ListNestedAttribute{
				Description: "List actions patterns.",
				Required:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							Description: "Type of supported action.",
							Required:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive(
									"drop",
									"forward",
									"worker",
								),
							},
						},
						"value": schema.ListAttribute{
							Required:    true,
							ElementType: types.StringType,
						},
					},
				},
			},
			"matchers": schema.ListNestedAttribute{
				Description: "Matching patterns to forward to your actions.",
				Required:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"field": schema.StringAttribute{
							Description: "Field for type matcher.",
							Required:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive("to"),
							},
						},
						"type": schema.StringAttribute{
							Description: "Type of matcher.",
							Required:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive("literal"),
							},
						},
						"value": schema.StringAttribute{
							Description: "Value for matcher.",
							Required:    true,
						},
					},
				},
			},
			"name": schema.StringAttribute{
				Description: "Routing rule name.",
				Optional:    true,
			},
			"enabled": schema.BoolAttribute{
				Description: "Routing rule status.",
				Computed:    true,
				Optional:    true,
				Default:     booldefault.StaticBool(true),
			},
			"priority": schema.Float64Attribute{
				Description: "Priority of the routing rule.",
				Computed:    true,
				Optional:    true,
				Validators: []validator.Float64{
					float64validator.AtLeast(0),
				},
				Default: float64default.StaticFloat64(0),
			},
		},
	}
}

func (r *EmailRoutingRuleResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *EmailRoutingRuleResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}

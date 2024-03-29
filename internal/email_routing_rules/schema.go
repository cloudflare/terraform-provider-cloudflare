// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_routing_rules

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r EmailRoutingRulesResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"zone_identifier": schema.StringAttribute{
				Description: "Identifier",
				Required:    true,
			},
			"rule_identifier": schema.StringAttribute{
				Description: "Routing rule identifier.",
				Optional:    true,
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
								stringvalidator.OneOfCaseInsensitive("drop", "forward", "worker"),
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
			"enabled": schema.BoolAttribute{
				Description: "Routing rule status.",
				Optional:    true,
			},
			"name": schema.StringAttribute{
				Description: "Routing rule name.",
				Optional:    true,
			},
			"priority": schema.Float64Attribute{
				Description: "Priority of the routing rule.",
				Optional:    true,
			},
			"id": schema.StringAttribute{
				Description: "Routing rule identifier.",
				Computed:    true,
			},
			"tag": schema.StringAttribute{
				Description: "Routing rule tag. (Deprecated, replaced by routing rule identifier)",
				Computed:    true,
			},
		},
	}
}

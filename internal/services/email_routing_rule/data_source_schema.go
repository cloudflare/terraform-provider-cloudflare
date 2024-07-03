// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_routing_rule

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = &EmailRoutingRuleDataSource{}
var _ datasource.DataSourceWithValidateConfig = &EmailRoutingRuleDataSource{}

func (r EmailRoutingRuleDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"zone_identifier": schema.StringAttribute{
				Description: "Identifier",
				Optional:    true,
			},
			"rule_identifier": schema.StringAttribute{
				Description: "Routing rule identifier.",
				Optional:    true,
			},
			"id": schema.StringAttribute{
				Description: "Routing rule identifier.",
				Optional:    true,
			},
			"actions": schema.ListNestedAttribute{
				Description: "List actions patterns.",
				Optional:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							Description: "Type of supported action.",
							Required:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive("drop", "forward", "worker"),
							},
						},
						"value": schema.StringAttribute{
							Required: true,
						},
					},
				},
			},
			"enabled": schema.BoolAttribute{
				Description: "Routing rule status.",
				Computed:    true,
				Optional:    true,
			},
			"matchers": schema.ListNestedAttribute{
				Description: "Matching patterns to forward to your actions.",
				Optional:    true,
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
			"priority": schema.Float64Attribute{
				Description: "Priority of the routing rule.",
				Computed:    true,
				Optional:    true,
			},
			"tag": schema.StringAttribute{
				Description: "Routing rule tag. (Deprecated, replaced by routing rule identifier)",
				Optional:    true,
			},
			"find_one_by": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"zone_identifier": schema.StringAttribute{
						Description: "Identifier",
						Required:    true,
					},
					"enabled": schema.BoolAttribute{
						Description: "Filter by enabled routing rules.",
						Optional:    true,
					},
					"page": schema.Float64Attribute{
						Description: "Page number of paginated results.",
						Computed:    true,
						Optional:    true,
					},
					"per_page": schema.Float64Attribute{
						Description: "Maximum number of results per page.",
						Computed:    true,
						Optional:    true,
					},
				},
			},
		},
	}
}

func (r *EmailRoutingRuleDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}

func (r *EmailRoutingRuleDataSource) ValidateConfig(ctx context.Context, req datasource.ValidateConfigRequest, resp *datasource.ValidateConfigResponse) {
}

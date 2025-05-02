// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_routing_rule

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = (*EmailRoutingRuleDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Routing rule identifier.",
				Computed:    true,
			},
			"rule_identifier": schema.StringAttribute{
				Description: "Routing rule identifier.",
				Optional:    true,
			},
			"zone_id": schema.StringAttribute{
				Description: "Identifier.",
				Required:    true,
			},
			"enabled": schema.BoolAttribute{
				Description: "Routing rule status.",
				Computed:    true,
			},
			"name": schema.StringAttribute{
				Description: "Routing rule name.",
				Computed:    true,
			},
			"priority": schema.Float64Attribute{
				Description: "Priority of the routing rule.",
				Computed:    true,
				Validators: []validator.Float64{
					float64validator.AtLeast(0),
				},
			},
			"tag": schema.StringAttribute{
				Description:        "Routing rule tag. (Deprecated, replaced by routing rule identifier)",
				Computed:           true,
				DeprecationMessage: "This attribute is deprecated.",
			},
			"actions": schema.ListNestedAttribute{
				Description: "List actions patterns.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[EmailRoutingRuleActionsDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							Description: "Type of supported action.\nAvailable values: \"drop\", \"forward\", \"worker\".",
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive(
									"drop",
									"forward",
									"worker",
								),
							},
						},
						"value": schema.ListAttribute{
							Computed:    true,
							CustomType:  customfield.NewListType[types.String](ctx),
							ElementType: types.StringType,
						},
					},
				},
			},
			"matchers": schema.ListNestedAttribute{
				Description: "Matching patterns to forward to your actions.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[EmailRoutingRuleMatchersDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							Description: "Type of matcher.\nAvailable values: \"literal\".",
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive("literal"),
							},
						},
						"field": schema.StringAttribute{
							Description: "Field for type matcher.\nAvailable values: \"to\".",
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive("to"),
							},
						},
						"value": schema.StringAttribute{
							Description: "Value for matcher.",
							Computed:    true,
						},
					},
				},
			},
			"filter": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"enabled": schema.BoolAttribute{
						Description: "Filter by enabled routing rules.",
						Optional:    true,
					},
				},
			},
		},
	}
}

func (d *EmailRoutingRuleDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *EmailRoutingRuleDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("rule_identifier"), path.MatchRoot("filter")),
	}
}

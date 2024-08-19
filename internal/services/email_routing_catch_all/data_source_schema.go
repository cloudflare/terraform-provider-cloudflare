// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_routing_catch_all

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = &EmailRoutingCatchAllDataSource{}

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"zone_identifier": schema.StringAttribute{
				Description: "Identifier",
				Required:    true,
			},
			"id": schema.StringAttribute{
				Description: "Routing rule identifier.",
				Optional:    true,
			},
			"name": schema.StringAttribute{
				Description: "Routing rule name.",
				Optional:    true,
			},
			"tag": schema.StringAttribute{
				Description: "Routing rule tag. (Deprecated, replaced by routing rule identifier)",
				Optional:    true,
			},
			"actions": schema.ListNestedAttribute{
				Description: "List actions for the catch-all routing rule.",
				Optional:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							Description: "Type of action for catch-all rule.",
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
							Optional:    true,
							ElementType: types.StringType,
						},
					},
				},
			},
			"matchers": schema.ListNestedAttribute{
				Description: "List of matchers for the catch-all routing rule.",
				Optional:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							Description: "Type of matcher. Default is 'all'.",
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive("all"),
							},
						},
					},
				},
			},
			"enabled": schema.BoolAttribute{
				Description: "Routing rule status.",
				Computed:    true,
				Optional:    true,
			},
		},
	}
}

func (d *EmailRoutingCatchAllDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *EmailRoutingCatchAllDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}

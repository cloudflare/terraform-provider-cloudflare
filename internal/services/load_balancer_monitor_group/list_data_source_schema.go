// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package load_balancer_monitor_group

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*LoadBalancerMonitorGroupsDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "Identifier.",
				Required:    true,
			},
			"max_items": schema.Int64Attribute{
				Description: "Max items to fetch, default: 1000",
				Optional:    true,
				Validators: []validator.Int64{
					int64validator.AtLeast(0),
				},
			},
			"result": schema.ListNestedAttribute{
				Description: "The items returned by the data source",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[LoadBalancerMonitorGroupsResultDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "The ID of the Monitor Group to use for checking the health of origins within this pool.",
							Computed:    true,
						},
						"description": schema.StringAttribute{
							Description: "A short description of the monitor group",
							Computed:    true,
						},
						"members": schema.ListNestedAttribute{
							Description: "List of monitors in this group",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectListType[LoadBalancerMonitorGroupsMembersDataSourceModel](ctx),
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"enabled": schema.BoolAttribute{
										Description: "Whether this monitor is enabled in the group",
										Computed:    true,
									},
									"monitor_id": schema.StringAttribute{
										Description: "The ID of the Monitor to use for checking the health of origins within this pool.",
										Computed:    true,
									},
									"monitoring_only": schema.BoolAttribute{
										Description: "Whether this monitor is used for monitoring only (does not affect pool health)",
										Computed:    true,
									},
									"must_be_healthy": schema.BoolAttribute{
										Description: "Whether this monitor must be healthy for the pool to be considered healthy",
										Computed:    true,
									},
									"created_at": schema.StringAttribute{
										Description: "The timestamp of when the monitor was added to the group",
										Computed:    true,
										CustomType:  timetypes.RFC3339Type{},
									},
									"updated_at": schema.StringAttribute{
										Description: "The timestamp of when the monitor group member was last updated",
										Computed:    true,
										CustomType:  timetypes.RFC3339Type{},
									},
								},
							},
						},
						"created_at": schema.StringAttribute{
							Description: "The timestamp of when the monitor group was created",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"updated_at": schema.StringAttribute{
							Description: "The timestamp of when the monitor group was last updated",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
					},
				},
			},
		},
	}
}

func (d *LoadBalancerMonitorGroupsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *LoadBalancerMonitorGroupsDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}

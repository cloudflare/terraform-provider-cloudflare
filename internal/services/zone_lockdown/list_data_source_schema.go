// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zone_lockdown

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = (*ZoneLockdownsDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"zone_id": schema.StringAttribute{
				Description: "Identifier",
				Required:    true,
			},
			"created_on": schema.StringAttribute{
				Description: "The timestamp of when the rule was created.",
				Optional:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"description": schema.StringAttribute{
				Description: "A string to search for in the description of existing rules.",
				Optional:    true,
			},
			"description_search": schema.StringAttribute{
				Description: "A string to search for in the description of existing rules.",
				Optional:    true,
			},
			"ip": schema.StringAttribute{
				Description: "A single IP address to search for in existing rules.",
				Optional:    true,
			},
			"ip_range_search": schema.StringAttribute{
				Description: "A single IP address range to search for in existing rules.",
				Optional:    true,
			},
			"ip_search": schema.StringAttribute{
				Description: "A single IP address to search for in existing rules.",
				Optional:    true,
			},
			"modified_on": schema.StringAttribute{
				Description: "The timestamp of when the rule was last modified.",
				Optional:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"priority": schema.Float64Attribute{
				Description: "The priority of the rule to control the processing order. A lower number indicates higher priority. If not provided, any rules with a configured priority will be processed before rules without a priority.",
				Optional:    true,
			},
			"uri_search": schema.StringAttribute{
				Description: "A single URI to search for in the list of URLs of existing rules.",
				Optional:    true,
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
				CustomType:  customfield.NewNestedObjectListType[ZoneLockdownsResultDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "The unique identifier of the Zone Lockdown rule.",
							Computed:    true,
						},
						"configurations": schema.ListNestedAttribute{
							Description: "A list of IP addresses or CIDR ranges that will be allowed to access the URLs specified in the Zone Lockdown rule. You can include any number of `ip` or `ip_range` configurations.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectListType[ZoneLockdownsConfigurationsDataSourceModel](ctx),
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"target": schema.StringAttribute{
										Description: "The configuration target. You must set the target to `ip` when specifying an IP address in the Zone Lockdown rule.",
										Computed:    true,
										Validators: []validator.String{
											stringvalidator.OneOfCaseInsensitive("ip", "ip_range"),
										},
									},
									"value": schema.StringAttribute{
										Description: "The IP address to match. This address will be compared to the IP address of incoming requests.",
										Computed:    true,
									},
								},
							},
						},
						"created_on": schema.StringAttribute{
							Description: "The timestamp of when the rule was created.",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"description": schema.StringAttribute{
							Description: "An informative summary of the rule.",
							Computed:    true,
						},
						"modified_on": schema.StringAttribute{
							Description: "The timestamp of when the rule was last modified.",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"paused": schema.BoolAttribute{
							Description: "When true, indicates that the rule is currently paused.",
							Computed:    true,
						},
						"urls": schema.ListAttribute{
							Description: "The URLs to include in the rule definition. You can use wildcards. Each entered URL will be escaped before use, which means you can only use simple wildcard patterns.",
							Computed:    true,
							CustomType:  customfield.NewListType[types.String](ctx),
							ElementType: types.StringType,
						},
					},
				},
			},
		},
	}
}

func (d *ZoneLockdownsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *ZoneLockdownsDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}

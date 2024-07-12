// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zone_lockdown

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = &ZoneLockdownDataSource{}
var _ datasource.DataSourceWithValidateConfig = &ZoneLockdownDataSource{}

func (r ZoneLockdownDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"zone_identifier": schema.StringAttribute{
				Description: "Identifier",
				Optional:    true,
			},
			"id": schema.StringAttribute{
				Description: "The unique identifier of the Zone Lockdown rule.",
				Computed:    true,
				Optional:    true,
			},
			"created_on": schema.StringAttribute{
				Description: "The timestamp of when the rule was created.",
				Computed:    true,
			},
			"description": schema.StringAttribute{
				Description: "An informative summary of the rule.",
				Computed:    true,
			},
			"modified_on": schema.StringAttribute{
				Description: "The timestamp of when the rule was last modified.",
				Computed:    true,
			},
			"paused": schema.BoolAttribute{
				Description: "When true, indicates that the rule is currently paused.",
				Computed:    true,
			},
			"urls": schema.ListAttribute{
				Description: "The URLs to include in the rule definition. You can use wildcards. Each entered URL will be escaped before use, which means you can only use simple wildcard patterns.",
				Computed:    true,
				ElementType: types.StringType,
			},
			"find_one_by": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"zone_identifier": schema.StringAttribute{
						Description: "Identifier",
						Required:    true,
					},
					"created_on": schema.StringAttribute{
						Description: "The timestamp of when the rule was created.",
						Optional:    true,
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
					},
					"page": schema.Float64Attribute{
						Description: "Page number of paginated results.",
						Computed:    true,
						Optional:    true,
					},
					"per_page": schema.Float64Attribute{
						Description: "The maximum number of results per page. You can only set the value to `1` or to a multiple of 5 such as `5`, `10`, `15`, or `20`.",
						Computed:    true,
						Optional:    true,
					},
					"priority": schema.Float64Attribute{
						Description: "The priority of the rule to control the processing order. A lower number indicates higher priority. If not provided, any rules with a configured priority will be processed before rules without a priority.",
						Optional:    true,
					},
					"uri_search": schema.StringAttribute{
						Description: "A single URI to search for in the list of URLs of existing rules.",
						Optional:    true,
					},
				},
			},
		},
	}
}

func (r *ZoneLockdownDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}

func (r *ZoneLockdownDataSource) ValidateConfig(ctx context.Context, req datasource.ValidateConfigRequest, resp *datasource.ValidateConfigResponse) {
}

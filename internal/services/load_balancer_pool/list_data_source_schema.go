// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package load_balancer_pool

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = &LoadBalancerPoolsDataSource{}
var _ datasource.DataSourceWithValidateConfig = &LoadBalancerPoolsDataSource{}

func (r LoadBalancerPoolsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "Identifier",
				Required:    true,
			},
			"monitor": schema.StringAttribute{
				Description: "The ID of the Monitor to use for checking the health of origins within this pool.",
				Optional:    true,
			},
			"max_items": schema.Int64Attribute{
				Description: "Max items to fetch, default: 1000",
				Optional:    true,
			},
			"items": schema.ListNestedAttribute{
				Description: "The items returned by the data source",
				Computed:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Computed: true,
						},
						"check_regions": schema.ListAttribute{
							Description: "A list of regions from which to run health checks. Null means every Cloudflare data center.",
							Computed:    true,
							ElementType: types.StringType,
						},
						"created_on": schema.StringAttribute{
							Computed: true,
						},
						"description": schema.StringAttribute{
							Description: "A human-readable description of the pool.",
							Computed:    true,
						},
						"disabled_at": schema.StringAttribute{
							Description: "This field shows up only if the pool is disabled. This field is set with the time the pool was disabled at.",
							Computed:    true,
						},
						"enabled": schema.BoolAttribute{
							Description: "Whether to enable (the default) or disable this pool. Disabled pools will not receive traffic and are excluded from health checks. Disabling a pool will cause any load balancers using it to failover to the next pool (if any).",
							Computed:    true,
						},
						"latitude": schema.Float64Attribute{
							Description: "The latitude of the data center containing the origins used in this pool in decimal degrees. If this is set, longitude must also be set.",
							Computed:    true,
						},
						"longitude": schema.Float64Attribute{
							Description: "The longitude of the data center containing the origins used in this pool in decimal degrees. If this is set, latitude must also be set.",
							Computed:    true,
						},
						"minimum_origins": schema.Int64Attribute{
							Description: "The minimum number of origins that must be healthy for this pool to serve traffic. If the number of healthy origins falls below this number, the pool will be marked unhealthy and will failover to the next available pool.",
							Computed:    true,
						},
						"modified_on": schema.StringAttribute{
							Computed: true,
						},
						"monitor": schema.StringAttribute{
							Description: "The ID of the Monitor to use for checking the health of origins within this pool.",
							Computed:    true,
						},
						"name": schema.StringAttribute{
							Description: "A short name (tag) for the pool. Only alphanumeric characters, hyphens, and underscores are allowed.",
							Computed:    true,
						},
						"notification_email": schema.StringAttribute{
							Description: "This field is now deprecated. It has been moved to Cloudflare's Centralized Notification service https://developers.cloudflare.com/fundamentals/notifications/. The email address to send health status notifications to. This can be an individual mailbox or a mailing list. Multiple emails can be supplied as a comma delimited list.",
							Computed:    true,
						},
						"origins": schema.ListNestedAttribute{
							Description: "The list of origins within this pool. Traffic directed at this pool is balanced across all currently healthy origins, provided the pool itself is healthy.",
							Computed:    true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"address": schema.StringAttribute{
										Description: "The IP address (IPv4 or IPv6) of the origin, or its publicly addressable hostname. Hostnames entered here should resolve directly to the origin, and not be a hostname proxied by Cloudflare. To set an internal/reserved address, virtual_network_id must also be set.",
										Computed:    true,
									},
									"disabled_at": schema.StringAttribute{
										Description: "This field shows up only if the origin is disabled. This field is set with the time the origin was disabled.",
										Computed:    true,
									},
									"enabled": schema.BoolAttribute{
										Description: "Whether to enable (the default) this origin within the pool. Disabled origins will not receive traffic and are excluded from health checks. The origin will only be disabled for the current pool.",
										Computed:    true,
									},
									"name": schema.StringAttribute{
										Description: "A human-identifiable name for the origin.",
										Computed:    true,
									},
									"virtual_network_id": schema.StringAttribute{
										Description: "The virtual network subnet ID the origin belongs in. Virtual network must also belong to the account.",
										Computed:    true,
									},
									"weight": schema.Float64Attribute{
										Description: "The weight of this origin relative to other origins in the pool. Based on the configured weight the total traffic is distributed among origins within the pool.\n- `origin_steering.policy=\"least_outstanding_requests\"`: Use weight to scale the origin's outstanding requests.\n- `origin_steering.policy=\"least_connections\"`: Use weight to scale the origin's open connections.",
										Computed:    true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func (r *LoadBalancerPoolsDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}

func (r *LoadBalancerPoolsDataSource) ValidateConfig(ctx context.Context, req datasource.ValidateConfigRequest, resp *datasource.ValidateConfigResponse) {
}

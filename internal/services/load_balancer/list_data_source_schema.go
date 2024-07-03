// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package load_balancer

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = &LoadBalancersDataSource{}
var _ datasource.DataSourceWithValidateConfig = &LoadBalancersDataSource{}

func (r LoadBalancersDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"zone_id": schema.StringAttribute{
				Required: true,
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
						"country_pools": schema.StringAttribute{
							Description: "A mapping of country codes to a list of pool IDs (ordered by their failover priority) for the given country. Any country not explicitly defined will fall back to using the corresponding region_pool mapping if it exists else to default_pools.",
							Computed:    true,
						},
						"created_on": schema.StringAttribute{
							Computed: true,
						},
						"default_pools": schema.ListAttribute{
							Description: "A list of pool IDs ordered by their failover priority. Pools defined here are used by default, or when region_pools are not configured for a given region.",
							Computed:    true,
							ElementType: types.StringType,
						},
						"description": schema.StringAttribute{
							Description: "Object description.",
							Computed:    true,
						},
						"enabled": schema.BoolAttribute{
							Description: "Whether to enable (the default) this load balancer.",
							Computed:    true,
						},
						"fallback_pool": schema.StringAttribute{
							Description: "The pool ID to use when all other pools are detected as unhealthy.",
							Computed:    true,
						},
						"modified_on": schema.StringAttribute{
							Computed: true,
						},
						"name": schema.StringAttribute{
							Description: "The DNS hostname to associate with your Load Balancer. If this hostname already exists as a DNS record in Cloudflare's DNS, the Load Balancer will take precedence and the DNS record will not be used.",
							Computed:    true,
						},
						"pop_pools": schema.StringAttribute{
							Description: "(Enterprise only): A mapping of Cloudflare PoP identifiers to a list of pool IDs (ordered by their failover priority) for the PoP (datacenter). Any PoPs not explicitly defined will fall back to using the corresponding country_pool, then region_pool mapping if it exists else to default_pools.",
							Computed:    true,
						},
						"proxied": schema.BoolAttribute{
							Description: "Whether the hostname should be gray clouded (false) or orange clouded (true).",
							Computed:    true,
						},
						"region_pools": schema.StringAttribute{
							Description: "A mapping of region codes to a list of pool IDs (ordered by their failover priority) for the given region. Any regions not explicitly defined will fall back to using default_pools.",
							Computed:    true,
						},
						"rules": schema.ListNestedAttribute{
							Description: "BETA Field Not General Access: A list of rules for this load balancer to execute.",
							Computed:    true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"condition": schema.StringAttribute{
										Description: "The condition expressions to evaluate. If the condition evaluates to true, the overrides or fixed_response in this rule will be applied. An empty condition is always true. For more details on condition expressions, please see https://developers.cloudflare.com/load-balancing/understand-basics/load-balancing-rules/expressions.",
										Computed:    true,
									},
									"disabled": schema.BoolAttribute{
										Description: "Disable this specific rule. It will no longer be evaluated by this load balancer.",
										Computed:    true,
									},
									"name": schema.StringAttribute{
										Description: "Name of this rule. Only used for human readability.",
										Computed:    true,
									},
									"priority": schema.Int64Attribute{
										Description: "The order in which rules should be executed in relation to each other. Lower values are executed first. Values do not need to be sequential. If no value is provided for any rule the array order of the rules field will be used to assign a priority.",
										Computed:    true,
									},
									"terminates": schema.BoolAttribute{
										Description: "If this rule's condition is true, this causes rule evaluation to stop after processing this rule.",
										Computed:    true,
									},
								},
							},
						},
						"session_affinity": schema.StringAttribute{
							Description: "Specifies the type of session affinity the load balancer should use unless specified as `\"none\"` or \"\" (default). The supported types are:\n- `\"cookie\"`: On the first request to a proxied load balancer, a cookie is generated, encoding information of which origin the request will be forwarded to. Subsequent requests, by the same client to the same load balancer, will be sent to the origin server the cookie encodes, for the duration of the cookie and as long as the origin server remains healthy. If the cookie has expired or the origin server is unhealthy, then a new origin server is calculated and used.\n- `\"ip_cookie\"`: Behaves the same as `\"cookie\"` except the initial origin selection is stable and based on the client's ip address.\n- `\"header\"`: On the first request to a proxied load balancer, a session key based on the configured HTTP headers (see `session_affinity_attributes.headers`) is generated, encoding the request headers used for storing in the load balancer session state which origin the request will be forwarded to. Subsequent requests to the load balancer with the same headers will be sent to the same origin server, for the duration of the session and as long as the origin server remains healthy. If the session has been idle for the duration of `session_affinity_ttl` seconds or the origin server is unhealthy, then a new origin server is calculated and used. See `headers` in `session_affinity_attributes` for additional required configuration.",
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive("none", "cookie", "ip_cookie", "header", ""),
							},
						},
						"session_affinity_ttl": schema.Float64Attribute{
							Description: "Time, in seconds, until a client's session expires after being created. Once the expiry time has been reached, subsequent requests may get sent to a different origin server. The accepted ranges per `session_affinity` policy are:\n- `\"cookie\"` / `\"ip_cookie\"`: The current default of 23 hours will be used unless explicitly set. The accepted range of values is between [1800, 604800].\n- `\"header\"`: The current default of 1800 seconds will be used unless explicitly set. The accepted range of values is between [30, 3600]. Note: With session affinity by header, sessions only expire after they haven't been used for the number of seconds specified.",
							Computed:    true,
						},
						"steering_policy": schema.StringAttribute{
							Description: "Steering Policy for this load balancer.\n- `\"off\"`: Use `default_pools`.\n- `\"geo\"`: Use `region_pools`/`country_pools`/`pop_pools`. For non-proxied requests, the country for `country_pools` is determined by `location_strategy`.\n- `\"random\"`: Select a pool randomly.\n- `\"dynamic_latency\"`: Use round trip time to select the closest pool in default_pools (requires pool health checks).\n- `\"proximity\"`: Use the pools' latitude and longitude to select the closest pool using the Cloudflare PoP location for proxied requests or the location determined by `location_strategy` for non-proxied requests.\n- `\"least_outstanding_requests\"`: Select a pool by taking into consideration `random_steering` weights, as well as each pool's number of outstanding requests. Pools with more pending requests are weighted proportionately less relative to others.\n- `\"least_connections\"`: Select a pool by taking into consideration `random_steering` weights, as well as each pool's number of open connections. Pools with more open connections are weighted proportionately less relative to others. Supported for HTTP/1 and HTTP/2 connections.\n- `\"\"`: Will map to `\"geo\"` if you use `region_pools`/`country_pools`/`pop_pools` otherwise `\"off\"`.",
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive("off", "geo", "random", "dynamic_latency", "proximity", "least_outstanding_requests", "least_connections", ""),
							},
						},
						"ttl": schema.Float64Attribute{
							Description: "Time to live (TTL) of the DNS entry for the IP address returned by this load balancer. This only applies to gray-clouded (unproxied) load balancers.",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func (r *LoadBalancersDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}

func (r *LoadBalancersDataSource) ValidateConfig(ctx context.Context, req datasource.ValidateConfigRequest, resp *datasource.ValidateConfigResponse) {
}

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package load_balancer_pool

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = &LoadBalancerPoolDataSource{}
var _ datasource.DataSourceWithValidateConfig = &LoadBalancerPoolDataSource{}

func (r LoadBalancerPoolDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "Identifier",
				Optional:    true,
			},
			"pool_id": schema.StringAttribute{
				Optional: true,
			},
			"created_on": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"disabled_at": schema.StringAttribute{
				Description: "This field shows up only if the pool is disabled. This field is set with the time the pool was disabled at.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"enabled": schema.BoolAttribute{
				Description: "Whether to enable (the default) or disable this pool. Disabled pools will not receive traffic and are excluded from health checks. Disabling a pool will cause any load balancers using it to failover to the next pool (if any).",
				Computed:    true,
			},
			"minimum_origins": schema.Int64Attribute{
				Description: "The minimum number of origins that must be healthy for this pool to serve traffic. If the number of healthy origins falls below this number, the pool will be marked unhealthy and will failover to the next available pool.",
				Computed:    true,
			},
			"modified_on": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"description": schema.StringAttribute{
				Description: "A human-readable description of the pool.",
				Computed:    true,
				Optional:    true,
			},
			"id": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"latitude": schema.Float64Attribute{
				Description: "The latitude of the data center containing the origins used in this pool in decimal degrees. If this is set, longitude must also be set.",
				Computed:    true,
				Optional:    true,
			},
			"longitude": schema.Float64Attribute{
				Description: "The longitude of the data center containing the origins used in this pool in decimal degrees. If this is set, latitude must also be set.",
				Computed:    true,
				Optional:    true,
			},
			"name": schema.StringAttribute{
				Description: "A short name (tag) for the pool. Only alphanumeric characters, hyphens, and underscores are allowed.",
				Computed:    true,
				Optional:    true,
			},
			"notification_email": schema.StringAttribute{
				Description: "This field is now deprecated. It has been moved to Cloudflare's Centralized Notification service https://developers.cloudflare.com/fundamentals/notifications/. The email address to send health status notifications to. This can be an individual mailbox or a mailing list. Multiple emails can be supplied as a comma delimited list.",
				Computed:    true,
				Optional:    true,
			},
			"check_regions": schema.ListAttribute{
				Description: "A list of regions from which to run health checks. Null means every Cloudflare data center.",
				Computed:    true,
				Optional:    true,
				ElementType: types.StringType,
			},
			"load_shedding": schema.SingleNestedAttribute{
				Description: "Configures load shedding policies and percentages for the pool.",
				Computed:    true,
				Optional:    true,
				Attributes: map[string]schema.Attribute{
					"default_percent": schema.Float64Attribute{
						Description: "The percent of traffic to shed from the pool, according to the default policy. Applies to new sessions and traffic without session affinity.",
						Computed:    true,
						Validators: []validator.Float64{
							float64validator.Between(0, 100),
						},
					},
					"default_policy": schema.StringAttribute{
						Description: "The default policy to use when load shedding. A random policy randomly sheds a given percent of requests. A hash policy computes a hash over the CF-Connecting-IP address and sheds all requests originating from a percent of IPs.",
						Computed:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("random", "hash"),
						},
					},
					"session_percent": schema.Float64Attribute{
						Description: "The percent of existing sessions to shed from the pool, according to the session policy.",
						Computed:    true,
						Validators: []validator.Float64{
							float64validator.Between(0, 100),
						},
					},
					"session_policy": schema.StringAttribute{
						Description: "Only the hash policy is supported for existing sessions (to avoid exponential decay).",
						Computed:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("hash"),
						},
					},
				},
			},
			"notification_filter": schema.SingleNestedAttribute{
				Description: "Filter pool and origin health notifications by resource type or health status. Use null to reset.",
				Computed:    true,
				Optional:    true,
				Attributes: map[string]schema.Attribute{
					"origin": schema.SingleNestedAttribute{
						Description: "Filter options for a particular resource type (pool or origin). Use null to reset.",
						Computed:    true,
						Optional:    true,
						Attributes: map[string]schema.Attribute{
							"disable": schema.BoolAttribute{
								Description: "If set true, disable notifications for this type of resource (pool or origin).",
								Computed:    true,
							},
							"healthy": schema.BoolAttribute{
								Description: "If present, send notifications only for this health status (e.g. false for only DOWN events). Use null to reset (all events).",
								Computed:    true,
								Optional:    true,
							},
						},
					},
					"pool": schema.SingleNestedAttribute{
						Description: "Filter options for a particular resource type (pool or origin). Use null to reset.",
						Computed:    true,
						Optional:    true,
						Attributes: map[string]schema.Attribute{
							"disable": schema.BoolAttribute{
								Description: "If set true, disable notifications for this type of resource (pool or origin).",
								Computed:    true,
							},
							"healthy": schema.BoolAttribute{
								Description: "If present, send notifications only for this health status (e.g. false for only DOWN events). Use null to reset (all events).",
								Computed:    true,
								Optional:    true,
							},
						},
					},
				},
			},
			"origin_steering": schema.SingleNestedAttribute{
				Description: "Configures origin steering for the pool. Controls how origins are selected for new sessions and traffic without session affinity.",
				Computed:    true,
				Optional:    true,
				Attributes: map[string]schema.Attribute{
					"policy": schema.StringAttribute{
						Description: "The type of origin steering policy to use.\n- `\"random\"`: Select an origin randomly.\n- `\"hash\"`: Select an origin by computing a hash over the CF-Connecting-IP address.\n- `\"least_outstanding_requests\"`: Select an origin by taking into consideration origin weights, as well as each origin's number of outstanding requests. Origins with more pending requests are weighted proportionately less relative to others.\n- `\"least_connections\"`: Select an origin by taking into consideration origin weights, as well as each origin's number of open connections. Origins with more open connections are weighted proportionately less relative to others. Supported for HTTP/1 and HTTP/2 connections.",
						Computed:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("random", "hash", "least_outstanding_requests", "least_connections"),
						},
					},
				},
			},
			"origins": schema.ListNestedAttribute{
				Description: "The list of origins within this pool. Traffic directed at this pool is balanced across all currently healthy origins, provided the pool itself is healthy.",
				Computed:    true,
				Optional:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"address": schema.StringAttribute{
							Description: "The IP address (IPv4 or IPv6) of the origin, or its publicly addressable hostname. Hostnames entered here should resolve directly to the origin, and not be a hostname proxied by Cloudflare. To set an internal/reserved address, virtual_network_id must also be set.",
							Computed:    true,
							Optional:    true,
						},
						"disabled_at": schema.StringAttribute{
							Description: "This field shows up only if the origin is disabled. This field is set with the time the origin was disabled.",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"enabled": schema.BoolAttribute{
							Description: "Whether to enable (the default) this origin within the pool. Disabled origins will not receive traffic and are excluded from health checks. The origin will only be disabled for the current pool.",
							Computed:    true,
						},
						"header": schema.SingleNestedAttribute{
							Description: "The request header is used to pass additional information with an HTTP request. Currently supported header is 'Host'.",
							Computed:    true,
							Optional:    true,
							Attributes: map[string]schema.Attribute{
								"host": schema.ListAttribute{
									Description: "The 'Host' header allows to override the hostname set in the HTTP request. Current support is 1 'Host' header override per origin.",
									Computed:    true,
									Optional:    true,
									ElementType: types.StringType,
								},
							},
						},
						"name": schema.StringAttribute{
							Description: "A human-identifiable name for the origin.",
							Computed:    true,
							Optional:    true,
						},
						"virtual_network_id": schema.StringAttribute{
							Description: "The virtual network subnet ID the origin belongs in. Virtual network must also belong to the account.",
							Computed:    true,
							Optional:    true,
						},
						"weight": schema.Float64Attribute{
							Description: "The weight of this origin relative to other origins in the pool. Based on the configured weight the total traffic is distributed among origins within the pool.\n- `origin_steering.policy=\"least_outstanding_requests\"`: Use weight to scale the origin's outstanding requests.\n- `origin_steering.policy=\"least_connections\"`: Use weight to scale the origin's open connections.",
							Computed:    true,
							Validators: []validator.Float64{
								float64validator.Between(0, 1),
							},
						},
					},
				},
			},
			"monitor": schema.StringAttribute{
				Description: "The ID of the Monitor to use for checking the health of origins within this pool.",
				Computed:    true,
				Optional:    true,
			},
			"filter": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"account_id": schema.StringAttribute{
						Description: "Identifier",
						Required:    true,
					},
					"monitor": schema.StringAttribute{
						Description: "The ID of the Monitor to use for checking the health of origins within this pool.",
						Optional:    true,
					},
				},
			},
		},
	}
}

func (r *LoadBalancerPoolDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}

func (r *LoadBalancerPoolDataSource) ValidateConfig(ctx context.Context, req datasource.ValidateConfigRequest, resp *datasource.ValidateConfigResponse) {
}

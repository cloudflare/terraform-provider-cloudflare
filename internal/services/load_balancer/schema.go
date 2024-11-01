// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package load_balancer

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/float64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithConfigValidators = (*LoadBalancerResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"zone_id": schema.StringAttribute{
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"fallback_pool": schema.StringAttribute{
				Description: "The pool ID to use when all other pools are detected as unhealthy.",
				Required:    true,
			},
			"name": schema.StringAttribute{
				Description: "The DNS hostname to associate with your Load Balancer. If this hostname already exists as a DNS record in Cloudflare's DNS, the Load Balancer will take precedence and the DNS record will not be used.",
				Required:    true,
			},
			"default_pools": schema.ListAttribute{
				Description: "A list of pool IDs ordered by their failover priority. Pools defined here are used by default, or when region_pools are not configured for a given region.",
				Required:    true,
				ElementType: types.StringType,
			},
			"description": schema.StringAttribute{
				Description: "Object description.",
				Optional:    true,
			},
			"session_affinity_ttl": schema.Float64Attribute{
				Description: "Time, in seconds, until a client's session expires after being created. Once the expiry time has been reached, subsequent requests may get sent to a different origin server. The accepted ranges per `session_affinity` policy are:\n- `\"cookie\"` / `\"ip_cookie\"`: The current default of 23 hours will be used unless explicitly set. The accepted range of values is between [1800, 604800].\n- `\"header\"`: The current default of 1800 seconds will be used unless explicitly set. The accepted range of values is between [30, 3600]. Note: With session affinity by header, sessions only expire after they haven't been used for the number of seconds specified.",
				Computed:    true,
				Optional:    true,
			},
			"ttl": schema.Float64Attribute{
				Description: "Time to live (TTL) of the DNS entry for the IP address returned by this load balancer. This only applies to gray-clouded (unproxied) load balancers.",
				Optional:    true,
				Computed:    true,
			},
			"country_pools": schema.MapAttribute{
				Description: "A mapping of country codes to a list of pool IDs (ordered by their failover priority) for the given country. Any country not explicitly defined will fall back to using the corresponding region_pool mapping if it exists else to default_pools.",
				Computed:    true,
				Optional:    true,
				Default:     mapdefault.StaticValue(types.MapValueMust(types.ListType{ElemType: types.StringType}, map[string]attr.Value{})),
				ElementType: types.ListType{
					ElemType: types.StringType,
				},
			},
			"networks": schema.ListAttribute{
				Description: "List of networks where Load Balancer or Pool is enabled.",
				Optional:    true,
				ElementType: types.StringType,
			},
			"pop_pools": schema.MapAttribute{
				Description: "(Enterprise only): A mapping of Cloudflare PoP identifiers to a list of pool IDs (ordered by their failover priority) for the PoP (datacenter). Any PoPs not explicitly defined will fall back to using the corresponding country_pool, then region_pool mapping if it exists else to default_pools.",
				Computed:    true,
				Optional:    true,
				Default:     mapdefault.StaticValue(types.MapValueMust(types.ListType{ElemType: types.StringType}, map[string]attr.Value{})),
				ElementType: types.ListType{
					ElemType: types.StringType,
				},
			},
			"region_pools": schema.MapAttribute{
				Description: "A mapping of region codes to a list of pool IDs (ordered by their failover priority) for the given region. Any regions not explicitly defined will fall back to using default_pools.",
				Computed:    true,
				Optional:    true,
				Default:     mapdefault.StaticValue(types.MapValueMust(types.ListType{ElemType: types.StringType}, map[string]attr.Value{})),
				ElementType: types.ListType{
					ElemType: types.StringType,
				},
			},
			"enabled": schema.BoolAttribute{
				Description: "Whether to enable (the default) this load balancer.",
				Computed:    true,
				Optional:    true,
				Default:     booldefault.StaticBool(true),
			},
			"proxied": schema.BoolAttribute{
				Description: "Whether the hostname should be gray clouded (false) or orange clouded (true).",
				Computed:    true,
				Optional:    true,
				Default:     booldefault.StaticBool(false),
			},
			"session_affinity": schema.StringAttribute{
				Description: "Specifies the type of session affinity the load balancer should use unless specified as `\"none\"`. The supported types are:\n- `\"cookie\"`: On the first request to a proxied load balancer, a cookie is generated, encoding information of which origin the request will be forwarded to. Subsequent requests, by the same client to the same load balancer, will be sent to the origin server the cookie encodes, for the duration of the cookie and as long as the origin server remains healthy. If the cookie has expired or the origin server is unhealthy, then a new origin server is calculated and used.\n- `\"ip_cookie\"`: Behaves the same as `\"cookie\"` except the initial origin selection is stable and based on the client's ip address.\n- `\"header\"`: On the first request to a proxied load balancer, a session key based on the configured HTTP headers (see `session_affinity_attributes.headers`) is generated, encoding the request headers used for storing in the load balancer session state which origin the request will be forwarded to. Subsequent requests to the load balancer with the same headers will be sent to the same origin server, for the duration of the session and as long as the origin server remains healthy. If the session has been idle for the duration of `session_affinity_ttl` seconds or the origin server is unhealthy, then a new origin server is calculated and used. See `headers` in `session_affinity_attributes` for additional required configuration.",
				Computed:    true,
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"none",
						"cookie",
						"ip_cookie",
						"header",
					),
				},
				Default: stringdefault.StaticString("none"),
			},
			"steering_policy": schema.StringAttribute{
				Description: "Steering Policy for this load balancer.\n- `\"off\"`: Use `default_pools`.\n- `\"geo\"`: Use `region_pools`/`country_pools`/`pop_pools`. For non-proxied requests, the country for `country_pools` is determined by `location_strategy`.\n- `\"random\"`: Select a pool randomly.\n- `\"dynamic_latency\"`: Use round trip time to select the closest pool in default_pools (requires pool health checks).\n- `\"proximity\"`: Use the pools' latitude and longitude to select the closest pool using the Cloudflare PoP location for proxied requests or the location determined by `location_strategy` for non-proxied requests.\n- `\"least_outstanding_requests\"`: Select a pool by taking into consideration `random_steering` weights, as well as each pool's number of outstanding requests. Pools with more pending requests are weighted proportionately less relative to others.\n- `\"least_connections\"`: Select a pool by taking into consideration `random_steering` weights, as well as each pool's number of open connections. Pools with more open connections are weighted proportionately less relative to others. Supported for HTTP/1 and HTTP/2 connections.\n- `\"\"`: Will map to `\"geo\"` if you use `region_pools`/`country_pools`/`pop_pools` otherwise `\"off\"`.",
				Computed:    true,
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"off",
						"geo",
						"random",
						"dynamic_latency",
						"proximity",
						"least_outstanding_requests",
						"least_connections",
						"",
					),
				},
				Default: stringdefault.StaticString(""),
			},
			"adaptive_routing": schema.SingleNestedAttribute{
				Description: "Controls features that modify the routing of requests to pools and origins in response to dynamic conditions, such as during the interval between active health monitoring requests. For example, zero-downtime failover occurs immediately when an origin becomes unavailable due to HTTP 521, 522, or 523 response codes. If there is another healthy origin in the same pool, the request is retried once against this alternate origin.",
				Computed:    true,
				Optional:    true,
				CustomType:  customfield.NewNestedObjectType[LoadBalancerAdaptiveRoutingModel](ctx),
				Attributes: map[string]schema.Attribute{
					"failover_across_pools": schema.BoolAttribute{
						Description: "Extends zero-downtime failover of requests to healthy origins from alternate pools, when no healthy alternate exists in the same pool, according to the failover order defined by traffic and origin steering. When set false (the default) zero-downtime failover will only occur between origins within the same pool. See `session_affinity_attributes` for control over when sessions are broken or reassigned.",
						Computed:    true,
						Optional:    true,
						Default:     booldefault.StaticBool(false),
					},
				},
			},
			"location_strategy": schema.SingleNestedAttribute{
				Description: "Controls location-based steering for non-proxied requests. See `steering_policy` to learn how steering is affected.",
				Computed:    true,
				Optional:    true,
				CustomType:  customfield.NewNestedObjectType[LoadBalancerLocationStrategyModel](ctx),
				Attributes: map[string]schema.Attribute{
					"mode": schema.StringAttribute{
						Description: "Determines the authoritative location when ECS is not preferred, does not exist in the request, or its GeoIP lookup is unsuccessful.\n- `\"pop\"`: Use the Cloudflare PoP location.\n- `\"resolver_ip\"`: Use the DNS resolver GeoIP location. If the GeoIP lookup is unsuccessful, use the Cloudflare PoP location.",
						Computed:    true,
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("pop", "resolver_ip"),
						},
						Default: stringdefault.StaticString("pop"),
					},
					"prefer_ecs": schema.StringAttribute{
						Description: "Whether the EDNS Client Subnet (ECS) GeoIP should be preferred as the authoritative location.\n- `\"always\"`: Always prefer ECS.\n- `\"never\"`: Never prefer ECS.\n- `\"proximity\"`: Prefer ECS only when `steering_policy=\"proximity\"`.\n- `\"geo\"`: Prefer ECS only when `steering_policy=\"geo\"`.",
						Computed:    true,
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive(
								"always",
								"never",
								"proximity",
								"geo",
							),
						},
						Default: stringdefault.StaticString("proximity"),
					},
				},
			},
			"random_steering": schema.SingleNestedAttribute{
				Description: "Configures pool weights.\n- `steering_policy=\"random\"`: A random pool is selected with probability proportional to pool weights.\n- `steering_policy=\"least_outstanding_requests\"`: Use pool weights to scale each pool's outstanding requests.\n- `steering_policy=\"least_connections\"`: Use pool weights to scale each pool's open connections.",
				Computed:    true,
				Optional:    true,
				CustomType:  customfield.NewNestedObjectType[LoadBalancerRandomSteeringModel](ctx),
				Attributes: map[string]schema.Attribute{
					"default_weight": schema.Float64Attribute{
						Description: "The default weight for pools in the load balancer that are not specified in the pool_weights map.",
						Computed:    true,
						Optional:    true,
						Validators: []validator.Float64{
							float64validator.Between(0, 1),
						},
						Default: float64default.StaticFloat64(1),
					},
					"pool_weights": schema.MapAttribute{
						Description: "A mapping of pool IDs to custom weights. The weight is relative to other pools in the load balancer.",
						Optional:    true,
						ElementType: types.Float64Type,
					},
				},
			},
			"rules": schema.ListNestedAttribute{
				Description: "BETA Field Not General Access: A list of rules for this load balancer to execute.",
				Computed:    true,
				Optional:    true,
				CustomType:  customfield.NewNestedObjectListType[LoadBalancerRulesModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"condition": schema.StringAttribute{
							Description: "The condition expressions to evaluate. If the condition evaluates to true, the overrides or fixed_response in this rule will be applied. An empty condition is always true. For more details on condition expressions, please see https://developers.cloudflare.com/load-balancing/understand-basics/load-balancing-rules/expressions.",
							Optional:    true,
						},
						"disabled": schema.BoolAttribute{
							Description: "Disable this specific rule. It will no longer be evaluated by this load balancer.",
							Computed:    true,
							Optional:    true,
							Default:     booldefault.StaticBool(false),
						},
						"fixed_response": schema.SingleNestedAttribute{
							Description: "A collection of fields used to directly respond to the eyeball instead of routing to a pool. If a fixed_response is supplied the rule will be marked as terminates.",
							Computed:    true,
							Optional:    true,
							CustomType:  customfield.NewNestedObjectType[LoadBalancerRulesFixedResponseModel](ctx),
							Attributes: map[string]schema.Attribute{
								"content_type": schema.StringAttribute{
									Description: "The http 'Content-Type' header to include in the response.",
									Optional:    true,
								},
								"location": schema.StringAttribute{
									Description: "The http 'Location' header to include in the response.",
									Optional:    true,
								},
								"message_body": schema.StringAttribute{
									Description: "Text to include as the http body.",
									Optional:    true,
								},
								"status_code": schema.Int64Attribute{
									Description: "The http status code to respond with.",
									Optional:    true,
								},
							},
						},
						"name": schema.StringAttribute{
							Description: "Name of this rule. Only used for human readability.",
							Optional:    true,
						},
						"overrides": schema.SingleNestedAttribute{
							Description: "A collection of overrides to apply to the load balancer when this rule's condition is true. All fields are optional.",
							Computed:    true,
							Optional:    true,
							CustomType:  customfield.NewNestedObjectType[LoadBalancerRulesOverridesModel](ctx),
							Attributes: map[string]schema.Attribute{
								"adaptive_routing": schema.SingleNestedAttribute{
									Description: "Controls features that modify the routing of requests to pools and origins in response to dynamic conditions, such as during the interval between active health monitoring requests. For example, zero-downtime failover occurs immediately when an origin becomes unavailable due to HTTP 521, 522, or 523 response codes. If there is another healthy origin in the same pool, the request is retried once against this alternate origin.",
									Computed:    true,
									Optional:    true,
									CustomType:  customfield.NewNestedObjectType[LoadBalancerRulesOverridesAdaptiveRoutingModel](ctx),
									Attributes: map[string]schema.Attribute{
										"failover_across_pools": schema.BoolAttribute{
											Description: "Extends zero-downtime failover of requests to healthy origins from alternate pools, when no healthy alternate exists in the same pool, according to the failover order defined by traffic and origin steering. When set false (the default) zero-downtime failover will only occur between origins within the same pool. See `session_affinity_attributes` for control over when sessions are broken or reassigned.",
											Computed:    true,
											Optional:    true,
											Default:     booldefault.StaticBool(false),
										},
									},
								},
								"country_pools": schema.MapAttribute{
									Description: "A mapping of country codes to a list of pool IDs (ordered by their failover priority) for the given country. Any country not explicitly defined will fall back to using the corresponding region_pool mapping if it exists else to default_pools.",
									Computed:    true,
									Optional:    true,
									Default:     mapdefault.StaticValue(types.MapValueMust(types.ListType{ElemType: types.StringType}, map[string]attr.Value{})),
									CustomType:  customfield.NewMapType[customfield.List[types.String]](ctx),
									ElementType: types.ListType{
										ElemType: types.StringType,
									},
								},
								"default_pools": schema.ListAttribute{
									Description: "A list of pool IDs ordered by their failover priority. Pools defined here are used by default, or when region_pools are not configured for a given region.",
									Optional:    true,
									ElementType: types.StringType,
								},
								"fallback_pool": schema.StringAttribute{
									Description: "The pool ID to use when all other pools are detected as unhealthy.",
									Optional:    true,
								},
								"location_strategy": schema.SingleNestedAttribute{
									Description: "Controls location-based steering for non-proxied requests. See `steering_policy` to learn how steering is affected.",
									Computed:    true,
									Optional:    true,
									CustomType:  customfield.NewNestedObjectType[LoadBalancerRulesOverridesLocationStrategyModel](ctx),
									Attributes: map[string]schema.Attribute{
										"mode": schema.StringAttribute{
											Description: "Determines the authoritative location when ECS is not preferred, does not exist in the request, or its GeoIP lookup is unsuccessful.\n- `\"pop\"`: Use the Cloudflare PoP location.\n- `\"resolver_ip\"`: Use the DNS resolver GeoIP location. If the GeoIP lookup is unsuccessful, use the Cloudflare PoP location.",
											Computed:    true,
											Optional:    true,
											Validators: []validator.String{
												stringvalidator.OneOfCaseInsensitive("pop", "resolver_ip"),
											},
											Default: stringdefault.StaticString("pop"),
										},
										"prefer_ecs": schema.StringAttribute{
											Description: "Whether the EDNS Client Subnet (ECS) GeoIP should be preferred as the authoritative location.\n- `\"always\"`: Always prefer ECS.\n- `\"never\"`: Never prefer ECS.\n- `\"proximity\"`: Prefer ECS only when `steering_policy=\"proximity\"`.\n- `\"geo\"`: Prefer ECS only when `steering_policy=\"geo\"`.",
											Computed:    true,
											Optional:    true,
											Validators: []validator.String{
												stringvalidator.OneOfCaseInsensitive(
													"always",
													"never",
													"proximity",
													"geo",
												),
											},
											Default: stringdefault.StaticString("proximity"),
										},
									},
								},
								"pop_pools": schema.MapAttribute{
									Description: "(Enterprise only): A mapping of Cloudflare PoP identifiers to a list of pool IDs (ordered by their failover priority) for the PoP (datacenter). Any PoPs not explicitly defined will fall back to using the corresponding country_pool, then region_pool mapping if it exists else to default_pools.",
									Computed:    true,
									Optional:    true,
									CustomType:  customfield.NewMapType[customfield.List[types.String]](ctx),
									ElementType: types.ListType{
										ElemType: types.StringType,
									},
								},
								"random_steering": schema.SingleNestedAttribute{
									Description: "Configures pool weights.\n- `steering_policy=\"random\"`: A random pool is selected with probability proportional to pool weights.\n- `steering_policy=\"least_outstanding_requests\"`: Use pool weights to scale each pool's outstanding requests.\n- `steering_policy=\"least_connections\"`: Use pool weights to scale each pool's open connections.",
									Computed:    true,
									Optional:    true,
									CustomType:  customfield.NewNestedObjectType[LoadBalancerRulesOverridesRandomSteeringModel](ctx),
									Attributes: map[string]schema.Attribute{
										"default_weight": schema.Float64Attribute{
											Description: "The default weight for pools in the load balancer that are not specified in the pool_weights map.",
											Computed:    true,
											Optional:    true,
											Validators: []validator.Float64{
												float64validator.Between(0, 1),
											},
											Default: float64default.StaticFloat64(1),
										},
										"pool_weights": schema.MapAttribute{
											Description: "A mapping of pool IDs to custom weights. The weight is relative to other pools in the load balancer.",
											Optional:    true,
											ElementType: types.Float64Type,
										},
									},
								},
								"region_pools": schema.MapAttribute{
									Description: "A mapping of region codes to a list of pool IDs (ordered by their failover priority) for the given region. Any regions not explicitly defined will fall back to using default_pools.",
									Computed:    true,
									Optional:    true,
									CustomType:  customfield.NewMapType[customfield.List[types.String]](ctx),
									ElementType: types.ListType{
										ElemType: types.StringType,
									},
								},
								"session_affinity": schema.StringAttribute{
									Description: "Specifies the type of session affinity the load balancer should use unless specified as `\"none\"`. The supported types are:\n- `\"cookie\"`: On the first request to a proxied load balancer, a cookie is generated, encoding information of which origin the request will be forwarded to. Subsequent requests, by the same client to the same load balancer, will be sent to the origin server the cookie encodes, for the duration of the cookie and as long as the origin server remains healthy. If the cookie has expired or the origin server is unhealthy, then a new origin server is calculated and used.\n- `\"ip_cookie\"`: Behaves the same as `\"cookie\"` except the initial origin selection is stable and based on the client's ip address.\n- `\"header\"`: On the first request to a proxied load balancer, a session key based on the configured HTTP headers (see `session_affinity_attributes.headers`) is generated, encoding the request headers used for storing in the load balancer session state which origin the request will be forwarded to. Subsequent requests to the load balancer with the same headers will be sent to the same origin server, for the duration of the session and as long as the origin server remains healthy. If the session has been idle for the duration of `session_affinity_ttl` seconds or the origin server is unhealthy, then a new origin server is calculated and used. See `headers` in `session_affinity_attributes` for additional required configuration.",
									Computed:    true,
									Optional:    true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive(
											"none",
											"cookie",
											"ip_cookie",
											"header",
										),
									},
									Default: stringdefault.StaticString("none"),
								},
								"session_affinity_attributes": schema.SingleNestedAttribute{
									Description: "Configures attributes for session affinity.",
									Computed:    true,
									Optional:    true,
									CustomType:  customfield.NewNestedObjectType[LoadBalancerRulesOverridesSessionAffinityAttributesModel](ctx),
									Attributes: map[string]schema.Attribute{
										"drain_duration": schema.Float64Attribute{
											Description: "Configures the drain duration in seconds. This field is only used when session affinity is enabled on the load balancer.",
											Computed:    true,
											Optional:    true,
										},
										"headers": schema.ListAttribute{
											Description: "Configures the names of HTTP headers to base session affinity on when header `session_affinity` is enabled. At least one HTTP header name must be provided. To specify the exact cookies to be used, include an item in the following format: `\"cookie:<cookie-name-1>,<cookie-name-2>\"` (example) where everything after the colon is a comma-separated list of cookie names. Providing only `\"cookie\"` will result in all cookies being used. The default max number of HTTP header names that can be provided depends on your plan: 5 for Enterprise, 1 for all other plans.",
											Optional:    true,
											ElementType: types.StringType,
										},
										"require_all_headers": schema.BoolAttribute{
											Description: "When header `session_affinity` is enabled, this option can be used to specify how HTTP headers on load balancing requests will be used. The supported values are:\n- `\"true\"`: Load balancing requests must contain *all* of the HTTP headers specified by the `headers` session affinity attribute, otherwise sessions aren't created.\n- `\"false\"`: Load balancing requests must contain *at least one* of the HTTP headers specified by the `headers` session affinity attribute, otherwise sessions aren't created.",
											Computed:    true,
											Optional:    true,
											Default:     booldefault.StaticBool(false),
										},
										"samesite": schema.StringAttribute{
											Description: "Configures the SameSite attribute on session affinity cookie. Value \"Auto\" will be translated to \"Lax\" or \"None\" depending if Always Use HTTPS is enabled. Note: when using value \"None\", the secure attribute can not be set to \"Never\".",
											Computed:    true,
											Optional:    true,
											Validators: []validator.String{
												stringvalidator.OneOfCaseInsensitive(
													"Auto",
													"Lax",
													"None",
													"Strict",
												),
											},
											Default: stringdefault.StaticString("Auto"),
										},
										"secure": schema.StringAttribute{
											Description: "Configures the Secure attribute on session affinity cookie. Value \"Always\" indicates the Secure attribute will be set in the Set-Cookie header, \"Never\" indicates the Secure attribute will not be set, and \"Auto\" will set the Secure attribute depending if Always Use HTTPS is enabled.",
											Computed:    true,
											Optional:    true,
											Validators: []validator.String{
												stringvalidator.OneOfCaseInsensitive(
													"Auto",
													"Always",
													"Never",
												),
											},
											Default: stringdefault.StaticString("Auto"),
										},
										"zero_downtime_failover": schema.StringAttribute{
											Description: "Configures the zero-downtime failover between origins within a pool when session affinity is enabled. This feature is currently incompatible with Argo, Tiered Cache, and Bandwidth Alliance. The supported values are:\n- `\"none\"`: No failover takes place for sessions pinned to the origin (default).\n- `\"temporary\"`: Traffic will be sent to another other healthy origin until the originally pinned origin is available; note that this can potentially result in heavy origin flapping.\n- `\"sticky\"`: The session affinity cookie is updated and subsequent requests are sent to the new origin. Note: Zero-downtime failover with sticky sessions is currently not supported for session affinity by header.",
											Computed:    true,
											Optional:    true,
											Validators: []validator.String{
												stringvalidator.OneOfCaseInsensitive(
													"none",
													"temporary",
													"sticky",
												),
											},
											Default: stringdefault.StaticString("none"),
										},
									},
								},
								"session_affinity_ttl": schema.Float64Attribute{
									Description: "Time, in seconds, until a client's session expires after being created. Once the expiry time has been reached, subsequent requests may get sent to a different origin server. The accepted ranges per `session_affinity` policy are:\n- `\"cookie\"` / `\"ip_cookie\"`: The current default of 23 hours will be used unless explicitly set. The accepted range of values is between [1800, 604800].\n- `\"header\"`: The current default of 1800 seconds will be used unless explicitly set. The accepted range of values is between [30, 3600]. Note: With session affinity by header, sessions only expire after they haven't been used for the number of seconds specified.",
									Computed:    true,
									Optional:    true,
								},
								"steering_policy": schema.StringAttribute{
									Description: "Steering Policy for this load balancer.\n- `\"off\"`: Use `default_pools`.\n- `\"geo\"`: Use `region_pools`/`country_pools`/`pop_pools`. For non-proxied requests, the country for `country_pools` is determined by `location_strategy`.\n- `\"random\"`: Select a pool randomly.\n- `\"dynamic_latency\"`: Use round trip time to select the closest pool in default_pools (requires pool health checks).\n- `\"proximity\"`: Use the pools' latitude and longitude to select the closest pool using the Cloudflare PoP location for proxied requests or the location determined by `location_strategy` for non-proxied requests.\n- `\"least_outstanding_requests\"`: Select a pool by taking into consideration `random_steering` weights, as well as each pool's number of outstanding requests. Pools with more pending requests are weighted proportionately less relative to others.\n- `\"least_connections\"`: Select a pool by taking into consideration `random_steering` weights, as well as each pool's number of open connections. Pools with more open connections are weighted proportionately less relative to others. Supported for HTTP/1 and HTTP/2 connections.\n- `\"\"`: Will map to `\"geo\"` if you use `region_pools`/`country_pools`/`pop_pools` otherwise `\"off\"`.",
									Computed:    true,
									Optional:    true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive(
											"off",
											"geo",
											"random",
											"dynamic_latency",
											"proximity",
											"least_outstanding_requests",
											"least_connections",
											"",
										),
									},
									Default: stringdefault.StaticString(""),
								},
								"ttl": schema.Float64Attribute{
									Description: "Time to live (TTL) of the DNS entry for the IP address returned by this load balancer. This only applies to gray-clouded (unproxied) load balancers.",
									Optional:    true,
									Computed:    true,
								},
							},
						},
						"priority": schema.Int64Attribute{
							Description: "The order in which rules should be executed in relation to each other. Lower values are executed first. Values do not need to be sequential. If no value is provided for any rule the array order of the rules field will be used to assign a priority.",
							Computed:    true,
							Optional:    true,
							Validators: []validator.Int64{
								int64validator.AtLeast(0),
							},
							Default: int64default.StaticInt64(0),
						},
						"terminates": schema.BoolAttribute{
							Description: "If this rule's condition is true, this causes rule evaluation to stop after processing this rule.",
							Computed:    true,
							Optional:    true,
							Default:     booldefault.StaticBool(false),
						},
					},
				},
			},
			"session_affinity_attributes": schema.SingleNestedAttribute{
				Description: "Configures attributes for session affinity.",
				Computed:    true,
				Optional:    true,
				CustomType:  customfield.NewNestedObjectType[LoadBalancerSessionAffinityAttributesModel](ctx),
				Attributes: map[string]schema.Attribute{
					"drain_duration": schema.Float64Attribute{
						Description: "Configures the drain duration in seconds. This field is only used when session affinity is enabled on the load balancer.",
						Computed:    true,
						Optional:    true,
					},
					"headers": schema.ListAttribute{
						Description: "Configures the names of HTTP headers to base session affinity on when header `session_affinity` is enabled. At least one HTTP header name must be provided. To specify the exact cookies to be used, include an item in the following format: `\"cookie:<cookie-name-1>,<cookie-name-2>\"` (example) where everything after the colon is a comma-separated list of cookie names. Providing only `\"cookie\"` will result in all cookies being used. The default max number of HTTP header names that can be provided depends on your plan: 5 for Enterprise, 1 for all other plans.",
						Optional:    true,
						ElementType: types.StringType,
					},
					"require_all_headers": schema.BoolAttribute{
						Description: "When header `session_affinity` is enabled, this option can be used to specify how HTTP headers on load balancing requests will be used. The supported values are:\n- `\"true\"`: Load balancing requests must contain *all* of the HTTP headers specified by the `headers` session affinity attribute, otherwise sessions aren't created.\n- `\"false\"`: Load balancing requests must contain *at least one* of the HTTP headers specified by the `headers` session affinity attribute, otherwise sessions aren't created.",
						Computed:    true,
						Optional:    true,
						// Default:     booldefault.StaticBool(false),
					},
					"samesite": schema.StringAttribute{
						Description: "Configures the SameSite attribute on session affinity cookie. Value \"Auto\" will be translated to \"Lax\" or \"None\" depending if Always Use HTTPS is enabled. Note: when using value \"None\", the secure attribute can not be set to \"Never\".",
						Computed:    true,
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive(
								"Auto",
								"Lax",
								"None",
								"Strict",
							),
						},
						// Default: stringdefault.StaticString("Auto"),
					},
					"secure": schema.StringAttribute{
						Description: "Configures the Secure attribute on session affinity cookie. Value \"Always\" indicates the Secure attribute will be set in the Set-Cookie header, \"Never\" indicates the Secure attribute will not be set, and \"Auto\" will set the Secure attribute depending if Always Use HTTPS is enabled.",
						Computed:    true,
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive(
								"Auto",
								"Always",
								"Never",
							),
						},
						// Default: stringdefault.StaticString("Auto"),
					},
					"zero_downtime_failover": schema.StringAttribute{
						Description: "Configures the zero-downtime failover between origins within a pool when session affinity is enabled. This feature is currently incompatible with Argo, Tiered Cache, and Bandwidth Alliance. The supported values are:\n- `\"none\"`: No failover takes place for sessions pinned to the origin (default).\n- `\"temporary\"`: Traffic will be sent to another other healthy origin until the originally pinned origin is available; note that this can potentially result in heavy origin flapping.\n- `\"sticky\"`: The session affinity cookie is updated and subsequent requests are sent to the new origin. Note: Zero-downtime failover with sticky sessions is currently not supported for session affinity by header.",
						Computed:    true,
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive(
								"none",
								"temporary",
								"sticky",
							),
						},
						// Default: stringdefault.StaticString("none"),
					},
				},
			},
			"created_on": schema.StringAttribute{
				Computed: true,
			},
			"modified_on": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

func (r *LoadBalancerResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *LoadBalancerResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_tunnel_cloudflared_config

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithConfigValidators = (*ZeroTrustTunnelCloudflaredConfigResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "UUID of the tunnel.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"tunnel_id": schema.StringAttribute{
				Description:   "UUID of the tunnel.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"account_id": schema.StringAttribute{
				Description:   "Identifier.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"config": schema.SingleNestedAttribute{
				Description: "The tunnel configuration and ingress rules.",
				Computed:    true,
				Optional:    true,
				CustomType:  customfield.NewNestedObjectType[ZeroTrustTunnelCloudflaredConfigConfigModel](ctx),
				Attributes: map[string]schema.Attribute{
					"ingress": schema.ListNestedAttribute{
						Description: "List of public hostname definitions. At least one ingress rule needs to be defined for the tunnel.",
						Computed:    true,
						Optional:    true,
						CustomType:  customfield.NewNestedObjectListType[ZeroTrustTunnelCloudflaredConfigConfigIngressModel](ctx),
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"hostname": schema.StringAttribute{
									Description: "Public hostname for this service.",
									Optional:    true,
								},
								"service": schema.StringAttribute{
									Description: "Protocol and address of destination server. Supported protocols: http://, https://, unix://, tcp://, ssh://, rdp://, unix+tls://, smb://. Alternatively can return a HTTP status code http_status:[code] e.g. 'http_status:404'.",
									Required:    true,
								},
								"origin_request": schema.SingleNestedAttribute{
									Description: "Configuration parameters for the public hostname specific connection settings between cloudflared and origin server.",
									Computed:    true,
									Optional:    true,
									CustomType:  customfield.NewNestedObjectType[ZeroTrustTunnelCloudflaredConfigConfigIngressOriginRequestModel](ctx),
									Attributes: map[string]schema.Attribute{
										"access": schema.SingleNestedAttribute{
											Description: "For all L7 requests to this hostname, cloudflared will validate each request's Cf-Access-Jwt-Assertion request header.",
											Computed:    true,
											Optional:    true,
											CustomType:  customfield.NewNestedObjectType[ZeroTrustTunnelCloudflaredConfigConfigIngressOriginRequestAccessModel](ctx),
											Attributes: map[string]schema.Attribute{
												"aud_tag": schema.ListAttribute{
													Description: "Access applications that are allowed to reach this hostname for this Tunnel. Audience tags can be identified in the dashboard or via the List Access policies API.",
													Required:    true,
													ElementType: types.StringType,
												},
												"team_name": schema.StringAttribute{
													Computed: true,
													Optional: true,
													Default:  stringdefault.StaticString("Your Zero Trust organization name."),
												},
												"required": schema.BoolAttribute{
													Description: "Deny traffic that has not fulfilled Access authorization.",
													Computed:    true,
													Optional:    true,
													Default:     booldefault.StaticBool(false),
												},
											},
										},
										"ca_pool": schema.StringAttribute{
											Description: "Path to the certificate authority (CA) for the certificate of your origin. This option should be used only if your certificate is not signed by Cloudflare.",
											Computed:    true,
											Optional:    true,
											Default:     stringdefault.StaticString(""),
										},
										"connect_timeout": schema.Int64Attribute{
											Description: "Timeout for establishing a new TCP connection to your origin server. This excludes the time taken to establish TLS, which is controlled by tlsTimeout.",
											Computed:    true,
											Optional:    true,
											Default:     int64default.StaticInt64(10),
										},
										"disable_chunked_encoding": schema.BoolAttribute{
											Description: "Disables chunked transfer encoding. Useful if you are running a WSGI server.",
											Optional:    true,
										},
										"http2_origin": schema.BoolAttribute{
											Description: "Attempt to connect to origin using HTTP2. Origin must be configured as https.",
											Optional:    true,
										},
										"http_host_header": schema.StringAttribute{
											Description: "Sets the HTTP Host header on requests sent to the local service.",
											Optional:    true,
										},
										"keep_alive_connections": schema.Int64Attribute{
											Description: "Maximum number of idle keepalive connections between Tunnel and your origin. This does not restrict the total number of concurrent connections.",
											Computed:    true,
											Optional:    true,
											Default:     int64default.StaticInt64(100),
										},
										"keep_alive_timeout": schema.Int64Attribute{
											Description: "Timeout after which an idle keepalive connection can be discarded.",
											Computed:    true,
											Optional:    true,
											Default:     int64default.StaticInt64(90),
										},
										"no_happy_eyeballs": schema.BoolAttribute{
											Description: "Disable the “happy eyeballs” algorithm for IPv4/IPv6 fallback if your local network has misconfigured one of the protocols.",
											Computed:    true,
											Optional:    true,
											Default:     booldefault.StaticBool(false),
										},
										"no_tls_verify": schema.BoolAttribute{
											Description: "Disables TLS verification of the certificate presented by your origin. Will allow any certificate from the origin to be accepted.",
											Computed:    true,
											Optional:    true,
											Default:     booldefault.StaticBool(false),
										},
										"origin_server_name": schema.StringAttribute{
											Description: "Hostname that cloudflared should expect from your origin server certificate.",
											Computed:    true,
											Optional:    true,
											Default:     stringdefault.StaticString(""),
										},
										"proxy_type": schema.StringAttribute{
											Description: `cloudflared starts a proxy server to translate HTTP traffic into TCP when proxying, for example, SSH or RDP. This configures what type of proxy will be started. Valid options are: "" for the regular proxy and "socks" for a SOCKS5 proxy.`,
											Computed:    true,
											Optional:    true,
											Default:     stringdefault.StaticString(""),
										},
										"tcp_keep_alive": schema.Int64Attribute{
											Description: "The timeout after which a TCP keepalive packet is sent on a connection between Tunnel and the origin server.",
											Computed:    true,
											Optional:    true,
											Default:     int64default.StaticInt64(30),
										},
										"tls_timeout": schema.Int64Attribute{
											Description: "Timeout for completing a TLS handshake to your origin server, if you have chosen to connect Tunnel to an HTTPS server.",
											Computed:    true,
											Optional:    true,
											Default:     int64default.StaticInt64(10),
										},
									},
								},
								"path": schema.StringAttribute{
									Description: "Requests with this path route to this public hostname.",
									Computed:    true,
									Optional:    true,
									Default:     stringdefault.StaticString(""),
								},
							},
						},
					},
					"origin_request": schema.SingleNestedAttribute{
						Description: "Configuration parameters for the public hostname specific connection settings between cloudflared and origin server.",
						Computed:    true,
						Optional:    true,
						CustomType:  customfield.NewNestedObjectType[ZeroTrustTunnelCloudflaredConfigConfigOriginRequestModel](ctx),
						Attributes: map[string]schema.Attribute{
							"access": schema.SingleNestedAttribute{
								Description: "For all L7 requests to this hostname, cloudflared will validate each request's Cf-Access-Jwt-Assertion request header.",
								Computed:    true,
								Optional:    true,
								CustomType:  customfield.NewNestedObjectType[ZeroTrustTunnelCloudflaredConfigConfigOriginRequestAccessModel](ctx),
								Attributes: map[string]schema.Attribute{
									"aud_tag": schema.ListAttribute{
										Description: "Access applications that are allowed to reach this hostname for this Tunnel. Audience tags can be identified in the dashboard or via the List Access policies API.",
										Required:    true,
										ElementType: types.StringType,
									},
									"team_name": schema.StringAttribute{
										Computed: true,
										Optional: true,
										Default:  stringdefault.StaticString("Your Zero Trust organization name."),
									},
									"required": schema.BoolAttribute{
										Description: "Deny traffic that has not fulfilled Access authorization.",
										Computed:    true,
										Optional:    true,
										Default:     booldefault.StaticBool(false),
									},
								},
							},
							"ca_pool": schema.StringAttribute{
								Description: "Path to the certificate authority (CA) for the certificate of your origin. This option should be used only if your certificate is not signed by Cloudflare.",
								Computed:    true,
								Optional:    true,
								Default:     stringdefault.StaticString(""),
							},
							"connect_timeout": schema.Int64Attribute{
								Description: "Timeout for establishing a new TCP connection to your origin server. This excludes the time taken to establish TLS, which is controlled by tlsTimeout.",
								Computed:    true,
								Optional:    true,
								Default:     int64default.StaticInt64(10),
							},
							"disable_chunked_encoding": schema.BoolAttribute{
								Description: "Disables chunked transfer encoding. Useful if you are running a WSGI server.",
								Optional:    true,
							},
							"http2_origin": schema.BoolAttribute{
								Description: "Attempt to connect to origin using HTTP2. Origin must be configured as https.",
								Optional:    true,
							},
							"http_host_header": schema.StringAttribute{
								Description: "Sets the HTTP Host header on requests sent to the local service.",
								Optional:    true,
							},
							"keep_alive_connections": schema.Int64Attribute{
								Description: "Maximum number of idle keepalive connections between Tunnel and your origin. This does not restrict the total number of concurrent connections.",
								Computed:    true,
								Optional:    true,
								Default:     int64default.StaticInt64(100),
							},
							"keep_alive_timeout": schema.Int64Attribute{
								Description: "Timeout after which an idle keepalive connection can be discarded.",
								Computed:    true,
								Optional:    true,
								Default:     int64default.StaticInt64(90),
							},
							"no_happy_eyeballs": schema.BoolAttribute{
								Description: "Disable the “happy eyeballs” algorithm for IPv4/IPv6 fallback if your local network has misconfigured one of the protocols.",
								Computed:    true,
								Optional:    true,
								Default:     booldefault.StaticBool(false),
							},
							"no_tls_verify": schema.BoolAttribute{
								Description: "Disables TLS verification of the certificate presented by your origin. Will allow any certificate from the origin to be accepted.",
								Computed:    true,
								Optional:    true,
								Default:     booldefault.StaticBool(false),
							},
							"origin_server_name": schema.StringAttribute{
								Description: "Hostname that cloudflared should expect from your origin server certificate.",
								Computed:    true,
								Optional:    true,
								Default:     stringdefault.StaticString(""),
							},
							"proxy_type": schema.StringAttribute{
								Description: `cloudflared starts a proxy server to translate HTTP traffic into TCP when proxying, for example, SSH or RDP. This configures what type of proxy will be started. Valid options are: "" for the regular proxy and "socks" for a SOCKS5 proxy.`,
								Computed:    true,
								Optional:    true,
								Default:     stringdefault.StaticString(""),
							},
							"tcp_keep_alive": schema.Int64Attribute{
								Description: "The timeout after which a TCP keepalive packet is sent on a connection between Tunnel and the origin server.",
								Computed:    true,
								Optional:    true,
								Default:     int64default.StaticInt64(30),
							},
							"tls_timeout": schema.Int64Attribute{
								Description: "Timeout for completing a TLS handshake to your origin server, if you have chosen to connect Tunnel to an HTTPS server.",
								Computed:    true,
								Optional:    true,
								Default:     int64default.StaticInt64(10),
							},
						},
					},
					"warp_routing": schema.SingleNestedAttribute{
						Description: "Enable private network access from WARP users to private network routes. This is enabled if the tunnel has an assigned route.",
						Optional:    true,
						Computed:    true,
						CustomType:  customfield.NewNestedObjectType[ZeroTrustTunnelCloudflaredConfigConfigWARPRoutingModel](ctx),
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{
								Computed: true,
								Optional: true,
								Default:  booldefault.StaticBool(false),
							},
						},
					},
				},
			},
			"created_at": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"source": schema.StringAttribute{
				Description: "Indicates if this is a locally or remotely configured tunnel. If `local`, manage the tunnel using a YAML file on the origin machine. If `cloudflare`, manage the tunnel's configuration on the Zero Trust dashboard.\nAvailable values: \"local\", \"cloudflare\".",
				Computed:    true,
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("local", "cloudflare"),
				},
				Default: stringdefault.StaticString("local"),
			},
			"version": schema.Int64Attribute{
				Description: "The version of the Tunnel Configuration.",
				Computed:    true,
			},
		},
	}
}

func (r *ZeroTrustTunnelCloudflaredConfigResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *ZeroTrustTunnelCloudflaredConfigResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}

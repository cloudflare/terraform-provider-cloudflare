// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_tunnel_cloudflared_config

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = (*ZeroTrustTunnelCloudflaredConfigDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "Identifier",
				Computed:    true,
			},
			"tunnel_id": schema.StringAttribute{
				Description: "UUID of the tunnel.",
				Computed:    true,
			},
			"created_at": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"source": schema.StringAttribute{
				Description: "Indicates if this is a locally or remotely configured tunnel. If `local`, manage the tunnel using a YAML file on the origin machine. If `cloudflare`, manage the tunnel's configuration on the Zero Trust dashboard.",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("local", "cloudflare"),
				},
			},
			"version": schema.Int64Attribute{
				Description: "The version of the Tunnel Configuration.",
				Computed:    true,
			},
			"config": schema.SingleNestedAttribute{
				Description: "The tunnel configuration and ingress rules.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[ZeroTrustTunnelCloudflaredConfigConfigDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"ingress": schema.ListNestedAttribute{
						Description: "List of public hostname definitions. At least one ingress rule needs to be defined for the tunnel.",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectListType[ZeroTrustTunnelCloudflaredConfigConfigIngressDataSourceModel](ctx),
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"hostname": schema.StringAttribute{
									Description: "Public hostname for this service.",
									Computed:    true,
								},
								"service": schema.StringAttribute{
									Description: "Protocol and address of destination server. Supported protocols: http://, https://, unix://, tcp://, ssh://, rdp://, unix+tls://, smb://. Alternatively can return a HTTP status code http_status:[code] e.g. 'http_status:404'.\n",
									Computed:    true,
								},
								"origin_request": schema.SingleNestedAttribute{
									Description: "Configuration parameters for the public hostname specific connection settings between cloudflared and origin server.",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectType[ZeroTrustTunnelCloudflaredConfigConfigIngressOriginRequestDataSourceModel](ctx),
									Attributes: map[string]schema.Attribute{
										"access": schema.SingleNestedAttribute{
											Description: "For all L7 requests to this hostname, cloudflared will validate each request's Cf-Access-Jwt-Assertion request header.",
											Computed:    true,
											CustomType:  customfield.NewNestedObjectType[ZeroTrustTunnelCloudflaredConfigConfigIngressOriginRequestAccessDataSourceModel](ctx),
											Attributes: map[string]schema.Attribute{
												"aud_tag": schema.ListAttribute{
													Description: "Access applications that are allowed to reach this hostname for this Tunnel. Audience tags can be identified in the dashboard or via the List Access policies API.",
													Computed:    true,
													CustomType:  customfield.NewListType[types.String](ctx),
													ElementType: types.StringType,
												},
												"team_name": schema.StringAttribute{
													Computed: true,
												},
												"required": schema.BoolAttribute{
													Description: "Deny traffic that has not fulfilled Access authorization.",
													Computed:    true,
												},
											},
										},
										"ca_pool": schema.StringAttribute{
											Description: "Path to the certificate authority (CA) for the certificate of your origin. This option should be used only if your certificate is not signed by Cloudflare.",
											Computed:    true,
										},
										"connect_timeout": schema.Int64Attribute{
											Description: "Timeout for establishing a new TCP connection to your origin server. This excludes the time taken to establish TLS, which is controlled by tlsTimeout.",
											Computed:    true,
										},
										"disable_chunked_encoding": schema.BoolAttribute{
											Description: "Disables chunked transfer encoding. Useful if you are running a WSGI server.",
											Computed:    true,
										},
										"http2_origin": schema.BoolAttribute{
											Description: "Attempt to connect to origin using HTTP2. Origin must be configured as https.",
											Computed:    true,
										},
										"http_host_header": schema.StringAttribute{
											Description: "Sets the HTTP Host header on requests sent to the local service.",
											Computed:    true,
										},
										"keep_alive_connections": schema.Int64Attribute{
											Description: "Maximum number of idle keepalive connections between Tunnel and your origin. This does not restrict the total number of concurrent connections.",
											Computed:    true,
										},
										"keep_alive_timeout": schema.Int64Attribute{
											Description: "Timeout after which an idle keepalive connection can be discarded.",
											Computed:    true,
										},
										"no_happy_eyeballs": schema.BoolAttribute{
											Description: "Disable the “happy eyeballs” algorithm for IPv4/IPv6 fallback if your local network has misconfigured one of the protocols.",
											Computed:    true,
										},
										"no_tls_verify": schema.BoolAttribute{
											Description: "Disables TLS verification of the certificate presented by your origin. Will allow any certificate from the origin to be accepted.",
											Computed:    true,
										},
										"origin_server_name": schema.StringAttribute{
											Description: "Hostname that cloudflared should expect from your origin server certificate.",
											Computed:    true,
										},
										"proxy_type": schema.StringAttribute{
											Description: "cloudflared starts a proxy server to translate HTTP traffic into TCP when proxying, for example, SSH or RDP. This configures what type of proxy will be started. Valid options are: \"\" for the regular proxy and \"socks\" for a SOCKS5 proxy.\n",
											Computed:    true,
										},
										"tcp_keep_alive": schema.Int64Attribute{
											Description: "The timeout after which a TCP keepalive packet is sent on a connection between Tunnel and the origin server.",
											Computed:    true,
										},
										"tls_timeout": schema.Int64Attribute{
											Description: "Timeout for completing a TLS handshake to your origin server, if you have chosen to connect Tunnel to an HTTPS server.",
											Computed:    true,
										},
									},
								},
								"path": schema.StringAttribute{
									Description: "Requests with this path route to this public hostname.",
									Computed:    true,
								},
							},
						},
					},
					"origin_request": schema.SingleNestedAttribute{
						Description: "Configuration parameters for the public hostname specific connection settings between cloudflared and origin server.",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectType[ZeroTrustTunnelCloudflaredConfigConfigOriginRequestDataSourceModel](ctx),
						Attributes: map[string]schema.Attribute{
							"access": schema.SingleNestedAttribute{
								Description: "For all L7 requests to this hostname, cloudflared will validate each request's Cf-Access-Jwt-Assertion request header.",
								Computed:    true,
								CustomType:  customfield.NewNestedObjectType[ZeroTrustTunnelCloudflaredConfigConfigOriginRequestAccessDataSourceModel](ctx),
								Attributes: map[string]schema.Attribute{
									"aud_tag": schema.ListAttribute{
										Description: "Access applications that are allowed to reach this hostname for this Tunnel. Audience tags can be identified in the dashboard or via the List Access policies API.",
										Computed:    true,
										CustomType:  customfield.NewListType[types.String](ctx),
										ElementType: types.StringType,
									},
									"team_name": schema.StringAttribute{
										Computed: true,
									},
									"required": schema.BoolAttribute{
										Description: "Deny traffic that has not fulfilled Access authorization.",
										Computed:    true,
									},
								},
							},
							"ca_pool": schema.StringAttribute{
								Description: "Path to the certificate authority (CA) for the certificate of your origin. This option should be used only if your certificate is not signed by Cloudflare.",
								Computed:    true,
							},
							"connect_timeout": schema.Int64Attribute{
								Description: "Timeout for establishing a new TCP connection to your origin server. This excludes the time taken to establish TLS, which is controlled by tlsTimeout.",
								Computed:    true,
							},
							"disable_chunked_encoding": schema.BoolAttribute{
								Description: "Disables chunked transfer encoding. Useful if you are running a WSGI server.",
								Computed:    true,
							},
							"http2_origin": schema.BoolAttribute{
								Description: "Attempt to connect to origin using HTTP2. Origin must be configured as https.",
								Computed:    true,
							},
							"http_host_header": schema.StringAttribute{
								Description: "Sets the HTTP Host header on requests sent to the local service.",
								Computed:    true,
							},
							"keep_alive_connections": schema.Int64Attribute{
								Description: "Maximum number of idle keepalive connections between Tunnel and your origin. This does not restrict the total number of concurrent connections.",
								Computed:    true,
							},
							"keep_alive_timeout": schema.Int64Attribute{
								Description: "Timeout after which an idle keepalive connection can be discarded.",
								Computed:    true,
							},
							"no_happy_eyeballs": schema.BoolAttribute{
								Description: "Disable the “happy eyeballs” algorithm for IPv4/IPv6 fallback if your local network has misconfigured one of the protocols.",
								Computed:    true,
							},
							"no_tls_verify": schema.BoolAttribute{
								Description: "Disables TLS verification of the certificate presented by your origin. Will allow any certificate from the origin to be accepted.",
								Computed:    true,
							},
							"origin_server_name": schema.StringAttribute{
								Description: "Hostname that cloudflared should expect from your origin server certificate.",
								Computed:    true,
							},
							"proxy_type": schema.StringAttribute{
								Description: "cloudflared starts a proxy server to translate HTTP traffic into TCP when proxying, for example, SSH or RDP. This configures what type of proxy will be started. Valid options are: \"\" for the regular proxy and \"socks\" for a SOCKS5 proxy.\n",
								Computed:    true,
							},
							"tcp_keep_alive": schema.Int64Attribute{
								Description: "The timeout after which a TCP keepalive packet is sent on a connection between Tunnel and the origin server.",
								Computed:    true,
							},
							"tls_timeout": schema.Int64Attribute{
								Description: "Timeout for completing a TLS handshake to your origin server, if you have chosen to connect Tunnel to an HTTPS server.",
								Computed:    true,
							},
						},
					},
					"warp_routing": schema.SingleNestedAttribute{
						Description: "Enable private network access from WARP users to private network routes. This is enabled if the tunnel has an assigned route.",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectType[ZeroTrustTunnelCloudflaredConfigConfigWARPRoutingDataSourceModel](ctx),
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{
								Computed: true,
							},
						},
					},
				},
			},
		},
	}
}

func (d *ZeroTrustTunnelCloudflaredConfigDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *ZeroTrustTunnelCloudflaredConfigDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package tunnel_config

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r TunnelConfigResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "Cloudflare account ID",
				Required:    true,
			},
			"tunnel_id": schema.StringAttribute{
				Description: "UUID of the tunnel.",
				Required:    true,
			},
			"config": schema.SingleNestedAttribute{
				Description: "The tunnel configuration and ingress rules.",
				Optional:    true,
				Attributes: map[string]schema.Attribute{
					"ingress": schema.ListNestedAttribute{
						Description: "List of public hostname definitions",
						Optional:    true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"hostname": schema.StringAttribute{
									Description: "Public hostname for this service.",
									Required:    true,
								},
								"service": schema.StringAttribute{
									Description: "Protocol and address of destination server. Supported protocols: http://, https://, unix://, tcp://, ssh://, rdp://, unix+tls://, smb://. Alternatively can return a HTTP status code http_status:[code] e.g. 'http_status:404'.\n",
									Required:    true,
								},
								"originrequest": schema.SingleNestedAttribute{
									Description: "Configuration parameters for the public hostname specific connection settings between cloudflared and origin server.",
									Optional:    true,
									Attributes: map[string]schema.Attribute{
										"access": schema.SingleNestedAttribute{
											Description: "For all L7 requests to this hostname, cloudflared will validate each request's Cf-Access-Jwt-Assertion request header.",
											Optional:    true,
											Attributes: map[string]schema.Attribute{
												"audtag": schema.ListAttribute{
													Description: "Access applications that are allowed to reach this hostname for this Tunnel. Audience tags can be identified in the dashboard or via the List Access policies API.",
													Required:    true,
													ElementType: types.StringType,
												},
												"teamname": schema.StringAttribute{
													Computed: true,
													Optional: true,
													Default:  stringdefault.StaticString("Your Zero Trust authentication domain."),
												},
												"required": schema.BoolAttribute{
													Description: "Deny traffic that has not fulfilled Access authorization.",
													Computed:    true,
													Optional:    true,
													Default:     booldefault.StaticBool(false),
												},
											},
										},
										"capool": schema.StringAttribute{
											Description: "Path to the certificate authority (CA) for the certificate of your origin. This option should be used only if your certificate is not signed by Cloudflare.",
											Computed:    true,
											Optional:    true,
											Default:     stringdefault.StaticString(""),
										},
										"connecttimeout": schema.Int64Attribute{
											Description: "Timeout for establishing a new TCP connection to your origin server. This excludes the time taken to establish TLS, which is controlled by tlsTimeout.",
											Computed:    true,
											Optional:    true,
											Default:     int64default.StaticInt64(10),
										},
										"disablechunkedencoding": schema.BoolAttribute{
											Description: "Disables chunked transfer encoding. Useful if you are running a WSGI server.",
											Optional:    true,
										},
										"http2origin": schema.BoolAttribute{
											Description: "Attempt to connect to origin using HTTP2. Origin must be configured as https.",
											Optional:    true,
										},
										"httphostheader": schema.StringAttribute{
											Description: "Sets the HTTP Host header on requests sent to the local service.",
											Optional:    true,
										},
										"keepaliveconnections": schema.Int64Attribute{
											Description: "Maximum number of idle keepalive connections between Tunnel and your origin. This does not restrict the total number of concurrent connections.",
											Computed:    true,
											Optional:    true,
											Default:     int64default.StaticInt64(100),
										},
										"keepalivetimeout": schema.Int64Attribute{
											Description: "Timeout after which an idle keepalive connection can be discarded.",
											Computed:    true,
											Optional:    true,
											Default:     int64default.StaticInt64(90),
										},
										"nohappyeyeballs": schema.BoolAttribute{
											Description: "Disable the “happy eyeballs” algorithm for IPv4/IPv6 fallback if your local network has misconfigured one of the protocols.",
											Computed:    true,
											Optional:    true,
											Default:     booldefault.StaticBool(false),
										},
										"notlsverify": schema.BoolAttribute{
											Description: "Disables TLS verification of the certificate presented by your origin. Will allow any certificate from the origin to be accepted.",
											Computed:    true,
											Optional:    true,
											Default:     booldefault.StaticBool(false),
										},
										"originservername": schema.StringAttribute{
											Description: "Hostname that cloudflared should expect from your origin server certificate.",
											Computed:    true,
											Optional:    true,
											Default:     stringdefault.StaticString(""),
										},
										"proxytype": schema.StringAttribute{
											Description: "cloudflared starts a proxy server to translate HTTP traffic into TCP when proxying, for example, SSH or RDP. This configures what type of proxy will be started. Valid options are: \"\" for the regular proxy and \"socks\" for a SOCKS5 proxy.\n",
											Computed:    true,
											Optional:    true,
											Default:     stringdefault.StaticString(""),
										},
										"tcpkeepalive": schema.Int64Attribute{
											Description: "The timeout after which a TCP keepalive packet is sent on a connection between Tunnel and the origin server.",
											Computed:    true,
											Optional:    true,
											Default:     int64default.StaticInt64(30),
										},
										"tlstimeout": schema.Int64Attribute{
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
					"originrequest": schema.SingleNestedAttribute{
						Description: "Configuration parameters for the public hostname specific connection settings between cloudflared and origin server.",
						Optional:    true,
						Attributes: map[string]schema.Attribute{
							"access": schema.SingleNestedAttribute{
								Description: "For all L7 requests to this hostname, cloudflared will validate each request's Cf-Access-Jwt-Assertion request header.",
								Optional:    true,
								Attributes: map[string]schema.Attribute{
									"audtag": schema.ListAttribute{
										Description: "Access applications that are allowed to reach this hostname for this Tunnel. Audience tags can be identified in the dashboard or via the List Access policies API.",
										Required:    true,
										ElementType: types.StringType,
									},
									"teamname": schema.StringAttribute{
										Computed: true,
										Optional: true,
										Default:  stringdefault.StaticString("Your Zero Trust authentication domain."),
									},
									"required": schema.BoolAttribute{
										Description: "Deny traffic that has not fulfilled Access authorization.",
										Computed:    true,
										Optional:    true,
										Default:     booldefault.StaticBool(false),
									},
								},
							},
							"capool": schema.StringAttribute{
								Description: "Path to the certificate authority (CA) for the certificate of your origin. This option should be used only if your certificate is not signed by Cloudflare.",
								Computed:    true,
								Optional:    true,
								Default:     stringdefault.StaticString(""),
							},
							"connecttimeout": schema.Int64Attribute{
								Description: "Timeout for establishing a new TCP connection to your origin server. This excludes the time taken to establish TLS, which is controlled by tlsTimeout.",
								Computed:    true,
								Optional:    true,
								Default:     int64default.StaticInt64(10),
							},
							"disablechunkedencoding": schema.BoolAttribute{
								Description: "Disables chunked transfer encoding. Useful if you are running a WSGI server.",
								Optional:    true,
							},
							"http2origin": schema.BoolAttribute{
								Description: "Attempt to connect to origin using HTTP2. Origin must be configured as https.",
								Optional:    true,
							},
							"httphostheader": schema.StringAttribute{
								Description: "Sets the HTTP Host header on requests sent to the local service.",
								Optional:    true,
							},
							"keepaliveconnections": schema.Int64Attribute{
								Description: "Maximum number of idle keepalive connections between Tunnel and your origin. This does not restrict the total number of concurrent connections.",
								Computed:    true,
								Optional:    true,
								Default:     int64default.StaticInt64(100),
							},
							"keepalivetimeout": schema.Int64Attribute{
								Description: "Timeout after which an idle keepalive connection can be discarded.",
								Computed:    true,
								Optional:    true,
								Default:     int64default.StaticInt64(90),
							},
							"nohappyeyeballs": schema.BoolAttribute{
								Description: "Disable the “happy eyeballs” algorithm for IPv4/IPv6 fallback if your local network has misconfigured one of the protocols.",
								Computed:    true,
								Optional:    true,
								Default:     booldefault.StaticBool(false),
							},
							"notlsverify": schema.BoolAttribute{
								Description: "Disables TLS verification of the certificate presented by your origin. Will allow any certificate from the origin to be accepted.",
								Computed:    true,
								Optional:    true,
								Default:     booldefault.StaticBool(false),
							},
							"originservername": schema.StringAttribute{
								Description: "Hostname that cloudflared should expect from your origin server certificate.",
								Computed:    true,
								Optional:    true,
								Default:     stringdefault.StaticString(""),
							},
							"proxytype": schema.StringAttribute{
								Description: "cloudflared starts a proxy server to translate HTTP traffic into TCP when proxying, for example, SSH or RDP. This configures what type of proxy will be started. Valid options are: \"\" for the regular proxy and \"socks\" for a SOCKS5 proxy.\n",
								Computed:    true,
								Optional:    true,
								Default:     stringdefault.StaticString(""),
							},
							"tcpkeepalive": schema.Int64Attribute{
								Description: "The timeout after which a TCP keepalive packet is sent on a connection between Tunnel and the origin server.",
								Computed:    true,
								Optional:    true,
								Default:     int64default.StaticInt64(30),
							},
							"tlstimeout": schema.Int64Attribute{
								Description: "Timeout for completing a TLS handshake to your origin server, if you have chosen to connect Tunnel to an HTTPS server.",
								Computed:    true,
								Optional:    true,
								Default:     int64default.StaticInt64(10),
							},
						},
					},
					"warp_routing": schema.SingleNestedAttribute{
						Description: "Enable private network access from WARP users to private network routes",
						Optional:    true,
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
		},
	}
}

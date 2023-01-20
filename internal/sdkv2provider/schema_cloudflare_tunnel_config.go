package sdkv2provider

import (
	"fmt"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCloudflareTunnelConfigSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"tunnel_id": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Identifier of the Tunnel to target for this configuration.",
		},
		consts.AccountIDSchemaKey: {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The account identifier to target for the resource.",
		},

		"config": {
			Type:        schema.TypeList,
			MaxItems:    1,
			Required:    true,
			Description: "Configuration block for Tunnel Configuration.",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"warp_routing": {
						Type:        schema.TypeList,
						Optional:    true,
						MaxItems:    1,
						Description: "If you're exposing a [private network](https://developers.cloudflare.com/cloudflare-one/connections/connect-apps/private-net/), you need to add the `warp-routing` key and set it to `true`.",
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"enabled": {
									Type:        schema.TypeBool,
									Optional:    true,
									Description: "Whether WARP routing is enabled.",
								},
							},
						},
					},
					"origin_request": {
						Type:     schema.TypeList,
						Optional: true,
						MaxItems: 1,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"connect_timeout": {
									Type:        schema.TypeString,
									Optional:    true,
									Description: "Timeout for establishing a new TCP connection to your origin server. This excludes the time taken to establish TLS, which is controlled by `tlsTimeout`.",
									Default:     "30s",
								},
								"tls_timeout": {
									Type:        schema.TypeString,
									Optional:    true,
									Description: "Timeout for completing a TLS handshake to your origin server, if you have chosen to connect Tunnel to an HTTPS server.",
									Default:     "10s",
								},
								"tcp_keep_alive": {
									Type:        schema.TypeString,
									Optional:    true,
									Description: "The timeout after which a TCP keepalive packet is sent on a connection between Tunnel and the origin server.",
									Default:     "30s",
								},
								"no_happy_eyeballs": {
									Type:        schema.TypeBool,
									Optional:    true,
									Description: "Disable the “happy eyeballs” algorithm for IPv4/IPv6 fallback if your local network has misconfigured one of the protocols.",
									Default:     false,
								},
								"keep_alive_connections": {
									Type:        schema.TypeInt,
									Optional:    true,
									Description: "Maximum number of idle keepalive connections between Tunnel and your origin. This does not restrict the total number of concurrent connections.",
									Default:     100,
								},
								"keep_alive_timeout": {
									Type:        schema.TypeString,
									Optional:    true,
									Description: "Timeout after which an idle keepalive connection can be discarded.",
									Default:     "1m30s",
								},
								"http_host_header": {
									Type:        schema.TypeString,
									Optional:    true,
									Description: "Sets the HTTP Host header on requests sent to the local service.",
									Default:     "",
								},
								"origin_server_name": {
									Type:        schema.TypeString,
									Optional:    true,
									Description: "Hostname that cloudflared should expect from your origin server certificate.",
									Default:     "",
								},
								"ca_pool": {
									Type:        schema.TypeString,
									Optional:    true,
									Description: "Path to the certificate authority (CA) for the certificate of your origin. This option should be used only if your certificate is not signed by Cloudflare.",
									Default:     "",
								},
								"no_tls_verify": {
									Type:        schema.TypeBool,
									Optional:    true,
									Description: "Disables TLS verification of the certificate presented by your origin. Will allow any certificate from the origin to be accepted.",
									Default:     false,
								},
								"disable_chunked_encoding": {
									Type:        schema.TypeBool,
									Optional:    true,
									Description: "Disables chunked transfer encoding. Useful if you are running a Web Server Gateway Interface (WSGI) server.",
									Default:     false,
								},
								"bastion_mode": {
									Type:        schema.TypeBool,
									Optional:    true,
									Description: "Runs as jump host.",
								},
								"proxy_address": {
									Type:        schema.TypeString,
									Optional:    true,
									Description: "cloudflared starts a proxy server to translate HTTP traffic into TCP when proxying, for example, SSH or RDP. This configures the listen address for that proxy.",
									Default:     "127.0.0.1",
								},
								"proxy_port": {
									Type:        schema.TypeInt,
									Optional:    true,
									Description: "cloudflared starts a proxy server to translate HTTP traffic into TCP when proxying, for example, SSH or RDP. This configures the listen port for that proxy. If set to zero, an unused port will randomly be chosen.",
									Default:     0,
								},
								"proxy_type": {
									Type:         schema.TypeString,
									Optional:     true,
									Description:  fmt.Sprintf("cloudflared starts a proxy server to translate HTTP traffic into TCP when proxying, for example, SSH or RDP. This configures what type of proxy will be started. %s", renderAvailableDocumentationValuesStringSlice([]string{"", "socks"})),
									ValidateFunc: validation.StringInSlice([]string{"", "socks"}, false),
									Default:      "",
								},
								"ip_rules": {
									Type:        schema.TypeSet,
									Optional:    true,
									Description: "IP rules for the proxy service.",
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											"prefix": {
												Type:        schema.TypeString,
												Optional:    true,
												Description: "IP rule prefix.",
											},
											"ports": {
												Type:        schema.TypeList,
												Optional:    true,
												Elem:        &schema.Schema{Type: schema.TypeInt},
												Description: "Ports to use within the IP rule.",
											},
											"allow": {
												Type:        schema.TypeBool,
												Optional:    true,
												Description: "Whether to allow the IP prefix.",
											},
										},
									},
								},
							},
						},
					},
					"ingress_rule": {
						Type:        schema.TypeList,
						Description: "Each incoming request received by cloudflared causes cloudflared to send a request to a local service. This section configures the rules that determine which requests are sent to which local services. [Read more](https://developers.cloudflare.com/cloudflare-one/connections/connect-apps/install-and-setup/tunnel-guide/local/local-management/ingress/)",
						Required:    true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"hostname": {
									Type:        schema.TypeString,
									Optional:    true,
									Description: "Hostname to match the incoming request with. If the hostname matches, the request will be sent to the service.",
								},
								"path": {
									Type:        schema.TypeString,
									Optional:    true,
									Description: "Path of the incoming request. If the path matches, the request will be sent to the local service.",
								},
								"service": {
									Type:        schema.TypeString,
									Required:    true,
									Description: "Name of the service to which the request will be sent.",
								},
							},
						},
					},
				},
			},
		},
	}
}

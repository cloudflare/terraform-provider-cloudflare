package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCloudflareTunnelConfigSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"tunnel_id": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "ID of the tunnel",
		},
		"account_id": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Cloudflare Account ID",
		},

		"config": {
			Type:     schema.TypeSet,
			MaxItems: 1,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"warp_routing": {
						Type:     schema.TypeList,
						Optional: true,
						MaxItems: 1,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"enabled": {
									Type:     schema.TypeBool,
									Required: true,
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
									Description: "HTTP proxy timeout for establishing a new connection",
								},
								"tls_timeout": {
									Type:        schema.TypeString,
									Optional:    true,
									Description: "HTTP proxy timeout for completing a TLS handshake",
								},
								"tcp_keep_alive": {
									Type:        schema.TypeString,
									Optional:    true,
									Description: "HTTP proxy TCP keepalive duration",
								},
								"no_happy_eyeballs": {
									Type:        schema.TypeBool,
									Optional:    true,
									Description: "HTTP proxy should disable 'happy eyeballs' for IPv4/v6 fallback",
								},
								"keep_alive_connections": {
									Type:        schema.TypeInt,
									Optional:    true,
									Description: "HTTP proxy maximum keepalive connection pool size",
								},
								"keep_alive_timeout": {
									Type:        schema.TypeString,
									Optional:    true,
									Description: "HTTP proxy timeout for closing an idle connection",
								},
								"http_host_header": {
									Type:        schema.TypeString,
									Optional:    true,
									Description: "Sets the HTTP Host header for the local webserver.",
								},
								"origin_server_name": {
									Type:        schema.TypeString,
									Optional:    true,
									Description: "Hostname on the origin server certificate.",
								},
								"ca_pool": {
									Type:        schema.TypeString,
									Optional:    true,
									Description: "Path to the CA for the certificate of your origin. This option should be used only if your certificate is not signed by Cloudflare.",
								},
								"no_tls_verify": {
									Type:        schema.TypeBool,
									Optional:    true,
									Description: "Disables TLS verification of the certificate presented by your origin.",
								},
								"disable_chunked_encoding": {
									Type:        schema.TypeBool,
									Optional:    true,
									Description: "Disables chunked transfer encoding.",
								},
								"bastion_mode": {
									Type:        schema.TypeBool,
									Optional:    true,
									Description: "Runs as jump host",
								},
								"proxy_address": {
									Type:        schema.TypeString,
									Optional:    true,
									Description: "Listen address for the proxy.",
								},
								"proxy_port": {
									Type:        schema.TypeInt,
									Optional:    true,
									Description: "Listen port for the proxy.",
								},
								"proxy_type": {
									Type:         schema.TypeString,
									Optional:     true,
									Description:  "Valid options are 'socks' or empty.",
									ValidateFunc: validation.StringInSlice([]string{"", "socks"}, false),
								},
								"ip_rules": {
									Type:        schema.TypeList,
									Optional:    true,
									Description: "IP rules for the proxy service",
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											"prefix": {
												Type:     schema.TypeString,
												Required: true,
											},
											"ports": {
												Type:     schema.TypeList,
												Required: true,
												Elem:     &schema.Schema{Type: schema.TypeInt},
											},
											"allow": {
												Type:     schema.TypeBool,
												Required: true,
											},
										},
									},
								},
							},
						},
					},
					"ingress_rule": {
						Type:     schema.TypeSet,
						Optional: true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"hostname": {Type: schema.TypeString, Required: true},
								"path":     {Type: schema.TypeString, Required: true},
								"service":  {Type: schema.TypeString, Required: true},
							},
						},
					},
				},
			},
		},
	}
}

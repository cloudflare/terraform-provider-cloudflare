package provider

import (
	"context"
	"time"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflareTunnelConfig() *schema.Resource {
	return &schema.Resource{
		Schema:        resourceCloudflareTunnelConfigSchema(),
		ReadContext:   resourceCloudflareTunnelConfigUpdate,
		UpdateContext: resourceCloudflareTeamsRuleUpdate,
	}
}

func buildTunnelConfig(d *schema.ResourceData) cloudflare.TunnelConfiguration {
	connectTimeOut, _ := time.ParseDuration(d.Get("connect_timeout").(string))
	tlsKeepAlive, _ := time.ParseDuration(d.Get("tls_timeout").(string))
	tcpKeepAlive, _ := time.ParseDuration(d.Get("tcp_keep_alive").(string))
	noHappyEyeballs := d.Get("no_happy_eyeballs").(bool)
	keepAliveConnections := d.Get("keep_alive_connections").(int)
	keepAliveTimeout, _ := time.ParseDuration(d.Get("keep_alive_timeout").(string))
	httpHostHeader := d.Get("http_host_header").(string)
	originServerName := d.Get("ca_pool").(string)
	noTLSVerify := d.Get("no_tls_verify").(bool)
	disableChuckedEncoding := d.Get("disabled_chunked_encoding").(bool)
	bastionMode := d.Get("bastion_mode").(bool)
	proxyAddress := d.Get("proxy_address").(string)
	proxyPort := uint(d.Get("proxy_port").(int))
	proxyType := d.Get("proxy_type").(string)
	origin := cloudflare.OriginRequestConfig{
		ConnectTimeout:         &connectTimeOut,
		TLSTimeout:             &tlsKeepAlive,
		TCPKeepAlive:           &tcpKeepAlive,
		NoHappyEyeballs:        &noHappyEyeballs,
		KeepAliveConnections:   &keepAliveConnections,
		KeepAliveTimeout:       &keepAliveTimeout,
		HTTPHostHeader:         &httpHostHeader,
		OriginServerName:       &originServerName,
		NoTLSVerify:            &noTLSVerify,
		DisableChunkedEncoding: &disableChuckedEncoding,
		BastionMode:            &bastionMode,
		ProxyAddress:           &proxyAddress,
		ProxyPort:              &proxyPort,
		ProxyType:              &proxyType,
	}
	ingressIpRules := []cloudflare.IngressIPRule{}
	for _, rule := range d.Get("ip_rules").([]interface{}) {
		prefix, _ := cloudflare.AnyPtr(rule.prefix).(*string)
		ingressIpRules = append(ingressIpRules, cloudflare.IngressIPRule{
			Prefix: prefix,
		})

	}
	return cloudflare.TunnelConfiguration{
		OriginRequest: origin,
	}
}

func resourceCloudflareTunnelConfigRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := d.Get("account_id").(string)
	tunnelID := d.Get("tunnel_id").(string)
	config, err := client.GetTunnelConfiguration(ctx, cloudflare.AccountIdentifier(accountID), tunnelID)
	if err != nil {
		return diag.Errorf("error find tunnel config %q: %w", d.Id(), err)
	}
	d.Set("tunnel_id", config.TunnelID)
	return nil
}

func resourceCloudflareTunnelConfigUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

}

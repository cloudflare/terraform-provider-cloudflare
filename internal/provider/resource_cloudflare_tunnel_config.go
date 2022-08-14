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
		ReadContext:   resourceCloudflareTunnelConfigRead,
		UpdateContext: resourceCloudflareTunnelConfigUpdate,
	}
}

func buildTunnelConfig(d *schema.ResourceData) cloudflare.TunnelConfiguration {
	connectTimeOut, _ := time.ParseDuration(d.Get("config.0.origin_request.0.connect_timeout").(string))
	tlsKeepAlive, _ := time.ParseDuration(d.Get("config.0.origin_request.0.tls_timeout").(string))
	tcpKeepAlive, _ := time.ParseDuration(d.Get("config.0.origin_request.0.tcp_keep_alive").(string))
	noHappyEyeballs := d.Get("config.0.origin_request.0.no_happy_eyeballs").(bool)
	keepAliveConnections := d.Get("config.0.origin_request.0.keep_alive_connections").(int)
	keepAliveTimeout, _ := time.ParseDuration(d.Get("config.0.origin_request.0.keep_alive_timeout").(string))
	httpHostHeader := d.Get("config.0.origin_request.0.http_host_header").(string)
	originServerName := d.Get("config.0.origin_request.0.ca_pool").(string)
	noTLSVerify := d.Get("config.0.origin_request.0.no_tls_verify").(bool)
	disableChuckedEncoding := d.Get("config.0.origin_request.0.disabled_chunked_encoding").(bool)
	bastionMode := d.Get("config.0.origin_request.0.bastion_mode").(bool)
	proxyAddress := d.Get("config.0.origin_request.0.proxy_address").(string)
	proxyPort := uint(d.Get("config.0.origin_request.0.proxy_port").(int))
	proxyType := d.Get("config.0.origin_request.0.proxy_type").(string)
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
	ingressIpRules := []cloudflare.UnvalidatedIngressRule{}
	if _, ok := d.GetOk("origin_request.0.ip_rules"); ok {
		if _, ok := d.GetOk("origin_request.0.ip_rules.0.prefix"); ok {
		}
	}

	warpRouting := cloudflare.WarpRoutingConfig{}
	if _, ok := d.GetOk("config.0.warp_routing"); ok {
		if _, ok := d.GetOk("config.0.warp_routing.enabled"); ok {
			warpRouting.Enabled = d.Get("config.0.warp_routing.enabled").(bool)
		}
	}

	return cloudflare.TunnelConfiguration{
		OriginRequest: origin,
		WarpRouting:   &warpRouting,
		Ingress:       ingressIpRules,
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
	client := meta.(*cloudflare.API)
	accountID := d.Get("account_id").(string)
	tunnelID := d.Get("tunnel_id").(string)
	tunnelConfig := buildTunnelConfig(d)
	tunnel := cloudflare.TunnelConfigurationParams{
		TunnelID: tunnelID,
		Config:   tunnelConfig,
	}
	_, err := client.UpdateTunnelConfiguration(ctx, cloudflare.AccountIdentifier(accountID), tunnel)
	if err != nil {
		return diag.Errorf("error updating tunnel config %q: %w", d.Id(), err)
	}
	return resourceCloudflareTunnelConfigRead(ctx, d, meta)
}

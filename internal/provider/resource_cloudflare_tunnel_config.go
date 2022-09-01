package provider

import (
	"context"
	"fmt"
	"time"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflareTunnelConfig() *schema.Resource {
	return &schema.Resource{
		Schema:        resourceCloudflareTunnelConfigSchema(),
		ReadContext:   resourceCloudflareTunnelConfigRead,
		CreateContext: resourceCloudflareTunnelConfigUpdate,
		UpdateContext: resourceCloudflareTunnelConfigUpdate,
		DeleteContext: resourceCloudflareTunnelConfigDelete,
	}
}

func buildTunnelConfig(d *schema.ResourceData) cloudflare.TunnelConfiguration {
	fmt.Println("Starting Tunnel config")
	// config := d.Get("config").(*schema.Set).List()[0].(map[string]interface{})
	// fmt.Printf("%#v\n", config["ingress_rule"].(map[string]interface{}))
	warpRouting := cloudflare.WarpRoutingConfig{}
	if _, ok := d.GetOk("config.0.warp_routing"); ok {
		if _, ok := d.GetOk("config.0.warp_routing.0.enabled"); ok {
			warpRouting.Enabled = d.Get("config.0.warp_routing.enabled").(bool)
		}
	}
	fmt.Println("Building origin")
	originRequest := "config.0.origin_request.0"

	connectTimeOut, _ := time.ParseDuration(d.Get(originRequest+".connect_timeout").(string))
	tlsKeepAlive, _ := time.ParseDuration(d.Get(originRequest+".tls_timeout").(string))
	tcpKeepAlive, _ := time.ParseDuration(d.Get(originRequest+".tcp_keep_alive").(string))
	noHappyEyeballs := d.Get(originRequest+".no_happy_eyeballs").(bool)
	keepAliveConnections := d.Get(originRequest+".keep_alive_connections").(int)
	keepAliveTimeout, _ := time.ParseDuration(d.Get(originRequest+".keep_alive_timeout").(string))
	httpHostHeader := d.Get(originRequest+".http_host_header").(string)
	originServerName := d.Get(originRequest+".ca_pool").(string)
	noTLSVerify := d.Get(originRequest+".no_tls_verify").(bool)
	disableChuckedEncoding := d.Get(originRequest+".disable_chunked_encoding").(bool)
	bastionMode := d.Get(originRequest+".bastion_mode").(bool)
	proxyAddress := d.Get(originRequest+".proxy_address").(string)
	proxyPort := uint(d.Get(originRequest+".proxy_port").(int))
	proxyType := d.Get(originRequest+".proxy_type").(string)
	// fmt.Printf("%s\n", connectTimeOut)
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

	
	var ipRules []cloudflare.IngressIPRule
	if items, ok := d.GetOk(originRequest+".ip_rules"); ok {
		fmt.Println("Building IP Rules")
		for _, item := range items.(*schema.Set).List() {
			rule := item.(map[string]interface{})
			newRule := cloudflare.IngressIPRule{}
			newRule.Prefix = cloudflare.StringPtr(rule["ip"].(string))
			newRule.Allow = rule["allow"].(bool)
			for _, value := range rule["ports"].([]interface{}) {
				newRule.Ports = append(newRule.Ports, value.(int))
			}
			ipRules = append(ipRules, newRule)
		}
	}
	origin.IPRules = ipRules

	
	var ingressRules []cloudflare.UnvalidatedIngressRule
	if items, ok := d.GetOk("config.0.ingress_rule"); ok {
		fmt.Println("Building Ingress Rules")
		for _, item := range items.(*schema.Set).List() {
			data := item.(map[string]interface{})
			fmt.Printf("%#v\n", data)
			ingressRules = append(ingressRules, cloudflare.UnvalidatedIngressRule{
				Hostname: data["hostname"].(string),
				Path:     data["path"].(string),
				Service:  data["service"].(string),
			})
		}
	}

	return cloudflare.TunnelConfiguration{
		OriginRequest: origin,
		WarpRouting:   &warpRouting,
		Ingress:       ingressRules,
	}
}

func resourceCloudflareTunnelConfigRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := d.Get("account_id").(string)
	tunnelID := d.Get("tunnel_id").(string)
	config, err := client.GetTunnelConfiguration(ctx, cloudflare.AccountIdentifier(accountID), tunnelID)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error getting tunnel config %q: %w", d.Id(), err))
	}
	d.SetId(config.TunnelID)
	return nil
}

func resourceCloudflareTunnelConfigUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := d.Get("account_id").(string)
	tunnelID := d.Get("tunnel_id").(string)
	tunnel := cloudflare.TunnelConfigurationParams{
		TunnelID: tunnelID,
		Config:   buildTunnelConfig(d),
	}
	// fmt.Printf("%#v\n", tunnel)
	_, err := client.UpdateTunnelConfiguration(ctx, cloudflare.AccountIdentifier(accountID), tunnel)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error creating / updating tunnel config %q: %w", tunnelID, err))
	}
	return resourceCloudflareTunnelConfigRead(ctx, d, meta)
}

func resourceCloudflareTunnelConfigDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := d.Get("account_id").(string)
	tunnel := cloudflare.TunnelConfigurationParams{
		TunnelID: d.Id(),
		Config:   cloudflare.TunnelConfiguration{},
	}
	_, err := client.UpdateTunnelConfiguration(ctx, cloudflare.AccountIdentifier(accountID), tunnel)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error deleting tunnel config %q: %w", d.Id(), err))
	}
	return resourceCloudflareTunnelConfigRead(ctx, d, meta)
}

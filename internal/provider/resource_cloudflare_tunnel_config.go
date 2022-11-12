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
	config := d.Get("config").(*schema.Set).List()[0].(map[string]interface{})
	fmt.Printf("%#v\n", config["ingress_rule"].(*schema.Set))
	warpRouting := cloudflare.WarpRoutingConfig{}
	if _, ok := d.GetOk("config.0.warp_routing"); ok {
		if _, ok := d.GetOk("config.0.warp_routing.0.enabled"); ok {
			warpRouting.Enabled = d.Get("config.0.warp_routing.enabled").(bool)
		}
	}
	fmt.Println("Building origin")
	originRequest := "config.0.origin_request.0"

	connectTimeOut, _ := time.ParseDuration(d.Get(originRequest + ".connect_timeout").(string))
	tlsKeepAlive, _ := time.ParseDuration(d.Get(originRequest + ".tls_timeout").(string))
	tcpKeepAlive, _ := time.ParseDuration(d.Get(originRequest + ".tcp_keep_alive").(string))
	keepAliveTimeout, _ := time.ParseDuration(d.Get(originRequest + ".keep_alive_timeout").(string))
	// fmt.Printf("%s\n", connectTimeOut)
	origin := cloudflare.OriginRequestConfig{
		ConnectTimeout:         &connectTimeOut,
		TLSTimeout:             &tlsKeepAlive,
		TCPKeepAlive:           &tcpKeepAlive,
		NoHappyEyeballs:        cloudflare.BoolPtr(d.Get(originRequest + ".no_happy_eyeballs").(bool)),
		KeepAliveConnections:   cloudflare.IntPtr(d.Get(originRequest + ".keep_alive_connections").(int)),
		KeepAliveTimeout:       &keepAliveTimeout,
		HTTPHostHeader:         cloudflare.StringPtr(d.Get(originRequest + ".http_host_header").(string)),
		OriginServerName:       cloudflare.StringPtr(d.Get(originRequest + ".origin_server_name").(string)),
		CAPool:                 cloudflare.StringPtr(d.Get(originRequest + ".ca_pool").(string)),
		NoTLSVerify:            cloudflare.BoolPtr(d.Get(originRequest + ".no_tls_verify").(bool)),
		DisableChunkedEncoding: cloudflare.BoolPtr(d.Get(originRequest + ".disable_chunked_encoding").(bool)),
		BastionMode:            cloudflare.BoolPtr(d.Get(originRequest + ".bastion_mode").(bool)),
		ProxyAddress:           cloudflare.StringPtr(d.Get(originRequest + ".proxy_address").(string)),
		ProxyPort:              cloudflare.UintPtr(uint(d.Get(originRequest + ".proxy_port").(int))),
		ProxyType:              cloudflare.StringPtr(d.Get(originRequest + ".proxy_type").(string)),
	}

	var ipRules []cloudflare.IngressIPRule
	if items, ok := d.GetOk(originRequest + ".ip_rules"); ok {
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
	_, err := client.UpdateTunnelConfiguration(ctx, cloudflare.AccountIdentifier(accountID), tunnel)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error creating / updating tunnel config %q: %w", tunnelID, err))
	}
	return resourceCloudflareTunnelConfigRead(ctx, d, meta)
}

func resourceCloudflareTunnelConfigDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := d.Get("account_id").(string)
	tunnelID := d.Get("tunnel_id").(string)
	err := client.DeleteTunnel(ctx, cloudflare.AccountIdentifier(accountID), tunnelID)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error deleting tunnel config %q: %w", d.Id(), err))
	}
	return nil
}

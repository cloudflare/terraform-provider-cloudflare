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
	// fmt.Printf("ingress_rule\n%#v\n", config["ingress_rule"].(*schema.Set))

	// fmt.Printf("origin_request\n%#v\n", config["origin_request"].(*schema.Set))

	warpRouting := cloudflare.WarpRoutingConfig{}
	warpConfig := config["warp_routing"].(*schema.Set).List()[0].(map[string]interface{})
	warpRouting.Enabled = warpConfig["enabled"].(bool)

	originConfig := config["origin_request"].(*schema.Set).List()[0].(map[string]interface{})

	connectTimeOut, _ := time.ParseDuration(originConfig["connect_timeout"].(string))
	tlsKeepAlive, _ := time.ParseDuration(originConfig["tls_timeout"].(string))
	tcpKeepAlive, _ := time.ParseDuration(originConfig["tcp_keep_alive"].(string))
	keepAliveTimeout, _ := time.ParseDuration(originConfig["keep_alive_timeout"].(string))
	// fmt.Printf("%s\n", connectTimeOut)
	origin := cloudflare.OriginRequestConfig{
		ConnectTimeout:         &connectTimeOut,
		TLSTimeout:             &tlsKeepAlive,
		TCPKeepAlive:           &tcpKeepAlive,
		NoHappyEyeballs:        cloudflare.BoolPtr(originConfig["no_happy_eyeballs"].(bool)),
		KeepAliveConnections:   cloudflare.IntPtr(originConfig["keep_alive_connections"].(int)),
		KeepAliveTimeout:       &keepAliveTimeout,
		HTTPHostHeader:         cloudflare.StringPtr(originConfig["http_host_header"].(string)),
		OriginServerName:       cloudflare.StringPtr(originConfig["origin_server_name"].(string)),
		CAPool:                 cloudflare.StringPtr(originConfig["ca_pool"].(string)),
		NoTLSVerify:            cloudflare.BoolPtr(originConfig["no_tls_verify"].(bool)),
		DisableChunkedEncoding: cloudflare.BoolPtr(originConfig["disable_chunked_encoding"].(bool)),
		BastionMode:            cloudflare.BoolPtr(originConfig["bastion_mode"].(bool)),
		ProxyAddress:           cloudflare.StringPtr(originConfig["proxy_address"].(string)),
		ProxyPort:              cloudflare.UintPtr(uint(originConfig["proxy_port"].(int))),
		ProxyType:              cloudflare.StringPtr(originConfig["proxy_type"].(string)),
	}

	var ipRules []cloudflare.IngressIPRule
	for _, ingressRule := range originConfig["ip_rules"].(*schema.Set).List() {
		ingressRuleConfig := ingressRule.(map[string]interface{})
		fmt.Printf("%s\n", ingressRuleConfig)
		ipRule := cloudflare.IngressIPRule{
			Prefix: cloudflare.StringPtr(ingressRuleConfig["prefix"].(string)),
			Allow:  ingressRuleConfig["allow"].(bool),
		}
		for _, value := range ingressRuleConfig["ports"].(*schema.Set).List() {
			ipRule.Ports = append(ipRule.Ports, value.(int))
		}
		ipRules = append(ipRules, ipRule)
	}
	origin.IPRules = ipRules

	var ingressRules []cloudflare.UnvalidatedIngressRule
	for _, ingressRule := range config["ingress_rule"].([]interface{}) {
		ingressRuleConfig := ingressRule.(map[string]interface{})
		ingressRule := cloudflare.UnvalidatedIngressRule{
			Service:  ingressRuleConfig["service"].(string),
			Hostname: ingressRuleConfig["hostname"].(string),
			Path:     ingressRuleConfig["path"].(string),
		}
		ingressRules = append(ingressRules, ingressRule)
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
	config, err := client.GetTunnelConfiguration(ctx, cloudflare.AccountIdentifier(accountID), d.Id())
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
	result, err := client.UpdateTunnelConfiguration(ctx, cloudflare.AccountIdentifier(accountID), tunnel)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error creating / updating tunnel config %q: %w", tunnelID, err))
	}

	d.SetId(result.TunnelID)
	return resourceCloudflareTunnelConfigRead(ctx, d, meta)
}

func resourceCloudflareTunnelConfigDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := d.Get("account_id").(string)

	err := client.DeleteTunnel(ctx, cloudflare.AccountIdentifier(accountID), d.Id())
	if err != nil {
		return diag.FromErr(fmt.Errorf("error deleting tunnel config %q: %w", d.Id(), err))
	}

	return nil
}

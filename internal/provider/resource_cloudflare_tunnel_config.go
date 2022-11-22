package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-log/tflog"
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
	warpRouting := cloudflare.WarpRoutingConfig{}
	if item, ok := d.GetOk("config.0.warp_routing.0"); ok {
		warpRouting = cloudflare.WarpRoutingConfig{
			Enabled: item.(map[string]interface{})["enabled"].(bool),
		}
	}

	originConfig := cloudflare.OriginRequestConfig{}
	if item, ok := d.GetOk("config.0.origin_request.0"); ok {
		originRequest := item.(map[string]interface{})
		fmt.Printf("%#v\n", d.Get("config.0.origin_request.0.connect_timeout"))
		if v, ok := originRequest["connect_timeout"]; ok {
			timeout, _ := time.ParseDuration(v.(string))
			originConfig.ConnectTimeout = &timeout
		}
		if v, ok := originRequest["tls_timeout"]; ok {
			timeout, _ := time.ParseDuration(v.(string))
			originConfig.TLSTimeout = &timeout
		}
		if v, ok := originRequest["tcp_keep_alive"]; ok {
			timeout, _ := time.ParseDuration(v.(string))
			originConfig.TCPKeepAlive = &timeout
		}
		if v, ok := originRequest["no_happy_eyeballs"]; ok {
			originConfig.NoHappyEyeballs = cloudflare.BoolPtr(v.(bool))
		}
		if v, ok := originRequest["keep_alive_connections"]; ok {
			originConfig.KeepAliveConnections = cloudflare.IntPtr(v.(int))
		}
		if v, ok := originRequest["keep_alive_timeout"]; ok {
			timeout, _ := time.ParseDuration(v.(string))
			originConfig.KeepAliveTimeout = &timeout
		}
		if v, ok := originRequest["http_host_header"]; ok {
			originConfig.HTTPHostHeader = cloudflare.StringPtr(v.(string))
		}
		if v, ok := originRequest["origin_server_name"]; ok {
			originConfig.OriginServerName = cloudflare.StringPtr(v.(string))
		}
		if v, ok := originRequest["ca_pool"]; ok {
			originConfig.CAPool = cloudflare.StringPtr(v.(string))
		}
		if v, ok := originRequest["no_tls_verify"]; ok {
			originConfig.NoTLSVerify = cloudflare.BoolPtr(v.(bool))
		}
		if v, ok := originRequest["disable_chunked_encoding"]; ok {
			originConfig.DisableChunkedEncoding = cloudflare.BoolPtr(v.(bool))
		}
		if v, ok := originRequest["bastion_mode"]; ok {
			originConfig.BastionMode = cloudflare.BoolPtr(v.(bool))
		}
		if v, ok := originRequest["proxy_address"]; ok {
			originConfig.ProxyAddress = cloudflare.StringPtr(v.(string))
		}
		if v, ok := originRequest["proxy_port"]; ok {
			originConfig.ProxyPort = cloudflare.UintPtr(uint(v.(int)))
		}
		if v, ok := originRequest["proxy_type"]; ok {
			originConfig.ProxyType = cloudflare.StringPtr(v.(string))
		}

		var ipRules []cloudflare.IngressIPRule
		if v, ok := originRequest["ip_rules"]; ok {
			for _, ingressRule := range v.([]interface{}) {
				ingressRuleConfig := ingressRule.(map[string]interface{})
				ipRule := cloudflare.IngressIPRule{
					Prefix: cloudflare.StringPtr(ingressRuleConfig["prefix"].(string)),
					Allow:  ingressRuleConfig["allow"].(bool),
				}
				for _, value := range ingressRuleConfig["ports"].([]interface{}) {
					ipRule.Ports = append(ipRule.Ports, value.(int))
				}
				ipRules = append(ipRules, ipRule)
			}
		}
		originConfig.IPRules = ipRules
	}

	var ingressRules []cloudflare.UnvalidatedIngressRule
	for _, ingressRule := range d.Get("config.0.ingress_rule").([]interface{}) {
		ingressRuleConfig := ingressRule.(map[string]interface{})
		ingressRule := cloudflare.UnvalidatedIngressRule{
			Service:  ingressRuleConfig["service"].(string),
			Hostname: ingressRuleConfig["hostname"].(string),
			Path:     ingressRuleConfig["path"].(string),
		}
		ingressRules = append(ingressRules, ingressRule)
	}

	return cloudflare.TunnelConfiguration{
		OriginRequest: originConfig,
		WarpRouting:   &warpRouting,
		Ingress:       ingressRules,
	}
}

func resourceCloudflareTunnelConfigRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := d.Get("account_id").(string)
	result, err := client.GetTunnelConfiguration(ctx, cloudflare.AccountIdentifier(accountID), d.Id())
	tflog.Debug(ctx, fmt.Sprintf("GetTunnelConfiguration: %+v", result))
	if err != nil {
		return diag.FromErr(fmt.Errorf("error getting tunnel config %q: %w", d.Id(), err))
	}
	d.SetId(result.TunnelID)
	config := result.Config
	var configMap []map[string]interface{}

	warpConfig := []map[string]interface{}{
		{
			"enabled": config.WarpRouting.Enabled,
		},
	}
	fmt.Printf("Happy Eyeballs %#v\n", config.OriginRequest.NoHappyEyeballs)
	originRequestConfig := []map[string]interface{}{
		{
			"connect_timeout":          optDateDurationToString(config.OriginRequest.ConnectTimeout),
			"tls_timeout":              optDateDurationToString(config.OriginRequest.TLSTimeout),
			"tcp_keep_alive":           optDateDurationToString(config.OriginRequest.TCPKeepAlive),
			"no_happy_eyeballs":        optBoolToString(config.OriginRequest.NoHappyEyeballs),
			"keep_alive_connections":   optIntToString(config.OriginRequest.KeepAliveConnections),
			"keep_alive_timeout":       optDateDurationToString(config.OriginRequest.KeepAliveTimeout),
			"http_host_header":         cloudflare.String(config.OriginRequest.HTTPHostHeader),
			"origin_server_name":       cloudflare.String(config.OriginRequest.OriginServerName),
			"ca_pool":                  cloudflare.String(config.OriginRequest.CAPool),
			"no_tls_verify":            optBoolToString(config.OriginRequest.NoTLSVerify),
			"disable_chunked_encoding": optBoolToString(config.OriginRequest.DisableChunkedEncoding),
			"bastion_mode":             optBoolToString(config.OriginRequest.BastionMode),
			"proxy_address":            cloudflare.String(config.OriginRequest.ProxyAddress),
			"proxy_port":               optUintToString(config.OriginRequest.ProxyPort),
			"proxy_type":               cloudflare.String(config.OriginRequest.ProxyType),
		},
	}

	var ipRules []map[string]interface{}
	for _, ipRule := range config.OriginRequest.IPRules {
		ipRules = append(ipRules, map[string]interface{}{
			"prefix": ipRule.Prefix,
			"allow":  ipRule.Allow,
			"ports":  ipRule.Ports,
		})
	}
	originRequestConfig[0]["ip_rules"] = ipRules

	var ingressRules []map[string]interface{}
	for _, ingressRule := range config.Ingress {
		ingressRules = append(ingressRules, map[string]interface{}{
			"service":  ingressRule.Service,
			"hostname": ingressRule.Hostname,
			"path":     ingressRule.Path,
		})
	}
	configMap = append(configMap, map[string]interface{}{
		"warp_routing":   warpConfig,
		"origin_request": originRequestConfig,
		"ingress_rule":   ingressRules,
	})
	d.Set("config", configMap)
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

func optBoolToString(b *bool) bool {
	if b != nil {
		fmt.Printf("optBoolToString %#v\n", *b)
		return *b
	}
	fmt.Println("optBoolToString was null")
	return false
}

func optIntToString(i *int) interface{} {
	if i != nil {
		return cloudflare.Int(i)
	}
	return nil
}
func optUintToString(u *uint) interface{} {
	if u != nil {
		return cloudflare.Uint(u)
	}
	return nil
}

func optDateDurationToString(d *time.Duration) string {
	if d != nil {
		return d.String()
	}
	return ""
}

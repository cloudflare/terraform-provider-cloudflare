package sdkv2provider

import (
	"context"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"
)

func resourceCloudflareTunnelConfig() *schema.Resource {
	return &schema.Resource{
		Schema:        resourceCloudflareTunnelConfigSchema(),
		ReadContext:   resourceCloudflareTunnelConfigRead,
		CreateContext: resourceCloudflareTunnelConfigUpdate,
		UpdateContext: resourceCloudflareTunnelConfigUpdate,
		DeleteContext: resourceCloudflareTunnelConfigDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceCloudflareTunnelConfigImport,
		},
		Description: heredoc.Doc(`
			Provides a Cloudflare Tunnel configuration resource.
		`),
	}
}

func buildTunnelOriginRequest(originRequest map[string]interface{}) (originConfig cloudflare.OriginRequestConfig) {
	if v, ok := originRequest["connect_timeout"]; ok {
		timeout, _ := time.ParseDuration(v.(string))
		originConfig.ConnectTimeout = &cloudflare.TunnelDuration{Duration: timeout}
	}
	if v, ok := originRequest["tls_timeout"]; ok {
		timeout, _ := time.ParseDuration(v.(string))
		originConfig.TLSTimeout = &cloudflare.TunnelDuration{Duration: timeout}
	}
	if v, ok := originRequest["tcp_keep_alive"]; ok {
		timeout, _ := time.ParseDuration(v.(string))
		originConfig.TCPKeepAlive = &cloudflare.TunnelDuration{Duration: timeout}
	}
	if v, ok := originRequest["no_happy_eyeballs"]; ok {
		originConfig.NoHappyEyeballs = cloudflare.BoolPtr(v.(bool))
	}
	if v, ok := originRequest["keep_alive_connections"]; ok {
		originConfig.KeepAliveConnections = cloudflare.IntPtr(v.(int))
	}
	if v, ok := originRequest["keep_alive_timeout"]; ok {
		timeout, _ := time.ParseDuration(v.(string))
		originConfig.KeepAliveTimeout = &cloudflare.TunnelDuration{Duration: timeout}
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
	if v, ok := originRequest["http2_origin"]; ok {
		originConfig.Http2Origin = cloudflare.BoolPtr(v.(bool))
	}
	var ipRules []cloudflare.IngressIPRule
	if v, ok := originRequest["ip_rules"]; ok {
		for _, ingressRule := range v.(*schema.Set).List() {
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
	if v, ok := originRequest["access"]; ok {
		if len(v.([]interface{})) != 0 {
			accessConfig := v.([]interface{})[0].(map[string]interface{})
			originConfig.Access = &cloudflare.AccessConfig{
				Required: accessConfig["required"].(bool),
				TeamName: accessConfig["team_name"].(string),
			}
			for _, value := range accessConfig["aud_tag"].(*schema.Set).List() {
				originConfig.Access.AudTag = append(originConfig.Access.AudTag, value.(string))
			}
		}
	}

	return
}

func parseOriginRequest(originRequest cloudflare.OriginRequestConfig) (returnValue []map[string]interface{}) {
	returnValue = append(returnValue, map[string]interface{}{
		"no_happy_eyeballs":        originRequest.NoHappyEyeballs,
		"keep_alive_connections":   originRequest.KeepAliveConnections,
		"http_host_header":         originRequest.HTTPHostHeader,
		"origin_server_name":       originRequest.OriginServerName,
		"ca_pool":                  originRequest.CAPool,
		"no_tls_verify":            originRequest.NoTLSVerify,
		"disable_chunked_encoding": originRequest.DisableChunkedEncoding,
		"bastion_mode":             originRequest.BastionMode,
		"proxy_address":            originRequest.ProxyAddress,
		"proxy_port":               originRequest.ProxyPort,
		"proxy_type":               originRequest.ProxyType,
		"http2_origin":             originRequest.Http2Origin,
	})
	if originRequest.ConnectTimeout != nil {
		returnValue[0]["connect_timeout"] = originRequest.ConnectTimeout.String()
	}
	if originRequest.TLSTimeout != nil {
		returnValue[0]["tls_timeout"] = originRequest.TLSTimeout.String()
	}
	if originRequest.TCPKeepAlive != nil {
		returnValue[0]["tcp_keep_alive"] = originRequest.TCPKeepAlive.String()
	}
	if originRequest.KeepAliveTimeout != nil {
		returnValue[0]["keep_alive_timeout"] = originRequest.KeepAliveTimeout.String()
	}
	var accessConfig []map[string]interface{}
	if originRequest.Access != nil {
		accessConfig = append(accessConfig, map[string]interface{}{
			"required":  originRequest.Access.Required,
			"team_name": originRequest.Access.TeamName,
			"aud_tag":   originRequest.Access.AudTag,
		})
	}
	returnValue[0]["access"] = accessConfig
	return
}

func buildTunnelConfig(d *schema.ResourceData) cloudflare.TunnelConfiguration {
	warpRouting := cloudflare.WarpRoutingConfig{}
	if item, ok := d.GetOk("config.0.warp_routing"); ok {
		warpRouting = cloudflare.WarpRoutingConfig{
			Enabled: item.([]interface{})[0].(map[string]interface{})["enabled"].(bool),
		}
	}

	originConfig := cloudflare.OriginRequestConfig{}
	if item, ok := d.GetOk("config.0.origin_request"); ok {
		originConfig = buildTunnelOriginRequest(item.([]interface{})[0].(map[string]interface{}))
	}

	var ingressRules []cloudflare.UnvalidatedIngressRule
	for _, ingressRule := range d.Get("config.0.ingress_rule").([]interface{}) {
		ingressRuleConfig := ingressRule.(map[string]interface{})
		ingressRule := cloudflare.UnvalidatedIngressRule{
			Service:  ingressRuleConfig["service"].(string),
			Hostname: ingressRuleConfig["hostname"].(string),
			Path:     ingressRuleConfig["path"].(string),
		}
		if v, ok := ingressRuleConfig["origin_request"]; ok {
			if len(v.([]interface{})) != 0 {
				ingressOriginConfig := buildTunnelOriginRequest(v.([]interface{})[0].(map[string]interface{}))
				ingressRule.OriginRequest = &ingressOriginConfig
			}
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
	accountID := d.Get(consts.AccountIDSchemaKey).(string)
	result, err := client.GetTunnelConfiguration(ctx, cloudflare.AccountIdentifier(accountID), d.Id())
	tflog.Debug(ctx, fmt.Sprintf("GetTunnelConfiguration: %+v", result))
	if err != nil {
		return diag.FromErr(fmt.Errorf("error getting tunnel config %q: %w", d.Id(), err))
	}

	d.SetId(result.TunnelID)
	config := result.Config
	var configMap []map[string]interface{}

	var warpConfigMap []map[string]interface{}

	emptyWarpRouting := cloudflare.WarpRoutingConfig{}
	if !reflect.DeepEqual(config.WarpRouting, &emptyWarpRouting) {
		warpConfigMap = append(warpConfigMap, map[string]interface{}{
			"enabled": config.WarpRouting.Enabled,
		})
	}

	var originRequestMap []map[string]interface{}
	emptyOriginRequest := cloudflare.OriginRequestConfig{}
	if !reflect.DeepEqual(config.OriginRequest, emptyOriginRequest) {
		originRequestMap = parseOriginRequest(config.OriginRequest)
		var ipRules []map[string]interface{}
		for _, ipRule := range config.OriginRequest.IPRules {
			ipRules = append(ipRules, map[string]interface{}{
				"prefix": ipRule.Prefix,
				"allow":  ipRule.Allow,
				"ports":  ipRule.Ports,
			})
		}
		originRequestMap[0]["ip_rules"] = ipRules
	}

	var ingressRules []map[string]interface{}
	for _, ingressRule := range config.Ingress {
		var rule = make(map[string]interface{})
		rule["service"] = ingressRule.Service
		rule["hostname"] = ingressRule.Hostname
		rule["path"] = ingressRule.Path
		var ingressOriginRequestMap []map[string]interface{}
		if ingressRule.OriginRequest != nil {
			ingressOriginRequestMap = parseOriginRequest(*ingressRule.OriginRequest)
			rule["origin_request"] = ingressOriginRequestMap
		}
		ingressRules = append(ingressRules, rule)
	}
	configMap = append(configMap, map[string]interface{}{
		"warp_routing":   warpConfigMap,
		"origin_request": originRequestMap,
		"ingress_rule":   ingressRules,
	})
	d.Set("config", configMap)
	return nil
}

func resourceCloudflareTunnelConfigUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := d.Get(consts.AccountIDSchemaKey).(string)
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
	accountID := d.Get(consts.AccountIDSchemaKey).(string)
	tunnelID := d.Get("tunnel_id").(string)

	// can't delete a tunnel config, so set an "empty" config instead
	tunnel := cloudflare.TunnelConfigurationParams{
		TunnelID: tunnelID,
		Config: cloudflare.TunnelConfiguration{
			OriginRequest: cloudflare.OriginRequestConfig{},
			WarpRouting: &cloudflare.WarpRoutingConfig{
				Enabled: false,
			},
			Ingress: []cloudflare.UnvalidatedIngressRule{
				{
					Service: "http_status:404",
				},
			},
		},
	}

	_, err := client.UpdateTunnelConfiguration(ctx, cloudflare.AccountIdentifier(accountID), tunnel)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error clearing tunnel config %q: %w", tunnelID, err))
	}

	d.SetId("")

	return nil
}

func resourceCloudflareTunnelConfigImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	client := meta.(*cloudflare.API)
	attributes := strings.Split(d.Id(), "/")

	if len(attributes) != 2 {
		return nil, fmt.Errorf("invalid id (\"%s\") specified, should be in format \"accountID/tunnelId\"", d.Id())
	}

	accountId, tunnelId := attributes[0], attributes[1]

	result, err := client.GetTunnelConfiguration(ctx, cloudflare.AccountIdentifier(accountId), tunnelId)
	tflog.Debug(ctx, fmt.Sprintf("GetTunnelConfiguration: %+v", result))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("failed to fetch Cloudflare Tunnel configuration %s", tunnelId))
	}

	d.Set(consts.AccountIDSchemaKey, accountId)
	d.Set("tunnel_id", result.TunnelID)
	d.SetId(result.TunnelID)

	resourceCloudflareTunnelConfigRead(ctx, d, meta)

	return []*schema.ResourceData{d}, nil
}

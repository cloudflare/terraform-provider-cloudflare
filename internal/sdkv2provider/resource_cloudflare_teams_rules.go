package sdkv2provider

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflareTeamsRule() *schema.Resource {
	return &schema.Resource{
		Schema:        resourceCloudflareTeamsRuleSchema(),
		ReadContext:   resourceCloudflareTeamsRuleRead,
		UpdateContext: resourceCloudflareTeamsRuleUpdate,
		CreateContext: resourceCloudflareTeamsRuleCreate,
		DeleteContext: resourceCloudflareTeamsRuleDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceCloudflareTeamsRuleImport,
		},
		Description: "Provides a Cloudflare Teams rule resource. Teams rules comprise secure web gateway policies.",
	}
}

const rulePrecedenceFactor int64 = 1000

func resourceCloudflareTeamsRuleRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := d.Get(consts.AccountIDSchemaKey).(string)

	rule, err := client.TeamsRule(ctx, accountID, d.Id())
	if err != nil {
		if strings.Contains(err.Error(), "invalid rule id") {
			tflog.Info(ctx, fmt.Sprintf("Teams Rule config %s does not exists", d.Id()))
			d.SetId("")
			return nil
		}
		return diag.FromErr(fmt.Errorf("error finding Teams Rule %q: %w", d.Id(), err))
	}
	if err := d.Set("name", rule.Name); err != nil {
		return diag.FromErr(fmt.Errorf("error parsing rule name"))
	}
	if err := d.Set("description", rule.Description); err != nil {
		return diag.FromErr(fmt.Errorf("error parsing rule description"))
	}
	if err := d.Set("precedence", apiToProviderRulePrecedence(rule.Precedence, rule.Name)); err != nil {
		return diag.FromErr(fmt.Errorf("error parsing rule precedence"))
	}
	if err := d.Set("enabled", rule.Enabled); err != nil {
		return diag.FromErr(fmt.Errorf("error parsing rule enablement"))
	}
	if err := d.Set("action", rule.Action); err != nil {
		return diag.FromErr(fmt.Errorf("error parsing rule action"))
	}
	if err := d.Set("filters", rule.Filters); err != nil {
		return diag.FromErr(fmt.Errorf("error parsing rule filters"))
	}
	if err := d.Set("traffic", rule.Traffic); err != nil {
		return diag.FromErr(fmt.Errorf("error parsing rule traffic"))
	}
	if err := d.Set("identity", rule.Identity); err != nil {
		return diag.FromErr(fmt.Errorf("error parsing rule identity"))
	}
	if err := d.Set("device_posture", rule.DevicePosture); err != nil {
		return diag.FromErr(fmt.Errorf("error parsing rule device_posture"))
	}
	if err := d.Set("version", int64(rule.Version)); err != nil {
		return diag.FromErr(fmt.Errorf("error parsing rule version"))
	}

	if err := d.Set("rule_settings", flattenTeamsRuleSettings(d, &rule.RuleSettings)); err != nil {
		return diag.FromErr(fmt.Errorf("error parsing rule settings"))
	}

	return nil
}

func resourceCloudflareTeamsRuleCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)

	accountID := d.Get(consts.AccountIDSchemaKey).(string)
	settings := inflateTeamsRuleSettings(d.Get("rule_settings"))

	var filters []cloudflare.TeamsFilterType
	for _, f := range d.Get("filters").([]interface{}) {
		filters = append(filters, cloudflare.TeamsFilterType(f.(string)))
	}

	ruleName := d.Get("name").(string)
	apiPrecedence := providerToApiRulePrecedence(int64(d.Get("precedence").(int)), ruleName)
	newTeamsRule := cloudflare.TeamsRule{
		Name:          ruleName,
		Description:   d.Get("description").(string),
		Precedence:    uint64(apiPrecedence),
		Enabled:       d.Get("enabled").(bool),
		Action:        cloudflare.TeamsGatewayAction(d.Get("action").(string)),
		Filters:       filters,
		Traffic:       d.Get("traffic").(string),
		Identity:      d.Get("identity").(string),
		DevicePosture: d.Get("device_posture").(string),
		Version:       uint64(d.Get("version").(int)),
	}

	if settings != nil {
		newTeamsRule.RuleSettings = *settings
	}

	tflog.Debug(ctx, fmt.Sprintf("Creating Cloudflare Teams Rule from struct: %+v", newTeamsRule))

	rule, err := client.TeamsCreateRule(ctx, accountID, newTeamsRule)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error creating Teams rule for account %q: %w", accountID, err))
	}

	d.SetId(rule.ID)
	return resourceCloudflareTeamsRuleRead(ctx, d, meta)
}

func resourceCloudflareTeamsRuleUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := d.Get(consts.AccountIDSchemaKey).(string)
	settings := inflateTeamsRuleSettings(d.Get("rule_settings"))

	var filters []cloudflare.TeamsFilterType
	for _, f := range d.Get("filters").([]interface{}) {
		filters = append(filters, cloudflare.TeamsFilterType(f.(string)))
	}

	ruleName := d.Get("name").(string)
	apiPrecedence := providerToApiRulePrecedence(int64(d.Get("precedence").(int)), ruleName)
	teamsRule := cloudflare.TeamsRule{
		ID:            d.Id(),
		Name:          ruleName,
		Description:   d.Get("description").(string),
		Precedence:    uint64(apiPrecedence),
		Enabled:       d.Get("enabled").(bool),
		Action:        cloudflare.TeamsGatewayAction(d.Get("action").(string)),
		Filters:       filters,
		Traffic:       d.Get("traffic").(string),
		Identity:      d.Get("identity").(string),
		DevicePosture: d.Get("device_posture").(string),
		Version:       uint64(d.Get("version").(int)),
	}

	if settings != nil {
		teamsRule.RuleSettings = *settings
	}
	tflog.Debug(ctx, fmt.Sprintf("Updating Cloudflare Teams rule from struct: %+v", teamsRule))

	updatedTeamsRule, err := client.TeamsUpdateRule(ctx, accountID, teamsRule.ID, teamsRule)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error updating Teams rule for account %q: %w", accountID, err))
	}
	if updatedTeamsRule.ID == "" {
		return diag.FromErr(fmt.Errorf("failed to find Teams Rule ID in update response; resource was empty"))
	}
	return resourceCloudflareTeamsRuleRead(ctx, d, meta)
}

func resourceCloudflareTeamsRuleDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	id := d.Id()
	accountID := d.Get(consts.AccountIDSchemaKey).(string)

	tflog.Debug(ctx, fmt.Sprintf("Deleting Cloudflare Teams Rule using ID: %s", id))

	err := client.TeamsDeleteRule(ctx, accountID, id)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error deleting Teams Rule for account %q: %w", accountID, err))
	}

	return resourceCloudflareTeamsRuleRead(ctx, d, meta)
}

func resourceCloudflareTeamsRuleImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	attributes := strings.SplitN(d.Id(), "/", 2)

	if len(attributes) != 2 {
		return nil, fmt.Errorf("invalid id (\"%s\") specified, should be in format \"accountID/teamsRuleID\"", d.Id())
	}

	accountID, teamsRuleID := attributes[0], attributes[1]

	tflog.Debug(ctx, fmt.Sprintf("Importing Cloudflare Teams Rule: id %s for account %s", teamsRuleID, accountID))

	d.Set(consts.AccountIDSchemaKey, accountID)
	d.SetId(teamsRuleID)

	resourceCloudflareTeamsRuleRead(ctx, d, meta)

	return []*schema.ResourceData{d}, nil
}

func flattenTeamsRuleSettings(d *schema.ResourceData, settings *cloudflare.TeamsRuleSettings) []interface{} {
	if _, ok := d.GetOkExists("block_page_enabled"); !ok &&
		len(settings.OverrideIPs) == 0 &&
		settings.BlockReason == "" &&
		settings.OverrideHost == "" &&
		settings.BISOAdminControls == nil &&
		settings.L4Override == nil &&
		len(settings.AddHeaders) == 0 &&
		settings.CheckSession == nil &&
		settings.InsecureDisableDNSSECValidation == false &&
		settings.EgressSettings == nil &&
		settings.UntrustedCertSettings == nil &&
		settings.PayloadLog == nil &&
		settings.IPCategories == false &&
		settings.AllowChildBypass == nil &&
		settings.BypassParentRule == nil &&
		settings.AuditSSH == nil &&
		settings.NotificationSettings == nil &&
		settings.ResolveDnsThroughCloudflare == nil &&
		settings.IgnoreCNAMECategoryMatches == nil &&
		settings.DnsResolverSettings == nil {
		return nil
	}

	result := map[string]interface{}{
		"block_page_enabled":                 settings.BlockPageEnabled,
		"block_page_reason":                  settings.BlockReason,
		"override_ips":                       settings.OverrideIPs,
		"override_host":                      settings.OverrideHost,
		"l4override":                         flattenTeamsL4Override(settings.L4Override),
		"biso_admin_controls":                flattenTeamsRuleBisoAdminControls(settings.BISOAdminControls),
		"check_session":                      flattenTeamsCheckSessionSettings(settings.CheckSession),
		"add_headers":                        flattenTeamsAddHeaders(settings.AddHeaders),
		"insecure_disable_dnssec_validation": settings.InsecureDisableDNSSECValidation,
		"egress":                             flattenTeamsEgressSettings(settings.EgressSettings),
		"untrusted_cert":                     flattenTeamsUntrustedCertSettings(settings.UntrustedCertSettings),
		"payload_log":                        flattenTeamsDlpPayloadLogSettings(settings.PayloadLog),
		"notification_settings":              flattenTeamsNotificationSettings(settings.NotificationSettings),
	}

	if settings.IPCategories {
		result["ip_categories"] = true
	}

	if settings.IgnoreCNAMECategoryMatches != nil {
		result["ignore_cname_category_matches"] = *settings.IgnoreCNAMECategoryMatches
	}

	if settings.ResolveDnsThroughCloudflare != nil {
		result["resolve_dns_through_cloudflare"] = *settings.ResolveDnsThroughCloudflare
	}

	if settings.AllowChildBypass != nil {
		result["allow_child_bypass"] = *settings.AllowChildBypass
	}

	if settings.BypassParentRule != nil {
		result["bypass_parent_rule"] = *settings.BypassParentRule
	}

	if settings.AuditSSH != nil {
		result["audit_ssh"] = flattenTeamsAuditSSHSettings(settings.AuditSSH)
	}

	if settings.DnsResolverSettings != nil {
		result["dns_resolvers"] = flattenTeamsDnsResolverSettings(settings.DnsResolverSettings)
	}

	return []interface{}{result}
}

func inflateTeamsRuleSettings(settings interface{}) *cloudflare.TeamsRuleSettings {
	settingsList := settings.([]interface{})
	if len(settingsList) != 1 {
		return nil
	}
	settingsMap := settingsList[0].(map[string]interface{})
	enabled := settingsMap["block_page_enabled"].(bool)
	reason := settingsMap["block_page_reason"].(string)

	var overrideIPs []string
	for _, ip := range settingsMap["override_ips"].([]interface{}) {
		overrideIPs = append(overrideIPs, ip.(string))
	}

	overrideHost := settingsMap["override_host"].(string)
	l4Override := inflateTeamsL4Override(settingsMap["l4override"].([]interface{}))

	bisoAdminControls := inflateTeamsRuleBisoAdminControls(settingsMap["biso_admin_controls"].([]interface{}))

	checkSessionSettings := inflateTeamsCheckSessionSettings(settingsMap["check_session"].([]interface{}))
	addHeaders := inflateTeamsAddHeaders(settingsMap["add_headers"].(map[string]interface{}))
	insecureDisableDNSSECValidation := settingsMap["insecure_disable_dnssec_validation"].(bool)
	egressSettings := inflateTeamsEgressSettings(settingsMap["egress"].([]interface{}))
	payloadLog := inflateTeamsDlpPayloadLogSettings(settingsMap["payload_log"].([]interface{}))
	untrustedCertSettings := inflateTeamsUntrustedCertSettings(settingsMap["untrusted_cert"].([]interface{}))
	notificationSettings := inflateTeamsNotificationSettings(settingsMap["notification_settings"])
	dnsResolverSettings := inflateTeamsDnsResolverSettings(settingsMap["dns_resolvers"].([]interface{}))

	ignoreCNAMECategoryMatches := readOptionalBooleanSettings(settingsMap, "ignore_cname_category_matches")
	allowChildBypass := readOptionalBooleanSettings(settingsMap, "allow_child_bypass")
	bypassParentRule := readOptionalBooleanSettings(settingsMap, "bypass_parent_rule")
	resolveDnsThroughCloudflare := readOptionalBooleanSettings(settingsMap, "resolve_dns_through_cloudflare")
	auditSSHSettings := inflateTeamsAuditSSHSettings(settingsMap["audit_ssh"].([]interface{}))
	ipCategories := readOptionalBooleanSettings(settingsMap, "ip_categories")

	result := &cloudflare.TeamsRuleSettings{
		BlockPageEnabled:                enabled,
		BlockReason:                     reason,
		OverrideIPs:                     overrideIPs,
		OverrideHost:                    overrideHost,
		L4Override:                      l4Override,
		BISOAdminControls:               bisoAdminControls,
		CheckSession:                    checkSessionSettings,
		AddHeaders:                      addHeaders,
		InsecureDisableDNSSECValidation: insecureDisableDNSSECValidation,
		EgressSettings:                  egressSettings,
		PayloadLog:                      payloadLog,
		UntrustedCertSettings:           untrustedCertSettings,
		NotificationSettings:            notificationSettings,
		ResolveDnsThroughCloudflare:     &resolveDnsThroughCloudflare,
		DnsResolverSettings:             dnsResolverSettings,
		IgnoreCNAMECategoryMatches:      &ignoreCNAMECategoryMatches,
		IPCategories:                    ipCategories,
		AuditSSH:                        auditSSHSettings,
	}

	// set optional settings if present, so api won't complain
	if allowChildBypass {
		result.AllowChildBypass = &allowChildBypass
	}
	if bypassParentRule {
		result.BypassParentRule = &bypassParentRule
	}

	return result
}

func flattenTeamsRuleBisoAdminControls(settings *cloudflare.TeamsBISOAdminControlSettings) []interface{} {
	if settings == nil {
		return nil
	}
	return []interface{}{map[string]interface{}{
		"disable_printing":   settings.DisablePrinting,
		"disable_copy_paste": settings.DisableCopyPaste,
		"disable_download":   settings.DisableDownload,
		"disable_upload":     settings.DisableUpload,
		"disable_keyboard":   settings.DisableKeyboard,
	}}
}

func flattenTeamsCheckSessionSettings(settings *cloudflare.TeamsCheckSessionSettings) []interface{} {
	if settings == nil {
		return nil
	}
	duration := settings.Duration.Duration.String()
	return []interface{}{map[string]interface{}{
		"enforce":  settings.Enforce,
		"duration": duration,
	}}
}

func inflateTeamsRuleBisoAdminControls(settings interface{}) *cloudflare.TeamsBISOAdminControlSettings {
	settingsList := settings.([]interface{})
	if len(settingsList) != 1 {
		return nil
	}
	settingsMap := settingsList[0].(map[string]interface{})
	disablePrinting := settingsMap["disable_printing"].(bool)
	disableCopyPaste := settingsMap["disable_copy_paste"].(bool)
	disableDownload := settingsMap["disable_download"].(bool)
	disableUpload := settingsMap["disable_upload"].(bool)
	disableKeyboard := settingsMap["disable_keyboard"].(bool)
	return &cloudflare.TeamsBISOAdminControlSettings{
		DisablePrinting:  disablePrinting,
		DisableCopyPaste: disableCopyPaste,
		DisableDownload:  disableDownload,
		DisableUpload:    disableUpload,
		DisableKeyboard:  disableKeyboard,
	}
}

func inflateTeamsCheckSessionSettings(settings interface{}) *cloudflare.TeamsCheckSessionSettings {
	settingsList := settings.([]interface{})
	if len(settingsList) != 1 {
		return nil
	}
	settingsMap := settingsList[0].(map[string]interface{})
	enforce := settingsMap["enforce"].(bool)
	durationString := settingsMap["duration"].(string)

	dur, err := time.ParseDuration(durationString)
	if err != nil {
		return nil
	}

	duration := cloudflare.Duration{Duration: dur}

	return &cloudflare.TeamsCheckSessionSettings{
		Enforce:  enforce,
		Duration: duration,
	}
}

func inflateTeamsAddHeaders(settings interface{}) http.Header {
	settingsMap := settings.(map[string]interface{})

	ret := http.Header{}
	for key := range settingsMap {
		v, ok := settingsMap[key].(string)
		if !ok {
			continue
		}
		s := strings.Split(v, ",")
		ret[key] = s
	}

	return ret
}

func flattenTeamsAddHeaders(settings http.Header) interface{} {
	if settings == nil {
		return nil
	}

	ret := make(map[string]interface{})
	for name, value := range settings {
		ret[name] = strings.Join(value[:], ",")
	}

	return ret
}

func flattenTeamsL4Override(settings *cloudflare.TeamsL4OverrideSettings) []interface{} {
	if settings == nil {
		return nil
	}
	return []interface{}{map[string]interface{}{
		"ip":   settings.IP,
		"port": settings.Port,
	}}
}

func inflateTeamsL4Override(settings interface{}) *cloudflare.TeamsL4OverrideSettings {
	settingsList := settings.([]interface{})
	if len(settingsList) != 1 {
		return nil
	}
	settingsMap := settingsList[0].(map[string]interface{})
	ip := settingsMap["ip"].(string)
	port := settingsMap["port"].(int)
	return &cloudflare.TeamsL4OverrideSettings{
		IP:   ip,
		Port: port,
	}
}

func flattenTeamsAuditSSHSettings(settings *cloudflare.AuditSSHRuleSettings) []interface{} {
	if settings == nil {
		return nil
	}
	return []interface{}{map[string]interface{}{
		"command_logging": settings.CommandLogging,
	}}
}

func flattenTeamsEgressSettings(settings *cloudflare.EgressSettings) []interface{} {
	if settings == nil {
		return nil
	}
	return []interface{}{map[string]interface{}{
		"ipv4":          settings.Ipv4,
		"ipv6":          settings.Ipv6Range,
		"ipv4_fallback": settings.Ipv4Fallback,
	}}
}

func inflateTeamsEgressSettings(settings interface{}) *cloudflare.EgressSettings {
	settingsList := settings.([]interface{})
	if len(settingsList) != 1 {
		return nil
	}
	settingsMap := settingsList[0].(map[string]interface{})
	ipv4 := settingsMap["ipv4"].(string)
	ipv6 := settingsMap["ipv6"].(string)
	ipv4Fallback := settingsMap["ipv4_fallback"].(string)
	return &cloudflare.EgressSettings{
		Ipv4:         ipv4,
		Ipv6Range:    ipv6,
		Ipv4Fallback: ipv4Fallback,
	}
}

func flattenTeamsDlpPayloadLogSettings(settings *cloudflare.TeamsDlpPayloadLogSettings) []interface{} {
	if settings == nil {
		return nil
	}
	return []interface{}{map[string]interface{}{
		"enabled": settings.Enabled,
	}}
}

func flattenTeamsNotificationSettings(settings *cloudflare.TeamsNotificationSettings) []interface{} {
	if settings == nil {
		return nil
	}
	enabled := false
	if settings.Enabled != nil {
		enabled = *settings.Enabled
	}

	return []interface{}{map[string]interface{}{
		"enabled":     enabled,
		"message":     settings.Message,
		"support_url": settings.SupportURL,
	}}
}

func flattenTeamsUntrustedCertSettings(settings *cloudflare.UntrustedCertSettings) []interface{} {
	if settings == nil {
		return nil
	}
	return []interface{}{map[string]interface{}{
		"action": settings.Action,
	}}
}

func flattenTeamsDnsResolverSettings(settings *cloudflare.TeamsDnsResolverSettings) []interface{} {
	if settings == nil {
		return nil
	}
	ret := map[string]interface{}{}
	if len(settings.V4Resolvers) != 0 {
		ret["ipv4"] = flattenTeamsDnsResolverAddressesV4(settings.V4Resolvers)
	}
	if len(settings.V6Resolvers) != 0 {
		ret["ipv6"] = flattenTeamsDnsResolverAddressesV6(settings.V6Resolvers)
	}
	return []interface{}{ret}
}

func inflateTeamsDnsResolverSettings(settings interface{}) *cloudflare.TeamsDnsResolverSettings {
	settingsList := settings.([]interface{})
	if len(settingsList) != 1 {
		return nil
	}
	settingsMap := settingsList[0].(map[string]interface{})
	ipv4 := inflateTeamsDnsResolverAddressesV4(settingsMap["ipv4"].([]interface{}))
	ipv6 := inflateTeamsDnsResolverAddressesV6(settingsMap["ipv6"].([]interface{}))
	return &cloudflare.TeamsDnsResolverSettings{
		V4Resolvers: ipv4,
		V6Resolvers: ipv6,
	}
}

func flattenTeamsDnsResolverAddressesV4(settings []cloudflare.TeamsDnsResolverAddressV4) []interface{} {
	ret := []interface{}{}
	for _, addr := range settings {
		part := map[string]interface{}{
			"ip":                            addr.IP,
			"port":                          addr.Port,
			"route_through_private_network": addr.RouteThroughPrivateNetwork,
			"vnet_id":                       addr.VnetID,
		}
		ret = append(ret, part)
	}
	return ret
}

func inflateTeamsDnsResolverAddressesV4(settings []interface{}) []cloudflare.TeamsDnsResolverAddressV4 {
	ret := []cloudflare.TeamsDnsResolverAddressV4{}
	for _, raw := range settings {
		part := raw.(map[string]interface{})
		newAddr := cloudflare.TeamsDnsResolverAddressV4{
			TeamsDnsResolverAddress: cloudflare.TeamsDnsResolverAddress{
				IP:                         part["ip"].(string),
				Port:                       cloudflare.IntPtr(part["port"].(int)),
				VnetID:                     part["vnet_id"].(string),
				RouteThroughPrivateNetwork: cloudflare.BoolPtr(part["route_through_private_network"].(bool)),
			},
		}
		ret = append(ret, newAddr)
	}
	return ret
}

func flattenTeamsDnsResolverAddressesV6(settings []cloudflare.TeamsDnsResolverAddressV6) []interface{} {
	ret := []interface{}{}
	for _, addr := range settings {
		part := map[string]interface{}{
			"ip":                            addr.IP,
			"port":                          addr.Port,
			"route_through_private_network": addr.RouteThroughPrivateNetwork,
			"vnet_id":                       addr.VnetID,
		}
		ret = append(ret, part)
	}
	return ret
}

func inflateTeamsDnsResolverAddressesV6(settings []interface{}) []cloudflare.TeamsDnsResolverAddressV6 {
	ret := []cloudflare.TeamsDnsResolverAddressV6{}
	for _, raw := range settings {
		part := raw.(map[string]interface{})
		newAddr := cloudflare.TeamsDnsResolverAddressV6{
			TeamsDnsResolverAddress: cloudflare.TeamsDnsResolverAddress{
				IP:                         part["ip"].(string),
				Port:                       cloudflare.IntPtr(part["port"].(int)),
				VnetID:                     part["vnet_id"].(string),
				RouteThroughPrivateNetwork: cloudflare.BoolPtr(part["route_through_private_network"].(bool)),
			},
		}
		ret = append(ret, newAddr)
	}
	return ret
}

func inflateTeamsDlpPayloadLogSettings(settings interface{}) *cloudflare.TeamsDlpPayloadLogSettings {
	settingsList := settings.([]interface{})
	if len(settingsList) != 1 {
		return nil
	}
	settingsMap := settingsList[0].(map[string]interface{})
	enabled := settingsMap["enabled"].(bool)
	return &cloudflare.TeamsDlpPayloadLogSettings{
		Enabled: enabled,
	}
}

func inflateTeamsUntrustedCertSettings(settings interface{}) *cloudflare.UntrustedCertSettings {
	settingsList := settings.([]interface{})
	if len(settingsList) != 1 {
		return nil
	}
	settingsMap := settingsList[0].(map[string]interface{})
	action := settingsMap["action"].(string)
	var actionValue cloudflare.TeamsGatewayUntrustedCertAction
	switch action {
	case "pass_through":
		actionValue = cloudflare.UntrustedCertPassthrough
	case "block":
		actionValue = cloudflare.UntrustedCertBlock
	case "error":
		actionValue = cloudflare.UntrustedCertError
	}

	return &cloudflare.UntrustedCertSettings{
		Action: actionValue,
	}
}

func inflateTeamsNotificationSettings(settings interface{}) *cloudflare.TeamsNotificationSettings {
	settingsList := settings.([]interface{})
	if len(settingsList) != 1 {
		return nil
	}
	settingsMap := settingsList[0].(map[string]interface{})
	enabled := settingsMap["enabled"].(bool)
	message := settingsMap["message"].(string)
	supportUrl := settingsMap["support_url"].(string)
	return &cloudflare.TeamsNotificationSettings{
		Enabled:    &enabled,
		Message:    message,
		SupportURL: supportUrl,
	}
}

func providerToApiRulePrecedence(provided int64, ruleName string) int64 {
	return provided*rulePrecedenceFactor + int64(hashCodeString(ruleName))%rulePrecedenceFactor
}

func apiToProviderRulePrecedence(apiPrecedence uint64, ruleName string) int64 {
	return (int64(apiPrecedence) - int64(hashCodeString(ruleName))%rulePrecedenceFactor) / rulePrecedenceFactor
}

func readOptionalBooleanSettings(settingsMap map[string]any, name string) bool {
	val, ok := settingsMap[name]
	if !ok {
		return false
	}
	return val.(bool)
}

func inflateTeamsAuditSSHSettings(settings interface{}) *cloudflare.AuditSSHRuleSettings {
	settingsList := settings.([]interface{})
	if len(settingsList) != 1 {
		return nil
	}
	settingsMap := settingsList[0].(map[string]interface{})
	logging := settingsMap["command_logging"].(bool)
	return &cloudflare.AuditSSHRuleSettings{
		CommandLogging: logging,
	}
}

package page_rule

import (
	"encoding/json"

	"github.com/cloudflare/cloudflare-go/v3/pagerules"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (m PageRuleModel) marshalCustom() (data []byte, err error) {
	if data, err = apijson.MarshalRoot(m); err != nil {
		return
	}
	if data, err = m.marshalTargetsAndActions(data); err != nil {
		return
	}
	return
}

func (m PageRuleModel) marshalCustomForUpdate(state PageRuleModel) (data []byte, err error) {
	if data, err = apijson.MarshalForUpdate(m, state); err != nil {
		return
	}
	if data, err = m.marshalTargetsAndActions(data); err != nil {
		return
	}
	return
}

func (m PageRuleModel) marshalTargetsAndActions(b []byte) (data []byte, err error) {
	var T struct {
		ID         string `json:"id,omitempty"`
		ZoneID     string `json:"zone_id,omitempty"`
		Priority   int64  `json:"priority,omitempty"`
		Status     string `json:"status,omitempty"`
		CreatedOn  string `json:"created_on,omitempty"`
		ModifiedOn string `json:"modified_on,omitempty"`
		Target     string `json:"target,omitempty"`
		Targets    []any  `json:"targets,omitempty"`
		Actions    any    `json:"actions,omitempty"`
	}
	if err = json.Unmarshal(b, &T); err != nil {
		return nil, err
	}

	T.Targets = []any{
		map[string]any{
			"target": "url",
			"constraint": map[string]any{
				"operator": "matches",
				"value":    T.Target,
			},
		},
	}
	T.Target = "" // omitempty

	T.Actions, err = m.Actions.Encode()
	if err != nil {
		return nil, err
	}

	return json.Marshal(T)
}

type PageRuleActionsModel struct {
	AlwaysUseHTTPS          types.Bool   `tfsdk:"always_use_https" json:"always_use_https,optional"`
	AutomaticHTTPSRewrites  types.String `tfsdk:"automatic_https_rewrites" json:"automatic_https_rewrites,optional"`
	BrowserCheck            types.String `tfsdk:"browser_check" json:"browser_check,optional"`
	BypassCacheOnCookie     types.String `tfsdk:"bypass_cache_on_cookie" json:"bypass_cache_on_cookie,optional"`
	CacheByDeviceType       types.String `tfsdk:"cache_by_device_type" json:"cache_by_device_type,optional"`
	CacheDeceptionArmor     types.String `tfsdk:"cache_deception_armor" json:"cache_deception_armor,optional"`
	CacheLevel              types.String `tfsdk:"cache_level" json:"cache_level,optional"`
	CacheOnCookie           types.String `tfsdk:"cache_on_cookie" json:"cache_on_cookie,optional"`
	DisableApps             types.Bool   `tfsdk:"disable_apps" json:"disable_apps,optional"`
	DisablePerformance      types.Bool   `tfsdk:"disable_performance" json:"disable_performance,optional"`
	DisableSecurity         types.Bool   `tfsdk:"disable_security" json:"disable_security,optional"`
	DisableZaraz            types.Bool   `tfsdk:"disable_zaraz" json:"disable_zaraz,optional"`
	EmailObfuscation        types.String `tfsdk:"email_obfuscation" json:"email_obfuscation,optional"`
	ExplicitCacheControl    types.String `tfsdk:"explicit_cache_control" json:"explicit_cache_control,optional"`
	HostHeaderOverride      types.String `tfsdk:"host_header_override" json:"host_header_override,optional"`
	IPGeolocation           types.String `tfsdk:"ip_geolocation" json:"ip_geolocation,optional"`
	Mirage                  types.String `tfsdk:"mirage" json:"mirage,optional"`
	OpportunisticEncryption types.String `tfsdk:"opportunistic_encryption" json:"opportunistic_encryption,optional"`
	OriginErrorPagePassThru types.String `tfsdk:"origin_error_page_pass_thru" json:"origin_error_page_pass_thru,optional"`
	Polish                  types.String `tfsdk:"polish" json:"polish,optional"`
	ResolveOverride         types.String `tfsdk:"resolve_override" json:"resolve_override,optional"`
	RespectStrongEtag       types.String `tfsdk:"respect_strong_etag" json:"respect_strong_etag,optional"`
	ResponseBuffering       types.String `tfsdk:"response_buffering" json:"response_buffering,optional"`
	RocketLoader            types.String `tfsdk:"rocket_loader" json:"rocket_loader,optional"`
	SSL                     types.String `tfsdk:"ssl" json:"ssl,optional"`
	SecurityLevel           types.String `tfsdk:"security_level" json:"security_level,optional"`
	SortQueryStringForCache types.String `tfsdk:"sort_query_string_for_cache" json:"sort_query_string_for_cache,optional"`
	TrueClientIPHeader      types.String `tfsdk:"true_client_ip_header" json:"true_client_ip_header,optional"`
	WAF                     types.String `tfsdk:"waf" json:"waf,optional"`
}

func (m *PageRuleActionsModel) Encode() (encoded []map[string]any, err error) {
	encoded = []map[string]any{}
	if m.AlwaysUseHTTPS.ValueBool() {
		encoded = append(encoded, map[string]any{"id": pagerules.PageRuleActionsIDAlwaysUseHTTPS})
	}
	if !m.AutomaticHTTPSRewrites.IsNull() {
		encoded = append(encoded, map[string]any{"id": pagerules.PageRuleActionsIDAutomaticHTTPSRewrites, "value": m.AutomaticHTTPSRewrites.String()})
	}
	if !m.BrowserCheck.IsNull() {
		encoded = append(encoded, map[string]any{"id": pagerules.PageRuleActionsIDBrowserCheck, "value": m.BrowserCheck.ValueString()})
	}
	if !m.BypassCacheOnCookie.IsNull() {
		encoded = append(encoded, map[string]any{"id": pagerules.PageRuleActionsIDBypassCacheOnCookie, "value": m.BypassCacheOnCookie.ValueString()})
	}
	if !m.CacheByDeviceType.IsNull() {
		encoded = append(encoded, map[string]any{"id": pagerules.PageRuleActionsIDCacheByDeviceType, "value": m.CacheByDeviceType.ValueString()})
	}
	if !m.CacheDeceptionArmor.IsNull() {
		encoded = append(encoded, map[string]any{"id": pagerules.PageRuleActionsIDCacheDeceptionArmor, "value": m.CacheDeceptionArmor.ValueString()})
	}
	if !m.CacheLevel.IsNull() {
		encoded = append(encoded, map[string]any{"id": pagerules.PageRuleActionsIDCacheLevel, "value": m.CacheLevel.ValueString()})
	}
	if !m.CacheOnCookie.IsNull() {
		encoded = append(encoded, map[string]any{"id": pagerules.PageRuleActionsIDCacheOnCookie, "value": m.CacheOnCookie.ValueString()})
	}
	if m.DisableApps.ValueBool() {
		encoded = append(encoded, map[string]any{"id": pagerules.PageRuleActionsIDDisableApps, "value": m.DisableApps.ValueBool()})
	}
	if m.DisablePerformance.ValueBool() {
		encoded = append(encoded, map[string]any{"id": pagerules.PageRuleActionsIDDisablePerformance})
	}
	if m.DisableSecurity.ValueBool() {
		encoded = append(encoded, map[string]any{"id": pagerules.PageRuleActionsIDDisableSecurity})
	}
	if m.DisableZaraz.ValueBool() {
		encoded = append(encoded, map[string]any{"id": pagerules.PageRuleActionsIDDisableZaraz})
	}
	if !m.EmailObfuscation.IsNull() {
		encoded = append(encoded, map[string]any{"id": pagerules.PageRuleActionsIDEmailObfuscation, "value": m.EmailObfuscation.ValueString()})
	}
	if !m.ExplicitCacheControl.IsNull() {
		encoded = append(encoded, map[string]any{"id": pagerules.PageRuleActionsIDExplicitCacheControl, "value": m.ExplicitCacheControl.ValueString()})
	}
	if !m.HostHeaderOverride.IsNull() {
		encoded = append(encoded, map[string]any{"id": pagerules.PageRuleActionsIDHostHeaderOverride, "value": m.HostHeaderOverride.ValueString()})
	}
	if !m.IPGeolocation.IsNull() {
		encoded = append(encoded, map[string]any{"id": pagerules.PageRuleActionsIDIPGeolocation, "value": m.IPGeolocation.ValueString()})
	}
	if !m.Mirage.IsNull() {
		encoded = append(encoded, map[string]any{"id": pagerules.PageRuleActionsIDMirage, "value": m.Mirage.ValueString()})
	}
	if !m.OpportunisticEncryption.IsNull() {
		encoded = append(encoded, map[string]any{"id": pagerules.PageRuleActionsIDOpportunisticEncryption, "value": m.OpportunisticEncryption.ValueString()})
	}
	if !m.OriginErrorPagePassThru.IsNull() {
		encoded = append(encoded, map[string]any{"id": pagerules.PageRuleActionsIDOriginErrorPagePassThru, "value": m.OriginErrorPagePassThru.ValueString()})
	}
	if !m.Polish.IsNull() {
		encoded = append(encoded, map[string]any{"id": pagerules.PageRuleActionsIDPolish, "value": m.Polish.ValueString()})
	}
	if !m.ResolveOverride.IsNull() {
		encoded = append(encoded, map[string]any{"id": pagerules.PageRuleActionsIDResolveOverride, "value": m.ResolveOverride.ValueString()})
	}
	if !m.RespectStrongEtag.IsNull() {
		encoded = append(encoded, map[string]any{"id": pagerules.PageRuleActionsIDRespectStrongEtag, "value": m.RespectStrongEtag.ValueString()})
	}
	if !m.ResponseBuffering.IsNull() {
		encoded = append(encoded, map[string]any{"id": pagerules.PageRuleActionsIDResponseBuffering, "value": m.ResponseBuffering.ValueString()})
	}
	if !m.RocketLoader.IsNull() {
		encoded = append(encoded, map[string]any{"id": pagerules.PageRuleActionsIDRocketLoader, "value": m.RocketLoader.ValueString()})
	}
	if !m.SecurityLevel.IsNull() {
		encoded = append(encoded, map[string]any{"id": pagerules.PageRuleActionsIDSecurityLevel, "value": m.SecurityLevel.ValueString()})
	}
	if !m.SortQueryStringForCache.IsNull() {
		encoded = append(encoded, map[string]any{"id": pagerules.PageRuleActionsIDSortQueryStringForCache, "value": m.SortQueryStringForCache.ValueString()})
	}
	if !m.SSL.IsNull() {
		encoded = append(encoded, map[string]any{"id": pagerules.PageRuleActionsIDSSL, "value": m.SSL.ValueString()})
	}
	if !m.TrueClientIPHeader.IsNull() {
		encoded = append(encoded, map[string]any{"id": pagerules.PageRuleActionsIDTrueClientIPHeader, "value": m.TrueClientIPHeader.ValueString()})
	}
	if !m.WAF.IsNull() {
		encoded = append(encoded, map[string]any{"id": pagerules.PageRuleActionsIDWAF, "value": m.WAF.ValueString()})
	}

	return
}

package page_rule

import (
	"context"
	"encoding/json"

	"github.com/cloudflare/cloudflare-go/v3/page_rules"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
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

type PageRuleActionsCacheKeyFieldsQueryStringModel struct {
	Include []types.String `tfsdk:"include" json:"include,optional,omitempty"`
	Exclude []types.String `tfsdk:"exclude" json:"exclude,optional,omitempty"`
}

type PageRuleActionsCacheKeyFieldsHeaderModel struct {
	CheckPresence []types.String `tfsdk:"check_presence" json:"check_presence,optional,omitempty"`
	Include       []types.String `tfsdk:"include" json:"include,optional,omitempty"`
	Exclude       []types.String `tfsdk:"exclude" json:"exclude,optional,omitempty"`
}

type PageRuleActionsCacheKeyFieldsHostModel struct {
	Resolved types.Bool `tfsdk:"resolved" json:"resolved,optional"`
}

type PageRuleActionsCacheKeyFieldsCookieModel struct {
	Include       []types.String `tfsdk:"include" json:"include,optional,omitempty"`
	CheckPresence []types.String `tfsdk:"check_presence" json:"check_presence,optional,omitempty"`
}

type PageRuleActionsCacheKeyFieldsUserModel struct {
	DeviceType types.Bool `tfsdk:"device_type" json:"device_type,optional"`
	Geo        types.Bool `tfsdk:"geo" json:"geo,optional"`
	Lang       types.Bool `tfsdk:"lang" json:"lang,optional"`
}

type PageRuleActionsCacheKeyFieldsModel struct {
	QueryString customfield.NestedObject[PageRuleActionsCacheKeyFieldsQueryStringModel] `tfsdk:"query_string" json:"query_string,optional"`
	Header      customfield.NestedObject[PageRuleActionsCacheKeyFieldsHeaderModel]      `tfsdk:"header" json:"header,optional"`
	Host        customfield.NestedObject[PageRuleActionsCacheKeyFieldsHostModel]        `tfsdk:"host" json:"host,optional"`
	Cookie      customfield.NestedObject[PageRuleActionsCacheKeyFieldsCookieModel]      `tfsdk:"cookie" json:"cookie,optional"`
	User        customfield.NestedObject[PageRuleActionsCacheKeyFieldsUserModel]        `tfsdk:"user" json:"user,optional"`
}

type PageRuleActionsForwardingURLModel struct {
	URL        types.String `tfsdk:"url" json:"url,required"`
	StatusCode types.Int64  `tfsdk:"status_code" json:"status_code,required"`
}

type PageRuleActionsModel struct {
	AlwaysUseHTTPS          types.Bool                                                   `tfsdk:"always_use_https" json:"always_use_https,optional"`
	AutomaticHTTPSRewrites  types.String                                                 `tfsdk:"automatic_https_rewrites" json:"automatic_https_rewrites,optional"`
	BrowserCacheTTL         types.Int64                                                  `tfsdk:"browser_cache_ttl" json:"browser_cache_ttl,optional"`
	BrowserCheck            types.String                                                 `tfsdk:"browser_check" json:"browser_check,optional"`
	BypassCacheOnCookie     types.String                                                 `tfsdk:"bypass_cache_on_cookie" json:"bypass_cache_on_cookie,optional"`
	CacheByDeviceType       types.String                                                 `tfsdk:"cache_by_device_type" json:"cache_by_device_type,optional"`
	CacheDeceptionArmor     types.String                                                 `tfsdk:"cache_deception_armor" json:"cache_deception_armor,optional"`
	CacheLevel              types.String                                                 `tfsdk:"cache_level" json:"cache_level,optional"`
	CacheOnCookie           types.String                                                 `tfsdk:"cache_on_cookie" json:"cache_on_cookie,optional"`
	CacheKeyFields          customfield.NestedObject[PageRuleActionsCacheKeyFieldsModel] `tfsdk:"cache_key_fields" json:"cache_key_fields,optional"`
	DisableApps             types.Bool                                                   `tfsdk:"disable_apps" json:"disable_apps,optional"`
	DisablePerformance      types.Bool                                                   `tfsdk:"disable_performance" json:"disable_performance,optional"`
	DisableSecurity         types.Bool                                                   `tfsdk:"disable_security" json:"disable_security,optional"`
	DisableZaraz            types.Bool                                                   `tfsdk:"disable_zaraz" json:"disable_zaraz,optional"`
	EdgeCacheTTL            types.Int64                                                  `tfsdk:"edge_cache_ttl" json:"edge_cache_ttl,optional"`
	EmailObfuscation        types.String                                                 `tfsdk:"email_obfuscation" json:"email_obfuscation,optional"`
	ExplicitCacheControl    types.String                                                 `tfsdk:"explicit_cache_control" json:"explicit_cache_control,optional"`
	ForwardingURL           customfield.NestedObject[PageRuleActionsForwardingURLModel]  `tfsdk:"forwarding_url" json:"forwarding_url,optional"`
	HostHeaderOverride      types.String                                                 `tfsdk:"host_header_override" json:"host_header_override,optional"`
	IPGeolocation           types.String                                                 `tfsdk:"ip_geolocation" json:"ip_geolocation,optional"`
	Mirage                  types.String                                                 `tfsdk:"mirage" json:"mirage,optional"`
	OpportunisticEncryption types.String                                                 `tfsdk:"opportunistic_encryption" json:"opportunistic_encryption,optional"`
	OriginErrorPagePassThru types.String                                                 `tfsdk:"origin_error_page_pass_thru" json:"origin_error_page_pass_thru,optional"`
	Polish                  types.String                                                 `tfsdk:"polish" json:"polish,optional"`
	ResolveOverride         types.String                                                 `tfsdk:"resolve_override" json:"resolve_override,optional"`
	RespectStrongEtag       types.String                                                 `tfsdk:"respect_strong_etag" json:"respect_strong_etag,optional"`
	ResponseBuffering       types.String                                                 `tfsdk:"response_buffering" json:"response_buffering,optional"`
	RocketLoader            types.String                                                 `tfsdk:"rocket_loader" json:"rocket_loader,optional"`
	SSL                     types.String                                                 `tfsdk:"ssl" json:"ssl,optional"`
	SecurityLevel           types.String                                                 `tfsdk:"security_level" json:"security_level,optional"`
	SortQueryStringForCache types.String                                                 `tfsdk:"sort_query_string_for_cache" json:"sort_query_string_for_cache,optional"`
	TrueClientIPHeader      types.String                                                 `tfsdk:"true_client_ip_header" json:"true_client_ip_header,optional"`
	WAF                     types.String                                                 `tfsdk:"waf" json:"waf,optional"`
}

func (m *PageRuleActionsModel) Encode() (encoded []map[string]any, err error) {
	encoded = []map[string]any{}
	if m.AlwaysUseHTTPS.ValueBool() {
		encoded = append(encoded, map[string]any{"id": page_rules.PageRuleActionsIDAlwaysUseHTTPS})
	}
	if !m.AutomaticHTTPSRewrites.IsNull() {
		encoded = append(encoded, map[string]any{"id": page_rules.PageRuleActionsIDAutomaticHTTPSRewrites, "value": m.AutomaticHTTPSRewrites.String()})
	}
	if !m.BrowserCacheTTL.IsNull() {
		encoded = append(encoded, map[string]any{"id": page_rules.PageRuleActionsIDBrowserCacheTTL, "value": m.BrowserCacheTTL.ValueInt64()})
	}
	if !m.BrowserCheck.IsNull() {
		encoded = append(encoded, map[string]any{"id": page_rules.PageRuleActionsIDBrowserCheck, "value": m.BrowserCheck.ValueString()})
	}
	if !m.BypassCacheOnCookie.IsNull() {
		encoded = append(encoded, map[string]any{"id": page_rules.PageRuleActionsIDBypassCacheOnCookie, "value": m.BypassCacheOnCookie.ValueString()})
	}
	if !m.CacheByDeviceType.IsNull() {
		encoded = append(encoded, map[string]any{"id": page_rules.PageRuleActionsIDCacheByDeviceType, "value": m.CacheByDeviceType.ValueString()})
	}
	if !m.CacheDeceptionArmor.IsNull() {
		encoded = append(encoded, map[string]any{"id": page_rules.PageRuleActionsIDCacheDeceptionArmor, "value": m.CacheDeceptionArmor.ValueString()})
	}
	if !m.CacheLevel.IsNull() {
		encoded = append(encoded, map[string]any{"id": page_rules.PageRuleActionsIDCacheLevel, "value": m.CacheLevel.ValueString()})
	}
	if !m.CacheOnCookie.IsNull() {
		encoded = append(encoded, map[string]any{"id": page_rules.PageRuleActionsIDCacheOnCookie, "value": m.CacheOnCookie.ValueString()})
	}
	if !m.CacheKeyFields.IsNull() {
		var ckf PageRuleActionsCacheKeyFieldsModel
		m.CacheKeyFields.As(context.TODO(), &ckf, basetypes.ObjectAsOptions{})

		var host PageRuleActionsCacheKeyFieldsHostModel
		ckf.Host.As(context.TODO(), &host, basetypes.ObjectAsOptions{})

		var user PageRuleActionsCacheKeyFieldsUserModel
		ckf.User.As(context.TODO(), &user, basetypes.ObjectAsOptions{})

		var qs PageRuleActionsCacheKeyFieldsQueryStringModel
		ckf.QueryString.As(context.TODO(), &qs, basetypes.ObjectAsOptions{})

		var header PageRuleActionsCacheKeyFieldsHeaderModel
		ckf.Header.As(context.TODO(), &header, basetypes.ObjectAsOptions{})

		var cookie PageRuleActionsCacheKeyFieldsCookieModel
		ckf.Cookie.As(context.TODO(), &cookie, basetypes.ObjectAsOptions{})

		// This page rule is also known as Cache Key in the schema documentation.
		// However, the API expects the "id" to be "cache_key_fields". So we are
		// hard coding it.
		encoded = append(encoded, map[string]any{
			"id": "cache_key_fields",
			"value": map[string]any{
				"cookie": map[string][]string{
					"include":        convertToStringSlice(cookie.Include),
					"check_presence": convertToStringSlice(cookie.CheckPresence),
				},
				"header": map[string][]string{
					"include":        convertToStringSlice(header.Include),
					"exclude":        convertToStringSlice(header.Exclude),
					"check_presence": convertToStringSlice(header.CheckPresence),
				},
				"host": map[string]bool{
					"resolved": host.Resolved.ValueBool(),
				},
				"query_string": map[string][]string{
					"include": convertToStringSlice(qs.Include),
					"exclude": convertToStringSlice(qs.Exclude),
				},
				"user": map[string]bool{
					"geo":         user.Geo.ValueBool(),
					"device_type": user.DeviceType.ValueBool(),
					"lang":        user.Lang.ValueBool(),
				},
			},
		})
	}
	if m.DisableApps.ValueBool() {
		encoded = append(encoded, map[string]any{"id": page_rules.PageRuleActionsIDDisableApps, "value": m.DisableApps.ValueBool()})
	}
	if m.DisablePerformance.ValueBool() {
		encoded = append(encoded, map[string]any{"id": page_rules.PageRuleActionsIDDisablePerformance})
	}
	if m.DisableSecurity.ValueBool() {
		encoded = append(encoded, map[string]any{"id": page_rules.PageRuleActionsIDDisableSecurity})
	}
	if m.DisableZaraz.ValueBool() {
		encoded = append(encoded, map[string]any{"id": page_rules.PageRuleActionsIDDisableZaraz})
	}
	if !m.EdgeCacheTTL.IsNull() {
		encoded = append(encoded, map[string]any{"id": page_rules.PageRuleActionsIDEdgeCacheTTL, "value": m.EdgeCacheTTL.ValueInt64()})
	}
	if !m.EmailObfuscation.IsNull() {
		encoded = append(encoded, map[string]any{"id": page_rules.PageRuleActionsIDEmailObfuscation, "value": m.EmailObfuscation.ValueString()})
	}
	if !m.ExplicitCacheControl.IsNull() {
		encoded = append(encoded, map[string]any{"id": page_rules.PageRuleActionsIDExplicitCacheControl, "value": m.ExplicitCacheControl.ValueString()})
	}
	if !m.ForwardingURL.IsNull() {
		var fw PageRuleActionsForwardingURLModel
		m.ForwardingURL.As(context.TODO(), &fw, basetypes.ObjectAsOptions{})
		encoded = append(encoded, map[string]any{
			"id": page_rules.PageRuleActionsIDForwardingURL,
			"value": map[string]any{
				"url":         fw.URL.ValueString(),
				"status_code": fw.StatusCode.ValueInt64(),
			},
		})
	}
	if !m.HostHeaderOverride.IsNull() {
		encoded = append(encoded, map[string]any{"id": page_rules.PageRuleActionsIDHostHeaderOverride, "value": m.HostHeaderOverride.ValueString()})
	}
	if !m.IPGeolocation.IsNull() {
		encoded = append(encoded, map[string]any{"id": page_rules.PageRuleActionsIDIPGeolocation, "value": m.IPGeolocation.ValueString()})
	}
	if !m.Mirage.IsNull() {
		encoded = append(encoded, map[string]any{"id": page_rules.PageRuleActionsIDMirage, "value": m.Mirage.ValueString()})
	}
	if !m.OpportunisticEncryption.IsNull() {
		encoded = append(encoded, map[string]any{"id": page_rules.PageRuleActionsIDOpportunisticEncryption, "value": m.OpportunisticEncryption.ValueString()})
	}
	if !m.OriginErrorPagePassThru.IsNull() {
		encoded = append(encoded, map[string]any{"id": page_rules.PageRuleActionsIDOriginErrorPagePassThru, "value": m.OriginErrorPagePassThru.ValueString()})
	}
	if !m.Polish.IsNull() {
		encoded = append(encoded, map[string]any{"id": page_rules.PageRuleActionsIDPolish, "value": m.Polish.ValueString()})
	}
	if !m.ResolveOverride.IsNull() {
		encoded = append(encoded, map[string]any{"id": page_rules.PageRuleActionsIDResolveOverride, "value": m.ResolveOverride.ValueString()})
	}
	if !m.RespectStrongEtag.IsNull() {
		encoded = append(encoded, map[string]any{"id": page_rules.PageRuleActionsIDRespectStrongEtag, "value": m.RespectStrongEtag.ValueString()})
	}
	if !m.ResponseBuffering.IsNull() {
		encoded = append(encoded, map[string]any{"id": page_rules.PageRuleActionsIDResponseBuffering, "value": m.ResponseBuffering.ValueString()})
	}
	if !m.RocketLoader.IsNull() {
		encoded = append(encoded, map[string]any{"id": page_rules.PageRuleActionsIDRocketLoader, "value": m.RocketLoader.ValueString()})
	}
	if !m.SecurityLevel.IsNull() {
		encoded = append(encoded, map[string]any{"id": page_rules.PageRuleActionsIDSecurityLevel, "value": m.SecurityLevel.ValueString()})
	}
	if !m.SortQueryStringForCache.IsNull() {
		encoded = append(encoded, map[string]any{"id": page_rules.PageRuleActionsIDSortQueryStringForCache, "value": m.SortQueryStringForCache.ValueString()})
	}
	if !m.SSL.IsNull() {
		encoded = append(encoded, map[string]any{"id": page_rules.PageRuleActionsIDSSL, "value": m.SSL.ValueString()})
	}
	if !m.TrueClientIPHeader.IsNull() {
		encoded = append(encoded, map[string]any{"id": page_rules.PageRuleActionsIDTrueClientIPHeader, "value": m.TrueClientIPHeader.ValueString()})
	}
	if !m.WAF.IsNull() {
		encoded = append(encoded, map[string]any{"id": page_rules.PageRuleActionsIDWAF, "value": m.WAF.ValueString()})
	}

	return
}

func convertToStringSlice(b []basetypes.StringValue) []string {
	ss := []string{}
	for _, v := range b {
		ss = append(ss, v.ValueString())
	}
	return ss
}

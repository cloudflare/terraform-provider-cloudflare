package page_rule

import (
	"context"
	"encoding/json"
	"strings"
	"time"

	"github.com/cloudflare/cloudflare-go/v4/page_rules"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/attr"
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
	Include []types.String `tfsdk:"include" json:"include,computed_optional,omitempty"`
	Exclude []types.String `tfsdk:"exclude" json:"exclude,computed_optional,omitempty"`
}

type PageRuleActionsCacheKeyFieldsHeaderModel struct {
	CheckPresence []types.String `tfsdk:"check_presence" json:"check_presence,computed_optional,omitempty"`
	Include       []types.String `tfsdk:"include" json:"include,computed_optional,omitempty"`
	Exclude       []types.String `tfsdk:"exclude" json:"exclude,computed_optional,omitempty"`
}

type PageRuleActionsCacheKeyFieldsHostModel struct {
	Resolved types.Bool `tfsdk:"resolved" json:"resolved,computed_optional"`
}

type PageRuleActionsCacheKeyFieldsCookieModel struct {
	Include       []types.String `tfsdk:"include" json:"include,computed_optional,omitempty"`
	CheckPresence []types.String `tfsdk:"check_presence" json:"check_presence,computed_optional,omitempty"`
}

type PageRuleActionsCacheKeyFieldsUserModel struct {
	DeviceType types.Bool `tfsdk:"device_type" json:"device_type,computed_optional"`
	Geo        types.Bool `tfsdk:"geo" json:"geo,computed_optional"`
	Lang       types.Bool `tfsdk:"lang" json:"lang,computed_optional"`
}

type PageRuleActionsCacheKeyFieldsModel struct {
	QueryString customfield.NestedObject[PageRuleActionsCacheKeyFieldsQueryStringModel] `tfsdk:"query_string" json:"query_string,optional"`
	Header      customfield.NestedObject[PageRuleActionsCacheKeyFieldsHeaderModel]      `tfsdk:"header" json:"header,optional"`
	Host        customfield.NestedObject[PageRuleActionsCacheKeyFieldsHostModel]        `tfsdk:"host" json:"host,computed_optional"`
	Cookie      customfield.NestedObject[PageRuleActionsCacheKeyFieldsCookieModel]      `tfsdk:"cookie" json:"cookie,optional"`
	User        customfield.NestedObject[PageRuleActionsCacheKeyFieldsUserModel]        `tfsdk:"user" json:"user,computed_optional"`
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
	CacheTTLByStatus        types.Dynamic                                                `tfsdk:"cache_ttl_by_status" json:"cache_ttl_by_status,optional"`
	DisableApps             types.Bool                                                   `tfsdk:"disable_apps" json:"disable_apps,optional"`
	DisablePerformance      types.Bool                                                   `tfsdk:"disable_performance" json:"disable_performance,optional"`
	DisableSecurity         types.Bool                                                   `tfsdk:"disable_security" json:"disable_security,optional"`
	DisableZaraz            types.Bool                                                   `tfsdk:"disable_zaraz" json:"disable_zaraz,optional"`
	EdgeCacheTTL            types.Int64                                                  `tfsdk:"edge_cache_ttl" json:"edge_cache_ttl,optional"`
	EmailObfuscation        types.String                                                 `tfsdk:"email_obfuscation" json:"email_obfuscation,optional"`
	ExplicitCacheControl    types.String                                                 `tfsdk:"explicit_cache_control" json:"explicit_cache_control,optional"`
	ForwardingURL           *PageRuleActionsForwardingURLModel                           `tfsdk:"forwarding_url" json:"forwarding_url,optional"`
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

		encoded = append(encoded, map[string]any{
			"id": page_rules.PageRuleActionsIDCacheKeyFields,
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
	if !m.CacheTTLByStatus.IsNull() {
		stringVal := m.CacheTTLByStatus.String()
		ttl := map[string]interface{}{}

		err = json.Unmarshal([]byte(stringVal), &ttl)
		if err != nil {
			return
		}
		value := map[string]any{}
		for k, v := range ttl {
			value[k] = v
		}
		encoded = append(encoded, map[string]any{"id": page_rules.PageRuleActionsIDCacheTTLByStatus, "value": value})
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
	if m.ForwardingURL != nil {
		encoded = append(encoded, map[string]any{
			"id": page_rules.PageRuleActionsIDForwardingURL,
			"value": map[string]any{
				"url":         m.ForwardingURL.URL.ValueString(),
				"status_code": m.ForwardingURL.StatusCode.ValueInt64(),
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

func UnmarshalPageRuleModel(b []byte) (*PageRuleModel, error) {
	var resp struct {
		Result struct {
			ID         string `json:"id"`
			ZoneID     string `json:"zone_id"`
			Priority   int64  `json:"priority"`
			Status     string `json:"status"`
			CreatedOn  string `json:"created_on"`
			ModifiedOn string `json:"modified_on"`
			Targets    []struct {
				Constraint struct {
					Operator string `json:"operator"`
					Value    string `json:"value"`
				} `json:"constraint"`
				Target string `json:"target"`
			} `json:"targets"`
			Actions []struct {
				ID    string          `json:"id"`
				Value json.RawMessage `json:"value"`
			} `json:"actions"`
		} `json:"result"`
	}

	if err := json.Unmarshal(b, &resp); err != nil {
		return nil, err
	}

	m := &PageRuleModel{
		ID:         types.StringValue(resp.Result.ID),
		Priority:   types.Int64Value(resp.Result.Priority),
		Status:     types.StringValue(resp.Result.Status),
		CreatedOn:  timetypes.NewRFC3339TimeValue(timeFromString(resp.Result.CreatedOn)),
		ModifiedOn: timetypes.NewRFC3339TimeValue(timeFromString(resp.Result.ModifiedOn)),
	}

	if len(resp.Result.Targets) > 0 {
		m.Target = types.StringValue(strings.TrimRight(resp.Result.Targets[0].Constraint.Value, "/"))
	}

	m.Actions = &PageRuleActionsModel{}
	for _, action := range resp.Result.Actions {
		switch action.ID {
		case "always_use_https":
			m.Actions.AlwaysUseHTTPS = types.BoolValue(true)
		case "automatic_https_rewrites":
			var val string
			_ = json.Unmarshal(action.Value, &val)
			m.Actions.AutomaticHTTPSRewrites = types.StringValue(val)
		case "browser_cache_ttl":
			var val int64
			_ = json.Unmarshal(action.Value, &val)
			m.Actions.BrowserCacheTTL = types.Int64Value(val)
		case "browser_check":
			var val string
			_ = json.Unmarshal(action.Value, &val)
			m.Actions.BrowserCheck = types.StringValue(val)
		case "bypass_cache_on_cookie":
			var val string
			_ = json.Unmarshal(action.Value, &val)
			m.Actions.BypassCacheOnCookie = types.StringValue(val)
		case "cache_by_device_type":
			var val string
			_ = json.Unmarshal(action.Value, &val)
			m.Actions.CacheByDeviceType = types.StringValue(val)
		case "cache_deception_armor":
			var val string
			_ = json.Unmarshal(action.Value, &val)
			m.Actions.CacheDeceptionArmor = types.StringValue(val)
		case "cache_level":
			var val string
			_ = json.Unmarshal(action.Value, &val)
			m.Actions.CacheLevel = types.StringValue(val)
		case "cache_on_cookie":
			var val string
			_ = json.Unmarshal(action.Value, &val)
			m.Actions.CacheOnCookie = types.StringValue(val)
		case "cache_key_fields":
			var ckf PageRuleActionsCacheKeyFieldsModel
			var val map[string]any
			_ = json.Unmarshal(action.Value, &val)
			for key, field := range val {
				switch key {
				case "cookie":
					var cookie PageRuleActionsCacheKeyFieldsCookieModel
					cookiesMap := field.(map[string]interface{})
					if inc, ok := cookiesMap["include"]; ok {
						includesSlice, ok := inc.([]interface{})
						if ok {
							for _, i := range includesSlice {
								cookie.Include = append(cookie.Include, types.StringValue(i.(string)))
							}
						}
					}
					if cp, ok := cookiesMap["check_presence"]; ok {
						cpSlice, ok := cp.([]interface{})
						if ok {
							for _, c := range cpSlice {
								cookie.CheckPresence = append(cookie.CheckPresence, types.StringValue(c.(string)))
							}
						}
					}
					if len(cookie.CheckPresence) != 0 || len(cookie.Include) != 0 {
						ckf.Cookie, _ = customfield.NewObject[PageRuleActionsCacheKeyFieldsCookieModel](context.Background(), &cookie)
					}
				case "header":
					var header PageRuleActionsCacheKeyFieldsHeaderModel
					headersMap := field.(map[string]interface{})
					if inc, ok := headersMap["include"]; ok {
						includesSlice, ok := inc.([]interface{})
						if ok {
							for _, i := range includesSlice {
								header.Include = append(header.Include, types.StringValue(i.(string)))
							}
						}
					}
					if exc, ok := headersMap["exclude"]; ok {
						excludesSlice, ok := exc.([]interface{})
						if ok {
							for _, e := range excludesSlice {
								header.Exclude = append(header.Exclude, types.StringValue(e.(string)))
							}
						}
					}
					if cp, ok := headersMap["check_presence"]; ok {
						cpSlice, ok := cp.([]interface{})
						if ok {
							for _, c := range cpSlice {
								header.CheckPresence = append(header.CheckPresence, types.StringValue(c.(string)))
							}
						}
					}
					if len(header.CheckPresence) != 0 || len(header.Include) != 0 || len(header.Exclude) != 0 {
						ckf.Header, _ = customfield.NewObject[PageRuleActionsCacheKeyFieldsHeaderModel](context.Background(), &header)
					}
				case "host":
					var host PageRuleActionsCacheKeyFieldsHostModel
					if resolved, ok := (field.(map[string]interface{}))["resolved"]; ok {
						isResolved, ok := resolved.(bool)
						if ok {
							host.Resolved = types.BoolValue(isResolved)
						}
					}
					ckf.Host, _ = customfield.NewObject[PageRuleActionsCacheKeyFieldsHostModel](context.Background(), &host)
				case "query_string":
					var queryString PageRuleActionsCacheKeyFieldsQueryStringModel
					qsMap := field.(map[string]interface{})
					if inc, ok := qsMap["include"]; ok {
						includesSlice, ok := inc.([]interface{})
						if ok {
							for _, i := range includesSlice {
								queryString.Include = append(queryString.Include, types.StringValue(i.(string)))
							}
						}
					}
					if exc, ok := qsMap["exclude"]; ok {
						excludesSlice, ok := exc.([]interface{})
						if ok {
							for _, e := range excludesSlice {
								queryString.Exclude = append(queryString.Exclude, types.StringValue(e.(string)))
							}
						}
					}
					ckf.QueryString, _ = customfield.NewObject[PageRuleActionsCacheKeyFieldsQueryStringModel](context.Background(), &queryString)
				case "user":
					var user PageRuleActionsCacheKeyFieldsUserModel
					userMap := field.(map[string]interface{})
					if geo, ok := userMap["geo"]; ok {
						isGeo, ok := geo.(bool)
						if ok {
							user.Geo = types.BoolValue(isGeo)
						}
					}
					if deviceType, ok := userMap["device_type"]; ok {
						isDeviceType, ok := deviceType.(bool)
						if ok {
							user.DeviceType = types.BoolValue(isDeviceType)
						}
					}
					if lang, ok := userMap["lang"]; ok {
						isLang, ok := lang.(bool)
						if ok {
							user.Lang = types.BoolValue(isLang)
						}
					}
					ckf.User, _ = customfield.NewObject[PageRuleActionsCacheKeyFieldsUserModel](context.Background(), &user)
				}
			}
			m.Actions.CacheKeyFields, _ = customfield.NewObject[PageRuleActionsCacheKeyFieldsModel](context.Background(), &ckf)
		case "cache_ttl_by_status":
			var val map[string]interface{}
			_ = json.Unmarshal(action.Value, &val)
			elements := make(map[string]attr.Value)
			aMap := make(map[string]attr.Type)
			for k, v := range val {
				aMap[k] = types.DynamicType
				switch v.(type) {
				case string:
					elements[k] = types.DynamicValue(basetypes.NewDynamicValue(types.StringValue(v.(string))))
				case float64:
					elements[k] = types.DynamicValue(basetypes.NewDynamicValue(types.Float64Value(v.(float64))))
				}

			}
			mapValue, _ := types.ObjectValue(aMap, elements)
			m.Actions.CacheTTLByStatus = types.DynamicValue(mapValue)
		case "disable_apps":
			m.Actions.DisableApps = types.BoolValue(true)
		case "disable_performance":
			m.Actions.DisablePerformance = types.BoolValue(true)
		case "disable_security":
			m.Actions.DisableSecurity = types.BoolValue(true)
		case "disable_zaraz":
			m.Actions.DisableZaraz = types.BoolValue(true)
		case "edge_cache_ttl":
			var val int64
			_ = json.Unmarshal(action.Value, &val)
			m.Actions.EdgeCacheTTL = types.Int64Value(val)
		case "email_obfuscation":
			var val string
			_ = json.Unmarshal(action.Value, &val)
			m.Actions.EmailObfuscation = types.StringValue(val)
		case "explicit_cache_control":
			var val string
			_ = json.Unmarshal(action.Value, &val)
			m.Actions.ExplicitCacheControl = types.StringValue(val)
		case "forwarding_url":
			var val struct {
				URL        string `json:"url"`
				StatusCode int64  `json:"status_code"`
			}
			_ = json.Unmarshal(action.Value, &val)
			m.Actions.ForwardingURL = &PageRuleActionsForwardingURLModel{
				URL:        types.StringValue(val.URL),
				StatusCode: types.Int64Value(val.StatusCode),
			}
		case "host_header_override":
			var val string
			_ = json.Unmarshal(action.Value, &val)
			m.Actions.HostHeaderOverride = types.StringValue(val)
		case "ip_geolocation":
			var val string
			_ = json.Unmarshal(action.Value, &val)
			m.Actions.IPGeolocation = types.StringValue(val)
		case "mirage":
			var val string
			_ = json.Unmarshal(action.Value, &val)
			m.Actions.Mirage = types.StringValue(val)
		case "opportunistic_encryption":
			var val string
			_ = json.Unmarshal(action.Value, &val)
			m.Actions.OpportunisticEncryption = types.StringValue(val)
		case "origin_error_page_pass_thru":
			var val string
			_ = json.Unmarshal(action.Value, &val)
			m.Actions.OriginErrorPagePassThru = types.StringValue(val)
		case "polish":
			var val string
			_ = json.Unmarshal(action.Value, &val)
			m.Actions.Polish = types.StringValue(val)
		case "resolve_override":
			var val string
			_ = json.Unmarshal(action.Value, &val)
			m.Actions.ResolveOverride = types.StringValue(val)
		case "respect_strong_etag":
			var val string
			_ = json.Unmarshal(action.Value, &val)
			m.Actions.RespectStrongEtag = types.StringValue(val)
		case "response_buffering":
			var val string
			_ = json.Unmarshal(action.Value, &val)
			m.Actions.ResponseBuffering = types.StringValue(val)
		case "rocket_loader":
			var val string
			_ = json.Unmarshal(action.Value, &val)
			m.Actions.RocketLoader = types.StringValue(val)
		case "ssl":
			var val string
			_ = json.Unmarshal(action.Value, &val)
			m.Actions.SSL = types.StringValue(val)
		case "security_level":
			var val string
			_ = json.Unmarshal(action.Value, &val)
			m.Actions.SecurityLevel = types.StringValue(val)
		case "sort_query_string_for_cache":
			var val string
			_ = json.Unmarshal(action.Value, &val)
			m.Actions.SortQueryStringForCache = types.StringValue(val)
		case "true_client_ip_header":
			var val string
			_ = json.Unmarshal(action.Value, &val)
			m.Actions.TrueClientIPHeader = types.StringValue(val)
		case "waf":
			var val string
			_ = json.Unmarshal(action.Value, &val)
			m.Actions.WAF = types.StringValue(val)
		}

	}
	return m, nil
}

func timeFromString(s string) time.Time {
	t, _ := time.Parse(time.RFC3339Nano, s)
	return t
}

func convertToStringSlice(b []basetypes.StringValue) []string {
	ss := []string{}
	for _, v := range b {
		ss = append(ss, v.ValueString())
	}
	return ss
}

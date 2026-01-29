package v500

import (
	"context"
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// parseJSONToSourceModel parses raw JSON into SourceCloudflarePageRuleModel.
// This is needed because json.Unmarshal doesn't work with framework types.
func parseJSONToSourceModel(ctx context.Context, rawJSON []byte) (SourceCloudflarePageRuleModel, diag.Diagnostics) {
	var diags diag.Diagnostics
	var model SourceCloudflarePageRuleModel

	// Parse JSON into generic map first
	var jsonMap map[string]interface{}
	if err := json.Unmarshal(rawJSON, &jsonMap); err != nil {
		diags.AddError(
			"Failed to unmarshal JSON",
			err.Error(),
		)
		return model, diags
	}

	// Manually map fields
	if id, ok := jsonMap["id"].(string); ok {
		model.ID = types.StringValue(id)
	} else {
		model.ID = types.StringNull()
	}

	if zoneID, ok := jsonMap["zone_id"].(string); ok {
		model.ZoneID = types.StringValue(zoneID)
	}

	if target, ok := jsonMap["target"].(string); ok {
		model.Target = types.StringValue(target)
	}

	if priority, ok := jsonMap["priority"].(float64); ok {
		model.Priority = types.Int64Value(int64(priority))
	} else {
		model.Priority = types.Int64Null()
	}

	if status, ok := jsonMap["status"].(string); ok {
		model.Status = types.StringValue(status)
	} else {
		model.Status = types.StringNull()
	}

	// Parse actions array (v4 format)
	if actionsRaw, ok := jsonMap["actions"].([]interface{}); ok && len(actionsRaw) > 0 {
		if actionsMap, ok := actionsRaw[0].(map[string]interface{}); ok {
			actionsModel := parseActionsFromJSON(actionsMap)
			model.Actions = []SourceActionsModel{actionsModel}
		}
	}

	return model, diags
}

// parseActionsFromJSON parses actions map into SourceActionsModel.
func parseActionsFromJSON(jsonMap map[string]interface{}) SourceActionsModel {
	var model SourceActionsModel

	// Helper to convert interface{} to types.StringValue
	strValue := func(key string) types.String {
		if val, ok := jsonMap[key].(string); ok {
			return types.StringValue(val)
		}
		return types.StringNull()
	}

	boolValue := func(key string) types.Bool {
		if val, ok := jsonMap[key].(bool); ok {
			return types.BoolValue(val)
		}
		return types.BoolNull()
	}

	int64Value := func(key string) types.Int64 {
		if val, ok := jsonMap[key].(float64); ok {
			return types.Int64Value(int64(val))
		}
		return types.Int64Null()
	}

	// Map all simple fields
	model.AlwaysUseHTTPS = boolValue("always_use_https")
	model.AutomaticHTTPSRewrites = strValue("automatic_https_rewrites")
	model.BrowserCacheTTL = strValue("browser_cache_ttl")
	model.BrowserCheck = strValue("browser_check")
	model.CacheByDeviceType = strValue("cache_by_device_type")
	model.CacheDeceptionArmor = strValue("cache_deception_armor")
	model.CacheLevel = strValue("cache_level")
	model.CacheOnCookie = strValue("cache_on_cookie")
	model.DisableApps = boolValue("disable_apps")
	model.DisablePerformance = boolValue("disable_performance")
	model.DisableRailgun = boolValue("disable_railgun")
	model.DisableSecurity = boolValue("disable_security")
	model.DisableZaraz = boolValue("disable_zaraz")
	model.EdgeCacheTTL = int64Value("edge_cache_ttl")
	model.EmailObfuscation = strValue("email_obfuscation")
	model.ExplicitCacheControl = strValue("explicit_cache_control")
	model.HostHeaderOverride = strValue("host_header_override")
	model.IPGeolocation = strValue("ip_geolocation")
	model.Mirage = strValue("mirage")
	model.OpportunisticEncryption = strValue("opportunistic_encryption")
	model.OriginErrorPagePassThru = strValue("origin_error_page_pass_thru")
	model.Polish = strValue("polish")
	model.ResolveOverride = strValue("resolve_override")
	model.RespectStrongEtag = strValue("respect_strong_etag")
	model.ResponseBuffering = strValue("response_buffering")
	model.RocketLoader = strValue("rocket_loader")
	model.SecurityLevel = strValue("security_level")
	model.ServerSideExclude = strValue("server_side_exclude")
	model.SortQueryStringForCache = strValue("sort_query_string_for_cache")
	model.SSL = strValue("ssl")
	model.TrueClientIPHeader = strValue("true_client_ip_header")
	model.WAF = strValue("waf")

	// Parse nested structures
	if cacheKeyFields, ok := jsonMap["cache_key_fields"].([]interface{}); ok && len(cacheKeyFields) > 0 {
		if ckfMap, ok := cacheKeyFields[0].(map[string]interface{}); ok {
			model.CacheKeyFields = parseCacheKeyFields(ckfMap)
		}
	}

	if cacheTTLByStatus, ok := jsonMap["cache_ttl_by_status"].([]interface{}); ok {
		model.CacheTTLByStatus = parseCacheTTLByStatus(cacheTTLByStatus)
	}

	if forwardingURL, ok := jsonMap["forwarding_url"].([]interface{}); ok && len(forwardingURL) > 0 {
		if fwdMap, ok := forwardingURL[0].(map[string]interface{}); ok {
			model.ForwardingURL = parseForwardingURL(fwdMap)
		}
	}

	if minify, ok := jsonMap["minify"].([]interface{}); ok && len(minify) > 0 {
		if minifyMap, ok := minify[0].(map[string]interface{}); ok {
			model.Minify = parseMinify(minifyMap)
		}
	}

	return model
}

// parseCacheKeyFields parses cache_key_fields nested structure
func parseCacheKeyFields(jsonMap map[string]interface{}) []SourceCacheKeyFieldsModel {
	var model SourceCacheKeyFieldsModel

	// Parse cookie
	if cookie, ok := jsonMap["cookie"].([]interface{}); ok && len(cookie) > 0 {
		if cookieMap, ok := cookie[0].(map[string]interface{}); ok {
			var cookieModel SourceCacheKeyFieldsCookieModel
			if checkPresence, ok := cookieMap["check_presence"].([]interface{}); ok {
				cookieModel.CheckPresence = parseStringSet(checkPresence)
			}
			if include, ok := cookieMap["include"].([]interface{}); ok {
				cookieModel.Include = parseStringSet(include)
			}
			model.Cookie = []SourceCacheKeyFieldsCookieModel{cookieModel}
		}
	}

	// Parse header
	if header, ok := jsonMap["header"].([]interface{}); ok && len(header) > 0 {
		if headerMap, ok := header[0].(map[string]interface{}); ok {
			var headerModel SourceCacheKeyFieldsHeaderModel
			if checkPresence, ok := headerMap["check_presence"].([]interface{}); ok {
				headerModel.CheckPresence = parseStringSet(checkPresence)
			}
			if exclude, ok := headerMap["exclude"].([]interface{}); ok {
				headerModel.Exclude = parseStringSet(exclude)
			}
			if include, ok := headerMap["include"].([]interface{}); ok {
				headerModel.Include = parseStringSet(include)
			}
			model.Header = []SourceCacheKeyFieldsHeaderModel{headerModel}
		}
	}

	// Parse host
	if host, ok := jsonMap["host"].([]interface{}); ok && len(host) > 0 {
		if hostMap, ok := host[0].(map[string]interface{}); ok {
			var hostModel SourceCacheKeyFieldsHostModel
			if resolved, ok := hostMap["resolved"].(bool); ok {
				hostModel.Resolved = types.BoolValue(resolved)
			}
			model.Host = []SourceCacheKeyFieldsHostModel{hostModel}
		}
	}

	// Parse query_string
	if queryString, ok := jsonMap["query_string"].([]interface{}); ok && len(queryString) > 0 {
		if qsMap, ok := queryString[0].(map[string]interface{}); ok {
			var qsModel SourceCacheKeyFieldsQueryStringModel
			if exclude, ok := qsMap["exclude"].([]interface{}); ok {
				qsModel.Exclude = parseStringSet(exclude)
			}
			if ignore, ok := qsMap["ignore"].(bool); ok {
				qsModel.Ignore = types.BoolValue(ignore)
			}
			if include, ok := qsMap["include"].([]interface{}); ok {
				qsModel.Include = parseStringSet(include)
			}
			model.QueryString = []SourceCacheKeyFieldsQueryStringModel{qsModel}
		}
	}

	// Parse user
	if user, ok := jsonMap["user"].([]interface{}); ok && len(user) > 0 {
		if userMap, ok := user[0].(map[string]interface{}); ok {
			var userModel SourceCacheKeyFieldsUserModel
			if deviceType, ok := userMap["device_type"].(bool); ok {
				userModel.DeviceType = types.BoolValue(deviceType)
			}
			if geo, ok := userMap["geo"].(bool); ok {
				userModel.Geo = types.BoolValue(geo)
			}
			if lang, ok := userMap["lang"].(bool); ok {
				userModel.Lang = types.BoolValue(lang)
			}
			model.User = []SourceCacheKeyFieldsUserModel{userModel}
		}
	}

	return []SourceCacheKeyFieldsModel{model}
}

// parseCacheTTLByStatus parses cache_ttl_by_status array
func parseCacheTTLByStatus(arr []interface{}) []SourceCacheTTLByStatusModel {
	var result []SourceCacheTTLByStatusModel
	for _, item := range arr {
		if itemMap, ok := item.(map[string]interface{}); ok {
			var model SourceCacheTTLByStatusModel
			if codes, ok := itemMap["codes"].(string); ok {
				model.Codes = types.StringValue(codes)
			}
			if ttl, ok := itemMap["ttl"].(float64); ok {
				model.TTL = types.Int64Value(int64(ttl))
			}
			result = append(result, model)
		}
	}
	return result
}

// parseForwardingURL parses forwarding_url nested structure
func parseForwardingURL(jsonMap map[string]interface{}) []SourceForwardingURLModel {
	var model SourceForwardingURLModel
	if url, ok := jsonMap["url"].(string); ok {
		model.URL = types.StringValue(url)
	}
	if statusCode, ok := jsonMap["status_code"].(float64); ok {
		model.StatusCode = types.Int64Value(int64(statusCode))
	}
	return []SourceForwardingURLModel{model}
}

// parseMinify parses minify nested structure
func parseMinify(jsonMap map[string]interface{}) []SourceMinifyModel {
	var model SourceMinifyModel
	if css, ok := jsonMap["css"].(string); ok {
		model.CSS = types.StringValue(css)
	}
	if html, ok := jsonMap["html"].(string); ok {
		model.HTML = types.StringValue(html)
	}
	if js, ok := jsonMap["js"].(string); ok {
		model.JS = types.StringValue(js)
	}
	return []SourceMinifyModel{model}
}

// parseStringList converts []interface{} to []types.String
func parseStringList(arr []interface{}) []types.String {
	var result []types.String
	for _, item := range arr {
		if str, ok := item.(string); ok {
			result = append(result, types.StringValue(str))
		}
	}
	return result
}

// parseStringSet converts []interface{} to types.Set
func parseStringSet(arr []interface{}) types.Set {
	var elements []string
	for _, item := range arr {
		if str, ok := item.(string); ok {
			elements = append(elements, str)
		}
	}

	if len(elements) == 0 {
		return types.SetNull(types.StringType)
	}

	set, _ := types.SetValueFrom(context.Background(), types.StringType, elements)
	return set
}

// parseJSONToTargetModel parses raw JSON (v5 format) into TargetPageRuleModel.
// This handles early v5 states where actions is already an object.
// Since the structure matches our target model (pointers, not customfield), this is straightforward.
func parseJSONToTargetModel(ctx context.Context, rawJSON []byte) (*TargetPageRuleModel, diag.Diagnostics) {
	var diags diag.Diagnostics
	var model TargetPageRuleModel

	// Parse JSON into generic map first
	var jsonMap map[string]interface{}
	if err := json.Unmarshal(rawJSON, &jsonMap); err != nil {
		diags.AddError(
			"Failed to unmarshal JSON",
			err.Error(),
		)
		return nil, diags
	}

	// Top-level fields
	if id, ok := jsonMap["id"].(string); ok {
		model.ID = types.StringValue(id)
	} else {
		model.ID = types.StringNull()
	}

	if zoneID, ok := jsonMap["zone_id"].(string); ok {
		model.ZoneID = types.StringValue(zoneID)
	}

	if target, ok := jsonMap["target"].(string); ok {
		model.Target = types.StringValue(target)
	}

	if priority, ok := jsonMap["priority"].(float64); ok {
		model.Priority = types.Int64Value(int64(priority))
	} else {
		model.Priority = types.Int64Null()
	}

	if status, ok := jsonMap["status"].(string); ok {
		model.Status = types.StringValue(status)
	} else {
		model.Status = types.StringNull()
	}

	// Parse actions object (v5 format - already an object, not array)
	if actionsRaw, ok := jsonMap["actions"].(map[string]interface{}); ok {
		model.Actions = parseTargetActionsFromJSON(actionsRaw)
	}

	return &model, diags
}

// parseTargetActionsFromJSON parses actions object (v5 format) into TargetActionsModel.
// For early v5 states, actions is already an object with the correct structure.
func parseTargetActionsFromJSON(jsonMap map[string]interface{}) *TargetActionsModel {
	var model TargetActionsModel

	// Initialize CacheTTLByStatus with proper type (will be overwritten if present in JSON)
	model.CacheTTLByStatus = types.MapNull(types.StringType)

	// Helper functions (same as before)
	strValue := func(key string) types.String {
		if val, ok := jsonMap[key].(string); ok {
			return types.StringValue(val)
		}
		return types.StringNull()
	}

	boolValue := func(key string) types.Bool {
		if val, ok := jsonMap[key].(bool); ok {
			return types.BoolValue(val)
		}
		return types.BoolNull()
	}

	int64Value := func(key string) types.Int64 {
		if val, ok := jsonMap[key].(float64); ok {
			return types.Int64Value(int64(val))
		}
		return types.Int64Null()
	}

	// Map all fields
	model.AlwaysUseHTTPS = boolValue("always_use_https")
	model.AutomaticHTTPSRewrites = strValue("automatic_https_rewrites")
	model.BrowserCacheTTL = int64Value("browser_cache_ttl")
	model.BrowserCheck = strValue("browser_check")
	model.BypassCacheOnCookie = strValue("bypass_cache_on_cookie")
	model.CacheByDeviceType = strValue("cache_by_device_type")
	model.CacheDeceptionArmor = strValue("cache_deception_armor")
	model.CacheLevel = strValue("cache_level")
	model.CacheOnCookie = strValue("cache_on_cookie")
	model.DisableApps = boolValue("disable_apps")
	model.DisablePerformance = boolValue("disable_performance")
	model.DisableSecurity = boolValue("disable_security")
	model.DisableZaraz = boolValue("disable_zaraz")
	model.EdgeCacheTTL = int64Value("edge_cache_ttl")
	model.EmailObfuscation = strValue("email_obfuscation")
	model.ExplicitCacheControl = strValue("explicit_cache_control")
	model.HostHeaderOverride = strValue("host_header_override")
	model.IPGeolocation = strValue("ip_geolocation")
	model.Mirage = strValue("mirage")
	model.OpportunisticEncryption = strValue("opportunistic_encryption")
	model.OriginErrorPagePassThru = strValue("origin_error_page_pass_thru")
	model.Polish = strValue("polish")
	model.ResolveOverride = strValue("resolve_override")
	model.RespectStrongEtag = strValue("respect_strong_etag")
	model.ResponseBuffering = strValue("response_buffering")
	model.RocketLoader = strValue("rocket_loader")
	model.SecurityLevel = strValue("security_level")
	model.SortQueryStringForCache = strValue("sort_query_string_for_cache")
	model.SSL = strValue("ssl")
	model.TrueClientIPHeader = strValue("true_client_ip_header")
	model.WAF = strValue("waf")

	// Nested structures (v5 format - objects, not arrays)
	if forwardingURL, ok := jsonMap["forwarding_url"].(map[string]interface{}); ok {
		model.ForwardingURL = parseTargetForwardingURL(forwardingURL)
	}

	if cacheKeyFields, ok := jsonMap["cache_key_fields"].(map[string]interface{}); ok {
		model.CacheKeyFields = parseTargetCacheKeyFields(cacheKeyFields)
	}

	if cacheTTLByStatus, ok := jsonMap["cache_ttl_by_status"].(map[string]interface{}); ok {
		model.CacheTTLByStatus = parseTargetCacheTTLByStatus(cacheTTLByStatus)
	}

	return &model
}

// parseTargetForwardingURL parses forwarding_url object.
func parseTargetForwardingURL(jsonMap map[string]interface{}) *TargetForwardingURLModel {
	var model TargetForwardingURLModel
	if url, ok := jsonMap["url"].(string); ok {
		model.URL = types.StringValue(url)
	}
	if statusCode, ok := jsonMap["status_code"].(float64); ok {
		model.StatusCode = types.Int64Value(int64(statusCode))
	}
	return &model
}

// parseTargetCacheKeyFields parses cache_key_fields object (v5 format - nested objects).
func parseTargetCacheKeyFields(jsonMap map[string]interface{}) *TargetCacheKeyFieldsModel {
	var model TargetCacheKeyFieldsModel

	if cookie, ok := jsonMap["cookie"].(map[string]interface{}); ok {
		cookieModel := &TargetCacheKeyFieldsCookieModel{}
		if checkPresence, ok := cookie["check_presence"].([]interface{}); ok {
			cookieModel.CheckPresence = parseStringListToTypesStringSlice(checkPresence)
		}
		if include, ok := cookie["include"].([]interface{}); ok {
			cookieModel.Include = parseStringListToTypesStringSlice(include)
		}
		model.Cookie = cookieModel
	}

	if header, ok := jsonMap["header"].(map[string]interface{}); ok {
		headerModel := &TargetCacheKeyFieldsHeaderModel{}
		if checkPresence, ok := header["check_presence"].([]interface{}); ok {
			headerModel.CheckPresence = parseStringListToTypesStringSlice(checkPresence)
		}
		if include, ok := header["include"].([]interface{}); ok {
			headerModel.Include = parseStringListToTypesStringSlice(include)
		}
		if exclude, ok := header["exclude"].([]interface{}); ok {
			headerModel.Exclude = parseStringListToTypesStringSlice(exclude)
		}
		model.Header = headerModel
	}

	if host, ok := jsonMap["host"].(map[string]interface{}); ok {
		hostModel := &TargetCacheKeyFieldsHostModel{}
		if resolved, ok := host["resolved"].(bool); ok {
			hostModel.Resolved = types.BoolValue(resolved)
		}
		model.Host = hostModel
	}

	if queryString, ok := jsonMap["query_string"].(map[string]interface{}); ok {
		qsModel := &TargetCacheKeyFieldsQueryStringModel{}
		if include, ok := queryString["include"].([]interface{}); ok {
			qsModel.Include = parseStringListToTypesStringSlice(include)
		}
		if exclude, ok := queryString["exclude"].([]interface{}); ok {
			qsModel.Exclude = parseStringListToTypesStringSlice(exclude)
		}
		// NOTE: The 'ignore' field from v4 was removed in v5 - not parsed
		model.QueryString = qsModel
	}

	if user, ok := jsonMap["user"].(map[string]interface{}); ok {
		userModel := &TargetCacheKeyFieldsUserModel{}
		if deviceType, ok := user["device_type"].(bool); ok {
			userModel.DeviceType = types.BoolValue(deviceType)
		}
		if geo, ok := user["geo"].(bool); ok {
			userModel.Geo = types.BoolValue(geo)
		}
		if lang, ok := user["lang"].(bool); ok {
			userModel.Lang = types.BoolValue(lang)
		}
		model.User = userModel
	}

	return &model
}

// parseTargetCacheTTLByStatus parses cache_ttl_by_status map (v5 format - map[string]string).
func parseTargetCacheTTLByStatus(jsonMap map[string]interface{}) types.Map {
	if len(jsonMap) == 0 {
		return types.MapNull(types.StringType)
	}

	elements := make(map[string]attr.Value, len(jsonMap))
	for key, val := range jsonMap {
		if strVal, ok := val.(string); ok {
			elements[key] = types.StringValue(strVal)
		}
	}

	// Use MapValueMust to ensure proper type specification
	return types.MapValueMust(types.StringType, elements)
}

// parseStringListToTypesStringSlice converts []interface{} of strings to []types.String.
func parseStringListToTypesStringSlice(arr []interface{}) []types.String {
	var result []types.String
	for _, item := range arr {
		if str, ok := item.(string); ok {
			result = append(result, types.StringValue(str))
		}
	}
	return result
}

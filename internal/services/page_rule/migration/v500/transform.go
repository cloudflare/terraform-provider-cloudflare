package v500

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Transform converts source (legacy v4 SDKv2) state to target (current v5 Plugin Framework) state.
// This function handles all field transformations, type conversions, and nested structure migrations.
func Transform(ctx context.Context, source SourceCloudflarePageRuleModel) (*TargetPageRuleModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	// Step 1: Validate required fields
	if source.ZoneID.IsNull() || source.ZoneID.IsUnknown() {
		diags.AddError(
			"Missing required field",
			"zone_id is required for page_rule migration. The source state is missing this field, which indicates corrupted state.",
		)
		return nil, diags
	}
	if source.Target.IsNull() || source.Target.IsUnknown() {
		diags.AddError(
			"Missing required field",
			"target is required for page_rule migration. The source state is missing this field, which indicates corrupted state.",
		)
		return nil, diags
	}

	// Step 2: Initialize target with direct copies
	target := &TargetPageRuleModel{
		ID:       source.ID,
		ZoneID:   source.ZoneID,
		Target:   source.Target,
		Priority: source.Priority,
	}

	// Step 3: Handle status field (CRITICAL: default changed from "active" to "disabled")
	// If v4 state has no status (or empty), set to "active" to preserve v4 behavior
	if source.Status.IsNull() || source.Status.IsUnknown() || source.Status.ValueString() == "" {
		target.Status = types.StringValue("active")
	} else {
		target.Status = source.Status
	}

	// Step 4: Set computed timestamp fields to null (will refresh from API)
	target.CreatedOn = timetypes.NewRFC3339Null()
	target.ModifiedOn = timetypes.NewRFC3339Null()

	// Step 5: Transform actions (CRITICAL: extract from array[0])
	if len(source.Actions) > 0 {
		targetActions, actionDiags := transformActions(ctx, source.Actions[0])
		diags.Append(actionDiags...)
		if diags.HasError() {
			return nil, diags
		}
		target.Actions = targetActions
	} else {
		diags.AddError(
			"Missing required field",
			"actions is required for page_rule migration. The source state is missing this field.",
		)
		return nil, diags
	}

	return target, diags
}

// transformActions converts source actions (from SDKv2 TypeList MaxItems:1) to target actions (SingleNestedAttribute).
// This handles all action fields including complex nested structures.
func transformActions(ctx context.Context, source SourceActionsModel) (*TargetActionsModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	target := &TargetActionsModel{}

	// Boolean fields: false → null (v4 had default: false, v5 stores null instead)
	target.AlwaysUseHTTPS = nullifyFalseBool(source.AlwaysUseHTTPS)
	target.DisableApps = nullifyFalseBool(source.DisableApps)
	target.DisablePerformance = nullifyFalseBool(source.DisablePerformance)
	target.DisableSecurity = nullifyFalseBool(source.DisableSecurity)
	target.DisableZaraz = nullifyFalseBool(source.DisableZaraz)
	// NOTE: source.DisableRailgun is intentionally NOT copied (deprecated)
	// NOTE: source.Minify is intentionally NOT copied (deprecated)

	// String fields: nullify empty strings (v4 may have empty strings, v5 wants null)
	target.AutomaticHTTPSRewrites = nullifyEmptyString(source.AutomaticHTTPSRewrites)
	target.BrowserCheck = nullifyEmptyString(source.BrowserCheck)
	target.BypassCacheOnCookie = nullifyEmptyString(source.BypassCacheOnCookie)
	target.CacheByDeviceType = nullifyEmptyString(source.CacheByDeviceType)
	target.CacheDeceptionArmor = nullifyEmptyString(source.CacheDeceptionArmor)
	target.CacheLevel = nullifyEmptyString(source.CacheLevel)
	target.CacheOnCookie = nullifyEmptyString(source.CacheOnCookie)
	target.EmailObfuscation = nullifyEmptyString(source.EmailObfuscation)
	target.ExplicitCacheControl = nullifyEmptyString(source.ExplicitCacheControl)
	target.HostHeaderOverride = nullifyEmptyString(source.HostHeaderOverride)
	target.IPGeolocation = nullifyEmptyString(source.IPGeolocation)
	target.Mirage = nullifyEmptyString(source.Mirage)
	target.OpportunisticEncryption = nullifyEmptyString(source.OpportunisticEncryption)
	target.OriginErrorPagePassThru = nullifyEmptyString(source.OriginErrorPagePassThru)
	target.Polish = nullifyEmptyString(source.Polish)
	target.ResolveOverride = nullifyEmptyString(source.ResolveOverride)
	target.RespectStrongEtag = nullifyEmptyString(source.RespectStrongEtag)
	target.ResponseBuffering = nullifyEmptyString(source.ResponseBuffering)
	target.RocketLoader = nullifyEmptyString(source.RocketLoader)
	target.SecurityLevel = nullifyEmptyString(source.SecurityLevel)
	target.SortQueryStringForCache = nullifyEmptyString(source.SortQueryStringForCache)
	target.SSL = nullifyEmptyString(source.SSL)
	target.TrueClientIPHeader = nullifyEmptyString(source.TrueClientIPHeader)
	target.WAF = nullifyEmptyString(source.WAF)

	// browser_cache_ttl: STRING in v4 → Int64 in v5
	if !source.BrowserCacheTTL.IsNull() && !source.BrowserCacheTTL.IsUnknown() {
		if intVal, err := strconv.ParseInt(source.BrowserCacheTTL.ValueString(), 10, 64); err == nil {
			target.BrowserCacheTTL = types.Int64Value(intVal)
		} else {
			// Invalid string, set to null
			target.BrowserCacheTTL = types.Int64Null()
		}
	} else {
		target.BrowserCacheTTL = types.Int64Null()
	}

	// edge_cache_ttl: Int64 (nullify 0 values - v4 default was 0, v5 uses null)
	target.EdgeCacheTTL = nullifyZeroInt64(source.EdgeCacheTTL)

	// forwarding_url: TypeList MaxItems:1 → SingleNestedAttribute
	if len(source.ForwardingURL) > 0 {
		target.ForwardingURL = &TargetForwardingURLModel{
			URL:        source.ForwardingURL[0].URL,
			StatusCode: source.ForwardingURL[0].StatusCode,
		}
	} else {
		target.ForwardingURL = nil
	}

	// cache_key_fields: 5-level deep nested transformation
	if len(source.CacheKeyFields) > 0 {
		ckf, ckfDiags := transformCacheKeyFields(ctx, source.CacheKeyFields[0])
		diags.Append(ckfDiags...)
		target.CacheKeyFields = ckf
	} else {
		target.CacheKeyFields = nil
	}

	// cache_ttl_by_status: Set[Object] → Map[String]
	if len(source.CacheTTLByStatus) > 0 {
		cacheTTLMap, cacheTTLDiags := transformCacheTTLByStatus(ctx, source.CacheTTLByStatus)
		diags.Append(cacheTTLDiags...)
		target.CacheTTLByStatus = cacheTTLMap
	} else {
		target.CacheTTLByStatus = types.MapNull(types.StringType)
	}

	return target, diags
}

// transformCacheKeyFields transforms 5-level deep cache_key_fields structure.
// Each level is a TypeList MaxItems:1 in v4 → SingleNestedAttribute (pointer) in v5.
func transformCacheKeyFields(ctx context.Context, source SourceCacheKeyFieldsModel) (*TargetCacheKeyFieldsModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	target := &TargetCacheKeyFieldsModel{}

	// Level 2: cookie (TypeList MaxItems:1 → pointer)
	if len(source.Cookie) > 0 {
		cookieTarget := &TargetCacheKeyFieldsCookieModel{}

		// Set → List conversion
		if !source.Cookie[0].CheckPresence.IsNull() && !source.Cookie[0].CheckPresence.IsUnknown() {
			checkPresence, setDiags := convertSetToStringSlice(ctx, source.Cookie[0].CheckPresence)
			diags.Append(setDiags...)
			if !diags.HasError() {
				cookieTarget.CheckPresence = checkPresence
			}
		}

		if !source.Cookie[0].Include.IsNull() && !source.Cookie[0].Include.IsUnknown() {
			include, setDiags := convertSetToStringSlice(ctx, source.Cookie[0].Include)
			diags.Append(setDiags...)
			if !diags.HasError() {
				cookieTarget.Include = include
			}
		}

		target.Cookie = cookieTarget
	} else {
		target.Cookie = nil
	}

	// Level 2: header (TypeList MaxItems:1 → pointer)
	if len(source.Header) > 0 {
		headerTarget := &TargetCacheKeyFieldsHeaderModel{}

		// Set → List conversions
		if !source.Header[0].CheckPresence.IsNull() && !source.Header[0].CheckPresence.IsUnknown() {
			checkPresence, setDiags := convertSetToStringSlice(ctx, source.Header[0].CheckPresence)
			diags.Append(setDiags...)
			if !diags.HasError() {
				headerTarget.CheckPresence = checkPresence
			}
		}

		if !source.Header[0].Include.IsNull() && !source.Header[0].Include.IsUnknown() {
			include, setDiags := convertSetToStringSlice(ctx, source.Header[0].Include)
			diags.Append(setDiags...)
			if !diags.HasError() {
				headerTarget.Include = include
			}
		}

		if !source.Header[0].Exclude.IsNull() && !source.Header[0].Exclude.IsUnknown() {
			exclude, setDiags := convertSetToStringSlice(ctx, source.Header[0].Exclude)
			diags.Append(setDiags...)
			if !diags.HasError() {
				headerTarget.Exclude = exclude
			}
		}

		target.Header = headerTarget
	} else {
		target.Header = nil
	}

	// Level 2: host (TypeList MaxItems:1 → pointer)
	if len(source.Host) > 0 {
		hostTarget := &TargetCacheKeyFieldsHostModel{
			Resolved: source.Host[0].Resolved,
		}
		target.Host = hostTarget
	} else {
		target.Host = nil
	}

	// Level 2: query_string (TypeList MaxItems:1 → pointer)
	// NOTE: The 'ignore' field from v4 is intentionally dropped (removed in v5)
	if len(source.QueryString) > 0 {
		qsTarget := &TargetCacheKeyFieldsQueryStringModel{}

		// Set → List conversions
		if !source.QueryString[0].Include.IsNull() && !source.QueryString[0].Include.IsUnknown() {
			include, setDiags := convertSetToStringSlice(ctx, source.QueryString[0].Include)
			diags.Append(setDiags...)
			if !diags.HasError() {
				qsTarget.Include = include
			}
		}

		if !source.QueryString[0].Exclude.IsNull() && !source.QueryString[0].Exclude.IsUnknown() {
			exclude, setDiags := convertSetToStringSlice(ctx, source.QueryString[0].Exclude)
			diags.Append(setDiags...)
			if !diags.HasError() {
				qsTarget.Exclude = exclude
			}
		}

		target.QueryString = qsTarget
	} else {
		target.QueryString = nil
	}

	// Level 2: user (TypeList MaxItems:1 → pointer)
	// CRITICAL: v4 may not have 'lang' field - must add with default false
	if len(source.User) > 0 {
		userTarget := &TargetCacheKeyFieldsUserModel{
			DeviceType: source.User[0].DeviceType,
			Geo:        source.User[0].Geo,
		}

		// Handle missing lang field (v4 may not have this)
		if !source.User[0].Lang.IsNull() && !source.User[0].Lang.IsUnknown() {
			userTarget.Lang = source.User[0].Lang
		} else {
			// Add lang = false if missing (prevents drift)
			userTarget.Lang = types.BoolValue(false)
		}

		target.User = userTarget
	} else {
		target.User = nil
	}

	return target, diags
}

// transformCacheTTLByStatus transforms cache_ttl_by_status from Set[Object] to Map[String].
// v4: [{codes: "200", ttl: 3600}, {codes: "404", ttl: 300}]
// v5: {"200": "3600", "404": "300"}
func transformCacheTTLByStatus(ctx context.Context, source []SourceCacheTTLByStatusModel) (types.Map, diag.Diagnostics) {
	var diags diag.Diagnostics

	// Build map from array of objects
	mapValues := make(map[string]attr.Value, len(source))
	for _, entry := range source {
		if !entry.Codes.IsNull() && !entry.Codes.IsUnknown() && !entry.TTL.IsNull() && !entry.TTL.IsUnknown() {
			// Key: codes (String)
			key := entry.Codes.ValueString()
			// Value: ttl as String (Int64 → String conversion)
			value := strconv.FormatInt(entry.TTL.ValueInt64(), 10)
			mapValues[key] = types.StringValue(value)
		}
	}

	if len(mapValues) == 0 {
		return types.MapNull(types.StringType), diags
	}

	mapValue, mapDiags := types.MapValue(types.StringType, mapValues)
	diags.Append(mapDiags...)
	return mapValue, diags
}

// ============================================================================
// Helper Functions
// ============================================================================

// nullifyFalseBool converts false → null for boolean fields.
// v4 had default: false, v5 stores null instead to reduce state size.
func nullifyFalseBool(val types.Bool) types.Bool {
	if val.IsNull() || val.IsUnknown() {
		return types.BoolNull()
	}
	if !val.ValueBool() {
		// false → null
		return types.BoolNull()
	}
	// true → true
	return val
}

// convertSetToStringSlice converts types.Set to []types.String for Set[String] → List[String] conversions.
// CRITICAL: Extract directly to []string then convert to []types.String to avoid attr.Value issues.
func convertSetToStringSlice(ctx context.Context, set types.Set) ([]types.String, diag.Diagnostics) {
	var diags diag.Diagnostics

	// Extract to []string first
	var rawStrings []string
	diags.Append(set.ElementsAs(ctx, &rawStrings, false)...)
	if diags.HasError() {
		return nil, diags
	}

	// Convert []string to []types.String
	result := make([]types.String, 0, len(rawStrings))
	for _, str := range rawStrings {
		result = append(result, types.StringValue(str))
	}
	return result, diags
}

// nullifyEmptyString converts empty strings to null.
// v4 may store empty strings for unset fields, v5 expects null.
func nullifyEmptyString(val types.String) types.String {
	if val.IsNull() || val.IsUnknown() {
		return types.StringNull()
	}
	if val.ValueString() == "" {
		// empty string → null
		return types.StringNull()
	}
	// non-empty string → keep as-is
	return val
}

// nullifyZeroInt64 converts 0 values to null for Int64 fields.
// v4 may have default value of 0, v5 expects null for unset fields.
func nullifyZeroInt64(val types.Int64) types.Int64 {
	if val.IsNull() || val.IsUnknown() {
		return types.Int64Null()
	}
	if val.ValueInt64() == 0 {
		// 0 → null
		return types.Int64Null()
	}
	// non-zero → keep as-is
	return val
}

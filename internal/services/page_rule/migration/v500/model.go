package v500

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ============================================================================
// Source Models (Legacy Provider - v4.x / SDKv2)
// ============================================================================

// FlexibleV0Model represents version 0 state with Dynamic actions field.
// This allows parsing BOTH v4 (actions as array) and early v5 (actions as object) formats.
type FlexibleV0Model struct {
	ID       types.String  `tfsdk:"id"`
	ZoneID   types.String  `tfsdk:"zone_id"`
	Target   types.String  `tfsdk:"target"`
	Priority types.Int64   `tfsdk:"priority"`
	Status   types.String  `tfsdk:"status"`
	Actions  types.Dynamic `tfsdk:"actions"` // Can be array OR object
}

// SourceCloudflarePageRuleModel represents the legacy page_rule resource state from v4.x provider (SDKv2).
// Schema version: 3 (treated as v3 for SDKv2 resources)
// Resource type: cloudflare_page_rule
//
// NOTE: In SDKv2, TypeList with MaxItems:1 is stored as an array with 1 element in state.
// This model reflects that storage format.
type SourceCloudflarePageRuleModel struct {
	ID       types.String `tfsdk:"id"`
	ZoneID   types.String `tfsdk:"zone_id"`
	Target   types.String `tfsdk:"target"`
	Priority types.Int64  `tfsdk:"priority"`
	Status   types.String `tfsdk:"status"`

	// SDKv2 TypeList MaxItems:1 is stored as []SourceActionsModel
	Actions []SourceActionsModel `tfsdk:"actions"`
}

// SourceActionsModel represents the actions block from v4.x provider (SDKv2).
// In v4, this was a TypeList MaxItems:1, stored as array[0] in state.
type SourceActionsModel struct {
	// Boolean fields (with default: false in v4)
	AlwaysUseHTTPS     types.Bool `tfsdk:"always_use_https"`
	DisableApps        types.Bool `tfsdk:"disable_apps"`
	DisablePerformance types.Bool `tfsdk:"disable_performance"`
	DisableRailgun     types.Bool `tfsdk:"disable_railgun"` // Deprecated - will be dropped
	DisableSecurity    types.Bool `tfsdk:"disable_security"`
	DisableZaraz       types.Bool `tfsdk:"disable_zaraz"`

	// String fields (on/off values)
	AutomaticHTTPSRewrites  types.String `tfsdk:"automatic_https_rewrites"`
	BrowserCheck            types.String `tfsdk:"browser_check"`
	CacheByDeviceType       types.String `tfsdk:"cache_by_device_type"`
	CacheDeceptionArmor     types.String `tfsdk:"cache_deception_armor"`
	EmailObfuscation        types.String `tfsdk:"email_obfuscation"`
	ExplicitCacheControl    types.String `tfsdk:"explicit_cache_control"`
	IPGeolocation           types.String `tfsdk:"ip_geolocation"`
	Mirage                  types.String `tfsdk:"mirage"`
	OpportunisticEncryption types.String `tfsdk:"opportunistic_encryption"`
	OriginErrorPagePassThru types.String `tfsdk:"origin_error_page_pass_thru"`
	RespectStrongEtag       types.String `tfsdk:"respect_strong_etag"`
	ResponseBuffering       types.String `tfsdk:"response_buffering"`
	RocketLoader            types.String `tfsdk:"rocket_loader"`
	ServerSideExclude       types.String `tfsdk:"server_side_exclude"`
	SortQueryStringForCache types.String `tfsdk:"sort_query_string_for_cache"`
	TrueClientIPHeader      types.String `tfsdk:"true_client_ip_header"`
	WAF                     types.String `tfsdk:"waf"`

	// String fields (other)
	BypassCacheOnCookie types.String `tfsdk:"bypass_cache_on_cookie"`
	CacheLevel          types.String `tfsdk:"cache_level"`
	CacheOnCookie       types.String `tfsdk:"cache_on_cookie"`
	HostHeaderOverride  types.String `tfsdk:"host_header_override"`
	Polish              types.String `tfsdk:"polish"`
	ResolveOverride     types.String `tfsdk:"resolve_override"`
	SecurityLevel       types.String `tfsdk:"security_level"`
	SSL                 types.String `tfsdk:"ssl"`

	// NOTE: browser_cache_ttl is STRING in v4, but Int64 in v5!
	BrowserCacheTTL types.String `tfsdk:"browser_cache_ttl"`

	// Numeric fields
	EdgeCacheTTL types.Int64 `tfsdk:"edge_cache_ttl"`

	// Complex nested structures (TypeList MaxItems:1 in v4 = array in state)
	ForwardingURL []SourceForwardingURLModel `tfsdk:"forwarding_url"`
	Minify        []SourceMinifyModel        `tfsdk:"minify"` // Deprecated - will be dropped

	// cache_key_fields: TypeList MaxItems:1 (5 levels deep!)
	CacheKeyFields []SourceCacheKeyFieldsModel `tfsdk:"cache_key_fields"`

	// cache_ttl_by_status: TypeSet in v4 (array of objects) → Map in v5
	CacheTTLByStatus []SourceCacheTTLByStatusModel `tfsdk:"cache_ttl_by_status"`
}

// SourceForwardingURLModel represents forwarding_url block from v4.x (TypeList MaxItems:1).
type SourceForwardingURLModel struct {
	URL        types.String `tfsdk:"url"`
	StatusCode types.Int64  `tfsdk:"status_code"`
}

// SourceMinifyModel represents minify block from v4.x (DEPRECATED - will be dropped).
type SourceMinifyModel struct {
	JS   types.String `tfsdk:"js"`
	CSS  types.String `tfsdk:"css"`
	HTML types.String `tfsdk:"html"`
}

// SourceCacheKeyFieldsModel represents cache_key_fields from v4.x (TypeList MaxItems:1).
// Each nested block is also TypeList MaxItems:1, creating 5 levels of nesting!
type SourceCacheKeyFieldsModel struct {
	Cookie      []SourceCacheKeyFieldsCookieModel      `tfsdk:"cookie"`       // TypeList MaxItems:1
	Header      []SourceCacheKeyFieldsHeaderModel      `tfsdk:"header"`       // TypeList MaxItems:1
	Host        []SourceCacheKeyFieldsHostModel        `tfsdk:"host"`         // TypeList MaxItems:1
	QueryString []SourceCacheKeyFieldsQueryStringModel `tfsdk:"query_string"` // TypeList MaxItems:1
	User        []SourceCacheKeyFieldsUserModel        `tfsdk:"user"`         // TypeList MaxItems:1
}

// SourceCacheKeyFieldsCookieModel represents cookie block (TypeList MaxItems:1).
type SourceCacheKeyFieldsCookieModel struct {
	CheckPresence types.Set `tfsdk:"check_presence"` // Set[String] in v4
	Include       types.Set `tfsdk:"include"`        // Set[String] in v4
}

// SourceCacheKeyFieldsHeaderModel represents header block (TypeList MaxItems:1).
type SourceCacheKeyFieldsHeaderModel struct {
	CheckPresence types.Set `tfsdk:"check_presence"` // Set[String] in v4
	Include       types.Set `tfsdk:"include"`        // Set[String] in v4
	Exclude       types.Set `tfsdk:"exclude"`        // Set[String] in v4
}

// SourceCacheKeyFieldsHostModel represents host block (TypeList MaxItems:1).
type SourceCacheKeyFieldsHostModel struct {
	Resolved types.Bool `tfsdk:"resolved"`
}

// SourceCacheKeyFieldsQueryStringModel represents query_string block (TypeList MaxItems:1).
type SourceCacheKeyFieldsQueryStringModel struct {
	Include types.Set  `tfsdk:"include"` // Set[String] in v4
	Exclude types.Set  `tfsdk:"exclude"` // Set[String] in v4
	Ignore  types.Bool `tfsdk:"ignore"`
}

// SourceCacheKeyFieldsUserModel represents user block (TypeList MaxItems:1).
// CRITICAL: v4 configs may NOT have the 'lang' field!
type SourceCacheKeyFieldsUserModel struct {
	DeviceType types.Bool `tfsdk:"device_type"`
	Geo        types.Bool `tfsdk:"geo"`
	Lang       types.Bool `tfsdk:"lang"` // May be null in v4 state!
}

// SourceCacheTTLByStatusModel represents a cache_ttl_by_status entry from v4.x (TypeSet).
// In v4, this was stored as an array of objects: [{codes: "200", ttl: 3600}, ...]
// In v5, this becomes a map: {"200": "3600", ...}
type SourceCacheTTLByStatusModel struct {
	Codes types.String `tfsdk:"codes"`
	TTL   types.Int64  `tfsdk:"ttl"`
}

// ============================================================================
// Target Models (Current Provider - v5.x+ / Plugin Framework)
// ============================================================================

// TargetPageRuleModel represents the current page_rule resource state from v5.x+ provider.
// Schema version: 500 (will be set when Version field is added to schema)
// Resource type: cloudflare_page_rule (unchanged)
//
// NOTE: This matches the PageRuleModel in parent package, duplicated here for migration package isolation.
type TargetPageRuleModel struct {
	ID         types.String      `tfsdk:"id"`
	ZoneID     types.String      `tfsdk:"zone_id"`
	Target     types.String      `tfsdk:"target"`
	Priority   types.Int64       `tfsdk:"priority"`
	Status     types.String      `tfsdk:"status"`
	CreatedOn  timetypes.RFC3339 `tfsdk:"created_on"`
	ModifiedOn timetypes.RFC3339 `tfsdk:"modified_on"`

	// In v5, actions is a pointer to single nested object (not array!)
	Actions *TargetActionsModel `tfsdk:"actions"`
}

// TargetActionsModel represents the actions object in v5.x+ provider.
// In v5, this is a SingleNestedAttribute (not TypeList MaxItems:1).
type TargetActionsModel struct {
	// Boolean fields (no defaults, false → null)
	AlwaysUseHTTPS     types.Bool `tfsdk:"always_use_https"`
	DisableApps        types.Bool `tfsdk:"disable_apps"`
	DisablePerformance types.Bool `tfsdk:"disable_performance"`
	DisableSecurity    types.Bool `tfsdk:"disable_security"`
	DisableZaraz       types.Bool `tfsdk:"disable_zaraz"`
	// NOTE: disable_railgun removed (deprecated)
	// NOTE: minify removed (deprecated)

	// String fields (on/off values)
	AutomaticHTTPSRewrites  types.String `tfsdk:"automatic_https_rewrites"`
	BrowserCheck            types.String `tfsdk:"browser_check"`
	CacheByDeviceType       types.String `tfsdk:"cache_by_device_type"`
	CacheDeceptionArmor     types.String `tfsdk:"cache_deception_armor"`
	EmailObfuscation        types.String `tfsdk:"email_obfuscation"`
	ExplicitCacheControl    types.String `tfsdk:"explicit_cache_control"`
	IPGeolocation           types.String `tfsdk:"ip_geolocation"`
	Mirage                  types.String `tfsdk:"mirage"`
	OpportunisticEncryption types.String `tfsdk:"opportunistic_encryption"`
	OriginErrorPagePassThru types.String `tfsdk:"origin_error_page_pass_thru"`
	RespectStrongEtag       types.String `tfsdk:"respect_strong_etag"`
	ResponseBuffering       types.String `tfsdk:"response_buffering"`
	RocketLoader            types.String `tfsdk:"rocket_loader"`
	SortQueryStringForCache types.String `tfsdk:"sort_query_string_for_cache"`
	TrueClientIPHeader      types.String `tfsdk:"true_client_ip_header"`
	WAF                     types.String `tfsdk:"waf"`

	// String fields (other)
	BypassCacheOnCookie types.String `tfsdk:"bypass_cache_on_cookie"`
	CacheLevel          types.String `tfsdk:"cache_level"`
	CacheOnCookie       types.String `tfsdk:"cache_on_cookie"`
	HostHeaderOverride  types.String `tfsdk:"host_header_override"`
	Polish              types.String `tfsdk:"polish"`
	ResolveOverride     types.String `tfsdk:"resolve_override"`
	SecurityLevel       types.String `tfsdk:"security_level"`
	SSL                 types.String `tfsdk:"ssl"`

	// NOTE: browser_cache_ttl is Int64 in v5 (was String in v4!)
	BrowserCacheTTL types.Int64 `tfsdk:"browser_cache_ttl"`

	// Numeric fields
	EdgeCacheTTL types.Int64 `tfsdk:"edge_cache_ttl"`

	// Complex nested structures (SingleNestedAttribute in v5)
	// NOTE: Using pointers instead of customfield.NestedObject to avoid type mismatch issues
	// during state upgrade. Terraform framework matches by tfsdk tags, not types.
	ForwardingURL *TargetForwardingURLModel `tfsdk:"forwarding_url"`

	// cache_key_fields: SingleNestedAttribute (pointer for JSON compatibility)
	CacheKeyFields *TargetCacheKeyFieldsModel `tfsdk:"cache_key_fields"`

	// cache_ttl_by_status: Map[String] in v5 (was Set[Object] in v4)
	CacheTTLByStatus types.Map `tfsdk:"cache_ttl_by_status"`
}

// TargetForwardingURLModel represents forwarding_url in v5 (SingleNestedAttribute).
type TargetForwardingURLModel struct {
	URL        types.String `tfsdk:"url"`
	StatusCode types.Int64  `tfsdk:"status_code"`
}

// TargetCacheKeyFieldsModel represents cache_key_fields in v5 (SingleNestedAttribute).
// NOTE: Using pointers for nested objects to avoid customfield type mismatch issues.
type TargetCacheKeyFieldsModel struct {
	Cookie      *TargetCacheKeyFieldsCookieModel      `tfsdk:"cookie"`
	Header      *TargetCacheKeyFieldsHeaderModel      `tfsdk:"header"`
	Host        *TargetCacheKeyFieldsHostModel        `tfsdk:"host"`
	QueryString *TargetCacheKeyFieldsQueryStringModel `tfsdk:"query_string"`
	User        *TargetCacheKeyFieldsUserModel        `tfsdk:"user"`
}

// TargetCacheKeyFieldsCookieModel represents cookie in v5.
type TargetCacheKeyFieldsCookieModel struct {
	CheckPresence []types.String `tfsdk:"check_presence"` // List[String] in v5 (was Set in v4)
	Include       []types.String `tfsdk:"include"`        // List[String] in v5 (was Set in v4)
}

// TargetCacheKeyFieldsHeaderModel represents header in v5.
type TargetCacheKeyFieldsHeaderModel struct {
	CheckPresence []types.String `tfsdk:"check_presence"` // List[String] in v5 (was Set in v4)
	Include       []types.String `tfsdk:"include"`        // List[String] in v5 (was Set in v4)
	Exclude       []types.String `tfsdk:"exclude"`        // List[String] in v5 (was Set in v4)
}

// TargetCacheKeyFieldsHostModel represents host in v5.
type TargetCacheKeyFieldsHostModel struct {
	Resolved types.Bool `tfsdk:"resolved"`
}

// TargetCacheKeyFieldsQueryStringModel represents query_string in v5.
// NOTE: The 'ignore' field from v4 was removed in v5 - do not include it
type TargetCacheKeyFieldsQueryStringModel struct {
	Include []types.String `tfsdk:"include"` // List[String] in v5 (was Set in v4)
	Exclude []types.String `tfsdk:"exclude"` // List[String] in v5 (was Set in v4)
}

// TargetCacheKeyFieldsUserModel represents user in v5.
// CRITICAL: Must always have lang field in v5!
type TargetCacheKeyFieldsUserModel struct {
	DeviceType types.Bool `tfsdk:"device_type"`
	Geo        types.Bool `tfsdk:"geo"`
	Lang       types.Bool `tfsdk:"lang"` // Required in v5 (may be missing in v4)
}

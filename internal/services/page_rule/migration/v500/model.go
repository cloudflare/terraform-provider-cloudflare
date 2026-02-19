package v500

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ============================================================================
// Source V4 Models (Legacy SDKv2 Provider)
// ============================================================================

// SourceV4PageRuleModel represents the page_rule state from v4.x provider (SDKv2).
// Schema version: 0
//
// In SDKv2, TypeList with MaxItems:1 is stored as an array with 1 element.
type SourceV4PageRuleModel struct {
	ID       types.String `tfsdk:"id"`
	ZoneID   types.String `tfsdk:"zone_id"`
	Target   types.String `tfsdk:"target"`
	Priority types.Int64  `tfsdk:"priority"`
	Status   types.String `tfsdk:"status"`

	CreatedOn  timetypes.RFC3339 `tfsdk:"created_on"`
	ModifiedOn timetypes.RFC3339 `tfsdk:"modified_on"`

	Actions []SourceV4ActionsModel `tfsdk:"actions"` // TypeList MaxItems:1 = array
}

// SourceV4ActionsModel represents the actions block from v4.x (TypeList MaxItems:1).
type SourceV4ActionsModel struct {
	// Boolean fields
	AlwaysUseHTTPS     types.Bool `tfsdk:"always_use_https"`
	DisableApps        types.Bool `tfsdk:"disable_apps"`
	DisablePerformance types.Bool `tfsdk:"disable_performance"`
	DisableRailgun     types.Bool `tfsdk:"disable_railgun"` // Deprecated in v5
	DisableSecurity    types.Bool `tfsdk:"disable_security"`
	DisableZaraz       types.Bool `tfsdk:"disable_zaraz"`

	// String fields (on/off)
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

	// browser_cache_ttl is STRING in v4, Int64 in v5
	BrowserCacheTTL types.String `tfsdk:"browser_cache_ttl"`

	// Numeric fields
	EdgeCacheTTL types.Int64 `tfsdk:"edge_cache_ttl"`

	// Nested structures (TypeList MaxItems:1 = array)
	ForwardingURL    []SourceV4ForwardingURLModel    `tfsdk:"forwarding_url"`
	Minify           []SourceV4MinifyModel           `tfsdk:"minify"` // Deprecated in v5
	CacheKeyFields   []SourceV4CacheKeyFieldsModel   `tfsdk:"cache_key_fields"`
	CacheTTLByStatus []SourceV4CacheTTLByStatusModel `tfsdk:"cache_ttl_by_status"` // TypeSet = array
}

type SourceV4ForwardingURLModel struct {
	URL        types.String `tfsdk:"url"`
	StatusCode types.Int64  `tfsdk:"status_code"`
}

type SourceV4MinifyModel struct {
	JS   types.String `tfsdk:"js"`
	CSS  types.String `tfsdk:"css"`
	HTML types.String `tfsdk:"html"`
}

type SourceV4CacheKeyFieldsModel struct {
	Cookie      []SourceV4CacheKeyFieldsCookieModel      `tfsdk:"cookie"`
	Header      []SourceV4CacheKeyFieldsHeaderModel      `tfsdk:"header"`
	Host        []SourceV4CacheKeyFieldsHostModel        `tfsdk:"host"`
	QueryString []SourceV4CacheKeyFieldsQueryStringModel `tfsdk:"query_string"`
	User        []SourceV4CacheKeyFieldsUserModel        `tfsdk:"user"`
}

type SourceV4CacheKeyFieldsCookieModel struct {
	CheckPresence types.Set `tfsdk:"check_presence"`
	Include       types.Set `tfsdk:"include"`
}

type SourceV4CacheKeyFieldsHeaderModel struct {
	CheckPresence types.Set `tfsdk:"check_presence"`
	Include       types.Set `tfsdk:"include"`
	Exclude       types.Set `tfsdk:"exclude"`
}

type SourceV4CacheKeyFieldsHostModel struct {
	Resolved types.Bool `tfsdk:"resolved"`
}

type SourceV4CacheKeyFieldsQueryStringModel struct {
	Include types.Set  `tfsdk:"include"`
	Exclude types.Set  `tfsdk:"exclude"`
	Ignore  types.Bool `tfsdk:"ignore"` // Removed in v5
}

type SourceV4CacheKeyFieldsUserModel struct {
	DeviceType types.Bool `tfsdk:"device_type"`
	Geo        types.Bool `tfsdk:"geo"`
	Lang       types.Bool `tfsdk:"lang"`
}

type SourceV4CacheTTLByStatusModel struct {
	Codes types.String `tfsdk:"codes"`
	TTL   types.Int64  `tfsdk:"ttl"`
}

// ============================================================================
// Target V5 Models (Current Plugin Framework Provider)
// ============================================================================

// TargetV5PageRuleModel represents the page_rule state from v5.x+ provider (Plugin Framework).
// Schema version: 500
type TargetV5PageRuleModel struct {
	ID         types.String      `tfsdk:"id"`
	ZoneID     types.String      `tfsdk:"zone_id"`
	Target     types.String      `tfsdk:"target"`
	Priority   types.Int64       `tfsdk:"priority"`
	Status     types.String      `tfsdk:"status"`
	CreatedOn  timetypes.RFC3339 `tfsdk:"created_on"`
	ModifiedOn timetypes.RFC3339 `tfsdk:"modified_on"`

	Actions *TargetV5ActionsModel `tfsdk:"actions"` // SingleNestedAttribute = pointer
}

// TargetV5ActionsModel represents the actions object in v5.x+ (SingleNestedAttribute).
type TargetV5ActionsModel struct {
	// Boolean fields
	AlwaysUseHTTPS     types.Bool `tfsdk:"always_use_https"`
	DisableApps        types.Bool `tfsdk:"disable_apps"`
	DisablePerformance types.Bool `tfsdk:"disable_performance"`
	DisableSecurity    types.Bool `tfsdk:"disable_security"`
	DisableZaraz       types.Bool `tfsdk:"disable_zaraz"`

	// String fields (on/off)
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

	// browser_cache_ttl is Int64 in v5 (was String in v4)
	BrowserCacheTTL types.Int64 `tfsdk:"browser_cache_ttl"`

	// Numeric fields
	EdgeCacheTTL types.Int64 `tfsdk:"edge_cache_ttl"`

	// Nested structures (SingleNestedAttribute = pointer)
	ForwardingURL  *TargetV5ForwardingURLModel  `tfsdk:"forwarding_url"`
	CacheKeyFields *TargetV5CacheKeyFieldsModel `tfsdk:"cache_key_fields"`

	// cache_ttl_by_status: Map[String] in v5 (was Set[Object] in v4)
	CacheTTLByStatus types.Map `tfsdk:"cache_ttl_by_status"`
}

type TargetV5ForwardingURLModel struct {
	URL        types.String `tfsdk:"url"`
	StatusCode types.Int64  `tfsdk:"status_code"`
}

type TargetV5CacheKeyFieldsModel struct {
	Cookie      *TargetV5CacheKeyFieldsCookieModel      `tfsdk:"cookie"`
	Header      *TargetV5CacheKeyFieldsHeaderModel      `tfsdk:"header"`
	Host        *TargetV5CacheKeyFieldsHostModel        `tfsdk:"host"`
	QueryString *TargetV5CacheKeyFieldsQueryStringModel `tfsdk:"query_string"`
	User        *TargetV5CacheKeyFieldsUserModel        `tfsdk:"user"`
}

type TargetV5CacheKeyFieldsCookieModel struct {
	CheckPresence []types.String `tfsdk:"check_presence"`
	Include       []types.String `tfsdk:"include"`
}

type TargetV5CacheKeyFieldsHeaderModel struct {
	CheckPresence []types.String `tfsdk:"check_presence"`
	Include       []types.String `tfsdk:"include"`
	Exclude       []types.String `tfsdk:"exclude"`
}

type TargetV5CacheKeyFieldsHostModel struct {
	Resolved types.Bool `tfsdk:"resolved"`
}

type TargetV5CacheKeyFieldsQueryStringModel struct {
	Include []types.String `tfsdk:"include"`
	Exclude []types.String `tfsdk:"exclude"`
}

type TargetV5CacheKeyFieldsUserModel struct {
	DeviceType types.Bool `tfsdk:"device_type"`
	Geo        types.Bool `tfsdk:"geo"`
	Lang       types.Bool `tfsdk:"lang"`
}

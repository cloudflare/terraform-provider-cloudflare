package v500

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ============================================================================
// Source V4 Models (Legacy Plugin Framework Provider)
// ============================================================================

// SourceV4RulesetModel represents the ruleset state from the v4 Plugin Framework provider.
// Schema version: 1 (v4 had an in-provider V0→V1 migration for ratelimit field rename)
//
// In Plugin Framework, ListNestedBlock with SizeAtMost(1) is stored as an array
// with 0 or 1 elements in state. []*T captures this correctly.
type SourceV4RulesetModel struct {
	AccountID   types.String         `tfsdk:"account_id"`
	Description types.String         `tfsdk:"description"`
	ID          types.String         `tfsdk:"id"`
	Kind        types.String         `tfsdk:"kind"`
	Name        types.String         `tfsdk:"name"`
	Phase       types.String         `tfsdk:"phase"`
	Rules       []*SourceV4RuleModel `tfsdk:"rules"`
	ZoneID      types.String         `tfsdk:"zone_id"`
}

// SourceV4RuleModel represents a single rule in the v4 state.
type SourceV4RuleModel struct {
	Action                 types.String                           `tfsdk:"action"`
	ActionParameters       []*SourceV4ActionParametersModel       `tfsdk:"action_parameters"`
	Description            types.String                           `tfsdk:"description"`
	Enabled                types.Bool                             `tfsdk:"enabled"`
	ExposedCredentialCheck []*SourceV4ExposedCredentialCheckModel `tfsdk:"exposed_credential_check"`
	Expression             types.String                           `tfsdk:"expression"`
	ID                     types.String                           `tfsdk:"id"`
	Logging                []*SourceV4LoggingModel                `tfsdk:"logging"`
	Ratelimit              []*SourceV4RatelimitModel              `tfsdk:"ratelimit"`
	Ref                    types.String                           `tfsdk:"ref"`
}

// SourceV4ActionParametersModel represents the action_parameters block in v4.
// Note: In v4, this was a ListNestedBlock (MaxItems:1), stored as array in state.
type SourceV4ActionParametersModel struct {
	AdditionalCacheablePorts types.Set                    `tfsdk:"additional_cacheable_ports"`
	AutomaticHTTPSRewrites   types.Bool                   `tfsdk:"automatic_https_rewrites"`
	AutoMinify               []*SourceV4AutoMinifyModel   `tfsdk:"autominify"`
	BIC                      types.Bool                   `tfsdk:"bic"`
	BrowserTTL               []*SourceV4BrowserTTLModel   `tfsdk:"browser_ttl"`
	Cache                    types.Bool                   `tfsdk:"cache"`
	CacheKey                 []*SourceV4CacheKeyModel     `tfsdk:"cache_key"`
	CacheReserve             []*SourceV4CacheReserveModel `tfsdk:"cache_reserve"`
	Content                  types.String                 `tfsdk:"content"`
	ContentType              types.String                 `tfsdk:"content_type"`
	CookieFields             types.Set                    `tfsdk:"cookie_fields"`
	DisableApps              types.Bool                   `tfsdk:"disable_apps"`
	DisableRailgun           types.Bool                   `tfsdk:"disable_railgun"`
	DisableRUM               types.Bool                   `tfsdk:"disable_rum"`
	DisableZaraz             types.Bool                   `tfsdk:"disable_zaraz"`
	EdgeTTL                  []*SourceV4EdgeTTLModel      `tfsdk:"edge_ttl"`
	EmailObfuscation         types.Bool                   `tfsdk:"email_obfuscation"`
	Fonts                    types.Bool                   `tfsdk:"fonts"`
	FromList                 []*SourceV4FromListModel     `tfsdk:"from_list"`
	FromValue                []*SourceV4FromValueModel    `tfsdk:"from_value"`
	Headers                  []*SourceV4HeaderModel       `tfsdk:"headers"`
	HostHeader               types.String                 `tfsdk:"host_header"`
	HotlinkProtection        types.Bool                   `tfsdk:"hotlink_protection"`
	ID                       types.String                 `tfsdk:"id"`
	Increment                types.Int64                  `tfsdk:"increment"`
	MatchedData              []*SourceV4MatchedDataModel  `tfsdk:"matched_data"`
	Mirage                   types.Bool                   `tfsdk:"mirage"`
	OpportunisticEncryption  types.Bool                   `tfsdk:"opportunistic_encryption"`
	Origin                   []*SourceV4OriginModel       `tfsdk:"origin"`
	OriginCacheControl       types.Bool                   `tfsdk:"origin_cache_control"`
	OriginErrorPagePassthru  types.Bool                   `tfsdk:"origin_error_page_passthru"`
	Overrides                []*SourceV4OverridesModel    `tfsdk:"overrides"`
	Phases                   types.Set                    `tfsdk:"phases"`
	Polish                   types.String                 `tfsdk:"polish"`
	Products                 types.Set                    `tfsdk:"products"`
	ReadTimeout              types.Int64                  `tfsdk:"read_timeout"`
	RequestFields            types.Set                    `tfsdk:"request_fields"`
	RespectStrongEtags       types.Bool                   `tfsdk:"respect_strong_etags"`
	Response                 []*SourceV4ResponseModel     `tfsdk:"response"`
	ResponseFields           types.Set                    `tfsdk:"response_fields"`
	RocketLoader             types.Bool                   `tfsdk:"rocket_loader"`
	RedirectsForAITraining   types.Bool                   `tfsdk:"redirects_for_ai_training"`
	Rules                    map[string]types.String      `tfsdk:"rules"`
	Ruleset                  types.String                 `tfsdk:"ruleset"`
	Rulesets                 types.Set                    `tfsdk:"rulesets"`
	SecurityLevel            types.String                 `tfsdk:"security_level"`
	ServerSideExcludes       types.Bool                   `tfsdk:"server_side_excludes"`
	ServeStale               []*SourceV4ServeStaleModel   `tfsdk:"serve_stale"`
	SNI                      []*SourceV4SNIModel          `tfsdk:"sni"`
	SSL                      types.String                 `tfsdk:"ssl"`
	StatusCode               types.Int64                  `tfsdk:"status_code"`
	SXG                      types.Bool                   `tfsdk:"sxg"`
	URI                      []*SourceV4URIModel          `tfsdk:"uri"`
	Algorithms               []*SourceV4AlgorithmModel    `tfsdk:"algorithms"`
}

type SourceV4AutoMinifyModel struct {
	HTML types.Bool `tfsdk:"html"`
	CSS  types.Bool `tfsdk:"css"`
	JS   types.Bool `tfsdk:"js"`
}

type SourceV4BrowserTTLModel struct {
	Mode    types.String `tfsdk:"mode"`
	Default types.Int64  `tfsdk:"default"`
}

type SourceV4CacheKeyModel struct {
	CacheByDeviceType       types.Bool                        `tfsdk:"cache_by_device_type"`
	IgnoreQueryStringsOrder types.Bool                        `tfsdk:"ignore_query_strings_order"`
	CacheDeceptionArmor     types.Bool                        `tfsdk:"cache_deception_armor"`
	CustomKey               []*SourceV4CacheKeyCustomKeyModel `tfsdk:"custom_key"`
}

type SourceV4CacheKeyCustomKeyModel struct {
	QueryString []*SourceV4QueryStringModel     `tfsdk:"query_string"`
	Header      []*SourceV4CustomKeyHeaderModel `tfsdk:"header"`
	Cookie      []*SourceV4CustomKeyCookieModel `tfsdk:"cookie"`
	User        []*SourceV4CustomKeyUserModel   `tfsdk:"user"`
	Host        []*SourceV4CustomKeyHostModel   `tfsdk:"host"`
}

type SourceV4QueryStringModel struct {
	Include types.Set `tfsdk:"include"`
	Exclude types.Set `tfsdk:"exclude"`
}

type SourceV4CustomKeyHeaderModel struct {
	Include       types.Set            `tfsdk:"include"`
	CheckPresence types.Set            `tfsdk:"check_presence"`
	ExcludeOrigin types.Bool           `tfsdk:"exclude_origin"`
	Contains      map[string]types.Set `tfsdk:"contains"`
}

type SourceV4CustomKeyCookieModel struct {
	Include       types.Set `tfsdk:"include"`
	CheckPresence types.Set `tfsdk:"check_presence"`
}

type SourceV4CustomKeyUserModel struct {
	DeviceType types.Bool `tfsdk:"device_type"`
	Geo        types.Bool `tfsdk:"geo"`
	Lang       types.Bool `tfsdk:"lang"`
}

type SourceV4CustomKeyHostModel struct {
	Resolved types.Bool `tfsdk:"resolved"`
}

type SourceV4CacheReserveModel struct {
	Eligible        types.Bool  `tfsdk:"eligible"`
	MinimumFileSize types.Int64 `tfsdk:"minimum_file_size"`
}

type SourceV4EdgeTTLModel struct {
	Mode          types.String                  `tfsdk:"mode"`
	Default       types.Int64                   `tfsdk:"default"`
	StatusCodeTTL []*SourceV4StatusCodeTTLModel `tfsdk:"status_code_ttl"`
}

type SourceV4StatusCodeTTLModel struct {
	StatusCode      types.Int64                     `tfsdk:"status_code"`
	Value           types.Int64                     `tfsdk:"value"`
	StatusCodeRange []*SourceV4StatusCodeRangeModel `tfsdk:"status_code_range"`
}

type SourceV4StatusCodeRangeModel struct {
	From types.Int64 `tfsdk:"from"`
	To   types.Int64 `tfsdk:"to"`
}

type SourceV4FromListModel struct {
	Name types.String `tfsdk:"name"`
	Key  types.String `tfsdk:"key"`
}

type SourceV4FromValueModel struct {
	StatusCode          types.Int64               `tfsdk:"status_code"`
	PreserveQueryString types.Bool                `tfsdk:"preserve_query_string"`
	TargetURL           []*SourceV4TargetURLModel `tfsdk:"target_url"`
}

type SourceV4TargetURLModel struct {
	Value      types.String `tfsdk:"value"`
	Expression types.String `tfsdk:"expression"`
}

type SourceV4HeaderModel struct {
	Name       types.String `tfsdk:"name"`
	Value      types.String `tfsdk:"value"`
	Expression types.String `tfsdk:"expression"`
	Operation  types.String `tfsdk:"operation"`
}

type SourceV4MatchedDataModel struct {
	PublicKey types.String `tfsdk:"public_key"`
}

type SourceV4OriginModel struct {
	Host types.String `tfsdk:"host"`
	Port types.Int64  `tfsdk:"port"`
}

type SourceV4OverridesModel struct {
	Enabled          types.Bool                        `tfsdk:"enabled"`
	Action           types.String                      `tfsdk:"action"`
	SensitivityLevel types.String                      `tfsdk:"sensitivity_level"`
	Categories       []*SourceV4OverridesCategoryModel `tfsdk:"categories"`
	Rules            []*SourceV4OverridesRuleModel     `tfsdk:"rules"`
}

type SourceV4OverridesCategoryModel struct {
	Category types.String `tfsdk:"category"`
	Action   types.String `tfsdk:"action"`
	Enabled  types.Bool   `tfsdk:"enabled"`
}

type SourceV4OverridesRuleModel struct {
	ID               types.String `tfsdk:"id"`
	Action           types.String `tfsdk:"action"`
	Enabled          types.Bool   `tfsdk:"enabled"`
	ScoreThreshold   types.Int64  `tfsdk:"score_threshold"`
	SensitivityLevel types.String `tfsdk:"sensitivity_level"`
}

type SourceV4ResponseModel struct {
	StatusCode  types.Int64  `tfsdk:"status_code"`
	ContentType types.String `tfsdk:"content_type"`
	Content     types.String `tfsdk:"content"`
}

type SourceV4ServeStaleModel struct {
	DisableStaleWhileUpdating types.Bool `tfsdk:"disable_stale_while_updating"`
}

type SourceV4SNIModel struct {
	Value types.String `tfsdk:"value"`
}

type SourceV4URIModel struct {
	Path   []*SourceV4URIPartModel `tfsdk:"path"`
	Query  []*SourceV4URIPartModel `tfsdk:"query"`
	Origin types.Bool              `tfsdk:"origin"`
}

type SourceV4URIPartModel struct {
	Value      types.String `tfsdk:"value"`
	Expression types.String `tfsdk:"expression"`
}

// SourceV4AlgorithmModel has Name as raw string (not types.String) - matches v4 model exactly.
type SourceV4AlgorithmModel struct {
	Name string `tfsdk:"name"`
}

type SourceV4RatelimitModel struct {
	Characteristics         types.Set    `tfsdk:"characteristics"`
	CountingExpression      types.String `tfsdk:"counting_expression"`
	MitigationTimeout       types.Int64  `tfsdk:"mitigation_timeout"`
	Period                  types.Int64  `tfsdk:"period"`
	RequestsPerPeriod       types.Int64  `tfsdk:"requests_per_period"`
	RequestsToOrigin        types.Bool   `tfsdk:"requests_to_origin"`
	ScorePerPeriod          types.Int64  `tfsdk:"score_per_period"`
	ScoreResponseHeaderName types.String `tfsdk:"score_response_header_name"`
}

type SourceV4ExposedCredentialCheckModel struct {
	PasswordExpression types.String `tfsdk:"password_expression"`
	UsernameExpression types.String `tfsdk:"username_expression"`
}

type SourceV4LoggingModel struct {
	Enabled types.Bool `tfsdk:"enabled"`
}

// ============================================================================
// Target V5 Models (Current Plugin Framework Provider)
// ============================================================================

// TargetV5RulesetModel represents the ruleset state for v5.x+ provider.
// Uses plain Go types (pointers, slices, maps) instead of customfield types
// since tfsdk reflection handles the conversion to the v5 schema's customfield types.
type TargetV5RulesetModel struct {
	ID          types.String         `tfsdk:"id"`
	AccountID   types.String         `tfsdk:"account_id"`
	ZoneID      types.String         `tfsdk:"zone_id"`
	Kind        types.String         `tfsdk:"kind"`
	Name        types.String         `tfsdk:"name"`
	Phase       types.String         `tfsdk:"phase"`
	Description types.String         `tfsdk:"description"`
	Rules       []*TargetV5RuleModel `tfsdk:"rules"`
	LastUpdated timetypes.RFC3339    `tfsdk:"last_updated"`
	Version     types.String         `tfsdk:"version"`
}

// TargetV5RuleModel matches v5 RulesetRulesModel tfsdk tags.
type TargetV5RuleModel struct {
	ID                     types.String                         `tfsdk:"id"`
	Action                 types.String                         `tfsdk:"action"`
	ActionParameters       *TargetV5ActionParametersModel       `tfsdk:"action_parameters"`
	Description            types.String                         `tfsdk:"description"`
	Enabled                types.Bool                           `tfsdk:"enabled"`
	ExposedCredentialCheck *TargetV5ExposedCredentialCheckModel `tfsdk:"exposed_credential_check"`
	Expression             types.String                         `tfsdk:"expression"`
	Logging                *TargetV5LoggingModel                `tfsdk:"logging"`
	Ratelimit              *TargetV5RatelimitModel              `tfsdk:"ratelimit"`
	Ref                    types.String                         `tfsdk:"ref"`
}

// TargetV5ActionParametersModel matches v5 RulesetRulesActionParametersModel tfsdk tags.
// Uses plain Go types:
//   - *T for customfield.NestedObject[T] (SingleNestedAttribute)
//   - []*T for customfield.NestedObjectList[T] (ListNestedAttribute)
//   - map[string]*T for customfield.NestedObjectMap[T] (MapNestedAttribute)
//   - []types.String for customfield.List[types.String] (ListAttribute)
//   - []types.Int64 for customfield.List[types.Int64] (ListAttribute)
//   - map[string][]types.String for customfield.Map[customfield.List[types.String]] (MapAttribute)
type TargetV5ActionParametersModel struct {
	// Nested objects (MaxItems:1 blocks → SingleNestedAttribute)
	Response     *TargetV5ResponseModel     `tfsdk:"response"`
	Autominify   *TargetV5AutoMinifyModel   `tfsdk:"autominify"`
	BrowserTTL   *TargetV5BrowserTTLModel   `tfsdk:"browser_ttl"`
	CacheKey     *TargetV5CacheKeyModel     `tfsdk:"cache_key"`
	CacheReserve *TargetV5CacheReserveModel `tfsdk:"cache_reserve"`
	EdgeTTL      *TargetV5EdgeTTLModel      `tfsdk:"edge_ttl"`
	FromList     *TargetV5FromListModel     `tfsdk:"from_list"`
	FromValue    *TargetV5FromValueModel    `tfsdk:"from_value"`
	MatchedData  *TargetV5MatchedDataModel  `tfsdk:"matched_data"`
	Overrides    *TargetV5OverridesModel    `tfsdk:"overrides"`
	Origin       *TargetV5OriginModel       `tfsdk:"origin"`
	SNI          *TargetV5SNIModel          `tfsdk:"sni"`
	ServeStale   *TargetV5ServeStaleModel   `tfsdk:"serve_stale"`
	URI          *TargetV5URIModel          `tfsdk:"uri"`

	// Nested object lists
	Algorithms               []*TargetV5AlgorithmModel        `tfsdk:"algorithms"`
	CookieFields             []*TargetV5FieldNameModel        `tfsdk:"cookie_fields"`
	RequestFields            []*TargetV5FieldNameModel        `tfsdk:"request_fields"`
	ResponseFields           []*TargetV5ResponseFieldModel    `tfsdk:"response_fields"`
	RawResponseFields        []*TargetV5RawResponseFieldModel `tfsdk:"raw_response_fields"`
	TransformedRequestFields []*TargetV5FieldNameModel        `tfsdk:"transformed_request_fields"`

	// Nested object map (headers: list → map keyed by name)
	Headers map[string]*TargetV5HeaderModel `tfsdk:"headers"`

	// Scalar fields
	ID                      types.String `tfsdk:"id"`
	HostHeader              types.String `tfsdk:"host_header"`
	Increment               types.Int64  `tfsdk:"increment"`
	Content                 types.String `tfsdk:"content"`
	ContentType             types.String `tfsdk:"content_type"`
	StatusCode              types.Int64  `tfsdk:"status_code"`
	AutomaticHTTPSRewrites  types.Bool   `tfsdk:"automatic_https_rewrites"`
	BIC                     types.Bool   `tfsdk:"bic"`
	Cache                   types.Bool   `tfsdk:"cache"`
	ContentConverter        types.Bool   `tfsdk:"content_converter"`
	DisableApps             types.Bool   `tfsdk:"disable_apps"`
	DisableRUM              types.Bool   `tfsdk:"disable_rum"`
	DisableZaraz            types.Bool   `tfsdk:"disable_zaraz"`
	EmailObfuscation        types.Bool   `tfsdk:"email_obfuscation"`
	Fonts                   types.Bool   `tfsdk:"fonts"`
	HotlinkProtection       types.Bool   `tfsdk:"hotlink_protection"`
	Mirage                  types.Bool   `tfsdk:"mirage"`
	OpportunisticEncryption types.Bool   `tfsdk:"opportunistic_encryption"`
	OriginCacheControl      types.Bool   `tfsdk:"origin_cache_control"`
	OriginErrorPagePassthru types.Bool   `tfsdk:"origin_error_page_passthru"`
	ReadTimeout             types.Int64  `tfsdk:"read_timeout"`
	RespectStrongEtags      types.Bool   `tfsdk:"respect_strong_etags"`
	RocketLoader            types.Bool   `tfsdk:"rocket_loader"`
	SecurityLevel           types.String `tfsdk:"security_level"`
	ServerSideExcludes      types.Bool   `tfsdk:"server_side_excludes"`
	SSL                     types.String `tfsdk:"ssl"`
	SXG                     types.Bool   `tfsdk:"sxg"`
	Polish                  types.String `tfsdk:"polish"`
	Ruleset                 types.String `tfsdk:"ruleset"`

	// New in v5 (absent in v4)
	AssetName              types.String `tfsdk:"asset_name"`
	RequestBodyBuffering   types.String `tfsdk:"request_body_buffering"`
	ResponseBodyBuffering  types.String `tfsdk:"response_body_buffering"`
	RedirectsForAITraining types.Bool   `tfsdk:"redirects_for_ai_training"`

	// Cache control fields (set_cache_settings action)
	StripETags        types.Bool `tfsdk:"strip_etags"`
	StripLastModified types.Bool `tfsdk:"strip_last_modified"`
	StripSetCookie    types.Bool `tfsdk:"strip_set_cookie"`

	// Cache directives
	MaxAge               *TargetV5CacheControlValueModel      `tfsdk:"max_age"`
	SMaxage              *TargetV5CacheControlValueModel      `tfsdk:"s_maxage"`
	StaleWhileRevalidate *TargetV5CacheControlValueModel      `tfsdk:"stale_while_revalidate"`
	StaleIfError         *TargetV5CacheControlValueModel      `tfsdk:"stale_if_error"`
	Private              *TargetV5CacheControlQualifiersModel `tfsdk:"private"`
	NoCache              *TargetV5CacheControlQualifiersModel `tfsdk:"no_cache"`
	MustRevalidate       *TargetV5CacheControlSimpleModel     `tfsdk:"must_revalidate"`
	ProxyRevalidate      *TargetV5CacheControlSimpleModel     `tfsdk:"proxy_revalidate"`
	MustUnderstand       *TargetV5CacheControlSimpleModel     `tfsdk:"must_understand"`
	NoTransform          *TargetV5CacheControlSimpleModel     `tfsdk:"no_transform"`
	Immutable            *TargetV5CacheControlSimpleModel     `tfsdk:"immutable"`
	NoStore              *TargetV5CacheControlSimpleModel     `tfsdk:"no_store"`
	Public               *TargetV5CacheControlSimpleModel     `tfsdk:"public"`

	// set_cache_tags fields
	Operation  types.String   `tfsdk:"operation"`
	Values     []types.String `tfsdk:"values"`
	Expression types.String   `tfsdk:"expression"`

	// List attributes (Set → List)
	Products                 []types.String `tfsdk:"products"`
	Phases                   []types.String `tfsdk:"phases"`
	Rulesets                 []types.String `tfsdk:"rulesets"`
	AdditionalCacheablePorts []types.Int64  `tfsdk:"additional_cacheable_ports"`

	// Map attribute (map[string]string → map[string]list)
	Rules map[string][]types.String `tfsdk:"rules"`
}

// TargetV5OverridesModel matches v5 RulesetRulesActionParametersOverridesModel.
type TargetV5OverridesModel struct {
	Action           types.String                      `tfsdk:"action"`
	Categories       []*TargetV5OverridesCategoryModel `tfsdk:"categories"`
	Enabled          types.Bool                        `tfsdk:"enabled"`
	Rules            []*TargetV5OverridesRuleModel     `tfsdk:"rules"`
	SensitivityLevel types.String                      `tfsdk:"sensitivity_level"`
}

type TargetV5OverridesCategoryModel struct {
	Category         types.String `tfsdk:"category"`
	Action           types.String `tfsdk:"action"`
	Enabled          types.Bool   `tfsdk:"enabled"`
	SensitivityLevel types.String `tfsdk:"sensitivity_level"` // New in v5
}

type TargetV5OverridesRuleModel struct {
	ID               types.String `tfsdk:"id"`
	Action           types.String `tfsdk:"action"`
	Enabled          types.Bool   `tfsdk:"enabled"`
	ScoreThreshold   types.Int64  `tfsdk:"score_threshold"`
	SensitivityLevel types.String `tfsdk:"sensitivity_level"`
}

type TargetV5OriginModel struct {
	Host types.String `tfsdk:"host"`
	Port types.Int64  `tfsdk:"port"`
}

type TargetV5SNIModel struct {
	Value types.String `tfsdk:"value"`
}

// TargetV5URIModel matches v5 RulesetRulesActionParametersURIModel.
// Note: v4 had origin bool in URI, v5 does NOT.
type TargetV5URIModel struct {
	Path  *TargetV5URIPartModel `tfsdk:"path"`
	Query *TargetV5URIPartModel `tfsdk:"query"`
}

type TargetV5URIPartModel struct {
	Value      types.String `tfsdk:"value"`
	Expression types.String `tfsdk:"expression"`
}

// TargetV5HeaderModel matches v5 RulesetRulesActionParametersHeadersModel.
// Note: name is NOT included (it becomes the map key in the map).
type TargetV5HeaderModel struct {
	Operation  types.String `tfsdk:"operation"`
	Value      types.String `tfsdk:"value"`
	Expression types.String `tfsdk:"expression"`
}

type TargetV5MatchedDataModel struct {
	PublicKey types.String `tfsdk:"public_key"`
}

type TargetV5ResponseModel struct {
	Content     types.String `tfsdk:"content"`
	ContentType types.String `tfsdk:"content_type"`
	StatusCode  types.Int64  `tfsdk:"status_code"`
}

type TargetV5AutoMinifyModel struct {
	CSS  types.Bool `tfsdk:"css"`
	HTML types.Bool `tfsdk:"html"`
	JS   types.Bool `tfsdk:"js"`
}

type TargetV5EdgeTTLModel struct {
	Default       types.Int64                   `tfsdk:"default"`
	Mode          types.String                  `tfsdk:"mode"`
	StatusCodeTTL []*TargetV5StatusCodeTTLModel `tfsdk:"status_code_ttl"`
}

type TargetV5StatusCodeTTLModel struct {
	Value           types.Int64                   `tfsdk:"value"`
	StatusCodeRange *TargetV5StatusCodeRangeModel `tfsdk:"status_code_range"`
	StatusCode      types.Int64                   `tfsdk:"status_code"`
}

type TargetV5StatusCodeRangeModel struct {
	From types.Int64 `tfsdk:"from"`
	To   types.Int64 `tfsdk:"to"`
}

type TargetV5BrowserTTLModel struct {
	Mode    types.String `tfsdk:"mode"`
	Default types.Int64  `tfsdk:"default"`
}

type TargetV5CacheKeyModel struct {
	CacheByDeviceType       types.Bool                      `tfsdk:"cache_by_device_type"`
	CacheDeceptionArmor     types.Bool                      `tfsdk:"cache_deception_armor"`
	CustomKey               *TargetV5CacheKeyCustomKeyModel `tfsdk:"custom_key"`
	IgnoreQueryStringsOrder types.Bool                      `tfsdk:"ignore_query_strings_order"`
}

type TargetV5CacheKeyCustomKeyModel struct {
	Cookie      *TargetV5CKCookieModel      `tfsdk:"cookie"`
	Header      *TargetV5CKHeaderModel      `tfsdk:"header"`
	Host        *TargetV5CKHostModel        `tfsdk:"host"`
	QueryString *TargetV5CKQueryStringModel `tfsdk:"query_string"`
	User        *TargetV5CKUserModel        `tfsdk:"user"`
}

type TargetV5CKCookieModel struct {
	CheckPresence []types.String `tfsdk:"check_presence"`
	Include       []types.String `tfsdk:"include"`
}

type TargetV5CKHeaderModel struct {
	CheckPresence []types.String            `tfsdk:"check_presence"`
	Contains      map[string][]types.String `tfsdk:"contains"`
	ExcludeOrigin types.Bool                `tfsdk:"exclude_origin"`
	Include       []types.String            `tfsdk:"include"`
}

type TargetV5CKHostModel struct {
	Resolved types.Bool `tfsdk:"resolved"`
}

type TargetV5CKQueryStringModel struct {
	Include *TargetV5QSIncludeModel `tfsdk:"include"`
	Exclude *TargetV5QSExcludeModel `tfsdk:"exclude"`
}

type TargetV5QSIncludeModel struct {
	List []types.String `tfsdk:"list"`
	All  types.Bool     `tfsdk:"all"`
}

type TargetV5QSExcludeModel struct {
	List []types.String `tfsdk:"list"`
	All  types.Bool     `tfsdk:"all"`
}

type TargetV5CKUserModel struct {
	DeviceType types.Bool `tfsdk:"device_type"`
	Geo        types.Bool `tfsdk:"geo"`
	Lang       types.Bool `tfsdk:"lang"`
}

type TargetV5CacheReserveModel struct {
	Eligible        types.Bool  `tfsdk:"eligible"`
	MinimumFileSize types.Int64 `tfsdk:"minimum_file_size"`
}

type TargetV5FromListModel struct {
	Key  types.String `tfsdk:"key"`
	Name types.String `tfsdk:"name"`
}

type TargetV5FromValueModel struct {
	PreserveQueryString types.Bool              `tfsdk:"preserve_query_string"`
	StatusCode          types.Int64             `tfsdk:"status_code"`
	TargetURL           *TargetV5TargetURLModel `tfsdk:"target_url"`
}

type TargetV5TargetURLModel struct {
	Value      types.String `tfsdk:"value"`
	Expression types.String `tfsdk:"expression"`
}

type TargetV5ServeStaleModel struct {
	DisableStaleWhileUpdating types.Bool `tfsdk:"disable_stale_while_updating"`
}

// Cache control directive models for set_cache_settings action
type TargetV5CacheControlValueModel struct {
	Operation      types.String `tfsdk:"operation"`
	Value          types.Int64  `tfsdk:"value"`
	CloudflareOnly types.Bool   `tfsdk:"cloudflare_only"`
}

type TargetV5CacheControlSimpleModel struct {
	Operation      types.String `tfsdk:"operation"`
	CloudflareOnly types.Bool   `tfsdk:"cloudflare_only"`
}

type TargetV5CacheControlQualifiersModel struct {
	Operation      types.String   `tfsdk:"operation"`
	Qualifiers     []types.String `tfsdk:"qualifiers"`
	CloudflareOnly types.Bool     `tfsdk:"cloudflare_only"`
}

// TargetV5FieldNameModel is used for cookie_fields, request_fields, transformed_request_fields.
type TargetV5FieldNameModel struct {
	Name types.String `tfsdk:"name"`
}

// TargetV5ResponseFieldModel is used for response_fields (has preserve_duplicates in v5).
type TargetV5ResponseFieldModel struct {
	Name               types.String `tfsdk:"name"`
	PreserveDuplicates types.Bool   `tfsdk:"preserve_duplicates"`
}

// TargetV5RawResponseFieldModel matches v5 RulesetRulesActionParametersRawResponseFieldsModel.
type TargetV5RawResponseFieldModel struct {
	Name               types.String `tfsdk:"name"`
	PreserveDuplicates types.Bool   `tfsdk:"preserve_duplicates"`
}

type TargetV5AlgorithmModel struct {
	Name types.String `tfsdk:"name"`
}

type TargetV5RatelimitModel struct {
	Characteristics         []types.String `tfsdk:"characteristics"`
	Period                  types.Int64    `tfsdk:"period"`
	CountingExpression      types.String   `tfsdk:"counting_expression"`
	MitigationTimeout       types.Int64    `tfsdk:"mitigation_timeout"`
	RequestsPerPeriod       types.Int64    `tfsdk:"requests_per_period"`
	RequestsToOrigin        types.Bool     `tfsdk:"requests_to_origin"`
	ScorePerPeriod          types.Int64    `tfsdk:"score_per_period"`
	ScoreResponseHeaderName types.String   `tfsdk:"score_response_header_name"`
}

type TargetV5ExposedCredentialCheckModel struct {
	PasswordExpression types.String `tfsdk:"password_expression"`
	UsernameExpression types.String `tfsdk:"username_expression"`
}

type TargetV5LoggingModel struct {
	Enabled types.Bool `tfsdk:"enabled"`
}

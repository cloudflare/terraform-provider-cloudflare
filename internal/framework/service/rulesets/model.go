package rulesets

import "github.com/hashicorp/terraform-plugin-framework/types"

type RulesetResourceModel struct {
	AccountID   types.String  `tfsdk:"account_id"`
	Description types.String  `tfsdk:"description"`
	ID          types.String  `tfsdk:"id"`
	Kind        types.String  `tfsdk:"kind"`
	Name        types.String  `tfsdk:"name"`
	Phase       types.String  `tfsdk:"phase"`
	Rules       []*RulesModel `tfsdk:"rules"`
	ZoneID      types.String  `tfsdk:"zone_id"`
}

type RulesModel struct {
	Version                types.String                   `tfsdk:"version"`
	Action                 types.String                   `tfsdk:"action"`
	ActionParameters       []*ActionParametersModel       `tfsdk:"action_parameters"`
	Description            types.String                   `tfsdk:"description"`
	Enabled                types.Bool                     `tfsdk:"enabled"`
	ExposedCredentialCheck []*ExposedCredentialCheckModel `tfsdk:"exposed_credential_check"`
	Expression             types.String                   `tfsdk:"expression"`
	ID                     types.String                   `tfsdk:"id"`
	LastUpdated            types.String                   `tfsdk:"last_updated"`
	Logging                []*LoggingModel                `tfsdk:"logging"`
	Ratelimit              []*RatelimitModel              `tfsdk:"ratelimit"`
	Ref                    types.String                   `tfsdk:"ref"`
}

type ActionParametersModel struct {
	Version                 types.String                                 `tfsdk:"version"`
	AutomaticHTTPSRewrites  types.Bool                                   `tfsdk:"automatic_https_rewrites"`
	AutoMinify              []*ActionParameterAutoMinifyModel            `tfsdk:"autominify"`
	BIC                     types.Bool                                   `tfsdk:"bic"`
	BrowserTTL              []*ActionParameterBrowserTTLModel            `tfsdk:"browser_ttl"`
	Cache                   types.Bool                                   `tfsdk:"cache"`
	CacheKey                []*ActionParameterCacheKeyModel              `tfsdk:"cache_key"`
	Content                 types.String                                 `tfsdk:"content"`
	ContentType             types.String                                 `tfsdk:"content_type"`
	CookieFields            types.Set                                    `tfsdk:"cookie_fields"`
	DisableApps             types.Bool                                   `tfsdk:"disable_apps"`
	DisableRailgun          types.Bool                                   `tfsdk:"disable_railgun"`
	DisableZaraz            types.Bool                                   `tfsdk:"disable_zaraz"`
	EdgeTTL                 []*ActionParameterEdgeTTLModel               `tfsdk:"edge_ttl"`
	EmailObfuscation        types.Bool                                   `tfsdk:"email_obfuscation"`
	FromList                []*ActionParameterFromListModel              `tfsdk:"from_list"`
	FromValue               []*ActionParameterFromValueModel             `tfsdk:"from_value"`
	Headers                 []*ActionParametersHeadersModel              `tfsdk:"headers"`
	HostHeader              types.String                                 `tfsdk:"host_header"`
	HotlinkProtection       types.Bool                                   `tfsdk:"hotlink_protection"`
	ID                      types.String                                 `tfsdk:"id"`
	Increment               types.Int64                                  `tfsdk:"increment"`
	MatchedData             []*ActionParametersMatchedDataModel          `tfsdk:"matched_data"`
	Mirage                  types.Bool                                   `tfsdk:"mirage"`
	OpportunisticEncryption types.Bool                                   `tfsdk:"opportunistic_encryption"`
	Origin                  []*ActionParameterOriginModel                `tfsdk:"origin"`
	OriginCacheControl      types.Bool                                   `tfsdk:"origin_cache_control"`
	OriginErrorPagePassthru types.Bool                                   `tfsdk:"origin_error_page_passthru"`
	Overrides               []*ActionParameterOverridesModel             `tfsdk:"overrides"`
	Phases                  types.Set                                    `tfsdk:"phases"`
	Polish                  types.String                                 `tfsdk:"polish"`
	Products                types.Set                                    `tfsdk:"products"`
	ReadTimeout             types.Int64                                  `tfsdk:"read_timeout"`
	RequestFields           types.Set                                    `tfsdk:"request_fields"`
	RespectStrongEtags      types.Bool                                   `tfsdk:"respect_strong_etags"`
	Response                []*ActionParameterResponseModel              `tfsdk:"response"`
	ResponseFields          types.Set                                    `tfsdk:"response_fields"`
	RocketLoader            types.Bool                                   `tfsdk:"rocket_loader"`
	Rules                   map[string]types.String                      `tfsdk:"rules"`
	Ruleset                 types.String                                 `tfsdk:"ruleset"`
	Rulesets                types.Set                                    `tfsdk:"rulesets"`
	SecurityLevel           types.String                                 `tfsdk:"security_level"`
	ServerSideExcludes      types.Bool                                   `tfsdk:"server_side_excludes"`
	ServeStale              []*ActionParameterServeStaleModel            `tfsdk:"serve_stale"`
	SNI                     []*ActionParameterSNIModel                   `tfsdk:"sni"`
	SSL                     types.String                                 `tfsdk:"ssl"`
	StatusCode              types.Int64                                  `tfsdk:"status_code"`
	SXG                     types.Bool                                   `tfsdk:"sxg"`
	URI                     []*ActionParametersURIModel                  `tfsdk:"uri"`
	Algorithms              []*ActionParametersCompressionAlgorithmModel `tfsdk:"algorithms"`
}

type ActionParameterOverridesModel struct {
	Enabled          types.Bool                                 `tfsdk:"enabled"`
	Action           types.String                               `tfsdk:"action"`
	SensitivityLevel types.String                               `tfsdk:"sensitivity_level"`
	Categories       []*ActionParameterOverridesCategoriesModel `tfsdk:"categories"`
	Rules            []*ActionParameterOverridesRulesModel      `tfsdk:"rules"`
}

type ActionParameterOverridesCategoriesModel struct {
	Category types.String `tfsdk:"category"`
	Action   types.String `tfsdk:"action"`
	Enabled  types.Bool   `tfsdk:"enabled"`
}

type ActionParameterOverridesRulesModel struct {
	ID               types.String `tfsdk:"id"`
	Action           types.String `tfsdk:"action"`
	Enabled          types.Bool   `tfsdk:"enabled"`
	ScoreThreshold   types.Int64  `tfsdk:"score_threshold"`
	SensitivityLevel types.String `tfsdk:"sensitivity_level"`
}

type ActionParameterOriginModel struct {
	Host types.String `tfsdk:"host"`
	Port types.Int64  `tfsdk:"port"`
}

type ActionParameterSNIModel struct {
	Value types.String `tfsdk:"value"`
}

type ActionParametersURIModel struct {
	Path   []*ActionParametersURIPartModel `tfsdk:"path"`
	Query  []*ActionParametersURIPartModel `tfsdk:"query"`
	Origin types.Bool                      `tfsdk:"origin"`
}

type ActionParametersURIPartModel struct {
	Value      types.String `tfsdk:"value"`
	Expression types.String `tfsdk:"expression"`
}

type ActionParametersHeadersModel struct {
	Name       types.String `tfsdk:"name"`
	Value      types.String `tfsdk:"value"`
	Expression types.String `tfsdk:"expression"`
	Operation  types.String `tfsdk:"operation"`
}

type ActionParametersMatchedDataModel struct {
	PublicKey types.String `tfsdk:"public_key"`
}

type ActionParameterResponseModel struct {
	StatusCode  types.Int64  `tfsdk:"status_code"`
	ContentType types.String `tfsdk:"content_type"`
	Content     types.String `tfsdk:"content"`
}

type ActionParameterAutoMinifyModel struct {
	HTML types.Bool `tfsdk:"html"`
	CSS  types.Bool `tfsdk:"css"`
	JS   types.Bool `tfsdk:"js"`
}

type ActionParameterEdgeTTLModel struct {
	Mode          types.String                                `tfsdk:"mode"`
	Default       types.Int64                                 `tfsdk:"default"`
	StatusCodeTTL []*ActionParameterEdgeTTLStatusCodeTTLModel `tfsdk:"status_code_ttl"`
}

type ActionParameterEdgeTTLStatusCodeTTLModel struct {
	StatusCode      types.Int64                                                `tfsdk:"status_code"`
	Value           types.Int64                                                `tfsdk:"value"`
	StatusCodeRange []*ActionParameterEdgeTTLStatusCodeTTLStatusCodeRangeModel `tfsdk:"status_code_range"`
}

type ActionParameterEdgeTTLStatusCodeTTLStatusCodeRangeModel struct {
	From types.Int64 `tfsdk:"from"`
	To   types.Int64 `tfsdk:"to"`
}

type ActionParameterBrowserTTLModel struct {
	Mode    types.String `tfsdk:"mode"`
	Default types.Int64  `tfsdk:"default"`
}

type ActionParameterServeStaleModel struct {
	DisableStaleWhileUpdating types.Bool `tfsdk:"disable_stale_while_updating"`
}

type ActionParameterCacheKeyModel struct {
	CacheByDeviceType       types.Bool                               `tfsdk:"cache_by_device_type"`
	IgnoreQueryStringsOrder types.Bool                               `tfsdk:"ignore_query_strings_order"`
	CacheDeceptionArmor     types.Bool                               `tfsdk:"cache_deception_armor"`
	CustomKey               []*ActionParameterCacheKeyCustomKeyModel `tfsdk:"custom_key"`
}

type ActionParameterCacheKeyCustomKeyModel struct {
	QueryString []*ActionParameterCacheKeyCustomKeyQueryStringModel `tfsdk:"query_string"`
	Header      []*ActionParameterCacheKeyCustomKeyHeaderModel      `tfsdk:"header"`
	Cookie      []*ActionParameterCacheKeyCustomKeyCookieModel      `tfsdk:"cookie"`
	User        []*ActionParameterCacheKeyCustomKeyUserModel        `tfsdk:"user"`
	Host        []*ActionParameterCacheKeyCustomKeyHostModel        `tfsdk:"host"`
}

type ActionParameterCacheKeyCustomKeyQueryStringModel struct {
	Include types.Set `tfsdk:"include"`
	Exclude types.Set `tfsdk:"exclude"`
}

type ActionParameterCacheKeyCustomKeyHeaderModel struct {
	Include       types.Set  `tfsdk:"include"`
	CheckPresence types.Set  `tfsdk:"check_presence"`
	ExcludeOrigin types.Bool `tfsdk:"exclude_origin"`
}

type ActionParameterCacheKeyCustomKeyCookieModel struct {
	Include       types.Set `tfsdk:"include"`
	CheckPresence types.Set `tfsdk:"check_presence"`
}

type ActionParameterCacheKeyCustomKeyUserModel struct {
	DeviceType types.Bool `tfsdk:"device_type"`
	Geo        types.Bool `tfsdk:"geo"`
	Lang       types.Bool `tfsdk:"lang"`
}

type ActionParameterCacheKeyCustomKeyHostModel struct {
	Resolved types.Bool `tfsdk:"resolved"`
}

type ActionParameterFromListModel struct {
	Name types.String `tfsdk:"name"`
	Key  types.String `tfsdk:"key"`
}

type ActionParameterFromValueModel struct {
	StatusCode          types.Int64                               `tfsdk:"status_code"`
	PreserveQueryString types.Bool                                `tfsdk:"preserve_query_string"`
	TargetURL           []*ActionParameterFromValueTargetURLModel `tfsdk:"target_url"`
}

type ActionParameterFromValueTargetURLModel struct {
	Value      types.String `tfsdk:"value"`
	Expression types.String `tfsdk:"expression"`
}

type ActionParametersCompressionAlgorithmModel struct {
	Name string `tfsdk:"name"`
}

type RatelimitModel struct {
	Characteristics         types.Set    `tfsdk:"characteristics"`
	CountingExpression      types.String `tfsdk:"counting_expression"`
	MitigationTimeout       types.Int64  `tfsdk:"mitigation_timeout"`
	Period                  types.Int64  `tfsdk:"period"`
	RequestsPerPeriod       types.Int64  `tfsdk:"requests_per_period"`
	RequestsToOrigin        types.Bool   `tfsdk:"requests_to_origin"`
	ScorePerPeriod          types.Int64  `tfsdk:"score_per_period"`
	ScoreResponseHeaderName types.String `tfsdk:"score_response_header_name"`
}

type ExposedCredentialCheckModel struct {
	PasswordExpression types.String `tfsdk:"password_expression"`
	UsernameExpression types.String `tfsdk:"username_expression"`
}

type LoggingModel struct {
	Enabled types.Bool `tfsdk:"enabled"`
}

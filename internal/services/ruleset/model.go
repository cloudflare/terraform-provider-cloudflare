// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ruleset

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type RulesetResultEnvelope struct {
	Result RulesetModel `json:"result"`
}

type RulesetModel struct {
	ID          types.String          `tfsdk:"id" json:"id,computed"`
	AccountID   types.String          `tfsdk:"account_id" path:"account_id,optional"`
	ZoneID      types.String          `tfsdk:"zone_id" path:"zone_id,optional"`
	Kind        types.String          `tfsdk:"kind" json:"kind,required"`
	Name        types.String          `tfsdk:"name" json:"name,required"`
	Phase       types.String          `tfsdk:"phase" json:"phase,required"`
	Rules       *[]*RulesetRulesModel `tfsdk:"rules" json:"rules,required"`
	Description types.String          `tfsdk:"description" json:"description,computed_optional"`
	LastUpdated timetypes.RFC3339     `tfsdk:"last_updated" json:"last_updated,computed" format:"date-time"`
	Version     types.String          `tfsdk:"version" json:"version,computed"`
}

type RulesetRulesModel struct {
	LastUpdated      timetypes.RFC3339                                           `tfsdk:"last_updated" json:"last_updated,computed" format:"date-time"`
	Version          types.String                                                `tfsdk:"version" json:"version,computed"`
	ID               types.String                                                `tfsdk:"id" json:"id,computed_optional"`
	Action           types.String                                                `tfsdk:"action" json:"action,computed_optional"`
	ActionParameters customfield.NestedObject[RulesetRulesActionParametersModel] `tfsdk:"action_parameters" json:"action_parameters,computed_optional"`
	Categories       customfield.List[types.String]                              `tfsdk:"categories" json:"categories,computed"`
	Description      types.String                                                `tfsdk:"description" json:"description,computed_optional"`
	Enabled          types.Bool                                                  `tfsdk:"enabled" json:"enabled,computed_optional"`
	Expression       types.String                                                `tfsdk:"expression" json:"expression,computed_optional"`
	Logging          customfield.NestedObject[RulesetRulesLoggingModel]          `tfsdk:"logging" json:"logging,computed_optional"`
	Ref              types.String                                                `tfsdk:"ref" json:"ref,computed_optional"`
}

type RulesetRulesActionParametersModel struct {
	Response                 customfield.NestedObject[RulesetRulesActionParametersResponseModel]           `tfsdk:"response" json:"response,computed_optional"`
	Algorithms               customfield.NestedObjectList[RulesetRulesActionParametersAlgorithmsModel]     `tfsdk:"algorithms" json:"algorithms,computed_optional"`
	ID                       types.String                                                                  `tfsdk:"id" json:"id,computed_optional"`
	MatchedData              customfield.NestedObject[RulesetRulesActionParametersMatchedDataModel]        `tfsdk:"matched_data" json:"matched_data,computed_optional"`
	Overrides                customfield.NestedObject[RulesetRulesActionParametersOverridesModel]          `tfsdk:"overrides" json:"overrides,computed_optional"`
	FromList                 customfield.NestedObject[RulesetRulesActionParametersFromListModel]           `tfsdk:"from_list" json:"from_list,computed_optional"`
	FromValue                customfield.NestedObject[RulesetRulesActionParametersFromValueModel]          `tfsdk:"from_value" json:"from_value,computed_optional"`
	Headers                  customfield.NestedObjectMap[RulesetRulesActionParametersHeadersModel]         `tfsdk:"headers" json:"headers,computed_optional"`
	URI                      customfield.NestedObject[RulesetRulesActionParametersURIModel]                `tfsdk:"uri" json:"uri,computed_optional"`
	HostHeader               types.String                                                                  `tfsdk:"host_header" json:"host_header,computed_optional"`
	Origin                   customfield.NestedObject[RulesetRulesActionParametersOriginModel]             `tfsdk:"origin" json:"origin,computed_optional"`
	SNI                      customfield.NestedObject[RulesetRulesActionParametersSNIModel]                `tfsdk:"sni" json:"sni,computed_optional"`
	Increment                types.Int64                                                                   `tfsdk:"increment" json:"increment,computed_optional"`
	Content                  types.String                                                                  `tfsdk:"content" json:"content,computed_optional"`
	ContentType              types.String                                                                  `tfsdk:"content_type" json:"content_type,computed_optional"`
	StatusCode               types.Float64                                                                 `tfsdk:"status_code" json:"status_code,computed_optional"`
	AutomaticHTTPSRewrites   types.Bool                                                                    `tfsdk:"automatic_https_rewrites" json:"automatic_https_rewrites,computed_optional"`
	Autominify               customfield.NestedObject[RulesetRulesActionParametersAutominifyModel]         `tfsdk:"autominify" json:"autominify,computed_optional"`
	Bic                      types.Bool                                                                    `tfsdk:"bic" json:"bic,computed_optional"`
	DisableApps              types.Bool                                                                    `tfsdk:"disable_apps" json:"disable_apps,computed_optional"`
	DisableRUM               types.Bool                                                                    `tfsdk:"disable_rum" json:"disable_rum,computed_optional"`
	DisableZaraz             types.Bool                                                                    `tfsdk:"disable_zaraz" json:"disable_zaraz,computed_optional"`
	EmailObfuscation         types.Bool                                                                    `tfsdk:"email_obfuscation" json:"email_obfuscation,computed_optional"`
	Fonts                    types.Bool                                                                    `tfsdk:"fonts" json:"fonts,computed_optional"`
	HotlinkProtection        types.Bool                                                                    `tfsdk:"hotlink_protection" json:"hotlink_protection,computed_optional"`
	Mirage                   types.Bool                                                                    `tfsdk:"mirage" json:"mirage,computed_optional"`
	OpportunisticEncryption  types.Bool                                                                    `tfsdk:"opportunistic_encryption" json:"opportunistic_encryption,computed_optional"`
	Polish                   types.String                                                                  `tfsdk:"polish" json:"polish,computed_optional"`
	RocketLoader             types.Bool                                                                    `tfsdk:"rocket_loader" json:"rocket_loader,computed_optional"`
	SecurityLevel            types.String                                                                  `tfsdk:"security_level" json:"security_level,computed_optional"`
	ServerSideExcludes       types.Bool                                                                    `tfsdk:"server_side_excludes" json:"server_side_excludes,computed_optional"`
	SSL                      types.String                                                                  `tfsdk:"ssl" json:"ssl,computed_optional"`
	Sxg                      types.Bool                                                                    `tfsdk:"sxg" json:"sxg,computed_optional"`
	Phases                   customfield.List[types.String]                                                `tfsdk:"phases" json:"phases,computed_optional"`
	Products                 customfield.List[types.String]                                                `tfsdk:"products" json:"products,computed_optional"`
	Rules                    customfield.Map[customfield.List[types.String]]                               `tfsdk:"rules" json:"rules,computed_optional"`
	Ruleset                  types.String                                                                  `tfsdk:"ruleset" json:"ruleset,computed_optional"`
	Rulesets                 customfield.List[types.String]                                                `tfsdk:"rulesets" json:"rulesets,computed_optional"`
	AdditionalCacheablePorts customfield.List[types.Int64]                                                 `tfsdk:"additional_cacheable_ports" json:"additional_cacheable_ports,computed_optional"`
	BrowserTTL               customfield.NestedObject[RulesetRulesActionParametersBrowserTTLModel]         `tfsdk:"browser_ttl" json:"browser_ttl,computed_optional"`
	Cache                    types.Bool                                                                    `tfsdk:"cache" json:"cache,computed_optional"`
	CacheKey                 customfield.NestedObject[RulesetRulesActionParametersCacheKeyModel]           `tfsdk:"cache_key" json:"cache_key,computed_optional"`
	CacheReserve             customfield.NestedObject[RulesetRulesActionParametersCacheReserveModel]       `tfsdk:"cache_reserve" json:"cache_reserve,computed_optional"`
	EdgeTTL                  customfield.NestedObject[RulesetRulesActionParametersEdgeTTLModel]            `tfsdk:"edge_ttl" json:"edge_ttl,computed_optional"`
	OriginCacheControl       types.Bool                                                                    `tfsdk:"origin_cache_control" json:"origin_cache_control,computed_optional"`
	OriginErrorPagePassthru  types.Bool                                                                    `tfsdk:"origin_error_page_passthru" json:"origin_error_page_passthru,computed_optional"`
	ReadTimeout              types.Int64                                                                   `tfsdk:"read_timeout" json:"read_timeout,computed_optional"`
	RespectStrongEtags       types.Bool                                                                    `tfsdk:"respect_strong_etags" json:"respect_strong_etags,computed_optional"`
	ServeStale               customfield.NestedObject[RulesetRulesActionParametersServeStaleModel]         `tfsdk:"serve_stale" json:"serve_stale,computed_optional"`
	CookieFields             customfield.NestedObjectList[RulesetRulesActionParametersCookieFieldsModel]   `tfsdk:"cookie_fields" json:"cookie_fields,computed_optional"`
	RequestFields            customfield.NestedObjectList[RulesetRulesActionParametersRequestFieldsModel]  `tfsdk:"request_fields" json:"request_fields,computed_optional"`
	ResponseFields           customfield.NestedObjectList[RulesetRulesActionParametersResponseFieldsModel] `tfsdk:"response_fields" json:"response_fields,computed_optional"`
}

type RulesetRulesActionParametersResponseModel struct {
	Content     types.String `tfsdk:"content" json:"content,required"`
	ContentType types.String `tfsdk:"content_type" json:"content_type,required"`
	StatusCode  types.Int64  `tfsdk:"status_code" json:"status_code,required"`
}

type RulesetRulesActionParametersAlgorithmsModel struct {
	Name types.String `tfsdk:"name" json:"name,computed_optional"`
}

type RulesetRulesActionParametersMatchedDataModel struct {
	PublicKey types.String `tfsdk:"public_key" json:"public_key,required"`
}

type RulesetRulesActionParametersOverridesModel struct {
	Action           types.String                                                                       `tfsdk:"action" json:"action,computed_optional"`
	Categories       customfield.NestedObjectList[RulesetRulesActionParametersOverridesCategoriesModel] `tfsdk:"categories" json:"categories,computed_optional"`
	Enabled          types.Bool                                                                         `tfsdk:"enabled" json:"enabled,computed_optional"`
	Rules            customfield.NestedObjectList[RulesetRulesActionParametersOverridesRulesModel]      `tfsdk:"rules" json:"rules,computed_optional"`
	SensitivityLevel types.String                                                                       `tfsdk:"sensitivity_level" json:"sensitivity_level,computed_optional"`
}

type RulesetRulesActionParametersOverridesCategoriesModel struct {
	Category         types.String `tfsdk:"category" json:"category,required"`
	Action           types.String `tfsdk:"action" json:"action,computed_optional"`
	Enabled          types.Bool   `tfsdk:"enabled" json:"enabled,computed_optional"`
	SensitivityLevel types.String `tfsdk:"sensitivity_level" json:"sensitivity_level,computed_optional"`
}

type RulesetRulesActionParametersOverridesRulesModel struct {
	ID               types.String `tfsdk:"id" json:"id,required"`
	Action           types.String `tfsdk:"action" json:"action,computed_optional"`
	Enabled          types.Bool   `tfsdk:"enabled" json:"enabled,computed_optional"`
	ScoreThreshold   types.Int64  `tfsdk:"score_threshold" json:"score_threshold,computed_optional"`
	SensitivityLevel types.String `tfsdk:"sensitivity_level" json:"sensitivity_level,computed_optional"`
}

type RulesetRulesActionParametersFromListModel struct {
	Key  types.String `tfsdk:"key" json:"key,computed_optional"`
	Name types.String `tfsdk:"name" json:"name,computed_optional"`
}

type RulesetRulesActionParametersFromValueModel struct {
	PreserveQueryString types.Bool                                                                    `tfsdk:"preserve_query_string" json:"preserve_query_string,computed_optional"`
	StatusCode          types.Float64                                                                 `tfsdk:"status_code" json:"status_code,computed_optional"`
	TargetURL           customfield.NestedObject[RulesetRulesActionParametersFromValueTargetURLModel] `tfsdk:"target_url" json:"target_url,computed_optional"`
}

type RulesetRulesActionParametersFromValueTargetURLModel struct {
	Value      types.String `tfsdk:"value" json:"value,computed_optional"`
	Expression types.String `tfsdk:"expression" json:"expression,computed_optional"`
}

type RulesetRulesActionParametersHeadersModel struct {
	Operation  types.String `tfsdk:"operation" json:"operation,required"`
	Value      types.String `tfsdk:"value" json:"value,computed_optional"`
	Expression types.String `tfsdk:"expression" json:"expression,computed_optional"`
}

type RulesetRulesActionParametersURIModel struct {
	Path  customfield.NestedObject[RulesetRulesActionParametersURIPathModel]  `tfsdk:"path" json:"path,computed_optional"`
	Query customfield.NestedObject[RulesetRulesActionParametersURIQueryModel] `tfsdk:"query" json:"query,computed_optional"`
}

type RulesetRulesActionParametersURIPathModel struct {
	Value      types.String `tfsdk:"value" json:"value,computed_optional"`
	Expression types.String `tfsdk:"expression" json:"expression,computed_optional"`
}

type RulesetRulesActionParametersURIQueryModel struct {
	Value      types.String `tfsdk:"value" json:"value,computed_optional"`
	Expression types.String `tfsdk:"expression" json:"expression,computed_optional"`
}

type RulesetRulesActionParametersOriginModel struct {
	Host types.String  `tfsdk:"host" json:"host,computed_optional"`
	Port types.Float64 `tfsdk:"port" json:"port,computed_optional"`
}

type RulesetRulesActionParametersSNIModel struct {
	Value types.String `tfsdk:"value" json:"value,required"`
}

type RulesetRulesActionParametersAutominifyModel struct {
	Css  types.Bool `tfsdk:"css" json:"css,computed_optional"`
	HTML types.Bool `tfsdk:"html" json:"html,computed_optional"`
	JS   types.Bool `tfsdk:"js" json:"js,computed_optional"`
}

type RulesetRulesActionParametersBrowserTTLModel struct {
	Mode    types.String `tfsdk:"mode" json:"mode,required"`
	Default types.Int64  `tfsdk:"default" json:"default,computed_optional"`
}

type RulesetRulesActionParametersCacheKeyModel struct {
	CacheByDeviceType       types.Bool                                                                   `tfsdk:"cache_by_device_type" json:"cache_by_device_type,computed_optional"`
	CacheDeceptionArmor     types.Bool                                                                   `tfsdk:"cache_deception_armor" json:"cache_deception_armor,computed_optional"`
	CustomKey               customfield.NestedObject[RulesetRulesActionParametersCacheKeyCustomKeyModel] `tfsdk:"custom_key" json:"custom_key,computed_optional"`
	IgnoreQueryStringsOrder types.Bool                                                                   `tfsdk:"ignore_query_strings_order" json:"ignore_query_strings_order,computed_optional"`
}

type RulesetRulesActionParametersCacheKeyCustomKeyModel struct {
	Cookie      customfield.NestedObject[RulesetRulesActionParametersCacheKeyCustomKeyCookieModel]      `tfsdk:"cookie" json:"cookie,computed_optional"`
	Header      customfield.NestedObject[RulesetRulesActionParametersCacheKeyCustomKeyHeaderModel]      `tfsdk:"header" json:"header,computed_optional"`
	Host        customfield.NestedObject[RulesetRulesActionParametersCacheKeyCustomKeyHostModel]        `tfsdk:"host" json:"host,computed_optional"`
	QueryString customfield.NestedObject[RulesetRulesActionParametersCacheKeyCustomKeyQueryStringModel] `tfsdk:"query_string" json:"query_string,computed_optional"`
	User        customfield.NestedObject[RulesetRulesActionParametersCacheKeyCustomKeyUserModel]        `tfsdk:"user" json:"user,computed_optional"`
}

type RulesetRulesActionParametersCacheKeyCustomKeyCookieModel struct {
	CheckPresence customfield.List[types.String] `tfsdk:"check_presence" json:"check_presence,computed_optional"`
	Include       customfield.List[types.String] `tfsdk:"include" json:"include,computed_optional"`
}

type RulesetRulesActionParametersCacheKeyCustomKeyHeaderModel struct {
	CheckPresence customfield.List[types.String]                  `tfsdk:"check_presence" json:"check_presence,computed_optional"`
	Contains      customfield.Map[customfield.List[types.String]] `tfsdk:"contains" json:"contains,computed_optional"`
	ExcludeOrigin types.Bool                                      `tfsdk:"exclude_origin" json:"exclude_origin,computed_optional"`
	Include       customfield.List[types.String]                  `tfsdk:"include" json:"include,computed_optional"`
}

type RulesetRulesActionParametersCacheKeyCustomKeyHostModel struct {
	Resolved types.Bool `tfsdk:"resolved" json:"resolved,computed_optional"`
}

type RulesetRulesActionParametersCacheKeyCustomKeyQueryStringModel struct {
	Exclude customfield.NestedObject[RulesetRulesActionParametersCacheKeyCustomKeyQueryStringExcludeModel] `tfsdk:"exclude" json:"exclude,computed_optional"`
	Include customfield.NestedObject[RulesetRulesActionParametersCacheKeyCustomKeyQueryStringIncludeModel] `tfsdk:"include" json:"include,computed_optional"`
}

type RulesetRulesActionParametersCacheKeyCustomKeyQueryStringExcludeModel struct {
	All  types.Bool                     `tfsdk:"all" json:"all,computed_optional"`
	List customfield.List[types.String] `tfsdk:"list" json:"list,computed_optional"`
}

type RulesetRulesActionParametersCacheKeyCustomKeyQueryStringIncludeModel struct {
	All  types.Bool                     `tfsdk:"all" json:"all,computed_optional"`
	List customfield.List[types.String] `tfsdk:"list" json:"list,computed_optional"`
}

type RulesetRulesActionParametersCacheKeyCustomKeyUserModel struct {
	DeviceType types.Bool `tfsdk:"device_type" json:"device_type,computed_optional"`
	Geo        types.Bool `tfsdk:"geo" json:"geo,computed_optional"`
	Lang       types.Bool `tfsdk:"lang" json:"lang,computed_optional"`
}

type RulesetRulesActionParametersCacheReserveModel struct {
	Eligible        types.Bool  `tfsdk:"eligible" json:"eligible,required"`
	MinimumFileSize types.Int64 `tfsdk:"minimum_file_size" json:"minimum_file_size,required"`
}

type RulesetRulesActionParametersEdgeTTLModel struct {
	Default       types.Int64                                               `tfsdk:"default" json:"default,required"`
	Mode          types.String                                              `tfsdk:"mode" json:"mode,required"`
	StatusCodeTTL *[]*RulesetRulesActionParametersEdgeTTLStatusCodeTTLModel `tfsdk:"status_code_ttl" json:"status_code_ttl,required"`
}

type RulesetRulesActionParametersEdgeTTLStatusCodeTTLModel struct {
	Value           types.Int64                                                                                    `tfsdk:"value" json:"value,required"`
	StatusCodeRange customfield.NestedObject[RulesetRulesActionParametersEdgeTTLStatusCodeTTLStatusCodeRangeModel] `tfsdk:"status_code_range" json:"status_code_range,computed_optional"`
	StatusCodeValue types.Int64                                                                                    `tfsdk:"status_code_value" json:"status_code_value,computed_optional"`
}

type RulesetRulesActionParametersEdgeTTLStatusCodeTTLStatusCodeRangeModel struct {
	From types.Int64 `tfsdk:"from" json:"from,required"`
	To   types.Int64 `tfsdk:"to" json:"to,required"`
}

type RulesetRulesActionParametersServeStaleModel struct {
	DisableStaleWhileUpdating types.Bool `tfsdk:"disable_stale_while_updating" json:"disable_stale_while_updating,required"`
}

type RulesetRulesActionParametersCookieFieldsModel struct {
	Name types.String `tfsdk:"name" json:"name,required"`
}

type RulesetRulesActionParametersRequestFieldsModel struct {
	Name types.String `tfsdk:"name" json:"name,required"`
}

type RulesetRulesActionParametersResponseFieldsModel struct {
	Name types.String `tfsdk:"name" json:"name,required"`
}

type RulesetRulesLoggingModel struct {
	Enabled types.Bool `tfsdk:"enabled" json:"enabled,required"`
}

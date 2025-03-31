// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ruleset

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type RulesetResultEnvelope struct {
	Result RulesetModel `json:"result"`
}

type RulesetModel struct {
	ID          types.String                                    `tfsdk:"id" json:"id,computed"`
	AccountID   types.String                                    `tfsdk:"account_id" path:"account_id,optional"`
	ZoneID      types.String                                    `tfsdk:"zone_id" path:"zone_id,optional"`
	Kind        types.String                                    `tfsdk:"kind" json:"kind,required"`
	Name        types.String                                    `tfsdk:"name" json:"name,required"`
	Phase       types.String                                    `tfsdk:"phase" json:"phase,required"`
	Description types.String                                    `tfsdk:"description" json:"description,computed_optional"`
	Rules       customfield.NestedObjectList[RulesetRulesModel] `tfsdk:"rules" json:"rules,computed_optional"`
}

func (m RulesetModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m RulesetModel) MarshalJSONForUpdate(state RulesetModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type RulesetRulesModel struct {
	ID                     types.String                                                      `tfsdk:"id" json:"id,optional"`
	Action                 types.String                                                      `tfsdk:"action" json:"action,optional"`
	ActionParameters       customfield.NestedObject[RulesetRulesActionParametersModel]       `tfsdk:"action_parameters" json:"action_parameters,computed_optional"`
	Categories             customfield.List[types.String]                                    `tfsdk:"categories" json:"categories,computed"`
	Description            types.String                                                      `tfsdk:"description" json:"description,computed_optional"`
	Enabled                types.Bool                                                        `tfsdk:"enabled" json:"enabled,computed_optional"`
	ExposedCredentialCheck customfield.NestedObject[RulesetRulesExposedCredentialCheckModel] `tfsdk:"exposed_credential_check" json:"exposed_credential_check,computed_optional"`
	Expression             types.String                                                      `tfsdk:"expression" json:"expression,optional"`
	Logging                jsontypes.Normalized                                              `tfsdk:"logging" json:"logging,optional"`
	Ratelimit              customfield.NestedObject[RulesetRulesRatelimitModel]              `tfsdk:"ratelimit" json:"ratelimit,computed_optional"`
	Ref                    types.String                                                      `tfsdk:"ref" json:"ref,optional"`
}

type RulesetRulesActionParametersModel struct {
	Response                 customfield.NestedObject[RulesetRulesActionParametersResponseModel]                     `tfsdk:"response" json:"response,computed_optional"`
	Algorithms               customfield.NestedObjectList[RulesetRulesActionParametersAlgorithmsModel]               `tfsdk:"algorithms" json:"algorithms,computed_optional"`
	ID                       types.String                                                                            `tfsdk:"id" json:"id,optional"`
	MatchedData              customfield.NestedObject[RulesetRulesActionParametersMatchedDataModel]                  `tfsdk:"matched_data" json:"matched_data,computed_optional"`
	Overrides                customfield.NestedObject[RulesetRulesActionParametersOverridesModel]                    `tfsdk:"overrides" json:"overrides,computed_optional"`
	FromList                 customfield.NestedObject[RulesetRulesActionParametersFromListModel]                     `tfsdk:"from_list" json:"from_list,computed_optional"`
	FromValue                customfield.NestedObject[RulesetRulesActionParametersFromValueModel]                    `tfsdk:"from_value" json:"from_value,computed_optional"`
	Headers                  customfield.NestedObjectMap[RulesetRulesActionParametersHeadersModel]                   `tfsdk:"headers" json:"headers,computed_optional"`
	URI                      customfield.NestedObject[RulesetRulesActionParametersURIModel]                          `tfsdk:"uri" json:"uri,computed_optional"`
	HostHeader               types.String                                                                            `tfsdk:"host_header" json:"host_header,optional"`
	Origin                   customfield.NestedObject[RulesetRulesActionParametersOriginModel]                       `tfsdk:"origin" json:"origin,computed_optional"`
	SNI                      customfield.NestedObject[RulesetRulesActionParametersSNIModel]                          `tfsdk:"sni" json:"sni,computed_optional"`
	Increment                types.Int64                                                                             `tfsdk:"increment" json:"increment,optional"`
	Content                  types.String                                                                            `tfsdk:"content" json:"content,optional"`
	ContentType              types.String                                                                            `tfsdk:"content_type" json:"content_type,optional"`
	StatusCode               types.Float64                                                                           `tfsdk:"status_code" json:"status_code,optional"`
	AutomaticHTTPSRewrites   types.Bool                                                                              `tfsdk:"automatic_https_rewrites" json:"automatic_https_rewrites,optional"`
	Autominify               customfield.NestedObject[RulesetRulesActionParametersAutominifyModel]                   `tfsdk:"autominify" json:"autominify,computed_optional"`
	BIC                      types.Bool                                                                              `tfsdk:"bic" json:"bic,optional"`
	DisableApps              types.Bool                                                                              `tfsdk:"disable_apps" json:"disable_apps,optional"`
	DisableRUM               types.Bool                                                                              `tfsdk:"disable_rum" json:"disable_rum,optional"`
	DisableZaraz             types.Bool                                                                              `tfsdk:"disable_zaraz" json:"disable_zaraz,optional"`
	EmailObfuscation         types.Bool                                                                              `tfsdk:"email_obfuscation" json:"email_obfuscation,optional"`
	Fonts                    types.Bool                                                                              `tfsdk:"fonts" json:"fonts,optional"`
	HotlinkProtection        types.Bool                                                                              `tfsdk:"hotlink_protection" json:"hotlink_protection,optional"`
	Mirage                   types.Bool                                                                              `tfsdk:"mirage" json:"mirage,optional"`
	OpportunisticEncryption  types.Bool                                                                              `tfsdk:"opportunistic_encryption" json:"opportunistic_encryption,optional"`
	Polish                   types.String                                                                            `tfsdk:"polish" json:"polish,optional"`
	RocketLoader             types.Bool                                                                              `tfsdk:"rocket_loader" json:"rocket_loader,optional"`
	SecurityLevel            types.String                                                                            `tfsdk:"security_level" json:"security_level,optional"`
	ServerSideExcludes       types.Bool                                                                              `tfsdk:"server_side_excludes" json:"server_side_excludes,optional"`
	SSL                      types.String                                                                            `tfsdk:"ssl" json:"ssl,optional"`
	SXG                      types.Bool                                                                              `tfsdk:"sxg" json:"sxg,optional"`
	Phases                   *[]types.String                                                                         `tfsdk:"phases" json:"phases,optional"`
	Products                 *[]types.String                                                                         `tfsdk:"products" json:"products,optional"`
	Rules                    *map[string]*[]types.String                                                             `tfsdk:"rules" json:"rules,optional"`
	Ruleset                  types.String                                                                            `tfsdk:"ruleset" json:"ruleset,optional"`
	Rulesets                 *[]types.String                                                                         `tfsdk:"rulesets" json:"rulesets,optional"`
	AdditionalCacheablePorts *[]types.Int64                                                                          `tfsdk:"additional_cacheable_ports" json:"additional_cacheable_ports,optional"`
	BrowserTTL               customfield.NestedObject[RulesetRulesActionParametersBrowserTTLModel]                   `tfsdk:"browser_ttl" json:"browser_ttl,computed_optional"`
	Cache                    types.Bool                                                                              `tfsdk:"cache" json:"cache,optional"`
	CacheKey                 customfield.NestedObject[RulesetRulesActionParametersCacheKeyModel]                     `tfsdk:"cache_key" json:"cache_key,computed_optional"`
	CacheReserve             customfield.NestedObject[RulesetRulesActionParametersCacheReserveModel]                 `tfsdk:"cache_reserve" json:"cache_reserve,computed_optional"`
	EdgeTTL                  customfield.NestedObject[RulesetRulesActionParametersEdgeTTLModel]                      `tfsdk:"edge_ttl" json:"edge_ttl,computed_optional"`
	OriginCacheControl       types.Bool                                                                              `tfsdk:"origin_cache_control" json:"origin_cache_control,optional"`
	OriginErrorPagePassthru  types.Bool                                                                              `tfsdk:"origin_error_page_passthru" json:"origin_error_page_passthru,optional"`
	ReadTimeout              types.Int64                                                                             `tfsdk:"read_timeout" json:"read_timeout,optional"`
	RespectStrongEtags       types.Bool                                                                              `tfsdk:"respect_strong_etags" json:"respect_strong_etags,optional"`
	ServeStale               customfield.NestedObject[RulesetRulesActionParametersServeStaleModel]                   `tfsdk:"serve_stale" json:"serve_stale,computed_optional"`
	CookieFields             customfield.NestedObjectList[RulesetRulesActionParametersCookieFieldsModel]             `tfsdk:"cookie_fields" json:"cookie_fields,computed_optional"`
	RawResponseFields        customfield.NestedObjectList[RulesetRulesActionParametersRawResponseFieldsModel]        `tfsdk:"raw_response_fields" json:"raw_response_fields,computed_optional"`
	RequestFields            customfield.NestedObjectList[RulesetRulesActionParametersRequestFieldsModel]            `tfsdk:"request_fields" json:"request_fields,computed_optional"`
	ResponseFields           customfield.NestedObjectList[RulesetRulesActionParametersResponseFieldsModel]           `tfsdk:"response_fields" json:"response_fields,computed_optional"`
	TransformedRequestFields customfield.NestedObjectList[RulesetRulesActionParametersTransformedRequestFieldsModel] `tfsdk:"transformed_request_fields" json:"transformed_request_fields,computed_optional"`
}

type RulesetRulesActionParametersResponseModel struct {
	Content     types.String `tfsdk:"content" json:"content,required"`
	ContentType types.String `tfsdk:"content_type" json:"content_type,required"`
	StatusCode  types.Int64  `tfsdk:"status_code" json:"status_code,required"`
}

type RulesetRulesActionParametersAlgorithmsModel struct {
	Name types.String `tfsdk:"name" json:"name,optional"`
}

type RulesetRulesActionParametersMatchedDataModel struct {
	PublicKey types.String `tfsdk:"public_key" json:"public_key,required"`
}

type RulesetRulesActionParametersOverridesModel struct {
	Action           types.String                                                                       `tfsdk:"action" json:"action,optional"`
	Categories       customfield.NestedObjectList[RulesetRulesActionParametersOverridesCategoriesModel] `tfsdk:"categories" json:"categories,computed_optional"`
	Enabled          types.Bool                                                                         `tfsdk:"enabled" json:"enabled,optional"`
	Rules            customfield.NestedObjectList[RulesetRulesActionParametersOverridesRulesModel]      `tfsdk:"rules" json:"rules,computed_optional"`
	SensitivityLevel types.String                                                                       `tfsdk:"sensitivity_level" json:"sensitivity_level,optional"`
}

type RulesetRulesActionParametersOverridesCategoriesModel struct {
	Category         types.String `tfsdk:"category" json:"category,required"`
	Action           types.String `tfsdk:"action" json:"action,optional"`
	Enabled          types.Bool   `tfsdk:"enabled" json:"enabled,optional"`
	SensitivityLevel types.String `tfsdk:"sensitivity_level" json:"sensitivity_level,optional"`
}

type RulesetRulesActionParametersOverridesRulesModel struct {
	ID               types.String `tfsdk:"id" json:"id,required"`
	Action           types.String `tfsdk:"action" json:"action,optional"`
	Enabled          types.Bool   `tfsdk:"enabled" json:"enabled,optional"`
	ScoreThreshold   types.Int64  `tfsdk:"score_threshold" json:"score_threshold,optional"`
	SensitivityLevel types.String `tfsdk:"sensitivity_level" json:"sensitivity_level,optional"`
}

type RulesetRulesActionParametersFromListModel struct {
	Key  types.String `tfsdk:"key" json:"key,optional"`
	Name types.String `tfsdk:"name" json:"name,optional"`
}

type RulesetRulesActionParametersFromValueModel struct {
	PreserveQueryString types.Bool                                                                    `tfsdk:"preserve_query_string" json:"preserve_query_string,optional"`
	StatusCode          types.Float64                                                                 `tfsdk:"status_code" json:"status_code,optional"`
	TargetURL           customfield.NestedObject[RulesetRulesActionParametersFromValueTargetURLModel] `tfsdk:"target_url" json:"target_url,computed_optional"`
}

type RulesetRulesActionParametersFromValueTargetURLModel struct {
	Value      types.String `tfsdk:"value" json:"value,optional"`
	Expression types.String `tfsdk:"expression" json:"expression,optional"`
}

type RulesetRulesActionParametersHeadersModel struct {
	Operation  types.String `tfsdk:"operation" json:"operation,required"`
	Value      types.String `tfsdk:"value" json:"value,optional"`
	Expression types.String `tfsdk:"expression" json:"expression,optional"`
}

type RulesetRulesActionParametersURIModel struct {
	Path  customfield.NestedObject[RulesetRulesActionParametersURIPathModel]  `tfsdk:"path" json:"path,computed_optional"`
	Query customfield.NestedObject[RulesetRulesActionParametersURIQueryModel] `tfsdk:"query" json:"query,computed_optional"`
}

type RulesetRulesActionParametersURIPathModel struct {
	Value      types.String `tfsdk:"value" json:"value,optional"`
	Expression types.String `tfsdk:"expression" json:"expression,optional"`
}

type RulesetRulesActionParametersURIQueryModel struct {
	Value      types.String `tfsdk:"value" json:"value,optional"`
	Expression types.String `tfsdk:"expression" json:"expression,optional"`
}

type RulesetRulesActionParametersOriginModel struct {
	Host types.String  `tfsdk:"host" json:"host,optional"`
	Port types.Float64 `tfsdk:"port" json:"port,optional"`
}

type RulesetRulesActionParametersSNIModel struct {
	Value types.String `tfsdk:"value" json:"value,required"`
}

type RulesetRulesActionParametersAutominifyModel struct {
	CSS  types.Bool `tfsdk:"css" json:"css,optional"`
	HTML types.Bool `tfsdk:"html" json:"html,optional"`
	JS   types.Bool `tfsdk:"js" json:"js,optional"`
}

type RulesetRulesActionParametersBrowserTTLModel struct {
	Mode    types.String `tfsdk:"mode" json:"mode,required"`
	Default types.Int64  `tfsdk:"default" json:"default,optional"`
}

type RulesetRulesActionParametersCacheKeyModel struct {
	CacheByDeviceType       types.Bool                                                                   `tfsdk:"cache_by_device_type" json:"cache_by_device_type,optional"`
	CacheDeceptionArmor     types.Bool                                                                   `tfsdk:"cache_deception_armor" json:"cache_deception_armor,optional"`
	CustomKey               customfield.NestedObject[RulesetRulesActionParametersCacheKeyCustomKeyModel] `tfsdk:"custom_key" json:"custom_key,computed_optional"`
	IgnoreQueryStringsOrder types.Bool                                                                   `tfsdk:"ignore_query_strings_order" json:"ignore_query_strings_order,optional"`
}

type RulesetRulesActionParametersCacheKeyCustomKeyModel struct {
	Cookie      customfield.NestedObject[RulesetRulesActionParametersCacheKeyCustomKeyCookieModel]      `tfsdk:"cookie" json:"cookie,computed_optional"`
	Header      customfield.NestedObject[RulesetRulesActionParametersCacheKeyCustomKeyHeaderModel]      `tfsdk:"header" json:"header,computed_optional"`
	Host        customfield.NestedObject[RulesetRulesActionParametersCacheKeyCustomKeyHostModel]        `tfsdk:"host" json:"host,computed_optional"`
	QueryString customfield.NestedObject[RulesetRulesActionParametersCacheKeyCustomKeyQueryStringModel] `tfsdk:"query_string" json:"query_string,computed_optional"`
	User        customfield.NestedObject[RulesetRulesActionParametersCacheKeyCustomKeyUserModel]        `tfsdk:"user" json:"user,computed_optional"`
}

type RulesetRulesActionParametersCacheKeyCustomKeyCookieModel struct {
	CheckPresence *[]types.String `tfsdk:"check_presence" json:"check_presence,optional"`
	Include       *[]types.String `tfsdk:"include" json:"include,optional"`
}

type RulesetRulesActionParametersCacheKeyCustomKeyHeaderModel struct {
	CheckPresence *[]types.String             `tfsdk:"check_presence" json:"check_presence,optional"`
	Contains      *map[string]*[]types.String `tfsdk:"contains" json:"contains,optional"`
	ExcludeOrigin types.Bool                  `tfsdk:"exclude_origin" json:"exclude_origin,optional"`
	Include       *[]types.String             `tfsdk:"include" json:"include,optional"`
}

type RulesetRulesActionParametersCacheKeyCustomKeyHostModel struct {
	Resolved types.Bool `tfsdk:"resolved" json:"resolved,optional"`
}

type RulesetRulesActionParametersCacheKeyCustomKeyQueryStringModel struct {
	Include customfield.NestedObject[RulesetRulesActionParametersCacheKeyCustomKeyQueryStringIncludeModel] `tfsdk:"include" json:"include,computed_optional"`
	Exclude customfield.NestedObject[RulesetRulesActionParametersCacheKeyCustomKeyQueryStringExcludeModel] `tfsdk:"exclude" json:"exclude,computed_optional"`
}

type RulesetRulesActionParametersCacheKeyCustomKeyQueryStringIncludeModel struct {
	List *[]types.String `tfsdk:"list" json:"list,optional"`
	All  types.Bool      `tfsdk:"all" json:"all,optional"`
}

type RulesetRulesActionParametersCacheKeyCustomKeyQueryStringExcludeModel struct {
	List *[]types.String `tfsdk:"list" json:"list,optional"`
	All  types.Bool      `tfsdk:"all" json:"all,optional"`
}

type RulesetRulesActionParametersCacheKeyCustomKeyUserModel struct {
	DeviceType types.Bool `tfsdk:"device_type" json:"device_type,optional"`
	Geo        types.Bool `tfsdk:"geo" json:"geo,optional"`
	Lang       types.Bool `tfsdk:"lang" json:"lang,optional"`
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
	Value           types.Int64                                                           `tfsdk:"value" json:"value,required"`
	StatusCodeRange *RulesetRulesActionParametersEdgeTTLStatusCodeTTLStatusCodeRangeModel `tfsdk:"status_code_range" json:"status_code_range,optional"`
	StatusCodeValue types.Int64                                                           `tfsdk:"status_code_value" json:"status_code_value,optional"`
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

type RulesetRulesActionParametersRawResponseFieldsModel struct {
	Name               types.String `tfsdk:"name" json:"name,required"`
	PreserveDuplicates types.Bool   `tfsdk:"preserve_duplicates" json:"preserve_duplicates,computed_optional"`
}

type RulesetRulesActionParametersRequestFieldsModel struct {
	Name types.String `tfsdk:"name" json:"name,required"`
}

type RulesetRulesActionParametersResponseFieldsModel struct {
	Name               types.String `tfsdk:"name" json:"name,required"`
	PreserveDuplicates types.Bool   `tfsdk:"preserve_duplicates" json:"preserve_duplicates,computed_optional"`
}

type RulesetRulesActionParametersTransformedRequestFieldsModel struct {
	Name types.String `tfsdk:"name" json:"name,required"`
}

type RulesetRulesExposedCredentialCheckModel struct {
	PasswordExpression types.String `tfsdk:"password_expression" json:"password_expression,required"`
	UsernameExpression types.String `tfsdk:"username_expression" json:"username_expression,required"`
}

type RulesetRulesRatelimitModel struct {
	Characteristics         *[]types.String `tfsdk:"characteristics" json:"characteristics,required"`
	Period                  types.Int64     `tfsdk:"period" json:"period,required"`
	CountingExpression      types.String    `tfsdk:"counting_expression" json:"counting_expression,optional"`
	MitigationTimeout       types.Int64     `tfsdk:"mitigation_timeout" json:"mitigation_timeout,optional"`
	RequestsPerPeriod       types.Int64     `tfsdk:"requests_per_period" json:"requests_per_period,optional"`
	RequestsToOrigin        types.Bool      `tfsdk:"requests_to_origin" json:"requests_to_origin,optional"`
	ScorePerPeriod          types.Int64     `tfsdk:"score_per_period" json:"score_per_period,optional"`
	ScoreResponseHeaderName types.String    `tfsdk:"score_response_header_name" json:"score_response_header_name,optional"`
}

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ruleset

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/rulesets"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type RulesetResultDataSourceEnvelope struct {
	Result RulesetDataSourceModel `json:"result,computed"`
}

type RulesetResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[RulesetDataSourceModel] `json:"result,computed"`
}

type RulesetDataSourceModel struct {
	AccountID   types.String                                              `tfsdk:"account_id" path:"account_id,optional"`
	RulesetID   types.String                                              `tfsdk:"ruleset_id" path:"ruleset_id,optional"`
	ZoneID      types.String                                              `tfsdk:"zone_id" path:"zone_id,optional"`
	Description types.String                                              `tfsdk:"description" json:"description,computed"`
	ID          types.String                                              `tfsdk:"id" json:"id,computed"`
	Kind        types.String                                              `tfsdk:"kind" json:"kind,computed"`
	LastUpdated timetypes.RFC3339                                         `tfsdk:"last_updated" json:"last_updated,computed" format:"date-time"`
	Name        types.String                                              `tfsdk:"name" json:"name,computed"`
	Phase       types.String                                              `tfsdk:"phase" json:"phase,computed"`
	Version     types.String                                              `tfsdk:"version" json:"version,computed"`
	Rules       customfield.NestedObjectList[RulesetRulesDataSourceModel] `tfsdk:"rules" json:"rules,computed"`
	Filter      *RulesetFindOneByDataSourceModel                          `tfsdk:"filter"`
}

func (m *RulesetDataSourceModel) toReadParams(_ context.Context) (params rulesets.RulesetGetParams, diags diag.Diagnostics) {
	params = rulesets.RulesetGetParams{}

	if !m.Filter.AccountID.IsNull() {
		params.AccountID = cloudflare.F(m.Filter.AccountID.ValueString())
	} else {
		params.ZoneID = cloudflare.F(m.Filter.ZoneID.ValueString())
	}

	return
}

func (m *RulesetDataSourceModel) toListParams(_ context.Context) (params rulesets.RulesetListParams, diags diag.Diagnostics) {
	params = rulesets.RulesetListParams{}

	if !m.Filter.AccountID.IsNull() {
		params.AccountID = cloudflare.F(m.Filter.AccountID.ValueString())
	} else {
		params.ZoneID = cloudflare.F(m.Filter.ZoneID.ValueString())
	}

	return
}

type RulesetRulesDataSourceModel struct {
	LastUpdated            timetypes.RFC3339                                                           `tfsdk:"last_updated" json:"last_updated,computed" format:"date-time"`
	Version                types.String                                                                `tfsdk:"version" json:"version,computed"`
	ID                     types.String                                                                `tfsdk:"id" json:"id,computed"`
	Action                 types.String                                                                `tfsdk:"action" json:"action,computed"`
	ActionParameters       customfield.NestedObject[RulesetRulesActionParametersDataSourceModel]       `tfsdk:"action_parameters" json:"action_parameters,computed"`
	Categories             customfield.List[types.String]                                              `tfsdk:"categories" json:"categories,computed"`
	Description            types.String                                                                `tfsdk:"description" json:"description,computed"`
	Enabled                types.Bool                                                                  `tfsdk:"enabled" json:"enabled,computed"`
	ExposedCredentialCheck customfield.NestedObject[RulesetRulesExposedCredentialCheckDataSourceModel] `tfsdk:"exposed_credential_check" json:"exposed_credential_check,computed"`
	Expression             types.String                                                                `tfsdk:"expression" json:"expression,computed"`
	Logging                customfield.NestedObject[RulesetRulesLoggingDataSourceModel]                `tfsdk:"logging" json:"logging,computed"`
	Ratelimit              customfield.NestedObject[RulesetRulesRatelimitDataSourceModel]              `tfsdk:"ratelimit" json:"ratelimit,computed"`
	Ref                    types.String                                                                `tfsdk:"ref" json:"ref,computed"`
}

type RulesetRulesActionParametersDataSourceModel struct {
	Response                 customfield.NestedObject[RulesetRulesActionParametersResponseDataSourceModel]           `tfsdk:"response" json:"response,computed"`
	Algorithms               customfield.NestedObjectList[RulesetRulesActionParametersAlgorithmsDataSourceModel]     `tfsdk:"algorithms" json:"algorithms,computed"`
	ID                       types.String                                                                            `tfsdk:"id" json:"id,computed"`
	MatchedData              customfield.NestedObject[RulesetRulesActionParametersMatchedDataDataSourceModel]        `tfsdk:"matched_data" json:"matched_data,computed"`
	Overrides                customfield.NestedObject[RulesetRulesActionParametersOverridesDataSourceModel]          `tfsdk:"overrides" json:"overrides,computed"`
	FromList                 customfield.NestedObject[RulesetRulesActionParametersFromListDataSourceModel]           `tfsdk:"from_list" json:"from_list,computed"`
	FromValue                customfield.NestedObject[RulesetRulesActionParametersFromValueDataSourceModel]          `tfsdk:"from_value" json:"from_value,computed"`
	Headers                  customfield.NestedObjectMap[RulesetRulesActionParametersHeadersDataSourceModel]         `tfsdk:"headers" json:"headers,computed"`
	URI                      customfield.NestedObject[RulesetRulesActionParametersURIDataSourceModel]                `tfsdk:"uri" json:"uri,computed"`
	HostHeader               types.String                                                                            `tfsdk:"host_header" json:"host_header,computed"`
	Origin                   customfield.NestedObject[RulesetRulesActionParametersOriginDataSourceModel]             `tfsdk:"origin" json:"origin,computed"`
	SNI                      customfield.NestedObject[RulesetRulesActionParametersSNIDataSourceModel]                `tfsdk:"sni" json:"sni,computed"`
	Increment                types.Int64                                                                             `tfsdk:"increment" json:"increment,computed"`
	Content                  types.String                                                                            `tfsdk:"content" json:"content,computed"`
	ContentType              types.String                                                                            `tfsdk:"content_type" json:"content_type,computed"`
	StatusCode               types.Float64                                                                           `tfsdk:"status_code" json:"status_code,computed"`
	AutomaticHTTPSRewrites   types.Bool                                                                              `tfsdk:"automatic_https_rewrites" json:"automatic_https_rewrites,computed"`
	Autominify               customfield.NestedObject[RulesetRulesActionParametersAutominifyDataSourceModel]         `tfsdk:"autominify" json:"autominify,computed"`
	BIC                      types.Bool                                                                              `tfsdk:"bic" json:"bic,computed"`
	DisableApps              types.Bool                                                                              `tfsdk:"disable_apps" json:"disable_apps,computed"`
	DisableRUM               types.Bool                                                                              `tfsdk:"disable_rum" json:"disable_rum,computed"`
	DisableZaraz             types.Bool                                                                              `tfsdk:"disable_zaraz" json:"disable_zaraz,computed"`
	EmailObfuscation         types.Bool                                                                              `tfsdk:"email_obfuscation" json:"email_obfuscation,computed"`
	Fonts                    types.Bool                                                                              `tfsdk:"fonts" json:"fonts,computed"`
	HotlinkProtection        types.Bool                                                                              `tfsdk:"hotlink_protection" json:"hotlink_protection,computed"`
	Mirage                   types.Bool                                                                              `tfsdk:"mirage" json:"mirage,computed"`
	OpportunisticEncryption  types.Bool                                                                              `tfsdk:"opportunistic_encryption" json:"opportunistic_encryption,computed"`
	Polish                   types.String                                                                            `tfsdk:"polish" json:"polish,computed"`
	RocketLoader             types.Bool                                                                              `tfsdk:"rocket_loader" json:"rocket_loader,computed"`
	SecurityLevel            types.String                                                                            `tfsdk:"security_level" json:"security_level,computed"`
	ServerSideExcludes       types.Bool                                                                              `tfsdk:"server_side_excludes" json:"server_side_excludes,computed"`
	SSL                      types.String                                                                            `tfsdk:"ssl" json:"ssl,computed"`
	SXG                      types.Bool                                                                              `tfsdk:"sxg" json:"sxg,computed"`
	Phases                   customfield.List[types.String]                                                          `tfsdk:"phases" json:"phases,computed"`
	Products                 customfield.List[types.String]                                                          `tfsdk:"products" json:"products,computed"`
	Rules                    customfield.Map[customfield.List[types.String]]                                         `tfsdk:"rules" json:"rules,computed"`
	Ruleset                  types.String                                                                            `tfsdk:"ruleset" json:"ruleset,computed"`
	Rulesets                 customfield.List[types.String]                                                          `tfsdk:"rulesets" json:"rulesets,computed"`
	AdditionalCacheablePorts customfield.List[types.Int64]                                                           `tfsdk:"additional_cacheable_ports" json:"additional_cacheable_ports,computed"`
	BrowserTTL               customfield.NestedObject[RulesetRulesActionParametersBrowserTTLDataSourceModel]         `tfsdk:"browser_ttl" json:"browser_ttl,computed"`
	Cache                    types.Bool                                                                              `tfsdk:"cache" json:"cache,computed"`
	CacheKey                 customfield.NestedObject[RulesetRulesActionParametersCacheKeyDataSourceModel]           `tfsdk:"cache_key" json:"cache_key,computed"`
	CacheReserve             customfield.NestedObject[RulesetRulesActionParametersCacheReserveDataSourceModel]       `tfsdk:"cache_reserve" json:"cache_reserve,computed"`
	EdgeTTL                  customfield.NestedObject[RulesetRulesActionParametersEdgeTTLDataSourceModel]            `tfsdk:"edge_ttl" json:"edge_ttl,computed"`
	OriginCacheControl       types.Bool                                                                              `tfsdk:"origin_cache_control" json:"origin_cache_control,computed"`
	OriginErrorPagePassthru  types.Bool                                                                              `tfsdk:"origin_error_page_passthru" json:"origin_error_page_passthru,computed"`
	ReadTimeout              types.Int64                                                                             `tfsdk:"read_timeout" json:"read_timeout,computed"`
	RespectStrongEtags       types.Bool                                                                              `tfsdk:"respect_strong_etags" json:"respect_strong_etags,computed"`
	ServeStale               customfield.NestedObject[RulesetRulesActionParametersServeStaleDataSourceModel]         `tfsdk:"serve_stale" json:"serve_stale,computed"`
	CookieFields             customfield.NestedObjectList[RulesetRulesActionParametersCookieFieldsDataSourceModel]   `tfsdk:"cookie_fields" json:"cookie_fields,computed"`
	RequestFields            customfield.NestedObjectList[RulesetRulesActionParametersRequestFieldsDataSourceModel]  `tfsdk:"request_fields" json:"request_fields,computed"`
	ResponseFields           customfield.NestedObjectList[RulesetRulesActionParametersResponseFieldsDataSourceModel] `tfsdk:"response_fields" json:"response_fields,computed"`
}

type RulesetRulesActionParametersResponseDataSourceModel struct {
	Content     types.String `tfsdk:"content" json:"content,computed"`
	ContentType types.String `tfsdk:"content_type" json:"content_type,computed"`
	StatusCode  types.Int64  `tfsdk:"status_code" json:"status_code,computed"`
}

type RulesetRulesActionParametersAlgorithmsDataSourceModel struct {
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type RulesetRulesActionParametersMatchedDataDataSourceModel struct {
	PublicKey types.String `tfsdk:"public_key" json:"public_key,computed"`
}

type RulesetRulesActionParametersOverridesDataSourceModel struct {
	Action           types.String                                                                                 `tfsdk:"action" json:"action,computed"`
	Categories       customfield.NestedObjectList[RulesetRulesActionParametersOverridesCategoriesDataSourceModel] `tfsdk:"categories" json:"categories,computed"`
	Enabled          types.Bool                                                                                   `tfsdk:"enabled" json:"enabled,computed"`
	Rules            customfield.NestedObjectList[RulesetRulesActionParametersOverridesRulesDataSourceModel]      `tfsdk:"rules" json:"rules,computed"`
	SensitivityLevel types.String                                                                                 `tfsdk:"sensitivity_level" json:"sensitivity_level,computed"`
}

type RulesetRulesActionParametersOverridesCategoriesDataSourceModel struct {
	Category         types.String `tfsdk:"category" json:"category,computed"`
	Action           types.String `tfsdk:"action" json:"action,computed"`
	Enabled          types.Bool   `tfsdk:"enabled" json:"enabled,computed"`
	SensitivityLevel types.String `tfsdk:"sensitivity_level" json:"sensitivity_level,computed"`
}

type RulesetRulesActionParametersOverridesRulesDataSourceModel struct {
	ID               types.String `tfsdk:"id" json:"id,computed"`
	Action           types.String `tfsdk:"action" json:"action,computed"`
	Enabled          types.Bool   `tfsdk:"enabled" json:"enabled,computed"`
	ScoreThreshold   types.Int64  `tfsdk:"score_threshold" json:"score_threshold,computed"`
	SensitivityLevel types.String `tfsdk:"sensitivity_level" json:"sensitivity_level,computed"`
}

type RulesetRulesActionParametersFromListDataSourceModel struct {
	Key  types.String `tfsdk:"key" json:"key,computed"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type RulesetRulesActionParametersFromValueDataSourceModel struct {
	PreserveQueryString types.Bool                                                                              `tfsdk:"preserve_query_string" json:"preserve_query_string,computed"`
	StatusCode          types.Float64                                                                           `tfsdk:"status_code" json:"status_code,computed"`
	TargetURL           customfield.NestedObject[RulesetRulesActionParametersFromValueTargetURLDataSourceModel] `tfsdk:"target_url" json:"target_url,computed"`
}

type RulesetRulesActionParametersFromValueTargetURLDataSourceModel struct {
	Value      types.String `tfsdk:"value" json:"value,computed"`
	Expression types.String `tfsdk:"expression" json:"expression,computed"`
}

type RulesetRulesActionParametersHeadersDataSourceModel struct {
	Operation  types.String `tfsdk:"operation" json:"operation,computed"`
	Value      types.String `tfsdk:"value" json:"value,computed"`
	Expression types.String `tfsdk:"expression" json:"expression,computed"`
}

type RulesetRulesActionParametersURIDataSourceModel struct {
	Path  customfield.NestedObject[RulesetRulesActionParametersURIPathDataSourceModel]  `tfsdk:"path" json:"path,computed"`
	Query customfield.NestedObject[RulesetRulesActionParametersURIQueryDataSourceModel] `tfsdk:"query" json:"query,computed"`
}

type RulesetRulesActionParametersURIPathDataSourceModel struct {
	Value      types.String `tfsdk:"value" json:"value,computed"`
	Expression types.String `tfsdk:"expression" json:"expression,computed"`
}

type RulesetRulesActionParametersURIQueryDataSourceModel struct {
	Value      types.String `tfsdk:"value" json:"value,computed"`
	Expression types.String `tfsdk:"expression" json:"expression,computed"`
}

type RulesetRulesActionParametersOriginDataSourceModel struct {
	Host types.String  `tfsdk:"host" json:"host,computed"`
	Port types.Float64 `tfsdk:"port" json:"port,computed"`
}

type RulesetRulesActionParametersSNIDataSourceModel struct {
	Value types.String `tfsdk:"value" json:"value,computed"`
}

type RulesetRulesActionParametersAutominifyDataSourceModel struct {
	CSS  types.Bool `tfsdk:"css" json:"css,computed"`
	HTML types.Bool `tfsdk:"html" json:"html,computed"`
	JS   types.Bool `tfsdk:"js" json:"js,computed"`
}

type RulesetRulesActionParametersBrowserTTLDataSourceModel struct {
	Mode    types.String `tfsdk:"mode" json:"mode,computed"`
	Default types.Int64  `tfsdk:"default" json:"default,computed"`
}

type RulesetRulesActionParametersCacheKeyDataSourceModel struct {
	CacheByDeviceType       types.Bool                                                                             `tfsdk:"cache_by_device_type" json:"cache_by_device_type,computed"`
	CacheDeceptionArmor     types.Bool                                                                             `tfsdk:"cache_deception_armor" json:"cache_deception_armor,computed"`
	CustomKey               customfield.NestedObject[RulesetRulesActionParametersCacheKeyCustomKeyDataSourceModel] `tfsdk:"custom_key" json:"custom_key,computed"`
	IgnoreQueryStringsOrder types.Bool                                                                             `tfsdk:"ignore_query_strings_order" json:"ignore_query_strings_order,computed"`
}

type RulesetRulesActionParametersCacheKeyCustomKeyDataSourceModel struct {
	Cookie      customfield.NestedObject[RulesetRulesActionParametersCacheKeyCustomKeyCookieDataSourceModel]      `tfsdk:"cookie" json:"cookie,computed"`
	Header      customfield.NestedObject[RulesetRulesActionParametersCacheKeyCustomKeyHeaderDataSourceModel]      `tfsdk:"header" json:"header,computed"`
	Host        customfield.NestedObject[RulesetRulesActionParametersCacheKeyCustomKeyHostDataSourceModel]        `tfsdk:"host" json:"host,computed"`
	QueryString customfield.NestedObject[RulesetRulesActionParametersCacheKeyCustomKeyQueryStringDataSourceModel] `tfsdk:"query_string" json:"query_string,computed"`
	User        customfield.NestedObject[RulesetRulesActionParametersCacheKeyCustomKeyUserDataSourceModel]        `tfsdk:"user" json:"user,computed"`
}

type RulesetRulesActionParametersCacheKeyCustomKeyCookieDataSourceModel struct {
	CheckPresence customfield.List[types.String] `tfsdk:"check_presence" json:"check_presence,computed"`
	Include       customfield.List[types.String] `tfsdk:"include" json:"include,computed"`
}

type RulesetRulesActionParametersCacheKeyCustomKeyHeaderDataSourceModel struct {
	CheckPresence customfield.List[types.String]                  `tfsdk:"check_presence" json:"check_presence,computed"`
	Contains      customfield.Map[customfield.List[types.String]] `tfsdk:"contains" json:"contains,computed"`
	ExcludeOrigin types.Bool                                      `tfsdk:"exclude_origin" json:"exclude_origin,computed"`
	Include       customfield.List[types.String]                  `tfsdk:"include" json:"include,computed"`
}

type RulesetRulesActionParametersCacheKeyCustomKeyHostDataSourceModel struct {
	Resolved types.Bool `tfsdk:"resolved" json:"resolved,computed"`
}

type RulesetRulesActionParametersCacheKeyCustomKeyQueryStringDataSourceModel struct {
	Exclude customfield.NestedObject[RulesetRulesActionParametersCacheKeyCustomKeyQueryStringExcludeDataSourceModel] `tfsdk:"exclude" json:"exclude,computed"`
	Include customfield.NestedObject[RulesetRulesActionParametersCacheKeyCustomKeyQueryStringIncludeDataSourceModel] `tfsdk:"include" json:"include,computed"`
}

type RulesetRulesActionParametersCacheKeyCustomKeyQueryStringExcludeDataSourceModel struct {
	All  types.Bool                     `tfsdk:"all" json:"all,computed"`
	List customfield.List[types.String] `tfsdk:"list" json:"list,computed"`
}

type RulesetRulesActionParametersCacheKeyCustomKeyQueryStringIncludeDataSourceModel struct {
	All  types.Bool                     `tfsdk:"all" json:"all,computed"`
	List customfield.List[types.String] `tfsdk:"list" json:"list,computed"`
}

type RulesetRulesActionParametersCacheKeyCustomKeyUserDataSourceModel struct {
	DeviceType types.Bool `tfsdk:"device_type" json:"device_type,computed"`
	Geo        types.Bool `tfsdk:"geo" json:"geo,computed"`
	Lang       types.Bool `tfsdk:"lang" json:"lang,computed"`
}

type RulesetRulesActionParametersCacheReserveDataSourceModel struct {
	Eligible        types.Bool  `tfsdk:"eligible" json:"eligible,computed"`
	MinimumFileSize types.Int64 `tfsdk:"minimum_file_size" json:"minimum_file_size,computed"`
}

type RulesetRulesActionParametersEdgeTTLDataSourceModel struct {
	Default       types.Int64                                                                                   `tfsdk:"default" json:"default,computed"`
	Mode          types.String                                                                                  `tfsdk:"mode" json:"mode,computed"`
	StatusCodeTTL customfield.NestedObjectList[RulesetRulesActionParametersEdgeTTLStatusCodeTTLDataSourceModel] `tfsdk:"status_code_ttl" json:"status_code_ttl,computed"`
}

type RulesetRulesActionParametersEdgeTTLStatusCodeTTLDataSourceModel struct {
	Value           types.Int64                                                                                              `tfsdk:"value" json:"value,computed"`
	StatusCodeRange customfield.NestedObject[RulesetRulesActionParametersEdgeTTLStatusCodeTTLStatusCodeRangeDataSourceModel] `tfsdk:"status_code_range" json:"status_code_range,computed"`
	StatusCodeValue types.Int64                                                                                              `tfsdk:"status_code_value" json:"status_code_value,computed"`
}

type RulesetRulesActionParametersEdgeTTLStatusCodeTTLStatusCodeRangeDataSourceModel struct {
	From types.Int64 `tfsdk:"from" json:"from,computed"`
	To   types.Int64 `tfsdk:"to" json:"to,computed"`
}

type RulesetRulesActionParametersServeStaleDataSourceModel struct {
	DisableStaleWhileUpdating types.Bool `tfsdk:"disable_stale_while_updating" json:"disable_stale_while_updating,computed"`
}

type RulesetRulesActionParametersCookieFieldsDataSourceModel struct {
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type RulesetRulesActionParametersRequestFieldsDataSourceModel struct {
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type RulesetRulesActionParametersResponseFieldsDataSourceModel struct {
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type RulesetRulesExposedCredentialCheckDataSourceModel struct {
	PasswordExpression types.String `tfsdk:"password_expression" json:"password_expression,computed"`
	UsernameExpression types.String `tfsdk:"username_expression" json:"username_expression,computed"`
}

type RulesetRulesLoggingDataSourceModel struct {
	Enabled types.Bool `tfsdk:"enabled" json:"enabled,computed"`
}

type RulesetRulesRatelimitDataSourceModel struct {
	Characteristics         customfield.List[types.String] `tfsdk:"characteristics" json:"characteristics,computed"`
	Period                  types.Int64                    `tfsdk:"period" json:"period,computed"`
	CountingExpression      types.String                   `tfsdk:"counting_expression" json:"counting_expression,computed"`
	MitigationTimeout       types.Int64                    `tfsdk:"mitigation_timeout" json:"mitigation_timeout,computed"`
	RequestsPerPeriod       types.Int64                    `tfsdk:"requests_per_period" json:"requests_per_period,computed"`
	RequestsToOrigin        types.Bool                     `tfsdk:"requests_to_origin" json:"requests_to_origin,computed"`
	ScorePerPeriod          types.Int64                    `tfsdk:"score_per_period" json:"score_per_period,computed"`
	ScoreResponseHeaderName types.String                   `tfsdk:"score_response_header_name" json:"score_response_header_name,computed"`
}

type RulesetFindOneByDataSourceModel struct {
	AccountID types.String `tfsdk:"account_id" path:"account_id,optional"`
	ZoneID    types.String `tfsdk:"zone_id" path:"zone_id,optional"`
}

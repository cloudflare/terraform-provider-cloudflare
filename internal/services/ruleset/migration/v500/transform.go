package v500

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Transform converts source (v4 Plugin Framework) state to target (v5 Plugin Framework) state.
// All field transformations are handled here:
//   - MaxItems:1 ListNestedBlock arrays → SingleNestedAttribute pointers
//   - TypeSet fields → TypeList fields
//   - Set[string] → List[{name: string}] (cookie_fields, request_fields, response_fields)
//   - headers: List[{name, ...}] → Map[name → {...}]
//   - query_string include/exclude: Set[string] → {list:[...]} or {all:true}
//   - disable_railgun: removed
//   - rules map: map[string]string (CSV) → map[string][]string
func Transform(ctx context.Context, source SourceV4RulesetModel) (*TargetV5RulesetModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	target := &TargetV5RulesetModel{
		ID:          source.ID,
		AccountID:   source.AccountID,
		ZoneID:      source.ZoneID,
		Kind:        source.Kind,
		Name:        source.Name,
		Phase:       source.Phase,
		Description: source.Description,
		// LastUpdated and Version are computed by the API - set to null
		LastUpdated: timetypes.NewRFC3339Null(),
		Version:     types.StringNull(),
	}

	// Transform rules
	if len(source.Rules) > 0 {
		rules := make([]*TargetV5RuleModel, 0, len(source.Rules))
		for _, srcRule := range source.Rules {
			if srcRule == nil {
				continue
			}
			rule, ruleDiags := transformRule(ctx, srcRule)
			diags.Append(ruleDiags...)
			if diags.HasError() {
				return nil, diags
			}
			rules = append(rules, rule)
		}
		target.Rules = rules
	} else {
		target.Rules = []*TargetV5RuleModel{}
	}

	return target, diags
}

// transformRule converts a single v4 rule to v5 format.
func transformRule(ctx context.Context, src *SourceV4RuleModel) (*TargetV5RuleModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	target := &TargetV5RuleModel{
		ID:          src.ID,
		Action:      src.Action,
		Description: src.Description,
		Enabled:     src.Enabled,
		Expression:  src.Expression,
		Ref:         src.Ref,
	}

	// action_parameters: ListNestedBlock MaxItems:1 → SingleNestedAttribute pointer
	if len(src.ActionParameters) > 0 {
		ap, apDiags := transformActionParameters(ctx, src.ActionParameters[0])
		diags.Append(apDiags...)
		if diags.HasError() {
			return nil, diags
		}
		target.ActionParameters = ap
	}

	// exposed_credential_check: ListNestedBlock MaxItems:1 → SingleNestedAttribute pointer
	if len(src.ExposedCredentialCheck) > 0 {
		ec := src.ExposedCredentialCheck[0]
		target.ExposedCredentialCheck = &TargetV5ExposedCredentialCheckModel{
			PasswordExpression: ec.PasswordExpression,
			UsernameExpression: ec.UsernameExpression,
		}
	}

	// logging: ListNestedBlock MaxItems:1 → SingleNestedAttribute pointer
	if len(src.Logging) > 0 {
		target.Logging = &TargetV5LoggingModel{
			Enabled: src.Logging[0].Enabled,
		}
	}

	// ratelimit: ListNestedBlock MaxItems:1 → SingleNestedAttribute pointer
	if len(src.Ratelimit) > 0 {
		rl, rlDiags := transformRatelimit(ctx, src.Ratelimit[0])
		diags.Append(rlDiags...)
		if diags.HasError() {
			return nil, diags
		}
		target.Ratelimit = rl
	}

	return target, diags
}

// transformActionParameters converts a v4 action_parameters block to v5 format.
func transformActionParameters(ctx context.Context, src *SourceV4ActionParametersModel) (*TargetV5ActionParametersModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	target := &TargetV5ActionParametersModel{
		// Direct scalar copies
		ID:                      src.ID,
		HostHeader:              src.HostHeader,
		Increment:               src.Increment,
		Content:                 src.Content,
		ContentType:             src.ContentType,
		StatusCode:              src.StatusCode,
		AutomaticHTTPSRewrites:  src.AutomaticHTTPSRewrites,
		BIC:                     src.BIC,
		Cache:                   src.Cache,
		DisableApps:             src.DisableApps,
		DisableRUM:              src.DisableRUM,
		DisableZaraz:            src.DisableZaraz,
		EmailObfuscation:        src.EmailObfuscation,
		Fonts:                   src.Fonts,
		HotlinkProtection:       src.HotlinkProtection,
		Mirage:                  src.Mirage,
		OpportunisticEncryption: src.OpportunisticEncryption,
		OriginCacheControl:      src.OriginCacheControl,
		OriginErrorPagePassthru: src.OriginErrorPagePassthru,
		ReadTimeout:             src.ReadTimeout,
		RespectStrongEtags:      src.RespectStrongEtags,
		RocketLoader:            src.RocketLoader,
		SecurityLevel:           src.SecurityLevel,
		ServerSideExcludes:      src.ServerSideExcludes,
		SSL:                     src.SSL,
		SXG:                     src.SXG,
		Polish:                  src.Polish,
		Ruleset:                 src.Ruleset,

		// New v5 fields absent in v4 — set to null
		AssetName:              types.StringNull(),
		ContentConverter:       types.BoolNull(),
		RequestBodyBuffering:   types.StringNull(),
		ResponseBodyBuffering:  types.StringNull(),
		RedirectsForAITraining: types.BoolNull(),

		// New v5 list fields absent in v4 — set to empty
		RawResponseFields:        nil,
		TransformedRequestFields: nil,
	}

	// Note: DisableRailgun is intentionally NOT copied (removed in v5)

	// products: TypeSet → []types.String
	if !src.Products.IsNull() && !src.Products.IsUnknown() {
		products, setDiags := convertSetToStringSlice(ctx, src.Products)
		diags.Append(setDiags...)
		target.Products = products
	}

	// phases: TypeSet → []types.String
	if !src.Phases.IsNull() && !src.Phases.IsUnknown() {
		phases, setDiags := convertSetToStringSlice(ctx, src.Phases)
		diags.Append(setDiags...)
		target.Phases = phases
	}

	// rulesets: TypeSet → []types.String
	if !src.Rulesets.IsNull() && !src.Rulesets.IsUnknown() {
		rulesets, setDiags := convertSetToStringSlice(ctx, src.Rulesets)
		diags.Append(setDiags...)
		target.Rulesets = rulesets
	}

	// additional_cacheable_ports: TypeSet[Int64] → []types.Int64
	if !src.AdditionalCacheablePorts.IsNull() && !src.AdditionalCacheablePorts.IsUnknown() {
		ports, portsDiags := convertSetToInt64Slice(ctx, src.AdditionalCacheablePorts)
		diags.Append(portsDiags...)
		target.AdditionalCacheablePorts = ports
	}

	// cookie_fields: TypeSet[String] → []*TargetV5FieldNameModel{{name: "val"}}
	if !src.CookieFields.IsNull() && !src.CookieFields.IsUnknown() {
		names, setDiags := convertSetToStringSlice(ctx, src.CookieFields)
		diags.Append(setDiags...)
		target.CookieFields = stringSliceToFieldNameList(names)
	}

	// request_fields: TypeSet[String] → []*TargetV5FieldNameModel{{name: "val"}}
	if !src.RequestFields.IsNull() && !src.RequestFields.IsUnknown() {
		names, setDiags := convertSetToStringSlice(ctx, src.RequestFields)
		diags.Append(setDiags...)
		target.RequestFields = stringSliceToFieldNameList(names)
	}

	// response_fields: TypeSet[String] → []*TargetV5ResponseFieldModel{{name: "val", preserve_duplicates: false}}
	if !src.ResponseFields.IsNull() && !src.ResponseFields.IsUnknown() {
		names, setDiags := convertSetToStringSlice(ctx, src.ResponseFields)
		diags.Append(setDiags...)
		target.ResponseFields = stringSliceToResponseFieldList(names)
	}

	// headers: []*SourceV4HeaderModel (list with name) → map[string]*TargetV5HeaderModel (keyed by name)
	target.Headers = transformHeaders(src.Headers)

	// rules map: map[string]string (CSV) → map[string][]types.String
	if src.Rules != nil {
		target.Rules = transformRulesMap(src.Rules)
	}

	// algorithms: []*SourceV4AlgorithmModel → []*TargetV5AlgorithmModel (name: raw string → types.String)
	target.Algorithms = transformAlgorithms(src.Algorithms)

	// uri: ListNestedBlock MaxItems:1 → *TargetV5URIModel (drops origin field)
	if len(src.URI) > 0 {
		target.URI = transformURI(src.URI[0])
	}

	// autominify: ListNestedBlock MaxItems:1 → pointer
	if len(src.AutoMinify) > 0 {
		am := src.AutoMinify[0]
		target.Autominify = &TargetV5AutoMinifyModel{
			CSS:  am.CSS,
			HTML: am.HTML,
			JS:   am.JS,
		}
	}

	// browser_ttl: ListNestedBlock MaxItems:1 → pointer
	if len(src.BrowserTTL) > 0 {
		bt := src.BrowserTTL[0]
		target.BrowserTTL = &TargetV5BrowserTTLModel{
			Mode:    bt.Mode,
			Default: bt.Default,
		}
	}

	// cache_key: ListNestedBlock MaxItems:1 → pointer
	if len(src.CacheKey) > 0 {
		ck, ckDiags := transformCacheKey(ctx, src.CacheKey[0])
		diags.Append(ckDiags...)
		target.CacheKey = ck
	}

	// cache_reserve: ListNestedBlock MaxItems:1 → pointer
	if len(src.CacheReserve) > 0 {
		cr := src.CacheReserve[0]
		target.CacheReserve = &TargetV5CacheReserveModel{
			Eligible:        cr.Eligible,
			MinimumFileSize: cr.MinimumFileSize,
		}
	}

	// edge_ttl: ListNestedBlock MaxItems:1 → pointer
	if len(src.EdgeTTL) > 0 {
		et, etDiags := transformEdgeTTL(src.EdgeTTL[0])
		diags.Append(etDiags...)
		target.EdgeTTL = et
	}

	// from_list: ListNestedBlock MaxItems:1 → pointer
	if len(src.FromList) > 0 {
		fl := src.FromList[0]
		target.FromList = &TargetV5FromListModel{
			Key:  fl.Key,
			Name: fl.Name,
		}
	}

	// from_value: ListNestedBlock MaxItems:1 → pointer
	if len(src.FromValue) > 0 {
		fv := src.FromValue[0]
		target.FromValue = &TargetV5FromValueModel{
			StatusCode:          fv.StatusCode,
			PreserveQueryString: fv.PreserveQueryString,
		}
		// target_url: ListNestedBlock MaxItems:1 inside from_value → pointer
		if len(fv.TargetURL) > 0 {
			tu := fv.TargetURL[0]
			target.FromValue.TargetURL = &TargetV5TargetURLModel{
				Value:      tu.Value,
				Expression: tu.Expression,
			}
		}
	}

	// matched_data: ListNestedBlock MaxItems:1 → pointer
	if len(src.MatchedData) > 0 {
		target.MatchedData = &TargetV5MatchedDataModel{
			PublicKey: src.MatchedData[0].PublicKey,
		}
	}

	// overrides: ListNestedBlock MaxItems:1 → pointer
	if len(src.Overrides) > 0 {
		target.Overrides = transformOverrides(src.Overrides[0])
	}

	// origin: ListNestedBlock MaxItems:1 → pointer
	if len(src.Origin) > 0 {
		o := src.Origin[0]
		target.Origin = &TargetV5OriginModel{
			Host: o.Host,
			Port: o.Port,
		}
	}

	// sni: ListNestedBlock MaxItems:1 → pointer
	if len(src.SNI) > 0 {
		target.SNI = &TargetV5SNIModel{
			Value: src.SNI[0].Value,
		}
	}

	// serve_stale: ListNestedBlock MaxItems:1 → pointer
	if len(src.ServeStale) > 0 {
		target.ServeStale = &TargetV5ServeStaleModel{
			DisableStaleWhileUpdating: src.ServeStale[0].DisableStaleWhileUpdating,
		}
	}

	// response: ListNestedBlock MaxItems:1 → pointer
	if len(src.Response) > 0 {
		r := src.Response[0]
		target.Response = &TargetV5ResponseModel{
			StatusCode:  r.StatusCode,
			ContentType: r.ContentType,
			Content:     r.Content,
		}
	}

	return target, diags
}

// transformHeaders converts v4 headers list to v5 headers map.
// v4: [{name: "X-Hdr", operation: "set", value: "val", expression: ""}]
// v5: {"X-Hdr": {operation: "set", value: "val", expression: ""}}
func transformHeaders(srcHeaders []*SourceV4HeaderModel) map[string]*TargetV5HeaderModel {
	if len(srcHeaders) == 0 {
		return nil
	}

	result := make(map[string]*TargetV5HeaderModel, len(srcHeaders))
	for _, h := range srcHeaders {
		if h == nil || h.Name.IsNull() || h.Name.IsUnknown() {
			continue
		}
		name := h.Name.ValueString()
		result[name] = &TargetV5HeaderModel{
			Operation:  h.Operation,
			Value:      h.Value,
			Expression: h.Expression,
		}
	}
	if len(result) == 0 {
		return nil
	}
	return result
}

// transformRulesMap converts the action_parameters.rules map from CSV strings to string slices.
// v4: {"efb7b8c949ac4650a09736fc376e9aee": "5de7edfa648c4d6891dc3e7f84534ffa,e3a567afc347477d9702d9047e97d760"}
// v5: {"efb7b8c949ac4650a09736fc376e9aee": ["5de7edfa648c4d6891dc3e7f84534ffa", "e3a567afc347477d9702d9047e97d760"]}
func transformRulesMap(srcRules map[string]types.String) map[string][]types.String {
	if len(srcRules) == 0 {
		return nil
	}

	result := make(map[string][]types.String, len(srcRules))
	for k, v := range srcRules {
		if v.IsNull() || v.IsUnknown() || v.ValueString() == "" {
			result[k] = []types.String{}
			continue
		}
		parts := strings.Split(v.ValueString(), ",")
		ruleIDs := make([]types.String, 0, len(parts))
		for _, p := range parts {
			trimmed := strings.TrimSpace(p)
			if trimmed != "" {
				ruleIDs = append(ruleIDs, types.StringValue(trimmed))
			}
		}
		result[k] = ruleIDs
	}
	return result
}

// transformAlgorithms converts v4 algorithms (Name as raw string) to v5 format (Name as types.String).
func transformAlgorithms(srcAlgs []*SourceV4AlgorithmModel) []*TargetV5AlgorithmModel {
	if len(srcAlgs) == 0 {
		return nil
	}

	result := make([]*TargetV5AlgorithmModel, 0, len(srcAlgs))
	for _, a := range srcAlgs {
		if a == nil {
			continue
		}
		result = append(result, &TargetV5AlgorithmModel{
			Name: types.StringValue(a.Name),
		})
	}
	return result
}

// transformURI converts v4 uri block to v5 format.
// v4: {origin: bool, path: [{value, expression}], query: [{value, expression}]}
// v5: {path: {value, expression}, query: {value, expression}} (origin removed)
func transformURI(src *SourceV4URIModel) *TargetV5URIModel {
	if src == nil {
		return nil
	}

	target := &TargetV5URIModel{}

	// path: ListNestedBlock MaxItems:1 → pointer
	if len(src.Path) > 0 {
		p := src.Path[0]
		target.Path = &TargetV5URIPartModel{
			Value:      p.Value,
			Expression: p.Expression,
		}
	}

	// query: ListNestedBlock MaxItems:1 → pointer
	if len(src.Query) > 0 {
		q := src.Query[0]
		target.Query = &TargetV5URIPartModel{
			Value:      q.Value,
			Expression: q.Expression,
		}
	}

	// origin (bool) is intentionally NOT copied — removed in v5

	return target
}

// transformOverrides converts v4 overrides block to v5 format.
// overrides is MaxItems:1 in v4; categories and rules are lists.
func transformOverrides(src *SourceV4OverridesModel) *TargetV5OverridesModel {
	if src == nil {
		return nil
	}

	target := &TargetV5OverridesModel{
		Action:           src.Action,
		Enabled:          src.Enabled,
		SensitivityLevel: src.SensitivityLevel,
	}

	// categories: list copy (no transformation needed)
	if len(src.Categories) > 0 {
		cats := make([]*TargetV5OverridesCategoryModel, 0, len(src.Categories))
		for _, c := range src.Categories {
			if c == nil {
				continue
			}
			cats = append(cats, &TargetV5OverridesCategoryModel{
				Category:         c.Category,
				Action:           c.Action,
				Enabled:          c.Enabled,
				SensitivityLevel: types.StringNull(), // new in v5
			})
		}
		target.Categories = cats
	}

	// rules: list copy (no transformation needed)
	if len(src.Rules) > 0 {
		rules := make([]*TargetV5OverridesRuleModel, 0, len(src.Rules))
		for _, r := range src.Rules {
			if r == nil {
				continue
			}
			rules = append(rules, &TargetV5OverridesRuleModel{
				ID:               r.ID,
				Action:           r.Action,
				Enabled:          r.Enabled,
				ScoreThreshold:   r.ScoreThreshold,
				SensitivityLevel: r.SensitivityLevel,
			})
		}
		target.Rules = rules
	}

	return target
}

// transformCacheKey converts v4 cache_key block to v5 format.
func transformCacheKey(ctx context.Context, src *SourceV4CacheKeyModel) (*TargetV5CacheKeyModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	target := &TargetV5CacheKeyModel{
		CacheByDeviceType:       src.CacheByDeviceType,
		IgnoreQueryStringsOrder: src.IgnoreQueryStringsOrder,
		CacheDeceptionArmor:     src.CacheDeceptionArmor,
	}

	// custom_key: ListNestedBlock MaxItems:1 → pointer
	if len(src.CustomKey) > 0 {
		ck, ckDiags := transformCustomKey(ctx, src.CustomKey[0])
		diags.Append(ckDiags...)
		target.CustomKey = ck
	}

	return target, diags
}

// transformCustomKey converts v4 custom_key to v5 format.
func transformCustomKey(ctx context.Context, src *SourceV4CacheKeyCustomKeyModel) (*TargetV5CacheKeyCustomKeyModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	target := &TargetV5CacheKeyCustomKeyModel{}

	// query_string: ListNestedBlock MaxItems:1 → pointer with include/exclude transformation
	if len(src.QueryString) > 0 {
		qs, qsDiags := transformCKQueryString(ctx, src.QueryString[0])
		diags.Append(qsDiags...)
		target.QueryString = qs
	}

	// header: ListNestedBlock MaxItems:1 → pointer
	if len(src.Header) > 0 {
		h, hDiags := transformCKHeader(ctx, src.Header[0])
		diags.Append(hDiags...)
		target.Header = h
	}

	// cookie: ListNestedBlock MaxItems:1 → pointer
	if len(src.Cookie) > 0 {
		c, cDiags := transformCKCookie(ctx, src.Cookie[0])
		diags.Append(cDiags...)
		target.Cookie = c
	}

	// user: ListNestedBlock MaxItems:1 → pointer (direct copy)
	if len(src.User) > 0 {
		u := src.User[0]
		target.User = &TargetV5CKUserModel{
			DeviceType: u.DeviceType,
			Geo:        u.Geo,
			Lang:       u.Lang,
		}
	}

	// host: ListNestedBlock MaxItems:1 → pointer (direct copy)
	if len(src.Host) > 0 {
		target.Host = &TargetV5CKHostModel{
			Resolved: src.Host[0].Resolved,
		}
	}

	return target, diags
}

// transformCKQueryString converts v4 custom_key.query_string to v5 format.
// v4: {include: Set[String], exclude: Set[String]}
// v5: {include: {list: [...]} or {all: true}, exclude: {list: [...]}}
func transformCKQueryString(ctx context.Context, src *SourceV4QueryStringModel) (*TargetV5CKQueryStringModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	target := &TargetV5CKQueryStringModel{}

	if !src.Include.IsNull() && !src.Include.IsUnknown() {
		includeStrings, setDiags := convertSetToStringSlice(ctx, src.Include)
		diags.Append(setDiags...)
		if !diags.HasError() {
			if len(includeStrings) == 1 && includeStrings[0].ValueString() == "*" {
				// Special case: ["*"] → {all: true}
				target.Include = &TargetV5QSIncludeModel{
					All:  types.BoolValue(true),
					List: nil,
				}
			} else {
				target.Include = &TargetV5QSIncludeModel{
					All:  types.BoolValue(false),
					List: includeStrings,
				}
			}
		}
	}

	if !src.Exclude.IsNull() && !src.Exclude.IsUnknown() {
		excludeStrings, setDiags := convertSetToStringSlice(ctx, src.Exclude)
		diags.Append(setDiags...)
		if !diags.HasError() {
			target.Exclude = &TargetV5QSExcludeModel{
				All:  types.BoolValue(false),
				List: excludeStrings,
			}
		}
	}

	return target, diags
}

// transformCKHeader converts v4 custom_key.header to v5 format.
func transformCKHeader(ctx context.Context, src *SourceV4CustomKeyHeaderModel) (*TargetV5CKHeaderModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	target := &TargetV5CKHeaderModel{
		ExcludeOrigin: src.ExcludeOrigin,
	}

	if !src.Include.IsNull() && !src.Include.IsUnknown() {
		include, setDiags := convertSetToStringSlice(ctx, src.Include)
		diags.Append(setDiags...)
		target.Include = include
	}

	if !src.CheckPresence.IsNull() && !src.CheckPresence.IsUnknown() {
		cp, setDiags := convertSetToStringSlice(ctx, src.CheckPresence)
		diags.Append(setDiags...)
		target.CheckPresence = cp
	}

	// contains: map[string]types.Set → map[string][]types.String
	if len(src.Contains) > 0 {
		contains := make(map[string][]types.String, len(src.Contains))
		for k, v := range src.Contains {
			if !v.IsNull() && !v.IsUnknown() {
				vals, setDiags := convertSetToStringSlice(ctx, v)
				diags.Append(setDiags...)
				contains[k] = vals
			}
		}
		target.Contains = contains
	}

	return target, diags
}

// transformCKCookie converts v4 custom_key.cookie to v5 format.
func transformCKCookie(ctx context.Context, src *SourceV4CustomKeyCookieModel) (*TargetV5CKCookieModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	target := &TargetV5CKCookieModel{}

	if !src.Include.IsNull() && !src.Include.IsUnknown() {
		include, setDiags := convertSetToStringSlice(ctx, src.Include)
		diags.Append(setDiags...)
		target.Include = include
	}

	if !src.CheckPresence.IsNull() && !src.CheckPresence.IsUnknown() {
		cp, setDiags := convertSetToStringSlice(ctx, src.CheckPresence)
		diags.Append(setDiags...)
		target.CheckPresence = cp
	}

	return target, diags
}

// transformEdgeTTL converts v4 edge_ttl block to v5 format.
// Handles status_code_ttl list and status_code_range: array[0] → pointer.
func transformEdgeTTL(src *SourceV4EdgeTTLModel) (*TargetV5EdgeTTLModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	target := &TargetV5EdgeTTLModel{
		Mode:    src.Mode,
		Default: src.Default,
	}

	if len(src.StatusCodeTTL) > 0 {
		sctList := make([]*TargetV5StatusCodeTTLModel, 0, len(src.StatusCodeTTL))
		for _, sct := range src.StatusCodeTTL {
			if sct == nil {
				continue
			}
			sctTarget := &TargetV5StatusCodeTTLModel{
				StatusCode: sct.StatusCode,
				Value:      sct.Value,
			}
			// status_code_range: ListNestedBlock MaxItems:1 → pointer
			if len(sct.StatusCodeRange) > 0 {
				r := sct.StatusCodeRange[0]
				sctTarget.StatusCodeRange = &TargetV5StatusCodeRangeModel{
					From: r.From,
					To:   r.To,
				}
			}
			sctList = append(sctList, sctTarget)
		}
		target.StatusCodeTTL = sctList
	}

	return target, diags
}

// transformRatelimit converts v4 ratelimit block to v5 format.
// Key change: characteristics TypeSet → []types.String (List).
func transformRatelimit(ctx context.Context, src *SourceV4RatelimitModel) (*TargetV5RatelimitModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	target := &TargetV5RatelimitModel{
		Period:                  src.Period,
		CountingExpression:      src.CountingExpression,
		MitigationTimeout:       src.MitigationTimeout,
		RequestsPerPeriod:       src.RequestsPerPeriod,
		RequestsToOrigin:        src.RequestsToOrigin,
		ScorePerPeriod:          src.ScorePerPeriod,
		ScoreResponseHeaderName: src.ScoreResponseHeaderName,
	}

	// characteristics: TypeSet → []types.String
	if !src.Characteristics.IsNull() && !src.Characteristics.IsUnknown() {
		chars, setDiags := convertSetToStringSlice(ctx, src.Characteristics)
		diags.Append(setDiags...)
		target.Characteristics = chars
	}

	return target, diags
}

// ============================================================================
// Helper Functions
// ============================================================================

// convertSetToStringSlice converts types.Set (string elements) to []types.String.
func convertSetToStringSlice(ctx context.Context, set types.Set) ([]types.String, diag.Diagnostics) {
	var diags diag.Diagnostics

	var rawStrings []string
	diags.Append(set.ElementsAs(ctx, &rawStrings, false)...)
	if diags.HasError() {
		return nil, diags
	}

	result := make([]types.String, 0, len(rawStrings))
	for _, str := range rawStrings {
		result = append(result, types.StringValue(str))
	}
	return result, diags
}

// convertSetToInt64Slice converts types.Set (int64 elements) to []types.Int64.
func convertSetToInt64Slice(ctx context.Context, set types.Set) ([]types.Int64, diag.Diagnostics) {
	var diags diag.Diagnostics

	var rawInts []int64
	diags.Append(set.ElementsAs(ctx, &rawInts, false)...)
	if diags.HasError() {
		return nil, diags
	}

	result := make([]types.Int64, 0, len(rawInts))
	for _, i := range rawInts {
		result = append(result, types.Int64Value(i))
	}
	return result, diags
}

// stringSliceToFieldNameList converts []types.String to []*TargetV5FieldNameModel.
// Used for cookie_fields, request_fields, transformed_request_fields.
func stringSliceToFieldNameList(names []types.String) []*TargetV5FieldNameModel {
	if len(names) == 0 {
		return nil
	}
	result := make([]*TargetV5FieldNameModel, 0, len(names))
	for _, name := range names {
		result = append(result, &TargetV5FieldNameModel{Name: name})
	}
	return result
}

// stringSliceToResponseFieldList converts []types.String to []*TargetV5ResponseFieldModel.
// Used for response_fields (which has preserve_duplicates in v5, defaulting to false).
func stringSliceToResponseFieldList(names []types.String) []*TargetV5ResponseFieldModel {
	if len(names) == 0 {
		return nil
	}
	result := make([]*TargetV5ResponseFieldModel, 0, len(names))
	for _, name := range names {
		result = append(result, &TargetV5ResponseFieldModel{
			Name:               name,
			PreserveDuplicates: types.BoolValue(false),
		})
	}
	return result
}

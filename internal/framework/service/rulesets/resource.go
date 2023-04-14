package rulesets

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"sort"
	"strings"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/framework/expanders"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/framework/flatteners"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

const (
	accountLevelRulesetDeleteURL = "https://api.cloudflare.com/#account-rulesets-delete-account-ruleset"
	zoneLevelRulesetDeleteURL    = "https://api.cloudflare.com/#zone-rulesets-delete-zone-ruleset"
	duplicateRulesetError        = "A similar configuration with rules already exists and overwriting will have unintended consequences. If you are migrating from the Dashboard, you will need to first remove the existing rules otherwise you can remove the existing phase yourself using the API (%s)."
)

var _ resource.Resource = &RulesetResource{}
var _ resource.ResourceWithImportState = &RulesetResource{}

func NewResource() resource.Resource {
	return &RulesetResource{}
}

type RulesetResource struct {
	client *cloudflare.API
}

func (r *RulesetResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ruleset"
}

func (r *RulesetResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*cloudflare.API)

	if !ok {
		resp.Diagnostics.AddError(
			"unexpected resource configure type",
			fmt.Sprintf("expected *cloudflare.API, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *RulesetResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *RulesetResourceModel

	diags := req.Plan.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	accountID := data.AccountID
	zoneID := data.ZoneID
	rulesetPhase := data.Phase.ValueString()

	var ruleset cloudflare.Ruleset
	var sempahoreErr error

	if accountID.ValueString() != "" {
		ruleset, sempahoreErr = r.client.GetAccountRulesetPhase(ctx, accountID.ValueString(), rulesetPhase)
	} else {
		ruleset, sempahoreErr = r.client.GetZoneRulesetPhase(ctx, zoneID.ValueString(), rulesetPhase)
	}

	if len(ruleset.Rules) > 0 {
		deleteRulesetURL := accountLevelRulesetDeleteURL
		if accountID.ValueString() == "" {
			deleteRulesetURL = zoneLevelRulesetDeleteURL
		}
		resp.Diagnostics.AddError(
			fmt.Sprintf("failed to create ruleset %q", rulesetPhase),
			fmt.Sprintf(duplicateRulesetError, deleteRulesetURL),
		)
		return
	}

	rulesetName := data.Name.ValueString()
	rulesetDescription := data.Description.ValueString()
	rulesetKind := data.Kind.ValueString()
	rs := cloudflare.Ruleset{
		Name:        rulesetName,
		Description: rulesetDescription,
		Kind:        rulesetKind,
		Phase:       rulesetPhase,
	}

	rulesetData := data.toRuleset()

	if len(rulesetData.Rules) > 0 {
		rs.Rules = rulesetData.Rules
	}

	if sempahoreErr == nil && len(ruleset.Rules) == 0 && ruleset.Description == "" {
		tflog.Debug(ctx, "default ruleset created by the UI with empty rules found, recreating from scratch")
		var deleteRulesetErr error
		if accountID.ValueString() != "" {
			deleteRulesetErr = r.client.DeleteAccountRuleset(ctx, accountID.ValueString(), ruleset.ID)
		} else {
			deleteRulesetErr = r.client.DeleteZoneRuleset(ctx, zoneID.ValueString(), ruleset.ID)
		}

		if deleteRulesetErr != nil {
			resp.Diagnostics.AddError("failed to delete ruleset", deleteRulesetErr.Error())
			return
		}
	}

	var rulesetCreateErr error
	if accountID.ValueString() != "" {
		ruleset, rulesetCreateErr = r.client.CreateAccountRuleset(ctx, accountID.ValueString(), rs)
	} else {
		ruleset, rulesetCreateErr = r.client.CreateZoneRuleset(ctx, zoneID.ValueString(), rs)
	}

	if rulesetCreateErr != nil {
		resp.Diagnostics.AddError(fmt.Sprintf("error creating ruleset %s", rulesetName), rulesetCreateErr.Error())
		return
	}

	rulesetEntryPoint := cloudflare.Ruleset{
		Description: rulesetDescription,
		Rules:       rulesetData.Rules,
	}

	var err error
	// For "custom" rulesets, we don't send a follow up PUT it to the entrypoint
	// endpoint.
	if rulesetKind != string(cloudflare.RulesetKindCustom) {
		if accountID.ValueString() != "" {
			ruleset, err = r.client.UpdateAccountRulesetPhase(ctx, accountID.ValueString(), rulesetPhase, rulesetEntryPoint)
		} else {
			ruleset, err = r.client.UpdateZoneRulesetPhase(ctx, zoneID.ValueString(), rulesetPhase, rulesetEntryPoint)
		}

		if err != nil {
			resp.Diagnostics.AddError(fmt.Sprintf("error updating ruleset phase entrypoint %s", rulesetName), err.Error())
			return
		}
	}

	if zoneID.ValueString() != "" {
		data.ZoneID = types.StringValue(zoneID.ValueString())
	} else {
		data.AccountID = types.StringValue(accountID.ValueString())
	}

	data.ID = types.StringValue(ruleset.ID)

	diags = resp.State.Set(ctx, toRulesetResourceModel(data.ZoneID, data.AccountID, ruleset))
	resp.Diagnostics.Append(diags...)
}

func (r *RulesetResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *RulesetResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	accountID := data.AccountID
	zoneID := data.ZoneID
	var err error
	var ruleset cloudflare.Ruleset

	if accountID.ValueString() != "" {
		ruleset, err = r.client.GetAccountRuleset(ctx, accountID.ValueString(), data.ID.ValueString())
	} else {
		ruleset, err = r.client.GetZoneRuleset(ctx, zoneID.ValueString(), data.ID.ValueString())
	}

	if err != nil {
		var notFoundError *cloudflare.NotFoundError
		if errors.As(err, &notFoundError) {
			tflog.Warn(ctx, fmt.Sprintf("removing ruleset from state because it is not present in the remote"))
			resp.State.RemoveResource(ctx)
			return
		}

		resp.Diagnostics.AddError(
			fmt.Sprintf("error reading ruleset ID %q", data.ID.ValueString()),
			err.Error(),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, toRulesetResourceModel(zoneID, accountID, ruleset))...)
}

func (r *RulesetResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var state *RulesetResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var plan *RulesetResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	accountID := plan.AccountID
	zoneID := plan.ZoneID.ValueString()

	remappedRules, e := remapPreservedRuleIDs(state, plan)
	if e != nil {
		resp.Diagnostics.AddError("failed to remap rule IDs from state", e.Error())
		return
	}

	var err error
	var rs cloudflare.Ruleset
	description := plan.Description.ValueString()
	if accountID.ValueString() != "" {
		rs, err = r.client.UpdateAccountRuleset(ctx, accountID.ValueString(), state.ID.ValueString(), description, remappedRules)
	} else {
		rs, err = r.client.UpdateZoneRuleset(ctx, zoneID, state.ID.ValueString(), description, remappedRules)
	}

	if err != nil {
		resp.Diagnostics.AddError(fmt.Sprintf("error updating ruleset with ID %q", state.ID.ValueString()), err.Error())
		return
	}

	plan.ID = types.StringValue(rs.ID)

	resp.Diagnostics.Append(resp.State.Set(ctx, toRulesetResourceModel(state.ZoneID, state.AccountID, rs))...)
}

func (r *RulesetResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *RulesetResourceModel

	diags := req.State.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)

	accountID := data.AccountID
	zoneID := data.ZoneID

	var err error

	if accountID.ValueString() != "" {
		err = r.client.DeleteAccountRuleset(ctx, accountID.ValueString(), data.ID.ValueString())
	} else {
		err = r.client.DeleteZoneRuleset(ctx, zoneID.ValueString(), data.ID.ValueString())
	}

	if err != nil {
		resp.Diagnostics.AddError(fmt.Sprintf("error deleting ruleset with ID %q", data.ID.ValueString()), err.Error())
		return
	}
}

func (r *RulesetResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, "/")
	resourceLevel, resourceIdentifier, rulesetID := idParts[0], idParts[1], idParts[2]

	if len(idParts) != 3 || resourceLevel == "" || resourceIdentifier == "" || rulesetID == "" {
		resp.Diagnostics.AddError(
			"invalid import identifier",
			fmt.Sprintf("expected import identifier to be resourceLevel/resourceIdentifier/rulesetID. got: %q", req.ID),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), rulesetID)...)
	if resourceLevel == "zone" {
		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("zone_id"), resourceIdentifier)...)
	} else {
		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("account_id"), resourceIdentifier)...)
	}
}

// toRulesetResourceModel is a method that takes the API payload
// (`cloudflare.Ruleset`) and builds the state representation in the form of
// `*RulesetResourceModel`.
//
// The reverse of this method is `toRuleset` which handles building an API
// representation using the proposed config.
func toRulesetResourceModel(zoneID, accountID basetypes.StringValue, in cloudflare.Ruleset) *RulesetResourceModel {
	data := RulesetResourceModel{
		ID:          types.StringValue(in.ID),
		Description: types.StringValue(in.Description),
		Name:        types.StringValue(in.Name),
		Kind:        types.StringValue(in.Kind),
		Phase:       types.StringValue(in.Phase),
	}

	var ruleState []*RulesModel
	for _, ruleResponse := range in.Rules {
		var rule RulesModel

		rule.ID = flatteners.String(ruleResponse.ID)
		rule.Action = flatteners.String(ruleResponse.Action)
		rule.Expression = flatteners.String(ruleResponse.Expression)
		rule.Description = types.StringValue(ruleResponse.Description)

		rule.Enabled = flatteners.Bool(ruleResponse.Enabled)
		rule.Version = flatteners.String(cloudflare.String(ruleResponse.Version))

		// action_parameters
		if !reflect.ValueOf(ruleResponse.ActionParameters).IsNil() {
			rule.ActionParameters = append(rule.ActionParameters, &ActionParametersModel{
				AutomaticHTTPSRewrites:  flatteners.Bool(ruleResponse.ActionParameters.AutomaticHTTPSRewrites),
				BIC:                     flatteners.Bool(ruleResponse.ActionParameters.BrowserIntegrityCheck),
				Cache:                   flatteners.Bool(ruleResponse.ActionParameters.Cache),
				Content:                 flatteners.String(ruleResponse.ActionParameters.Content),
				ContentType:             flatteners.String(ruleResponse.ActionParameters.ContentType),
				DisableApps:             flatteners.Bool(ruleResponse.ActionParameters.DisableApps),
				DisableRailgun:          flatteners.Bool(ruleResponse.ActionParameters.DisableRailgun),
				DisableZaraz:            flatteners.Bool(ruleResponse.ActionParameters.DisableZaraz),
				EmailObfuscation:        flatteners.Bool(ruleResponse.ActionParameters.EmailObfuscation),
				HostHeader:              flatteners.String(ruleResponse.ActionParameters.HostHeader),
				HotlinkProtection:       flatteners.Bool(ruleResponse.ActionParameters.HotLinkProtection),
				ID:                      flatteners.String(ruleResponse.ActionParameters.ID),
				Increment:               flatteners.Int64(int64(ruleResponse.ActionParameters.Increment)),
				Mirage:                  flatteners.Bool(ruleResponse.ActionParameters.Mirage),
				OpportunisticEncryption: flatteners.Bool(ruleResponse.ActionParameters.OpportunisticEncryption),
				RocketLoader:            flatteners.Bool(ruleResponse.ActionParameters.RocketLoader),
				Ruleset:                 flatteners.String(ruleResponse.ActionParameters.Ruleset),
				ServerSideExcludes:      flatteners.Bool(ruleResponse.ActionParameters.ServerSideExcludes),
				StatusCode:              flatteners.Int64(int64(ruleResponse.ActionParameters.StatusCode)),
				SXG:                     flatteners.Bool(ruleResponse.ActionParameters.SXG),
				OriginErrorPagePassthru: flatteners.Bool(ruleResponse.ActionParameters.OriginErrorPagePassthru),
				RespectStrongEtags:      flatteners.Bool(ruleResponse.ActionParameters.RespectStrongETags),
				Version:                 flatteners.String(cloudflare.String(ruleResponse.ActionParameters.Version)),
			})

			if !reflect.ValueOf(ruleResponse.ActionParameters.Polish).IsNil() {
				rule.ActionParameters[0].Polish = flatteners.String(ruleResponse.ActionParameters.Polish.String())
			}

			if !reflect.ValueOf(ruleResponse.ActionParameters.SecurityLevel).IsNil() {
				rule.ActionParameters[0].SecurityLevel = flatteners.String(ruleResponse.ActionParameters.SecurityLevel.String())
			}

			if !reflect.ValueOf(ruleResponse.ActionParameters.SSL).IsNil() {
				rule.ActionParameters[0].SSL = flatteners.String(ruleResponse.ActionParameters.SSL.String())
			}

			var phases []attr.Value
			for _, s := range ruleResponse.ActionParameters.Phases {
				phases = append(phases, types.StringValue(s))
			}
			rule.ActionParameters[0].Phases = flatteners.StringSet(phases)

			var products []attr.Value
			for _, s := range ruleResponse.ActionParameters.Products {
				products = append(products, types.StringValue(s))
			}
			rule.ActionParameters[0].Products = flatteners.StringSet(products)

			var cookieFields []attr.Value
			for _, s := range ruleResponse.ActionParameters.CookieFields {
				cookieFields = append(cookieFields, types.StringValue(s.Name))
			}
			rule.ActionParameters[0].CookieFields = flatteners.StringSet(cookieFields)

			var rulesets []attr.Value
			for _, s := range ruleResponse.ActionParameters.Rulesets {
				rulesets = append(rulesets, types.StringValue(s))
			}
			rule.ActionParameters[0].Rulesets = flatteners.StringSet(rulesets)

			var requestFields []attr.Value
			for _, s := range ruleResponse.ActionParameters.RequestFields {
				requestFields = append(requestFields, types.StringValue(s.Name))
			}
			rule.ActionParameters[0].RequestFields = flatteners.StringSet(requestFields)

			var responseFields []attr.Value
			for _, s := range ruleResponse.ActionParameters.ResponseFields {
				responseFields = append(responseFields, types.StringValue(s.Name))
			}
			rule.ActionParameters[0].ResponseFields = flatteners.StringSet(responseFields)

			if !reflect.ValueOf(ruleResponse.ActionParameters.Overrides).IsNil() {
				var overrides []*ActionParameterOverridesModel

				var ruleOverrides []*ActionParameterOverridesRulesModel
				for _, r := range ruleResponse.ActionParameters.Overrides.Rules {
					ruleOverrides = append(ruleOverrides, &ActionParameterOverridesRulesModel{
						ID:               flatteners.String(r.ID),
						Action:           flatteners.String(r.Action),
						ScoreThreshold:   flatteners.Int64(int64(r.ScoreThreshold)),
						Enabled:          flatteners.Bool(r.Enabled),
						SensitivityLevel: flatteners.String(r.SensitivityLevel),
					})
				}

				var categoryOverrides []*ActionParameterOverridesCategoriesModel
				for _, c := range ruleResponse.ActionParameters.Overrides.Categories {
					categoryOverrides = append(categoryOverrides, &ActionParameterOverridesCategoriesModel{
						Category: flatteners.String(c.Category),
						Action:   flatteners.String(c.Action),
						Enabled:  flatteners.Bool(c.Enabled),
					})
				}

				override := &ActionParameterOverridesModel{
					SensitivityLevel: flatteners.String(ruleResponse.ActionParameters.Overrides.SensitivityLevel),
					Action:           flatteners.String(ruleResponse.ActionParameters.Overrides.Action),
					Enabled:          flatteners.Bool(ruleResponse.ActionParameters.Overrides.Enabled),
				}

				if len(ruleOverrides) > 0 {
					override.Rules = ruleOverrides
				}

				if len(categoryOverrides) > 0 {
					override.Categories = categoryOverrides
				}

				overrides = append(overrides, override)

				rule.ActionParameters[0].Overrides = overrides
			}

			if ruleResponse.ActionParameters.Rules != nil {
				result := make(map[string]basetypes.StringValue, 0)
				for k, v := range ruleResponse.ActionParameters.Rules {
					result[k] = types.StringValue(strings.Join(v, ","))
				}
				rule.ActionParameters[0].Rules = result
			}

			if ruleResponse.ActionParameters.Response != nil {
				rule.ActionParameters[0].Response = []*ActionParameterResponseModel{{
					StatusCode:  flatteners.Int64(int64(ruleResponse.ActionParameters.Response.StatusCode)),
					ContentType: flatteners.String(ruleResponse.ActionParameters.Response.ContentType),
					Content:     flatteners.String(ruleResponse.ActionParameters.Response.Content),
				}}
			}

			if ruleResponse.ActionParameters.AutoMinify != nil {
				rule.ActionParameters[0].AutoMinify = []*ActionParameterAutoMinifyModel{{
					HTML: flatteners.Bool(&ruleResponse.ActionParameters.AutoMinify.HTML),
					CSS:  flatteners.Bool(&ruleResponse.ActionParameters.AutoMinify.CSS),
					JS:   flatteners.Bool(&ruleResponse.ActionParameters.AutoMinify.JS),
				}}
			}

			if ruleResponse.ActionParameters.MatchedData != nil && ruleResponse.ActionParameters.MatchedData.PublicKey != "" {
				rule.ActionParameters[0].MatchedData = []*ActionParametersMatchedDataModel{{
					PublicKey: types.StringValue(ruleResponse.ActionParameters.MatchedData.PublicKey),
				}}
			}

			if ruleResponse.ActionParameters.BrowserTTL != nil {
				var defaultVal basetypes.Int64Value
				if cloudflare.Uint(ruleResponse.ActionParameters.BrowserTTL.Default) > 0 {
					defaultVal = types.Int64Value(int64(cloudflare.Uint(ruleResponse.ActionParameters.BrowserTTL.Default)))
				}

				rule.ActionParameters[0].BrowserTTL = []*ActionParameterBrowserTTLModel{{
					Mode:    types.StringValue(ruleResponse.ActionParameters.BrowserTTL.Mode),
					Default: defaultVal,
				}}
			}

			if ruleResponse.ActionParameters.CacheKey != nil {
				rule.ActionParameters[0].CacheKey = []*ActionParameterCacheKeyModel{{
					CacheByDeviceType:       flatteners.Bool(ruleResponse.ActionParameters.CacheKey.CacheByDeviceType),
					CacheDeceptionArmor:     flatteners.Bool(ruleResponse.ActionParameters.CacheKey.CacheDeceptionArmor),
					IgnoreQueryStringsOrder: flatteners.Bool(ruleResponse.ActionParameters.CacheKey.IgnoreQueryStringsOrder),
				}}

				if ruleResponse.ActionParameters.CacheKey.CustomKey != nil {
					key := &ActionParameterCacheKeyCustomKeyModel{}

					if ruleResponse.ActionParameters.CacheKey.CustomKey.User != nil {
						key.User = []*ActionParameterCacheKeyCustomKeyUserModel{{
							DeviceType: flatteners.Bool(ruleResponse.ActionParameters.CacheKey.CustomKey.User.DeviceType),
							Geo:        flatteners.Bool(ruleResponse.ActionParameters.CacheKey.CustomKey.User.Geo),
							Lang:       flatteners.Bool(ruleResponse.ActionParameters.CacheKey.CustomKey.User.Lang),
						}}
					}

					if ruleResponse.ActionParameters.CacheKey.CustomKey.Host != nil {
						key.Host = []*ActionParameterCacheKeyCustomKeyHostModel{{
							Resolved: flatteners.Bool(ruleResponse.ActionParameters.CacheKey.CustomKey.Host.Resolved),
						}}
					}

					if ruleResponse.ActionParameters.CacheKey.CustomKey.Cookie != nil {
						include, _ := basetypes.NewSetValueFrom(context.Background(), types.StringType, ruleResponse.ActionParameters.CacheKey.CustomKey.Cookie.Include)
						checkPresence, _ := basetypes.NewSetValueFrom(context.Background(), types.StringType, ruleResponse.ActionParameters.CacheKey.CustomKey.Cookie.CheckPresence)
						key.Cookie = []*ActionParameterCacheKeyCustomKeyCookieModel{{
							Include:       include,
							CheckPresence: checkPresence,
						}}
					}

					if ruleResponse.ActionParameters.CacheKey.CustomKey.Header != nil {
						include, _ := basetypes.NewSetValueFrom(context.Background(), types.StringType, ruleResponse.ActionParameters.CacheKey.CustomKey.Header.Include)
						checkPresence, _ := basetypes.NewSetValueFrom(context.Background(), types.StringType, ruleResponse.ActionParameters.CacheKey.CustomKey.Header.CheckPresence)
						if len(include.Elements()) > 0 || len(checkPresence.Elements()) > 0 {
							key.Header = []*ActionParameterCacheKeyCustomKeyHeaderModel{{
								Include:       include,
								CheckPresence: checkPresence,
								ExcludeOrigin: flatteners.Bool(ruleResponse.ActionParameters.CacheKey.CustomKey.Header.ExcludeOrigin),
							}}
						}
					}

					if ruleResponse.ActionParameters.CacheKey.CustomKey.Query != nil {
						include, _ := basetypes.NewSetValueFrom(context.Background(), types.StringType, ruleResponse.ActionParameters.CacheKey.CustomKey.Query.Include)
						exclude, _ := basetypes.NewSetValueFrom(context.Background(), types.StringType, ruleResponse.ActionParameters.CacheKey.CustomKey.Query.Exclude)

						if ruleResponse.ActionParameters.CacheKey.CustomKey.Query.Include != nil && ruleResponse.ActionParameters.CacheKey.CustomKey.Query.Include.All {
							include, _ = basetypes.NewSetValueFrom(context.Background(), types.StringType, []string{"*"})
						}

						if ruleResponse.ActionParameters.CacheKey.CustomKey.Query.Exclude != nil && ruleResponse.ActionParameters.CacheKey.CustomKey.Query.Exclude.All {
							exclude, _ = basetypes.NewSetValueFrom(context.Background(), types.StringType, []string{"*"})
						}

						key.QueryString = []*ActionParameterCacheKeyCustomKeyQueryStringModel{{
							Include: include,
							Exclude: exclude,
						}}
					}

					rule.ActionParameters[0].CacheKey[0].CustomKey = []*ActionParameterCacheKeyCustomKeyModel{key}
				}
			}

			if ruleResponse.ActionParameters.EdgeTTL != nil {
				var defaultVal basetypes.Int64Value
				if cloudflare.Uint(ruleResponse.ActionParameters.EdgeTTL.Default) > 0 {
					defaultVal = types.Int64Value(int64(cloudflare.Uint(ruleResponse.ActionParameters.EdgeTTL.Default)))
				}

				rule.ActionParameters[0].EdgeTTL = []*ActionParameterEdgeTTLModel{{
					Mode:    types.StringValue(ruleResponse.ActionParameters.EdgeTTL.Mode),
					Default: defaultVal,
				}}

				var statusCodeTTLs []*ActionParameterEdgeTTLStatusCodeTTLModel
				for _, sct := range ruleResponse.ActionParameters.EdgeTTL.StatusCodeTTL {
					var sctrange []*ActionParameterEdgeTTLStatusCodeTTLStatusCodeRangeModel

					if sct.StatusCodeRange != nil {
						sctrange = append(sctrange, &ActionParameterEdgeTTLStatusCodeTTLStatusCodeRangeModel{
							To:   flatteners.Int64(int64(cloudflare.Uint(sct.StatusCodeRange.To))),
							From: flatteners.Int64(int64(cloudflare.Uint(sct.StatusCodeRange.From))),
						})
					}
					statusCodeTTLs = append(statusCodeTTLs, &ActionParameterEdgeTTLStatusCodeTTLModel{
						StatusCode:      flatteners.Int64(int64(cloudflare.Uint(sct.StatusCodeValue))),
						Value:           flatteners.Int64(int64(cloudflare.Int(sct.Value))),
						StatusCodeRange: sctrange,
					})
				}
				rule.ActionParameters[0].EdgeTTL[0].StatusCodeTTL = statusCodeTTLs
			}

			if ruleResponse.ActionParameters.ServeStale != nil {
				rule.ActionParameters[0].ServeStale = []*ActionParameterServeStaleModel{{
					DisableStaleWhileUpdating: types.BoolValue(*ruleResponse.ActionParameters.ServeStale.DisableStaleWhileUpdating),
				}}
			}

			if ruleResponse.ActionParameters.FromList != nil {
				rule.ActionParameters[0].FromList = []*ActionParameterFromListModel{{
					Name: types.StringValue(ruleResponse.ActionParameters.FromList.Name),
					Key:  types.StringValue(ruleResponse.ActionParameters.FromList.Key),
				}}
			}

			if ruleResponse.ActionParameters.Origin != nil {
				rule.ActionParameters[0].Origin = []*ActionParameterOriginModel{{
					Host: types.StringValue(ruleResponse.ActionParameters.Origin.Host),
					Port: types.Int64Value(int64(ruleResponse.ActionParameters.Origin.Port)),
				}}
			}

			if ruleResponse.ActionParameters.SNI != nil && ruleResponse.ActionParameters.SNI.Value != "" {
				rule.ActionParameters[0].SNI = []*ActionParameterSNIModel{{
					Value: types.StringValue(ruleResponse.ActionParameters.SNI.Value),
				}}
			}

			if ruleResponse.ActionParameters.URI != nil {
				rule.ActionParameters[0].URI = []*ActionParametersURIModel{{
					Origin: flatteners.Bool(ruleResponse.ActionParameters.URI.Origin),
				}}

				if ruleResponse.ActionParameters.URI.Path != nil {
					rule.ActionParameters[0].URI[0].Path = []*ActionParametersURIPartModel{{
						Value:      flatteners.String(ruleResponse.ActionParameters.URI.Path.Value),
						Expression: flatteners.String(ruleResponse.ActionParameters.URI.Path.Expression),
					}}
				}

				if ruleResponse.ActionParameters.URI.Query != nil {
					if ruleResponse.ActionParameters.URI.Query.Expression != "" {
						rule.ActionParameters[0].URI[0].Query = []*ActionParametersURIPartModel{{
							Value:      types.StringNull(),
							Expression: flatteners.String(ruleResponse.ActionParameters.URI.Query.Expression),
						}}
					} else {
						rule.ActionParameters[0].URI[0].Query = []*ActionParametersURIPartModel{{
							Value:      types.StringValue(cloudflare.String(ruleResponse.ActionParameters.URI.Query.Value)),
							Expression: flatteners.String(ruleResponse.ActionParameters.URI.Query.Expression),
						}}
					}
				}
			}

			if ruleResponse.ActionParameters.Headers != nil {
				sortedHeaders := make([]string, 0, len(ruleResponse.ActionParameters.Headers))
				for k := range ruleResponse.ActionParameters.Headers {
					sortedHeaders = append(sortedHeaders, k)
				}
				sort.Strings(sortedHeaders)

				var headers []*ActionParametersHeadersModel
				for _, name := range sortedHeaders {
					headers = append(headers, &ActionParametersHeadersModel{
						Name:       types.StringValue(name),
						Value:      flatteners.String(ruleResponse.ActionParameters.Headers[name].Value),
						Expression: flatteners.String(ruleResponse.ActionParameters.Headers[name].Expression),
						Operation:  flatteners.String(ruleResponse.ActionParameters.Headers[name].Operation),
					})
				}
				rule.ActionParameters[0].Headers = headers
			}

			if ruleResponse.ActionParameters.FromValue != nil {
				rule.ActionParameters[0].FromValue = []*ActionParameterFromValueModel{{
					StatusCode:          flatteners.Int64(int64(ruleResponse.ActionParameters.FromValue.StatusCode)),
					PreserveQueryString: flatteners.Bool(&ruleResponse.ActionParameters.FromValue.PreserveQueryString),
					TargetURL: []*ActionParameterFromValueTargetURLModel{{
						Value:      flatteners.String(ruleResponse.ActionParameters.FromValue.TargetURL.Value),
						Expression: flatteners.String(ruleResponse.ActionParameters.FromValue.TargetURL.Expression),
					}},
				}}
			}
		}

		// ratelimit
		if !reflect.ValueOf(ruleResponse.RateLimit).IsNil() {
			var rlCharacteristicsKeys []attr.Value
			for _, s := range ruleResponse.RateLimit.Characteristics {
				rlCharacteristicsKeys = append(rlCharacteristicsKeys, types.StringValue(s))
			}

			rule.Ratelimit = append(rule.Ratelimit, &RatelimitModel{
				Characteristics:         types.SetValueMust(types.StringType, rlCharacteristicsKeys),
				Period:                  flatteners.Int64(int64(ruleResponse.RateLimit.Period)),
				RequestsPerPeriod:       flatteners.Int64(int64(ruleResponse.RateLimit.RequestsPerPeriod)),
				RequestsToOrigin:        flatteners.Bool(cloudflare.BoolPtr(ruleResponse.RateLimit.RequestsToOrigin)),
				MitigationTimeout:       flatteners.Int64(int64(ruleResponse.RateLimit.MitigationTimeout)),
				ScorePerPeriod:          flatteners.Int64(int64(ruleResponse.RateLimit.ScorePerPeriod)),
				ScoreResponseHeaderName: flatteners.String(ruleResponse.RateLimit.ScoreResponseHeaderName),
				CountingExpression:      flatteners.String(ruleResponse.RateLimit.CountingExpression),
			})
		}

		// exposed credential check
		if !reflect.ValueOf(ruleResponse.ExposedCredentialCheck).IsNil() {
			rule.ExposedCredentialCheck = append(rule.ExposedCredentialCheck, &ExposedCredentialCheckModel{
				UsernameExpression: types.StringValue(ruleResponse.ExposedCredentialCheck.UsernameExpression),
				PasswordExpression: types.StringValue(ruleResponse.ExposedCredentialCheck.PasswordExpression),
			})
		}

		// logging
		if !reflect.ValueOf(ruleResponse.Logging).IsNil() {
			rule.Logging = append(rule.Logging, &LoggingModel{Enabled: types.BoolValue(*ruleResponse.Logging.Enabled)})
		}

		ruleState = append(ruleState, &rule)
	}

	data.Rules = ruleState

	data.ZoneID = flatteners.String(zoneID.ValueString())
	data.AccountID = flatteners.String(accountID.ValueString())

	return &data
}

// toRuleset is a method that takes the proposed configuration changes
// (`*RulesetResourceModel`) and builds the API representation in the form of
// `cloudflare.Ruleset`.
//
// The reverse of this method is `toRulesetResourceModel` which handles building
// a state representation using the API response.
func (r *RulesetResourceModel) toRuleset() cloudflare.Ruleset {
	var rs cloudflare.Ruleset
	var rules []cloudflare.RulesetRule

	rs.ID = r.ID.ValueString()
	for _, rule := range r.Rules {
		rules = append(rules, rule.toRulesetRule())
	}

	rs.Rules = rules

	return rs
}

// toRulesetRule takes a state representation of a Ruleset Rule and transforms
// it into an API representation.
func (r *RulesModel) toRulesetRule() cloudflare.RulesetRule {
	rr := cloudflare.RulesetRule{
		Action:      r.Action.ValueString(),
		Expression:  r.Expression.ValueString(),
		Description: r.Description.ValueString(),
	}

	if !r.ID.IsNull() {
		rr.ID = r.ID.ValueString()
	}

	if !r.Enabled.IsNull() {
		rr.Enabled = cloudflare.BoolPtr(r.Enabled.ValueBool())
	}

	if !r.Ref.IsNull() {
		rr.Ref = r.Ref.ValueString()
	}

	if !r.Version.IsNull() {
		if r.Version.ValueString() != "" {
			rr.Version = cloudflare.StringPtr(r.Version.ValueString())
		}
	}

	for _, ap := range r.ActionParameters {
		rr.ActionParameters = &cloudflare.RulesetRuleActionParameters{}

		if len(ap.Rules) > 0 {
			ruleMap := map[string][]string{}
			for key, ruleIDs := range ap.Rules {
				s := strings.Split(ruleIDs.ValueString(), ",")
				ruleMap[key] = s
				rr.ActionParameters.Rules = ruleMap
			}
		}

		if len(expanders.StringSet(ap.Phases)) > 0 {
			rr.ActionParameters.Phases = expanders.StringSet(ap.Phases)
		}

		if len(expanders.StringSet(ap.Products)) > 0 {
			rr.ActionParameters.Products = expanders.StringSet(ap.Products)
		}

		if len(expanders.StringSet(ap.Rulesets)) > 0 {
			rr.ActionParameters.Rulesets = expanders.StringSet(ap.Rulesets)
		}

		if !ap.ID.IsNull() {
			rr.ActionParameters.ID = ap.ID.ValueString()
		}

		if !ap.Content.IsNull() {
			rr.ActionParameters.Content = ap.Content.ValueString()
		}

		if !ap.ContentType.IsNull() {
			rr.ActionParameters.ContentType = ap.ContentType.ValueString()
		}

		if !ap.HostHeader.IsNull() {
			rr.ActionParameters.HostHeader = ap.HostHeader.ValueString()
		}

		if !ap.Ruleset.IsNull() {
			rr.ActionParameters.Ruleset = ap.Ruleset.ValueString()
		}

		if !ap.Version.IsNull() {
			if ap.Version.ValueString() != "" {
				rr.ActionParameters.Version = cloudflare.StringPtr(ap.Version.ValueString())
			}
		}

		if !ap.Increment.IsNull() {
			rr.ActionParameters.Increment = int(ap.Increment.ValueInt64())
		}

		if !ap.StatusCode.IsNull() {
			rr.ActionParameters.StatusCode = uint16(ap.StatusCode.ValueInt64())
		}

		if !ap.AutomaticHTTPSRewrites.IsNull() {
			rr.ActionParameters.AutomaticHTTPSRewrites = cloudflare.BoolPtr(ap.AutomaticHTTPSRewrites.ValueBool())
		}

		if !ap.BIC.IsNull() {
			rr.ActionParameters.BrowserIntegrityCheck = cloudflare.BoolPtr(ap.BIC.ValueBool())
		}

		if !ap.Cache.IsNull() {
			rr.ActionParameters.Cache = cloudflare.BoolPtr(ap.Cache.ValueBool())
		}

		if !ap.DisableApps.IsNull() {
			rr.ActionParameters.DisableApps = cloudflare.BoolPtr(ap.DisableApps.ValueBool())
		}

		if !ap.DisableRailgun.IsNull() {
			rr.ActionParameters.DisableRailgun = cloudflare.BoolPtr(ap.DisableRailgun.ValueBool())
		}

		if !ap.DisableZaraz.IsNull() {
			rr.ActionParameters.DisableZaraz = cloudflare.BoolPtr(ap.DisableZaraz.ValueBool())
		}

		if !ap.EmailObfuscation.IsNull() {
			rr.ActionParameters.EmailObfuscation = cloudflare.BoolPtr(ap.EmailObfuscation.ValueBool())
		}

		if !ap.HotlinkProtection.IsNull() {
			rr.ActionParameters.HotLinkProtection = cloudflare.BoolPtr(ap.HotlinkProtection.ValueBool())
		}

		if !ap.Mirage.IsNull() {
			rr.ActionParameters.Mirage = cloudflare.BoolPtr(ap.Mirage.ValueBool())
		}

		if !ap.OpportunisticEncryption.IsNull() {
			rr.ActionParameters.OpportunisticEncryption = cloudflare.BoolPtr(ap.OpportunisticEncryption.ValueBool())
		}

		if !ap.RocketLoader.IsNull() {
			rr.ActionParameters.RocketLoader = cloudflare.BoolPtr(ap.RocketLoader.ValueBool())
		}

		if !ap.ServerSideExcludes.IsNull() {
			rr.ActionParameters.ServerSideExcludes = cloudflare.BoolPtr(ap.ServerSideExcludes.ValueBool())
		}

		if !ap.SXG.IsNull() {
			rr.ActionParameters.SXG = cloudflare.BoolPtr(ap.SXG.ValueBool())
		}

		if !ap.Polish.IsNull() {
			polish, _ := cloudflare.PolishFromString(ap.Polish.ValueString())
			rr.ActionParameters.Polish = polish
		}

		if !ap.SecurityLevel.IsNull() {
			securityLevel, _ := cloudflare.SecurityLevelFromString(ap.SecurityLevel.ValueString())
			rr.ActionParameters.SecurityLevel = securityLevel
		}

		if !ap.SSL.IsNull() {
			ssl, _ := cloudflare.SSLFromString(ap.SSL.ValueString())
			rr.ActionParameters.SSL = ssl
		}

		if !ap.OriginErrorPagePassthru.IsNull() {
			rr.ActionParameters.OriginErrorPagePassthru = cloudflare.BoolPtr(ap.OriginErrorPagePassthru.ValueBool())
		}

		if !ap.RespectStrongEtags.IsNull() {
			rr.ActionParameters.RespectStrongETags = cloudflare.BoolPtr(ap.RespectStrongEtags.ValueBool())
		}

		if len(ap.Overrides) > 0 {
			var overrides cloudflare.RulesetRuleActionParametersOverrides
			var ruleOverrides []cloudflare.RulesetRuleActionParametersRules
			var categoryOverrides []cloudflare.RulesetRuleActionParametersCategories

			for _, ro := range ap.Overrides[0].Rules {
				rule := cloudflare.RulesetRuleActionParametersRules{
					ID:               ro.ID.ValueString(),
					Action:           ro.Action.ValueString(),
					SensitivityLevel: ro.SensitivityLevel.ValueString(),
				}

				if !ro.ScoreThreshold.IsNull() {
					rule.ScoreThreshold = int(ro.ScoreThreshold.ValueInt64())
				}

				if !ro.Enabled.IsNull() {
					rule.Enabled = cloudflare.BoolPtr(ro.Enabled.ValueBool())
				}

				ruleOverrides = append(ruleOverrides, rule)
			}
			overrides.Rules = ruleOverrides

			for _, co := range ap.Overrides[0].Categories {
				category := cloudflare.RulesetRuleActionParametersCategories{
					Category: co.Category.ValueString(),
				}

				if !co.Action.IsNull() {
					category.Action = co.Action.ValueString()
				}

				if !co.Enabled.IsNull() {
					category.Enabled = cloudflare.BoolPtr(co.Enabled.ValueBool())
				}

				categoryOverrides = append(categoryOverrides, category)
			}
			overrides.Categories = categoryOverrides

			if !ap.Overrides[0].Action.IsNull() {
				overrides.Action = ap.Overrides[0].Action.ValueString()
			}

			if !ap.Overrides[0].SensitivityLevel.IsNull() {
				overrides.SensitivityLevel = ap.Overrides[0].SensitivityLevel.ValueString()
			}

			if !ap.Overrides[0].Enabled.IsNull() {
				overrides.Enabled = cloudflare.BoolPtr(ap.Overrides[0].Enabled.ValueBool())
			}

			rr.ActionParameters.Overrides = &overrides
		}

		if len(ap.MatchedData) > 0 {
			rr.ActionParameters.MatchedData = &cloudflare.RulesetRuleActionParametersMatchedData{
				PublicKey: ap.MatchedData[0].PublicKey.ValueString(),
			}
		}

		if len(ap.Response) > 0 {
			response := cloudflare.RulesetRuleActionParametersBlockResponse{
				ContentType: ap.Response[0].ContentType.ValueString(),
				Content:     ap.Response[0].Content.ValueString(),
				StatusCode:  uint16(ap.Response[0].StatusCode.ValueInt64()),
			}
			rr.ActionParameters.Response = &response
		}

		if len(ap.AutoMinify) > 0 {
			autominify := cloudflare.RulesetRuleActionParametersAutoMinify{
				HTML: ap.AutoMinify[0].HTML.ValueBool(),
				CSS:  ap.AutoMinify[0].CSS.ValueBool(),
				JS:   ap.AutoMinify[0].JS.ValueBool(),
			}
			rr.ActionParameters.AutoMinify = &autominify
		}

		if len(ap.BrowserTTL) > 0 {
			browserTTL := cloudflare.RulesetRuleActionParametersBrowserTTL{
				Mode: ap.BrowserTTL[0].Mode.ValueString(),
			}

			if !ap.BrowserTTL[0].Default.IsNull() {
				browserTTL.Default = cloudflare.UintPtr(uint(ap.BrowserTTL[0].Default.ValueInt64()))
			}

			rr.ActionParameters.BrowserTTL = &browserTTL
		}

		if len(ap.ServeStale) > 0 && !ap.ServeStale[0].DisableStaleWhileUpdating.IsNull() {
			rr.ActionParameters.ServeStale = &cloudflare.RulesetRuleActionParametersServeStale{
				DisableStaleWhileUpdating: cloudflare.BoolPtr(ap.ServeStale[0].DisableStaleWhileUpdating.ValueBool()),
			}
		}

		if len(ap.FromList) > 0 {
			fromList := cloudflare.RulesetRuleActionParametersFromList{
				Name: ap.FromList[0].Name.ValueString(),
				Key:  ap.FromList[0].Key.ValueString(),
			}
			rr.ActionParameters.FromList = &fromList
		}

		if len(ap.Origin) > 0 {
			origin := cloudflare.RulesetRuleActionParametersOrigin{
				Host: ap.Origin[0].Host.ValueString(),
				Port: uint16(ap.Origin[0].Port.ValueInt64()),
			}
			rr.ActionParameters.Origin = &origin
		}

		if len(ap.SNI) > 0 {
			sni := cloudflare.RulesetRuleActionParametersSni{
				Value: ap.SNI[0].Value.ValueString(),
			}
			rr.ActionParameters.SNI = &sni
		}

		if len(ap.URI) > 0 {
			uri := &cloudflare.RulesetRuleActionParametersURI{}

			if !ap.URI[0].Origin.IsNull() {
				uri.Origin = cloudflare.BoolPtr(ap.URI[0].Origin.ValueBool())
			}

			if len(ap.URI[0].Path) > 0 {
				uri.Path = &cloudflare.RulesetRuleActionParametersURIPath{
					Expression: ap.URI[0].Path[0].Expression.ValueString(),
					Value:      ap.URI[0].Path[0].Value.ValueString(),
				}
			}

			if len(ap.URI[0].Query) > 0 {
				if ap.URI[0].Query[0].Expression.ValueString() != "" {
					uri.Query = &cloudflare.RulesetRuleActionParametersURIQuery{
						Expression: ap.URI[0].Query[0].Expression.ValueString(),
					}
				} else {
					uri.Query = &cloudflare.RulesetRuleActionParametersURIQuery{
						Value: cloudflare.StringPtr(ap.URI[0].Query[0].Value.ValueString()),
					}
				}
			}

			rr.ActionParameters.URI = uri
		}

		if len(ap.Headers) > 0 {
			headers := map[string]cloudflare.RulesetRuleActionParametersHTTPHeader{}
			for _, header := range ap.Headers {
				headers[header.Name.ValueString()] = cloudflare.RulesetRuleActionParametersHTTPHeader{
					Operation:  header.Operation.ValueString(),
					Value:      header.Value.ValueString(),
					Expression: header.Expression.ValueString(),
				}
			}

			rr.ActionParameters.Headers = headers
		}

		if len(ap.CacheKey) > 0 {
			key := cloudflare.RulesetRuleActionParametersCacheKey{}

			if !ap.CacheKey[0].IgnoreQueryStringsOrder.IsNull() {
				key.IgnoreQueryStringsOrder = cloudflare.BoolPtr(ap.CacheKey[0].IgnoreQueryStringsOrder.ValueBool())
			}

			if !ap.CacheKey[0].CacheByDeviceType.IsNull() {
				key.CacheByDeviceType = cloudflare.BoolPtr(ap.CacheKey[0].CacheByDeviceType.ValueBool())
			}

			if !ap.CacheKey[0].CacheDeceptionArmor.IsNull() {
				key.CacheDeceptionArmor = cloudflare.BoolPtr(ap.CacheKey[0].CacheDeceptionArmor.ValueBool())
			}

			if len(ap.CacheKey[0].CustomKey) > 0 {
				customKey := cloudflare.RulesetRuleActionParametersCustomKey{}

				if len(ap.CacheKey[0].CustomKey[0].QueryString) > 0 {
					includeQueryList := expanders.StringSet(ap.CacheKey[0].CustomKey[0].QueryString[0].Include)
					excludeQueryList := expanders.StringSet(ap.CacheKey[0].CustomKey[0].QueryString[0].Exclude)

					if len(includeQueryList) > 0 {
						if len(includeQueryList) == 1 && includeQueryList[0] == "*" {
							customKey.Query = &cloudflare.RulesetRuleActionParametersCustomKeyQuery{
								Include: &cloudflare.RulesetRuleActionParametersCustomKeyList{
									All: true,
								},
							}
						} else {
							customKey.Query = &cloudflare.RulesetRuleActionParametersCustomKeyQuery{
								Include: &cloudflare.RulesetRuleActionParametersCustomKeyList{
									List: includeQueryList,
								},
							}
						}
					}

					if len(excludeQueryList) > 0 {
						if len(excludeQueryList) == 1 && excludeQueryList[0] == "*" {
							customKey.Query = &cloudflare.RulesetRuleActionParametersCustomKeyQuery{
								Exclude: &cloudflare.RulesetRuleActionParametersCustomKeyList{
									All: true,
								},
							}
						} else {
							customKey.Query = &cloudflare.RulesetRuleActionParametersCustomKeyQuery{
								Exclude: &cloudflare.RulesetRuleActionParametersCustomKeyList{
									List: excludeQueryList,
								},
							}
						}
					}
				}

				if len(ap.CacheKey[0].CustomKey[0].Header) > 0 {
					includeQueryList := expanders.StringSet(ap.CacheKey[0].CustomKey[0].Header[0].Include)
					checkPresenceList := expanders.StringSet(basetypes.SetValue(ap.CacheKey[0].CustomKey[0].Header[0].CheckPresence))

					customKey.Header = &cloudflare.RulesetRuleActionParametersCustomKeyHeader{
						RulesetRuleActionParametersCustomKeyFields: cloudflare.RulesetRuleActionParametersCustomKeyFields{
							Include:       includeQueryList,
							CheckPresence: checkPresenceList,
						},
						ExcludeOrigin: cloudflare.BoolPtr(ap.CacheKey[0].CustomKey[0].Header[0].ExcludeOrigin.ValueBool()),
					}
				}

				if len(ap.CacheKey[0].CustomKey[0].Cookie) > 0 {
					includeQueryList := expanders.StringSet(ap.CacheKey[0].CustomKey[0].Cookie[0].Include)
					checkPresenceList := expanders.StringSet(basetypes.SetValue(ap.CacheKey[0].CustomKey[0].Cookie[0].CheckPresence))

					if len(includeQueryList) > 0 || len(checkPresenceList) > 0 {
						customKey.Cookie = &cloudflare.RulesetRuleActionParametersCustomKeyCookie{
							Include:       includeQueryList,
							CheckPresence: checkPresenceList,
						}
					}
				}

				if len(ap.CacheKey[0].CustomKey[0].User) > 0 &&
					(!ap.CacheKey[0].CustomKey[0].User[0].DeviceType.IsNull() ||
						!ap.CacheKey[0].CustomKey[0].User[0].Geo.IsNull() ||
						!ap.CacheKey[0].CustomKey[0].User[0].Lang.IsNull()) {
					customKey.User = &cloudflare.RulesetRuleActionParametersCustomKeyUser{}

					if !ap.CacheKey[0].CustomKey[0].User[0].DeviceType.IsNull() {
						customKey.User.DeviceType = cloudflare.BoolPtr(ap.CacheKey[0].CustomKey[0].User[0].DeviceType.ValueBool())
					}

					if !ap.CacheKey[0].CustomKey[0].User[0].Geo.IsNull() {
						customKey.User.Geo = cloudflare.BoolPtr(ap.CacheKey[0].CustomKey[0].User[0].Geo.ValueBool())
					}

					if !ap.CacheKey[0].CustomKey[0].User[0].Lang.IsNull() {
						customKey.User.Lang = cloudflare.BoolPtr(ap.CacheKey[0].CustomKey[0].User[0].Lang.ValueBool())
					}
				}

				if len(ap.CacheKey[0].CustomKey[0].Host) > 0 && !ap.CacheKey[0].CustomKey[0].Host[0].Resolved.IsNull() {
					customKey.Host = &cloudflare.RulesetRuleActionParametersCustomKeyHost{
						Resolved: cloudflare.BoolPtr(ap.CacheKey[0].CustomKey[0].Host[0].Resolved.ValueBool()),
					}
				}

				key.CustomKey = &customKey
			}

			rr.ActionParameters.CacheKey = &key
		}

		if len(ap.EdgeTTL) > 0 {
			edgeTTL := &cloudflare.RulesetRuleActionParametersEdgeTTL{
				Mode: ap.EdgeTTL[0].Mode.ValueString(),
			}

			if !ap.EdgeTTL[0].Default.IsNull() {
				edgeTTL.Default = cloudflare.UintPtr(uint(ap.EdgeTTL[0].Default.ValueInt64()))
			}

			var statusCodeTTLs []cloudflare.RulesetRuleActionParametersStatusCodeTTL
			for _, sct := range ap.EdgeTTL[0].StatusCodeTTL {
				config := cloudflare.RulesetRuleActionParametersStatusCodeTTL{}

				if sct.StatusCodeRange != nil {
					config.StatusCodeRange = &cloudflare.RulesetRuleActionParametersStatusCodeRange{}

					if sct.StatusCodeRange[0].From.ValueInt64() > 0 {
						config.StatusCodeRange.From = cloudflare.UintPtr(uint(sct.StatusCodeRange[0].From.ValueInt64()))
					}

					if sct.StatusCodeRange[0].To.ValueInt64() > 0 {
						config.StatusCodeRange.To = cloudflare.UintPtr(uint(sct.StatusCodeRange[0].To.ValueInt64()))
					}
				}

				if !sct.StatusCode.IsNull() {
					config.StatusCodeValue = cloudflare.UintPtr(uint(sct.StatusCode.ValueInt64()))
				}

				config.Value = cloudflare.IntPtr(int(sct.Value.ValueInt64()))
				statusCodeTTLs = append(statusCodeTTLs, config)
			}

			edgeTTL.StatusCodeTTL = statusCodeTTLs
			rr.ActionParameters.EdgeTTL = edgeTTL
		}

		if len(ap.FromValue) > 0 {
			from := &cloudflare.RulesetRuleActionParametersFromValue{}

			if !ap.FromValue[0].StatusCode.IsNull() {
				from.StatusCode = uint16(ap.FromValue[0].StatusCode.ValueInt64())
			}

			if !ap.FromValue[0].PreserveQueryString.IsNull() {
				from.PreserveQueryString = ap.FromValue[0].PreserveQueryString.ValueBool()
			}

			from.TargetURL.Expression = ap.FromValue[0].TargetURL[0].Expression.ValueString()
			from.TargetURL.Value = ap.FromValue[0].TargetURL[0].Value.ValueString()

			rr.ActionParameters.FromValue = from
		}

		apCookieFields := expanders.StringSet(ap.CookieFields)
		if len(apCookieFields) > 0 {
			for _, cookie := range apCookieFields {
				rr.ActionParameters.CookieFields = append(rr.ActionParameters.CookieFields, cloudflare.RulesetActionParametersLogCustomField{Name: cookie})
			}
		}

		apRequestFields := expanders.StringSet(ap.RequestFields)
		if len(apRequestFields) > 0 {
			for _, request := range apRequestFields {
				rr.ActionParameters.RequestFields = append(rr.ActionParameters.RequestFields, cloudflare.RulesetActionParametersLogCustomField{Name: request})
			}
		}

		apResponseFields := expanders.StringSet(ap.ResponseFields)
		if len(apResponseFields) > 0 {
			for _, request := range apResponseFields {
				rr.ActionParameters.ResponseFields = append(rr.ActionParameters.ResponseFields, cloudflare.RulesetActionParametersLogCustomField{Name: request})
			}
		}
	}

	for _, rl := range r.Ratelimit {
		rr.RateLimit = &cloudflare.RulesetRuleRateLimit{
			Characteristics:         expanders.StringSet(rl.Characteristics),
			Period:                  int(rl.Period.ValueInt64()),
			RequestsPerPeriod:       int(rl.RequestsPerPeriod.ValueInt64()),
			ScorePerPeriod:          int(rl.ScorePerPeriod.ValueInt64()),
			ScoreResponseHeaderName: rl.ScoreResponseHeaderName.ValueString(),
			MitigationTimeout:       int(rl.MitigationTimeout.ValueInt64()),
			CountingExpression:      rl.CountingExpression.ValueString(),
			RequestsToOrigin:        rl.RequestsToOrigin.ValueBool(),
		}
	}

	for _, l := range r.Logging {
		rr.Logging = &cloudflare.RulesetRuleLogging{
			Enabled: cloudflare.BoolPtr(l.Enabled.ValueBool()),
		}
	}

	for _, e := range r.ExposedCredentialCheck {
		rr.ExposedCredentialCheck = &cloudflare.RulesetRuleExposedCredentialCheck{
			UsernameExpression: e.UsernameExpression.ValueString(),
			PasswordExpression: e.PasswordExpression.ValueString(),
		}
	}

	return rr
}

// ruleIDs is a lookup table for rule IDs with two operations, add and pop. We
// use add to populate the table from the old value of rules. We use pop to look
// up the ID for the new value of a rule (and remove it from the table).
// Internally, both operations serialize the rule to JSON and use the resulting
// string as the lookup key; the ID itself is excluded from the JSON. If a
// ruleset has multiple copies of the same rule, the copies have a single lookup
// key associated with multiple IDs; we preserve order when adding and popping
// the IDs.
type ruleIDs struct {
	ids map[string][]string
}

// add stores an ID for the given rule.
func (r *ruleIDs) add(rule cloudflare.RulesetRule) error {
	if rule.ID == "" {
		// This is unexpected. We only invoke this function for the old
		// values of rules, which have their IDs populated.
		return errors.New("unable to determine ID of existing rule")
	}

	id := rule.ID
	rule.ID = ""

	// For the purposes of preserving existing rule IDs, we don't want to
	// include computed fields as a part of the key value.
	rule.Version = nil

	data, err := json.Marshal(rule)
	if err != nil {
		return err
	}

	key := string(data[:])

	r.ids[key] = append(r.ids[key], id)
	return nil
}

// pop removes an ID for the given rule and returns it. Multiple IDs are
// returned in the order they were added. If no ID was found for the rule, pop
// returns an empty string.
func (r *ruleIDs) pop(rule cloudflare.RulesetRule) (string, error) {
	rule.ID = ""

	// For the purposes of preserving existing rule IDs, we don't want to
	// include computed fields as a part of the key value.
	rule.Version = nil

	data, err := json.Marshal(rule)
	if err != nil {
		return "", err
	}

	key := string(data[:])

	ids := r.ids[key]
	if len(ids) == 0 {
		return "", nil
	}

	id, ids := ids[0], ids[1:]
	r.ids[key] = ids

	return id, nil
}

// empty returns true if the store does not contain any rule IDs.
func (r *ruleIDs) empty() bool {
	return len(r.ids) == 0
}

func newRuleIDs(rulesetRules []cloudflare.RulesetRule) (ruleIDs, error) {
	r := ruleIDs{make(map[string][]string)}

	for _, rule := range rulesetRules {
		err := r.add(rule)
		if err != nil {
			return ruleIDs{}, err
		}
	}

	return r, nil
}

func remapPreservedRuleIDs(state, plan *RulesetResourceModel) ([]cloudflare.RulesetRule, error) {
	currentRuleset := state.toRuleset()
	plannedRuleset := plan.toRuleset()

	ids, err := newRuleIDs(currentRuleset.Rules)
	if err != nil {
		return nil, err
	}

	if ids.empty() {
		// There are no rule IDs when the ruleset is first created.
		return plannedRuleset.Rules, nil
	}

	for i := range plannedRuleset.Rules {
		rule := &plannedRuleset.Rules[i]
		rule.ID, err = ids.pop(*rule)
		if err != nil {
			return nil, err
		}
	}

	return plannedRuleset.Rules, nil
}

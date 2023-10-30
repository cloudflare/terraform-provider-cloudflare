package rulesets

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"sort"
	"strings"
	"time"

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
	duplicateRulesetError = "A similar configuration with rules already exists and overwriting will have unintended consequences. If you are migrating from the Dashboard, you will need to first import the existing rules using cf-terraforming. You can find details about how to do this at https://developers.cloudflare.com/terraform/additional-configurations/ddos-managed-rulesets/#optional-delete-existing-rulesets-to-start-from-scratch"
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

	rulesetName := data.Name.ValueString()
	rulesetDescription := data.Description.ValueString()
	rulesetKind := data.Kind.ValueString()
	rulesetPhase := data.Phase.ValueString()

	var identifier *cloudflare.ResourceContainer
	if accountID.ValueString() != "" {
		identifier = cloudflare.AccountIdentifier(accountID.ValueString())
	} else {
		identifier = cloudflare.ZoneIdentifier(zoneID.ValueString())
	}

	ruleset, semaphoreErr := r.client.GetEntrypointRuleset(ctx, identifier, rulesetPhase)

	// If an entrypoint ruleset with the same kind already exists, we should
	// prevent the user from accidentally overriding their existing
	// configuration, since only one entrypoint ruleset for each phase can exist
	// in an account or zone. If the existing entrypoint ruleset is empty, then
	// we should remove it, as it was probably created by the UI.
	//
	// This logic does not apply to non-entrypoint rulesets, such as custom
	// rulesets, as it is possible to have multiple of these rulesets for a
	// phase in an account or zone.
	//
	// We rely on the fact that GetAccountRulesetPhase and GetZoneRulesetPhase
	// only return entrypoint rulesets to check this. If the kind of the ruleset
	// being created does not match the kind of the ruleset returned by that
	// function, then the ruleset being created is not an entrypoint ruleset.
	if rulesetKind == ruleset.Kind {
		if len(ruleset.Rules) > 0 {
			resp.Diagnostics.AddError(
				fmt.Sprintf("failed to create ruleset %q", rulesetPhase),
				duplicateRulesetError,
			)
			return
		}

		if semaphoreErr == nil && len(ruleset.Rules) == 0 && ruleset.Description == "" {
			tflog.Debug(ctx, "default entrypoint ruleset created by the UI with empty rules found, recreating from scratch")

			err := r.client.DeleteRuleset(ctx, identifier, ruleset.ID)
			if err != nil {
				resp.Diagnostics.AddError(
					fmt.Sprintf("failed to delete existing entrypoint ruleset with ID %q", ruleset.ID),
					err.Error(),
				)
				return
			}
		}
	}

	rs := cloudflare.CreateRulesetParams{
		Name:        rulesetName,
		Description: rulesetDescription,
		Kind:        rulesetKind,
		Phase:       rulesetPhase,
	}

	rulesetData := data.toRuleset(ctx)

	if len(rulesetData.Rules) > 0 {
		rs.Rules = rulesetData.Rules
	}

	ruleset, rulesetCreateErr := r.client.CreateRuleset(ctx, identifier, rs)

	if rulesetCreateErr != nil {
		resp.Diagnostics.AddError(fmt.Sprintf("error creating ruleset %s", rulesetName), rulesetCreateErr.Error())
		return
	}

	params := cloudflare.UpdateEntrypointRulesetParams{
		Phase:       rulesetPhase,
		Description: rulesetDescription,
		Rules:       rulesetData.Rules,
	}

	var err error
	// For "custom" rulesets, we don't send a follow up PUT it to the entrypoint
	// endpoint.
	if rulesetKind != string(cloudflare.RulesetKindCustom) {
		ruleset, err = r.client.UpdateEntrypointRuleset(ctx, identifier, params)
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

	diags = resp.State.Set(ctx, toRulesetResourceModel(ctx, data.ZoneID, data.AccountID, ruleset))
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

	var identifier *cloudflare.ResourceContainer
	if accountID.ValueString() != "" {
		identifier = cloudflare.AccountIdentifier(accountID.ValueString())
	} else {
		identifier = cloudflare.ZoneIdentifier(zoneID.ValueString())
	}

	ruleset, err := r.client.GetRuleset(ctx, identifier, data.ID.ValueString())
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

	resp.Diagnostics.Append(resp.State.Set(ctx, toRulesetResourceModel(ctx, zoneID, accountID, ruleset))...)
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

	remappedRules, e := remapPreservedRuleRefs(ctx, state, plan)
	if e != nil {
		resp.Diagnostics.AddError("failed to remap rule IDs from state", e.Error())
		return
	}

	var identifier *cloudflare.ResourceContainer
	if accountID.ValueString() != "" {
		identifier = cloudflare.AccountIdentifier(accountID.ValueString())
	} else {
		identifier = cloudflare.ZoneIdentifier(zoneID)
	}

	params := cloudflare.UpdateRulesetParams{
		ID:          state.ID.ValueString(),
		Description: plan.Description.ValueString(),
		Rules:       remappedRules,
	}
	rs, err := r.client.UpdateRuleset(ctx, identifier, params)
	if err != nil {
		resp.Diagnostics.AddError(fmt.Sprintf("error updating ruleset with ID %q", state.ID.ValueString()), err.Error())
		return
	}

	plan.ID = types.StringValue(rs.ID)

	resp.Diagnostics.Append(resp.State.Set(ctx, toRulesetResourceModel(ctx, state.ZoneID, state.AccountID, rs))...)
}

func (r *RulesetResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *RulesetResourceModel

	diags := req.State.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)

	accountID := data.AccountID
	zoneID := data.ZoneID

	var identifier *cloudflare.ResourceContainer
	if accountID.ValueString() != "" {
		identifier = cloudflare.AccountIdentifier(accountID.ValueString())
	} else {
		identifier = cloudflare.ZoneIdentifier(zoneID.ValueString())
	}

	err := r.client.DeleteRuleset(ctx, identifier, data.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(fmt.Sprintf("error deleting ruleset with ID %q", data.ID.ValueString()), err.Error())
		return
	}
}

func (r *RulesetResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, "/")

	if len(idParts) != 3 || idParts[0] == "" || idParts[1] == "" || idParts[2] == "" {
		resp.Diagnostics.AddError(
			"invalid import identifier",
			fmt.Sprintf("expected import identifier to be resourceLevel/resourceIdentifier/rulesetID. got: %q", req.ID),
		)
		return
	}
	resourceLevel, resourceIdentifier, rulesetID := idParts[0], idParts[1], idParts[2]

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
func toRulesetResourceModel(ctx context.Context, zoneID, accountID basetypes.StringValue, in cloudflare.Ruleset) *RulesetResourceModel {
	data := RulesetResourceModel{
		ID:          types.StringValue(in.ID),
		Description: types.StringValue(in.Description),
		Name:        types.StringValue(in.Name),
		Kind:        types.StringValue(in.Kind),
		Phase:       types.StringValue(in.Phase),
	}

	var ruleState []*RulesModel
	for _, ruleResponse := range in.Rules {
		rule := RulesModel{
			ID:          flatteners.String(ruleResponse.ID),
			Ref:         flatteners.String(ruleResponse.Ref),
			Action:      flatteners.String(ruleResponse.Action),
			Expression:  flatteners.String(ruleResponse.Expression),
			Description: types.StringValue(ruleResponse.Description),
			Enabled:     flatteners.Bool(ruleResponse.Enabled),
			Version:     flatteners.String(cloudflare.String(ruleResponse.Version)),
		}

		if ruleResponse.LastUpdated != nil {
			rule.LastUpdated = types.StringValue(ruleResponse.LastUpdated.String())
		} else {
			rule.LastUpdated = types.StringNull()
		}

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
				OriginCacheControl:      flatteners.Bool(ruleResponse.ActionParameters.OriginCacheControl),
				OriginErrorPagePassthru: flatteners.Bool(ruleResponse.ActionParameters.OriginErrorPagePassthru),
				RespectStrongEtags:      flatteners.Bool(ruleResponse.ActionParameters.RespectStrongETags),
				ReadTimeout:             flatteners.Int64(int64(cloudflare.Uint(ruleResponse.ActionParameters.ReadTimeout))),
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

			var ports []attr.Value
			for _, s := range ruleResponse.ActionParameters.AdditionalCacheablePorts {
				ports = append(ports, types.Int64Value((int64(s))))
			}
			rule.ActionParameters[0].AdditionalCacheablePorts = flatteners.Int64Set(ports)

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
						include, _ := basetypes.NewSetValueFrom(ctx, types.StringType, ruleResponse.ActionParameters.CacheKey.CustomKey.Cookie.Include)
						checkPresence, _ := basetypes.NewSetValueFrom(ctx, types.StringType, ruleResponse.ActionParameters.CacheKey.CustomKey.Cookie.CheckPresence)
						key.Cookie = []*ActionParameterCacheKeyCustomKeyCookieModel{{
							Include:       include,
							CheckPresence: checkPresence,
						}}
					}

					if ruleResponse.ActionParameters.CacheKey.CustomKey.Header != nil {
						include, _ := basetypes.NewSetValueFrom(ctx, types.StringType, ruleResponse.ActionParameters.CacheKey.CustomKey.Header.Include)
						checkPresence, _ := basetypes.NewSetValueFrom(ctx, types.StringType, ruleResponse.ActionParameters.CacheKey.CustomKey.Header.CheckPresence)
						var excludeOrigin types.Bool
						if !reflect.ValueOf(ruleResponse.ActionParameters.CacheKey.CustomKey.Header.ExcludeOrigin).IsNil() {
							excludeOrigin = flatteners.Bool(ruleResponse.ActionParameters.CacheKey.CustomKey.Header.ExcludeOrigin)
						} else {
							excludeOrigin = types.BoolNull()
						}
						if len(include.Elements()) > 0 || len(checkPresence.Elements()) > 0 || excludeOrigin.ValueBool() {
							key.Header = []*ActionParameterCacheKeyCustomKeyHeaderModel{{
								Include:       include,
								CheckPresence: checkPresence,
								ExcludeOrigin: excludeOrigin,
							}}
						}
					}

					if ruleResponse.ActionParameters.CacheKey.CustomKey.Query != nil {
						include, _ := basetypes.NewSetValueFrom(ctx, types.StringType, ruleResponse.ActionParameters.CacheKey.CustomKey.Query.Include)
						exclude, _ := basetypes.NewSetValueFrom(ctx, types.StringType, ruleResponse.ActionParameters.CacheKey.CustomKey.Query.Exclude)

						if ruleResponse.ActionParameters.CacheKey.CustomKey.Query.Include != nil {
							if ruleResponse.ActionParameters.CacheKey.CustomKey.Query.Include.All {
								include, _ = basetypes.NewSetValueFrom(ctx, types.StringType, []string{"*"})
							} else {
								include, _ = basetypes.NewSetValueFrom(ctx, types.StringType, ruleResponse.ActionParameters.CacheKey.CustomKey.Query.Include.List)
							}
						}

						if ruleResponse.ActionParameters.CacheKey.CustomKey.Query.Exclude != nil {
							if ruleResponse.ActionParameters.CacheKey.CustomKey.Query.Exclude.All {
								exclude, _ = basetypes.NewSetValueFrom(ctx, types.StringType, []string{"*"})
							} else {
								exclude, _ = basetypes.NewSetValueFrom(ctx, types.StringType, ruleResponse.ActionParameters.CacheKey.CustomKey.Query.Exclude.List)
							}
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

					var sctValue basetypes.Int64Value
					if sct.Value != nil {
						sctValue = types.Int64Value(int64(cloudflare.Int(sct.Value)))
					} else {
						sctValue = types.Int64Null()
					}

					statusCodeTTLs = append(statusCodeTTLs, &ActionParameterEdgeTTLStatusCodeTTLModel{
						StatusCode:      flatteners.Int64(int64(cloudflare.Uint(sct.StatusCodeValue))),
						Value:           sctValue,
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
					Host: flatteners.String(ruleResponse.ActionParameters.Origin.Host),
					Port: flatteners.Int64(int64(ruleResponse.ActionParameters.Origin.Port)),
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
					StatusCode: flatteners.Int64(int64(ruleResponse.ActionParameters.FromValue.StatusCode)),
					TargetURL: []*ActionParameterFromValueTargetURLModel{{
						Value:      flatteners.String(ruleResponse.ActionParameters.FromValue.TargetURL.Value),
						Expression: flatteners.String(ruleResponse.ActionParameters.FromValue.TargetURL.Expression),
					}},
				}}

				if !reflect.ValueOf(ruleResponse.ActionParameters.FromValue.PreserveQueryString).IsNil() {
					rule.ActionParameters[0].FromValue[0].PreserveQueryString = flatteners.Bool(ruleResponse.ActionParameters.FromValue.PreserveQueryString)
				} else {
					rule.ActionParameters[0].FromValue[0].PreserveQueryString = types.BoolNull()
				}
			}

			if len(ruleResponse.ActionParameters.Algorithms) > 0 {
				var algos []*ActionParametersCompressionAlgorithmModel = nil

				for _, algo := range ruleResponse.ActionParameters.Algorithms {
					newAlgo := ActionParametersCompressionAlgorithmModel{
						Name: algo.Name,
					}
					algos = append(algos, &newAlgo)
				}

				rule.ActionParameters[0].Algorithms = algos
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
				MitigationTimeout:       types.Int64Value(int64(ruleResponse.RateLimit.MitigationTimeout)),
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
func (r *RulesetResourceModel) toRuleset(ctx context.Context) cloudflare.Ruleset {
	var rs cloudflare.Ruleset
	var rules []cloudflare.RulesetRule

	rs.ID = r.ID.ValueString()
	for _, rule := range r.Rules {
		rules = append(rules, rule.toRulesetRule(ctx))
	}

	rs.Rules = rules

	return rs
}

// toRulesetRule takes a state representation of a Ruleset Rule and transforms
// it into an API representation.
func (r *RulesModel) toRulesetRule(ctx context.Context) cloudflare.RulesetRule {
	rr := cloudflare.RulesetRule{
		ID:          r.ID.ValueString(),
		Ref:         r.Ref.ValueString(),
		Version:     r.Version.ValueStringPointer(),
		Action:      r.Action.ValueString(),
		Expression:  r.Expression.ValueString(),
		Description: r.Description.ValueString(),
	}

	if !r.Enabled.IsNull() {
		rr.Enabled = cloudflare.BoolPtr(r.Enabled.ValueBool())
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

		if len(expanders.StringSet(ctx, ap.Phases)) > 0 {
			rr.ActionParameters.Phases = expanders.StringSet(ctx, ap.Phases)
		}

		if len(expanders.StringSet(ctx, ap.Products)) > 0 {
			rr.ActionParameters.Products = expanders.StringSet(ctx, ap.Products)
		}

		if len(expanders.StringSet(ctx, ap.Rulesets)) > 0 {
			rr.ActionParameters.Rulesets = expanders.StringSet(ctx, ap.Rulesets)
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

		if !ap.AdditionalCacheablePorts.IsNull() {
			rr.ActionParameters.AdditionalCacheablePorts = expanders.Int64Set(ctx, ap.AdditionalCacheablePorts)
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

		if !ap.ReadTimeout.IsNull() {
			rr.ActionParameters.ReadTimeout = cloudflare.UintPtr(uint(ap.ReadTimeout.ValueInt64()))
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

		if !ap.OriginCacheControl.IsNull() {
			rr.ActionParameters.OriginCacheControl = cloudflare.BoolPtr(ap.OriginCacheControl.ValueBool())
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
					includeQueryList := expanders.StringSet(ctx, ap.CacheKey[0].CustomKey[0].QueryString[0].Include)
					excludeQueryList := expanders.StringSet(ctx, ap.CacheKey[0].CustomKey[0].QueryString[0].Exclude)

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
					includeQueryList := expanders.StringSet(ctx, ap.CacheKey[0].CustomKey[0].Header[0].Include)
					checkPresenceList := expanders.StringSet(ctx, basetypes.SetValue(ap.CacheKey[0].CustomKey[0].Header[0].CheckPresence))

					customKey.Header = &cloudflare.RulesetRuleActionParametersCustomKeyHeader{
						RulesetRuleActionParametersCustomKeyFields: cloudflare.RulesetRuleActionParametersCustomKeyFields{
							Include:       includeQueryList,
							CheckPresence: checkPresenceList,
						},
						ExcludeOrigin: cloudflare.BoolPtr(ap.CacheKey[0].CustomKey[0].Header[0].ExcludeOrigin.ValueBool()),
					}
				}

				if len(ap.CacheKey[0].CustomKey[0].Cookie) > 0 {
					includeQueryList := expanders.StringSet(ctx, ap.CacheKey[0].CustomKey[0].Cookie[0].Include)
					checkPresenceList := expanders.StringSet(ctx, basetypes.SetValue(ap.CacheKey[0].CustomKey[0].Cookie[0].CheckPresence))

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
				from.PreserveQueryString = cloudflare.BoolPtr(ap.FromValue[0].PreserveQueryString.ValueBool())
			}

			from.TargetURL.Expression = ap.FromValue[0].TargetURL[0].Expression.ValueString()
			from.TargetURL.Value = ap.FromValue[0].TargetURL[0].Value.ValueString()

			rr.ActionParameters.FromValue = from
		}

		apCookieFields := expanders.StringSet(ctx, ap.CookieFields)
		if len(apCookieFields) > 0 {
			for _, cookie := range apCookieFields {
				rr.ActionParameters.CookieFields = append(rr.ActionParameters.CookieFields, cloudflare.RulesetActionParametersLogCustomField{Name: cookie})
			}
		}

		apRequestFields := expanders.StringSet(ctx, ap.RequestFields)
		if len(apRequestFields) > 0 {
			for _, request := range apRequestFields {
				rr.ActionParameters.RequestFields = append(rr.ActionParameters.RequestFields, cloudflare.RulesetActionParametersLogCustomField{Name: request})
			}
		}

		apResponseFields := expanders.StringSet(ctx, ap.ResponseFields)
		if len(apResponseFields) > 0 {
			for _, request := range apResponseFields {
				rr.ActionParameters.ResponseFields = append(rr.ActionParameters.ResponseFields, cloudflare.RulesetActionParametersLogCustomField{Name: request})
			}
		}

		if len(ap.Algorithms) > 0 {
			for _, algo := range ap.Algorithms {
				newAlgo := cloudflare.RulesetRuleActionParametersCompressionAlgorithm{
					Name: algo.Name,
				}
				rr.ActionParameters.Algorithms = append(rr.ActionParameters.Algorithms, newAlgo)
			}
		}
	}

	for _, rl := range r.Ratelimit {
		rr.RateLimit = &cloudflare.RulesetRuleRateLimit{
			Characteristics:         expanders.StringSet(ctx, rl.Characteristics),
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

	if !r.LastUpdated.IsNull() {
		if lastUpdated, err := time.Parse(
			"2006-01-02 15:04:05.999999999 -0700 MST",
			r.LastUpdated.ValueString(),
		); err == nil {
			rr.LastUpdated = &lastUpdated
		}
	}

	return rr
}

// ruleRefs is a lookup table for rule IDs with two operations, add and pop.

// We use add to populate the table from the old value of rules. We use pop to
// look up the ref for the new value of a rule (and remove it from the table).
//
// Internally, both operations serialize the rule to JSON and use the resulting
// string as the lookup key; the ref itself and other computed fields are
// excluded from the JSON.
//
// If a ruleset has multiple copies of the same rule, the copies have a single
// lookup key associated with multiple refs; we preserve order when adding and
// popping the refs.
type ruleRefs struct {
	refs map[string][]string
}

// newRuleRefs creates a new ruleRefs.
func newRuleRefs(rulesetRules []cloudflare.RulesetRule, explicitRefs map[string]struct{}) (ruleRefs, error) {
	r := ruleRefs{make(map[string][]string)}
	for _, rule := range rulesetRules {
		if rule.Ref == "" {
			// This is unexpected. We only invoke this function for the old
			// values of rules, which have their refs populated.
			return ruleRefs{}, errors.New("unable to determine ID or ref of existing rule")
		}

		if _, ok := explicitRefs[rule.Ref]; ok {
			// We should not add explicitly-set refs, to avoid them being
			// "stolen" by other rules.
			continue
		}

		if err := r.add(rule); err != nil {
			return ruleRefs{}, err
		}
	}

	return r, nil
}

// add stores a ref for the given rule.
func (r *ruleRefs) add(rule cloudflare.RulesetRule) error {
	key, err := ruleToKey(rule)
	if err != nil {
		return err
	}

	r.refs[key] = append(r.refs[key], rule.Ref)
	return nil
}

// pop removes a ref for the given rule and returns it. If no ref was found for
// the rule, pop returns an empty string.
func (r *ruleRefs) pop(rule cloudflare.RulesetRule) (string, error) {
	key, err := ruleToKey(rule)
	if err != nil {
		return "", err
	}

	refs := r.refs[key]
	if len(refs) == 0 {
		return "", nil
	}

	ref, refs := refs[0], refs[1:]
	r.refs[key] = refs

	return ref, nil
}

// isEmpty returns true if the store does not contain any rule refs.
func (r *ruleRefs) isEmpty() bool {
	return len(r.refs) == 0
}

// ruleToKey converts a ruleset rule to a key that can be used to track
// equivalent rules. Internally, it serializes the rule to JSON after removing
// computed fields.
func ruleToKey(rule cloudflare.RulesetRule) (string, error) {
	// For the purposes of preserving existing rule refs, we don't want to
	// include computed fields as a part of the key value.
	rule.ID = ""
	rule.Ref = ""
	rule.Version = nil
	rule.LastUpdated = nil

	data, err := json.Marshal(rule)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

// remapPreservedRuleRefs tries to preserve the refs of rules that have not
// changed in the ruleset, while also allowing users to explicitly set the ref
// if they choose to.
func remapPreservedRuleRefs(ctx context.Context, state, plan *RulesetResourceModel) ([]cloudflare.RulesetRule, error) {
	currentRuleset := state.toRuleset(ctx)
	plannedRuleset := plan.toRuleset(ctx)

	plannedExplicitRefs := make(map[string]struct{})
	for _, rule := range plannedRuleset.Rules {
		if rule.Ref != "" {
			plannedExplicitRefs[rule.Ref] = struct{}{}
		}
	}

	refs, err := newRuleRefs(currentRuleset.Rules, plannedExplicitRefs)
	if err != nil {
		return nil, err
	}

	if refs.isEmpty() {
		// There are no rule refs when the ruleset is first created.
		return plannedRuleset.Rules, nil
	}

	for i := range plannedRuleset.Rules {
		rule := &plannedRuleset.Rules[i]

		// We should not override refs that have been explicitly set.
		if rule.Ref == "" {
			if rule.Ref, err = refs.pop(*rule); err != nil {
				return nil, err
			}
		}
	}

	return plannedRuleset.Rules, nil
}

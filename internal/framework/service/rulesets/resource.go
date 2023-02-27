package rulesets

import (
	"context"
	"errors"
	"fmt"
	"reflect"
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
			"unexpected resource configure yype",
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
			_, err = r.client.UpdateAccountRulesetPhase(ctx, accountID.ValueString(), rulesetPhase, rulesetEntryPoint)
		} else {
			_, err = r.client.UpdateZoneRulesetPhase(ctx, zoneID.ValueString(), rulesetPhase, rulesetEntryPoint)
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

	diags = resp.State.Set(ctx, data)
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

	ruleset := plan.toRuleset()

	var err error
	var rs cloudflare.Ruleset
	description := plan.Description.ValueString()
	if accountID.ValueString() != "" {
		rs, err = r.client.UpdateAccountRuleset(ctx, accountID.ValueString(), state.ID.ValueString(), description, ruleset.Rules)
	} else {
		rs, err = r.client.UpdateZoneRuleset(ctx, zoneID, state.ID.ValueString(), description, ruleset.Rules)
	}

	if err != nil {
		resp.Diagnostics.AddError(fmt.Sprintf("error updating ruleset with ID %q", state.ID.ValueString()), err.Error())
		return
	}

	plan.ID = types.StringValue(rs.ID)

	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
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

func toRulesetResourceModel(zoneID, accountID basetypes.StringValue, in cloudflare.Ruleset) *RulesetResourceModel {
	data := RulesetResourceModel{
		ID:          types.StringValue(in.ID),
		Description: types.StringValue(in.Description),
		Name:        types.StringValue(in.Name),
		Kind:        types.StringValue(in.Kind),
		Phase:       types.StringValue(in.Phase),
	}

	var rules []*RulesModel
	for _, inRule := range in.Rules {
		var rule RulesModel

		rule.Action = types.StringValue(inRule.Action)
		rule.Expression = types.StringValue(inRule.Expression)
		rule.Description = types.StringValue(inRule.Description)
		rule.Enabled = types.BoolValue(inRule.Enabled)

		// action_parameters
		if !reflect.ValueOf(inRule.ActionParameters).IsNil() {
			rule.ActionParameters = append(rule.ActionParameters, &ActionParametersModel{
				Response: []*ActionParameterResponseModel{{
					StatusCode:  types.Int64Value(int64(inRule.ActionParameters.Response.StatusCode)),
					ContentType: types.StringValue(inRule.ActionParameters.Response.ContentType),
					Content:     types.StringValue(inRule.ActionParameters.Response.Content),
				}},
			})

			var cookieFields []attr.Value
			for _, s := range inRule.ActionParameters.CookieFields {
				cookieFields = append(cookieFields, types.StringValue(s.Name))
			}
			rule.ActionParameters[0].CookieFields = flatteners.StringSet(cookieFields)
		}

		// ratelimit
		if !reflect.ValueOf(inRule.RateLimit).IsNil() {
			var rlCharacteristicsKeys []attr.Value
			for _, s := range inRule.RateLimit.Characteristics {
				rlCharacteristicsKeys = append(rlCharacteristicsKeys, types.StringValue(s))
			}

			rule.Ratelimit = append(rule.Ratelimit, &RatelimitModel{
				Characteristics:   types.SetValueMust(types.StringType, rlCharacteristicsKeys),
				Period:            types.Int64Value(int64(inRule.RateLimit.Period)),
				RequestsPerPeriod: types.Int64Value(int64(inRule.RateLimit.RequestsPerPeriod)),
				RequestsToOrigin:  types.BoolValue(inRule.RateLimit.RequestsToOrigin),
				MitigationTimeout: types.Int64Value(int64(inRule.RateLimit.MitigationTimeout)),
			})

			if inRule.RateLimit.ScorePerPeriod > 0 {
				rule.Ratelimit[0].ScorePerPeriod = types.Int64Value(int64(inRule.RateLimit.ScorePerPeriod))
			}

			if inRule.RateLimit.ScoreResponseHeaderName != "" {
				rule.Ratelimit[0].ScoreResponseHeaderName = types.StringValue(inRule.RateLimit.ScoreResponseHeaderName)
			}

			if inRule.RateLimit.CountingExpression != "" {
				rule.Ratelimit[0].CountingExpression = types.StringValue(inRule.RateLimit.CountingExpression)
			}
		}

		// logging
		if !reflect.ValueOf(inRule.Logging).IsNil() {
			rule.Logging = append(rule.Logging, &LoggingModel{Enabled: types.BoolValue(*inRule.Logging.Enabled)})
		}

		rules = append(rules, &rule)
	}

	data.Rules = rules

	if zoneID.ValueString() != "" {
		data.ZoneID = types.StringValue(zoneID.ValueString())
	} else {
		data.AccountID = types.StringValue(accountID.ValueString())
	}

	return &data
}

func (r *RulesetResourceModel) toRuleset() cloudflare.Ruleset {
	var rs cloudflare.Ruleset
	var rules []cloudflare.RulesetRule

	rs.ID = r.ID.ValueString()
	for _, rule := range r.Rules {
		newRule := cloudflare.RulesetRule{
			ID:          rule.ID.ValueString(),
			Action:      rule.Action.ValueString(),
			Expression:  rule.Expression.ValueString(),
			Description: rule.Description.ValueString(),
		}

		if !rule.Enabled.IsNull() {
			newRule.Enabled = rule.Enabled.ValueBool()
		}

		if !rule.ID.IsNull() {
			newRule.ID = rule.ID.ValueString()
		}

		if !rule.Ref.IsNull() {
			newRule.Ref = rule.Ref.ValueString()
		}

		if !rule.Version.IsNull() {
			newRule.Version = rule.Version.ValueString()
		}

		for i, ap := range rule.ActionParameters {
			response := cloudflare.RulesetRuleActionParametersBlockResponse{
				ContentType: ap.Response[i].ContentType.ValueString(),
				Content:     ap.Response[i].Content.ValueString(),
				StatusCode:  uint16(ap.Response[i].StatusCode.ValueInt64()),
			}
			newRule.ActionParameters = &cloudflare.RulesetRuleActionParameters{
				ID:       ap.ID.ValueString(),
				Version:  ap.Version.ValueString(),
				Response: &response,
			}

			apCookieFields := expanders.StringSet(ap.CookieFields)
			if len(apCookieFields) > 0 {
				for _, cookie := range apCookieFields {
					newRule.ActionParameters.CookieFields = append(newRule.ActionParameters.CookieFields, cloudflare.RulesetActionParametersLogCustomField{Name: cookie})
				}
			}
		}

		for _, rl := range rule.Ratelimit {
			newRule.RateLimit = &cloudflare.RulesetRuleRateLimit{
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

		rules = append(rules, newRule)
	}
	rs.Rules = rules

	return rs
}

package rulesets

import (
	"context"
	"errors"
	"fmt"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

const (
	accountLevelRulesetDeleteURL = "https://api.cloudflare.com/#account-rulesets-delete-account-ruleset"
	zoneLevelRulesetDeleteURL    = "https://api.cloudflare.com/#zone-rulesets-delete-zone-ruleset"
	duplicateRulesetError        = "failed to create ruleset %q as a similar configuration with rules already exists and overwriting will have unintended consequences. If you are migrating from the Dashboard, you will need to first remove the existing rules otherwise you can remove the existing phase yourself using the API (%s)."
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
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *cloudflare.API, got: %T. Please report this issue to the provider developers.", req.ProviderData),
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

	if !accountID.IsNull() || accountID.ValueString() != "" {
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
			fmt.Sprintf(duplicateRulesetError, rulesetPhase, deleteRulesetURL),
			fmt.Sprintf(duplicateRulesetError, rulesetPhase, deleteRulesetURL),
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

	rules := data.toRulesetRules()

	if len(rules) > 0 {
		rs.Rules = rules
	}

	if sempahoreErr == nil && len(ruleset.Rules) == 0 && ruleset.Description == "" {
		tflog.Debug(ctx, "default ruleset created by the UI with empty rules found, recreating from scratch")
		var deleteRulesetErr error
		if !accountID.IsNull() || accountID.ValueString() != "" {
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
	if !accountID.IsNull() || accountID.ValueString() != "" {
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
		Rules:       rules,
	}

	var err error
	// For "custom" rulesets, we don't send a follow up PUT it to the entrypoint
	// endpoint.
	if rulesetKind != string(cloudflare.RulesetKindCustom) {
		if !accountID.IsNull() || accountID.ValueString() != "" {
			_, err = r.client.UpdateAccountRulesetPhase(ctx, accountID.ValueString(), rulesetPhase, rulesetEntryPoint)
		} else {
			_, err = r.client.UpdateZoneRulesetPhase(ctx, zoneID.ValueString(), rulesetPhase, rulesetEntryPoint)
		}

		if err != nil {
			resp.Diagnostics.AddError(fmt.Sprintf("error updating ruleset phase entrypoint %s", rulesetName), err.Error())
			return
		}
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
	zoneID := data.ZoneID.ValueString()
	var err error
	var ruleset cloudflare.Ruleset

	if !accountID.IsNull() || accountID.ValueString() != "" {
		ruleset, err = r.client.GetAccountRuleset(ctx, accountID.ValueString(), data.ID.ValueString())
	} else {
		ruleset, err = r.client.GetZoneRuleset(ctx, zoneID, data.ID.ValueString())
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

	// set updated ruleset with the tfsdk model
	tflog.Debug(ctx, "ruleset", map[string]interface{}{"id": ruleset.ID})

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
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

	rules := plan.toRulesetRules()

	var err error
	var rs cloudflare.Ruleset
	description := plan.Description.ValueString()
	if !accountID.IsNull() || accountID.ValueString() != "" {
		rs, err = r.client.UpdateAccountRuleset(ctx, accountID.ValueString(), state.ID.ValueString(), description, rules)
	} else {
		rs, err = r.client.UpdateZoneRuleset(ctx, zoneID, state.ID.ValueString(), description, rules)
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

	if !accountID.IsNull() || accountID.ValueString() != "" {
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
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *RulesetResourceModel) toRulesetRules() []cloudflare.RulesetRule {
	var rules []cloudflare.RulesetRule

	for _, rule := range r.Rules {
		newRule := cloudflare.RulesetRule{
			ID:          rule.ID.ValueString(),
			Action:      rule.Action.ValueString(),
			Expression:  rule.Expression.ValueString(),
			Description: rule.Description.ValueString(),
		}

		if !rule.ID.IsNull() {
			newRule.Enabled = cloudflare.BoolPtr(rule.Enabled.ValueBool())
		} else {
			newRule.Enabled = nil
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

		for _, ap := range rule.ActionParameters {
			newRule.ActionParameters = &cloudflare.RulesetRuleActionParameters{
				ID:      ap.ID.ValueString(),
				Version: ap.Version.ValueString(),
			}
		}

		rules = append(rules, newRule)
	}

	return rules
}

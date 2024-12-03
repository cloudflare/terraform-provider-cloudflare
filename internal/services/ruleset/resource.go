// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ruleset

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v3"
	"github.com/cloudflare/cloudflare-go/v3/option"
	"github.com/cloudflare/cloudflare-go/v3/rulesets"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/importpath"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/logging"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.ResourceWithConfigure = (*RulesetResource)(nil)
var _ resource.ResourceWithModifyPlan = (*RulesetResource)(nil)
var _ resource.ResourceWithImportState = (*RulesetResource)(nil)

func NewResource() resource.Resource {
	return &RulesetResource{}
}

// RulesetResource defines the resource implementation.
type RulesetResource struct {
	client *cloudflare.Client
}

func (r *RulesetResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ruleset"
}

func (r *RulesetResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*cloudflare.Client)

	if !ok {
		resp.Diagnostics.AddError(
			"unexpected resource configure type",
			fmt.Sprintf("Expected *cloudflare.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *RulesetResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *RulesetModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	dataBytes, err := data.MarshalJSON()
	if err != nil {
		resp.Diagnostics.AddError("failed to serialize http request", err.Error())
		return
	}
	res := new(http.Response)
	env := RulesetResultEnvelope{*data}
	params := rulesets.RulesetNewParams{}

	if !data.AccountID.IsNull() {
		params.AccountID = cloudflare.F(data.AccountID.ValueString())
	} else {
		params.ZoneID = cloudflare.F(data.ZoneID.ValueString())
	}

	_, err = r.client.Rulesets.New(
		ctx,
		params,
		option.WithRequestBody("application/json", dataBytes),
		option.WithResponseBodyInto(&res),
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to make http request", err.Error())
		return
	}
	bytes, _ := io.ReadAll(res.Body)
	err = apijson.UnmarshalComputed(bytes, &env)
	if err != nil {
		resp.Diagnostics.AddError("failed to deserialize http request", err.Error())
		return
	}

	data = &env.Result

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *RulesetResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *RulesetModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state *RulesetModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	remapPreservedRuleRefs, err := remapPreservedRuleRefs(state, data)
	if err != nil {
		resp.Diagnostics.AddError("failed to remap preserved rule references", err.Error())
		return
	}

	data.Rules = &remapPreservedRuleRefs

	dataBytes, err := data.MarshalJSONForUpdate(*state)
	if err != nil {
		resp.Diagnostics.AddError("failed to serialize http request", err.Error())
		return
	}
	res := new(http.Response)
	env := RulesetResultEnvelope{*data}
	params := rulesets.RulesetUpdateParams{}

	if !data.AccountID.IsNull() {
		params.AccountID = cloudflare.F(data.AccountID.ValueString())
	} else {
		params.ZoneID = cloudflare.F(data.ZoneID.ValueString())
	}

	_, err = r.client.Rulesets.Update(
		ctx,
		data.ID.ValueString(),
		params,
		option.WithRequestBody("application/json", dataBytes),
		option.WithResponseBodyInto(&res),
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to make http request", err.Error())
		return
	}
	bytes, _ := io.ReadAll(res.Body)
	err = apijson.UnmarshalComputed(bytes, &env)
	if err != nil {
		resp.Diagnostics.AddError("failed to deserialize http request", err.Error())
		return
	}
	data = &env.Result

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *RulesetResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *RulesetModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	res := new(http.Response)
	env := RulesetResultEnvelope{*data}
	params := rulesets.RulesetGetParams{}

	if !data.AccountID.IsNull() {
		params.AccountID = cloudflare.F(data.AccountID.ValueString())
	} else {
		params.ZoneID = cloudflare.F(data.ZoneID.ValueString())
	}

	_, err := r.client.Rulesets.Get(
		ctx,
		data.ID.ValueString(),
		params,
		option.WithResponseBodyInto(&res),
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to make http request", err.Error())
		return
	}
	bytes, _ := io.ReadAll(res.Body)
	err = apijson.UnmarshalComputed(bytes, &env)
	if err != nil {
		resp.Diagnostics.AddError("failed to deserialize http request", err.Error())
		return
	}

	data = &env.Result

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *RulesetResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *RulesetModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	params := rulesets.RulesetDeleteParams{}

	if !data.AccountID.IsNull() {
		params.AccountID = cloudflare.F(data.AccountID.ValueString())
	} else {
		params.ZoneID = cloudflare.F(data.ZoneID.ValueString())
	}

	err := r.client.Rulesets.Delete(
		ctx,
		data.ID.ValueString(),
		params,
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to make http request", err.Error())
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *RulesetResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	var data *RulesetModel = new(RulesetModel)
	params := rulesets.RulesetGetParams{}

	path_accounts_or_zones, path_account_id_or_zone_id := "", ""
	path_ruleset_id := ""
	diags := importpath.ParseImportID(
		req.ID,
		"<{accounts|zones}/{account_id|zone_id}>/<ruleset_id>",
		&path_accounts_or_zones,
		&path_account_id_or_zone_id,
		&path_ruleset_id,
	)
	resp.Diagnostics.Append(diags...)
	switch path_accounts_or_zones {
	case "accounts":
		params.AccountID = cloudflare.F(path_account_id_or_zone_id)
		data.AccountID = types.StringValue(path_account_id_or_zone_id)
	case "zones":
		params.ZoneID = cloudflare.F(path_account_id_or_zone_id)
		data.ZoneID = types.StringValue(path_account_id_or_zone_id)
	default:
		resp.Diagnostics.AddError("invalid discriminator segment - <{accounts|zones}/{account_id|zone_id}>", "expected discriminator to be one of {accounts|zones}")
	}
	if resp.Diagnostics.HasError() {
		return
	}

	data.ID = types.StringValue(path_ruleset_id)

	res := new(http.Response)
	env := RulesetResultEnvelope{*data}
	_, err := r.client.Rulesets.Get(
		ctx,
		path_ruleset_id,
		params,
		option.WithResponseBodyInto(&res),
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to make http request", err.Error())
		return
	}
	bytes, _ := io.ReadAll(res.Body)
	err = apijson.Unmarshal(bytes, &env)
	if err != nil {
		resp.Diagnostics.AddError("failed to deserialize http request", err.Error())
		return
	}
	data = &env.Result

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *RulesetResource) ModifyPlan(_ context.Context, _ resource.ModifyPlanRequest, _ *resource.ModifyPlanResponse) {
}

// ruleRefs is a lookup table for rule IDs with two operations, add and pop.
//
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
func newRuleRefs(rulesetRules *[]*RulesetRulesModel, explicitRefs map[string]struct{}) (ruleRefs, error) {
	r := ruleRefs{make(map[string][]string)}
	thing := rulesetRules
	for _, rule := range *thing {
		if rule.Ref.IsNull() {
			// This is unexpected. We only invoke this function for the old
			// values of rules, which have their refs populated.
			return ruleRefs{}, errors.New("unable to determine ID or ref of existing rule")
		}

		if _, ok := explicitRefs[rule.Ref.ValueString()]; ok {
			// We should not add explicitly-set refs, to avoid them being
			// "stolen" by other rules.
			continue
		}

		if err := r.add(rule, rule.Ref); err != nil {
			return ruleRefs{}, err
		}
	}

	return r, nil
}

// add stores a ref for the given rule.
func (r *ruleRefs) add(rule *RulesetRulesModel, ruleRef basetypes.StringValue) error {
	key, err := ruleToKey(rule)
	if err != nil {
		return err
	}

	r.refs[key.ValueString()] = append(r.refs[key.ValueString()], ruleRef.ValueString())
	return nil
}

// pop removes a ref for the given rule and returns it. If no ref was found for
// the rule, pop returns an `null` value.
func (r *ruleRefs) pop(rule *RulesetRulesModel) (basetypes.StringValue, error) {
	key, err := ruleToKey(rule)
	if err != nil {
		return types.StringNull(), err
	}

	refs := r.refs[key.ValueString()]
	if len(refs) == 0 {
		return types.StringNull(), nil
	}

	ref, refs := refs[0], refs[1:]

	r.refs[key.ValueString()] = refs

	return types.StringValue(ref), nil
}

// isEmpty returns true if the store does not contain any rule refs.
func (r *ruleRefs) isEmpty() bool {
	return len(r.refs) == 0
}

// ruleToKey converts a ruleset rule to a key that can be used to track
// equivalent rules. Internally, it serializes the rule to JSON after removing
// computed fields.
func ruleToKey(rule *RulesetRulesModel) (basetypes.StringValue, error) {
	// For the purposes of preserving existing rule refs, we don't want to
	// include computed fields as a part of the key value.
	rule.ID = types.StringNull()
	rule.Ref = types.StringNull()
	rule.Version = types.StringNull()
	rule.LastUpdated = timetypes.RFC3339{}

	data, err := apijson.Marshal(rule)
	if err != nil {
		return types.StringNull(), err
	}

	return types.StringValue(string(data)), nil
}

// remapPreservedRuleRefs tries to preserve the refs of rules that have not
// changed in the ruleset, while also allowing users to explicitly set the ref
// if they choose to.
func remapPreservedRuleRefs(state, plan *RulesetModel) ([]*RulesetRulesModel, error) {
	currentRuleset := state
	plannedRuleset := plan

	plannedExplicitRefs := make(map[string]struct{})
	plannedRules := plannedRuleset.Rules

	for _, rule := range *plannedRules {
		if !rule.Ref.IsNull() {
			plannedExplicitRefs[rule.Ref.ValueString()] = struct{}{}
		}
	}

	refs, err := newRuleRefs(currentRuleset.Rules, plannedExplicitRefs)
	if err != nil {
		return nil, err
	}

	if refs.isEmpty() {
		// There are no rule refs when the ruleset is first created.
		return *plannedRuleset.Rules, nil
	}

	for i := range *plannedRules {
		thing := *plannedRuleset.Rules
		rule := thing[i]

		// We should not override refs that have been explicitly set.
		if rule.Ref.IsUnknown() {

			if rule.Ref, err = refs.pop(rule); err != nil {
				return nil, err
			}
		}
	}

	return *plannedRules, nil
}

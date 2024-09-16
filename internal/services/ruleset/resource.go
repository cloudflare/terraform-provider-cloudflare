// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ruleset

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/rulesets"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/importpath"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/logging"
	"github.com/hashicorp/terraform-plugin-framework/resource"
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

	dataBytes, err := apijson.Marshal(data)
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

	jsonRemappedRuleRefs, _ := apijson.Marshal(remapPreservedRuleRefs)
	var tmp *[]*RulesetRulesModel
	apijson.UnmarshalComputed(jsonRemappedRuleRefs, &tmp)

	data.Rules = tmp

	dataBytes, err := apijson.MarshalForUpdate(data, state)
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
	var data *RulesetModel
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
	case "zones":
		params.ZoneID = cloudflare.F(path_account_id_or_zone_id)
	default:
		resp.Diagnostics.AddError("invalid discriminator segment - <{accounts|zones}/{account_id|zone_id}>", "expected discriminator to be one of {accounts|zones}")
	}
	if resp.Diagnostics.HasError() {
		return
	}

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
	err = apijson.UnmarshalComputed(bytes, &env)
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
func newRuleRefs(rulesetRules []rulesets.RulesetNewResponseRule, explicitRefs map[string]struct{}) (ruleRefs, error) {
	r := ruleRefs{make(map[string][]string)}
	for _, rule := range rulesetRules {
		if rule.Ref == "" {
			// This is unexpected. We only invoke this function for the old
			// values of rules, which have their refs populated.
			return ruleRefs{}, errors.New(fmt.Sprintf("unable to determine ID or ref of existing rule. rule %+v", rule))
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
func (r *ruleRefs) add(rule rulesets.RulesetNewResponseRule) error {
	key, err := ruleToKey(rule)
	if err != nil {
		return err
	}

	r.refs[key] = append(r.refs[key], rule.Ref)
	return nil
}

// pop removes a ref for the given rule and returns it. If no ref was found for
// the rule, pop returns an empty string.
func (r *ruleRefs) pop(rule rulesets.RulesetNewResponseRule) (string, error) {
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
func ruleToKey(rule rulesets.RulesetNewResponseRule) (string, error) {
	rule.Ref = ""
	rule.ID = ""
	rule.Version = ""
	rule.LastUpdated = time.Time{}

	d, _ := apijson.Marshal(rule)
	return string(d), nil
}

// remapPreservedRuleRefs tries to preserve the refs of rules that have not
// changed in the ruleset, while also allowing users to explicitly set the ref
// if they choose to.
func remapPreservedRuleRefs(state, plan *RulesetModel) ([]rulesets.RulesetNewResponseRule, error) {
	var currentRuleset rulesets.RulesetNewResponse
	var plannedRuleset rulesets.RulesetNewResponse

	jsonState, _ := apijson.Marshal(state)
	jsonPlan, _ := apijson.Marshal(plan)

	err := apijson.Unmarshal(jsonState, &currentRuleset)
	if err != nil {
		return nil, err
	}

	err = apijson.Unmarshal(jsonPlan, &plannedRuleset)
	if err != nil {
		return nil, err
	}

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

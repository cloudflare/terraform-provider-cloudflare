// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ruleset_rule

import (
	"context"
	"fmt"
	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/rulesets"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijsoncustom"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/importpath"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/logging"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/ruleset"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"io"
	"net/http"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.ResourceWithConfigure = (*RulesetRuleResource)(nil)
var _ resource.ResourceWithImportState = (*RulesetRuleResource)(nil)

func NewResource() resource.Resource {
	return &RulesetRuleResource{}
}

// RulesetRuleResource defines the resource implementation.
type RulesetRuleResource struct {
	client *cloudflare.Client
}

func (r *RulesetRuleResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ruleset_rule"
}

func (r *RulesetRuleResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *RulesetRuleResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *RulesetRuleModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	postRuleRes := new(http.Response)
	postRuleEnv := ruleset.RulesetResultEnvelope{}
	postRuleParams := rulesets.RuleNewParams{}

	if !data.AccountID.IsNull() {
		postRuleParams.AccountID = cloudflare.F(data.AccountID.ValueString())
	} else {
		postRuleParams.ZoneID = cloudflare.F(data.ZoneID.ValueString())
	}

	dataBytes, err := data.MarshalJSON()
	if err != nil {
		resp.Diagnostics.AddError("failed to serialize http request", err.Error())
		return
	}

	resp.Diagnostics.AddWarning("kur4ence", fmt.Sprintf("sending body %#v", string(dataBytes)))

	_, err = r.client.Rulesets.Rules.New(
		ctx,
		data.RulesetID.ValueString(),
		postRuleParams,
		option.WithRequestBody("application/json", dataBytes),
		option.WithResponseBodyInto(&postRuleRes),
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to update ruleset", err.Error())
		return
	}

	updateResultBytes, _ := io.ReadAll(postRuleRes.Body)
	err = apijsoncustom.Unmarshal(updateResultBytes, &postRuleEnv)
	if err != nil {
		resp.Diagnostics.AddError("failed to deserialize updated ruleset", err.Error())
		return
	}

	updatedRulesSlice, diags := postRuleEnv.Result.Rules.AsStructSliceT(ctx)
	if diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}

	if len(updatedRulesSlice) > 0 {
		for _, rule := range updatedRulesSlice {
			if rule.Expression.ValueString() == data.Expression.ValueString() && !rule.ID.IsNull() {
				data.ID = rule.ID
				data.Ref = rule.Ref
				data.Logging = rule.Logging
				break
			}
		}
	}

	if data.ID.IsNull() {
		resp.Diagnostics.AddError("API Error", "Unable to determine the ID of the created rule")
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *RulesetRuleResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *RulesetRuleModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Since there's no GET endpoint for individual rules, we need to read the entire ruleset
	res := new(http.Response)
	env := ruleset.RulesetResultEnvelope{}
	params := rulesets.RulesetGetParams{}

	if !data.AccountID.IsNull() {
		params.AccountID = cloudflare.F(data.AccountID.ValueString())
	} else {
		params.ZoneID = cloudflare.F(data.ZoneID.ValueString())
	}
	_, err := r.client.Rulesets.Get(
		ctx,
		data.RulesetID.ValueString(),
		params,
		option.WithResponseBodyInto(&res),
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if res != nil && res.StatusCode == 404 {
		resp.Diagnostics.AddWarning("Resource not found", "The ruleset was not found on the server and will be removed from state.")
		resp.State.RemoveResource(ctx)
		return
	}
	if err != nil {
		resp.Diagnostics.AddError("failed to make http request", err.Error())
		return
	}

	bytes, _ := io.ReadAll(res.Body)
	err = apijsoncustom.Unmarshal(bytes, &env)
	if err != nil {
		resp.Diagnostics.AddError("failed to deserialize http request", err.Error())
		return
	}

	// Find our rule in the ruleset
	var foundRule *ruleset.RulesetRulesModel
	ruleID := data.ID.ValueString()

	rules, diags := env.Result.Rules.AsStructSliceT(ctx)
	if diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}

	for _, rule := range rules {
		if rule.ID.ValueString() == ruleID {
			foundRule = &rule
			break
		}
	}

	if foundRule == nil {
		// Rule not found, it may have been deleted outside of Terraform
		resp.State.RemoveResource(ctx)
		return
	}

	// Update the model with the current rule data
	data.Action = foundRule.Action
	data.Expression = foundRule.Expression
	data.Description = foundRule.Description
	data.Enabled = foundRule.Enabled
	data.ActionParameters = foundRule.ActionParameters
	data.ExposedCredentialCheck = foundRule.ExposedCredentialCheck
	data.Logging = foundRule.Logging
	data.Ratelimit = foundRule.Ratelimit
	data.Ref = foundRule.Ref

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *RulesetRuleResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *RulesetRuleModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state *RulesetRuleModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	dataBytes, err := data.MarshalJSONForUpdate(*state)
	if err != nil {
		resp.Diagnostics.AddError("failed to serialize http request", err.Error())
		return
	}
	res := new(http.Response)
	env := ruleset.RulesetResultEnvelope{}
	params := rulesets.RuleEditParams{}

	if !data.AccountID.IsNull() {
		params.AccountID = cloudflare.F(data.AccountID.ValueString())
	} else {
		params.ZoneID = cloudflare.F(data.ZoneID.ValueString())
	}

	_, err = r.client.Rulesets.Rules.Edit(
		ctx,
		data.RulesetID.ValueString(),
		state.ID.ValueString(),
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
	err = apijsoncustom.Unmarshal(bytes, &env)
	if err != nil {
		resp.Diagnostics.AddError("failed to deserialize http request", err.Error())
		return
	}

	var foundRule *ruleset.RulesetRulesModel

	rules, diags := env.Result.Rules.AsStructSliceT(ctx)
	if diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}

	for _, rule := range rules {
		if rule.Expression.ValueString() == data.Expression.ValueString() && !rule.ID.IsNull() {
			foundRule = &rule
			break
		}
	}

	if foundRule == nil {
		resp.Diagnostics.AddError("Rule not found", "The rule was not found in the ruleset after update")
		return
	}

	data.ID = foundRule.ID
	data.Ref = foundRule.Ref
	data.Action = foundRule.Action
	data.ActionParameters = foundRule.ActionParameters
	data.Description = foundRule.Description
	data.Enabled = foundRule.Enabled
	data.ExposedCredentialCheck = foundRule.ExposedCredentialCheck
	data.Expression = foundRule.Expression
	data.Logging = foundRule.Logging
	data.Ratelimit = foundRule.Ratelimit

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *RulesetRuleResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *RulesetRuleModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	res := new(http.Response)
	//env := ruleset.RulesetResultEnvelope{}
	params := rulesets.RuleDeleteParams{}

	if !data.AccountID.IsNull() {
		params.AccountID = cloudflare.F(data.AccountID.ValueString())
	} else {
		params.ZoneID = cloudflare.F(data.ZoneID.ValueString())
	}

	_, err := r.client.Rulesets.Rules.Delete(
		ctx,
		data.RulesetID.ValueString(),
		data.ID.ValueString(),
		params,
		option.WithResponseBodyInto(&res),
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if res != nil && res.StatusCode == 404 {
		resp.Diagnostics.AddWarning("Resource not found", "The resource was not found on the server and will be removed from state.")
		resp.State.RemoveResource(ctx)
		return
	}
	if err != nil {
		resp.Diagnostics.AddError("failed to make http request", err.Error())
		return
	}

	resp.State.RemoveResource(ctx)
}

func (r *RulesetRuleResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	var data *RulesetRuleModel = new(RulesetRuleModel)
	params := rulesets.RulesetGetParams{}

	path_accounts_or_zones, path_account_id_or_zone_id := "", ""
	path_ruleset_id, path_rule_id := "", ""

	diags := importpath.ParseImportID(
		req.ID,
		"<{accounts|zones}/{account_id|zone_id}>/<ruleset_id>/<rule_id>",
		&path_accounts_or_zones,
		&path_account_id_or_zone_id,
		&path_ruleset_id,
		&path_rule_id,
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

	data.ID = types.StringValue(path_rule_id)
	data.RulesetID = types.StringValue(path_ruleset_id)

	res := new(http.Response)
	env := ruleset.RulesetResultEnvelope{}

	if !data.AccountID.IsNull() {
		params.AccountID = cloudflare.F(data.AccountID.ValueString())
	} else {
		params.ZoneID = cloudflare.F(data.ZoneID.ValueString())
	}
	_, err := r.client.Rulesets.Get(
		ctx,
		data.RulesetID.ValueString(),
		params,
		option.WithResponseBodyInto(&res),
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if res != nil && res.StatusCode == 404 {
		resp.Diagnostics.AddWarning("Resource not found", "The ruleset was not found on the server and will be removed from state.")
		resp.State.RemoveResource(ctx)
		return
	}
	if err != nil {
		resp.Diagnostics.AddError("failed to make http request", err.Error())
		return
	}

	bytes, _ := io.ReadAll(res.Body)
	err = apijsoncustom.Unmarshal(bytes, &env)
	if err != nil {
		resp.Diagnostics.AddError("failed to deserialize http request", err.Error())
		return
	}

	// Find our rule in the ruleset
	var foundRule *ruleset.RulesetRulesModel
	ruleID := data.ID.ValueString()

	rules, diags := env.Result.Rules.AsStructSliceT(ctx)
	if diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}

	for _, rule := range rules {
		if rule.ID.ValueString() == ruleID {
			foundRule = &rule
			break
		}
	}

	if foundRule == nil {
		// Rule not found
		return
	}

	// Update the model with the current rule data
	data.Action = foundRule.Action
	data.Expression = foundRule.Expression
	data.Description = foundRule.Description
	data.Enabled = foundRule.Enabled
	data.ActionParameters = foundRule.ActionParameters
	data.ExposedCredentialCheck = foundRule.ExposedCredentialCheck
	data.Logging = foundRule.Logging
	data.Ratelimit = foundRule.Ratelimit
	data.Ref = foundRule.Ref

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ruleset

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/rulesets"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijsoncustom"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/importpath"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/logging"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
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
	bytes = transformQueryStringJSON(bytes)
	err = apijsoncustom.UnmarshalComputed(bytes, &env)
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
	bytes = transformQueryStringJSON(bytes)
	err = apijsoncustom.UnmarshalComputed(bytes, &env)
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
	if res != nil && res.StatusCode == 404 {
		resp.Diagnostics.AddWarning("Resource not found", "The resource was not found on the server and will be removed from state.")
		resp.State.RemoveResource(ctx)
		return
	}
	if err != nil {
		resp.Diagnostics.AddError("failed to make http request", err.Error())
		return
	}
	bytes, _ := io.ReadAll(res.Body)

	// Transform query_string structure in raw JSON before unmarshaling
	// API returns: "include": ["item1", "item2"]
	// Schema expects: "include": { "list": ["item1", "item2"] }
	bytes = transformQueryStringJSON(bytes)

	err = apijsoncustom.Unmarshal(bytes, &env)
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
	bytes = transformQueryStringJSON(bytes)
	err = apijsoncustom.Unmarshal(bytes, &env)
	if err != nil {
		resp.Diagnostics.AddError("failed to deserialize http request", err.Error())
		return
	}
	data = &env.Result

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

const privateStateKeyExistingRuleset = "existing_ruleset"

func (r *RulesetResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	var state *RulesetModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var plan *RulesetModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Mark every plan that has state. A replacement plans its create half with no
	// state, exactly like a plain create, and reads the mark to tell them apart
	if !req.State.Raw.IsNull() {
		resp.Diagnostics.Append(resp.Private.SetKey(ctx, privateStateKeyExistingRuleset, []byte("true"))...)
		if resp.Diagnostics.HasError() {
			return
		}
	}

	// Validate the plan once the planned rules below have been populated
	defer func() {
		if resp.Diagnostics.HasError() {
			return
		}

		// A ruleset that is not changing is not written during apply, so skip dry-run
		if !req.State.Raw.IsNull() && !resp.Plan.Raw.IsNull() && req.State.Raw.Equal(resp.Plan.Raw) {
			return
		}

		// Check if enough of the plan is known to build a request for the dry-run
		if !req.Config.Raw.IsFullyKnown() {
			return
		}

		r.validateWithDryRun(ctx, req, state, plan, &resp.Diagnostics)
	}()

	// Do nothing if there is no state or no plan.
	if state == nil || plan == nil {
		return
	}

	// Do nothing if there is no rules attribute in the state or the plan.
	if state.Rules.IsNullOrUnknown() || plan.Rules.IsNullOrUnknown() {
		return
	}

	rulesByRef := make(map[string]*RulesetRulesModel)

	stateRules, diags := state.Rules.AsStructSliceT(ctx)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	for i := range stateRules {
		if ref := stateRules[i].Ref.ValueString(); ref != "" {
			rulesByRef[ref] = &stateRules[i]
		}
	}

	planRules, diags := plan.Rules.AsStructSliceT(ctx)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	for i := range planRules {
		if ref := planRules[i].Ref.ValueString(); ref != "" {
			if stateRule, ok := rulesByRef[ref]; ok {
				// If the rule's ref matches a rule from the state, populate its
				// planned ID using the matching rule.
				if planRules[i].ID.IsUnknown() {
					planRules[i].ID = stateRule.ID
				}

				// If the rule's action is unchanged, populate its planned
				// logging attribute using the matching rule from the state.
				if planRules[i].Logging.IsUnknown() &&
					stateRule.Action.Equal(planRules[i].Action) {
					planRules[i].Logging = stateRule.Logging
				}
			}
		}
	}

	plan.Rules, diags = customfield.NewObjectList(ctx, planRules)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.Plan.Set(ctx, plan)...)
}

// transformQueryStringJSON transforms query_string.include from API format to schema format in raw JSON
// The Cloudflare API returns include as a direct array, but the v5 schema expects it wrapped in an object
// API format: { "include": ["param1", "param2"] }
// Schema format: { "include": { "list": ["param1", "param2"] } }
func transformQueryStringJSON(jsonBytes []byte) []byte {
	jsonStr := string(jsonBytes)
	parsed := gjson.Parse(jsonStr)

	// Navigate to rules array in the result
	rules := parsed.Get("result.rules")
	if !rules.Exists() || !rules.IsArray() {
		return jsonBytes
	}

	modified := false
	rules.ForEach(func(key, rule gjson.Result) bool {
		rulePath := fmt.Sprintf("result.rules.%s", key.String())

		// Check if action_parameters.cache_key.custom_key.query_string.include exists
		includePath := rulePath + ".action_parameters.cache_key.custom_key.query_string.include"
		includeValue := gjson.Get(jsonStr, includePath)

		if includeValue.Exists() && includeValue.IsArray() {
			// Transform array to object with "list" field
			// Check if it's a wildcard case: ["*"]
			if len(includeValue.Array()) == 1 && includeValue.Array()[0].String() == "*" {
				// Convert ["*"] to { "all": true }
				newInclude := map[string]interface{}{
					"all": true,
				}
				jsonStr, _ = sjson.Set(jsonStr, includePath, newInclude)
				modified = true
			} else {
				// Convert ["item1", "item2"] to { "list": ["item1", "item2"] }
				newInclude := map[string]interface{}{
					"list": includeValue.Value(),
				}
				jsonStr, _ = sjson.Set(jsonStr, includePath, newInclude)
				modified = true
			}
		}

		// Also handle exclude if present
		excludePath := rulePath + ".action_parameters.cache_key.custom_key.query_string.exclude"
		excludeValue := gjson.Get(jsonStr, excludePath)

		if excludeValue.Exists() && excludeValue.IsArray() {
			// Convert ["item1", "item2"] to { "list": ["item1", "item2"] }
			newExclude := map[string]interface{}{
				"list": excludeValue.Value(),
			}
			jsonStr, _ = sjson.Set(jsonStr, excludePath, newExclude)
			modified = true
		}

		return true
	})

	if modified {
		return []byte(jsonStr)
	}
	return jsonBytes
}

func (r *RulesetResource) validateWithDryRun(
	ctx context.Context,
	req resource.ModifyPlanRequest,
	state *RulesetModel,
	plan *RulesetModel,
	diagnostics *diag.Diagnostics,
) {
	if r.client == nil {
		return
	}

	requestOptions := []option.RequestOption{
		option.WithQuery("dry_run", "true"),
		option.WithMiddleware(logging.Middleware(ctx)),
	}

	var err error
	switch {
	// Terraform plans a create as a null state
	case req.State.Raw.IsNull():
		// skip the dry run when creating as part of a replacement
		if replacesExistingRuleset(ctx, req) {
			return
		}

		dataBytes, marshalErr := plan.MarshalJSON()
		if marshalErr != nil {
			diagnostics.AddError("failed to serialize request for validation", marshalErr.Error())
			return
		}

		params := rulesets.RulesetNewParams{}
		if !plan.AccountID.IsNull() {
			params.AccountID = cloudflare.F(plan.AccountID.ValueString())
		} else {
			params.ZoneID = cloudflare.F(plan.ZoneID.ValueString())
		}

		_, err = r.client.Rulesets.New(ctx, params,
			append(requestOptions, option.WithRequestBody("application/json", dataBytes))...)

	// Terraform plans a destroy as a null plan
	case req.Plan.Raw.IsNull():
		return

	// Terraform carries out a replacement as a destroy followed by a create that it
	// plans in a separate call, skip dry run
	case replacesRuleset(state, plan):
		return

	// Updating a ruleset that already exists
	default:
		dataBytes, marshalErr := plan.MarshalJSONForUpdate(*state)
		if marshalErr != nil {
			diagnostics.AddError("failed to serialize request for validation", marshalErr.Error())
			return
		}

		params := rulesets.RulesetUpdateParams{}
		if !plan.AccountID.IsNull() {
			params.AccountID = cloudflare.F(plan.AccountID.ValueString())
		} else {
			params.ZoneID = cloudflare.F(plan.ZoneID.ValueString())
		}

		_, err = r.client.Rulesets.Update(ctx, plan.ID.ValueString(), params,
			append(requestOptions, option.WithRequestBody("application/json", dataBytes))...)
	}

	if err == nil {
		return
	}

	var apiErr *cloudflare.Error
	if errors.As(err, &apiErr) && requestWasEvaluated(apiErr.StatusCode) {
		diagnostics.AddError(
			"the request was rejected during validation",
			"The request was validated without being applied, so nothing has changed.\n\n"+err.Error(),
		)

		return
	}

	diagnostics.AddWarning(
		"could not validate the request",
		"The request was not validated, so this plan does not confirm that apply will succeed.\n\n"+err.Error(),
	)
}

// requestWasEvaluated reports whether the API inspected the request and rejected it,
// rather than failing before inspection
func requestWasEvaluated(statusCode int) bool {
	switch statusCode {
	case http.StatusUnauthorized, http.StatusForbidden, http.StatusTooManyRequests:
		return false
	}

	return statusCode >= http.StatusBadRequest && statusCode < http.StatusInternalServerError
}

// replacesExistingRuleset checks whether a planned create is the second half of a
// replacement, by looking for the mark ModifyPlan leaves on a ruleset that is there
func replacesExistingRuleset(ctx context.Context, req resource.ModifyPlanRequest) bool {
	existing, _ := req.Private.GetKey(ctx, privateStateKeyExistingRuleset)

	return existing != nil
}

// replacesRuleset reports whether changing from the state to the plan replaces the
// ruleset instead of updating it in place
func replacesRuleset(state, plan *RulesetModel) bool {
	return !plan.AccountID.Equal(state.AccountID) ||
		!plan.ZoneID.Equal(state.ZoneID) ||
		!plan.Kind.Equal(state.Kind) ||
		!plan.Name.Equal(state.Name) ||
		!plan.Phase.Equal(state.Phase)
}

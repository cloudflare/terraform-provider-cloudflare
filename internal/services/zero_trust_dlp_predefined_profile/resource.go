// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_dlp_predefined_profile

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/importpath"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/logging"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.ResourceWithConfigure = (*ZeroTrustDLPPredefinedProfileResource)(nil)
var _ resource.ResourceWithModifyPlan = (*ZeroTrustDLPPredefinedProfileResource)(nil)
var _ resource.ResourceWithImportState = (*ZeroTrustDLPPredefinedProfileResource)(nil)

func NewResource() resource.Resource {
	return &ZeroTrustDLPPredefinedProfileResource{}
}

// ZeroTrustDLPPredefinedProfileResource defines the resource implementation.
type ZeroTrustDLPPredefinedProfileResource struct {
	client *cloudflare.Client
}

func (r *ZeroTrustDLPPredefinedProfileResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_zero_trust_dlp_predefined_profile"
}

func (r *ZeroTrustDLPPredefinedProfileResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *ZeroTrustDLPPredefinedProfileResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *ZeroTrustDLPPredefinedProfileModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve entries from plan to avoid inconsistency when API returns an empty list.
	// On first Create there is no prior state, so ModifyPlan cannot copy state→plan for
	// entries. The API always returns the full entries list (possibly empty []), but the
	// plan has entries=null when the user omits the deprecated attribute. We restore the
	// plan value so state matches what Terraform expected after apply.
	planEntries := data.Entries

	dataBytes, err := data.MarshalJSON()
	if err != nil {
		resp.Diagnostics.AddError("failed to serialize http request", err.Error())
		return
	}
	res := new(http.Response)
	env := ZeroTrustDLPPredefinedProfileResultEnvelope{*data}
	_, err = r.client.ZeroTrust.DLP.Profiles.Predefined.Update(
		ctx,
		data.ProfileID.ValueString(),
		zero_trust.DLPProfilePredefinedUpdateParams{
			AccountID: cloudflare.F(data.AccountID.ValueString()),
		},
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
	data.ID = data.ProfileID

	// Restore entries from plan when the API returned an empty list but the plan had
	// entries=null (user omitted the deprecated attribute). This prevents the framework
	// from raising "Provider produced inconsistent result after apply" on first create.
	if planEntries.IsNull() && data.Entries.IsNull() == false && len(data.Entries.Elements()) == 0 {
		data.Entries = planEntries
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ZeroTrustDLPPredefinedProfileResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *ZeroTrustDLPPredefinedProfileModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	var state *ZeroTrustDLPPredefinedProfileModel

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
	env := ZeroTrustDLPPredefinedProfileResultEnvelope{*data}
	_, err = r.client.ZeroTrust.DLP.Profiles.Predefined.Update(
		ctx,
		data.ProfileID.ValueString(),
		zero_trust.DLPProfilePredefinedUpdateParams{
			AccountID: cloudflare.F(data.AccountID.ValueString()),
		},
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
	data.ID = data.ProfileID

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ZeroTrustDLPPredefinedProfileResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *ZeroTrustDLPPredefinedProfileModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve enabled_entries from prior state to avoid drift when API returns null
	priorEnabledEntries := data.EnabledEntries

	res := new(http.Response)
	env := ZeroTrustDLPPredefinedProfileResultEnvelope{*data}
	_, err := r.client.ZeroTrust.DLP.Profiles.Predefined.Get(
		ctx,
		data.ProfileID.ValueString(),
		zero_trust.DLPProfilePredefinedGetParams{
			AccountID: cloudflare.F(data.AccountID.ValueString()),
		},
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
	err = apijson.Unmarshal(bytes, &env)
	if err != nil {
		resp.Diagnostics.AddError("failed to deserialize http request", err.Error())
		return
	}
	data = &env.Result
	data.ID = data.ProfileID

	// Restore enabled_entries from prior state when the API returns null or omits the field.
	//
	// The Cloudflare API returns "enabled_entries": null (or omits the field) for profiles
	// where no entries are explicitly enabled. apijson.Unmarshal with Always update behaviour
	// decodes a JSON null into a non-nil pointer to a nil slice (*[]types.String where *ptr
	// is nil) — not a Go nil pointer. The Terraform Framework treats these differently: a
	// non-nil pointer to a nil slice is not the same as an empty list [].
	//
	// We detect both cases (nil pointer and non-nil pointer-to-nil) and restore the prior
	// state value so that config == state after every refresh cycle.
	apiReturnedNullOrAbsent := data.EnabledEntries == nil || *data.EnabledEntries == nil
	if apiReturnedNullOrAbsent && priorEnabledEntries != nil {
		data.EnabledEntries = priorEnabledEntries
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ZeroTrustDLPPredefinedProfileResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *ZeroTrustDLPPredefinedProfileModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	_, err := r.client.ZeroTrust.DLP.Profiles.Predefined.Delete(
		ctx,
		data.ProfileID.ValueString(),
		zero_trust.DLPProfilePredefinedDeleteParams{
			AccountID: cloudflare.F(data.AccountID.ValueString()),
		},
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to make http request", err.Error())
		return
	}
	data.ID = data.ProfileID

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ZeroTrustDLPPredefinedProfileResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	var data = new(ZeroTrustDLPPredefinedProfileModel)

	path_account_id := ""
	path_profile_id := ""
	diags := importpath.ParseImportID(
		req.ID,
		"<account_id>/<profile_id>",
		&path_account_id,
		&path_profile_id,
	)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.AccountID = types.StringValue(path_account_id)
	data.ProfileID = types.StringValue(path_profile_id)

	res := new(http.Response)
	env := ZeroTrustDLPPredefinedProfileResultEnvelope{*data}
	_, err := r.client.ZeroTrust.DLP.Profiles.Predefined.Get(
		ctx,
		path_profile_id,
		zero_trust.DLPProfilePredefinedGetParams{
			AccountID: cloudflare.F(path_account_id),
		},
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
	data.ID = data.ProfileID

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ZeroTrustDLPPredefinedProfileResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	// Suppress drift on the deprecated `entries` attribute when the user has not set it
	// in their config (i.e. they are using `enabled_entries` instead, which is the correct
	// v5 approach).
	//
	// Background: the Cloudflare API always returns the full entries list on every GET.
	// Because `entries` is Computed+Optional, Terraform diffs state (API value) against
	// config (null/absent) and plans to remove the entries on every plan cycle — perpetual
	// drift. Users who do set `entries` explicitly in their config are unaffected: when
	// config is non-null, we leave the plan unchanged so their intent is respected.
	//
	// This hook runs at plan time (before apply), so it does not affect what is written
	// to state — it only prevents Terraform from proposing a removal it would immediately
	// undo on the next refresh.
	if req.Plan.Raw.IsNull() {
		// Resource is being destroyed — don't interfere.
		return
	}

	var configEntries, stateEntries customfield.NestedObjectList[ZeroTrustDLPPredefinedProfileEntriesModel]

	// Read config value for entries. If config is null/absent (user omitted entries),
	// copy the state value into the plan to suppress the spurious removal diff.
	configDiags := req.Config.GetAttribute(ctx, path.Root("entries"), &configEntries)
	resp.Diagnostics.Append(configDiags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if !configEntries.IsNull() {
		// User has entries in their config — respect it, don't interfere.
		return
	}

	// Config omits entries. Copy whatever is in state into the plan so Terraform
	// sees no diff for this attribute.
	stateDiags := req.State.GetAttribute(ctx, path.Root("entries"), &stateEntries)
	resp.Diagnostics.Append(stateDiags...)
	if resp.Diagnostics.HasError() {
		return
	}

	setPlanDiags := resp.Plan.SetAttribute(ctx, path.Root("entries"), stateEntries)
	resp.Diagnostics.Append(setPlanDiags...)
}

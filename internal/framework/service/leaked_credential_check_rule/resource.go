package leaked_credential_check_rule

import (
	"context"
	"fmt"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/framework/muxclient"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ resource.Resource                = &LeakedCredentialCheckRuleResource{}
	_ resource.ResourceWithImportState = &LeakedCredentialCheckRuleResource{}
)

func NewResource() resource.Resource {
	return &LeakedCredentialCheckRuleResource{}
}

type LeakedCredentialCheckRuleResource struct {
	client *muxclient.Client
}

func (r *LeakedCredentialCheckRuleResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_leaked_credential_check_rule"
}

func (r *LeakedCredentialCheckRuleResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*muxclient.Client)

	if !ok {
		resp.Diagnostics.AddError(
			"unexpected resource configure type",
			fmt.Sprintf("Expected *muxclient.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *LeakedCredentialCheckRuleResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data LeakedCredentialCheckRulesModel
	diags := req.Plan.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	detection, err := r.client.V1.LeakedCredentialCheckCreateDetection(ctx, cloudflare.ZoneIdentifier(data.ZoneID.ValueString()), cloudflare.LeakedCredentialCheckCreateDetectionParams{
		Username: data.Username.ValueString(),
		Password: data.Password.ValueString(),
	})
	if err != nil {
		resp.Diagnostics.AddError("Error creating a user-defined detection patter for Leaked Credential Check", err.Error())
		return
	}

	data.ID = types.StringValue(detection.ID)

	diags = resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *LeakedCredentialCheckRuleResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state LeakedCredentialCheckRulesModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	zoneID := state.ZoneID.ValueString()
	var foundRule cloudflare.LeakedCredentialCheckDetectionEntry
	rules, err := r.client.V1.LeakedCredentialCheckListDetections(ctx, cloudflare.ZoneIdentifier(zoneID), cloudflare.LeakedCredentialCheckListDetectionsParams{})
	if err != nil {
		resp.Diagnostics.AddError("Error listing Leaked Credential Check user-defined detection patterns", err.Error())
		return
	}

	// leaked credentials doens't offer a single get operation so
	// loop until we find the matching ID.
	for _, rule := range rules {
		if rule.ID == state.ID.ValueString() {
			foundRule = rule
			break
		}
	}

	state.Password = types.StringValue(foundRule.Password)
	state.Username = types.StringValue(foundRule.Username)
	state.ID = types.StringValue(foundRule.ID)
	state.ZoneID = types.StringValue(zoneID)

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

func (r *LeakedCredentialCheckRuleResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data LeakedCredentialCheckRulesModel
	diags := req.Plan.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	zoneID := cloudflare.ZoneIdentifier(data.ZoneID.ValueString())
	_, err := r.client.V1.LeakedCredentialCheckUpdateDetection(ctx, zoneID, cloudflare.LeakedCredentialCheckUpdateDetectionParams{
		LeakedCredentialCheckDetectionEntry: cloudflare.LeakedCredentialCheckDetectionEntry{
			ID:       data.ID.ValueString(),
			Username: data.Username.ValueString(),
			Password: data.Password.ValueString(),
		},
	})
	if err != nil {
		resp.Diagnostics.AddError("Error fetching Leaked Credential Check user-defined detection patterns", err.Error())
		return
	}

	diags = resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *LeakedCredentialCheckRuleResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state LeakedCredentialCheckRulesModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	zoneID := cloudflare.ZoneIdentifier(state.ZoneID.ValueString())
	deleteParam := cloudflare.LeakedCredentialCheckDeleteDetectionParams{DetectionID: state.ID.ValueString()}
	_, err := r.client.V1.LeakedCredentialCheckDeleteDetection(ctx, zoneID, deleteParam)
	if err != nil {
		resp.Diagnostics.AddError("Error deleting a user-defined detection patter for Leaked Credential Check", err.Error())
		return
	}
}

func (r *LeakedCredentialCheckRuleResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), req.ID)...)
	return
}

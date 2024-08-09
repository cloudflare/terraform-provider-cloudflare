package zero_trust_risk_score_integration

import (
	"context"
	"fmt"
	"strings"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/framework/muxclient"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &RiskScoreIntegrationResource{}

func NewResource() resource.Resource {
	return &RiskScoreIntegrationResource{}
}

// RiskScoreIntegrationResource defines the resource implementation.
type RiskScoreIntegrationResource struct {
	client *muxclient.Client
}

func (r *RiskScoreIntegrationResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_zero_trust_risk_score_integration"
}

func (r *RiskScoreIntegrationResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *RiskScoreIntegrationResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *RiskScoreIntegrationModel

	// Read Terraform plan data into the model.
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// When an integration is created, is it active
	if !data.Active.ValueBool() {
		resp.Diagnostics.AddError("failed to create risk score integration", "integration must be created with active=true")
		return
	}

	accountID := data.AccountID.ValueString()

	integration, err := r.client.V1.CreateRiskScoreIntegration(ctx, cloudflare.AccountIdentifier(accountID),
		cloudflare.RiskScoreIntegrationCreateRequest{
			IntegrationType: data.IntegrationType.ValueString(),
			TenantUrl:       data.TenantUrl.ValueString(),
			ReferenceID:     data.ReferenceID.ValueString(),
		},
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to create risk score integration", err.Error())
		return
	}

	data.ID = types.StringValue(integration.ID)
	data.ReferenceID = types.StringValue(integration.ReferenceID)
	data.Active = types.BoolPointerValue(integration.Active)
	data.WellKnownUrl = types.StringValue(integration.WellKnownUrl)
	// Save state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *RiskScoreIntegrationResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *RiskScoreIntegrationModel

	// Read Terraform state into the model.
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("Preparing to read Risk Score Integration: %+v", data))

	accountID := data.AccountID.ValueString()

	integration, err := r.client.V1.GetRiskScoreIntegration(ctx, cloudflare.AccountIdentifier(accountID), data.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("failed reading risk score integration", err.Error())
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("Read Risk Score Integration: %+v", data))

	data.IntegrationType = types.StringValue(integration.IntegrationType)
	data.TenantUrl = types.StringValue(integration.TenantUrl)
	data.ReferenceID = types.StringValue(integration.ReferenceID)
	data.Active = types.BoolPointerValue(integration.Active)
	data.WellKnownUrl = types.StringValue(integration.WellKnownUrl)
	// Save state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *RiskScoreIntegrationResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan_data *RiskScoreIntegrationModel
	var state_data *RiskScoreIntegrationModel

	// Read Terraform state & plan data into the model.
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan_data)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &state_data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	accountID := plan_data.AccountID.ValueString()

	tflog.Debug(ctx, fmt.Sprintf("Updating Risk Score Integration from struct: %+v", plan_data))

	integration, err := r.client.V1.UpdateRiskScoreIntegration(ctx, cloudflare.AccountIdentifier(accountID), state_data.ID.ValueString(),
		cloudflare.RiskScoreIntegrationUpdateRequest{
			IntegrationType: plan_data.IntegrationType.ValueString(),
			TenantUrl:       plan_data.TenantUrl.ValueString(),
			ReferenceID:     plan_data.ReferenceID.ValueString(),
			Active:          plan_data.Active.ValueBoolPointer(),
		},
	)
	if err != nil {
		resp.Diagnostics.AddError("failed updating risk score integration", err.Error())
		return
	}

	plan_data.ID = types.StringValue(integration.ID)
	plan_data.ReferenceID = types.StringValue(integration.ReferenceID)
	plan_data.WellKnownUrl = types.StringValue(integration.WellKnownUrl)
	// Save state
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan_data)...)
}

func (r *RiskScoreIntegrationResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *RiskScoreIntegrationModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	accountID := data.AccountID.ValueString()

	err := r.client.V1.DeleteRiskScoreIntegration(ctx, cloudflare.AccountIdentifier(accountID), data.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("failed to delete risk score integration", err.Error())
		return
	}

	// Delete state
	// We don't call "resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)"
	// This means that the state does not get set (and is deleted)
}

func (r *RiskScoreIntegrationResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idparts := strings.Split(req.ID, "/")
	if len(idparts) != 2 {
		resp.Diagnostics.AddError("error importing Risk Score Integration", "invalid ID specified. Please specify the ID as \"account_id/integration_id\"")
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(
		ctx, path.Root("account_id"), idparts[0],
	)...)
	resp.Diagnostics.Append(resp.State.SetAttribute(
		ctx, path.Root("id"), idparts[1],
	)...)
}

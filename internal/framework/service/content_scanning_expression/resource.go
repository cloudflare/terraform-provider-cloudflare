package content_scanning_expression

import (
	"context"
	"fmt"
	"strings"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/framework/muxclient"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ resource.Resource                = &ContentScanningExpressionResource{}
	_ resource.ResourceWithImportState = &ContentScanningExpressionResource{}
)

func NewResource() resource.Resource {
	return &ContentScanningExpressionResource{}
}

type ContentScanningExpressionResource struct {
	client *muxclient.Client
}

func (r *ContentScanningExpressionResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_content_scanning_expression"
}

func (r *ContentScanningExpressionResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *ContentScanningExpressionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data ContentScanningExpressionModel
	diags := req.Plan.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	params := cloudflare.ContentScanningAddCustomExpressionsParams{
		Payloads: []cloudflare.ContentScanningCustomPayload{
			{
				Payload: data.Payload.ValueString(),
			},
		},
	}
	expressions, err := r.client.V1.ContentScanningAddCustomExpressions(ctx, cloudflare.ZoneIdentifier(data.ZoneID.ValueString()), params)
	if err != nil {
		resp.Diagnostics.AddError("Error creating a custom scan expression for Content Scanning", err.Error())
		return
	}

	// The Add API returns a list of all exiting custom scan expression
	// loop until we find the newly created one, matching on payload
	// payload uniqueness is enforced by the service
	for _, exp := range expressions {
		if exp.Payload == data.Payload.ValueString() {
			data.ID = types.StringValue(exp.ID)
			break
		}
	}

	diags = resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *ContentScanningExpressionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state ContentScanningExpressionModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	zoneID := state.ZoneID.ValueString()
	var foundExp cloudflare.ContentScanningCustomExpression
	expressions, err := r.client.V1.ContentScanningListCustomExpressions(ctx, cloudflare.ZoneIdentifier(zoneID), cloudflare.ContentScanningListCustomExpressionsParams{})
	if err != nil {
		resp.Diagnostics.AddError("Error listing customs scan expressions for Content Scanning", err.Error())
		return
	}

	// content scanning doens't offer a single get operation so
	// loop until we find the matching ID.
	for _, exp := range expressions {
		if exp.ID == state.ID.ValueString() {
			foundExp = exp
			break
		}
	}

	state.ID = types.StringValue(foundExp.ID)
	state.Payload = types.StringValue(foundExp.Payload)

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

func (r *ContentScanningExpressionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan ContentScanningExpressionModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var state ContentScanningExpressionModel
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	zoneID := cloudflare.ZoneIdentifier(plan.ZoneID.ValueString())
	plan.ID = state.ID

	// API does not offer an update operation so we use delete/create
	delParams := cloudflare.ContentScanningDeleteCustomExpressionsParams{ID: plan.ID.ValueString()}
	_, err := r.client.V1.ContentScanningDeleteCustomExpression(ctx, zoneID, delParams)
	if err != nil {
		resp.Diagnostics.AddError("Error in Update while deleting custom scan expression for Content Scanning", err.Error())
		return
	}
	createParams := cloudflare.ContentScanningAddCustomExpressionsParams{
		Payloads: []cloudflare.ContentScanningCustomPayload{
			{
				Payload: plan.Payload.ValueString(),
			},
		},
	}
	expressions, err := r.client.V1.ContentScanningAddCustomExpressions(ctx, zoneID, createParams)
	if err != nil {
		resp.Diagnostics.AddError("Error in Update while creating a custom scan expression for Content Scanning", err.Error())
		return
	}

	// The Add API returns a list of all exiting custom scan expression
	// loop until we find the newly created one, matching on payload
	// payload uniqueness is enforced by the service
	for _, exp := range expressions {
		if exp.Payload == plan.Payload.ValueString() {
			plan.ID = types.StringValue(exp.ID)
			break
		}
	}

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r *ContentScanningExpressionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state ContentScanningExpressionModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	zoneID := cloudflare.ZoneIdentifier(state.ZoneID.ValueString())
	deleteParam := cloudflare.ContentScanningDeleteCustomExpressionsParams{ID: state.ID.ValueString()}
	_, err := r.client.V1.ContentScanningDeleteCustomExpression(ctx, zoneID, deleteParam)
	if err != nil {
		resp.Diagnostics.AddError("Error deleting custom scan expression for Content Scanning", err.Error())
		return
	}
}

func (r *ContentScanningExpressionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idparts := strings.Split(req.ID, "/")
	if len(idparts) != 2 {
		resp.Diagnostics.AddError("error importing content scanning custom expression", "invalid ID specified. Please specify the ID as \"zone_id/resource_id\"")
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(
		ctx, path.Root("zone_id"), idparts[0],
	)...)
	resp.Diagnostics.Append(resp.State.SetAttribute(
		ctx, path.Root("id"), idparts[1],
	)...)
}

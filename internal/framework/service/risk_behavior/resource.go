package risk_behavior

import (
	"context"
	"fmt"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/framework/muxclient"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &RiskBehaviorResource{}

func NewResource() resource.Resource {
	return &RiskBehaviorResource{}
}

// RiskBehaviorResource defines the resource implementation.
type RiskBehaviorResource struct {
	client *muxclient.Client
}

func (r *RiskBehaviorResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_risk_behavior"
}

func (r *RiskBehaviorResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *RiskBehaviorResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *RiskBehaviorModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	accountId := data.AccountID.ValueString()

	behaviorsMap, err := ConvertBehaviorsTtoC(data.Behaviors)
	if err != nil {
		resp.Diagnostics.AddError("invalid risk level", err.Error())
		return
	}

	behaviors, err := r.client.V1.UpdateBehaviors(ctx, accountId,
		cloudflare.Behaviors{Behaviors: behaviorsMap},
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to create risk behaviors", err.Error())
		return
	}

	retained := map[string]cloudflare.Behavior{}
	for k, b := range behaviors.Behaviors {
		_, ok := behaviorsMap[k]
		if ok {
			retained[k] = b
		}
	}
	behaviorsSet := ConvertBehaviorsCtoT(retained)

	data.AccountID = types.StringValue(accountId)
	data.Behaviors = behaviorsSet
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *RiskBehaviorResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *RiskBehaviorModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	accountId := data.AccountID.ValueString()

	behaviors, err := r.client.V1.Behaviors(ctx, accountId)
	if err != nil {
		resp.Diagnostics.AddError("failed reading risk behaviors", err.Error())
		return
	}

	retained := map[string]cloudflare.Behavior{}
	for k, b := range behaviors.Behaviors {
		if containsBehavior(data.Behaviors, k) {
			retained[k] = b
		}
	}
	behaviorsSet := ConvertBehaviorsCtoT(retained)

	data.AccountID = types.StringValue(accountId)
	data.Behaviors = behaviorsSet
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func containsBehavior(s []RiskBehaviorBehaviorModel, n string) bool {
	for _, a := range s {
		if a.Name.ValueString() == n {
			return true
		}
	}
	return false
}

func ConvertBehaviorsTtoC(b []RiskBehaviorBehaviorModel) (map[string]cloudflare.Behavior, error) {
	behaviorsMap := map[string]cloudflare.Behavior{}
	for _, b := range b {
		riskLevel, err := cloudflare.RiskLevelFromString(b.RiskLevel.ValueString())
		if err != nil {
			return nil, err
		}

		enabled := b.Enabled.ValueBool()

		behavior := cloudflare.Behavior{
			Enabled:   &enabled,
			RiskLevel: *riskLevel,
		}

		behaviorsMap[b.Name.ValueString()] = behavior
	}

	return behaviorsMap, nil
}

func ConvertBehaviorsCtoT(b map[string]cloudflare.Behavior) []RiskBehaviorBehaviorModel {
	behaviorsSet := []RiskBehaviorBehaviorModel{}

	for k, b := range b {
		behavior := RiskBehaviorBehaviorModel{
			Enabled:   types.BoolPointerValue(b.Enabled),
			Name:      types.StringValue(k),
			RiskLevel: types.StringValue(fmt.Sprint(b.RiskLevel)),
		}

		behaviorsSet = append(behaviorsSet, behavior)
	}

	return behaviorsSet
}

func (r *RiskBehaviorResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *RiskBehaviorModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	accountId := data.AccountID.ValueString()

	behaviorsMap, err := ConvertBehaviorsTtoC(data.Behaviors)
	if err != nil {
		resp.Diagnostics.AddError("invalid risk level", err.Error())
		return
	}

	behaviors, err := r.client.V1.UpdateBehaviors(ctx, accountId,
		cloudflare.Behaviors{Behaviors: behaviorsMap},
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to update risk behaviors", err.Error())
		return
	}

	retained := map[string]cloudflare.Behavior{}
	for k, b := range behaviors.Behaviors {
		_, ok := behaviorsMap[k]
		if ok {
			retained[k] = b
		}
	}

	behaviorsSet := ConvertBehaviorsCtoT(retained)

	data.AccountID = types.StringValue(accountId)
	data.Behaviors = behaviorsSet
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *RiskBehaviorResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *RiskBehaviorModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// tflog.Debug(ctx, "Resetting all zero trust risk behaviors to enabled: false, risk_level: low")

	behaviors, err := r.client.V1.Behaviors(ctx, data.AccountID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("failed to get risk behaviors", err.Error())
		return
	}

	// set all risk behavior values to false/low before running update
	for _, behavior := range behaviors.Behaviors {
		behavior.Enabled = cloudflare.BoolPtr(false)
		behavior.RiskLevel = cloudflare.Low
	}

	_, err = r.client.V1.UpdateBehaviors(ctx, data.AccountID.ValueString(), behaviors)
	if err != nil {
		resp.Diagnostics.AddError("failed to reset risk behaviors", err.Error())
		return
	}
}

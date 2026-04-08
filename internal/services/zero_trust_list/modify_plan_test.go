package zero_trust_list

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

// buildState constructs a tfsdk.State from a ZeroTrustListModel using the
// resource schema. Items may be nil (resource has no items configured).
func buildState(t *testing.T, items *[]*ZeroTrustListItemsModel) tfsdk.State {
	t.Helper()
	ctx := context.Background()
	schema := ResourceSchema(ctx)
	model := &ZeroTrustListModel{
		ID:        types.StringValue("test-id"),
		AccountID: types.StringValue("test-account"),
		Type:      types.StringValue("SERIAL"),
		Name:      types.StringValue("test-list"),
		Items:     items,
	}
	state := tfsdk.State{Schema: schema}
	diags := state.Set(ctx, model)
	if diags.HasError() {
		t.Fatalf("failed to set state: %v", diags)
	}
	return state
}

// buildPlan constructs a tfsdk.Plan from a ZeroTrustListModel using the
// resource schema. Items may be nil (resource has no items configured).
func buildPlan(t *testing.T, items *[]*ZeroTrustListItemsModel) tfsdk.Plan {
	t.Helper()
	s := buildState(t, items)
	return tfsdk.Plan{Schema: s.Schema, Raw: s.Raw}
}

// buildNullState returns a null tfsdk.State (simulates first-apply / destroy).
func buildNullState(t *testing.T) tfsdk.State {
	t.Helper()
	ctx := context.Background()
	schema := ResourceSchema(ctx)
	return tfsdk.State{
		Schema: schema,
		Raw:    tftypes.NewValue(schema.Type().TerraformType(ctx), nil),
	}
}

// buildNullPlan returns a null tfsdk.Plan (simulates destroy).
func buildNullPlan(t *testing.T) tfsdk.Plan {
	t.Helper()
	ctx := context.Background()
	schema := ResourceSchema(ctx)
	return tfsdk.Plan{
		Schema: schema,
		Raw:    tftypes.NewValue(schema.Type().TerraformType(ctx), nil),
	}
}

// TestModifyPlan_NullPlanReturnsEarly verifies that ModifyPlan is a no-op on destroy
// (null plan), consistent with the req.Plan.Raw.IsNull() guard.
func TestModifyPlan_NullPlanReturnsEarly(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	r := &ZeroTrustListResource{}
	req := resource.ModifyPlanRequest{
		State: buildState(t, &[]*ZeroTrustListItemsModel{
			{Value: types.StringValue("10.0.0.1")},
		}),
		Plan: buildNullPlan(t),
	}
	resp := &resource.ModifyPlanResponse{Plan: req.Plan}

	r.ModifyPlan(ctx, req, resp)

	if resp.Diagnostics.HasError() {
		t.Errorf("expected no error on null plan, got: %v", resp.Diagnostics)
	}
}

// TestModifyPlan_NullStateReturnsEarly verifies that ModifyPlan is a no-op on first
// create (null state), consistent with the req.State.Raw.IsNull() guard.
func TestModifyPlan_NullStateReturnsEarly(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	r := &ZeroTrustListResource{}
	req := resource.ModifyPlanRequest{
		State: buildNullState(t),
		Plan:  buildPlan(t, &[]*ZeroTrustListItemsModel{{Value: types.StringValue("10.0.0.1")}}),
	}
	resp := &resource.ModifyPlanResponse{Plan: req.Plan}

	r.ModifyPlan(ctx, req, resp)

	if resp.Diagnostics.HasError() {
		t.Errorf("expected no error on null state, got: %v", resp.Diagnostics)
	}
}

// TestModifyPlan_SameItemsSuppressesDiff verifies that identical plan and state items
// result in the plan being normalized to state (no-op diff suppression).
func TestModifyPlan_SameItemsSuppressesDiff(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	items := &[]*ZeroTrustListItemsModel{
		{Value: types.StringValue("10.0.0.1")},
		{Value: types.StringValue("10.0.0.2")},
	}

	stateObj := buildState(t, items)
	req := resource.ModifyPlanRequest{
		State: stateObj,
		Plan:  buildPlan(t, items),
	}
	resp := &resource.ModifyPlanResponse{Plan: req.Plan}

	r := &ZeroTrustListResource{}
	r.ModifyPlan(ctx, req, resp)

	if resp.Diagnostics.HasError() {
		t.Fatalf("unexpected error: %v", resp.Diagnostics)
	}
	// Plan raw should equal state raw — items normalized to state after suppression.
	if !resp.Plan.Raw.Equal(stateObj.Raw) {
		t.Error("expected plan to be normalized to state after diff suppression")
	}
}

// TestModifyPlan_DifferentItemsDoesNotSuppressDiff verifies that when plan and state
// items differ, ModifyPlan does NOT overwrite the plan — the diff is preserved.
func TestModifyPlan_DifferentItemsDoesNotSuppressDiff(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	stateObj := buildState(t, &[]*ZeroTrustListItemsModel{
		{Value: types.StringValue("10.0.0.1")},
	})
	planObj := buildPlan(t, &[]*ZeroTrustListItemsModel{
		{Value: types.StringValue("10.0.0.2")}, // different
	})
	originalPlanRaw := planObj.Raw

	req := resource.ModifyPlanRequest{
		State: stateObj,
		Plan:  planObj,
	}
	resp := &resource.ModifyPlanResponse{Plan: req.Plan}

	r := &ZeroTrustListResource{}
	r.ModifyPlan(ctx, req, resp)

	if resp.Diagnostics.HasError() {
		t.Fatalf("unexpected error: %v", resp.Diagnostics)
	}
	// Plan raw must be unchanged — ModifyPlan must not suppress a real diff.
	if !resp.Plan.Raw.Equal(originalPlanRaw) {
		t.Error("ModifyPlan must not suppress a diff when items actually changed")
	}
}

// TestModifyPlan_NilPlanItemsSkipsComparison verifies that when plan has no items
// (nil), ModifyPlan does not attempt a hash comparison or panic.
func TestModifyPlan_NilPlanItemsSkipsComparison(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	req := resource.ModifyPlanRequest{
		State: buildState(t, &[]*ZeroTrustListItemsModel{{Value: types.StringValue("10.0.0.1")}}),
		Plan:  buildPlan(t, nil), // removing all items
	}
	resp := &resource.ModifyPlanResponse{Plan: req.Plan}

	r := &ZeroTrustListResource{}
	r.ModifyPlan(ctx, req, resp)

	if resp.Diagnostics.HasError() {
		t.Errorf("unexpected error when plan items are nil: %v", resp.Diagnostics)
	}
}

// TestModifyPlan_NilStateItemsSkipsComparison verifies that when state has no items
// (nil), ModifyPlan does not attempt a hash comparison or panic.
func TestModifyPlan_NilStateItemsSkipsComparison(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	req := resource.ModifyPlanRequest{
		State: buildState(t, nil), // first time adding items
		Plan:  buildPlan(t, &[]*ZeroTrustListItemsModel{{Value: types.StringValue("10.0.0.1")}}),
	}
	resp := &resource.ModifyPlanResponse{Plan: req.Plan}

	r := &ZeroTrustListResource{}
	r.ModifyPlan(ctx, req, resp)

	if resp.Diagnostics.HasError() {
		t.Errorf("unexpected error when state items are nil: %v", resp.Diagnostics)
	}
}

// TestModifyPlan_ReorderedItemsSuppressesDiff verifies that items in a different
// order in the plan are treated as equivalent (set semantics) and the plan is
// normalized to the state order.
func TestModifyPlan_ReorderedItemsSuppressesDiff(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	stateItems := &[]*ZeroTrustListItemsModel{
		{Value: types.StringValue("10.0.0.1")},
		{Value: types.StringValue("10.0.0.2")},
		{Value: types.StringValue("10.0.0.3")},
	}
	planItems := &[]*ZeroTrustListItemsModel{
		{Value: types.StringValue("10.0.0.3")},
		{Value: types.StringValue("10.0.0.1")},
		{Value: types.StringValue("10.0.0.2")},
	}

	stateObj := buildState(t, stateItems)
	req := resource.ModifyPlanRequest{
		State: stateObj,
		Plan:  buildPlan(t, planItems),
	}
	resp := &resource.ModifyPlanResponse{Plan: req.Plan}

	r := &ZeroTrustListResource{}
	r.ModifyPlan(ctx, req, resp)

	if resp.Diagnostics.HasError() {
		t.Fatalf("unexpected error: %v", resp.Diagnostics)
	}
	// Plan raw should equal state raw — reorder diff suppressed, normalized to state.
	if !resp.Plan.Raw.Equal(stateObj.Raw) {
		t.Error("expected plan to be normalized to state after reorder suppression")
	}
}

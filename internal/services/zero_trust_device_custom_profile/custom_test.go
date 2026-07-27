package zero_trust_device_custom_profile

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestNormalizeSplitTunnelList_ResolvesUnknownToNull(t *testing.T) {
	ctx := context.Background()

	// Build a plan list with one entry where host and description are unknown
	// (simulates user setting only address).
	planEntries := []ZeroTrustDeviceCustomProfileExcludeModel{
		{
			Address:     types.StringValue("10.0.0.0/8"),
			Description: types.StringUnknown(),
			Host:        types.StringUnknown(),
		},
	}
	planList := customfield.NewObjectListMust(ctx, planEntries)

	// State list is empty (new element, no prior state).
	stateList := customfield.NullObjectList[ZeroTrustDeviceCustomProfileExcludeModel](ctx)

	result, diags := normalizeSplitTunnelList(ctx, planList, stateList)
	if diags.HasError() {
		t.Fatalf("unexpected diagnostics: %v", diags)
	}

	if result.IsNull() || result.IsUnknown() {
		t.Fatal("expected non-null, non-unknown result")
	}

	elements := result.Elements()
	if len(elements) != 1 {
		t.Fatalf("expected 1 element, got %d", len(elements))
	}

	// Extract the normalized entry.
	var entries []ZeroTrustDeviceCustomProfileExcludeModel
	diags = result.ElementsAs(ctx, &entries, false)
	if diags.HasError() {
		t.Fatalf("failed to extract entries: %v", diags)
	}

	entry := entries[0]
	if entry.Address.ValueString() != "10.0.0.0/8" {
		t.Errorf("expected address '10.0.0.0/8', got '%s'", entry.Address.ValueString())
	}
	if !entry.Host.IsNull() {
		t.Errorf("expected host to be null, got '%s' (unknown=%v)", entry.Host.ValueString(), entry.Host.IsUnknown())
	}
	if !entry.Description.IsNull() {
		t.Errorf("expected description to be null, got '%s' (unknown=%v)", entry.Description.ValueString(), entry.Description.IsUnknown())
	}
}

func TestNormalizeSplitTunnelList_PreservesStateForExistingEntries(t *testing.T) {
	ctx := context.Background()

	// Plan has an existing entry with unknown host (user only set address).
	planEntries := []ZeroTrustDeviceCustomProfileExcludeModel{
		{
			Address:     types.StringValue("10.0.0.0/8"),
			Description: types.StringUnknown(),
			Host:        types.StringUnknown(),
		},
	}
	planList := customfield.NewObjectListMust(ctx, planEntries)

	// State has the same entry with null host and description (from API response).
	stateEntries := []ZeroTrustDeviceCustomProfileExcludeModel{
		{
			Address:     types.StringValue("10.0.0.0/8"),
			Description: types.StringNull(),
			Host:        types.StringNull(),
		},
	}
	stateList := customfield.NewObjectListMust(ctx, stateEntries)

	result, diags := normalizeSplitTunnelList(ctx, planList, stateList)
	if diags.HasError() {
		t.Fatalf("unexpected diagnostics: %v", diags)
	}

	var entries []ZeroTrustDeviceCustomProfileExcludeModel
	diags = result.ElementsAs(ctx, &entries, false)
	if diags.HasError() {
		t.Fatalf("failed to extract entries: %v", diags)
	}

	entry := entries[0]
	if entry.Address.ValueString() != "10.0.0.0/8" {
		t.Errorf("expected address '10.0.0.0/8', got '%s'", entry.Address.ValueString())
	}
	if !entry.Host.IsNull() {
		t.Errorf("expected host to be null (from state), got unknown=%v value='%s'", entry.Host.IsUnknown(), entry.Host.ValueString())
	}
	if !entry.Description.IsNull() {
		t.Errorf("expected description to be null (from state), got unknown=%v value='%s'", entry.Description.IsUnknown(), entry.Description.ValueString())
	}
}

func TestNormalizeSplitTunnelList_MixedNewAndExistingEntries(t *testing.T) {
	ctx := context.Background()

	// Plan: existing entry (index 0) + new entry (index 1).
	planEntries := []ZeroTrustDeviceCustomProfileExcludeModel{
		{
			Address:     types.StringValue("10.0.0.0/8"),
			Description: types.StringUnknown(),
			Host:        types.StringUnknown(),
		},
		{
			Address:     types.StringUnknown(),
			Description: types.StringValue("Company domain"),
			Host:        types.StringValue("*.example.com"),
		},
	}
	planList := customfield.NewObjectListMust(ctx, planEntries)

	// State only has 1 entry.
	stateEntries := []ZeroTrustDeviceCustomProfileExcludeModel{
		{
			Address:     types.StringValue("10.0.0.0/8"),
			Description: types.StringNull(),
			Host:        types.StringNull(),
		},
	}
	stateList := customfield.NewObjectListMust(ctx, stateEntries)

	result, diags := normalizeSplitTunnelList(ctx, planList, stateList)
	if diags.HasError() {
		t.Fatalf("unexpected diagnostics: %v", diags)
	}

	var entries []ZeroTrustDeviceCustomProfileExcludeModel
	diags = result.ElementsAs(ctx, &entries, false)
	if diags.HasError() {
		t.Fatalf("failed to extract entries: %v", diags)
	}

	if len(entries) != 2 {
		t.Fatalf("expected 2 entries, got %d", len(entries))
	}

	// Entry 0: existing, unknown fields resolved from state (null).
	if !entries[0].Host.IsNull() {
		t.Errorf("entry[0].host: expected null, got unknown=%v value='%s'", entries[0].Host.IsUnknown(), entries[0].Host.ValueString())
	}
	if !entries[0].Description.IsNull() {
		t.Errorf("entry[0].description: expected null, got unknown=%v", entries[0].Description.IsUnknown())
	}

	// Entry 1: new, unknown address resolved to null (no state at index 1).
	if !entries[1].Address.IsNull() {
		t.Errorf("entry[1].address: expected null, got unknown=%v value='%s'", entries[1].Address.IsUnknown(), entries[1].Address.ValueString())
	}
	if entries[1].Host.ValueString() != "*.example.com" {
		t.Errorf("entry[1].host: expected '*.example.com', got '%s'", entries[1].Host.ValueString())
	}
	if entries[1].Description.ValueString() != "Company domain" {
		t.Errorf("entry[1].description: expected 'Company domain', got '%s'", entries[1].Description.ValueString())
	}
}

func TestNormalizeSplitTunnelList_NullAndUnknownListsPassThrough(t *testing.T) {
	ctx := context.Background()

	nullList := customfield.NullObjectList[ZeroTrustDeviceCustomProfileExcludeModel](ctx)
	unknownList := customfield.UnknownObjectList[ZeroTrustDeviceCustomProfileExcludeModel](ctx)
	emptyState := customfield.NullObjectList[ZeroTrustDeviceCustomProfileExcludeModel](ctx)

	result, diags := normalizeSplitTunnelList(ctx, nullList, emptyState)
	if diags.HasError() {
		t.Fatalf("unexpected diagnostics: %v", diags)
	}
	if !result.IsNull() {
		t.Error("expected null list to pass through")
	}

	result, diags = normalizeSplitTunnelList(ctx, unknownList, emptyState)
	if diags.HasError() {
		t.Fatalf("unexpected diagnostics: %v", diags)
	}
	if !result.IsUnknown() {
		t.Error("expected unknown list to pass through")
	}
}

func TestNormalizeSplitTunnelList_NoChangeWhenAllKnown(t *testing.T) {
	ctx := context.Background()

	// All attributes are known -- nothing to normalize.
	planEntries := []ZeroTrustDeviceCustomProfileExcludeModel{
		{
			Address:     types.StringValue("10.0.0.0/8"),
			Description: types.StringValue("desc"),
			Host:        types.StringNull(),
		},
	}
	planList := customfield.NewObjectListMust(ctx, planEntries)
	stateList := customfield.NullObjectList[ZeroTrustDeviceCustomProfileExcludeModel](ctx)

	result, diags := normalizeSplitTunnelList(ctx, planList, stateList)
	if diags.HasError() {
		t.Fatalf("unexpected diagnostics: %v", diags)
	}

	// Should return the original list unchanged.
	if !result.Equal(planList) {
		t.Error("expected unchanged list when all attributes are known")
	}
}

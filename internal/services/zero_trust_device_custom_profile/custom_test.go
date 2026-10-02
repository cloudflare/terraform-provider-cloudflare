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

// TestNormalizeSplitTunnelList_HostEntryDoesNotInheritAddressFromState reproduces
// GH-7307: when a host entry is inserted before existing address entries, plan[i]
// has host=known and address=unknown while state[i] is the displaced address entry.
// The code must NOT inherit the state address into the host entry — that would
// send both fields to the API, triggering error 2049 ("host and Address both cannot
// be present").
func TestNormalizeSplitTunnelList_HostEntryDoesNotInheritAddressFromState(t *testing.T) {
	ctx := context.Background()

	// Plan: host entry prepended in front of an existing address entry.
	// plan[0] = new host entry — address is unknown (optional+computed, not set by user).
	// plan[1] = existing address entry — host is unknown.
	planEntries := []ZeroTrustDeviceCustomProfileIncludeModel{
		{
			Host:        types.StringValue("example.com"),
			Address:     types.StringUnknown(),
			Description: types.StringValue("Example host"),
		},
		{
			Address:     types.StringValue("172.64.128.0/20"),
			Host:        types.StringUnknown(),
			Description: types.StringValue("Gateway initial resolved IPs"),
		},
	}
	planList := customfield.NewObjectListMust(ctx, planEntries)

	// State: before the edit, only the address entry existed at index 0.
	stateEntries := []ZeroTrustDeviceCustomProfileIncludeModel{
		{
			Address:     types.StringValue("172.64.128.0/20"),
			Host:        types.StringNull(),
			Description: types.StringValue("Gateway initial resolved IPs"),
		},
	}
	stateList := customfield.NewObjectListMust(ctx, stateEntries)

	result, diags := normalizeSplitTunnelList(ctx, planList, stateList)
	if diags.HasError() {
		t.Fatalf("unexpected diagnostics: %v", diags)
	}

	var entries []ZeroTrustDeviceCustomProfileIncludeModel
	diags = result.ElementsAs(ctx, &entries, false)
	if diags.HasError() {
		t.Fatalf("failed to extract entries: %v", diags)
	}

	if len(entries) != 2 {
		t.Fatalf("expected 2 entries, got %d", len(entries))
	}

	// entries[0] is the new host entry. address must be null — NOT inherited from
	// the displaced state[0] address entry, because host is already set.
	e0 := entries[0]
	if e0.Host.ValueString() != "example.com" {
		t.Errorf("entries[0].host: expected 'example.com', got %q", e0.Host.ValueString())
	}
	if !e0.Address.IsNull() {
		t.Errorf("entries[0].address: expected null (host is set — must not inherit state address), got unknown=%v value=%q",
			e0.Address.IsUnknown(), e0.Address.ValueString())
	}

	// entries[1] is the existing address entry. host must be null (from state or
	// null fallback — either is correct since state[1] doesn't exist).
	e1 := entries[1]
	if e1.Address.ValueString() != "172.64.128.0/20" {
		t.Errorf("entries[1].address: expected '172.64.128.0/20', got %q", e1.Address.ValueString())
	}
	if !e1.Host.IsNull() {
		t.Errorf("entries[1].host: expected null, got unknown=%v value=%q", e1.Host.IsUnknown(), e1.Host.ValueString())
	}
}

// TestNormalizeSplitTunnelList_AddressEntryDoesNotInheritHostFromState is the
// symmetric case: an address entry inserted before an existing host entry must
// not inherit the displaced host value.
func TestNormalizeSplitTunnelList_AddressEntryDoesNotInheritHostFromState(t *testing.T) {
	ctx := context.Background()

	// Plan: address entry prepended in front of an existing host entry.
	planEntries := []ZeroTrustDeviceCustomProfileExcludeModel{
		{
			Address:     types.StringValue("10.0.0.0/8"),
			Host:        types.StringUnknown(),
			Description: types.StringValue("Corporate network"),
		},
		{
			Host:        types.StringValue("example.com"),
			Address:     types.StringUnknown(),
			Description: types.StringValue("Corporate app"),
		},
	}
	planList := customfield.NewObjectListMust(ctx, planEntries)

	// State: before the edit, only the host entry existed at index 0.
	stateEntries := []ZeroTrustDeviceCustomProfileExcludeModel{
		{
			Host:        types.StringValue("example.com"),
			Address:     types.StringNull(),
			Description: types.StringValue("Corporate app"),
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

	// entries[0] is the new address entry. host must be null — NOT inherited from
	// the displaced state[0] host entry.
	e0 := entries[0]
	if e0.Address.ValueString() != "10.0.0.0/8" {
		t.Errorf("entries[0].address: expected '10.0.0.0/8', got %q", e0.Address.ValueString())
	}
	if !e0.Host.IsNull() {
		t.Errorf("entries[0].host: expected null (address is set — must not inherit state host), got unknown=%v value=%q",
			e0.Host.IsUnknown(), e0.Host.ValueString())
	}

	// entries[1] is the existing host entry.
	e1 := entries[1]
	if e1.Host.ValueString() != "example.com" {
		t.Errorf("entries[1].host: expected 'example.com', got %q", e1.Host.ValueString())
	}
	if !e1.Address.IsNull() {
		t.Errorf("entries[1].address: expected null, got unknown=%v value=%q", e1.Address.IsUnknown(), e1.Address.ValueString())
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

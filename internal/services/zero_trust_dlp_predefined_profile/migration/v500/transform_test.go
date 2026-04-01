package v500_test

import (
	"context"
	"testing"

	v500 "github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_dlp_predefined_profile/migration/v500"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// TestTransform_AllEntriesDisabled verifies that when a v4 predefined profile has entry
// blocks but all are disabled, the upgraded state has enabled_entries set to an empty
// slice (not nil). This prevents a spurious `+ enabled_entries = []` plan diff on the
// first terraform plan after the moved block fires the state upgrade.
func TestTransform_AllEntriesDisabled(t *testing.T) {
	ctx := context.Background()

	source := v500.SourceCloudflareDLPProfileModel{
		ID:        types.StringValue("profile-uuid"),
		AccountID: types.StringValue("account-uuid"),
		Name:      types.StringValue("Credentials and Secrets"),
		Type:      types.StringValue("predefined"),
		Entry: []v500.SourceEntryModel{
			{ID: types.StringValue("entry-1"), Name: types.StringValue("Entry 1"), Enabled: types.BoolValue(false), Pattern: nil},
			{ID: types.StringValue("entry-2"), Name: types.StringValue("Entry 2"), Enabled: types.BoolValue(false), Pattern: nil},
			{ID: types.StringValue("entry-3"), Name: types.StringValue("Entry 3"), Enabled: types.BoolValue(false), Pattern: nil},
		},
		AllowedMatchCount: types.Int64Value(0),
		OCREnabled:        types.BoolValue(false),
	}

	result, diags := v500.Transform(ctx, source)
	if diags.HasError() {
		t.Fatalf("Transform returned unexpected errors: %v", diags)
	}

	// EnabledEntries must be a non-nil pointer to an empty slice — not nil.
	// A nil pointer would cause `+ enabled_entries = []` drift on the first plan
	// after state upgrade when the config has `enabled_entries = []`.
	if result.EnabledEntries == nil {
		t.Fatal("EnabledEntries is nil; expected a pointer to an empty slice for all-disabled case")
	}
	if len(*result.EnabledEntries) != 0 {
		t.Fatalf("EnabledEntries should be empty; got %v", *result.EnabledEntries)
	}
}

// TestTransform_SomeEntriesEnabled verifies that enabled entry IDs are correctly
// extracted and disabled entries are omitted.
func TestTransform_SomeEntriesEnabled(t *testing.T) {
	ctx := context.Background()

	source := v500.SourceCloudflareDLPProfileModel{
		ID:        types.StringValue("profile-uuid"),
		AccountID: types.StringValue("account-uuid"),
		Name:      types.StringValue("Test Profile"),
		Type:      types.StringValue("predefined"),
		Entry: []v500.SourceEntryModel{
			{ID: types.StringValue("entry-enabled-1"), Name: types.StringValue("E1"), Enabled: types.BoolValue(true), Pattern: nil},
			{ID: types.StringValue("entry-disabled"), Name: types.StringValue("E2"), Enabled: types.BoolValue(false), Pattern: nil},
			{ID: types.StringValue("entry-enabled-2"), Name: types.StringValue("E3"), Enabled: types.BoolValue(true), Pattern: nil},
		},
		AllowedMatchCount: types.Int64Value(0),
		OCREnabled:        types.BoolValue(false),
	}

	result, diags := v500.Transform(ctx, source)
	if diags.HasError() {
		t.Fatalf("Transform returned unexpected errors: %v", diags)
	}

	if result.EnabledEntries == nil {
		t.Fatal("EnabledEntries is nil; expected non-nil")
	}
	if len(*result.EnabledEntries) != 2 {
		t.Fatalf("Expected 2 enabled entries, got %d: %v", len(*result.EnabledEntries), *result.EnabledEntries)
	}
	if (*result.EnabledEntries)[0].ValueString() != "entry-enabled-1" {
		t.Errorf("First enabled entry should be 'entry-enabled-1', got %q", (*result.EnabledEntries)[0].ValueString())
	}
	if (*result.EnabledEntries)[1].ValueString() != "entry-enabled-2" {
		t.Errorf("Second enabled entry should be 'entry-enabled-2', got %q", (*result.EnabledEntries)[1].ValueString())
	}
}

// TestTransform_NoEntries verifies that when a v4 profile has no entry blocks at all,
// EnabledEntries is nil (the resource was never configured with entries).
func TestTransform_NoEntries(t *testing.T) {
	ctx := context.Background()

	source := v500.SourceCloudflareDLPProfileModel{
		ID:                types.StringValue("profile-uuid"),
		AccountID:         types.StringValue("account-uuid"),
		Name:              types.StringValue("Empty Profile"),
		Type:              types.StringValue("predefined"),
		Entry:             nil,
		AllowedMatchCount: types.Int64Value(0),
		OCREnabled:        types.BoolValue(false),
	}

	result, diags := v500.Transform(ctx, source)
	if diags.HasError() {
		t.Fatalf("Transform returned unexpected errors: %v", diags)
	}

	if result.EnabledEntries != nil {
		t.Fatalf("EnabledEntries should be nil when no entry blocks exist, got %v", *result.EnabledEntries)
	}
}

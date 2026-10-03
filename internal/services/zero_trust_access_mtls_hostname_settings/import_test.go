package zero_trust_access_mtls_hostname_settings

import "testing"

// Cloudflare account IDs and zone IDs are both 32-character hex strings, so the
// scope of an import ID cannot be inferred from the ID itself. This resource
// previously branched on len(importID) == 32 to choose between account and zone
// scope, which made the zone branch unreachable for every real ID and silently
// wrote a zone ID into account_id. The import ID therefore carries an explicit
// accounts/zones discriminator.
const (
	testAccountID = "023e105f4ecef8ad9ca31a8372d0c353"
	testZoneID    = "b9c9ae341a26b5b692441bb6573015b9"
)

// Guards the premise behind the discriminator: if these ever differ in length,
// the ambiguity that makes an explicit scope necessary no longer holds.
func TestImportScopeIsNotInferableFromLength(t *testing.T) {
	if len(testAccountID) != len(testZoneID) {
		t.Fatalf("fixture IDs differ in length (%d vs %d)", len(testAccountID), len(testZoneID))
	}
}

func TestResolveImportScope(t *testing.T) {
	for _, tc := range []struct {
		name          string
		importID      string
		wantErr       bool
		wantAccountID string
		wantZoneID    string
	}{
		{
			name:          "account scope",
			importID:      "accounts/" + testAccountID,
			wantAccountID: testAccountID,
		},
		{
			// Unreachable before the discriminator was introduced: a 32-character
			// zone ID took the account branch.
			name:       "zone scope",
			importID:   "zones/" + testZoneID,
			wantZoneID: testZoneID,
		},
		{
			// Previously accepted and silently resolved to account scope.
			name:     "bare account id is rejected",
			importID: testAccountID,
			wantErr:  true,
		},
		{
			// Previously accepted and silently written to account_id.
			name:     "bare zone id is rejected",
			importID: testZoneID,
			wantErr:  true,
		},
		{
			name:     "too many segments",
			importID: "accounts/" + testAccountID + "/extra",
			wantErr:  true,
		},
		{
			name:     "unknown discriminator",
			importID: "organizations/" + testAccountID,
			wantErr:  true,
		},
		{
			name:     "empty id",
			importID: "",
			wantErr:  true,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			accountID, zoneID, diags := resolveImportScope(tc.importID)

			if tc.wantErr {
				if !diags.HasError() {
					t.Fatalf("expected an error for %q, got account_id=%v zone_id=%v",
						tc.importID, accountID, zoneID)
				}
				return
			}
			if diags.HasError() {
				t.Fatalf("unexpected error for %q: %v", tc.importID, diags)
			}

			// Exactly one scope must be populated; the other stays null so the
			// resource never sends both to the API.
			if tc.wantAccountID != "" {
				if accountID.ValueString() != tc.wantAccountID {
					t.Errorf("account_id: got %q, want %q", accountID.ValueString(), tc.wantAccountID)
				}
				if !zoneID.IsNull() {
					t.Errorf("zone_id: got %q, want null", zoneID.ValueString())
				}
			}
			if tc.wantZoneID != "" {
				if zoneID.ValueString() != tc.wantZoneID {
					t.Errorf("zone_id: got %q, want %q", zoneID.ValueString(), tc.wantZoneID)
				}
				if !accountID.IsNull() {
					t.Errorf("account_id: got %q, want null", accountID.ValueString())
				}
			}
		})
	}
}

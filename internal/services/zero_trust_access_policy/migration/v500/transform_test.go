package v500

import (
	"context"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

// TestTransform_ZoneIDOnly_AccountIDIsNull reproduces APIX-851:
// When a v4 cloudflare_access_policy used zone_id (no account_id), the
// Transform function used to produce a v5 state with null account_id. This
// caused "missing required account_id parameter" on any subsequent CRUD
// operation because the Go SDK requires a non-empty account_id to construct
// the API URL: POST accounts/{account_id}/access/policies.
//
// Customer scenario:
//  1. v4 config: zone_id + application_id (zone-scoped app policy)
//  2. Migration to v5: cloudflare_zero_trust_access_policy (reusable policy)
//  3. v5 state has null account_id → Read() fails even when config has account_id hardcoded
//     because Read() sources account_id from state, not config.
//
// Fix (APIX-851): Transform must either:
//
//	(a) emit a diagnostic error explaining zone_id-only policies require manual account_id, or
//	(b) resolve account_id from another source (CLOUDFLARE_ACCOUNT_ID env var).
//
// This test exercises path (a) — no env var set, expect a clear diagnostic error.
// TestTransform_ZoneIDOnly_AccountIDFromEnvVar covers path (b).
func TestTransform_ZoneIDOnly_AccountIDIsNull(t *testing.T) {
	// Defensively clear the env var in case the test runner has it set;
	// path (a) requires no fallback to be available.
	t.Setenv("CLOUDFLARE_ACCOUNT_ID", "")
	ctx := context.Background()

	v4State := SourceAccessPolicyModel{
		ID:              types.StringValue("test-policy-id"),
		AccountID:       types.StringNull(),
		ZoneID:          types.StringValue("zone-abc123"),
		ApplicationID:   types.StringValue("app-def456"),
		Name:            types.StringValue("test-policy"),
		Decision:        types.StringValue("allow"),
		SessionDuration: types.StringValue("24h"),
		Precedence:      types.Int64Value(1),
	}

	v5State, diags := Transform(ctx, v4State)

	// Either path is acceptable per APIX-851 contract:
	//   (a) clear diagnostic error  → v5State==nil, diags.HasError()
	//   (b) populated account_id    → v5State!=nil, !diags.HasError()
	if diags.HasError() {
		// Path (a): make sure the error names APIX-851 so the user knows where to look.
		msg := diags.Errors()[0].Summary() + " " + diags.Errors()[0].Detail()
		if !strings.Contains(msg, "APIX-851") {
			t.Errorf("expected diagnostic error to reference APIX-851, got: %s", msg)
		}
		return
	}

	if v5State == nil {
		t.Fatalf("Transform returned nil state and no error diagnostic")
	}
	if v5State.AccountID.IsNull() || v5State.AccountID.ValueString() == "" {
		t.Errorf("APIX-851: expected account_id to be populated (or a diagnostic error) after migration "+
			"from zone_id-scoped policy, but got null/empty. zone_id was %q, application_id was %q",
			v4State.ZoneID.ValueString(), v4State.ApplicationID.ValueString())
	}
}

// TestTransform_ZoneIDOnly_AccountIDFromEnvVar verifies the CLOUDFLARE_ACCOUNT_ID
// fallback path of the APIX-851 fix: when the v4 state is zone-scoped and no
// account_id is in state, the migration uses CLOUDFLARE_ACCOUNT_ID and emits a
// Warning diagnostic (not an Error).
func TestTransform_ZoneIDOnly_AccountIDFromEnvVar(t *testing.T) {
	t.Setenv("CLOUDFLARE_ACCOUNT_ID", "env-acct-789")
	ctx := context.Background()

	v4State := SourceAccessPolicyModel{
		ID:              types.StringValue("test-policy-id"),
		AccountID:       types.StringNull(),
		ZoneID:          types.StringValue("zone-abc123"),
		ApplicationID:   types.StringValue("app-def456"),
		Name:            types.StringValue("test-policy"),
		Decision:        types.StringValue("allow"),
		SessionDuration: types.StringValue("24h"),
		Precedence:      types.Int64Value(1),
	}

	v5State, diags := Transform(ctx, v4State)
	if diags.HasError() {
		t.Fatalf("did not expect an error with CLOUDFLARE_ACCOUNT_ID set, got: %v", diags)
	}
	if v5State == nil {
		t.Fatalf("Transform returned nil state without an error diagnostic")
	}
	if got := v5State.AccountID.ValueString(); got != "env-acct-789" {
		t.Errorf("expected account_id to be derived from CLOUDFLARE_ACCOUNT_ID env var (env-acct-789), got %q", got)
	}
	if diags.WarningsCount() == 0 {
		t.Errorf("expected a Warning diagnostic naming the env-var fallback, got none")
	}
}

// TestTransform_AccountIDPreservedWhenSet verifies that when a v4 state has
// account_id set, it is correctly preserved in the v5 state.
// This is a control test for APIX-851.
func TestTransform_AccountIDPreservedWhenSet(t *testing.T) {
	ctx := context.Background()

	v4State := SourceAccessPolicyModel{
		ID:              types.StringValue("test-policy-id"),
		AccountID:       types.StringValue("acct-123456"),
		ZoneID:          types.StringNull(),
		ApplicationID:   types.StringNull(),
		Name:            types.StringValue("test-policy"),
		Decision:        types.StringValue("allow"),
		SessionDuration: types.StringValue("24h"),
		Precedence:      types.Int64Value(1),
	}

	v5State, diags := Transform(ctx, v4State)
	if diags.HasError() {
		t.Fatalf("unexpected diagnostics: %v", diags)
	}

	if v5State.AccountID.ValueString() != "acct-123456" {
		t.Errorf("expected account_id to be preserved as 'acct-123456', got %q (null=%v)",
			v5State.AccountID.ValueString(), v5State.AccountID.IsNull())
	}
}

// TestTransform_BothZoneIDAndAccountID_AccountIDPreserved verifies that when
// both zone_id and account_id are set in v4 state (some users had both), the
// account_id is correctly preserved in v5.
func TestTransform_BothZoneIDAndAccountID_AccountIDPreserved(t *testing.T) {
	ctx := context.Background()

	v4State := SourceAccessPolicyModel{
		ID:              types.StringValue("test-policy-id"),
		AccountID:       types.StringValue("acct-123456"),
		ZoneID:          types.StringValue("zone-abc123"),
		ApplicationID:   types.StringValue("app-def456"),
		Name:            types.StringValue("test-policy"),
		Decision:        types.StringValue("allow"),
		SessionDuration: types.StringValue("24h"),
		Precedence:      types.Int64Value(1),
	}

	v5State, diags := Transform(ctx, v4State)
	if diags.HasError() {
		t.Fatalf("unexpected diagnostics: %v", diags)
	}

	if v5State.AccountID.ValueString() != "acct-123456" {
		t.Errorf("expected account_id to be preserved as 'acct-123456', got %q (null=%v)",
			v5State.AccountID.ValueString(), v5State.AccountID.IsNull())
	}
}

func TestTransformConditions_SkipsEmptyAuthMethodAndCommonName(t *testing.T) {
	ctx := context.Background()

	conds, diags := transformConditions(ctx, []SourceConditionGroupModel{
		{
			AuthMethod: types.StringValue(""),
			CommonName: types.StringValue("   "),
		},
	})

	if diags.HasError() {
		t.Fatalf("unexpected diagnostics: %v", diags)
	}

	if len(conds) != 0 {
		t.Fatalf("expected 0 conditions, got %d: %#v", len(conds), conds)
	}
}

func TestTransformConditions_IncludesNonEmptyAuthMethodAndCommonName(t *testing.T) {
	ctx := context.Background()

	conds, diags := transformConditions(ctx, []SourceConditionGroupModel{
		{
			AuthMethod: types.StringValue(" otp "),
			CommonName: types.StringValue(" device1.example.com "),
		},
	})

	if diags.HasError() {
		t.Fatalf("unexpected diagnostics: %v", diags)
	}

	if len(conds) != 2 {
		t.Fatalf("expected 2 conditions, got %d: %#v", len(conds), conds)
	}

	var gotAuthMethod, gotCommonName bool
	for _, c := range conds {
		if c.AuthMethod != nil {
			gotAuthMethod = true
			if c.AuthMethod.AuthMethod.ValueString() != "otp" {
				t.Fatalf("expected trimmed auth_method 'otp', got %q", c.AuthMethod.AuthMethod.ValueString())
			}
		}

		if c.CommonName != nil {
			gotCommonName = true
			if c.CommonName.CommonName.ValueString() != "device1.example.com" {
				t.Fatalf("expected trimmed common_name 'device1.example.com', got %q", c.CommonName.CommonName.ValueString())
			}
		}
	}

	if !gotAuthMethod {
		t.Fatal("expected auth_method condition, got none")
	}

	if !gotCommonName {
		t.Fatal("expected common_name condition, got none")
	}
}

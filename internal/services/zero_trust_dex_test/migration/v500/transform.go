package v500

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Transform converts source (legacy v4) state to target (current v5) state.
// This function is shared by both UpgradeFromV4 and MoveState handlers.
func Transform(ctx context.Context, source SourceCloudflareDeviceDexTestModel) (*ZeroTrustDEXTestModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	// Step 1: Validate required fields
	if source.AccountID.IsNull() || source.AccountID.IsUnknown() {
		diags.AddError(
			"Missing required field",
			"account_id is required for zero_trust_dex_test migration. The source state is missing this field, which indicates corrupted state.",
		)
		return nil, diags
	}
	if source.Name.IsNull() || source.Name.IsUnknown() {
		diags.AddError(
			"Missing required field",
			"name is required for zero_trust_dex_test migration. The source state is missing this field, which indicates corrupted state.",
		)
		return nil, diags
	}

	// Step 2: Initialize target with direct copies (pass-through fields)
	target := &ZeroTrustDEXTestModel{
		ID:          source.ID,
		AccountID:   source.AccountID,
		Name:        source.Name,
		Description: source.Description,
		Interval:    source.Interval,
		Enabled:     source.Enabled,
	}

	// Step 3: Handle new fields in v5
	// test_id: Copy from id field (both should have same value)
	target.TestID = source.ID

	// targeted: Set to Null (API will populate)
	target.Targeted = types.BoolNull()

	// target_policies: Set to Null (API will populate)
	target.TargetPolicies = customfield.NullObjectList[ZeroTrustDEXTestTargetPoliciesModel](ctx)

	// Step 4: Handle structure transformation - data field
	// v4: []SourceDEXTestDataModel (array with MaxItems:1)
	// v5: *ZeroTrustDEXTestDataModel (pointer)
	if len(source.Data) > 0 {
		sourceData := source.Data[0]

		// Transform array element to pointer
		target.Data = &ZeroTrustDEXTestDataModel{
			Kind:   sourceData.Kind,
			Host:   sourceData.Host,
			Method: sourceData.Method,
		}

		// Note: method field is optional and only present for kind="http"
		// We copy it regardless - v5 schema marks it as optional and will handle validation
	} else {
		// Edge case: data array is empty (shouldn't happen in valid state)
		diags.AddError(
			"Invalid state",
			"data field is required but empty in source state. This indicates corrupted state.",
		)
		return nil, diags
	}

	// Step 5: Drop deprecated fields (don't copy to v5)
	// source.Updated is intentionally not copied (removed in v5)
	// source.Created is intentionally not copied (removed in v5)

	return target, diags
}

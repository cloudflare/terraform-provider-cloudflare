package v500

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Transform converts source (legacy v4) state to target (current v5) state.
// This function is shared by both UpgradeFromV4 and MoveState handlers.
func Transform(ctx context.Context, source *SourceAccessServiceTokenModel) (*TargetAccessServiceTokenModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	// Step 1: Validate required fields
	if source.Name.IsNull() || source.Name.IsUnknown() {
		diags.AddError(
			"Missing required field",
			"name is required for access_service_token migration. The source state is missing this field, which indicates corrupted state.",
		)
		return nil, diags
	}

	// Step 2: Initialize target with direct copies
	target := &TargetAccessServiceTokenModel{
		ID:           source.ID,
		AccountID:    source.AccountID,
		ZoneID:       source.ZoneID,
		Name:         source.Name,
		ClientID:     source.ClientID,
		ClientSecret: source.ClientSecret,
		Duration:     source.Duration,
	}

	// Step 3: Handle type conversion: client_secret_version (Int64 → Float64)
	if !source.ClientSecretVersion.IsNull() && !source.ClientSecretVersion.IsUnknown() {
		target.ClientSecretVersion = types.Float64Value(float64(source.ClientSecretVersion.ValueInt64()))
	} else {
		// Apply v5 default value
		target.ClientSecretVersion = types.Float64Value(1.0)
	}

	// Step 4: Handle timestamp conversions (String → RFC3339)
	if !source.ExpiresAt.IsNull() && !source.ExpiresAt.IsUnknown() {
		val, valDiags := timetypes.NewRFC3339Value(source.ExpiresAt.ValueString())
		diags.Append(valDiags...)
		target.ExpiresAt = val
	} else {
		target.ExpiresAt = timetypes.NewRFC3339Null()
	}

	if !source.PreviousClientSecretExpiresAt.IsNull() && !source.PreviousClientSecretExpiresAt.IsUnknown() {
		val, valDiags := timetypes.NewRFC3339Value(source.PreviousClientSecretExpiresAt.ValueString())
		diags.Append(valDiags...)
		target.PreviousClientSecretExpiresAt = val
	} else {
		target.PreviousClientSecretExpiresAt = timetypes.NewRFC3339Null()
	}

	// Step 5: Drop deprecated fields (don't migrate them)
	// source.MinDaysForRenewal is intentionally not copied to target

	return target, diags
}

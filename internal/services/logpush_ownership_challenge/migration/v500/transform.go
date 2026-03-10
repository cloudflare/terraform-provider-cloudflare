package v500

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Transform converts source (v4 SDKv2) state to target (v5 Plugin Framework) state.
//
// This is one of the simplest transformations:
//   - account_id, zone_id, destination_conf: direct copies (no changes)
//   - ownership_challenge_filename: renamed to filename in v5
//   - message, valid: set to Null (new computed fields, API will populate)
//
// This function is used by UpgradeFromV4 handler.
func Transform(ctx context.Context, source SourceCloudflareLogpushOwnershipChallengeModel) (*TargetLogpushOwnershipChallengeModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	// Validate required fields
	if source.DestinationConf.IsNull() || source.DestinationConf.IsUnknown() {
		diags.AddError(
			"Missing required field",
			"destination_conf is required for logpush_ownership_challenge migration. The source state is missing this field, which indicates corrupted state.",
		)
		return nil, diags
	}

	// Initialize target model
	target := &TargetLogpushOwnershipChallengeModel{}

	// Step 1: Direct field copies (all v4 fields that survive to v5)
	target.AccountID = source.AccountID
	target.ZoneID = source.ZoneID
	target.DestinationConf = source.DestinationConf

	// Step 2: Rename ownership_challenge_filename → filename
	target.Filename = source.OwnershipChallengeFilename

	// Step 3: Set new v5 computed fields to Null (API will populate on next read/create)
	target.Message = types.StringNull()
	target.Valid = types.BoolNull()

	return target, diags
}

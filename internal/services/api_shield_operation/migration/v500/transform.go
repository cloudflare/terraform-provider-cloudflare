package v500

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
)

// Transform converts source (v4 SDKv2) state to target (v5 Plugin Framework) state.
//
// This is one of the simplest transformations:
// - All v4 fields are direct copies (no renames, no type changes)
// - New v5 fields are all Computed and set to Null (API will populate on read)
//
// This function is used by UpgradeFromV4 handler.
func Transform(ctx context.Context, source SourceCloudflareAPIShieldOperationModel) (*TargetAPIShieldOperationModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	// Validate required fields
	if source.ZoneID.IsNull() || source.ZoneID.IsUnknown() {
		diags.AddError(
			"Missing required field",
			"zone_id is required for api_shield_operation migration. The source state is missing this field, which indicates corrupted state.",
		)
		return nil, diags
	}

	if source.Method.IsNull() || source.Method.IsUnknown() {
		diags.AddError(
			"Missing required field",
			"method is required for api_shield_operation migration. The source state is missing this field, which indicates corrupted state.",
		)
		return nil, diags
	}

	if source.Host.IsNull() || source.Host.IsUnknown() {
		diags.AddError(
			"Missing required field",
			"host is required for api_shield_operation migration. The source state is missing this field, which indicates corrupted state.",
		)
		return nil, diags
	}

	if source.Endpoint.IsNull() || source.Endpoint.IsUnknown() {
		diags.AddError(
			"Missing required field",
			"endpoint is required for api_shield_operation migration. The source state is missing this field, which indicates corrupted state.",
		)
		return nil, diags
	}

	// Initialize target model
	target := &TargetAPIShieldOperationModel{}

	// Step 1: Direct field copies (all v4 fields)
	// These are all unchanged between v4 and v5
	target.ID = source.ID
	target.ZoneID = source.ZoneID
	target.Method = source.Method
	target.Host = source.Host
	target.Endpoint = source.Endpoint

	// Step 2: New v5 Computed fields
	// CRITICAL: OperationID must be set to ID value (not Null) because the Read function
	// uses OperationID to make API calls. In v5, ID and OperationID are the same value.
	target.OperationID = source.ID
	target.LastUpdated = timetypes.NewRFC3339Null()
	target.Features = customfield.NullObject[TargetAPIShieldOperationFeaturesModel](ctx)

	return target, diags
}

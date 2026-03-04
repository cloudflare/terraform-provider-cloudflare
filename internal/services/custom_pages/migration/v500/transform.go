package v500

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Transform converts source (legacy v4) state to target (current v5) state.
//
// Transformations performed:
// 1. type → identifier (field rename)
// 2. state: add "default" if missing (Optional → Required)
// 3. Pass through: url, zone_id, account_id
// 4. Set new computed fields to Null (API will populate)
func Transform(ctx context.Context, source SourceCloudflareCustomPagesModel) (*TargetCustomPagesModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	// Initialize target model
	target := &TargetCustomPagesModel{}

	// Transformation 1: type → identifier (field rename)
	// Both are String, both Required, same enum values
	target.Identifier = source.Type

	// Transformation 2: state field - add default if missing
	// v4: Optional (might be null)
	// v5: Required (must have value)
	if source.State.IsNull() || source.State.IsUnknown() {
		// v4 didn't have state, set default per tf-migrate logic
		target.State = types.StringValue("default")
	} else {
		// v4 had state, keep it
		target.State = source.State
	}

	// Transformation 3: Pass through unchanged fields
	// id: Computed in both v4 and v5
	target.ID = source.ID

	// url: Required (v4) → Optional+Computed (v5)
	// v4 state always has url (was required), so just pass through
	target.URL = source.URL

	// zone_id and account_id: Optional in both v4 and v5
	target.ZoneID = source.ZoneID
	target.AccountID = source.AccountID

	// Transformation 4: New computed fields in v5
	// These didn't exist in v4 state, set to Null
	// API will populate them on first refresh
	target.CreatedOn = timetypes.NewRFC3339Null()
	target.ModifiedOn = timetypes.NewRFC3339Null()
	target.Description = types.StringNull()
	target.PreviewTarget = types.StringNull()
	target.RequiredTokens = customfield.NullList[types.String](ctx)

	return target, diags
}

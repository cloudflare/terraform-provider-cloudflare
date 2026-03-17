package v500

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Transform converts source (v4) state to target (v5) state.
//
// Key transformations:
// 1. enforce_twofactor: top-level bool → settings.enforce_twofactor (wrapped in nested settings object)
// 2. settings.abuse_contact_email: NEW field, set to null (will be refreshed from API)
// 3. unit: NEW nested object, set to null (will be refreshed from API)
// 4. managed_by: NEW nested object, set to null (will be refreshed from API)
// 5. created_on: NEW computed timestamp, set to null (will be refreshed from API)
// 6. id, name, type: direct pass-through (types match)
func Transform(ctx context.Context, source SourceCloudflareAccountModel) (*TargetAccountModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	// Step 1: Validate required fields
	if source.Name.IsNull() || source.Name.IsUnknown() {
		diags.AddError(
			"Missing required field",
			"name is required for account migration. The source state is missing this field, which indicates corrupted state.",
		)
		return nil, diags
	}

	// Step 2: Build the settings nested object from enforce_twofactor
	// In v4, enforce_twofactor was a top-level attribute.
	// In v5, it's nested inside settings along with abuse_contact_email.
	settings := TargetAccountSettingsModel{
		EnforceTwofactor:  source.EnforceTwofactor,
		AbuseContactEmail: types.StringNull(), // NEW in v5, will be refreshed from API
	}
	settingsObj, settingsDiags := customfield.NewObject(ctx, &settings)
	diags.Append(settingsDiags...)
	if diags.HasError() {
		return nil, diags
	}

	// Step 3: Set new computed nested objects to null (will be refreshed from API)
	unitNull := customfield.NullObject[TargetAccountUnitModel](ctx)
	managedByNull := customfield.NullObject[TargetAccountManagedByModel](ctx)

	// Step 4: Build target model
	target := &TargetAccountModel{
		ID:        source.ID,
		Name:      source.Name,
		Type:      source.Type,
		Settings:  settingsObj,
		Unit:      unitNull,
		ManagedBy: managedByNull,
		CreatedOn: timetypes.NewRFC3339Null(), // NEW in v5, will be refreshed from API
	}

	return target, diags
}

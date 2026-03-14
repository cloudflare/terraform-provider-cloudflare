package v500

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Transform converts a source cloudflare_byo_ip_prefix state (v4 SDKv2) to the target
// cloudflare_byo_ip_prefix state (v5 Plugin Framework).
//
// Field mapping:
//   - id:          copied from source id (which was set from prefix_id in v4)
//   - account_id:  copied unchanged
//   - description: copied unchanged
//   - prefix_id:   dropped (was mapped to id in v4; v5 uses id directly)
//   - advertisement: dropped (replaced by computed-only 'advertised' bool in v5)
//
// New v5 required fields (asn, cidr) and all computed fields are set to null.
// The v5 provider will populate them from the API on the next plan/apply.
func Transform(_ context.Context, source SourceCloudflareByoIPPrefixModel) (*TargetByoIPPrefixModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	target := &TargetByoIPPrefixModel{
		// id: In v4, d.SetId(prefix_id) sets the resource ID to the prefix_id value.
		// The v4 state stores this as "id". Copy it directly to v5 id.
		ID: source.ID,

		// Pass-through fields
		AccountID:   source.AccountID,
		Description: source.Description,

		// New v5 required fields: user must add these manually after migration.
		// The v5 provider will fetch them from the API during the first plan/refresh.
		ASN:  types.Int64Null(),
		CIDR: types.StringNull(),

		// New v5 optional fields: not present in v4, leave as null.
		LOADocumentID:       types.StringNull(),
		DelegateLOACreation: types.BoolNull(),

		// Computed fields: v5 provider will populate from API.
		Advertised:               types.BoolNull(),
		AdvertisedModifiedAt:     timetypes.NewRFC3339Null(),
		Approved:                 types.StringNull(),
		CreatedAt:                timetypes.NewRFC3339Null(),
		IrrValidationState:       types.StringNull(),
		ModifiedAt:               timetypes.NewRFC3339Null(),
		OnDemandEnabled:          types.BoolNull(),
		OnDemandLocked:           types.BoolNull(),
		OwnershipValidationState: types.StringNull(),
		OwnershipValidationToken: types.StringNull(),
		RPKIValidationState:      types.StringNull(),
	}

	return target, diags
}

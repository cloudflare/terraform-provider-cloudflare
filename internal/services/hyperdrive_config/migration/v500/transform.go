package v500

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Transform converts a source cloudflare_hyperdrive_config state (v4) to target cloudflare_hyperdrive_config state (v500).
func Transform(ctx context.Context, source SourceCloudflareHyperdriveConfigModel) (*TargetHyperdriveConfigModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	// Validate required fields
	if source.AccountID.IsNull() || source.AccountID.IsUnknown() {
		diags.AddError("Missing required field", "account_id is required for hyperdrive_config migration")
		return nil, diags
	}
	if source.Name.IsNull() || source.Name.IsUnknown() {
		diags.AddError("Missing required field", "name is required for hyperdrive_config migration")
		return nil, diags
	}

	target := &TargetHyperdriveConfigModel{
		// Direct copies
		ID:        source.ID,
		AccountID: source.AccountID,
		Name:      source.Name,

		// New optional fields: initialize as null
		OriginConnectionLimit: types.Int64Null(),
		MTLS:                  nil,

		// New computed fields: initialize as null, API will populate on next plan/apply
		CreatedOn:  timetypes.NewRFC3339Null(),
		ModifiedOn: timetypes.NewRFC3339Null(),
	}

	// Origin: transform nested object
	if source.Origin != nil {
		target.Origin = transformOrigin(source.Origin)
	}

	// Caching: v4 uses types.Object, v5 uses *TargetHyperdriveConfigCachingModel
	target.Caching = transformCaching(source.Caching)

	return target, diags
}

// transformOrigin converts the v4 origin to v5 origin, adding the new service_id field as null.
func transformOrigin(source *SourceCloudflareHyperdriveConfigOriginModel) *TargetHyperdriveConfigOriginModel {
	return &TargetHyperdriveConfigOriginModel{
		// Direct copies - all existing fields pass through unchanged
		Database:           source.Database,
		Host:               source.Host,
		Password:           source.Password,
		Port:               source.Port,
		Scheme:             source.Scheme,
		User:               source.User,
		AccessClientID:     source.AccessClientID,
		AccessClientSecret: source.AccessClientSecret,

		// New in v5: initialize as null
		ServiceID: types.StringNull(),
	}
}

// transformCaching converts the v4 caching types.Object to v5 *TargetHyperdriveConfigCachingModel.
// In v4, caching was Optional+Computed and stored as types.Object.
// In v5, caching is Optional and stored as a pointer to a struct.
func transformCaching(source types.Object) *TargetHyperdriveConfigCachingModel {
	if source.IsNull() || source.IsUnknown() {
		return nil
	}

	// Extract attributes from the types.Object
	attrs := source.Attributes()
	if len(attrs) == 0 {
		return nil
	}

	target := &TargetHyperdriveConfigCachingModel{}

	// disabled
	if v, ok := attrs["disabled"]; ok {
		target.Disabled = attrValueToBool(v)
	} else {
		target.Disabled = types.BoolNull()
	}

	// max_age
	if v, ok := attrs["max_age"]; ok {
		target.MaxAge = attrValueToInt64(v)
	} else {
		target.MaxAge = types.Int64Null()
	}

	// stale_while_revalidate
	if v, ok := attrs["stale_while_revalidate"]; ok {
		target.StaleWhileRevalidate = attrValueToInt64(v)
	} else {
		target.StaleWhileRevalidate = types.Int64Null()
	}

	return target
}

// attrValueToBool converts an attr.Value to types.Bool.
func attrValueToBool(v attr.Value) types.Bool {
	if v == nil || v.IsNull() {
		return types.BoolNull()
	}
	if v.IsUnknown() {
		return types.BoolUnknown()
	}
	if bv, ok := v.(types.Bool); ok {
		return bv
	}
	return types.BoolNull()
}

// attrValueToInt64 converts an attr.Value to types.Int64.
func attrValueToInt64(v attr.Value) types.Int64 {
	if v == nil || v.IsNull() {
		return types.Int64Null()
	}
	if v.IsUnknown() {
		return types.Int64Unknown()
	}
	if iv, ok := v.(types.Int64); ok {
		return iv
	}
	return types.Int64Null()
}

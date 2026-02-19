package v500

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Transform converts a v4 custom_hostname_fallback_origin state to v5 format.
//
// Field transformations:
// - created_at: types.String → timetypes.RFC3339 (v4 stored as plain string, v5 uses RFC3339 type)
// - updated_at: types.String → timetypes.RFC3339 (v4 stored as plain string, v5 uses RFC3339 type)
// - errors: types.List → customfield.List[types.String] (v4 used plain list, v5 uses typed list)
// - All other fields: Direct copy (same types in v4 and v5)
func Transform(ctx context.Context, source SourceCustomHostnameFallbackOriginModel) (*TargetCustomHostnameFallbackOriginModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	// Convert timestamp strings to RFC3339 type
	// v4 stored these as plain strings, v5 uses timetypes.RFC3339
	createdAt := timetypes.NewRFC3339Null()
	if !source.CreatedAt.IsNull() && !source.CreatedAt.IsUnknown() {
		var d diag.Diagnostics
		createdAt, d = timetypes.NewRFC3339Value(source.CreatedAt.ValueString())
		diags.Append(d...)
	}

	updatedAt := timetypes.NewRFC3339Null()
	if !source.UpdatedAt.IsNull() && !source.UpdatedAt.IsUnknown() {
		var d diag.Diagnostics
		updatedAt, d = timetypes.NewRFC3339Value(source.UpdatedAt.ValueString())
		diags.Append(d...)
	}

	// Convert errors list to customfield.List
	// v4 used types.List, v5 uses customfield.List[types.String]
	errors := customfield.NullList[types.String](ctx)
	if !source.Errors.IsNull() && !source.Errors.IsUnknown() {
		var d diag.Diagnostics
		errors, d = customfield.NewList[types.String](ctx, source.Errors.Elements())
		diags.Append(d...)
	}

	target := &TargetCustomHostnameFallbackOriginModel{
		ID:        source.ID,
		ZoneID:    source.ZoneID,
		Origin:    source.Origin,
		CreatedAt: createdAt,
		Status:    source.Status,
		UpdatedAt: updatedAt,
		Errors:    errors,
	}

	return target, diags
}

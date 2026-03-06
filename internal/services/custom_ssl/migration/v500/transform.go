package v500

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Transform converts a v4 source state (SourceCloudflareCustomSSLModel) to the v5 target
// state (TargetCustomSSLModel).
//
// Transformation summary:
//   - custom_ssl_options[0].{certificate,private_key,bundle_method,type} → hoisted to top level
//   - custom_ssl_options[0].geo_restrictions (plain string) → geo_restrictions.label (nested object)
//   - custom_ssl_priority → dropped (write-only in v4, absent in v5)
//   - priority (types.Int64) → types.Float64
//   - uploaded_on / modified_on / expires_on (types.String) → timetypes.RFC3339
//   - hosts (types.List) → customfield.List[types.String]
//   - New v5-only computed fields → null (will be populated from API on next refresh)
func Transform(ctx context.Context, source SourceCloudflareCustomSSLModel) (*TargetCustomSSLModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	// Validate required field.
	if source.ZoneID.IsNull() || source.ZoneID.IsUnknown() {
		diags.AddError(
			"Missing required field",
			"zone_id is required for custom_ssl migration. The source state is missing this field, which indicates corrupted state.",
		)
		return nil, diags
	}

	target := &TargetCustomSSLModel{
		ID:     source.ID,
		ZoneID: source.ZoneID,
		// computed — copied directly
		Issuer:    source.Issuer,
		Signature: source.Signature,
		Status:    source.Status,
	}

	// -----------------------------------------------------------------------
	// Hoist fields from custom_ssl_options[0] block to top-level.
	// -----------------------------------------------------------------------
	if len(source.CustomSSLOptions) > 0 {
		opts := source.CustomSSLOptions[0]

		target.Certificate = opts.Certificate
		target.PrivateKey = opts.PrivateKey

		// bundle_method: optional in v4, computed+optional with default "ubiquitous" in v5.
		if !opts.BundleMethod.IsNull() && !opts.BundleMethod.IsUnknown() {
			target.BundleMethod = opts.BundleMethod
		} else {
			target.BundleMethod = types.StringValue("ubiquitous")
		}

		// type: optional in v4, computed+optional with default "legacy_custom" in v5.
		if !opts.Type.IsNull() && !opts.Type.IsUnknown() {
			target.Type = opts.Type
		} else {
			target.Type = types.StringValue("legacy_custom")
		}

		// geo_restrictions: TypeString in v4 → SingleNestedAttribute{label} in v5.
		// Empty string means no restriction was set; leave as nil.
		if !opts.GeoRestrictions.IsNull() && !opts.GeoRestrictions.IsUnknown() &&
			opts.GeoRestrictions.ValueString() != "" {
			target.GeoRestrictions = &TargetCustomSSLGeoRestrictionsModel{
				Label: opts.GeoRestrictions,
			}
		}
		// else: target.GeoRestrictions stays nil
	}

	// -----------------------------------------------------------------------
	// priority: Int64 → Float64
	// -----------------------------------------------------------------------
	if !source.Priority.IsNull() && !source.Priority.IsUnknown() {
		target.Priority = types.Float64Value(float64(source.Priority.ValueInt64()))
	} else {
		target.Priority = types.Float64Value(0)
	}

	// -----------------------------------------------------------------------
	// Timestamp fields: TypeString → timetypes.RFC3339
	// -----------------------------------------------------------------------
	target.UploadedOn = convertStringToRFC3339(source.UploadedOn)
	target.ModifiedOn = convertStringToRFC3339(source.ModifiedOn)
	target.ExpiresOn = convertStringToRFC3339(source.ExpiresOn)

	// -----------------------------------------------------------------------
	// hosts: types.List → customfield.List[types.String]
	// -----------------------------------------------------------------------
	if !source.Hosts.IsNull() && !source.Hosts.IsUnknown() {
		var hostStrings []string
		diags.Append(source.Hosts.ElementsAs(ctx, &hostStrings, false)...)
		if diags.HasError() {
			return nil, diags
		}
		hostAttrs := make([]types.String, len(hostStrings))
		for i, h := range hostStrings {
			hostAttrs[i] = types.StringValue(h)
		}
		hostList, listDiags := customfield.NewList[types.String](ctx, hostAttrs)
		diags.Append(listDiags...)
		if diags.HasError() {
			return nil, diags
		}
		target.Hosts = hostList
	} else {
		target.Hosts = customfield.NullList[types.String](ctx)
	}

	// -----------------------------------------------------------------------
	// New v5-only computed fields — set to null; API will populate on refresh.
	// -----------------------------------------------------------------------
	target.PolicyRestrictions = types.StringNull()
	target.Policy = types.StringNull()
	target.Deploy = types.StringNull()
	target.KeylessServer = customfield.NullObject[TargetCustomSSLKeylessServerModel](ctx)

	// custom_ssl_priority is intentionally NOT migrated — it was a write-only
	// reprioritization list in v4 that does not exist in v5.

	return target, diags
}

// convertStringToRFC3339 converts a types.String value containing an RFC3339 timestamp
// to a timetypes.RFC3339 value. Null/unknown strings produce a null RFC3339 value.
func convertStringToRFC3339(s types.String) timetypes.RFC3339 {
	if s.IsNull() || s.IsUnknown() {
		return timetypes.NewRFC3339Null()
	}
	v := s.ValueString()
	if v == "" {
		return timetypes.NewRFC3339Null()
	}
	return timetypes.NewRFC3339ValueMust(v)
}

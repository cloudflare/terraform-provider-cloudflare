// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package v500

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Transform converts a v4 (SDKv2) state to a v5 (Plugin Framework) state
// This handles all field transformations needed for the migration
func Transform(ctx context.Context, source SourceCloudflareZoneDNSSECModel) (*TargetZoneDNSSECModel, diag.Diagnostics) {
	var diags diag.Diagnostics
	target := &TargetZoneDNSSECModel{}

	tflog.Debug(ctx, "Starting zone_dnssec state transformation from v4 to v5")

	// Direct pass-through fields (same type, no transformation)
	target.ID = source.ID
	target.ZoneID = source.ZoneID
	target.Algorithm = source.Algorithm
	target.Digest = source.Digest
	target.DigestAlgorithm = source.DigestAlgorithm
	target.DigestType = source.DigestType
	target.DS = source.DS
	target.KeyType = source.KeyType
	target.PublicKey = source.PublicKey

	// Type conversion: Int64 → Float64 for flags
	if !source.Flags.IsNull() && !source.Flags.IsUnknown() {
		flagsValue := float64(source.Flags.ValueInt64())
		target.Flags = types.Float64Value(flagsValue)
		tflog.Debug(ctx, "Converted flags from Int64 to Float64", map[string]interface{}{
			"source_value": source.Flags.ValueInt64(),
			"target_value": flagsValue,
		})
	} else {
		target.Flags = types.Float64Null()
	}

	// Type conversion: Int64 → Float64 for key_tag
	if !source.KeyTag.IsNull() && !source.KeyTag.IsUnknown() {
		keyTagValue := float64(source.KeyTag.ValueInt64())
		target.KeyTag = types.Float64Value(keyTagValue)
		tflog.Debug(ctx, "Converted key_tag from Int64 to Float64", map[string]interface{}{
			"source_value": source.KeyTag.ValueInt64(),
			"target_value": keyTagValue,
		})
	} else {
		target.KeyTag = types.Float64Null()
	}

	// Date format conversion: RFC1123Z → RFC3339 for modified_on
	if !source.ModifiedOn.IsNull() && !source.ModifiedOn.IsUnknown() {
		modifiedOnStr := source.ModifiedOn.ValueString()
		if modifiedOnStr != "" {
			// Try to parse RFC1123Z format (v4 format)
			t, err := time.Parse(time.RFC1123Z, modifiedOnStr)
			if err != nil {
				// If parsing fails, try RFC3339 (might already be in correct format)
				t, err = time.Parse(time.RFC3339, modifiedOnStr)
				if err != nil {
					// If both fail, set to null and log warning
					tflog.Warn(ctx, "Failed to parse modified_on timestamp, setting to null", map[string]interface{}{
						"value": modifiedOnStr,
						"error": err.Error(),
					})
					target.ModifiedOn = types.StringNull()
				} else {
					// Already in RFC3339, use as-is
					target.ModifiedOn = types.StringValue(modifiedOnStr)
					tflog.Debug(ctx, "modified_on already in RFC3339 format", map[string]interface{}{
						"value": modifiedOnStr,
					})
				}
			} else {
				// Successfully parsed RFC1123Z, convert to RFC3339
				rfc3339Value := t.Format(time.RFC3339)
				target.ModifiedOn = types.StringValue(rfc3339Value)
				tflog.Debug(ctx, "Converted modified_on from RFC1123Z to RFC3339", map[string]interface{}{
					"source_format": modifiedOnStr,
					"target_format": rfc3339Value,
				})
			}
		} else {
			target.ModifiedOn = types.StringNull()
		}
	} else {
		target.ModifiedOn = types.StringNull()
	}

	// Status field: Set to null (optional-only field limitation)
	// Users will need to manually set status after migration if desired
	target.Status = types.StringNull()
	if !source.Status.IsNull() && !source.Status.IsUnknown() {
		tflog.Debug(ctx, "Setting status to null during migration (optional-only field limitation)", map[string]interface{}{
			"source_status": source.Status.ValueString(),
		})
	}

	// New v5 fields that don't exist in v4 - set to null
	target.DNSSECMultiSigner = types.BoolNull()
	target.DNSSECPresigned = types.BoolNull()
	target.DNSSECUseNsec3 = types.BoolNull()

	tflog.Debug(ctx, "Completed zone_dnssec state transformation", map[string]interface{}{
		"zone_id": target.ZoneID.ValueString(),
	})

	return target, diags
}

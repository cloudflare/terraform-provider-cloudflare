package v500

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Transform converts source (v4 SDKv2) state to target (v5 Framework) state.
//
// Key transformations:
// - hostnames: types.Set → *[]types.String (Set to List conversion)
// - requested_validity: types.Int64 → types.Float64 (type conversion + default 5475 if null)
// - min_days_for_renewal: Dropped (removed in v5)
//
// All other fields are direct copies.
func Transform(ctx context.Context, source SourceCloudflareOriginCACertificateModel) (*TargetOriginCACertificateModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	// Initialize target model
	target := &TargetOriginCACertificateModel{}

	// Step 1: Direct copies (no transformation)
	target.ID = source.ID
	target.Csr = source.Csr
	target.RequestType = source.RequestType
	target.Certificate = source.Certificate
	target.ExpiresOn = source.ExpiresOn

	// Step 2: Convert hostnames from Set to List
	if !source.Hostnames.IsNull() && !source.Hostnames.IsUnknown() {
		hostnames, err := convertSetToStringSlice(ctx, source.Hostnames)
		if err.HasError() {
			diags.Append(err...)
			return nil, diags
		}
		target.Hostnames = &hostnames
	} else {
		// If null/unknown, set to nil (null list)
		target.Hostnames = nil
	}

	// Step 3: Convert requested_validity from Int64 to Float64
	// Apply default value of 5475 if null/unknown
	if !source.RequestedValidity.IsNull() && !source.RequestedValidity.IsUnknown() {
		target.RequestedValidity = types.Float64Value(float64(source.RequestedValidity.ValueInt64()))
	} else {
		// Default value for requested_validity in v5 is 5475
		target.RequestedValidity = types.Float64Value(5475.0)
	}

	// Step 4: Drop deprecated fields
	// source.MinDaysForRenewal is intentionally not copied (removed in v5)

	return target, diags
}

// convertSetToStringSlice extracts string elements from a types.Set and converts to []types.String.
//
// This is used for converting the v4 hostnames field (TypeSet) to v5 format (*[]types.String).
// The order of elements may change as Sets are unordered, but this is acceptable.
func convertSetToStringSlice(ctx context.Context, set types.Set) ([]types.String, diag.Diagnostics) {
	var diags diag.Diagnostics

	// Extract Set elements to native Go []string first
	var rawStrings []string
	diags.Append(set.ElementsAs(ctx, &rawStrings, false)...)
	if diags.HasError() {
		return nil, diags
	}

	// Convert []string to []types.String
	result := make([]types.String, len(rawStrings))
	for i, str := range rawStrings {
		result[i] = types.StringValue(str)
	}

	return result, diags
}

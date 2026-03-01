// Package v500 implements state migration from legacy provider (v4) to current provider (v5)
// for the cloudflare_load_balancer_monitor resource.
package v500

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Transform converts source (legacy v4) state to target (current v5) state.
//
// Key transformations:
// 1. Header field: TypeSet (nested) → MapAttribute
// 2. Default values: Add v5 defaults for fields that were optional without defaults in v4
// 3. Direct copy: Pass-through fields (strings, ints, bools)
func Transform(ctx context.Context, source *SourceLoadBalancerMonitorModel) (*TargetLoadBalancerMonitorModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	tflog.Debug(ctx, "Starting transformation from v4 to v5")

	// Step 1: Validate required fields
	if source.AccountID.IsNull() || source.AccountID.IsUnknown() {
		diags.AddError(
			"Missing required field",
			"account_id is required for load_balancer_monitor migration. The source state is missing this field, which indicates corrupted state.",
		)
		return nil, diags
	}

	// Step 2: Initialize target with direct copies (pass-through fields)
	target := &TargetLoadBalancerMonitorModel{
		// Critical identifiers
		ID:        source.ID,
		AccountID: source.AccountID,

		// Numeric fields (Int64 compatible, no conversion needed)
		Interval: source.Interval,
		Retries:  source.Retries,
		Timeout:  source.Timeout,

		// Simple string fields
		Method: source.Method,
		Path:   source.Path,
		Type:   source.Type,

		// Computed timestamp fields
		CreatedOn:  source.CreatedOn,
		ModifiedOn: source.ModifiedOn,
	}

	// Step 2a: Handle optional Int64 fields where 0 means "unset" and should become null
	// These fields are Optional with no defaults, and the API returns null when not configured.
	// Converting 0 → null prevents drift after migration.

	// Port: 0 is not a valid port number, so 0 means "use default" → null
	if !source.Port.IsNull() && source.Port.ValueInt64() == 0 {
		target.Port = types.Int64Null()
		tflog.Debug(ctx, "Converted port from 0 to null (0 is not a valid port)")
	} else {
		target.Port = source.Port
	}

	// ConsecutiveDown: 0 means "not configured" → null
	if !source.ConsecutiveDown.IsNull() && source.ConsecutiveDown.ValueInt64() == 0 {
		target.ConsecutiveDown = types.Int64Null()
		tflog.Debug(ctx, "Converted consecutive_down from 0 to null (0 means unconfigured)")
	} else {
		target.ConsecutiveDown = source.ConsecutiveDown
	}

	// ConsecutiveUp: 0 means "not configured" → null
	if !source.ConsecutiveUp.IsNull() && source.ConsecutiveUp.ValueInt64() == 0 {
		target.ConsecutiveUp = types.Int64Null()
		tflog.Debug(ctx, "Converted consecutive_up from 0 to null (0 means unconfigured)")
	} else {
		target.ConsecutiveUp = source.ConsecutiveUp
	}

	// Step 3: Add v5 defaults for fields that were optional without defaults in v4
	// These defaults prevent unnecessary PATCH operations after migration

	// String fields with default ""
	if source.Description.IsNull() {
		target.Description = types.StringValue("")
		tflog.Debug(ctx, "Set description default: empty string")
	} else {
		target.Description = source.Description
	}

	if source.ExpectedBody.IsNull() {
		target.ExpectedBody = types.StringValue("")
		tflog.Debug(ctx, "Set expected_body default: empty string")
	} else {
		target.ExpectedBody = source.ExpectedBody
	}

	if source.ExpectedCodes.IsNull() {
		target.ExpectedCodes = types.StringValue("")
		tflog.Debug(ctx, "Set expected_codes default: empty string")
	} else {
		target.ExpectedCodes = source.ExpectedCodes
	}

	if source.ProbeZone.IsNull() {
		target.ProbeZone = types.StringValue("")
		tflog.Debug(ctx, "Set probe_zone default: empty string")
	} else {
		target.ProbeZone = source.ProbeZone
	}

	// Bool fields with default false
	if source.AllowInsecure.IsNull() {
		target.AllowInsecure = types.BoolValue(false)
		tflog.Debug(ctx, "Set allow_insecure default: false")
	} else {
		target.AllowInsecure = source.AllowInsecure
	}

	if source.FollowRedirects.IsNull() {
		target.FollowRedirects = types.BoolValue(false)
		tflog.Debug(ctx, "Set follow_redirects default: false")
	} else {
		target.FollowRedirects = source.FollowRedirects
	}

	// Step 4: Transform header field (TypeSet → MapAttribute)
	// This is the primary transformation challenge
	if !source.Header.IsNull() && !source.Header.IsUnknown() {
		headerMap, headerDiags := transformHeader(ctx, source.Header)
		diags.Append(headerDiags...)
		if diags.HasError() {
			tflog.Error(ctx, "Failed to transform header field")
			return nil, diags
		}
		target.Header = headerMap
		tflog.Debug(ctx, "Transformed header field from Set to Map")
	} else {
		target.Header = nil
		tflog.Debug(ctx, "No header field to transform (null or unknown)")
	}

	tflog.Info(ctx, "Transformation from v4 to v5 completed successfully")
	return target, diags
}

// transformHeader converts v4 header TypeSet structure to v5 MapAttribute structure.
//
// v4 structure (TypeSet with nested Resource):
//   header {
//     header = "Host"
//     values = ["example.com"]
//   }
//
// v4 state storage:
//   [{"header": "Host", "values": ["example.com"]}, {"header": "User-Agent", "values": ["Bot"]}]
//
// v5 structure (MapAttribute with ListType element):
//   header = {
//     "Host"       = ["example.com"]
//     "User-Agent" = ["Bot"]
//   }
//
// v5 model type: *map[string]*[]types.String
func transformHeader(ctx context.Context, sourceHeader types.Set) (*map[string]*[]types.String, diag.Diagnostics) {
	var diags diag.Diagnostics

	tflog.Debug(ctx, "Starting header transformation")

	// Extract Set elements to slice of SourceHeaderItem structs
	var headerItems []SourceHeaderItem
	diags.Append(sourceHeader.ElementsAs(ctx, &headerItems, false)...)
	if diags.HasError() {
		tflog.Error(ctx, "Failed to extract header Set elements")
		return nil, diags
	}

	if len(headerItems) == 0 {
		tflog.Debug(ctx, "No header items to transform")
		return nil, diags
	}

	tflog.Debug(ctx, "Extracted header items", map[string]interface{}{
		"count": len(headerItems),
	})

	// Build the target map structure
	headerMap := make(map[string]*[]types.String)

	for i, item := range headerItems {
		// Get header name
		headerName := item.Header.ValueString()
		if headerName == "" {
			tflog.Warn(ctx, "Skipping header item with empty name", map[string]interface{}{
				"index": i,
			})
			continue
		}

		// Extract values Set to slice of types.String
		var values []types.String
		diags.Append(item.Values.ElementsAs(ctx, &values, false)...)
		if diags.HasError() {
			tflog.Error(ctx, "Failed to extract header values", map[string]interface{}{
				"header_name": headerName,
			})
			return nil, diags
		}

		if len(values) == 0 {
			tflog.Warn(ctx, "Skipping header with no values", map[string]interface{}{
				"header_name": headerName,
			})
			continue
		}

		// Store in map (pointer to slice as per v5 model)
		headerMap[headerName] = &values

		tflog.Debug(ctx, "Transformed header item", map[string]interface{}{
			"header_name":  headerName,
			"values_count": len(values),
		})
	}

	if len(headerMap) == 0 {
		tflog.Debug(ctx, "No valid headers after transformation")
		return nil, diags
	}

	tflog.Info(ctx, "Header transformation completed", map[string]interface{}{
		"headers_count": len(headerMap),
	})

	return &headerMap, diags
}

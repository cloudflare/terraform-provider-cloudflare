package v500

import (
	"context"
	"strings"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/migrations"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Transform converts source (v4 flat structure) state to target (v5 nested structure) state.
//
// Major transformations:
// - Flat HTTP/TCP fields → nested http_config or tcp_config based on type
// - Header Set of objects → Map
// - CheckRegions List → *[]types.String
//
// This function is called by UpgradeFromV4 handler.
func Transform(ctx context.Context, source SourceHealthcheckModel) (*TargetHealthcheckModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	// Step 1: Validate required fields
	if source.ZoneID.IsNull() || source.ZoneID.IsUnknown() {
		diags.AddError(
			"Missing required field",
			"zone_id is required for healthcheck migration. The source state is missing this field, which indicates corrupted state.",
		)
		return nil, diags
	}
	if source.Name.IsNull() || source.Name.IsUnknown() {
		diags.AddError(
			"Missing required field",
			"name is required for healthcheck migration. The source state is missing this field, which indicates corrupted state.",
		)
		return nil, diags
	}
	if source.Address.IsNull() || source.Address.IsUnknown() {
		diags.AddError(
			"Missing required field",
			"address is required for healthcheck migration. The source state is missing this field, which indicates corrupted state.",
		)
		return nil, diags
	}

	// Step 2: Initialize target with direct copies
	target := &TargetHealthcheckModel{
		ID:                   source.ID,
		ZoneID:               source.ZoneID,
		Name:                 source.Name,
		Description:          migrations.FalseyStringToNull(source.Description),
		Address:              source.Address,
		Type:                 source.Type,
		Suspended:            source.Suspended,
		ConsecutiveFails:     source.ConsecutiveFails,
		ConsecutiveSuccesses: source.ConsecutiveSuccesses,
		Interval:             source.Interval,
		Retries:              source.Retries,
		Timeout:              source.Timeout,
	}

	// Convert CheckRegions from List to *[]types.String
	if !source.CheckRegions.IsNull() && !source.CheckRegions.IsUnknown() {
		var regions []types.String
		diags.Append(source.CheckRegions.ElementsAs(ctx, &regions, false)...)
		if diags.HasError() {
			return nil, diags
		}
		target.CheckRegions = &regions
	}

	// Convert timestamps (String → RFC3339)
	// The framework handles the conversion automatically via the timetypes.RFC3339 type
	if !source.CreatedOn.IsNull() && !source.CreatedOn.IsUnknown() {
		createdOn, parseDiags := timetypes.NewRFC3339Value(source.CreatedOn.ValueString())
		diags.Append(parseDiags...)
		if !parseDiags.HasError() {
			target.CreatedOn = createdOn
		}
	}
	if !source.ModifiedOn.IsNull() && !source.ModifiedOn.IsUnknown() {
		modifiedOn, parseDiags := timetypes.NewRFC3339Value(source.ModifiedOn.ValueString())
		diags.Append(parseDiags...)
		if !parseDiags.HasError() {
			target.ModifiedOn = modifiedOn
		}
	}

	// Step 3: Determine healthcheck type and create appropriate config
	healthcheckType := "HTTP" // Default if not specified
	if !source.Type.IsNull() && !source.Type.IsUnknown() {
		healthcheckType = strings.ToUpper(source.Type.ValueString())
	}

	tflog.Debug(ctx, "Transforming healthcheck config", map[string]interface{}{
		"type": healthcheckType,
	})

	// Branch based on type
	if healthcheckType == "TCP" {
		// Create TCP config
		tcpConfig, tcpDiags := transformTCPConfig(ctx, source)
		diags.Append(tcpDiags...)
		if diags.HasError() {
			return nil, diags
		}
		target.TCPConfig = tcpConfig

		tflog.Debug(ctx, "Created tcp_config", map[string]interface{}{
			"has_method": !source.Method.IsNull(),
			"has_port":   !source.Port.IsNull(),
		})
	} else {
		// Create HTTP config (for HTTP and HTTPS types)
		httpConfig, httpDiags := transformHTTPConfig(ctx, source)
		diags.Append(httpDiags...)
		if diags.HasError() {
			return nil, diags
		}
		target.HTTPConfig = httpConfig

		tflog.Debug(ctx, "Created http_config", map[string]interface{}{
			"has_method":  !source.Method.IsNull(),
			"has_port":    !source.Port.IsNull(),
			"has_path":    !source.Path.IsNull(),
			"has_headers": !source.Header.IsNull(),
		})
	}

	return target, diags
}

// transformHTTPConfig creates http_config nested object from flat v4 fields.
//
// Moves these fields from root to http_config:
// - method, port, path, expected_codes, expected_body
// - follow_redirects, allow_insecure
// - header (Set → Map transformation)
func transformHTTPConfig(ctx context.Context, source SourceHealthcheckModel) (customfield.NestedObject[TargetHTTPConfigModel], diag.Diagnostics) {
	var diags diag.Diagnostics

	// Create HTTP config model with values from source
	httpConfig := TargetHTTPConfigModel{
		Method:          source.Method,
		Port:            source.Port,
		Path:            source.Path,
		ExpectedBody:    source.ExpectedBody,
		FollowRedirects: source.FollowRedirects,
		AllowInsecure:   source.AllowInsecure,
	}

	// Convert ExpectedCodes List to *[]types.String
	if !source.ExpectedCodes.IsNull() && !source.ExpectedCodes.IsUnknown() {
		var codes []types.String
		diags.Append(source.ExpectedCodes.ElementsAs(ctx, &codes, false)...)
		if diags.HasError() {
			return customfield.NullObject[TargetHTTPConfigModel](ctx), diags
		}
		httpConfig.ExpectedCodes = &codes
	}

	// Transform header Set → Map
	headerMap, headerDiags := transformHeaderSetToMap(ctx, source.Header)
	diags.Append(headerDiags...)
	if diags.HasError() {
		return customfield.NullObject[TargetHTTPConfigModel](ctx), diags
	}
	httpConfig.Header = headerMap

	// Wrap in customfield.NestedObject
	nestedObj, nestedDiags := customfield.NewObject(ctx, &httpConfig)
	diags.Append(nestedDiags...)
	if diags.HasError() {
		return customfield.NullObject[TargetHTTPConfigModel](ctx), diags
	}

	return nestedObj, diags
}

// transformTCPConfig creates tcp_config nested object from flat v4 fields.
//
// Moves these fields from root to tcp_config:
// - method, port
func transformTCPConfig(ctx context.Context, source SourceHealthcheckModel) (customfield.NestedObject[TargetTCPConfigModel], diag.Diagnostics) {
	var diags diag.Diagnostics

	// Create TCP config model with values from source
	tcpConfig := TargetTCPConfigModel{
		Method: source.Method,
		Port:   source.Port,
	}

	// Wrap in customfield.NestedObject
	nestedObj, nestedDiags := customfield.NewObject(ctx, &tcpConfig)
	diags.Append(nestedDiags...)
	if diags.HasError() {
		return customfield.NullObject[TargetTCPConfigModel](ctx), diags
	}

	return nestedObj, diags
}

// transformHeaderSetToMap converts v4 header Set structure to v5 Map structure.
//
// v4 format (Set of objects):
//
//	[{header: "Host", values: ["example.com"]}, {header: "User-Agent", values: ["Bot"]}]
//
// v5 format (Map):
//
//	{"Host": ["example.com"], "User-Agent": ["Bot"]}
//
// Returns nil if source header Set is null/empty.
func transformHeaderSetToMap(ctx context.Context, headerSet types.Set) (*map[string]*[]types.String, diag.Diagnostics) {
	var diags diag.Diagnostics

	// If header Set is null or unknown, return nil
	if headerSet.IsNull() || headerSet.IsUnknown() {
		return nil, diags
	}

	// Extract header items from Set
	var headerItems []SourceHeaderModel
	diags.Append(headerSet.ElementsAs(ctx, &headerItems, false)...)
	if diags.HasError() {
		return nil, diags
	}

	// If no headers, return nil
	if len(headerItems) == 0 {
		return nil, diags
	}

	// Convert to map
	headerMap := make(map[string]*[]types.String)

	for _, item := range headerItems {
		// Skip if header name is missing
		if item.Header.IsNull() || item.Header.IsUnknown() {
			tflog.Warn(ctx, "Skipping header with missing name")
			continue
		}

		headerName := item.Header.ValueString()

		// Skip if values are missing
		if item.Values.IsNull() || item.Values.IsUnknown() {
			tflog.Warn(ctx, "Skipping header with missing values", map[string]interface{}{
				"header": headerName,
			})
			continue
		}

		// Extract values from Set to []types.String
		var values []types.String
		valuesDiags := item.Values.ElementsAs(ctx, &values, false)
		if valuesDiags.HasError() {
			// Log warning but continue processing other headers
			tflog.Warn(ctx, "Failed to extract values for header, skipping", map[string]interface{}{
				"header": headerName,
				"error":  valuesDiags.Errors()[0].Summary(),
			})
			continue
		}

		// Add to map
		headerMap[headerName] = &values

		tflog.Debug(ctx, "Transformed header", map[string]interface{}{
			"name":        headerName,
			"value_count": len(values),
		})
	}

	// If we ended up with no valid headers, return nil
	if len(headerMap) == 0 {
		return nil, diags
	}

	return &headerMap, diags
}

package v500

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ============================================================================
// Source Models (Legacy Provider - v4.x - SDKv2)
// ============================================================================

// SourceHealthcheckModel represents the legacy healthcheck state from v4.x provider.
// Schema version: 0 (implicit in v4)
// Resource type: cloudflare_healthcheck
//
// IMPORTANT: v4 has FLAT structure - all HTTP/TCP fields are at root level.
// The v5 structure nests these into http_config or tcp_config based on type.
type SourceHealthcheckModel struct {
	// Core fields (stay at root in both v4 and v5)
	ID                   types.String `tfsdk:"id"`
	ZoneID               types.String `tfsdk:"zone_id"`
	Name                 types.String `tfsdk:"name"`
	Description          types.String `tfsdk:"description"`
	Address              types.String `tfsdk:"address"`
	Type                 types.String `tfsdk:"type"`
	CheckRegions         types.List   `tfsdk:"check_regions"` // List of strings
	Suspended            types.Bool   `tfsdk:"suspended"`
	ConsecutiveFails     types.Int64  `tfsdk:"consecutive_fails"`
	ConsecutiveSuccesses types.Int64  `tfsdk:"consecutive_successes"`
	Interval             types.Int64  `tfsdk:"interval"`
	Retries              types.Int64  `tfsdk:"retries"`
	Timeout              types.Int64  `tfsdk:"timeout"`
	CreatedOn            types.String `tfsdk:"created_on"`
	ModifiedOn           types.String `tfsdk:"modified_on"`

	// FLAT HTTP/TCP fields (at root in v4, will move to nested config in v5)
	// These fields exist at root level in v4 and need to be moved into
	// http_config or tcp_config based on the type field value.
	Method          types.String `tfsdk:"method"`           // → http_config.method OR tcp_config.method
	Port            types.Int64  `tfsdk:"port"`             // → http_config.port OR tcp_config.port
	Path            types.String `tfsdk:"path"`             // → http_config.path (HTTP only)
	ExpectedCodes   types.List   `tfsdk:"expected_codes"`   // → http_config.expected_codes (HTTP only)
	ExpectedBody    types.String `tfsdk:"expected_body"`    // → http_config.expected_body (HTTP only)
	FollowRedirects types.Bool   `tfsdk:"follow_redirects"` // → http_config.follow_redirects (HTTP only)
	AllowInsecure   types.Bool   `tfsdk:"allow_insecure"`   // → http_config.allow_insecure (HTTP only)

	// Header as Set (v4 format) - will convert to Map in v5
	// v4: Set of objects with {header: string, values: []string}
	// v5: Map[string][]string
	Header types.Set `tfsdk:"header"` // → http_config.header (as Map)
}

// SourceHeaderModel represents a single header entry in the v4 Set structure.
// v4 stores headers as a Set of objects with "header" and "values" fields.
// v5 converts this to a Map[string][]string.
type SourceHeaderModel struct {
	Header types.String `tfsdk:"header"` // Header name (e.g., "Host")
	Values types.Set    `tfsdk:"values"` // Set of string values
}

// ============================================================================
// Target Models (Current Provider - v5.x+ - Framework)
// ============================================================================

// TargetHealthcheckModel represents the current healthcheck state from v5.x+ provider.
// Schema version: 500
// Resource type: cloudflare_healthcheck
//
// This matches the HealthcheckModel in internal/services/healthcheck/model.go
type TargetHealthcheckModel struct {
	ID                   types.String                                         `tfsdk:"id"`
	ZoneID               types.String                                         `tfsdk:"zone_id"`
	Address              types.String                                         `tfsdk:"address"`
	Name                 types.String                                         `tfsdk:"name"`
	Description          types.String                                         `tfsdk:"description"`
	CheckRegions         *[]types.String                                      `tfsdk:"check_regions"`
	ConsecutiveFails     types.Int64                                          `tfsdk:"consecutive_fails"`
	ConsecutiveSuccesses types.Int64                                          `tfsdk:"consecutive_successes"`
	Interval             types.Int64                                          `tfsdk:"interval"`
	Retries              types.Int64                                          `tfsdk:"retries"`
	Suspended            types.Bool                                           `tfsdk:"suspended"`
	Timeout              types.Int64                                          `tfsdk:"timeout"`
	Type                 types.String                                         `tfsdk:"type"`
	HTTPConfig           customfield.NestedObject[TargetHTTPConfigModel]      `tfsdk:"http_config"`
	TCPConfig            customfield.NestedObject[TargetTCPConfigModel]       `tfsdk:"tcp_config"`
	CreatedOn            timetypes.RFC3339                                    `tfsdk:"created_on"`
	FailureReason        types.String                                         `tfsdk:"failure_reason"`
	ModifiedOn           timetypes.RFC3339                                    `tfsdk:"modified_on"`
	Status               types.String                                         `tfsdk:"status"`
}

// TargetHTTPConfigModel represents the http_config nested object in v5.
// This matches HealthcheckHTTPConfigModel in internal/services/healthcheck/model.go
type TargetHTTPConfigModel struct {
	AllowInsecure   types.Bool                  `tfsdk:"allow_insecure"`
	ExpectedBody    types.String                `tfsdk:"expected_body"`
	ExpectedCodes   *[]types.String             `tfsdk:"expected_codes"`
	FollowRedirects types.Bool                  `tfsdk:"follow_redirects"`
	Header          *map[string]*[]types.String `tfsdk:"header"` // Map format in v5
	Method          types.String                `tfsdk:"method"`
	Path            types.String                `tfsdk:"path"`
	Port            types.Int64                 `tfsdk:"port"`
}

// TargetTCPConfigModel represents the tcp_config nested object in v5.
// This matches HealthcheckTCPConfigModel in internal/services/healthcheck/model.go
type TargetTCPConfigModel struct {
	Method types.String `tfsdk:"method"`
	Port   types.Int64  `tfsdk:"port"`
}

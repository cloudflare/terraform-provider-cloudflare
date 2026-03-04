// Package v500 implements state migration from legacy provider (v4) to current provider (v5)
// for the cloudflare_load_balancer_monitor resource.
//
// Schema version: 0 (v4) → 500 (v5)
// Resource name: cloudflare_load_balancer_monitor (unchanged)
package v500

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ============================================================================
// Source Models (Legacy Provider - v4.x)
// ============================================================================

// SourceLoadBalancerMonitorModel represents the legacy resource state from v4.x provider.
// Schema version: 0 (implicit in SDKv2)
// Resource type: cloudflare_load_balancer_monitor
//
// This model matches the v4 schema structure with SDKv2 types.
type SourceLoadBalancerMonitorModel struct {
	// Required fields
	ID        types.String `tfsdk:"id"`
	AccountID types.String `tfsdk:"account_id"`

	// Optional numeric fields (TypeInt in v4)
	// These will be converted to Int64 in v5 (but stored as float64 in JSON)
	Interval        types.Int64 `tfsdk:"interval"`
	Port            types.Int64 `tfsdk:"port"`
	Retries         types.Int64 `tfsdk:"retries"`
	Timeout         types.Int64 `tfsdk:"timeout"`
	ConsecutiveDown types.Int64 `tfsdk:"consecutive_down"`
	ConsecutiveUp   types.Int64 `tfsdk:"consecutive_up"`

	// Optional string fields
	Description   types.String `tfsdk:"description"`
	Method        types.String `tfsdk:"method"`
	Path          types.String `tfsdk:"path"`
	ProbeZone     types.String `tfsdk:"probe_zone"`
	Type          types.String `tfsdk:"type"`
	ExpectedBody  types.String `tfsdk:"expected_body"`
	ExpectedCodes types.String `tfsdk:"expected_codes"`

	// Optional boolean fields
	AllowInsecure   types.Bool `tfsdk:"allow_insecure"`
	FollowRedirects types.Bool `tfsdk:"follow_redirects"`

	// Computed timestamp fields
	CreatedOn  types.String `tfsdk:"created_on"`
	ModifiedOn types.String `tfsdk:"modified_on"`

	// Complex nested field (TypeSet in v4)
	// Stored as array of objects in state: [{"header": "Host", "values": ["example.com"]}]
	Header types.Set `tfsdk:"header"`
}

// SourceHeaderItem represents a single header entry from the v4 provider's TypeSet.
// This is the nested structure within the header Set.
//
// v4 schema:
//   header {
//     header = "Host"
//     values = ["example.com"]
//   }
//
// v4 state storage:
//   [{"header": "Host", "values": ["example.com"]}]
type SourceHeaderItem struct {
	Header types.String `tfsdk:"header"` // The header name (e.g., "Host", "User-Agent")
	Values types.Set    `tfsdk:"values"` // TypeSet of strings (e.g., ["example.com"])
}

// ============================================================================
// Target Models (Current Provider - v5.x+)
// ============================================================================

// TargetLoadBalancerMonitorModel represents the current resource state from v5.x+ provider.
// Schema version: 500 (controlled by migrations.schema version)
// Resource type: cloudflare_load_balancer_monitor
//
// This model matches the model in the parent package's model.go file.
// We duplicate it here to keep the migration package self-contained.
type TargetLoadBalancerMonitorModel struct {
	// Required fields
	ID        types.String `tfsdk:"id"`
	AccountID types.String `tfsdk:"account_id"`

	// Optional numeric fields (Int64Attribute in v5)
	ConsecutiveDown types.Int64 `tfsdk:"consecutive_down"`
	ConsecutiveUp   types.Int64 `tfsdk:"consecutive_up"`
	Port            types.Int64 `tfsdk:"port"`
	Interval        types.Int64 `tfsdk:"interval"`
	Retries         types.Int64 `tfsdk:"retries"`
	Timeout         types.Int64 `tfsdk:"timeout"`

	// Complex field: Map instead of Set (MAJOR CHANGE from v4)
	// v5 schema: MapAttribute with ListType element
	// v5 state: {"Host": ["example.com"], "User-Agent": ["Bot"]}
	// Model type: pointer to map of string to pointer to slice of types.String
	Header *map[string]*[]types.String `tfsdk:"header"`

	// Optional boolean fields (with defaults in v5)
	AllowInsecure   types.Bool `tfsdk:"allow_insecure"`
	FollowRedirects types.Bool `tfsdk:"follow_redirects"`

	// Optional string fields (with defaults in v5)
	Description   types.String `tfsdk:"description"`
	ExpectedBody  types.String `tfsdk:"expected_body"`
	ExpectedCodes types.String `tfsdk:"expected_codes"`
	Method        types.String `tfsdk:"method"`
	Path          types.String `tfsdk:"path"`
	ProbeZone     types.String `tfsdk:"probe_zone"`
	Type          types.String `tfsdk:"type"`

	// Computed timestamp fields
	CreatedOn  types.String `tfsdk:"created_on"`
	ModifiedOn types.String `tfsdk:"modified_on"`
}

package v500

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ============================================================================
// Source Models (Legacy Provider - v4.x)
// ============================================================================

// SourceCloudflareDeviceDexTestModel represents the legacy resource state from v4.x provider.
// Schema version: 0 (SDKv2 default)
// Resource type: cloudflare_device_dex_test
//
// This model represents the v4 state structure where the resource was named
// cloudflare_device_dex_test (deprecated name). The v4 provider also had
// cloudflare_zero_trust_dex_test as an alias with the same schema.
type SourceCloudflareDeviceDexTestModel struct {
	ID          types.String             `tfsdk:"id"`
	AccountID   types.String             `tfsdk:"account_id"`
	Name        types.String             `tfsdk:"name"`
	Description types.String             `tfsdk:"description"`
	Interval    types.String             `tfsdk:"interval"`
	Enabled     types.Bool               `tfsdk:"enabled"`
	Data        []SourceDEXTestDataModel `tfsdk:"data"` // TypeList MaxItems:1 stored as array
	Updated     types.String             `tfsdk:"updated"`
	Created     types.String             `tfsdk:"created"`
}

// SourceDEXTestDataModel represents the nested data block structure from v4.x provider.
// In the legacy provider, this is a TypeList with MaxItems: 1, so it's stored as an array in state.
type SourceDEXTestDataModel struct {
	Kind   types.String `tfsdk:"kind"`
	Host   types.String `tfsdk:"host"`
	Method types.String `tfsdk:"method"` // Optional: only present for kind="http"
}

// ============================================================================
// Target Models (Current Provider - v5.x+)
// ============================================================================

// ZeroTrustDEXTestModel represents the current resource state from v5.x+ provider.
// Schema version: 500 (with migrations enabled) or 1 (base v5)
// Resource type: cloudflare_zero_trust_dex_test
//
// Note: This matches the model in the parent package's model.go file.
// We duplicate it here to keep the migration package self-contained.
type ZeroTrustDEXTestModel struct {
	ID             types.String                                                      `tfsdk:"id"`
	TestID         types.String                                                      `tfsdk:"test_id"`
	AccountID      types.String                                                      `tfsdk:"account_id"`
	Enabled        types.Bool                                                        `tfsdk:"enabled"`
	Interval       types.String                                                      `tfsdk:"interval"`
	Name           types.String                                                      `tfsdk:"name"`
	Data           *ZeroTrustDEXTestDataModel                                        `tfsdk:"data"`
	Description    types.String                                                      `tfsdk:"description"`
	TargetPolicies customfield.NestedObjectList[ZeroTrustDEXTestTargetPoliciesModel] `tfsdk:"target_policies"`
	Targeted       types.Bool                                                        `tfsdk:"targeted"`
}

// ZeroTrustDEXTestDataModel represents the nested data structure from v5.x+ provider.
// In v5, this is a SingleNestedAttribute, so it's a pointer to a single object.
type ZeroTrustDEXTestDataModel struct {
	Host   types.String `tfsdk:"host"`
	Kind   types.String `tfsdk:"kind"`
	Method types.String `tfsdk:"method"`
}

// ZeroTrustDEXTestTargetPoliciesModel represents the target_policies nested structure.
// This is a new field in v5 that didn't exist in v4.
type ZeroTrustDEXTestTargetPoliciesModel struct {
	ID      types.String `tfsdk:"id"`
	Default types.Bool   `tfsdk:"default"`
	Name    types.String `tfsdk:"name"`
}

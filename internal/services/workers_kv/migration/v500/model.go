package v500

import (
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ============================================================================
// Source Models (Legacy Provider - v4.x)
// ============================================================================

// SourceCloudflareWorkersKVModel represents the legacy resource state from v4.x provider.
// Schema version: 0 (SDKv2 default)
// Resource type: cloudflare_workers_kv
//
// This matches the v4 SDKv2 schema which used "key" instead of "key_name".
type SourceCloudflareWorkersKVModel struct {
	ID          types.String `tfsdk:"id"`
	AccountID   types.String `tfsdk:"account_id"`
	NamespaceID types.String `tfsdk:"namespace_id"`
	Key         types.String `tfsdk:"key"`   // v4 field name (renamed to key_name in v5)
	Value       types.String `tfsdk:"value"`
}

// ============================================================================
// Target Models (Current Provider - v5.x+)
// ============================================================================

// TargetWorkersKVModel represents the current resource state from v5.x+ provider.
// Schema version: 500
// Resource type: cloudflare_workers_kv
//
// This matches the v5 Plugin Framework schema with "key_name" and optional "metadata".
// Note: This should match WorkersKVModel in the parent package's model.go file.
type TargetWorkersKVModel struct {
	ID          types.String         `tfsdk:"id"`
	KeyName     types.String         `tfsdk:"key_name"`   // Renamed from "key" in v4
	AccountID   types.String         `tfsdk:"account_id"`
	NamespaceID types.String         `tfsdk:"namespace_id"`
	Value       types.String         `tfsdk:"value"`
	Metadata    jsontypes.Normalized `tfsdk:"metadata"`   // New in v5, optional
}

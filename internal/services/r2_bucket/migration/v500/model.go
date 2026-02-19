package v500

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ============================================================================
// Source Models (Legacy Provider - v4.x)
// ============================================================================

// SourceR2BucketModel represents the legacy resource state from v4.x provider.
// Schema version: 0 (Framework resource, no versioning in v4)
// Resource type: cloudflare_r2_bucket
//
// v4 schema only had 4 fields: id, name, account_id, location
type SourceR2BucketModel struct {
	ID        types.String `tfsdk:"id"`
	Name      types.String `tfsdk:"name"`
	AccountID types.String `tfsdk:"account_id"`
	Location  types.String `tfsdk:"location"`
}

// ============================================================================
// Target Models (Current Provider - v5.x+)
// ============================================================================

// TargetR2BucketModel represents the current resource state from v5.x+ provider.
// Schema version: 500
// Resource type: cloudflare_r2_bucket
//
// v5 schema added 3 new fields: jurisdiction, storage_class, creation_date
type TargetR2BucketModel struct {
	ID           types.String `tfsdk:"id"`
	Name         types.String `tfsdk:"name"`
	AccountID    types.String `tfsdk:"account_id"`
	Location     types.String `tfsdk:"location"`
	Jurisdiction types.String `tfsdk:"jurisdiction"`  // NEW in v5: default "default"
	StorageClass types.String `tfsdk:"storage_class"` // NEW in v5: default "Standard"
	CreationDate types.String `tfsdk:"creation_date"` // NEW in v5: computed, API-assigned
}

package v500

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// SourceModel mirrors the prior (implicit schema_version=0) shape from main:
// id was Int64; the v500 schema requires String to align with the v7 SDK.
type SourceModel struct {
	ID           types.Int64       `tfsdk:"id"`
	AccountID    types.String      `tfsdk:"account_id"`
	IsRegex      types.Bool        `tfsdk:"is_regex"`
	Pattern      types.String      `tfsdk:"pattern"`
	PatternType  types.String      `tfsdk:"pattern_type"`
	Comments     types.String      `tfsdk:"comments"`
	CreatedAt    timetypes.RFC3339 `tfsdk:"created_at"`
	LastModified timetypes.RFC3339 `tfsdk:"last_modified"`
}

// TargetModel mirrors the v500 schema.
type TargetModel struct {
	ID           types.String      `tfsdk:"id"`
	AccountID    types.String      `tfsdk:"account_id"`
	IsRegex      types.Bool        `tfsdk:"is_regex"`
	Pattern      types.String      `tfsdk:"pattern"`
	PatternType  types.String      `tfsdk:"pattern_type"`
	Comments     types.String      `tfsdk:"comments"`
	CreatedAt    timetypes.RFC3339 `tfsdk:"created_at"`
	LastModified timetypes.RFC3339 `tfsdk:"last_modified"`
	ModifiedAt   timetypes.RFC3339 `tfsdk:"modified_at"`
}

package v500

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// SourceModel is the prior (implicit schema_version=0) shape from main:
// id is Int64 and the request body lives in a nested `body` list block.
type SourceModel struct {
	ID           types.Int64       `tfsdk:"id"`
	AccountID    types.String      `tfsdk:"account_id"`
	Body         *[]*SourceBody    `tfsdk:"body"`
	Comments     types.String      `tfsdk:"comments"`
	IsRecent     types.Bool        `tfsdk:"is_recent"`
	IsRegex      types.Bool        `tfsdk:"is_regex"`
	IsSimilarity types.Bool        `tfsdk:"is_similarity"`
	Pattern      types.String      `tfsdk:"pattern"`
	CreatedAt    timetypes.RFC3339 `tfsdk:"created_at"`
	LastModified timetypes.RFC3339 `tfsdk:"last_modified"`
}

type SourceBody struct {
	IsRecent     types.Bool   `tfsdk:"is_recent"`
	IsRegex      types.Bool   `tfsdk:"is_regex"`
	IsSimilarity types.Bool   `tfsdk:"is_similarity"`
	Pattern      types.String `tfsdk:"pattern"`
	Comments     types.String `tfsdk:"comments"`
}

// TargetModel mirrors the v500 schema: id is String and trusted-domain fields
// are flat at the top level (no `body` block).
type TargetModel struct {
	ID           types.String      `tfsdk:"id"`
	AccountID    types.String      `tfsdk:"account_id"`
	IsRecent     types.Bool        `tfsdk:"is_recent"`
	IsRegex      types.Bool        `tfsdk:"is_regex"`
	IsSimilarity types.Bool        `tfsdk:"is_similarity"`
	Pattern      types.String      `tfsdk:"pattern"`
	Comments     types.String      `tfsdk:"comments"`
	CreatedAt    timetypes.RFC3339 `tfsdk:"created_at"`
	LastModified timetypes.RFC3339 `tfsdk:"last_modified"`
	ModifiedAt   timetypes.RFC3339 `tfsdk:"modified_at"`
}

package v500

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ============================================================================
// Source Models (Legacy Provider - v4.x)
// ============================================================================

// SourceCloudflareSnippetModel represents the legacy cloudflare_snippet state from v4.x provider.
// Schema version: 0 (SDKv2 default)
// Resource type: cloudflare_snippet
//
// In v4, the snippet resource had:
// - name: The snippet name (renamed to snippet_name in v5)
// - main_module: Top-level attribute (moved to metadata.main_module in v5)
// - files: Block-style nested resources (converted to list attribute in v5)
// - created_on, modified_on: Computed timestamp strings
type SourceCloudflareSnippetModel struct {
	ZoneID     types.String             `tfsdk:"zone_id"`
	Name       types.String             `tfsdk:"name"`
	MainModule types.String             `tfsdk:"main_module"`
	Files      []SourceSnippetFileModel `tfsdk:"files"`
	CreatedOn  types.String             `tfsdk:"created_on"`
	ModifiedOn types.String             `tfsdk:"modified_on"`
}

// SourceSnippetFileModel represents a file block from the v4 snippet resource.
type SourceSnippetFileModel struct {
	Name    types.String `tfsdk:"name"`
	Content types.String `tfsdk:"content"`
}

// ============================================================================
// Target Models (Current Provider - v5.x+)
// ============================================================================

// TargetSnippetModel represents the current cloudflare_snippet state from v5.x+ provider.
// Schema version: 500
// Resource type: cloudflare_snippet
//
// Types must match the v5 schema exactly so that resp.State.Set() succeeds:
// - Files uses *[]TargetSnippetFileModel (framework converts to SnippetFileType via reflection)
// - Timestamps use timetypes.RFC3339 (matches schema CustomType)
type TargetSnippetModel struct {
	SnippetName types.String                `tfsdk:"snippet_name"`
	ZoneID      types.String                `tfsdk:"zone_id"`
	Files       *[]TargetSnippetFileModel   `tfsdk:"files"`
	Metadata    *TargetSnippetMetadataModel `tfsdk:"metadata"`
	CreatedOn   timetypes.RFC3339           `tfsdk:"created_on"`
	ModifiedOn  timetypes.RFC3339           `tfsdk:"modified_on"`
}

// TargetSnippetFileModel represents a file in the v5 files list attribute.
type TargetSnippetFileModel struct {
	Name    types.String `tfsdk:"name"`
	Content types.String `tfsdk:"content"`
}

// TargetSnippetMetadataModel represents the metadata nested object in v5.
type TargetSnippetMetadataModel struct {
	MainModule types.String `tfsdk:"main_module"`
}

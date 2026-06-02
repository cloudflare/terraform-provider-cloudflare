package v500

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// SourceModel mirrors the prior (implicit schema_version=0) shape from main:
// id was Int64; the v500 schema requires String to align with the v7 SDK.
type SourceModel struct {
	ID                      types.Int64       `tfsdk:"id"`
	AccountID               types.String      `tfsdk:"account_id"`
	Email                   types.String      `tfsdk:"email"`
	IsEmailRegex            types.Bool        `tfsdk:"is_email_regex"`
	Name                    types.String      `tfsdk:"name"`
	Comments                types.String      `tfsdk:"comments"`
	CreatedAt               timetypes.RFC3339 `tfsdk:"created_at"`
	DirectoryID             types.Int64       `tfsdk:"directory_id"`
	DirectoryNodeID         types.Int64       `tfsdk:"directory_node_id"`
	ExternalDirectoryNodeID types.String      `tfsdk:"external_directory_node_id"`
	LastModified            timetypes.RFC3339 `tfsdk:"last_modified"`
	Provenance              types.String      `tfsdk:"provenance"`
}

// TargetModel mirrors the v500 schema.
type TargetModel struct {
	ID                      types.String      `tfsdk:"id"`
	AccountID               types.String      `tfsdk:"account_id"`
	Email                   types.String      `tfsdk:"email"`
	IsEmailRegex            types.Bool        `tfsdk:"is_email_regex"`
	Name                    types.String      `tfsdk:"name"`
	Comments                types.String      `tfsdk:"comments"`
	DirectoryID             types.Int64       `tfsdk:"directory_id"`
	DirectoryNodeID         types.Int64       `tfsdk:"directory_node_id"`
	ExternalDirectoryNodeID types.String      `tfsdk:"external_directory_node_id"`
	Provenance              types.String      `tfsdk:"provenance"`
	CreatedAt               timetypes.RFC3339 `tfsdk:"created_at"`
	LastModified            timetypes.RFC3339 `tfsdk:"last_modified"`
	ModifiedAt              timetypes.RFC3339 `tfsdk:"modified_at"`
}

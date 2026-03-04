package v500

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
)

// Transform converts source (legacy v4) snippet state to target (current v5) snippet state.
// This function is used by UpgradeFromV4 to handle the v4→v5 migration.
//
// Key transformations:
// - name → snippet_name (field rename)
// - main_module → metadata.main_module (restructured into nested object)
// - files blocks → files list attribute (blocks to list of objects)
// - created_on, modified_on: String → timetypes.RFC3339 (custom type conversion)
func Transform(ctx context.Context, source SourceCloudflareSnippetModel) (*TargetSnippetModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	// Validate required fields
	if source.ZoneID.IsNull() || source.ZoneID.IsUnknown() {
		diags.AddError(
			"Missing required field",
			"zone_id is required for snippet migration. The source state is missing this field.",
		)
		return nil, diags
	}
	if source.Name.IsNull() || source.Name.IsUnknown() {
		diags.AddError(
			"Missing required field",
			"name is required for snippet migration. The source state is missing this field.",
		)
		return nil, diags
	}

	target := &TargetSnippetModel{
		ZoneID: source.ZoneID,
		// Field rename: name → snippet_name
		SnippetName: source.Name,
	}

	// Timestamps: String → timetypes.RFC3339
	if !source.CreatedOn.IsNull() && !source.CreatedOn.IsUnknown() {
		target.CreatedOn = timetypes.NewRFC3339ValueMust(source.CreatedOn.ValueString())
	} else {
		target.CreatedOn = timetypes.NewRFC3339Null()
	}
	if !source.ModifiedOn.IsNull() && !source.ModifiedOn.IsUnknown() {
		target.ModifiedOn = timetypes.NewRFC3339ValueMust(source.ModifiedOn.ValueString())
	} else {
		target.ModifiedOn = timetypes.NewRFC3339Null()
	}

	// Restructure: main_module → metadata.main_module
	if !source.MainModule.IsNull() && !source.MainModule.IsUnknown() {
		target.Metadata = &TargetSnippetMetadataModel{
			MainModule: source.MainModule,
		}
	}

	// Convert files blocks to list of file models
	if len(source.Files) > 0 {
		files := make([]TargetSnippetFileModel, len(source.Files))
		for i, f := range source.Files {
			files[i] = TargetSnippetFileModel{
				Name:    f.Name,
				Content: f.Content,
			}
		}
		target.Files = &files
	}

	return target, diags
}

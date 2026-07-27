package v500

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Transform converts a source cloudflare_d1_database state (v4) to target cloudflare_d1_database state (v500).
func Transform(ctx context.Context, source SourceCloudflareD1DatabaseModel) (*TargetD1DatabaseModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	// Validate required fields
	if source.AccountID.IsNull() || source.AccountID.IsUnknown() {
		diags.AddError("Missing required field", "account_id is required for d1_database migration")
		return nil, diags
	}
	if source.Name.IsNull() || source.Name.IsUnknown() {
		diags.AddError("Missing required field", "name is required for d1_database migration")
		return nil, diags
	}

	target := &TargetD1DatabaseModel{
		// Direct copies
		ID:        source.ID,
		AccountID: source.AccountID,
		Name:      source.Name,
		Version:   source.Version,

		// Copy id to uuid -- v5 uses uuid for API calls (Read, Update, Delete)
		UUID: source.ID,

		// New optional fields: initialize as null
		Jurisdiction:        types.StringNull(),
		PrimaryLocationHint: types.StringNull(),

		// read_replication did not exist in v4. The D1 API returns
		// {mode:"disabled"} as the default on read but rejects null on
		// update. Initialize to the API default so the provider never
		// sends null in a PUT request.
		ReadReplication: &TargetD1DatabaseReadReplicationModel{
			Mode: types.StringValue("disabled"),
		},

		// New computed fields: initialize as null, API will populate on next plan/apply
		CreatedAt: timetypes.NewRFC3339Null(),
		FileSize:  types.Float64Null(),
		NumTables: types.Float64Null(),
	}

	return target, diags
}

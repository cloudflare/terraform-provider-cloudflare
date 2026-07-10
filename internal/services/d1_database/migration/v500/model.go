package v500

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// SourceCloudflareD1DatabaseModel represents the source cloudflare_d1_database state structure.
// This corresponds to schema_version=0 from the legacy (framework) cloudflare provider v4.
// Used by UpgradeFromLegacyV0 to parse legacy state.
type SourceCloudflareD1DatabaseModel struct {
	ID        types.String `tfsdk:"id"`
	AccountID types.String `tfsdk:"account_id"`
	Name      types.String `tfsdk:"name"`
	Version   types.String `tfsdk:"version"`
}

// TargetD1DatabaseModel represents the target cloudflare_d1_database state structure (v500).
// Must match the v5 D1DatabaseModel structure exactly.
type TargetD1DatabaseModel struct {
	ID                  types.String                       `tfsdk:"id"`
	UUID                types.String                       `tfsdk:"uuid"`
	AccountID           types.String                       `tfsdk:"account_id"`
	Name                types.String                       `tfsdk:"name"`
	Jurisdiction        types.String                       `tfsdk:"jurisdiction"`
	PrimaryLocationHint types.String                       `tfsdk:"primary_location_hint"`
	ReadReplication     *TargetD1DatabaseReadReplicationModel `tfsdk:"read_replication"`
	CreatedAt           timetypes.RFC3339                  `tfsdk:"created_at"`
	FileSize            types.Float64                      `tfsdk:"file_size"`
	NumTables           types.Float64                      `tfsdk:"num_tables"`
	Version             types.String                       `tfsdk:"version"`
}

// TargetD1DatabaseReadReplicationModel represents the read_replication nested object (v500).
// Must match D1DatabaseReadReplicationModel structure exactly.
type TargetD1DatabaseReadReplicationModel struct {
	Mode types.String `tfsdk:"mode"`
}

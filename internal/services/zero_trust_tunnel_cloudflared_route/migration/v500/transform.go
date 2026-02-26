package v500

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Transform converts source (legacy v4) state to target (current v5) state.
// This function is shared by both UpgradeFromV4 and MoveState handlers.
//
// Note on ID: The v4 provider stored the ID as a network CIDR (or an MD5
// checksum of "network/vnet_id" when a virtual network is set). The v5
// provider uses a UUID from the API. We carry the v4 ID forward as-is;
// Read() detects the non-UUID format, looks up the route via the List API
// (filtering by network + optional virtual_network_id), and replaces the
// legacy ID with the real UUID in state.
func Transform(ctx context.Context, source *SourceTunnelRouteModel) (*TargetTunnelRouteModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	// Copy all fields from v4 to v5
	target := &TargetTunnelRouteModel{
		ID:               source.ID,
		AccountID:        source.AccountID,
		TunnelID:         source.TunnelID,
		Network:          source.Network,
		VirtualNetworkID: source.VirtualNetworkID,
	}

	// comment: set to empty string if null (v5 has default "")
	if !source.Comment.IsNull() && !source.Comment.IsUnknown() {
		target.Comment = source.Comment
	} else {
		target.Comment = types.StringValue("")
	}

	// created_at: not in v4, set to null (API will repopulate on next refresh)
	target.CreatedAt = timetypes.NewRFC3339Null()

	// deleted_at: not in v4, set to null (null = not deleted)
	target.DeletedAt = timetypes.NewRFC3339Null()

	return target, diags
}

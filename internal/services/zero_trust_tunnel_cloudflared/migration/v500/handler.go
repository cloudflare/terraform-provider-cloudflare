package v500

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// UpgradeFromV4 upgrades state from schema version 0 (v4 SDKv2) to version 500.
// It renames secret→tunnel_secret, drops cname/tunnel_token, and initialises
// all new v5 computed fields to null.
func UpgradeFromV4(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	var source SourceTunnelCloudflaredModel
	resp.Diagnostics.Append(req.State.Get(ctx, &source)...)
	if resp.Diagnostics.HasError() {
		return
	}

	target, diags := Transform(ctx, &source)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, target)...)
}

// UpgradeFromV5 upgrades state from schema version 1 (v5.0–v5.23) to version 500.
// v5.24.0 removed the is_pending_reconnect field from the connections nested object.
// We parse old state using the previous model (which includes that field) and
// write new state using the current model (which omits it).
func UpgradeFromV5(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	// Parse old state — TargetTunnelCloudflaredModel includes is_pending_reconnect
	// so the framework can decode the v5.0–v5.23 raw state correctly.
	var old TargetTunnelCloudflaredModel
	resp.Diagnostics.Append(req.State.Get(ctx, &old)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Build new state without is_pending_reconnect.
	current := &CurrentTunnelCloudflaredModel{
		ID:              old.ID,
		AccountID:       old.AccountID,
		ConfigSrc:       old.ConfigSrc,
		Name:            old.Name,
		TunnelSecret:    old.TunnelSecret,
		AccountTag:      old.AccountTag,
		ConnsActiveAt:   old.ConnsActiveAt,
		ConnsInactiveAt: old.ConnsInactiveAt,
		CreatedAt:       old.CreatedAt,
		DeletedAt:       old.DeletedAt,
		RemoteConfig:    old.RemoteConfig,
		Status:          old.Status,
		TunType:         old.TunType,
		Metadata:        old.Metadata,
	}

	if old.Connections.IsNull() || old.Connections.IsUnknown() {
		current.Connections = customfield.NullObjectList[CurrentTunnelCloudflaredConnectionsModel](ctx)
	} else {
		oldConns, diags := old.Connections.AsStructSliceT(ctx)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}
		newConns := make([]CurrentTunnelCloudflaredConnectionsModel, len(oldConns))
		for i, c := range oldConns {
			newConns[i] = CurrentTunnelCloudflaredConnectionsModel{
				ID:            c.ID,
				ClientID:      c.ClientID,
				ClientVersion: c.ClientVersion,
				ColoName:      c.ColoName,
				OpenedAt:      c.OpenedAt,
				OriginIP:      c.OriginIP,
				UUID:          c.UUID,
			}
		}
		var diags2 diag.Diagnostics
		current.Connections, diags2 = customfield.NewObjectList(ctx, newConns)
		resp.Diagnostics.Append(diags2...)
		if resp.Diagnostics.HasError() {
			return
		}
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, current)...)
}

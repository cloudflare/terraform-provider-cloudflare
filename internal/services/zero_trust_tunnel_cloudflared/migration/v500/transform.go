package v500

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Transform converts a v4 SourceTunnelCloudflaredModel into the v5 target model.
//
// Field mapping:
//   - id          → id          (UUID, unchanged between v4 and v5)
//   - account_id  → account_id
//   - name        → name
//   - secret      → tunnel_secret  (renamed)
//   - config_src  → config_src     (defaults to "local" if absent)
//   - cname       → (dropped, not in v5)
//   - tunnel_token→ (dropped, not in v5)
//
// All new v5 computed fields are set to null so the provider's Read populates them.
func Transform(ctx context.Context, source *SourceTunnelCloudflaredModel) (*TargetTunnelCloudflaredModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	target := &TargetTunnelCloudflaredModel{
		ID:        source.ID,
		AccountID: source.AccountID,
		Name:      source.Name,
	}

	// Rename: secret → tunnel_secret
	target.TunnelSecret = source.Secret

	// Preserve config_src if set; default to "local" to prevent plan drift
	if !source.ConfigSrc.IsNull() && !source.ConfigSrc.IsUnknown() && source.ConfigSrc.ValueString() != "" {
		target.ConfigSrc = source.ConfigSrc
	} else {
		target.ConfigSrc = types.StringValue("local")
	}

	// New v5 computed fields — null until provider Read populates them
	target.AccountTag = types.StringNull()
	target.ConnsActiveAt = timetypes.NewRFC3339Null()
	target.ConnsInactiveAt = timetypes.NewRFC3339Null()
	target.CreatedAt = timetypes.NewRFC3339Null()
	target.DeletedAt = timetypes.NewRFC3339Null()
	target.RemoteConfig = types.BoolNull()
	target.Status = types.StringNull()
	target.TunType = types.StringNull()
	target.Connections = customfield.NullObjectList[TargetTunnelCloudflaredConnectionsModel](ctx)
	target.Metadata = jsontypes.NewNormalizedNull()

	return target, diags
}

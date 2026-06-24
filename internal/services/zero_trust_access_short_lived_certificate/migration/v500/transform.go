package v500

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/migrations"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Transform converts source (legacy v4) state to target (current v5) state.
// This function is shared by both UpgradeFromV0 and MoveState handlers.
func Transform(ctx context.Context, source SourceAccessCACertificateModel) (*TargetAccessShortLivedCertificateModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	tflog.Debug(ctx, "Transforming access_ca_certificate state from v4 to v5",
		map[string]interface{}{
			"source_application_id": source.ApplicationID.ValueString(),
		})

	target := &TargetAccessShortLivedCertificateModel{
		ID:    source.ID,
		AppID: source.ApplicationID, // Rename: application_id -> app_id
		// Convert empty-string Computed values to null (v4 had Optional+Computed, v5 has Optional only)
		AccountID: migrations.FalseyStringToNull(source.AccountID),
		ZoneID:    migrations.FalseyStringToNull(source.ZoneID),
		AUD:       source.AUD,
		PublicKey: source.PublicKey,
	}

	return target, diags
}

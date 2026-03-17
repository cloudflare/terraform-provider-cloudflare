package v500

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func Transform(_ context.Context, source SourceLeakedCredentialCheckModel) (*TargetLeakedCredentialCheckModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	if source.ZoneID.IsNull() || source.ZoneID.IsUnknown() {
		diags.AddError(
			"Missing required field",
			"zone_id is required for leaked_credential_check migration.",
		)
		return nil, diags
	}

	target := &TargetLeakedCredentialCheckModel{
		ZoneID:  source.ZoneID,
		Enabled: source.Enabled,
	}

	return target, diags
}

package v500

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func Transform(_ context.Context, source SourceLeakedCredentialCheckRuleModel) (*TargetLeakedCredentialCheckRuleModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	if source.ZoneID.IsNull() || source.ZoneID.IsUnknown() {
		diags.AddError(
			"Missing required field",
			"zone_id is required for leaked_credential_check_rule migration.",
		)
		return nil, diags
	}

	if source.Username.IsNull() || source.Username.IsUnknown() {
		diags.AddError(
			"Missing required field",
			"username is required for leaked_credential_check_rule migration from v4 state.",
		)
		return nil, diags
	}

	if source.Password.IsNull() || source.Password.IsUnknown() {
		diags.AddError(
			"Missing required field",
			"password is required for leaked_credential_check_rule migration from v4 state.",
		)
		return nil, diags
	}

	target := &TargetLeakedCredentialCheckRuleModel{
		ID:       source.ID,
		ZoneID:   source.ZoneID,
		Username: source.Username,
		Password: source.Password,
	}

	return target, diags
}

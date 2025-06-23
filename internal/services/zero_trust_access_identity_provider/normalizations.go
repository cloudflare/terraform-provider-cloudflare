package zero_trust_access_identity_provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
)

func normalizeReadZeroTrustIDPData(ctx context.Context, apiValue *ZeroTrustAccessIdentityProviderModel, state *tfsdk.State) diag.Diagnostics {
	var stateValue ZeroTrustAccessIdentityProviderModel
	d := state.Get(ctx, &stateValue)
	if apiValue.Type.ValueString() == "azureAD" &&
		(apiValue.Config != nil && !apiValue.Config.ConditionalAccessEnabled.ValueBool()) &&
		(stateValue.Config != nil && !stateValue.Config.ConditionalAccessEnabled.ValueBool()) {
		apiValue.Config.ConditionalAccessEnabled = stateValue.Config.ConditionalAccessEnabled
	}
	return d
}

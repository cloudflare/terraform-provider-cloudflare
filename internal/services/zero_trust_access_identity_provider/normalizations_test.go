package zero_trust_access_identity_provider

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestNormalizeReadZeroTrustIDPConfigLoginPageAutoPrompt(t *testing.T) {
	t.Parallel()

	for _, tc := range []struct {
		name  string
		api   types.Bool
		state types.Bool
		want  types.Bool
	}{
		{"omitted_preserves_false", types.BoolNull(), types.BoolValue(false), types.BoolValue(false)},
		{"false_preserves_null", types.BoolValue(false), types.BoolNull(), types.BoolNull()},
		{"null_stays_null", types.BoolNull(), types.BoolNull(), types.BoolNull()},
		{"false_stays_false", types.BoolValue(false), types.BoolValue(false), types.BoolValue(false)},
		{"enable_drift_is_visible", types.BoolValue(true), types.BoolValue(false), types.BoolValue(true)},
		{"disable_drift_is_visible", types.BoolValue(false), types.BoolValue(true), types.BoolValue(false)},
		{"omitted_does_not_preserve_true", types.BoolNull(), types.BoolValue(true), types.BoolNull()},
		{"true_stays_true", types.BoolValue(true), types.BoolValue(true), types.BoolValue(true)},
	} {
		t.Run(tc.name, func(t *testing.T) {
			data := ZeroTrustAccessIdentityProviderModel{Config: &ZeroTrustAccessIdentityProviderConfigModel{
				LoginPageAutoPrompt: tc.api,
			}}
			state := ZeroTrustAccessIdentityProviderModel{Config: &ZeroTrustAccessIdentityProviderConfigModel{
				LoginPageAutoPrompt: tc.state,
			}}
			if diags := normalizeReadZeroTrustIDPConfigData(context.Background(), &data, &state); diags.HasError() {
				t.Fatalf("normalize config: %v", diags)
			}
			if !data.Config.LoginPageAutoPrompt.Equal(tc.want) {
				t.Errorf("login_page_auto_prompt = %v, want %v", data.Config.LoginPageAutoPrompt, tc.want)
			}
		})
	}
}

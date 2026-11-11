package zero_trust_device_default_profile_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_device_default_profile"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// policy_id never changes for the default profile, so plans must keep the
// value from state instead of showing it as "(known after apply)" (BUGS-2129).
func TestZeroTrustDeviceDefaultProfilePolicyIDKeepsState(t *testing.T) {
	t.Parallel()
	attr, ok := zero_trust_device_default_profile.ResourceSchema(context.TODO()).Attributes["policy_id"].(schema.StringAttribute)
	if !ok {
		t.Fatal("policy_id is not a StringAttribute")
	}

	plan := func(state types.String) types.String {
		resp := &planmodifier.StringResponse{PlanValue: types.StringUnknown()}
		for _, m := range attr.PlanModifiers {
			m.PlanModifyString(context.TODO(), planmodifier.StringRequest{
				StateValue: state,
				PlanValue:  resp.PlanValue,
			}, resp)
		}
		return resp.PlanValue
	}

	if got := plan(types.StringValue("0197d1b3-fb51-71e6-ba89-d9459caad255")); got.ValueString() != "0197d1b3-fb51-71e6-ba89-d9459caad255" {
		t.Fatalf("update plan: policy_id = %s, want the state value", got)
	}
	if got := plan(types.StringNull()); !got.IsUnknown() {
		t.Fatalf("create plan: policy_id = %s, want unknown", got)
	}
}

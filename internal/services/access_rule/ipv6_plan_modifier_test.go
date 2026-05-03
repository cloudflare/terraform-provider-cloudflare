package access_rule

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestIPv6CanonicalPlanModifier(t *testing.T) {
	type tc struct {
		state      string
		plan       string
		expectKept bool // true if planmodifier should overwrite plan with state
	}

	cases := map[string]tc{
		"compressed_vs_long_cidr": {
			state:      "2001:0db8:0000:0000:0000:0000:0000:0000/32",
			plan:       "2001:db8::/32",
			expectKept: true,
		},
		"compressed_vs_long_address": {
			state:      "2001:0db8:0000:0000:0000:0000:0000:0001",
			plan:       "2001:db8::1",
			expectKept: true,
		},
		"different_networks": {
			state:      "2001:db8::/32",
			plan:       "2001:db9::/32",
			expectKept: false,
		},
		"different_prefix_length": {
			state:      "2001:db8::/32",
			plan:       "2001:db8::/48",
			expectKept: false,
		},
		"identical_cidr": {
			state:      "2001:db8::/32",
			plan:       "2001:db8::/32",
			expectKept: false, // short-circuited, plan unchanged
		},
		"ipv4_unaffected": {
			state:      "192.0.2.1/24",
			plan:       "192.0.2.0/24",
			expectKept: false,
		},
		"non_ip_value_unaffected": {
			state:      "US",
			plan:       "DE",
			expectKept: false,
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			req := planmodifier.StringRequest{
				Path:       path.Root("test"),
				StateValue: types.StringValue(c.state),
				PlanValue:  types.StringValue(c.plan),
			}
			resp := &planmodifier.StringResponse{PlanValue: types.StringValue(c.plan)}

			ipv6CanonicalValue().PlanModifyString(context.Background(), req, resp)

			got := resp.PlanValue.ValueString()
			if c.expectKept && got != c.state {
				t.Errorf("expected plan to be overwritten with state %q, got %q", c.state, got)
			}
			if !c.expectKept && got != c.plan {
				t.Errorf("expected plan to remain %q, got %q", c.plan, got)
			}
		})
	}
}

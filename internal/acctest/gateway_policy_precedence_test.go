package acctest

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
)

func TestGatewayPolicyPrecedencePlanCheck(t *testing.T) {
	const address = "cloudflare_zero_trust_gateway_policy.allow"
	for _, tc := range []struct {
		name     string
		before   string
		after    string
		expected map[string]GatewayPolicyPrecedenceChange
		wantErr  bool
	}{
		{name: "unchanged", before: "10000", after: "10000"},
		{name: "unreviewed_reorder", before: "10000", after: "10", wantErr: true},
		{name: "unreviewed_hash_removal", before: "10116667", after: "10116", wantErr: true},
		{
			name: "explicit_change", before: "10116667", after: "10116",
			expected: map[string]GatewayPolicyPrecedenceChange{address: {Before: 10116667, After: 10116}},
		},
		{
			name: "wrong_source", before: "10116781", after: "10116", wantErr: true,
			expected: map[string]GatewayPolicyPrecedenceChange{address: {Before: 10116667, After: 10116}},
		},
		{
			name: "wrong_target", before: "10116667", after: "10", wantErr: true,
			expected: map[string]GatewayPolicyPrecedenceChange{address: {Before: 10116667, After: 10116}},
		},
		{
			name: "wrong_resource", before: "10116667", after: "10116", wantErr: true,
			expected: map[string]GatewayPolicyPrecedenceChange{"cloudflare_zero_trust_gateway_policy.other": {Before: 10116667, After: 10116}},
		},
		{name: "unknown_target", before: "10000", after: "null", wantErr: true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			var plan tfjson.Plan
			err := json.Unmarshal([]byte(fmt.Sprintf(`{
				"format_version": "1.0",
				"resource_changes": [{
					"address": %q,
					"mode": "managed",
					"type": "cloudflare_zero_trust_gateway_policy",
					"name": "allow",
					"change": {
						"actions": ["update"],
						"before": {"precedence": %s},
						"after": {"precedence": %s}
					}
				}]
			}`, address, tc.before, tc.after)), &plan)
			if err != nil {
				t.Fatal(err)
			}
			// The acceptance harness uses json.Number; plain JSON decoding uses
			// float64. Exercise both representations of the same plan.
			encoded, err := json.Marshal(&plan)
			if err != nil {
				t.Fatal(err)
			}
			for _, useJSONNumber := range []bool{false, true} {
				var checkedPlan tfjson.Plan
				checkedPlan.UseJSONNumber(useJSONNumber)
				if err := json.NewDecoder(strings.NewReader(string(encoded))).Decode(&checkedPlan); err != nil {
					t.Fatal(err)
				}
				var resp plancheck.CheckPlanResponse
				gatewayPolicyPrecedencePlanCheck{expected: tc.expected}.CheckPlan(context.Background(), plancheck.CheckPlanRequest{Plan: &checkedPlan}, &resp)
				if (resp.Error != nil) != tc.wantErr {
					t.Fatalf("UseJSONNumber=%t: CheckPlan error = %v, want error = %t", useJSONNumber, resp.Error, tc.wantErr)
				}
			}
		})
	}
}

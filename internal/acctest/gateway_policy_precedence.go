package acctest

import (
	"context"
	"encoding/json"
	"fmt"
	"hash/crc32"
	"reflect"

	"github.com/hashicorp/terraform-plugin-testing/plancheck"
)

// GatewayPolicyPrecedenceChange explicitly describes a reviewed migration update.
type GatewayPolicyPrecedenceChange struct {
	Before int64
	After  int64
}

// GatewayPolicyV4PrecedenceChange describes the exact change from the v4
// provider's name-based API encoding to an explicitly configured v5 precedence.
// This is a test expectation, not a heuristic for decoding arbitrary state.
func GatewayPolicyV4PrecedenceChange(ruleName string, configured int64) GatewayPolicyPrecedenceChange {
	return GatewayPolicyPrecedenceChange{
		Before: configured*1000 + int64(crc32.ChecksumIEEE([]byte(ruleName)))%1000,
		After:  configured,
	}
}

type gatewayPolicyPrecedencePlanCheck struct {
	expected map[string]GatewayPolicyPrecedenceChange
}

func (c gatewayPolicyPrecedencePlanCheck) CheckPlan(_ context.Context, req plancheck.CheckPlanRequest, resp *plancheck.CheckPlanResponse) {
	for _, rc := range req.Plan.ResourceChanges {
		if rc.Type != "cloudflare_zero_trust_gateway_policy" {
			continue
		}
		before, beforeOK := rc.Change.Before.(map[string]interface{})
		after, afterOK := rc.Change.After.(map[string]interface{})
		if !beforeOK || !afterOK {
			continue // The general migration checker handles creation/deletion.
		}
		if reflect.DeepEqual(before["precedence"], after["precedence"]) {
			continue
		}
		expected, ok := c.expected[rc.Address]
		if ok && gatewayPolicyPrecedenceMatches(before["precedence"], expected.Before) && gatewayPolicyPrecedenceMatches(after["precedence"], expected.After) {
			continue
		}
		resp.Error = fmt.Errorf("unexpected Gateway policy precedence change for %s: %v -> %v", rc.Address, before["precedence"], after["precedence"])
		return
	}
}

func gatewayPolicyPrecedenceMatches(value interface{}, expected int64) bool {
	switch number := value.(type) {
	case json.Number:
		actual, err := number.Int64()
		return err == nil && actual == expected
	case float64:
		return number == float64(expected)
	default:
		return false
	}
}

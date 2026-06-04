package v500_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/load_balancer_pool"
	v500 "github.com/cloudflare/terraform-provider-cloudflare/internal/services/load_balancer_pool/migration/v500"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/stretchr/testify/require"
)

// TestUpgradeFromV0Ambiguous_V5ObjectShape is a regression test for #7098.
//
// Before the fix, the slot-0 upgrader was registered with PriorSchema set to
// the v4 (SDKv2) source schema, which expects `load_shedding` and
// `origin_steering` as list blocks. When the user's state was written by
// v5.0–v5.7 (object shape) at schema_version=0, the Plugin Framework rejected
// the state pre-handler with:
//
//	AttributeName("origin_steering"): invalid JSON, expected "[", got "{"
//
// With the fix, PriorSchema is nil and the handler probes the raw JSON itself
// against the target schema. Object-shaped state should pass through cleanly.
func TestUpgradeFromV0Ambiguous_V5ObjectShape(t *testing.T) {
	ctx := context.Background()

	// Minimal v5-shaped state (post v5.0, schema_version=0). `load_shedding`
	// and `origin_steering` are SINGLE objects — this is the exact shape that
	// broke for the user in #7098.
	stateJSON := map[string]interface{}{
		"id":         "pool-1",
		"account_id": "acct-123",
		"name":       "test-pool",
		"enabled":    true,
		"origins": []interface{}{
			map[string]interface{}{
				"name":    "origin-1",
				"address": "1.2.3.4",
			},
		},
		"origin_steering": map[string]interface{}{
			"policy": "random",
		},
		"load_shedding": map[string]interface{}{
			"default_percent": 0,
			"default_policy":  "random",
			"session_percent": 0,
			"session_policy":  "hash",
		},
	}
	raw, err := json.Marshal(stateJSON)
	require.NoError(t, err)

	req := resource.UpgradeStateRequest{
		// PriorSchema=nil in production: req.State is empty.
		RawState: &tfprotov6.RawState{JSON: raw},
	}
	resp := &resource.UpgradeStateResponse{
		State: tfsdk.State{
			Schema: load_balancer_pool.ResourceSchema(ctx),
		},
	}

	v500.UpgradeFromV0Ambiguous(ctx, req, resp)

	require.False(t, resp.Diagnostics.HasError(),
		"UpgradeFromV0Ambiguous should not error on v5 object-shape state, got: %v",
		resp.Diagnostics)
}

// TestUpgradeFromV0Ambiguous_NilRawState verifies that a nil raw state surfaces
// a clear diagnostic rather than panicking.
func TestUpgradeFromV0Ambiguous_NilRawState(t *testing.T) {
	ctx := context.Background()

	req := resource.UpgradeStateRequest{
		RawState: nil,
	}
	resp := &resource.UpgradeStateResponse{
		State: tfsdk.State{
			Schema: load_balancer_pool.ResourceSchema(ctx),
		},
	}

	v500.UpgradeFromV0Ambiguous(ctx, req, resp)

	require.True(t, resp.Diagnostics.HasError(),
		"Expected error diagnostic for nil RawState")
}

// This file is intentionally NOT code-generated.

package v500_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/load_balancer"
	v500 "github.com/cloudflare/terraform-provider-cloudflare/internal/services/load_balancer/migration/v500"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

const v4LoadBalancerWithOverrideAffinityState = `{
	"id": "699d98642c564d2e855e9661899b7252",
	"zone_id": "023e105f4ecef8ad9ca31a8372d0c353",
	"name": "lb.example.com",
	"default_pool_ids": ["17b5962d775c646f3f9725cbc7a53df4"],
	"fallback_pool_id": "17b5962d775c646f3f9725cbc7a53df4",
	"session_affinity": "cookie",
	"session_affinity_attributes": [{
		"samesite": "Auto",
		"secure": "Auto",
		"drain_duration": 60,
		"zero_downtime_failover": "none",
		"headers": [],
		"require_all_headers": false
	}],
	"rules": [{
		"name": "header rule",
		"priority": 0,
		"disabled": false,
		"condition": "http.request.uri.path contains \"/api\"",
		"terminates": false,
		"overrides": [{
			"session_affinity": "header",
			"session_affinity_attributes": [{
				"samesite": "Strict",
				"secure": "Always",
				"zero_downtime_failover": "sticky",
				"headers": ["x-session"],
				"require_all_headers": true
			}]
		}],
		"fixed_response": []
	}]
}`

// The v4 rules.overrides.session_affinity_attributes block has no
// drain_duration, so the source model for it must not either, or every v4
// load balancer with an override affinity block fails to upgrade.
func TestLoadBalancerUpgradeFromV4OverrideSessionAffinitySetsState(t *testing.T) {
	ctx := context.Background()
	resp := resource.UpgradeStateResponse{State: tfsdk.State{Schema: load_balancer.ResourceSchema(ctx)}}
	v500.UpgradeFromV1Ambiguous(ctx, resource.UpgradeStateRequest{
		RawState: &tfprotov6.RawState{JSON: []byte(v4LoadBalancerWithOverrideAffinityState)},
	}, &resp)
	if resp.Diagnostics.HasError() {
		t.Fatalf("UpgradeFromV1Ambiguous failed: %v", resp.Diagnostics)
	}

	var got load_balancer.LoadBalancerModel
	if diags := resp.State.Get(ctx, &got); diags.HasError() {
		t.Fatalf("could not decode upgraded state: %v", diags)
	}

	top, diags := got.SessionAffinityAttributes.Value(ctx)
	if diags.HasError() || top == nil {
		t.Fatalf("session_affinity_attributes missing after upgrade: %v", diags)
	}
	if !top.DrainDuration.Equal(types.Float64Value(60)) {
		t.Errorf("session_affinity_attributes.drain_duration = %s, want 60", top.DrainDuration)
	}

	rules, diags := got.Rules.AsStructSliceT(ctx)
	if diags.HasError() || len(rules) != 1 {
		t.Fatalf("rules = %+v, want one rule: %v", rules, diags)
	}
	overrides, diags := rules[0].Overrides.Value(ctx)
	if diags.HasError() || overrides == nil {
		t.Fatalf("rules[0].overrides missing after upgrade: %v", diags)
	}
	attrs, diags := overrides.SessionAffinityAttributes.Value(ctx)
	if diags.HasError() || attrs == nil {
		t.Fatalf("rules[0].overrides.session_affinity_attributes missing after upgrade: %v", diags)
	}
	if !attrs.Samesite.Equal(types.StringValue("Strict")) || !attrs.ZeroDowntimeFailover.Equal(types.StringValue("sticky")) {
		t.Errorf("override session_affinity_attributes = %+v", attrs)
	}
	if !attrs.DrainDuration.IsNull() {
		t.Errorf("override session_affinity_attributes.drain_duration = %s, want null", attrs.DrainDuration)
	}
}

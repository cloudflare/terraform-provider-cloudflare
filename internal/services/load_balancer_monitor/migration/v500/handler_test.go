package v500_test

import (
	"context"
	"strings"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/load_balancer_monitor"
	v500 "github.com/cloudflare/terraform-provider-cloudflare/internal/services/load_balancer_monitor/migration/v500"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

func TestUpgradeFromV0DisambiguatesHeaderShapes(t *testing.T) {
	ctx := context.Background()
	upgrade := v500.UpgradeFromV0(load_balancer_monitor.ResourceSchema(ctx))
	response := func() resource.UpgradeStateResponse {
		return resource.UpgradeStateResponse{
			State: tfsdk.State{Schema: load_balancer_monitor.ResourceSchema(ctx)},
		}
	}

	for _, test := range []struct {
		name             string
		json             string
		expectHostHeader bool
		expectV4Defaults bool
		expectV5NoOp     bool
	}{
		{
			name:             "early v5 object",
			json:             `{"account_id":"account","type":"http","header":{"Host":["example.com"]}}`,
			expectHostHeader: true,
			expectV5NoOp:     true,
		},
		{
			name:             "v4 collection",
			json:             `{"account_id":"account","type":"http","expected_codes":null,"header":[{"header":"Host","values":["example.com"]}]}`,
			expectHostHeader: true,
			expectV4Defaults: true,
		},
		{
			name:             "v4 header omitted",
			json:             `{"account_id":"account","type":"http","expected_codes":null}`,
			expectV4Defaults: true,
		},
		{
			name:             "v4 header null",
			json:             `{"account_id":"account","type":"http","expected_codes":null,"header":null}`,
			expectV4Defaults: true,
		},
		{
			name:             "v4 header empty",
			json:             `{"account_id":"account","type":"http","expected_codes":null,"header":[]}`,
			expectV4Defaults: true,
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			resp := response()
			upgrade(ctx, resource.UpgradeStateRequest{
				RawState: &tfprotov6.RawState{JSON: []byte(test.json)},
			}, &resp)
			if resp.Diagnostics.HasError() {
				t.Fatalf("unexpected diagnostics: %s", resp.Diagnostics)
			}
			var state v500.TargetLoadBalancerMonitorModel
			diags := (tfsdk.State{Raw: resp.State.Raw, Schema: load_balancer_monitor.ResourceSchema(ctx)}).Get(ctx, &state)
			if diags.HasError() {
				t.Fatalf("could not decode target state: %s", diags)
			}
			if test.expectHostHeader {
				if state.Header == nil {
					t.Fatalf("header was not preserved in target map: %#v", state.Header)
				}
				hostValues := (*state.Header)["Host"]
				if hostValues == nil || len(*hostValues) != 1 || (*hostValues)[0].ValueString() != "example.com" {
					t.Fatalf("header was not preserved in target map: %#v", state.Header)
				}
			} else if state.Header != nil {
				t.Fatalf("expected no target header, got %#v", state.Header)
			}
			if test.expectV4Defaults && (state.ExpectedCodes.IsNull() || state.ExpectedCodes.ValueString() != "") {
				t.Fatalf("expected v4 defaults to be applied, got expected_codes=%s", state.ExpectedCodes)
			}
			if test.expectV5NoOp && !state.ExpectedCodes.IsNull() {
				t.Fatalf("expected early-v5 state to bypass v4 defaults, got expected_codes=%s", state.ExpectedCodes)
			}
		})
	}
}

func TestUpgradeFromV0RejectsMissingAndUnknownState(t *testing.T) {
	ctx := context.Background()
	upgrade := v500.UpgradeFromV0(load_balancer_monitor.ResourceSchema(ctx))
	response := func() resource.UpgradeStateResponse {
		return resource.UpgradeStateResponse{
			State: tfsdk.State{Schema: load_balancer_monitor.ResourceSchema(ctx)},
		}
	}

	for _, test := range []struct {
		name string
		raw  *tfprotov6.RawState
	}{
		{name: "missing", raw: nil},
		{name: "empty", raw: &tfprotov6.RawState{}},
		{name: "neither shape", raw: &tfprotov6.RawState{JSON: []byte(`{"account_id":"account","header":true}`)}},
	} {
		t.Run(test.name, func(t *testing.T) {
			resp := response()
			upgrade(ctx, resource.UpgradeStateRequest{RawState: test.raw}, &resp)
			if !resp.Diagnostics.HasError() || resp.Diagnostics[0].Summary() == "" || resp.Diagnostics[0].Detail() == "" {
				t.Fatal("expected a useful migration diagnostic")
			}
			if test.name == "neither shape" && (!strings.Contains(resp.Diagnostics[0].Summary(), "Unrecognized") || !strings.Contains(resp.Diagnostics[0].Detail(), "early-v5") || !strings.Contains(resp.Diagnostics[0].Detail(), "v4")) {
				t.Fatalf("diagnostic does not identify both attempted shapes: %s", resp.Diagnostics)
			}
		})
	}
}

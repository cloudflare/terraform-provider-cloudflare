package flagship_flag

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

func TestUpgradeFlagshipFlagV500ToV501FrameworkDecode(t *testing.T) {
	ctx := context.Background()
	priorSchema := resourceSchemaV500(ctx)
	rawState := &tfprotov6.RawState{JSON: []byte(`{
		"account_id":"acct",
		"app_id":"app",
		"flag_key":null,
		"default_variation":"disabled",
		"enabled":true,
		"key":"example-flag",
		"variations":{"enabled":"true","disabled":"false"},
		"rules":[],
		"description":"",
		"type":"boolean",
		"updated_at":"2026-08-13T00:00:00Z",
		"updated_by":"test@example.com"
	}`)}
	rawValue, err := rawState.Unmarshal(priorSchema.Type().TerraformType(ctx))
	if err != nil {
		t.Fatalf("decode prior raw state: %v", err)
	}
	priorState := tfsdk.State{Raw: rawValue, Schema: priorSchema}

	r := &FlagshipFlagResource{}
	upgrader, ok := r.UpgradeState(ctx)[500]
	if !ok {
		t.Fatal("no upgrader registered for version 500")
	}
	resp := resource.UpgradeStateResponse{State: tfsdk.State{Schema: ResourceSchema(ctx)}}
	upgrader.StateUpgrader(ctx, resource.UpgradeStateRequest{RawState: rawState, State: &priorState}, &resp)
	if resp.Diagnostics.HasError() {
		t.Fatalf("upgrade failed: %v", resp.Diagnostics)
	}
	if len(resp.Diagnostics) != 1 {
		t.Fatalf("diagnostics = %v, want one migration warning", resp.Diagnostics)
	}
	if got := resp.Diagnostics[0].Summary(); got != "Flagship flag state upgraded" {
		t.Fatalf("warning summary = %q", got)
	}

	var got FlagshipFlagModel
	if diags := resp.State.Get(ctx, &got); diags.HasError() {
		t.Fatalf("decode upgraded state: %v", diags)
	}
	if got.Key.ValueString() != "example-flag" {
		t.Fatalf("key = %q", got.Key.ValueString())
	}
	if _, exists := ResourceSchema(ctx).Attributes["flag_key"]; exists {
		t.Fatal("flag_key remains in the v501 schema")
	}
}

func TestUpgradeFlagshipFlagV500FallsBackToLegacyFlagKey(t *testing.T) {
	upgraded := upgradeFlagshipFlagV500ToV501(flagshipFlagModelV500{
		Key:     types.StringNull(),
		FlagKey: types.StringValue("legacy-key"),
	})
	if upgraded.Key.ValueString() != "legacy-key" {
		t.Fatalf("key = %q, want legacy-key", upgraded.Key.ValueString())
	}
}

func TestUpgradeFlagshipFlagV500PrefersCanonicalKey(t *testing.T) {
	upgraded := upgradeFlagshipFlagV500ToV501(flagshipFlagModelV500{
		Key:     types.StringValue("example-flag"),
		FlagKey: types.StringValue("wrong-flag"),
	})
	if upgraded.Key.ValueString() != "example-flag" {
		t.Fatalf("key = %q, want example-flag", upgraded.Key.ValueString())
	}
}

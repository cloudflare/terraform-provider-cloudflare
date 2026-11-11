// This file is intentionally NOT code-generated.

package v500_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/workers_script"
	v500 "github.com/cloudflare/terraform-provider-cloudflare/internal/services/workers_script/migration/v500"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

const v4WorkerScriptState = `{
	"id": "my-worker",
	"name": "my-worker",
	"account_id": "f037e56e89293a057740de681ac9abbe",
	"content": "export default { fetch() { return new Response('ok') } }",
	"module": true,
	"plain_text_binding": [{"name": "GREETING", "text": "hello"}]
}`

// The target model must convert into the live schema, which gained
// code_update_strategy, or both UpgradeState and the moved block fail with a
// Value Conversion Error.
func TestWorkersScriptMigrationFromV4SetsState(t *testing.T) {
	ctx := context.Background()
	targetSchema := workers_script.ResourceSchema(ctx)
	upgrader := workers_script.NewResource().(resource.ResourceWithUpgradeState).UpgradeState(ctx)[0]

	for _, tc := range []struct {
		name    string
		migrate func(t *testing.T) tfsdk.State
	}{
		{
			name: "UpgradeFromV0",
			migrate: func(t *testing.T) tfsdk.State {
				sourceState := decodeState(t, ctx, *upgrader.PriorSchema, v4WorkerScriptState)
				resp := resource.UpgradeStateResponse{State: tfsdk.State{Schema: targetSchema}}
				upgrader.StateUpgrader(ctx, resource.UpgradeStateRequest{State: &sourceState}, &resp)
				if resp.Diagnostics.HasError() {
					t.Fatalf("UpgradeFromV0 failed: %v", resp.Diagnostics)
				}
				return resp.State
			},
		},
		{
			name: "MoveState",
			migrate: func(t *testing.T) tfsdk.State {
				sourceState := decodeState(t, ctx, v500.SourceWorkerScriptSchema(), v4WorkerScriptState)
				resp := resource.MoveStateResponse{TargetState: tfsdk.State{Schema: targetSchema}}
				v500.MoveState(ctx, resource.MoveStateRequest{
					SourceTypeName:        "cloudflare_worker_script",
					SourceProviderAddress: "registry.terraform.io/cloudflare/cloudflare",
					SourceState:           &sourceState,
				}, &resp)
				if resp.Diagnostics.HasError() {
					t.Fatalf("MoveState failed: %v", resp.Diagnostics)
				}
				return resp.TargetState
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			state := tc.migrate(t)
			var got workers_script.WorkersScriptModel
			if diags := state.Get(ctx, &got); diags.HasError() {
				t.Fatalf("could not decode migrated state: %v", diags)
			}
			if !got.ScriptName.Equal(types.StringValue("my-worker")) {
				t.Errorf("script_name = %s, want \"my-worker\"", got.ScriptName)
			}
			bindings, diags := got.Bindings.AsStructSliceT(ctx)
			if diags.HasError() || len(bindings) != 1 || !bindings[0].Name.Equal(types.StringValue("GREETING")) {
				t.Errorf("bindings = %+v, want one GREETING binding", bindings)
			}
			if got.CodeUpdateStrategy != nil {
				t.Errorf("code_update_strategy = %+v, want null", got.CodeUpdateStrategy)
			}
		})
	}
}

func decodeState(t *testing.T, ctx context.Context, s schema.Schema, jsonState string) tfsdk.State {
	t.Helper()
	raw, err := (&tfprotov6.RawState{JSON: []byte(jsonState)}).Unmarshal(s.Type().TerraformType(ctx))
	if err != nil {
		t.Fatalf("could not decode v4 state: %v", err)
	}
	return tfsdk.State{Raw: raw, Schema: s}
}

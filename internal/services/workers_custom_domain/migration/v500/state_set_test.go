// This file is intentionally NOT code-generated.

package v500_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/workers_custom_domain"
	v500 "github.com/cloudflare/terraform-provider-cloudflare/internal/services/workers_custom_domain/migration/v500"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

const v4WorkerDomainState = `{
	"id": "11111111-1111-4111-8111-111111111111",
	"account_id": "f037e56e89293a057740de681ac9abbe",
	"hostname": "api.example.com",
	"service": "my-worker",
	"zone_id": "023e105f4ecef8ad9ca31a8372d0c353",
	"environment": "production",
	"zone_name": "example.com",
	"cert_id": "22222222-2222-4222-8222-222222222222"
}`

// The target model must convert into the live schema, which gained enabled and
// previews_enabled, or both UpgradeState and the moved block fail with a Value
// Conversion Error.
func TestWorkersCustomDomainMigrationFromV4SetsState(t *testing.T) {
	ctx := context.Background()
	sourceSchema := v500.SourceSchemaV0()
	targetSchema := workers_custom_domain.ResourceSchema(ctx)
	raw, err := (&tfprotov6.RawState{JSON: []byte(v4WorkerDomainState)}).Unmarshal(sourceSchema.Type().TerraformType(ctx))
	if err != nil {
		t.Fatalf("could not decode v4 state: %v", err)
	}
	sourceState := tfsdk.State{Raw: raw, Schema: sourceSchema}

	for _, tc := range []struct {
		name    string
		migrate func(t *testing.T) tfsdk.State
	}{
		{
			name: "UpgradeStateV0toV500",
			migrate: func(t *testing.T) tfsdk.State {
				resp := resource.UpgradeStateResponse{State: tfsdk.State{Schema: targetSchema}}
				v500.UpgradeStateV0toV500(sourceSchema).StateUpgrader(ctx, resource.UpgradeStateRequest{State: &sourceState}, &resp)
				if resp.Diagnostics.HasError() {
					t.Fatalf("UpgradeStateV0toV500 failed: %v", resp.Diagnostics)
				}
				return resp.State
			},
		},
		{
			name: "MoveStateV4toV500",
			migrate: func(t *testing.T) tfsdk.State {
				resp := resource.MoveStateResponse{TargetState: tfsdk.State{Schema: targetSchema}}
				v500.MoveStateV4toV500(ctx, resource.MoveStateRequest{
					SourceTypeName:        "cloudflare_worker_domain",
					SourceProviderAddress: "registry.terraform.io/cloudflare/cloudflare",
					SourceState:           &sourceState,
				}, &resp)
				if resp.Diagnostics.HasError() {
					t.Fatalf("MoveStateV4toV500 failed: %v", resp.Diagnostics)
				}
				return resp.TargetState
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			state := tc.migrate(t)
			var got workers_custom_domain.WorkersCustomDomainModel
			if diags := state.Get(ctx, &got); diags.HasError() {
				t.Fatalf("could not decode migrated state: %v", diags)
			}
			if !got.Hostname.Equal(types.StringValue("api.example.com")) || !got.Environment.Equal(types.StringValue("production")) {
				t.Errorf("hostname = %s, environment = %s", got.Hostname, got.Environment)
			}
			if !got.Enabled.IsNull() || !got.PreviewsEnabled.IsNull() {
				t.Errorf("enabled = %s, previews_enabled = %s, want null", got.Enabled, got.PreviewsEnabled)
			}
		})
	}
}

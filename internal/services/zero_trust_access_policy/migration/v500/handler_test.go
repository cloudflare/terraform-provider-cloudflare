// This file is intentionally NOT code-generated.

package v500_test

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-go/tftypes"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_access_policy"
	v500 "github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_access_policy/migration/v500"
)

// Early v5 releases (v5.0-v5.6) stored the computed app_count, which the
// resource no longer has. Their schema_version=0 state must still upgrade.
func TestUpgradeFromSchemaV0_DropsRemovedAppCount(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	schema := zero_trust_access_policy.ResourceSchema(ctx)
	req := resource.UpgradeStateRequest{
		RawState: &tfprotov6.RawState{JSON: []byte(`{
			"id": "policy-id",
			"account_id": "account-id",
			"name": "test",
			"decision": "allow",
			"app_count": 2
		}`)},
	}
	resp := resource.UpgradeStateResponse{State: tfsdk.State{Schema: schema}}

	v500.UpgradeFromSchemaV0(ctx, req, &resp)

	if resp.Diagnostics.HasError() {
		t.Fatalf("upgrade failed: %v", resp.Diagnostics)
	}

	var name types.String
	resp.Diagnostics.Append(resp.State.GetAttribute(ctx, path.Root("name"), &name)...)
	if resp.Diagnostics.HasError() {
		t.Fatalf("read upgraded state: %v", resp.Diagnostics)
	}
	if name.ValueString() != "test" {
		t.Fatalf("name = %q, want %q", name.ValueString(), "test")
	}
}

// TargetAccessPolicyModel must match the live schema, or the moved block from
// cloudflare_access_policy fails when it writes the target state.
func TestMoveFromAccessPolicy_WritesTargetState(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	sourceSchema := v500.SourceAccessPolicySchema()
	sourceValue, err := (&tfprotov6.RawState{JSON: []byte(`{
		"id": "policy-id",
		"account_id": "account-id",
		"application_id": "app-id",
		"name": "test",
		"precedence": 1,
		"decision": "allow"
	}`)}).Unmarshal(sourceSchema.Type().TerraformType(ctx))
	if err != nil {
		t.Fatalf("build v4 source state: %v", err)
	}

	targetSchema := zero_trust_access_policy.ResourceSchema(ctx)
	req := resource.MoveStateRequest{
		SourceTypeName:        "cloudflare_access_policy",
		SourceProviderAddress: "registry.terraform.io/cloudflare/cloudflare",
		SourceState:           &tfsdk.State{Schema: sourceSchema, Raw: sourceValue},
	}
	resp := resource.MoveStateResponse{
		TargetState: tfsdk.State{Schema: targetSchema, Raw: tftypes.NewValue(targetSchema.Type().TerraformType(ctx), nil)},
	}

	v500.MoveFromAccessPolicy(ctx, req, &resp)

	if resp.Diagnostics.HasError() {
		t.Fatalf("move failed: %v", resp.Diagnostics)
	}

	var name types.String
	resp.Diagnostics.Append(resp.TargetState.GetAttribute(ctx, path.Root("name"), &name)...)
	if resp.Diagnostics.HasError() {
		t.Fatalf("read moved state: %v", resp.Diagnostics)
	}
	if name.ValueString() != "test" {
		t.Fatalf("name = %q, want %q", name.ValueString(), "test")
	}
}

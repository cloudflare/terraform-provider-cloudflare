// This file is intentionally NOT code-generated.

package v500_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/api_shield_operation"
	v500 "github.com/cloudflare/terraform-provider-cloudflare/internal/services/api_shield_operation/migration/v500"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

const v4APIShieldOperationState = `{
	"id": "11111111-1111-4111-8111-111111111111",
	"zone_id": "023e105f4ecef8ad9ca31a8372d0c353",
	"method": "GET",
	"host": "api.example.com",
	"endpoint": "/api/v1/users/{var1}"
}`

// The target model must convert into the live schema, which dropped
// features.schema_info.learned_available and active_schema.is_learned, or
// UpgradeFromV4 fails with a Value Conversion Error (#7436).
func TestAPIShieldOperationMigrationFromV4SetsState(t *testing.T) {
	ctx := context.Background()
	sourceSchema := v500.SourceCloudflareAPIShieldOperationSchema()
	targetSchema := api_shield_operation.ResourceSchema(ctx)
	raw, err := (&tfprotov6.RawState{JSON: []byte(v4APIShieldOperationState)}).Unmarshal(sourceSchema.Type().TerraformType(ctx))
	if err != nil {
		t.Fatalf("could not decode v4 state: %v", err)
	}
	sourceState := tfsdk.State{Raw: raw, Schema: sourceSchema}

	resp := resource.UpgradeStateResponse{State: tfsdk.State{Schema: targetSchema}}
	v500.UpgradeFromV4(ctx, resource.UpgradeStateRequest{State: &sourceState}, &resp)
	if resp.Diagnostics.HasError() {
		t.Fatalf("UpgradeFromV4 failed: %v", resp.Diagnostics)
	}

	var got api_shield_operation.APIShieldOperationModel
	if diags := resp.State.Get(ctx, &got); diags.HasError() {
		t.Fatalf("could not decode migrated state: %v", diags)
	}
	id := types.StringValue("11111111-1111-4111-8111-111111111111")
	if !got.ID.Equal(id) || !got.OperationID.Equal(id) {
		t.Errorf("id = %s, operation_id = %s, want %s", got.ID, got.OperationID, id)
	}
	if !got.Endpoint.Equal(types.StringValue("/api/v1/users/{var1}")) || !got.Method.Equal(types.StringValue("GET")) {
		t.Errorf("endpoint = %s, method = %s", got.Endpoint, got.Method)
	}
	if !got.Features.IsNull() {
		t.Errorf("features = %s, want null", got.Features)
	}
}

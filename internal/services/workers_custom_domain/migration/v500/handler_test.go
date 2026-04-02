package v500

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

func TestUpgradeStateV0toV500(t *testing.T) {
	ctx := context.Background()
	sourceSchema := SourceSchemaV0()

	raw := createTerraformValue(t, sourceSchema, map[string]interface{}{
		"id":          "domain-id",
		"account_id":  "account-id",
		"hostname":    "api.example.com",
		"service":     "worker-service",
		"zone_id":     "zone-id",
		"environment": "production",
		"zone_name":   "example.com",
	})

	state := tfsdk.State{Raw: raw, Schema: sourceSchema}
	req := resource.UpgradeStateRequest{State: &state}
	resp := &resource.UpgradeStateResponse{State: tfsdk.State{Schema: sourceSchema}}

	upgrader := UpgradeStateV0toV500(sourceSchema)
	upgrader.StateUpgrader(ctx, req, resp)

	if resp.Diagnostics.HasError() {
		t.Fatalf("unexpected diagnostics: %v", resp.Diagnostics)
	}

	var got TargetWorkersCustomDomainModel
	if diags := resp.State.Get(ctx, &got); diags.HasError() {
		t.Fatalf("failed to read upgraded state: %v", diags)
	}

	if got.Hostname.ValueString() != "api.example.com" ||
		got.Environment.ValueString() != "production" ||
		got.ZoneName.ValueString() != "example.com" {
		t.Fatalf("unexpected upgraded state: %+v", got)
	}
}

func TestMoveStateV4toV500(t *testing.T) {
	ctx := context.Background()
	sourceSchema := SourceSchemaV0()

	raw := createTerraformValue(t, sourceSchema, map[string]interface{}{
		"id":          "domain-id",
		"account_id":  "account-id",
		"hostname":    "api.example.com",
		"service":     "worker-service",
		"zone_id":     "zone-id",
		"environment": "staging",
		"zone_name":   "example.com",
	})

	sourceState := tfsdk.State{Raw: raw, Schema: sourceSchema}
	req := resource.MoveStateRequest{
		SourceTypeName:        "cloudflare_worker_domain",
		SourceProviderAddress: "registry.terraform.io/cloudflare/cloudflare",
		SourceSchemaVersion:   0,
		SourceState:           &sourceState,
	}
	resp := &resource.MoveStateResponse{TargetState: tfsdk.State{Schema: sourceSchema}}

	MoveStateV4toV500(ctx, req, resp)

	if resp.Diagnostics.HasError() {
		t.Fatalf("unexpected diagnostics: %v", resp.Diagnostics)
	}

	var got TargetWorkersCustomDomainModel
	if diags := resp.TargetState.Get(ctx, &got); diags.HasError() {
		t.Fatalf("failed to read moved state: %v", diags)
	}

	if got.Environment.ValueString() != "staging" || got.Hostname.ValueString() != "api.example.com" {
		t.Fatalf("unexpected moved state: %+v", got)
	}
}

func TestUpgradeStateV1toV500_NoOp(t *testing.T) {
	ctx := context.Background()
	targetSchema := SourceSchemaV0()

	raw := createTerraformValue(t, targetSchema, map[string]interface{}{
		"id":          "domain-id",
		"account_id":  "account-id",
		"hostname":    "api.example.com",
		"service":     "worker-service",
		"zone_id":     "zone-id",
		"environment": "production",
		"zone_name":   "example.com",
	})

	state := tfsdk.State{Raw: raw, Schema: targetSchema}
	req := resource.UpgradeStateRequest{State: &state}
	resp := &resource.UpgradeStateResponse{State: tfsdk.State{Schema: targetSchema}}

	upgrader := UpgradeStateV1toV500(targetSchema)
	upgrader.StateUpgrader(ctx, req, resp)

	if resp.Diagnostics.HasError() {
		t.Fatalf("unexpected diagnostics: %v", resp.Diagnostics)
	}

	var got TargetWorkersCustomDomainModel
	if diags := resp.State.Get(ctx, &got); diags.HasError() {
		t.Fatalf("failed to read no-op upgraded state: %v", diags)
	}

	if got.ID.ValueString() != "domain-id" || got.ZoneName.ValueString() != "example.com" {
		t.Fatalf("unexpected no-op upgraded state: %+v", got)
	}
}

func TestUpgradeStateV0toV500_NullOptionalField(t *testing.T) {
	ctx := context.Background()
	sourceSchema := SourceSchemaV0()

	raw := createTerraformValue(t, sourceSchema, map[string]interface{}{
		"id":          "domain-id",
		"account_id":  "account-id",
		"hostname":    "api.example.com",
		"service":     "worker-service",
		"zone_id":     "zone-id",
		"environment": nil,
		"zone_name":   nil,
	})

	state := tfsdk.State{Raw: raw, Schema: sourceSchema}
	req := resource.UpgradeStateRequest{State: &state}
	resp := &resource.UpgradeStateResponse{State: tfsdk.State{Schema: sourceSchema}}

	upgrader := UpgradeStateV0toV500(sourceSchema)
	upgrader.StateUpgrader(ctx, req, resp)

	if resp.Diagnostics.HasError() {
		t.Fatalf("unexpected diagnostics: %v", resp.Diagnostics)
	}

	var got TargetWorkersCustomDomainModel
	if diags := resp.State.Get(ctx, &got); diags.HasError() {
		t.Fatalf("failed to read upgraded state: %v", diags)
	}

	if !got.Environment.IsNull() {
		t.Fatalf("expected null environment after upgrade, got: %v", got.Environment)
	}

	if !got.ZoneName.IsNull() {
		t.Fatalf("expected null zone_name after upgrade, got: %v", got.ZoneName)
	}
}

func TestMoveStateV4toV500_IgnoresUnknownSourceType(t *testing.T) {
	ctx := context.Background()
	sourceSchema := SourceSchemaV0()

	raw := createTerraformValue(t, sourceSchema, map[string]interface{}{
		"id":          "domain-id",
		"account_id":  "account-id",
		"hostname":    "api.example.com",
		"service":     "worker-service",
		"zone_id":     "zone-id",
		"environment": "staging",
		"zone_name":   "example.com",
	})

	sourceState := tfsdk.State{Raw: raw, Schema: sourceSchema}
	req := resource.MoveStateRequest{
		SourceTypeName:        "cloudflare_workers_custom_domain",
		SourceProviderAddress: "registry.terraform.io/cloudflare/cloudflare",
		SourceSchemaVersion:   0,
		SourceState:           &sourceState,
	}
	resp := &resource.MoveStateResponse{TargetState: tfsdk.State{Schema: sourceSchema}}

	MoveStateV4toV500(ctx, req, resp)

	if resp.Diagnostics.HasError() {
		t.Fatalf("unexpected diagnostics: %v", resp.Diagnostics)
	}

	if !resp.TargetState.Raw.IsNull() {
		t.Fatalf("expected target state to remain unset for unknown source type")
	}
}

func TestMoveStateV4toV500_IgnoresNonCloudflareProvider(t *testing.T) {
	ctx := context.Background()
	sourceSchema := SourceSchemaV0()

	raw := createTerraformValue(t, sourceSchema, map[string]interface{}{
		"id":          "domain-id",
		"account_id":  "account-id",
		"hostname":    "api.example.com",
		"service":     "worker-service",
		"zone_id":     "zone-id",
		"environment": "staging",
		"zone_name":   "example.com",
	})

	sourceState := tfsdk.State{Raw: raw, Schema: sourceSchema}
	req := resource.MoveStateRequest{
		SourceTypeName:        "cloudflare_worker_domain",
		SourceProviderAddress: "registry.terraform.io/hashicorp/random",
		SourceSchemaVersion:   0,
		SourceState:           &sourceState,
	}
	resp := &resource.MoveStateResponse{TargetState: tfsdk.State{Schema: sourceSchema}}

	MoveStateV4toV500(ctx, req, resp)

	if resp.Diagnostics.HasError() {
		t.Fatalf("unexpected diagnostics: %v", resp.Diagnostics)
	}

	if !resp.TargetState.Raw.IsNull() {
		t.Fatalf("expected target state to remain unset for non-cloudflare provider")
	}
}

func TestIsCloudflareProvider(t *testing.T) {
	testCases := []struct {
		name string
		addr string
		want bool
	}{
		{name: "registry address", addr: "registry.terraform.io/cloudflare/cloudflare", want: true},
		{name: "short cloudflare address", addr: "cloudflare/cloudflare", want: true},
		{name: "different provider", addr: "registry.terraform.io/hashicorp/random", want: false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if got := isCloudflareProvider(tc.addr); got != tc.want {
				t.Fatalf("isCloudflareProvider(%q) = %t, want %t", tc.addr, got, tc.want)
			}
		})
	}
}

func createTerraformValue(t *testing.T, s schema.Schema, values map[string]interface{}) tftypes.Value {
	t.Helper()

	attrTypes := make(map[string]tftypes.Type)
	for name := range s.Attributes {
		attrTypes[name] = tftypes.String
	}

	objectType := tftypes.Object{AttributeTypes: attrTypes}

	valueMap := make(map[string]tftypes.Value)
	for name, attrType := range attrTypes {
		valueMap[name] = tftypes.NewValue(attrType, nil)
	}
	for name, val := range values {
		if val != nil {
			valueMap[name] = tftypes.NewValue(attrTypes[name], val)
		}
	}

	return tftypes.NewValue(objectType, valueMap)
}

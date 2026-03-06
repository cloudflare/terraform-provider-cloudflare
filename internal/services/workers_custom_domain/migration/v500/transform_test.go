package v500

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestTransformV4toV500_Identity(t *testing.T) {
	ctx := context.Background()

	source := SourceV4WorkersCustomDomainModel{
		ID:          types.StringValue("domain-id"),
		AccountID:   types.StringValue("account-id"),
		Hostname:    types.StringValue("api.example.com"),
		Service:     types.StringValue("worker-service"),
		ZoneID:      types.StringValue("zone-id"),
		Environment: types.StringValue("production"),
		ZoneName:    types.StringValue("example.com"),
	}

	target, diags := TransformV4toV500(ctx, source)
	if diags.HasError() {
		t.Fatalf("unexpected diagnostics: %v", diags)
	}

	if target.ID != source.ID ||
		target.AccountID != source.AccountID ||
		target.Hostname != source.Hostname ||
		target.Service != source.Service ||
		target.ZoneID != source.ZoneID ||
		target.Environment != source.Environment ||
		target.ZoneName != source.ZoneName {
		t.Fatalf("identity mapping failed: got %+v, source %+v", target, source)
	}
}

func TestTransformV4toV500_NullAndEmptyValues(t *testing.T) {
	ctx := context.Background()

	source := SourceV4WorkersCustomDomainModel{
		ID:          types.StringNull(),
		AccountID:   types.StringValue("account-id"),
		Hostname:    types.StringValue(""),
		Service:     types.StringValue("worker-service"),
		ZoneID:      types.StringValue("zone-id"),
		Environment: types.StringNull(),
		ZoneName:    types.StringNull(),
	}

	target, diags := TransformV4toV500(ctx, source)
	if diags.HasError() {
		t.Fatalf("unexpected diagnostics: %v", diags)
	}

	if !target.ID.IsNull() {
		t.Fatalf("expected null id, got: %v", target.ID)
	}

	if target.Hostname.ValueString() != "" {
		t.Fatalf("expected empty hostname to be preserved, got: %q", target.Hostname.ValueString())
	}

	if !target.Environment.IsNull() {
		t.Fatalf("expected null environment, got: %v", target.Environment)
	}

	if !target.ZoneName.IsNull() {
		t.Fatalf("expected null zone_name, got: %v", target.ZoneName)
	}
}

func TestTransformV4toV500_UnknownValues(t *testing.T) {
	ctx := context.Background()

	source := SourceV4WorkersCustomDomainModel{
		ID:          types.StringUnknown(),
		AccountID:   types.StringValue("account-id"),
		Hostname:    types.StringUnknown(),
		Service:     types.StringValue("worker-service"),
		ZoneID:      types.StringValue("zone-id"),
		Environment: types.StringUnknown(),
		ZoneName:    types.StringUnknown(),
	}

	target, diags := TransformV4toV500(ctx, source)
	if diags.HasError() {
		t.Fatalf("unexpected diagnostics: %v", diags)
	}

	if !target.ID.IsUnknown() || !target.Hostname.IsUnknown() || !target.Environment.IsUnknown() || !target.ZoneName.IsUnknown() {
		t.Fatalf("expected unknown values to be preserved, got: %+v", target)
	}
}

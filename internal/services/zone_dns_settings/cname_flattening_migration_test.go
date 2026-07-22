package zone_dns_settings

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zone_setting"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestFlattenAllCNAMEsFromLegacyValue(t *testing.T) {
	t.Parallel()

	for _, test := range []struct {
		name  string
		value types.String
		want  types.Bool
	}{
		{name: "on", value: types.StringValue("on"), want: types.BoolValue(true)},
		{name: "flatten all", value: types.StringValue("flatten_all"), want: types.BoolValue(true)},
		{name: "off", value: types.StringValue("off"), want: types.BoolValue(false)},
		{name: "apex", value: types.StringValue("apex"), want: types.BoolValue(false)},
		{name: "flatten at root", value: types.StringValue("flatten_at_root"), want: types.BoolValue(false)},
	} {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			got, diags := flattenAllCNAMEsFromLegacyValue(context.Background(), customfield.RawNormalizedDynamicValueFrom(test.value))
			if diags.HasError() {
				t.Fatalf("unexpected diagnostics: %v", diags)
			}
			if !got.Equal(test.want) {
				t.Errorf("flattenAllCNAMEsFromLegacyValue(%q) = %t, want %t", test.value.ValueString(), got.ValueBool(), test.want.ValueBool())
			}
		})
	}
}

func TestMoveCNAMEFlatteningState(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	sourceSchema := zone_setting.ResourceSchema(ctx)
	sourceState := tfsdk.State{Schema: sourceSchema}
	diags := sourceState.Set(ctx, zone_setting.ZoneSettingModel{
		ID:        types.StringValue(zone_setting.CNAMEFlatteningSettingID),
		SettingID: types.StringValue(zone_setting.CNAMEFlatteningSettingID),
		ZoneID:    types.StringValue("zone-id"),
		Value:     customfield.RawNormalizedDynamicValueFrom(types.StringValue("flatten_all")),
	})
	if diags.HasError() {
		t.Fatalf("failed to create source state: %v", diags)
	}

	targetSchema := ResourceSchema(ctx)
	resp := &resource.MoveStateResponse{TargetState: tfsdk.State{Schema: targetSchema}}
	moveCNAMEFlatteningState(ctx, resource.MoveStateRequest{
		SourceTypeName: "cloudflare_zone_setting",
		SourceState:    &sourceState,
	}, resp)
	if resp.Diagnostics.HasError() {
		t.Fatalf("unexpected diagnostics: %v", resp.Diagnostics)
	}

	var target ZoneDNSSettingsModel
	if diags := resp.TargetState.Get(ctx, &target); diags.HasError() {
		t.Fatalf("failed to read target state: %v", diags)
	}
	if target.ZoneID.ValueString() != "zone-id" || !target.FlattenAllCNAMEs.ValueBool() {
		t.Fatalf("unexpected migrated state: %+v", target)
	}
}

func TestFlattenAllCNAMEsFromLegacyValueRejectsUnsupportedValue(t *testing.T) {
	t.Parallel()

	_, diags := flattenAllCNAMEsFromLegacyValue(context.Background(), customfield.RawNormalizedDynamicValueFrom(types.StringValue("unknown")))
	if !diags.HasError() {
		t.Fatal("expected an error diagnostic")
	}
}

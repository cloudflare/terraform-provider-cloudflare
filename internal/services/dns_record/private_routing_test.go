package dns_record_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/dns_record"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

const privateRoutingTimestamp = "2026-10-07T08:48:25Z"

func privateRoutingModel(ctx context.Context, privateRouting types.Bool) dns_record.DNSRecordModel {
	return dns_record.DNSRecordModel{
		ID:             types.StringValue("record-id"),
		ZoneID:         types.StringValue("zone-id"),
		Name:           types.StringValue("terraform.example.com"),
		Type:           types.StringValue("A"),
		Content:        types.StringValue("192.0.2.1"),
		PrivateRouting: privateRouting,
		Proxied:        types.BoolValue(false),
		TTL:            types.Float64Value(300),
		Tags:           customfield.NewSetMust[types.String](ctx, nil),
		Settings:       customfield.NullObject[dns_record.DNSRecordSettingsModel](ctx),
		CreatedOn:      timetypes.NewRFC3339ValueMust(privateRoutingTimestamp),
		ModifiedOn:     timetypes.NewRFC3339ValueMust(privateRoutingTimestamp),
		Proxiable:      types.BoolValue(true),
		Meta:           jsontypes.NewNormalizedValue("{}"),
	}
}

func TestDNSRecordModifyPlan_PrivateRoutingChangesInvalidateModifiedOn(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		name        string
		state       types.Bool
		plan        types.Bool
		wantUnknown bool
	}{
		{"unchanged_null", types.BoolNull(), types.BoolNull(), false},
		{"unchanged_false", types.BoolValue(false), types.BoolValue(false), false},
		{"unchanged_true", types.BoolValue(true), types.BoolValue(true), false},
		{"enable_from_null", types.BoolNull(), types.BoolValue(true), true},
		{"enable_from_false", types.BoolValue(false), types.BoolValue(true), true},
		{"disable_to_false", types.BoolValue(true), types.BoolValue(false), true},
		{"remove_true", types.BoolValue(true), types.BoolNull(), true},
		{"remove_false", types.BoolValue(false), types.BoolNull(), true},
		{"set_false_from_null", types.BoolNull(), types.BoolValue(false), true},
	} {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			stateModel := privateRoutingModel(ctx, tt.state)
			state := tfsdk.State{Schema: dns_record.ResourceSchema(ctx)}
			if diags := state.Set(ctx, &stateModel); diags.HasError() {
				t.Fatalf("set state: %v", diags)
			}

			planModel := privateRoutingModel(ctx, tt.plan)
			planModel.ModifiedOn = timetypes.NewRFC3339Unknown()
			plan := tfsdk.Plan{Schema: dns_record.ResourceSchema(ctx)}
			if diags := plan.Set(ctx, &planModel); diags.HasError() {
				t.Fatalf("set plan: %v", diags)
			}

			response := &resource.ModifyPlanResponse{Plan: plan}
			resourceUnderTest, ok := dns_record.NewResource().(*dns_record.DNSRecordResource)
			if !ok {
				t.Fatal("DNS record resource has unexpected type")
			}
			resourceUnderTest.ModifyPlan(ctx, resource.ModifyPlanRequest{State: state, Plan: plan}, response)
			if response.Diagnostics.HasError() {
				t.Fatalf("modify plan: %v", response.Diagnostics)
			}

			var result dns_record.DNSRecordModel
			if diags := response.Plan.Get(ctx, &result); diags.HasError() {
				t.Fatalf("get plan: %v", diags)
			}
			if result.ModifiedOn.IsUnknown() != tt.wantUnknown {
				t.Errorf("modified_on = %s, want unknown = %v", result.ModifiedOn, tt.wantUnknown)
			}
			if !tt.wantUnknown && result.ModifiedOn.ValueString() != privateRoutingTimestamp {
				t.Errorf("unchanged record lost modified_on: %s", result.ModifiedOn)
			}
		})
	}
}

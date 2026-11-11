package dns_record_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/dns_record"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestDNSRecordModifyPlan_CNAMEContentNormalization(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		name        string
		state       string
		plan        string
		wantContent string
		wantUnknown bool
	}{
		{"trailing_dot", "x.aws", "x.aws.", "x.aws", false},
		{"case_and_trailing_dot", "x.aws", "X.AWS.", "x.aws", false},
		{"case_only", "x.aws", "X.aws", "x.aws", false},
		{"real_change", "x.aws", "y.aws.", "y.aws.", true},
		{"two_trailing_dots", "x.aws", "x.aws..", "x.aws..", true},
		{"dot_only_vs_empty", "", ".", ".", true},
	} {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			stateModel := privateRoutingModel(ctx, types.BoolNull())
			stateModel.Type = types.StringValue("CNAME")
			stateModel.Content = types.StringValue(tt.state)
			state := tfsdk.State{Schema: dns_record.ResourceSchema(ctx)}
			if diags := state.Set(ctx, &stateModel); diags.HasError() {
				t.Fatalf("set state: %v", diags)
			}

			planModel := privateRoutingModel(ctx, types.BoolNull())
			planModel.Type = types.StringValue("CNAME")
			planModel.Content = types.StringValue(tt.plan)
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
			if got := result.Content.ValueString(); got != tt.wantContent {
				t.Errorf("content = %q, want %q", got, tt.wantContent)
			}
			if result.ModifiedOn.IsUnknown() != tt.wantUnknown {
				t.Errorf("modified_on = %s, want unknown = %v", result.ModifiedOn, tt.wantUnknown)
			}
		})
	}
}

package zero_trust_access_application

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestSessionDurationValidation(t *testing.T) {
	t.Parallel()

	for _, tc := range []struct {
		name      string
		appType   types.String
		duration  types.String
		wantError bool
	}{
		{"default", types.StringValue("infrastructure"), types.StringValue("24h"), false},
		{"seconds", types.StringValue("infrastructure"), types.StringValue("30s"), false},
		{"compound", types.StringValue("infrastructure"), types.StringValue("2h45m"), false},
		{"fractional_minutes_whole_seconds", types.StringValue("infrastructure"), types.StringValue("1.5m"), false},
		{"fractional_seconds", types.StringValue("infrastructure"), types.StringValue("1.5s"), true},
		{"milliseconds", types.StringValue("infrastructure"), types.StringValue("1000ms"), true},
		{"microseconds", types.StringValue("infrastructure"), types.StringValue("1000000us"), true},
		{"nanoseconds", types.StringValue("infrastructure"), types.StringValue("1000000000ns"), true},
		{"invalid_unit", types.StringValue("infrastructure"), types.StringValue("1d"), true},
		{"omitted", types.StringValue("infrastructure"), types.StringNull(), false},
		{"unknown_duration", types.StringValue("infrastructure"), types.StringUnknown(), false},
		{"unknown_type", types.StringUnknown(), types.StringValue("300ms"), false},
		{"self_hosted_milliseconds", types.StringValue("self_hosted"), types.StringValue("300ms"), false},
		{"self_hosted_fractional_seconds", types.StringValue("self_hosted"), types.StringValue("1.5s"), false},
		{"self_hosted_invalid_unit", types.StringValue("self_hosted"), types.StringValue("1d"), true},
		{"unsupported_type", types.StringValue("bookmark"), types.StringValue("24h"), true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			resourceSchema := ResourceSchema(ctx)
			durationSchema := resourceSchema.Attributes["session_duration"].(schema.StringAttribute)
			state := tfsdk.State{Schema: schema.Schema{Attributes: map[string]schema.Attribute{
				"type":             resourceSchema.Attributes["type"],
				"session_duration": durationSchema,
			}}}
			if diags := state.Set(ctx, &struct {
				Type            types.String `tfsdk:"type"`
				SessionDuration types.String `tfsdk:"session_duration"`
			}{tc.appType, tc.duration}); diags.HasError() {
				t.Fatalf("build config: %v", diags)
			}
			req := validator.StringRequest{
				Config:         tfsdk.Config{Schema: state.Schema, Raw: state.Raw},
				ConfigValue:    tc.duration,
				Path:           path.Root("session_duration"),
				PathExpression: path.MatchRoot("session_duration"),
			}
			var resp validator.StringResponse
			for _, v := range durationSchema.Validators {
				v.ValidateString(ctx, req, &resp)
			}
			if resp.Diagnostics.HasError() != tc.wantError {
				t.Errorf("validation errors = %v, want error %v", resp.Diagnostics, tc.wantError)
			}
		})
	}
}

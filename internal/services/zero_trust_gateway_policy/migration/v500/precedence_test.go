package v500_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_gateway_policy"
	v500 "github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_gateway_policy/migration/v500"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

// Schema migration must preserve source precedence, including any hash suffixes:
// removing a suffix can collapse distinct rules or invert manually assigned order.
func TestGatewayPolicyMigrationPreservesPrecedence(t *testing.T) {
	for _, migrationPath := range []string{"MoveState", "UpgradeFromVersion0"} {
		t.Run(migrationPath, func(t *testing.T) {
			for _, tc := range []struct {
				name        string
				ruleNames   [2]string
				actions     [2]string
				precedences [2]int64
			}{
				{
					// BUGS-2126: different live rules shared configured precedence 10116.
					name:        "ticket_values_remain_distinct",
					ruleNames:   [2]string{"GitLab Staging", "Harbor Staging"},
					actions:     [2]string{"allow", "allow"},
					precedences: [2]int64{10116667, 10116781},
				},
				{
					// API-assigned values are not necessarily v4 hash encodings.
					name:        "allow_stays_before_block_across_threshold",
					ruleNames:   [2]string{"Specific allow", "Catch-all block"},
					actions:     [2]string{"allow", "block"},
					precedences: [2]int64{9999, 10000},
				},
				{
					name:        "ordinary_values_are_unchanged",
					ruleNames:   [2]string{"Specific allow", "Catch-all block"},
					actions:     [2]string{"allow", "block"},
					precedences: [2]int64{1000, 2000},
				},
			} {
				t.Run(tc.name, func(t *testing.T) {
					var got [2]int64
					for i, precedence := range tc.precedences {
						state := migrateGatewayPolicyPrecedence(t, migrationPath, tc.ruleNames[i], tc.actions[i], precedence)
						if state.Precedence.IsNull() || state.Precedence.IsUnknown() {
							t.Fatalf("%s: migrated precedence must be known, got %s", tc.ruleNames[i], state.Precedence)
						}
						got[i] = state.Precedence.ValueInt64()
						if got[i] != precedence {
							t.Errorf("%s: precedence changed from %d to %d during schema migration", tc.ruleNames[i], precedence, got[i])
						}
						if !state.Name.Equal(types.StringValue(tc.ruleNames[i])) || !state.Action.Equal(types.StringValue(tc.actions[i])) {
							t.Fatalf("migration changed rule identity or action: name=%s action=%s", state.Name, state.Action)
						}
					}
					if got[0] == got[1] {
						t.Errorf("distinct source precedences %v collapsed to %v", tc.precedences, got)
					}
					if got[0] >= got[1] {
						t.Errorf("policy order was not preserved: %q must precede %q; source=%v migrated=%v", tc.ruleNames[0], tc.ruleNames[1], tc.precedences, got)
					}
				})
			}
		})
	}
}

func migrateGatewayPolicyPrecedence(t *testing.T, migrationPath, name, action string, precedence int64) zero_trust_gateway_policy.ZeroTrustGatewayPolicyModel {
	t.Helper()
	rawState := &tfprotov6.RawState{JSON: []byte(fmt.Sprintf(`{
		"id": "11111111-1111-4111-8111-111111111111",
		"account_id": "f037e56e89293a057740de681ac9abbe",
		"name": %q,
		"description": "Gateway precedence migration regression",
		"precedence": %d,
		"enabled": true,
		"action": %q,
		"filters": ["l4"],
		"traffic": "net.dst.ip in {192.0.2.1}",
		"identity": "",
		"device_posture": "",
		"version": 1,
		"rule_settings": []
	}`, name, precedence, action))}
	got, err := migrateGatewayPolicyRawState(context.Background(), migrationPath, rawState)
	if err != nil {
		t.Fatal(err)
	}
	return got
}

func migrateGatewayPolicyRawState(ctx context.Context, migrationPath string, rawState *tfprotov6.RawState) (zero_trust_gateway_policy.ZeroTrustGatewayPolicyModel, error) {
	var got zero_trust_gateway_policy.ZeroTrustGatewayPolicyModel
	sourceSchema := v500.SourceCloudflareTeamsRuleSchema()
	targetSchema := zero_trust_gateway_policy.ResourceSchema(ctx)

	var targetState tfsdk.State
	switch migrationPath {
	case "MoveState":
		raw, err := rawState.Unmarshal(sourceSchema.Type().TerraformType(ctx))
		if err != nil {
			return got, fmt.Errorf("could not decode v4 source state: %w", err)
		}
		sourceState := tfsdk.State{Raw: raw, Schema: sourceSchema}
		resp := resource.MoveStateResponse{TargetState: tfsdk.State{Schema: targetSchema}}
		v500.MoveState(ctx, resource.MoveStateRequest{
			SourceTypeName:        "cloudflare_teams_rule",
			SourceProviderAddress: "registry.terraform.io/cloudflare/cloudflare",
			SourceSchemaVersion:   0,
			SourceState:           &sourceState,
		}, &resp)
		if resp.Diagnostics.HasError() {
			return got, fmt.Errorf("MoveState failed: %v", resp.Diagnostics)
		}
		targetState = resp.TargetState
	case "UpgradeFromVersion0":
		resp := resource.UpgradeStateResponse{State: tfsdk.State{Schema: targetSchema}}
		v500.UpgradeFromVersion0(ctx, resource.UpgradeStateRequest{RawState: rawState}, &resp)
		if resp.Diagnostics.HasError() {
			return got, fmt.Errorf("UpgradeFromVersion0 failed: %v", resp.Diagnostics)
		}
		targetState = resp.State
	default:
		return got, fmt.Errorf("unknown migration path %q", migrationPath)
	}

	if diags := targetState.Get(ctx, &got); diags.HasError() {
		return got, fmt.Errorf("could not decode migrated state: %v", diags)
	}
	return got, nil
}

func TestGatewayPolicyTransformPreservesNullAndUnknownPrecedence(t *testing.T) {
	for _, tc := range []struct {
		name       string
		precedence types.Int64
	}{
		{name: "null", precedence: types.Int64Null()},
		{name: "unknown", precedence: types.Int64Unknown()},
	} {
		t.Run(tc.name, func(t *testing.T) {
			got, diags := v500.Transform(context.Background(), &v500.SourceCloudflareTeamsRuleModel{
				AccountID:  types.StringValue("f037e56e89293a057740de681ac9abbe"),
				Name:       types.StringValue("Specific allow"),
				Action:     types.StringValue("allow"),
				Precedence: tc.precedence,
			})
			if diags.HasError() {
				t.Fatalf("Transform failed: %v", diags)
			}
			if !got.Precedence.Equal(tc.precedence) {
				t.Errorf("precedence changed from %s to %s", tc.precedence, got.Precedence)
			}
		})
	}
}

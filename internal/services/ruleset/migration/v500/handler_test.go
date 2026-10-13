package v500_test

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/ruleset"
	v500 "github.com/cloudflare/terraform-provider-cloudflare/internal/services/ruleset/migration/v500"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	providerschema "github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/stretchr/testify/require"
)

// Exercise the actual provider protocol so a wrong PriorSchema fails before the
// handler, exactly as in #7135. No API calls, credentials, or customer state are used.
func TestRulesetUpgradeResourceStateV4(t *testing.T) {
	t.Parallel()
	fields := []struct {
		name      string
		params    map[string]interface{}
		wantAsset types.String
		wantBool  types.Bool
	}{
		{"older_v4_absent", map[string]interface{}{}, types.StringNull(), types.BoolNull()},
		{"v4_52_8_asset_null", map[string]interface{}{"asset_name": nil}, types.StringNull(), types.BoolNull()},
		{"v4_52_8_asset_value", map[string]interface{}{"asset_name": "custom-error"}, types.StringValue("custom-error"), types.BoolNull()},
		{"v4_52_9_null_padding", map[string]interface{}{"asset_name": nil, "content_converter": nil}, types.StringNull(), types.BoolNull()},
		{"converter_false", map[string]interface{}{"asset_name": nil, "content_converter": false}, types.StringNull(), types.BoolValue(false)},
		{"converter_true", map[string]interface{}{"asset_name": nil, "content_converter": true}, types.StringNull(), types.BoolValue(true)},
		{"both_populated", map[string]interface{}{"asset_name": "custom-error", "content_converter": true}, types.StringValue("custom-error"), types.BoolValue(true)},
	}
	for _, version := range []int64{0, 1} {
		for _, tc := range fields {
			t.Run(fmt.Sprintf("version_%d/%s", version, tc.name), func(t *testing.T) {
				t.Parallel()
				params := make(map[string]interface{}, len(tc.params)+1)
				for key, value := range tc.params {
					params[key] = value
				}
				rule := map[string]interface{}{
					"id": "rule-id", "ref": "stable-ref", "action": "set_config", "expression": "true", "enabled": true,
					"action_parameters": []interface{}{params},
					"logging":           []interface{}{}, "ratelimit": []interface{}{}, "exposed_credential_check": []interface{}{},
				}
				if version == 0 {
					// These fields were removed by v4's own 0→1 upgrade.
					rule["version"] = "7"
					rule["last_updated"] = "2026-01-01T00:00:00Z"
					params["version"] = "3"
				}
				input := syntheticRulesetState([]interface{}{rule})
				response := upgradeThroughProtocol(t, version, input)
				require.Empty(t, response.Diagnostics)
				var result v500.TargetV5RulesetModel
				state := upgradedState(t, response)
				diags := state.Get(context.Background(), &result)
				require.False(t, diags.HasError(), "%v", diags)
				require.Len(t, result.Rules, 1)
				got := result.Rules[0]
				require.NotNil(t, got.ActionParameters)
				require.True(t, tc.wantAsset.Equal(got.ActionParameters.AssetName))
				require.True(t, tc.wantBool.Equal(got.ActionParameters.ContentConverter))
				require.Equal(t, "stable-ref", got.Ref.ValueString())
				require.Equal(t, "rule-id", got.ID.ValueString())
				require.Equal(t, "true", got.Expression.ValueString())
				require.True(t, got.Enabled.ValueBool())
				if version == 0 {
					var decoded map[string]tftypes.Value
					require.NoError(t, state.Raw.As(&decoded))
					var outRules []tftypes.Value
					require.NoError(t, decoded["rules"].As(&outRules))
					var outRule map[string]tftypes.Value
					require.NoError(t, outRules[0].As(&outRule))
					require.NotContains(t, outRule, "version")
					require.NotContains(t, outRule, "last_updated")
					var outParams map[string]tftypes.Value
					require.NoError(t, outRule["action_parameters"].As(&outParams))
					require.NotContains(t, outParams, "version")
				}
			})
		}
	}
}

func TestRulesetUpgradeResourceStateV5Unchanged(t *testing.T) {
	t.Parallel()
	for _, version := range []int64{0, 1} {
		t.Run(fmt.Sprintf("version_%d", version), func(t *testing.T) {
			t.Parallel()
			input := syntheticRulesetState([]interface{}{map[string]interface{}{
				"action": "set_config", "expression": "true", "enabled": true, "ref": "keep-me",
				"action_parameters": map[string]interface{}{
					"content_converter": true,
					// This list exists in both formats and must not imply v4.
					"algorithms": []interface{}{map[string]interface{}{"name": "gzip"}},
				},
			}})
			raw, err := json.Marshal(input)
			require.NoError(t, err)
			schema := ruleset.ResourceSchema(context.Background())
			want, err := (&tfprotov6.RawState{JSON: raw}).Unmarshal(schema.Type().TerraformType(context.Background()))
			require.NoError(t, err)
			response := upgradeThroughProtocol(t, version, input)
			require.Empty(t, response.Diagnostics)
			got := upgradedState(t, response)
			require.True(t, want.Equal(got.Raw), "v5 state must be preserved")
		})
	}
}

func TestRulesetUpgradeResourceStateV4LaterRuleAndOtherBlocks(t *testing.T) {
	t.Parallel()
	for _, field := range []string{"action_parameters", "logging", "ratelimit", "exposed_credential_check"} {
		t.Run(field, func(t *testing.T) {
			t.Parallel()
			secondRule := map[string]interface{}{"action": "block", "expression": "false"}
			secondRule[field] = []interface{}{}
			if field == "logging" {
				secondRule[field] = []interface{}{map[string]interface{}{"enabled": false}}
			}
			response := upgradeThroughProtocol(t, 0, syntheticRulesetState([]interface{}{
				map[string]interface{}{"action": "block", "expression": "true", "action_parameters": nil},
				secondRule,
			}))
			require.Empty(t, response.Diagnostics)
			var result v500.TargetV5RulesetModel
			state := upgradedState(t, response)
			diags := state.Get(context.Background(), &result)
			require.False(t, diags.HasError(), "%v", diags)
			require.Len(t, result.Rules, 2)
			require.Equal(t, "false", result.Rules[1].Expression.ValueString())
			if field == "logging" {
				require.NotNil(t, result.Rules[1].Logging)
				require.False(t, result.Rules[1].Logging.Enabled.ValueBool())
			}
		})
	}
}

func TestRulesetUpgradeResourceStateRejectsUnknownFields(t *testing.T) {
	t.Parallel()
	// Do not fix schema drift by silently discarding arbitrary unknown fields,
	// including null fields: the same key might later hold configuration.
	input := syntheticRulesetState([]interface{}{map[string]interface{}{
		"action": "set_config", "expression": "true",
		"action_parameters": []interface{}{map[string]interface{}{"future_field": nil}},
	}})
	response := upgradeThroughProtocol(t, 1, input)
	require.NotEmpty(t, response.Diagnostics)
	require.Contains(t, response.Diagnostics[0].Detail, "future_field")
}

func syntheticRulesetState(rules []interface{}) map[string]interface{} {
	return map[string]interface{}{
		"id": "ruleset-id", "zone_id": "zone-id", "name": "migration-test",
		"kind": "zone", "phase": "http_config_settings", "rules": rules,
	}
}

func upgradeThroughProtocol(t *testing.T, version int64, input map[string]interface{}) *tfprotov6.UpgradeResourceStateResponse {
	t.Helper()
	raw, err := json.Marshal(input)
	require.NoError(t, err)
	server := providerserver.NewProtocol6(&rulesetMigrationTestProvider{})()
	response, err := server.UpgradeResourceState(context.Background(), &tfprotov6.UpgradeResourceStateRequest{
		TypeName: "cloudflare_ruleset", Version: version, RawState: &tfprotov6.RawState{JSON: raw},
	})
	require.NoError(t, err)
	return response
}

func upgradedState(t *testing.T, response *tfprotov6.UpgradeResourceStateResponse) tfsdk.State {
	t.Helper()
	require.NotNil(t, response.UpgradedState)
	schema := ruleset.ResourceSchema(context.Background())
	raw, err := response.UpgradedState.Unmarshal(schema.Type().TerraformType(context.Background()))
	require.NoError(t, err)
	return tfsdk.State{Schema: schema, Raw: raw}
}

type rulesetMigrationTestProvider struct{}

func (*rulesetMigrationTestProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "cloudflare"
	resp.Version = "test"
}

func (*rulesetMigrationTestProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = providerschema.Schema{}
}

func (*rulesetMigrationTestProvider) Configure(context.Context, provider.ConfigureRequest, *provider.ConfigureResponse) {
}

func (*rulesetMigrationTestProvider) DataSources(context.Context) []func() datasource.DataSource {
	return nil
}

func (*rulesetMigrationTestProvider) Resources(context.Context) []func() resource.Resource {
	return []func() resource.Resource{ruleset.NewResource}
}

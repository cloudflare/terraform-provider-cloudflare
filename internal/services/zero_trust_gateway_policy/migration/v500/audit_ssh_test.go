// This file is intentionally NOT code-generated.

package v500_test

import (
	"context"
	"testing"

	v500 "github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_gateway_policy/migration/v500"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

const v4AuditSSHRuleState = `{
	"id": "11111111-1111-4111-8111-111111111111",
	"account_id": "f037e56e89293a057740de681ac9abbe",
	"name": "Audit SSH",
	"description": "log ssh commands",
	"precedence": 100,
	"enabled": true,
	"action": "audit_ssh",
	"filters": ["l4"],
	"traffic": "net.dst.port == 22",
	"identity": "",
	"device_posture": "",
	"version": 1,
	"rule_settings": [{
		"block_page_enabled": false,
		"audit_ssh": [{"command_logging": true}]
	}]
}`

// The live schema has no rule_settings.audit_ssh, so the target model must not
// either, or every v4 policy migration fails with a Value Conversion Error.
// A v4 audit_ssh value is dropped.
func TestGatewayPolicyMigrationDropsAuditSSH(t *testing.T) {
	for _, migrationPath := range []string{"MoveState", "UpgradeFromVersion0"} {
		t.Run(migrationPath, func(t *testing.T) {
			ctx := context.Background()
			got, err := migrateGatewayPolicyRawState(ctx, migrationPath, &tfprotov6.RawState{JSON: []byte(v4AuditSSHRuleState)})
			if err != nil {
				t.Fatal(err)
			}
			if !got.Name.Equal(types.StringValue("Audit SSH")) {
				t.Errorf("name = %s", got.Name)
			}
			rs, diags := got.RuleSettings.Value(ctx)
			if diags.HasError() || rs == nil {
				t.Fatalf("rule_settings missing after migration: %v", diags)
			}
			if !rs.BlockPageEnabled.Equal(types.BoolValue(false)) {
				t.Errorf("rule_settings.block_page_enabled = %s, want false", rs.BlockPageEnabled)
			}
		})
	}
}

func TestGatewayPolicyTransformWarnsOnDroppedAuditSSH(t *testing.T) {
	_, diags := v500.Transform(context.Background(), &v500.SourceCloudflareTeamsRuleModel{
		AccountID: types.StringValue("f037e56e89293a057740de681ac9abbe"),
		Name:      types.StringValue("Audit SSH"),
		Action:    types.StringValue("audit_ssh"),
		RuleSettings: []v500.SourceRuleSettingsModel{{
			AuditSSH: []v500.SourceAuditSSHModel{{CommandLogging: types.BoolValue(true)}},
		}},
	})
	if diags.HasError() {
		t.Fatalf("Transform failed: %v", diags)
	}
	if diags.WarningsCount() != 1 {
		t.Errorf("want one warning for the dropped audit_ssh, got %v", diags)
	}
}

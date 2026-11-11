// This file is intentionally NOT code-generated.

package v500_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_access_service_token"
	v500 "github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_access_service_token/migration/v500"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

const v4ServiceTokenState = `{
	"id": "11111111-1111-4111-8111-111111111111",
	"account_id": "f037e56e89293a057740de681ac9abbe",
	"name": "ci-token",
	"client_id": "abc123.access",
	"client_secret": "secret",
	"expires_at": "2027-01-01T00:00:00Z",
	"min_days_for_renewal": 30,
	"duration": "8760h"
}`

// The target model must convert into the live schema, which gained the
// leak_* attributes, or both UpgradeState and the moved block fail with a
// Value Conversion Error.
func TestAccessServiceTokenMigrationFromV4SetsState(t *testing.T) {
	ctx := context.Background()
	sourceSchema := v500.SourceAccessServiceTokenSchema()
	targetSchema := zero_trust_access_service_token.ResourceSchema(ctx)
	raw, err := (&tfprotov6.RawState{JSON: []byte(v4ServiceTokenState)}).Unmarshal(sourceSchema.Type().TerraformType(ctx))
	if err != nil {
		t.Fatalf("could not decode v4 state: %v", err)
	}
	sourceState := tfsdk.State{Raw: raw, Schema: sourceSchema}

	for _, tc := range []struct {
		name    string
		migrate func(t *testing.T) tfsdk.State
	}{
		{
			name: "UpgradeFromV4",
			migrate: func(t *testing.T) tfsdk.State {
				resp := resource.UpgradeStateResponse{State: tfsdk.State{Schema: targetSchema}}
				v500.UpgradeFromV4(ctx, resource.UpgradeStateRequest{State: &sourceState}, &resp)
				if resp.Diagnostics.HasError() {
					t.Fatalf("UpgradeFromV4 failed: %v", resp.Diagnostics)
				}
				return resp.State
			},
		},
		{
			name: "MoveState",
			migrate: func(t *testing.T) tfsdk.State {
				resp := resource.MoveStateResponse{TargetState: tfsdk.State{Schema: targetSchema}}
				v500.MoveState(ctx, resource.MoveStateRequest{
					SourceTypeName:        "cloudflare_access_service_token",
					SourceProviderAddress: "registry.terraform.io/cloudflare/cloudflare",
					SourceState:           &sourceState,
				}, &resp)
				if resp.Diagnostics.HasError() {
					t.Fatalf("MoveState failed: %v", resp.Diagnostics)
				}
				return resp.TargetState
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			state := tc.migrate(t)
			var got zero_trust_access_service_token.ZeroTrustAccessServiceTokenModel
			if diags := state.Get(ctx, &got); diags.HasError() {
				t.Fatalf("could not decode migrated state: %v", diags)
			}
			if !got.Name.Equal(types.StringValue("ci-token")) || !got.ClientID.Equal(types.StringValue("abc123.access")) {
				t.Errorf("name = %s, client_id = %s", got.Name, got.ClientID)
			}
			if got.ExpiresAt.ValueString() != "2027-01-01T00:00:00Z" {
				t.Errorf("expires_at = %s", got.ExpiresAt)
			}
			if !got.LeakReportURL.IsNull() || !got.LeakReportedAt.IsNull() || !got.LeakReporter.IsNull() {
				t.Errorf("leak_report_url = %s, leak_reported_at = %s, leak_reporter = %s, want null",
					got.LeakReportURL, got.LeakReportedAt, got.LeakReporter)
			}
		})
	}
}

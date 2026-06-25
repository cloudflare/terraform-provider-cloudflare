package zero_trust_access_ai_controls_mcp_portal

import (
	"context"
	"sort"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// TestUpgradeMcpPortalV500ToV501 verifies the list -> set state migration for the
// servers attribute: every element is preserved (by server_id), scalar fields are
// copied through, and the null/unknown cases are handled without error.
func TestUpgradeMcpPortalV500ToV501(t *testing.T) {
	ctx := context.Background()

	t.Run("converts servers list to set preserving elements", func(t *testing.T) {
		priorServers := []ZeroTrustAccessAIControlsMcpPortalServersModel{
			{ServerID: types.StringValue("alpha"), DefaultDisabled: types.BoolValue(true), OnBehalf: types.BoolValue(true)},
			{ServerID: types.StringValue("beta"), DefaultDisabled: types.BoolValue(false), OnBehalf: types.BoolValue(false)},
		}
		prior := zeroTrustAccessAIControlsMcpPortalModelV500{
			ID:      types.StringValue("mcp-sandbox"),
			Name:    types.StringValue("MCP Gateway (Sandbox)"),
			Servers: customfield.NewObjectListMust(ctx, priorServers),
		}

		upgraded, diags := upgradeMcpPortalV500ToV501(ctx, prior)
		if diags.HasError() {
			t.Fatalf("unexpected diagnostics: %v", diags)
		}

		if upgraded.ID.ValueString() != "mcp-sandbox" {
			t.Errorf("id = %q, want %q", upgraded.ID.ValueString(), "mcp-sandbox")
		}
		if upgraded.Name.ValueString() != "MCP Gateway (Sandbox)" {
			t.Errorf("name = %q, want %q", upgraded.Name.ValueString(), "MCP Gateway (Sandbox)")
		}

		got, d := upgraded.Servers.AsStructSliceT(ctx)
		if d.HasError() {
			t.Fatalf("AsStructSliceT: %v", d)
		}
		ids := make([]string, 0, len(got))
		for _, s := range got {
			ids = append(ids, s.ServerID.ValueString())
		}
		sort.Strings(ids)
		want := []string{"alpha", "beta"}
		if len(ids) != len(want) {
			t.Fatalf("got %d servers, want %d", len(ids), len(want))
		}
		for i := range want {
			if ids[i] != want[i] {
				t.Errorf("server_id[%d] = %q, want %q", i, ids[i], want[i])
			}
		}
	})

	t.Run("null servers stays null", func(t *testing.T) {
		prior := zeroTrustAccessAIControlsMcpPortalModelV500{
			ID:      types.StringValue("p"),
			Servers: customfield.NullObjectList[ZeroTrustAccessAIControlsMcpPortalServersModel](ctx),
		}

		upgraded, diags := upgradeMcpPortalV500ToV501(ctx, prior)
		if diags.HasError() {
			t.Fatalf("unexpected diagnostics: %v", diags)
		}
		if !upgraded.Servers.IsNull() {
			t.Errorf("expected servers to remain null")
		}
	})
}

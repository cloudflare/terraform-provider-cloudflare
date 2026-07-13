package zero_trust_access_ai_controls_mcp_portal

import (
	"context"
	"sort"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
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

	t.Run("unknown servers stays unknown", func(t *testing.T) {
		prior := zeroTrustAccessAIControlsMcpPortalModelV500{
			ID:      types.StringValue("p"),
			Servers: customfield.UnknownObjectList[ZeroTrustAccessAIControlsMcpPortalServersModel](ctx),
		}

		upgraded, diags := upgradeMcpPortalV500ToV501(ctx, prior)
		if diags.HasError() {
			t.Fatalf("unexpected diagnostics: %v", diags)
		}
		if !upgraded.Servers.IsUnknown() {
			t.Errorf("expected servers to remain unknown")
		}
	})

	t.Run("preserves unknown nested values", func(t *testing.T) {
		serverType := customfield.NewNestedObjectListType[ZeroTrustAccessAIControlsMcpPortalServersModel](ctx).ElementType().(types.ObjectType)
		updatedPromptsType := serverType.AttrTypes["updated_prompts"].(types.ListType)
		updatedToolsType := serverType.AttrTypes["updated_tools"].(types.ListType)
		server, diags := types.ObjectValue(serverType.AttrTypes, map[string]attr.Value{
			"server_id":        types.StringValue("alpha"),
			"default_disabled": types.BoolValue(false),
			"on_behalf":        types.BoolValue(true),
			"updated_prompts":  types.ListNull(updatedPromptsType.ElemType),
			"updated_tools":    types.ListUnknown(updatedToolsType.ElemType),
		})
		if diags.HasError() {
			t.Fatalf("create server value: %v", diags)
		}
		servers, diags := customfield.NewObjectListFromAttributes[ZeroTrustAccessAIControlsMcpPortalServersModel](
			ctx,
			[]attr.Value{server},
		)
		if diags.HasError() {
			t.Fatalf("create servers list: %v", diags)
		}

		upgraded, diags := upgradeMcpPortalV500ToV501(ctx, zeroTrustAccessAIControlsMcpPortalModelV500{
			ID:      types.StringValue("p"),
			Servers: servers,
		})
		if diags.HasError() {
			t.Fatalf("unexpected diagnostics: %v", diags)
		}

		got := upgraded.Servers.Elements()[0].(types.Object)
		if !got.Attributes()["updated_tools"].IsUnknown() {
			t.Errorf("expected updated_tools to remain unknown")
		}
		if !got.Attributes()["updated_prompts"].IsNull() {
			t.Errorf("expected updated_prompts to remain null")
		}
	})
}

// TestUpgradeMcpPortalV500ToV501_FrameworkDecode exercises the real upgrade path
// that runs on user `terraform apply`: it decodes raw v500 JSON through the List
// source schema, converts it to a Set, and re-encodes it into the v501 schema.
// Terraform persists both List and Set values as JSON arrays, so this also covers
// state written by v5.22.0 after the unversioned switch to Set.
// This de-risks the highest-blast-radius path: a type mismatch here would fail
// apply for every existing user on upgrade.
func TestUpgradeMcpPortalV500ToV501_FrameworkDecode(t *testing.T) {
	ctx := context.Background()

	priorSchema := resourceSchemaV500(ctx)
	rawState := &tfprotov6.RawState{JSON: []byte(`{
		"id":"mcp-sandbox",
		"account_id":"acct",
		"hostname":"mcp.example.com",
		"name":"MCP Gateway (Sandbox)",
		"description":null,
		"allow_code_mode":true,
		"secure_web_gateway":false,
		"servers":[
			{"server_id":"alpha","default_disabled":true,"on_behalf":true,"updated_prompts":null,"updated_tools":null},
			{"server_id":"beta","default_disabled":false,"on_behalf":false,"updated_prompts":null,"updated_tools":null}
		],
		"created_at":null,
		"created_by":null,
		"modified_at":null,
		"modified_by":null
	}`)}
	rawValue, err := rawState.Unmarshal(priorSchema.Type().TerraformType(ctx))
	if err != nil {
		t.Fatalf("decode prior raw state: %v", err)
	}
	priorState := tfsdk.State{Raw: rawValue, Schema: priorSchema}

	r := &ZeroTrustAccessAIControlsMcpPortalResource{}
	upgrader, ok := r.UpgradeState(ctx)[500]
	if !ok {
		t.Fatal("no upgrader registered for prior version 500")
	}

	req := resource.UpgradeStateRequest{RawState: rawState, State: &priorState}
	resp := resource.UpgradeStateResponse{State: tfsdk.State{Schema: ResourceSchema(ctx)}}
	upgrader.StateUpgrader(ctx, req, &resp)
	if resp.Diagnostics.HasError() {
		t.Fatalf("upgrade failed: %v", resp.Diagnostics)
	}

	var out ZeroTrustAccessAIControlsMcpPortalModel
	if diags := resp.State.Get(ctx, &out); diags.HasError() {
		t.Fatalf("decode upgraded state: %v", diags)
	}

	if out.ID.ValueString() != "mcp-sandbox" || out.Name.ValueString() != "MCP Gateway (Sandbox)" {
		t.Errorf("scalar fields not preserved: id=%q name=%q", out.ID.ValueString(), out.Name.ValueString())
	}
	if out.Servers.IsNull() || out.Servers.IsUnknown() {
		t.Fatal("upgraded servers should be a known set")
	}

	servers, d := out.Servers.AsStructSliceT(ctx)
	if d.HasError() {
		t.Fatalf("AsStructSliceT: %v", d)
	}
	ids := make([]string, 0, len(servers))
	for _, s := range servers {
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

	// A Set keys elements by the whole object, so the list -> set migration must
	// carry the nested scalar overrides through, not just server_id. Assert the
	// per-element values survive the framework decode/encode round-trip.
	byID := make(map[string]ZeroTrustAccessAIControlsMcpPortalServersModel, len(servers))
	for _, s := range servers {
		byID[s.ServerID.ValueString()] = s
	}
	if a := byID["alpha"]; !a.DefaultDisabled.ValueBool() || !a.OnBehalf.ValueBool() {
		t.Errorf("alpha scalars not preserved: default_disabled=%v on_behalf=%v", a.DefaultDisabled.ValueBool(), a.OnBehalf.ValueBool())
	}
	if b := byID["beta"]; b.DefaultDisabled.ValueBool() || b.OnBehalf.ValueBool() {
		t.Errorf("beta scalars not preserved: default_disabled=%v on_behalf=%v", b.DefaultDisabled.ValueBool(), b.OnBehalf.ValueBool())
	}
}

// TestUpgradeMcpPortalV500ToV501_FrameworkDecodeNullAndEmpty covers the
// no-attached-servers cases through the real upgrade path. A portal with a null
// (or empty) servers list is the common case; most portals start with none, so
// the list -> set conversion must round-trip those through req.State.Get /
// resp.State.Set without error, rather than only being exercised for the
// populated case.
func TestUpgradeMcpPortalV500ToV501_FrameworkDecodeNullAndEmpty(t *testing.T) {
	ctx := context.Background()

	cases := []struct {
		name     string
		servers  customfield.NestedObjectList[ZeroTrustAccessAIControlsMcpPortalServersModel]
		wantNull bool
	}{
		{
			name:     "null servers",
			servers:  customfield.NullObjectList[ZeroTrustAccessAIControlsMcpPortalServersModel](ctx),
			wantNull: true,
		},
		{
			name:    "empty servers",
			servers: customfield.NewObjectListMust(ctx, []ZeroTrustAccessAIControlsMcpPortalServersModel{}),
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			priorSchema := resourceSchemaV500(ctx)
			priorState := tfsdk.State{Schema: priorSchema}
			prior := zeroTrustAccessAIControlsMcpPortalModelV500{
				ID:        types.StringValue("p"),
				AccountID: types.StringValue("acct"),
				Hostname:  types.StringValue("mcp.example.com"),
				Name:      types.StringValue("no servers"),
				Servers:   tc.servers,
			}
			if diags := priorState.Set(ctx, prior); diags.HasError() {
				t.Fatalf("encode prior state: %v", diags)
			}

			r := &ZeroTrustAccessAIControlsMcpPortalResource{}
			upgrader, ok := r.UpgradeState(ctx)[500]
			if !ok {
				t.Fatal("no upgrader registered for prior version 500")
			}

			req := resource.UpgradeStateRequest{State: &priorState}
			resp := resource.UpgradeStateResponse{State: tfsdk.State{Schema: ResourceSchema(ctx)}}
			upgrader.StateUpgrader(ctx, req, &resp)
			if resp.Diagnostics.HasError() {
				t.Fatalf("upgrade failed: %v", resp.Diagnostics)
			}

			var out ZeroTrustAccessAIControlsMcpPortalModel
			if diags := resp.State.Get(ctx, &out); diags.HasError() {
				t.Fatalf("decode upgraded state: %v", diags)
			}

			if tc.wantNull {
				if !out.Servers.IsNull() {
					t.Errorf("expected servers to remain null after upgrade")
				}
				return
			}

			// Empty (but not null) must upgrade to a known, empty set.
			if out.Servers.IsNull() || out.Servers.IsUnknown() {
				t.Fatalf("expected a known servers set, got null=%v unknown=%v", out.Servers.IsNull(), out.Servers.IsUnknown())
			}
			servers, d := out.Servers.AsStructSliceT(ctx)
			if d.HasError() {
				t.Fatalf("AsStructSliceT: %v", d)
			}
			if len(servers) != 0 {
				t.Errorf("got %d servers, want 0", len(servers))
			}
		})
	}
}

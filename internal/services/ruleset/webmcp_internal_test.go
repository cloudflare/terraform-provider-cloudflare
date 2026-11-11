package ruleset

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijsoncustom"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
)

func rulesetWithActionParameters(ctx context.Context, t *testing.T, params RulesetRulesActionParametersModel) RulesetModel {
	t.Helper()
	rules, diags := customfield.NewObjectList(ctx, []RulesetRulesModel{{
		Action:           types.StringValue("set_config"),
		Expression:       types.StringValue("true"),
		ActionParameters: customfield.NewObjectMust(ctx, &params),
	}})
	if diags.HasError() {
		t.Fatalf("building rules: %v", diags)
	}
	return RulesetModel{
		Kind:  types.StringValue("zone"),
		Name:  types.StringValue("test"),
		Phase: types.StringValue("http_config_settings"),
		Rules: rules,
	}
}

// An explicit empty pack list must reach the API, since it differs from omitting the field.
func TestWebMCPActionParametersEncoding(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	for _, tt := range []struct {
		name        string
		enabled     types.Bool
		packs       customfield.List[types.String]
		wantEnabled string
		wantPacks   string
	}{
		{
			name:    "omitted",
			enabled: types.BoolNull(),
			packs:   customfield.NullList[types.String](ctx),
		},
		{
			name:        "disabled with empty packs",
			enabled:     types.BoolValue(false),
			packs:       customfield.NewListMust[types.String](ctx, []attr.Value{}),
			wantEnabled: "false",
			wantPacks:   "[]",
		},
		{
			name:    "enabled with packs",
			enabled: types.BoolValue(true),
			packs: customfield.NewListMust[types.String](ctx, []attr.Value{
				types.StringValue("c2pa"),
				types.StringValue("mcp-server-client"),
			}),
			wantEnabled: "true",
			wantPacks:   `["c2pa","mcp-server-client"]`,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			model := rulesetWithActionParameters(ctx, t, RulesetRulesActionParametersModel{
				WebMCPEnabled: tt.enabled,
				WebMCPPacks:   tt.packs,
			})

			body, err := model.MarshalJSON()
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}

			params := gjson.GetBytes(body, "rules.0.action_parameters")
			if got := params.Get("webmcp_enabled").Raw; got != tt.wantEnabled {
				t.Errorf("webmcp_enabled = %q, want %q (body: %s)", got, tt.wantEnabled, body)
			}
			if got := params.Get("webmcp_packs").Raw; got != tt.wantPacks {
				t.Errorf("webmcp_packs = %q, want %q (body: %s)", got, tt.wantPacks, body)
			}
		})
	}
}

func TestWebMCPActionParametersDecoding(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	for _, tt := range []struct {
		name      string
		params    string
		wantNull  bool
		wantPacks []string
	}{
		{name: "omitted", params: `{}`, wantNull: true},
		{name: "empty", params: `{"webmcp_enabled":false,"webmcp_packs":[]}`, wantPacks: []string{}},
		{name: "selected", params: `{"webmcp_enabled":true,"webmcp_packs":["c2pa","mcp-server-client"]}`, wantPacks: []string{"c2pa", "mcp-server-client"}},
	} {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			raw := []byte(`{"result":{"id":"abc","kind":"zone","name":"test","phase":"http_config_settings","rules":[{"action":"set_config","expression":"true","action_parameters":` + tt.params + `}]}}`)

			var env RulesetResultEnvelope
			if err := apijsoncustom.Unmarshal(raw, &env); err != nil {
				t.Fatalf("unmarshal: %v", err)
			}

			rules, diags := env.Result.Rules.AsStructSliceT(ctx)
			if diags.HasError() || len(rules) != 1 {
				t.Fatalf("rules: %v %v", rules, diags)
			}
			params, diags := rules[0].ActionParameters.Value(ctx)
			if diags.HasError() || params == nil {
				t.Fatalf("action_parameters: %v", diags)
			}

			if tt.wantNull {
				if !params.WebMCPPacks.IsNull() {
					t.Errorf("webmcp_packs = %v, want null", params.WebMCPPacks)
				}
				if !params.WebMCPEnabled.IsNull() {
					t.Errorf("webmcp_enabled = %v, want null", params.WebMCPEnabled)
				}
				return
			}

			if params.WebMCPPacks.IsNull() {
				t.Fatalf("webmcp_packs is null, want %v", tt.wantPacks)
			}
			got, diags := params.WebMCPPacks.Value(ctx)
			if diags.HasError() {
				t.Fatalf("packs value: %v", diags)
			}
			if len(got) != len(tt.wantPacks) {
				t.Fatalf("webmcp_packs = %v, want %v", got, tt.wantPacks)
			}
			for i, want := range tt.wantPacks {
				if got[i].ValueString() != want {
					t.Errorf("webmcp_packs[%d] = %q, want %q", i, got[i].ValueString(), want)
				}
			}
		})
	}
}

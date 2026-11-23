// This file is intentionally NOT code-generated.

package snippet_rules

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// The rules on Cloudflare, as the endpoint answers them: the result is the list.
const liveRules = `{"success":true,"errors":[],"messages":[],"result":[` +
	`{"id":"rule-1","snippet_name":"redirect_old","expression":"http.request.uri.path eq \"/old\"","description":"Old path","enabled":true,"last_updated":"2026-01-02T00:00:00Z"},` +
	`{"id":"rule-2","snippet_name":"add_header","expression":"http.host eq \"example.com\"","description":"","enabled":false,"last_updated":"2026-01-03T00:00:00Z"}` +
	`]}`

func rulesServer(t *testing.T) *SnippetRulesResource {
	t.Helper()
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if req.Method == http.MethodGet && req.URL.Path == "/zones/test-zone/snippets/snippet_rules" {
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, liveRules)
			return
		}
		t.Errorf("unexpected request: %s %s", req.Method, req.URL.Path)
		w.WriteHeader(http.StatusNotFound)
	}))
	t.Cleanup(server.Close)
	return &SnippetRulesResource{client: cloudflare.NewClient(
		option.WithBaseURL(server.URL),
		option.WithAPIToken("test-token"),
		option.WithMaxRetries(0),
	)}
}

func assertLiveRules(t *testing.T, state tfsdk.State) {
	t.Helper()
	var got SnippetRulesModel
	if diags := state.Get(context.Background(), &got); diags.HasError() {
		t.Fatalf("decode state: %v", diags)
	}
	if got.Rules == nil || len(*got.Rules) != 2 {
		t.Fatalf("rules = %v, want the two rules on Cloudflare", got.Rules)
	}
	first, second := (*got.Rules)[0], (*got.Rules)[1]
	if !first.SnippetName.Equal(types.StringValue("redirect_old")) ||
		!first.Expression.Equal(types.StringValue(`http.request.uri.path eq "/old"`)) ||
		!first.Enabled.Equal(types.BoolValue(true)) {
		t.Errorf("rule 1 = %+v, want redirect_old as on Cloudflare", first)
	}
	if !second.SnippetName.Equal(types.StringValue("add_header")) || !second.Enabled.Equal(types.BoolValue(false)) {
		t.Errorf("rule 2 = %+v, want add_header, disabled, as on Cloudflare", second)
	}
}

// A rule changed outside Terraform must reach the state on refresh, or a plan never
// shows the drift.
func TestReadRefreshesRulesChangedOutsideTerraform(t *testing.T) {
	t.Parallel()

	r := rulesServer(t)
	ctx := context.Background()
	state := tfsdk.State{Schema: ResourceSchema(ctx)}
	stale := []*SnippetRulesRulesModel{{
		ID:          types.StringValue("rule-1"),
		SnippetName: types.StringValue("redirect_old"),
		Expression:  types.StringValue(`http.request.uri.path eq "/before"`),
		Description: types.StringValue("Old path"),
		Enabled:     types.BoolValue(true),
	}}
	if diags := state.Set(ctx, &SnippetRulesModel{ID: types.StringValue("test-zone"), ZoneID: types.StringValue("test-zone"), Rules: &stale}); diags.HasError() {
		t.Fatalf("seed state: %v", diags)
	}

	readResp := resource.ReadResponse{State: state}
	r.Read(ctx, resource.ReadRequest{State: state}, &readResp)
	if readResp.Diagnostics.HasError() {
		t.Fatalf("read failed: %v", readResp.Diagnostics)
	}
	assertLiveRules(t, readResp.State)
}

// Import must bring the rules into the state, or the first plan after it proposes to
// write rules that are already there.
func TestImportReadsTheRules(t *testing.T) {
	t.Parallel()

	r := rulesServer(t)
	ctx := context.Background()
	importResp := resource.ImportStateResponse{State: tfsdk.State{Schema: ResourceSchema(ctx)}}
	r.ImportState(ctx, resource.ImportStateRequest{ID: "test-zone"}, &importResp)
	if importResp.Diagnostics.HasError() {
		t.Fatalf("import failed: %v", importResp.Diagnostics)
	}
	assertLiveRules(t, importResp.State)
}

package zero_trust_access_ai_controls_mcp_portal

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
)

// TestInjectServerIDs verifies the Read/Create/Update/Import fix that copies each
// servers[].id into servers[].server_id in the raw response before unmarshal. The
// portal API returns identity under "id"; the schema/write path use "server_id".
// Injecting the key lets apijson populate server_id natively so the servers set is
// keyed correctly rather than collapsing elements that differ only by a null
// server_id.
func TestInjectServerIDs(t *testing.T) {
	serverIDs := func(body []byte) []string {
		var env struct {
			Result struct {
				Servers []map[string]any `json:"servers"`
			} `json:"result"`
		}
		if err := json.Unmarshal(body, &env); err != nil {
			t.Fatalf("unmarshal result: %v", err)
		}
		out := make([]string, 0, len(env.Result.Servers))
		for _, s := range env.Result.Servers {
			if v, ok := s["server_id"].(string); ok {
				out = append(out, v)
			} else {
				out = append(out, "<none>")
			}
		}
		return out
	}

	t.Run("injects server_id from id, preserves order", func(t *testing.T) {
		in := []byte(`{"result":{"id":"p","servers":[{"id":"alpha","name":"A"},{"id":"beta","name":"B"}]}}`)
		got := serverIDs(injectServerIDs(in))
		want := []string{"alpha", "beta"}
		for i := range want {
			if got[i] != want[i] {
				t.Errorf("servers[%d].server_id = %q, want %q", i, got[i], want[i])
			}
		}
	})

	t.Run("does not clobber an existing server_id", func(t *testing.T) {
		in := []byte(`{"result":{"id":"p","servers":[{"id":"alpha","server_id":"explicit"}]}}`)
		got := serverIDs(injectServerIDs(in))
		if got[0] != "explicit" {
			t.Errorf("server_id = %q, want %q (must not overwrite)", got[0], "explicit")
		}
	})

	t.Run("no servers key leaves body unchanged", func(t *testing.T) {
		in := []byte(`{"result":{"id":"p","name":"x"}}`)
		if string(injectServerIDs(in)) != string(in) {
			t.Errorf("expected unchanged body")
		}
	})

	t.Run("unparseable body left unchanged", func(t *testing.T) {
		in := []byte(`not json`)
		if string(injectServerIDs(in)) != string(in) {
			t.Errorf("expected unchanged body on parse error")
		}
	})

	t.Run("server without id left without server_id", func(t *testing.T) {
		in := []byte(`{"result":{"servers":[{"name":"A"}]}}`)
		got := serverIDs(injectServerIDs(in))
		if got[0] != "<none>" {
			t.Errorf("expected no server_id injected when id absent, got %q", got[0])
		}
	})

	t.Run("explicit null id is treated as absent", func(t *testing.T) {
		in := []byte(`{"result":{"servers":[{"id":null,"name":"A"}]}}`)
		got := serverIDs(injectServerIDs(in))
		if got[0] != "<none>" {
			t.Errorf("expected no server_id injected when id is null, got %q", got[0])
		}
	})
}

// TestPortalReadRoundTripServersKnown proves that a representative portal Read
// response decodes into a fully-known servers Set after injectServerIDs: every
// element has a known server_id (mapped from the response "id") and known
// default_disabled / on_behalf (returned by the API). This is the property the
// Set model relies on — set membership is hashed by the whole element, so an
// unknown/absent nested value would produce churn ("known after apply" or
// remove/add) instead of an empty plan. The backing API always serializes both
// booleans, so the decoded element is fully known here.
func TestPortalReadRoundTripServersKnown(t *testing.T) {
	ctx := context.Background()

	body := []byte(`{"result":{` +
		`"id":"mcp-sandbox","account_id":"acct","name":"MCP Gateway (Sandbox)","hostname":"mcp.example.com",` +
		`"servers":[` +
		`{"id":"alpha","default_disabled":true,"on_behalf":true},` +
		`{"id":"beta","default_disabled":false,"on_behalf":false}` +
		`]}}`)

	body = injectServerIDs(body)

	var env ZeroTrustAccessAIControlsMcpPortalResultEnvelope
	if err := apijson.Unmarshal(body, &env); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}

	if env.Result.Servers.IsNull() || env.Result.Servers.IsUnknown() {
		t.Fatal("servers set should be known after decode")
	}

	servers, d := env.Result.Servers.AsStructSliceT(ctx)
	if d.HasError() {
		t.Fatalf("AsStructSliceT: %v", d)
	}
	if len(servers) != 2 {
		t.Fatalf("got %d servers, want 2", len(servers))
	}

	for i, s := range servers {
		if s.ServerID.IsNull() || s.ServerID.IsUnknown() {
			t.Errorf("servers[%d].server_id not known: %#v", i, s.ServerID)
		}
		if s.DefaultDisabled.IsNull() || s.DefaultDisabled.IsUnknown() {
			t.Errorf("servers[%d].default_disabled not known: %#v", i, s.DefaultDisabled)
		}
		if s.OnBehalf.IsNull() || s.OnBehalf.IsUnknown() {
			t.Errorf("servers[%d].on_behalf not known: %#v", i, s.OnBehalf)
		}
	}
}

package zero_trust_access_ai_controls_mcp_portal

import (
	"encoding/json"
	"testing"
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
}

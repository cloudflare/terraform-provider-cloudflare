package zero_trust_access_ai_controls_mcp_portal

import (
	"context"
	"sort"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
)

func TestMcpPortalReadPopulatesServerIDs(t *testing.T) {
	t.Parallel()

	apiResponse := []byte(`{
		"result": {
			"id": "portal",
			"hostname": "portal.example.com",
			"name": "Portal",
			"servers": [
				{
					"server_id": "alpha",
					"default_disabled": false,
					"on_behalf": true,
					"updated_prompts": null,
					"updated_tools": null
				},
				{
					"server_id": "beta",
					"default_disabled": false,
					"on_behalf": true,
					"updated_prompts": null,
					"updated_tools": null
				}
			]
		}
	}`)

	var envelope ZeroTrustAccessAIControlsMcpPortalResultEnvelope
	if err := apijson.Unmarshal(apiResponse, &envelope); err != nil {
		t.Fatalf("failed to unmarshal API response: %v", err)
	}

	servers, diags := envelope.Result.Servers.AsStructSliceT(context.Background())
	if diags.HasError() {
		t.Fatalf("failed to decode servers: %v", diags)
	}
	if len(servers) != 2 {
		t.Fatalf("got %d servers, want 2", len(servers))
	}

	serverIDs := make([]string, 0, len(servers))
	for _, server := range servers {
		serverIDs = append(serverIDs, server.ServerID.ValueString())
	}
	sort.Strings(serverIDs)

	want := []string{"alpha", "beta"}
	for i := range want {
		if serverIDs[i] != want[i] {
			t.Errorf("server_id[%d] = %q, want %q", i, serverIDs[i], want[i])
		}
	}
}

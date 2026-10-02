package zero_trust_access_ai_controls_mcp_server

import (
	"encoding/json"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
)

// marshalMcpServerUpdate uses field-level diffing because this PUT endpoint
// applies partial-update semantics. Credential omission is significant:
// client_secret presence rotates the secret, while omission preserves it.
func marshalMcpServerUpdate(plan, state ZeroTrustAccessAIControlsMcpServerModel) ([]byte, error) {
	body, err := apijson.MarshalForPatch(plan, state)
	if err != nil {
		return nil, err
	}

	var fields map[string]json.RawMessage
	if err := json.Unmarshal(body, &fields); err != nil {
		return nil, err
	}
	if plan.ClientSecret.IsNull() || plan.ClientSecret.IsUnknown() {
		delete(fields, "client_secret")
	}
	if plan.AuthCredentials.IsNull() || plan.AuthCredentials.IsUnknown() {
		delete(fields, "auth_credentials")
	}

	return json.Marshal(fields)
}

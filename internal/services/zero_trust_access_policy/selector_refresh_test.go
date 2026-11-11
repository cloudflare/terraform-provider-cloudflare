package zero_trust_access_policy

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stretchr/testify/require"
)

func accessPolicyState(t *testing.T, ctx context.Context, field, rules string) tfsdk.State {
	t.Helper()
	var model ZeroTrustAccessPolicyModel
	require.NoError(t, apijson.Unmarshal([]byte(`{"id":"policy-id","name":"Policy","decision":"allow","`+field+`":`+rules+`}`), &model))
	model.AccountID = types.StringValue("account-id")
	state := tfsdk.State{Schema: ResourceSchema(ctx)}
	diags := state.Set(ctx, &model)
	require.False(t, diags.HasError(), "%v", diags)
	return state
}

func accessPolicyRules(t *testing.T, ctx context.Context, state tfsdk.State, field string) attr.Value {
	t.Helper()
	var model ZeroTrustAccessPolicyModel
	diags := state.Get(ctx, &model)
	require.False(t, diags.HasError(), "%v", diags)
	switch field {
	case "include":
		return model.Include
	case "exclude":
		return model.Exclude
	default:
		return model.Require
	}
}

// The API returns rules in its own order, which can differ from the order of
// the set in state. Refresh must not carry a selector from one rule into
// another, either as an empty phantom object or as a merged selector.
func TestAccessPolicyRefreshDoesNotMergeSelectorsAcrossRules(t *testing.T) {
	tests := []struct {
		name  string
		state string
		api   string
	}{
		{
			name:  "ip and device_posture",
			state: `[{"ip":{"ip":"192.0.2.10/32"}},{"device_posture":{"integration_uid":"posture-id"}}]`,
			api:   `[{"device_posture":{"integration_uid":"posture-id"}},{"ip":{"ip":"192.0.2.10/32"}}]`,
		},
		{
			name:  "email and oidc",
			state: `[{"email":{"email":"user@example.com"}},{"oidc":{"claim_name":"groups","claim_value":"admins","identity_provider_id":"idp-id"}}]`,
			api:   `[{"oidc":{"claim_name":"groups","claim_value":"admins","identity_provider_id":"idp-id"}},{"email":{"email":"user@example.com"}}]`,
		},
		{
			name:  "group and service token",
			state: `[{"group":{"id":"group-id"}},{"any_valid_service_token":{}}]`,
			api:   `[{"any_valid_service_token":{}},{"group":{"id":"group-id"}}]`,
		},
		{
			name:  "state holds a phantom selector",
			state: `[{"ip":{"ip":"192.0.2.10/32"},"device_posture":{"integration_uid":null}}]`,
			api:   `[{"ip":{"ip":"192.0.2.10/32"}}]`,
		},
	}
	for _, field := range []string{"include", "exclude", "require"} {
		for _, tt := range tests {
			t.Run(field+"/"+tt.name, func(t *testing.T) {
				ctx := context.Background()
				server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
					if req.Method != http.MethodGet || req.URL.Path != "/accounts/account-id/access/policies/policy-id" {
						t.Errorf("unexpected request: %s %s", req.Method, req.URL)
						http.Error(w, "unexpected request", http.StatusBadRequest)
						return
					}
					w.Header().Set("Content-Type", "application/json")
					_ = json.NewEncoder(w).Encode(map[string]any{
						"success": true,
						"result": map[string]any{
							"id": "policy-id", "name": "Policy", "decision": "allow",
							"created_at": "2026-01-01T00:00:00Z", "updated_at": "2026-01-01T00:00:00Z",
							"session_duration": "24h", "reusable": true, "app_count": 0,
							field: json.RawMessage(tt.api),
						},
					})
				}))
				defer server.Close()
				r := &ZeroTrustAccessPolicyResource{client: cloudflare.NewClient(option.WithBaseURL(server.URL), option.WithAPIToken("test-token"))}
				state := accessPolicyState(t, ctx, field, tt.state)
				response := &resource.ReadResponse{State: state}
				r.Read(ctx, resource.ReadRequest{State: state}, response)
				require.False(t, response.Diagnostics.HasError(), "%v", response.Diagnostics)
				expected := accessPolicyRules(t, ctx, accessPolicyState(t, ctx, field, tt.api), field)
				actual := accessPolicyRules(t, ctx, response.State, field)
				require.True(t, expected.Equal(actual), "expected %s, got %s", expected, actual)
			})
		}
	}
}

func TestPruneEmptyDevicePostureAndOIDCSelectors(t *testing.T) {
	phantoms := `[` +
		`{"ip":{"ip":"192.0.2.10/32"},"device_posture":{"integration_uid":null,"account_id":null}},` +
		`{"email":{"email":"user@example.com"},"oidc":{"claim_name":null,"claim_value":null,"identity_provider_id":null}},` +
		`{"device_posture":{"integration_uid":"posture-id"}},` +
		`{"oidc":{"claim_name":"groups","claim_value":"admins","identity_provider_id":"idp-id"}}]`
	clean := `[` +
		`{"ip":{"ip":"192.0.2.10/32"}},` +
		`{"email":{"email":"user@example.com"}},` +
		`{"device_posture":{"integration_uid":"posture-id"}},` +
		`{"oidc":{"claim_name":"groups","claim_value":"admins","identity_provider_id":"idp-id"}}]`
	for _, field := range []string{"include", "exclude", "require"} {
		t.Run(field, func(t *testing.T) {
			ctx := context.Background()
			var data, expected ZeroTrustAccessPolicyModel
			require.NoError(t, apijson.Unmarshal([]byte(`{"`+field+`":`+phantoms+`}`), &data))
			require.NoError(t, apijson.Unmarshal([]byte(`{"`+field+`":`+clean+`}`), &expected))
			diags := pruneEmptyConditionSelectors(ctx, &data)
			require.False(t, diags.HasError(), "%v", diags)
			require.True(t, expected.Include.Equal(data.Include), "include: expected %s, got %s", expected.Include, data.Include)
			require.True(t, expected.Exclude.Equal(data.Exclude), "exclude: expected %s, got %s", expected.Exclude, data.Exclude)
			require.True(t, expected.Require.Equal(data.Require), "require: expected %s, got %s", expected.Require, data.Require)
		})
	}
}

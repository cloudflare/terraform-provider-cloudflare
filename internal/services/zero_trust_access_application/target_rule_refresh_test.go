package zero_trust_access_application

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stretchr/testify/require"
)

func targetCriteriaState(t *testing.T, ctx context.Context, criteria string) tfsdk.State {
	t.Helper()
	var model ZeroTrustAccessApplicationModel
	require.NoError(t, apijson.Unmarshal([]byte(`{"id":"app-id","type":"infrastructure","name":"SSH","target_criteria":`+criteria+`}`), &model))
	model.AccountID = types.StringValue("account-id")
	state := tfsdk.State{Schema: ResourceSchema(ctx)}
	diags := state.Set(ctx, &model)
	require.False(t, diags.HasError(), "%v", diags)
	return state
}

func assertTargetCriteriaState(t *testing.T, ctx context.Context, expected, actual tfsdk.State) {
	t.Helper()
	var expectedCriteria, actualCriteria types.List
	require.False(t, expected.GetAttribute(ctx, path.Root("target_criteria"), &expectedCriteria).HasError())
	require.False(t, actual.GetAttribute(ctx, path.Root("target_criteria"), &actualCriteria).HasError())
	require.True(t, expectedCriteria.Equal(actualCriteria), "expected %s, got %s", expectedCriteria, actualCriteria)
}

func TestTargetCriteriaRefreshAndImport(t *testing.T) {
	allRules := `[{"port":22,"protocol":"SSH","include":{"target_attributes":{"hostname":["web-*"]},"tags":{"environment":["production","staging"]}},"require":{"tags":{"team":["engineering"]}},"exclude":{"target_attributes":{"hostname":["test-*"]}}}]`
	tagsOnly := `[{"port":22,"protocol":"SSH","include":{"tags":{"environment":["production"]}}}]`
	hostnamesOnly := `[{"port":22,"protocol":"SSH","include":{"target_attributes":{"hostname":["web-*"]}}}]`
	legacy := `[{"port":22,"protocol":"SSH","target_attributes":{"hostname":["server"]}}]`
	grandfathered := `[{"port":22,"protocol":"SSH","target_attributes":{"hostname":["server"],"environment":["production","staging"]}}]`
	tests := []struct {
		name   string
		before string
		after  string
	}{
		{name: "unchanged rules", before: allRules, after: allRules},
		{name: "unchanged tags only", before: tagsOnly, after: tagsOnly},
		{name: "unchanged hostnames only", before: hostnamesOnly, after: hostnamesOnly},
		{name: "removed rules and hostname selector", before: allRules, after: tagsOnly},
		{name: "removed tags selector", before: allRules, after: hostnamesOnly},
		{name: "rules replaced with legacy criteria", before: allRules, after: legacy},
		{name: "legacy replaced with rules", before: legacy, after: tagsOnly},
		{name: "grandfathered legacy map", before: grandfathered, after: grandfathered},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
				if req.Method != http.MethodGet || req.URL.Path != "/accounts/account-id/access/apps/app-id" {
					t.Errorf("unexpected request: %s %s", req.Method, req.URL)
					http.Error(w, "unexpected request", http.StatusBadRequest)
					return
				}
				w.Header().Set("Content-Type", "application/json")
				_ = json.NewEncoder(w).Encode(map[string]any{
					"success": true,
					"result": map[string]any{
						"id": "app-id", "type": "infrastructure", "name": "SSH",
						"target_criteria": json.RawMessage(tt.after),
					},
				})
			}))
			defer server.Close()
			r := &ZeroTrustAccessApplicationResource{client: cloudflare.NewClient(option.WithBaseURL(server.URL), option.WithAPIToken("test-token"))}
			state := targetCriteriaState(t, ctx, tt.before)
			expected := targetCriteriaState(t, ctx, tt.after)
			// Both the first refresh and subsequent refreshes must represent
			// omitted selectors as null and preserve the API's exact values.
			for i := 0; i < 2; i++ {
				response := &resource.ReadResponse{State: state}
				r.Read(ctx, resource.ReadRequest{State: state}, response)
				require.False(t, response.Diagnostics.HasError(), "%v", response.Diagnostics)
				assertTargetCriteriaState(t, ctx, expected, response.State)
				state = response.State
			}
			imported := &resource.ImportStateResponse{State: tfsdk.State{Schema: ResourceSchema(ctx)}}
			r.ImportState(ctx, resource.ImportStateRequest{ID: "accounts/account-id/app-id"}, imported)
			require.False(t, imported.Diagnostics.HasError(), "%v", imported.Diagnostics)
			assertTargetCriteriaState(t, ctx, expected, imported.State)
		})
	}
}

func TestTargetCriteriaLegacyAttributesPreservedOnUpdate(t *testing.T) {
	body := `{"name":"before","target_criteria":[{"port":22,"protocol":"SSH","target_attributes":{"hostname":["server"],"environment":["production","staging"]}}]}`
	var state ZeroTrustAccessApplicationModel
	require.NoError(t, apijson.Unmarshal([]byte(body), &state))
	plan := state
	plan.Name = types.StringValue("after")
	encoded, err := plan.MarshalJSONForUpdate(state)
	require.NoError(t, err)
	var expected, actual map[string]json.RawMessage
	require.NoError(t, json.Unmarshal([]byte(body), &expected))
	require.NoError(t, json.Unmarshal(encoded, &actual))
	require.JSONEq(t, string(expected["target_criteria"]), string(actual["target_criteria"]))
	require.JSONEq(t, `"after"`, string(actual["name"]))
}

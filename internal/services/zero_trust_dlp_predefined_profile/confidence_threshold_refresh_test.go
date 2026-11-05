package zero_trust_dlp_predefined_profile

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

func TestConfidenceThresholdRefreshWhenAPIOmitsField(t *testing.T) {
	ctx := context.Background()
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodGet || req.URL.Path != "/accounts/account-id/dlp/profiles/predefined/profile-id/config" {
			t.Errorf("unexpected request: %s %s", req.Method, req.URL)
			http.Error(w, "unexpected request", http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]any{
			"success": true,
			"result": map[string]any{
				"id":              "profile-id",
				"profile_id":      "profile-id",
				"name":            "Predefined",
				"enabled_entries": []string{},
			},
		})
	}))
	defer server.Close()

	var model ZeroTrustDLPPredefinedProfileModel
	require.NoError(t, apijson.Unmarshal([]byte(`{"confidence_threshold":"low"}`), &model))
	model.AccountID = types.StringValue("account-id")
	model.ProfileID = types.StringValue("profile-id")
	state := tfsdk.State{Schema: ResourceSchema(ctx)}
	require.False(t, state.Set(ctx, &model).HasError())

	r := &ZeroTrustDLPPredefinedProfileResource{client: cloudflare.NewClient(option.WithBaseURL(server.URL), option.WithAPIToken("test-token"))}
	response := &resource.ReadResponse{State: state}
	r.Read(ctx, resource.ReadRequest{State: state}, response)
	require.False(t, response.Diagnostics.HasError(), "%v", response.Diagnostics)

	var threshold types.String
	require.False(t, response.State.GetAttribute(ctx, path.Root("confidence_threshold"), &threshold).HasError())
	require.Equal(t, "low", threshold.ValueString(), "confidence_threshold drifted on refresh: expected %q, got %q", "low", threshold.ValueString())
}

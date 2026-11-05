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
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stretchr/testify/require"
)

const allNullLandingPageDesign = `{"button_color":null,"button_text_color":null,"image_url":null,"message":null,"title":null}`

func appLauncherState(t *testing.T, ctx context.Context, landingPageDesign string) tfsdk.State {
	t.Helper()
	body := `{"id":"app-id","type":"app_launcher","name":"App Launcher"`
	if landingPageDesign != "" {
		body += `,"landing_page_design":` + landingPageDesign
	}
	body += `}`
	var model ZeroTrustAccessApplicationModel
	require.NoError(t, apijson.Unmarshal([]byte(body), &model))
	model.AccountID = types.StringValue("account-id")
	state := tfsdk.State{Schema: ResourceSchema(ctx)}
	diags := state.Set(ctx, &model)
	require.False(t, diags.HasError(), "%v", diags)
	return state
}

func landingPageDesignServer(t *testing.T, landingPageDesign string) *httptest.Server {
	t.Helper()
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodGet || req.URL.Path != "/accounts/account-id/access/apps/app-id" {
			t.Errorf("unexpected request: %s %s", req.Method, req.URL)
			http.Error(w, "unexpected request", http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]any{
			"success": true,
			"result": map[string]any{
				"id": "app-id", "type": "app_launcher", "name": "App Launcher",
				"landing_page_design": json.RawMessage(landingPageDesign),
			},
		})
	}))
}

func landingPageDesignAttribute(t *testing.T, ctx context.Context, state tfsdk.State) customfield.NestedObject[ZeroTrustAccessApplicationLandingPageDesignModel] {
	t.Helper()
	var value customfield.NestedObject[ZeroTrustAccessApplicationLandingPageDesignModel]
	diags := state.GetAttribute(ctx, path.Root("landing_page_design"), &value)
	require.False(t, diags.HasError(), "%v", diags)
	return value
}

// The API reads landing_page_design back as an object whose fields are all null
// when no design is set (for example on accounts without App Launcher
// customization). Refresh must store that as null so a configuration that omits
// the attribute does not plan a change on every run.
func TestLandingPageDesignAllNullRefreshesAsNull(t *testing.T) {
	tests := []struct {
		name   string
		before string
	}{
		{name: "state already null", before: ""},
		{name: "state holds the all-null object", before: allNullLandingPageDesign},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			server := landingPageDesignServer(t, allNullLandingPageDesign)
			defer server.Close()
			r := &ZeroTrustAccessApplicationResource{client: cloudflare.NewClient(option.WithBaseURL(server.URL), option.WithAPIToken("test-token"))}
			state := appLauncherState(t, ctx, tt.before)
			for i := 0; i < 2; i++ {
				response := &resource.ReadResponse{State: state}
				r.Read(ctx, resource.ReadRequest{State: state}, response)
				require.False(t, response.Diagnostics.HasError(), "%v", response.Diagnostics)
				require.True(t, landingPageDesignAttribute(t, ctx, response.State).IsNull(), "refresh %d: expected null landing_page_design", i+1)
				state = response.State
			}
		})
	}
}

func TestLandingPageDesignWithValuesIsPreserved(t *testing.T) {
	ctx := context.Background()
	design := `{"button_color":"#ff0000","button_text_color":"#ffffff","image_url":"https://example.com/logo.png","message":"Hello","title":"Welcome!"}`
	server := landingPageDesignServer(t, design)
	defer server.Close()
	r := &ZeroTrustAccessApplicationResource{client: cloudflare.NewClient(option.WithBaseURL(server.URL), option.WithAPIToken("test-token"))}
	state := appLauncherState(t, ctx, design)
	response := &resource.ReadResponse{State: state}
	r.Read(ctx, resource.ReadRequest{State: state}, response)
	require.False(t, response.Diagnostics.HasError(), "%v", response.Diagnostics)
	expected := landingPageDesignAttribute(t, ctx, appLauncherState(t, ctx, design))
	actual := landingPageDesignAttribute(t, ctx, response.State)
	require.False(t, actual.IsNull())
	require.True(t, expected.Equal(actual), "expected %s, got %s", expected, actual)
}

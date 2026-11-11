package zero_trust_access_service_token_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_access_service_token"
	"github.com/hashicorp/terraform-plugin-framework/action"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

const (
	testAccountID = "023e105f4ecef8ad9ca31a8372d0c353"
	testTokenID   = "f174e90a-fafe-4643-bbbc-4a0ed4fc8415"
)

// refreshResponse is what the API returns from the refresh endpoint: the
// token with its new expiry. It never includes the client secret.
const refreshResponse = `{
  "success": true,
  "errors": [],
  "messages": [],
  "result": {
    "id": "` + testTokenID + `",
    "client_id": "88bf3b6d86161464f6509f7219099e57.access",
    "name": "CI/CD token",
    "duration": "8760h",
    "expires_at": "2027-11-06T16:00:00Z",
    "created_at": "2026-11-06T16:00:00Z",
    "updated_at": "2026-11-06T16:00:00Z"
  }
}`

func newRefreshAction(t *testing.T, baseURL string) action.Action {
	t.Helper()
	a := zero_trust_access_service_token.NewRefreshAction()
	configurable, ok := a.(action.ActionWithConfigure)
	if !ok {
		t.Fatal("refresh action doesn't implement ActionWithConfigure")
	}
	resp := action.ConfigureResponse{}
	configurable.Configure(context.Background(), action.ConfigureRequest{
		ProviderData: cloudflare.NewClient(
			option.WithBaseURL(baseURL),
			option.WithAPIToken("test-token"),
			option.WithMaxRetries(0),
		),
	}, &resp)
	if resp.Diagnostics.HasError() {
		t.Fatalf("Configure failed: %v", resp.Diagnostics)
	}
	return a
}

func refreshConfig(t *testing.T, a action.Action) tfsdk.Config {
	t.Helper()
	ctx := context.Background()
	schemaResp := action.SchemaResponse{}
	a.Schema(ctx, action.SchemaRequest{}, &schemaResp)
	if schemaResp.Diagnostics.HasError() {
		t.Fatalf("Schema failed: %v", schemaResp.Diagnostics)
	}
	objType, ok := schemaResp.Schema.Type().TerraformType(ctx).(tftypes.Object)
	if !ok {
		t.Fatalf("action schema type is %T, expected tftypes.Object", schemaResp.Schema.Type().TerraformType(ctx))
	}
	return tfsdk.Config{
		Schema: schemaResp.Schema,
		Raw: tftypes.NewValue(objType, map[string]tftypes.Value{
			"account_id":       tftypes.NewValue(tftypes.String, testAccountID),
			"service_token_id": tftypes.NewValue(tftypes.String, testTokenID),
		}),
	}
}

func TestRefreshActionMetadata(t *testing.T) {
	resp := action.MetadataResponse{}
	zero_trust_access_service_token.NewRefreshAction().Metadata(context.Background(), action.MetadataRequest{ProviderTypeName: "cloudflare"}, &resp)
	if want := "cloudflare_zero_trust_access_service_token_refresh"; resp.TypeName != want {
		t.Errorf("TypeName = %q, want %q", resp.TypeName, want)
	}
}

func TestRefreshActionInvoke(t *testing.T) {
	var method, path string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		method, path = r.Method, r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(refreshResponse))
	}))
	defer srv.Close()

	a := newRefreshAction(t, srv.URL)
	var progress []string
	resp := action.InvokeResponse{
		SendProgress: func(e action.InvokeProgressEvent) { progress = append(progress, e.Message) },
	}
	a.Invoke(context.Background(), action.InvokeRequest{Config: refreshConfig(t, a)}, &resp)

	if resp.Diagnostics.HasError() {
		t.Fatalf("Invoke failed: %v", resp.Diagnostics)
	}
	if method != http.MethodPost {
		t.Errorf("method = %s, want POST", method)
	}
	if want := "/accounts/" + testAccountID + "/access/service_tokens/" + testTokenID + "/refresh"; path != want {
		t.Errorf("path = %s, want %s", path, want)
	}
	if len(progress) != 1 || !strings.Contains(progress[0], "2027-11-06T16:00:00Z") {
		t.Errorf("progress = %q, want one message with the new expiry", progress)
	}
}

func TestRefreshActionInvokeAPIError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"success":false,"errors":[{"code":12128,"message":"service token not found"}],"messages":[],"result":null}`))
	}))
	defer srv.Close()

	a := newRefreshAction(t, srv.URL)
	resp := action.InvokeResponse{}
	a.Invoke(context.Background(), action.InvokeRequest{Config: refreshConfig(t, a)}, &resp)

	if !resp.Diagnostics.HasError() {
		t.Fatal("expected an error diagnostic for a 404 from the API")
	}
}

func TestRefreshActionInvokeWithoutClient(t *testing.T) {
	a := zero_trust_access_service_token.NewRefreshAction()
	resp := action.InvokeResponse{}
	a.Invoke(context.Background(), action.InvokeRequest{Config: refreshConfig(t, a)}, &resp)

	if !resp.Diagnostics.HasError() {
		t.Fatal("expected an error diagnostic when the provider client is missing")
	}
}

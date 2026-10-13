package zero_trust_access_identity_provider

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

func TestImportStateLoginPageAutoPrompt(t *testing.T) {
	t.Parallel()

	for _, tc := range []struct {
		name   string
		config string
		want   bool
	}{
		{name: "omitted", config: `{}`, want: false},
		{name: "null", config: `{"login_page_auto_prompt":null}`, want: false},
		{name: "false", config: `{"login_page_auto_prompt":false}`, want: false},
		{name: "true", config: `{"login_page_auto_prompt":true}`, want: true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
				if req.Method != http.MethodGet || req.URL.Path != "/accounts/test-account/access/identity_providers/test-idp" {
					t.Errorf("unexpected request: %s %s", req.Method, req.URL.Path)
					w.WriteHeader(http.StatusNotFound)
					return
				}
				w.Header().Set("Content-Type", "application/json")
				fmt.Fprintf(w, `{"success":true,"errors":[],"messages":[],"result":{"id":"test-idp","name":"test","type":"github","config":%s}}`, tc.config)
			}))
			defer server.Close()

			r := &ZeroTrustAccessIdentityProviderResource{client: cloudflare.NewClient(
				option.WithBaseURL(server.URL),
				option.WithAPIToken("test-token"),
				option.WithMaxRetries(0),
			)}
			ctx := context.Background()
			resp := resource.ImportStateResponse{State: tfsdk.State{Schema: ResourceSchema(ctx)}}
			r.ImportState(ctx, resource.ImportStateRequest{ID: "accounts/test-account/test-idp"}, &resp)
			if resp.Diagnostics.HasError() {
				t.Fatalf("import failed: %v", resp.Diagnostics)
			}
			var imported ZeroTrustAccessIdentityProviderModel
			if diags := resp.State.Get(ctx, &imported); diags.HasError() {
				t.Fatalf("decode imported state: %v", diags)
			}
			if imported.Config == nil {
				t.Fatal("imported config is nil")
			}
			if !imported.Config.LoginPageAutoPrompt.Equal(types.BoolValue(tc.want)) {
				t.Errorf("login_page_auto_prompt = %v, want %v", imported.Config.LoginPageAutoPrompt, tc.want)
			}
		})
	}
}

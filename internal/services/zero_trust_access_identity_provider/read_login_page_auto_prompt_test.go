package zero_trust_access_identity_provider

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// State written before login_page_auto_prompt existed holds null. Access stores
// and returns the field only for passkeys, so Read must fill the default.
func TestReadLoginPageAutoPromptOmitted(t *testing.T) {
	t.Parallel()

	for _, tc := range []struct {
		name   string
		idp    string
		config string
		want   types.Bool
	}{
		{
			name:   "azure_ad_omitted",
			idp:    "azureAD",
			config: `{"azure_cloud":"default","client_id":"client-id","conditional_access_enabled":false,"directory_id":"directory-id","pkce_enabled":true,"redirect_url":"https://example.cloudflareaccess.com/cdn-cgi/access/callback","support_groups":true}`,
			want:   types.BoolValue(false),
		},
		{
			name:   "passkeys_true",
			idp:    "passkeys",
			config: `{"redirect_url":"https://example.cloudflareaccess.com/cdn-cgi/access/callback","login_page_auto_prompt":true}`,
			want:   types.BoolValue(true),
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
				if req.Method != http.MethodGet || req.URL.Path != "/accounts/test-account/access/identity_providers/test-idp" {
					t.Errorf("unexpected request: %s %s", req.Method, req.URL.Path)
					w.WriteHeader(http.StatusNotFound)
					return
				}
				w.Header().Set("Content-Type", "application/json")
				fmt.Fprintf(w, `{"success":true,"errors":[],"messages":[],"result":{"id":"test-idp","uid":"test-idp","type":%q,"name":"test","version":"1","config":%s,"scim_config":{"scim_base_url":"https://example.cloudflareaccess.com/scim/test-idp","enabled":false,"user_deprovision":false,"seat_deprovision":false,"group_member_deprovision":false}}}`, tc.idp, tc.config)
			}))
			defer server.Close()

			r := &ZeroTrustAccessIdentityProviderResource{client: cloudflare.NewClient(
				option.WithBaseURL(server.URL),
				option.WithAPIToken("test-token"),
				option.WithMaxRetries(0),
			)}
			ctx := context.Background()
			prior := tfsdk.State{Schema: ResourceSchema(ctx)}
			if diags := prior.Set(ctx, &ZeroTrustAccessIdentityProviderModel{
				ID:        types.StringValue("test-idp"),
				AccountID: types.StringValue("test-account"),
				Name:      types.StringValue("test"),
				Type:      types.StringValue(tc.idp),
				Config: &ZeroTrustAccessIdentityProviderConfigModel{
					LoginPageAutoPrompt: types.BoolNull(),
				},
				SCIMConfig:         customfield.NullObject[ZeroTrustAccessIdentityProviderSCIMConfigModel](ctx),
				SAMLCertificateSet: customfield.NullObject[ZeroTrustAccessIdentityProviderSAMLCertificateSetModel](ctx),
			}); diags.HasError() {
				t.Fatalf("set prior state: %v", diags)
			}

			resp := resource.ReadResponse{State: prior}
			r.Read(ctx, resource.ReadRequest{State: prior}, &resp)
			if resp.Diagnostics.HasError() {
				t.Fatalf("read failed: %v", resp.Diagnostics)
			}
			var refreshed ZeroTrustAccessIdentityProviderModel
			if diags := resp.State.Get(ctx, &refreshed); diags.HasError() {
				t.Fatalf("decode refreshed state: %v", diags)
			}
			if refreshed.Config == nil {
				t.Fatal("refreshed config is nil")
			}
			if !refreshed.Config.LoginPageAutoPrompt.Equal(tc.want) {
				t.Errorf("login_page_auto_prompt = %v, want %v", refreshed.Config.LoginPageAutoPrompt, tc.want)
			}
		})
	}
}

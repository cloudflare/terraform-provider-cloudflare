package internal_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/terraform-provider-cloudflare/internal"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

const testProviderVersion = "1.0"
const testTerraformVersion = "1.9.0"

// providerConfig builds a tfsdk.Config from the provider schema, with every
// attribute null unless overridden in attrs.
func providerConfig(t *testing.T, ctx context.Context, attrs map[string]tftypes.Value) tfsdk.Config {
	t.Helper()

	schema := internal.ProviderSchema(ctx)
	objType, ok := schema.Type().TerraformType(ctx).(tftypes.Object)
	if !ok {
		t.Fatalf("provider schema type is %T, expected tftypes.Object", schema.Type().TerraformType(ctx))
	}

	values := map[string]tftypes.Value{}
	for name, attrType := range objType.AttributeTypes {
		if v, ok := attrs[name]; ok {
			values[name] = v
			continue
		}
		values[name] = tftypes.NewValue(attrType, nil)
	}

	for name := range attrs {
		if _, ok := objType.AttributeTypes[name]; !ok {
			t.Fatalf("provider schema has no attribute %q", name)
		}
	}

	return tfsdk.Config{Schema: schema, Raw: tftypes.NewValue(objType, values)}
}

// configureAndCaptureUserAgent configures the provider against a local test
// server and returns the User-Agent header the resulting client sends.
func configureAndCaptureUserAgent(t *testing.T, attrs map[string]tftypes.Value) string {
	t.Helper()
	ctx := context.Background()

	var userAgent string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userAgent = r.Header.Get("User-Agent")
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"success":true,"errors":[],"messages":[],"result":[]}`))
	}))
	defer srv.Close()

	if _, ok := attrs[consts.BaseURLSchemaKey]; !ok {
		if attrs == nil {
			attrs = map[string]tftypes.Value{}
		}
		attrs[consts.BaseURLSchemaKey] = tftypes.NewValue(tftypes.String, srv.URL)
	}

	resp := provider.ConfigureResponse{}
	internal.NewProvider(testProviderVersion)().Configure(ctx, provider.ConfigureRequest{
		TerraformVersion: testTerraformVersion,
		Config:           providerConfig(t, ctx, attrs),
	}, &resp)
	if resp.Diagnostics.HasError() {
		t.Fatalf("Configure failed: %v", resp.Diagnostics)
	}

	client, ok := resp.ResourceData.(*cloudflare.Client)
	if !ok {
		t.Fatalf("ResourceData is %T, expected *cloudflare.Client", resp.ResourceData)
	}

	// The response body is irrelevant; we only care about what went out on the
	// wire, so any decoding error is ignored.
	var out interface{}
	_ = client.Get(ctx, "zones", nil, &out)

	if userAgent == "" {
		t.Fatal("no request reached the test server")
	}
	return userAgent
}

// unsetEnv removes a variable for the duration of the test.
func unsetEnv(t *testing.T, key string) {
	t.Helper()
	if original, ok := os.LookupEnv(key); ok {
		t.Cleanup(func() { os.Setenv(key, original) })
		os.Unsetenv(key)
	}
}

func TestProviderUserAgentOperatorSuffix(t *testing.T) {
	suffixKey := consts.UserAgentOperatorSuffixSchemaKey

	tests := []struct {
		name       string
		attrs      map[string]tftypes.Value
		env        *string
		expectTail string
	}{
		{
			// Regression: the suffix used to be formatted with
			// StringValue.String(), which wraps the value in literal quotes.
			name:       "configured suffix is sent verbatim",
			attrs:      map[string]tftypes.Value{suffixKey: tftypes.NewValue(tftypes.String, "mycorp/1.0")},
			expectTail: " mycorp/1.0",
		},
		{
			// Regression: IsNull() alone let unknown values through as the
			// literal string "<unknown>".
			name:       "unknown suffix falls back to the Terraform version",
			attrs:      map[string]tftypes.Value{suffixKey: tftypes.NewValue(tftypes.String, tftypes.UnknownValue)},
			expectTail: " terraform/" + testTerraformVersion,
		},
		{
			name:       "suffix can come from the environment",
			env:        stringPtr("mycorp/1.0"),
			expectTail: " mycorp/1.0",
		},
		{
			name:       "configuration takes precedence over the environment",
			attrs:      map[string]tftypes.Value{suffixKey: tftypes.NewValue(tftypes.String, "config/1.0")},
			env:        stringPtr("env/1.0"),
			expectTail: " config/1.0",
		},
		{
			name:       "no suffix falls back to the Terraform version",
			expectTail: " terraform/" + testTerraformVersion,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if tc.env != nil {
				t.Setenv(consts.UserAgentOperatorSuffixEnvVarKey, *tc.env)
			} else {
				unsetEnv(t, consts.UserAgentOperatorSuffixEnvVarKey)
			}

			got := configureAndCaptureUserAgent(t, tc.attrs)

			if !strings.HasPrefix(got, "terraform-provider-cloudflare/"+testProviderVersion+" terraform-plugin-framework") {
				t.Errorf("unexpected user agent prefix: %q", got)
			}
			if !strings.HasSuffix(got, tc.expectTail) {
				t.Errorf("expected user agent %q to end with %q", got, tc.expectTail)
			}
		})
	}
}

func stringPtr(s string) *string {
	return &s
}

// Actions get their client from ActionData. Codegen rewrites Configure in
// provider.go, so this fails if a regeneration drops that line.
func TestProviderConfigureSetsActionData(t *testing.T) {
	ctx := context.Background()
	resp := provider.ConfigureResponse{}
	internal.NewProvider(testProviderVersion)().Configure(ctx, provider.ConfigureRequest{
		TerraformVersion: testTerraformVersion,
		Config: providerConfig(t, ctx, map[string]tftypes.Value{
			consts.APITokenSchemaKey: tftypes.NewValue(tftypes.String, "test-token"),
		}),
	}, &resp)
	if resp.Diagnostics.HasError() {
		t.Fatalf("Configure failed: %v", resp.Diagnostics)
	}
	if _, ok := resp.ActionData.(*cloudflare.Client); !ok {
		t.Fatalf("ActionData is %T, expected *cloudflare.Client", resp.ActionData)
	}
}

func TestProviderServesActionSchemas(t *testing.T) {
	server := providerserver.NewProtocol6(internal.NewProvider(testProviderVersion)())()
	resp, err := server.GetProviderSchema(context.Background(), &tfprotov6.GetProviderSchemaRequest{})
	if err != nil {
		t.Fatalf("GetProviderSchema: %v", err)
	}
	for _, d := range resp.Diagnostics {
		if d.Severity == tfprotov6.DiagnosticSeverityError {
			t.Fatalf("GetProviderSchema diagnostic: %s: %s", d.Summary, d.Detail)
		}
	}
	if _, ok := resp.ActionSchemas["cloudflare_zero_trust_access_service_token_refresh"]; !ok {
		t.Fatalf("action cloudflare_zero_trust_access_service_token_refresh is not served; got %d action schemas", len(resp.ActionSchemas))
	}
}

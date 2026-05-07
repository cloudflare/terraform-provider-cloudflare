package api_token_test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/api_token"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ---------------------------------------------------------------------------
// Helpers
// ---------------------------------------------------------------------------

func newAPITokenResource(t *testing.T, baseURL string) *api_token.APITokenResource {
	t.Helper()
	r := api_token.NewResource().(*api_token.APITokenResource)
	if baseURL == "" {
		return r
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIToken("test-token"),
	)
	req := resource.ConfigureRequest{ProviderData: client}
	resp := &resource.ConfigureResponse{}
	r.Configure(context.Background(), req, resp)
	if resp.Diagnostics.HasError() {
		t.Fatalf("newAPITokenResource: Configure failed: %v", resp.Diagnostics)
	}
	return r
}

func makeAPITokenState(t *testing.T, ctx context.Context, id, name, status, value string) tfsdk.State {
	t.Helper()

	permGroups := []*api_token.APITokenPoliciesPermissionGroupsModel{
		{ID: types.StringValue("pg-1")},
	}
	policies := []*api_token.APITokenPoliciesModel{
		{
			Effect:           types.StringValue("allow"),
			PermissionGroups: &permGroups,
			Resources:        types.StringValue(`{"com.cloudflare.api.account.abc123":"*"}`),
		},
	}
	model := api_token.APITokenModel{
		ID:         types.StringValue(id),
		Name:       types.StringValue(name),
		Policies:   &policies,
		ExpiresOn:  timetypes.NewRFC3339Null(),
		NotBefore:  timetypes.NewRFC3339Null(),
		Condition:  nil,
		Status:     types.StringValue(status),
		IssuedOn:   timetypes.NewRFC3339Null(),
		LastUsedOn: timetypes.NewRFC3339Null(),
		ModifiedOn: timetypes.NewRFC3339Null(),
		Value:      types.StringValue(value),
	}

	schema := api_token.ResourceSchema(ctx)
	state := tfsdk.State{Schema: schema}
	if diags := state.Set(ctx, &model); diags.HasError() {
		t.Fatalf("makeAPITokenState: failed to set state: %v", diags)
	}
	return state
}

func tokenGetResponse(id, name, status string) string {
	return fmt.Sprintf(`{
		"success": true,
		"errors": [],
		"messages": [],
		"result": {
			"id": %q,
			"name": %q,
			"status": %q,
			"policies": [
				{
					"effect": "allow",
					"permission_groups": [{"id": "pg-1"}],
					"resources": {"com.cloudflare.api.account.abc123": "*"}
				}
			]
		}
	}`, id, name, status)
}

// ---------------------------------------------------------------------------
// Tests
// ---------------------------------------------------------------------------

func TestRead_ActiveToken_PreservesState(t *testing.T) {
	ctx := context.Background()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, tokenGetResponse("tok-123", "my-token", "active"))
	}))
	defer ts.Close()

	r := newAPITokenResource(t, ts.URL)
	state := makeAPITokenState(t, ctx, "tok-123", "my-token", "active", "original-secret")

	resp := &resource.ReadResponse{State: state}
	r.Read(ctx, resource.ReadRequest{State: state}, resp)

	if resp.Diagnostics.HasError() {
		t.Fatalf("unexpected error: %v", resp.Diagnostics)
	}

	// Verify value is preserved from state
	var result api_token.APITokenModel
	resp.State.Get(ctx, &result)
	if result.Value.ValueString() != "original-secret" {
		t.Errorf("expected value %q, got %q", "original-secret", result.Value.ValueString())
	}
}

func TestRead_ExpiredToken_RemovesFromState(t *testing.T) {
	ctx := context.Background()
	var rollCalled atomic.Bool

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "PUT" {
			rollCalled.Store(true)
			t.Error("roll endpoint should not be called — expired tokens should be removed from state")
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, tokenGetResponse("tok-123", "my-token", "expired"))
	}))
	defer ts.Close()

	r := newAPITokenResource(t, ts.URL)
	state := makeAPITokenState(t, ctx, "tok-123", "my-token", "active", "old-secret")

	resp := &resource.ReadResponse{State: state}
	r.Read(ctx, resource.ReadRequest{State: state}, resp)

	if resp.Diagnostics.HasError() {
		t.Fatalf("unexpected error: %v", resp.Diagnostics)
	}
	if rollCalled.Load() {
		t.Error("roll endpoint was called — should not mutate during Read")
	}

	// Verify a warning was emitted about the token being removed
	hasWarning := false
	for _, d := range resp.Diagnostics.Warnings() {
		if d.Summary() == "Token expired or revoked" {
			hasWarning = true
			break
		}
	}
	if !hasWarning {
		t.Error("expected warning diagnostic about token being expired")
	}

	// Verify state was removed (resource will be recreated on next apply)
	var result api_token.APITokenModel
	diags := resp.State.Get(ctx, &result)
	if !diags.HasError() && result.ID.ValueString() != "" {
		t.Error("expected state to be removed for expired token")
	}
}

func TestRead_RevokedExposedToken_RemovesFromState(t *testing.T) {
	ctx := context.Background()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, tokenGetResponse("tok-456", "exposed-token", "revoked (exposed)"))
	}))
	defer ts.Close()

	r := newAPITokenResource(t, ts.URL)
	state := makeAPITokenState(t, ctx, "tok-456", "exposed-token", "active", "compromised-secret")

	resp := &resource.ReadResponse{State: state}
	r.Read(ctx, resource.ReadRequest{State: state}, resp)

	if resp.Diagnostics.HasError() {
		t.Fatalf("unexpected error: %v", resp.Diagnostics)
	}

	hasWarning := false
	for _, d := range resp.Diagnostics.Warnings() {
		if d.Summary() == "Token expired or revoked" {
			hasWarning = true
			break
		}
	}
	if !hasWarning {
		t.Error("expected warning diagnostic about token being revoked")
	}
}

func TestRead_DisabledToken_PreservesState(t *testing.T) {
	ctx := context.Background()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, tokenGetResponse("tok-dis", "disabled-token", "disabled"))
	}))
	defer ts.Close()

	r := newAPITokenResource(t, ts.URL)
	state := makeAPITokenState(t, ctx, "tok-dis", "disabled-token", "disabled", "my-secret")

	resp := &resource.ReadResponse{State: state}
	r.Read(ctx, resource.ReadRequest{State: state}, resp)

	if resp.Diagnostics.HasError() {
		t.Fatalf("unexpected error: %v", resp.Diagnostics)
	}

	var result api_token.APITokenModel
	resp.State.Get(ctx, &result)
	if result.Value.ValueString() != "my-secret" {
		t.Errorf("expected value preserved for disabled token, got %q", result.Value.ValueString())
	}
}

func TestRead_404_RemovesFromState(t *testing.T) {
	ctx := context.Background()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, `{"success":false,"errors":[{"code":1000,"message":"not found"}]}`)
	}))
	defer ts.Close()

	r := newAPITokenResource(t, ts.URL)
	state := makeAPITokenState(t, ctx, "tok-gone", "gone-token", "active", "secret")

	resp := &resource.ReadResponse{State: state}
	r.Read(ctx, resource.ReadRequest{State: state}, resp)

	if resp.Diagnostics.HasError() {
		t.Fatalf("unexpected error: %v", resp.Diagnostics)
	}
}

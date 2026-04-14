package account_token_test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/account_token"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ---------------------------------------------------------------------------
// Helpers
// ---------------------------------------------------------------------------

func newAccountTokenResource(t *testing.T, baseURL string) *account_token.AccountTokenResource {
	t.Helper()
	r := account_token.NewResource().(*account_token.AccountTokenResource)
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
		t.Fatalf("newAccountTokenResource: Configure failed: %v", resp.Diagnostics)
	}
	return r
}

func makeAccountTokenState(t *testing.T, ctx context.Context, id, accountID, name, status, value string) tfsdk.State {
	t.Helper()

	permGroups := []*account_token.AccountTokenPoliciesPermissionGroupsModel{
		{ID: types.StringValue("pg-1")},
	}
	policies := []*account_token.AccountTokenPoliciesModel{
		{
			Effect:           types.StringValue("allow"),
			PermissionGroups: &permGroups,
			Resources:        types.StringValue(`{"com.cloudflare.api.account.abc123":"*"}`),
		},
	}
	model := account_token.AccountTokenModel{
		ID:         types.StringValue(id),
		AccountID:  types.StringValue(accountID),
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

	schema := account_token.ResourceSchema(ctx)
	state := tfsdk.State{Schema: schema}
	if diags := state.Set(ctx, &model); diags.HasError() {
		t.Fatalf("makeAccountTokenState: failed to set state: %v", diags)
	}
	return state
}

func accountTokenGetResponse(id, name, status string) string {
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

func TestRead_ActiveAccountToken_PreservesState(t *testing.T) {
	ctx := context.Background()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, accountTokenGetResponse("tok-123", "my-account-token", "active"))
	}))
	defer ts.Close()

	r := newAccountTokenResource(t, ts.URL)
	state := makeAccountTokenState(t, ctx, "tok-123", "acct-abc", "my-account-token", "active", "original-secret")

	resp := &resource.ReadResponse{State: state}
	r.Read(ctx, resource.ReadRequest{State: state}, resp)

	if resp.Diagnostics.HasError() {
		t.Fatalf("unexpected error: %v", resp.Diagnostics)
	}

	var result account_token.AccountTokenModel
	resp.State.Get(ctx, &result)
	if result.Value.ValueString() != "original-secret" {
		t.Errorf("expected value %q, got %q", "original-secret", result.Value.ValueString())
	}
}

func TestRead_ExpiredAccountToken_RemovesFromState(t *testing.T) {
	ctx := context.Background()
	var rollCalled atomic.Bool

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "PUT" {
			rollCalled.Store(true)
			t.Error("roll endpoint should not be called — expired tokens should be removed from state")
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, accountTokenGetResponse("tok-123", "my-account-token", "expired"))
	}))
	defer ts.Close()

	r := newAccountTokenResource(t, ts.URL)
	state := makeAccountTokenState(t, ctx, "tok-123", "acct-abc", "my-account-token", "active", "old-secret")

	resp := &resource.ReadResponse{State: state}
	r.Read(ctx, resource.ReadRequest{State: state}, resp)

	if resp.Diagnostics.HasError() {
		t.Fatalf("unexpected error: %v", resp.Diagnostics)
	}
	if rollCalled.Load() {
		t.Error("roll endpoint was called — should not mutate during Read")
	}

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
}

func TestRead_RevokedExposedAccountToken_RemovesFromState(t *testing.T) {
	ctx := context.Background()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, accountTokenGetResponse("tok-456", "exposed-token", "revoked (exposed)"))
	}))
	defer ts.Close()

	r := newAccountTokenResource(t, ts.URL)
	state := makeAccountTokenState(t, ctx, "tok-456", "acct-abc", "exposed-token", "active", "compromised-secret")

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

func TestRead_DisabledAccountToken_PreservesState(t *testing.T) {
	ctx := context.Background()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, accountTokenGetResponse("tok-dis", "disabled-token", "disabled"))
	}))
	defer ts.Close()

	r := newAccountTokenResource(t, ts.URL)
	state := makeAccountTokenState(t, ctx, "tok-dis", "acct-abc", "disabled-token", "disabled", "my-secret")

	resp := &resource.ReadResponse{State: state}
	r.Read(ctx, resource.ReadRequest{State: state}, resp)

	if resp.Diagnostics.HasError() {
		t.Fatalf("unexpected error: %v", resp.Diagnostics)
	}

	var result account_token.AccountTokenModel
	resp.State.Get(ctx, &result)
	if result.Value.ValueString() != "my-secret" {
		t.Errorf("expected value preserved for disabled token, got %q", result.Value.ValueString())
	}
}

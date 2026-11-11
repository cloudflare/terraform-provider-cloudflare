package email_routing_settings_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/email_routing_settings"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func subaddressState(t *testing.T, value types.Bool) tfsdk.State {
	t.Helper()
	state := tfsdk.State{Schema: email_routing_settings.ResourceSchema(context.Background())}
	model := email_routing_settings.EmailRoutingSettingsModel{
		ID: types.StringValue("zone-id"), ZoneID: types.StringValue("zone-id"),
		SupportSubaddress: value,
	}
	if diags := state.Set(context.Background(), &model); diags.HasError() {
		t.Fatal(diags)
	}
	return state
}

func subaddressResource(t *testing.T, handler http.HandlerFunc) *email_routing_settings.EmailRoutingSettingsResource {
	t.Helper()
	server := httptest.NewServer(handler)
	t.Cleanup(server.Close)
	r := email_routing_settings.NewResource().(*email_routing_settings.EmailRoutingSettingsResource)
	response := &resource.ConfigureResponse{}
	r.Configure(context.Background(), resource.ConfigureRequest{ProviderData: cloudflare.NewClient(
		option.WithBaseURL(server.URL), option.WithMaxRetries(0),
	)}, response)
	if response.Diagnostics.HasError() {
		t.Fatal(response.Diagnostics)
	}
	return r
}

func subaddressResponse(w http.ResponseWriter, value bool) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]any{
		"success": true,
		"result": map[string]any{
			"id": "zone-id", "name": "example.com", "enabled": true,
			"support_subaddress": value, "status": "ready", "skip_wizard": false,
			"created": "2026-10-28T10:00:00Z", "modified": "2026-10-28T11:00:00Z",
		},
	})
}

func checkSubaddressPatch(t *testing.T, req *http.Request, want bool) {
	t.Helper()
	if req.Method != http.MethodPatch || req.URL.Path != "/zones/zone-id/email/routing" {
		t.Errorf("request = %s %s; want PATCH /zones/zone-id/email/routing", req.Method, req.URL)
	}
	var body map[string]any
	if err := json.NewDecoder(req.Body).Decode(&body); err != nil {
		t.Fatal(err)
	}
	if len(body) != 1 || body["support_subaddress"] != want {
		t.Errorf("request body = %#v; want only support_subaddress=%t", body, want)
	}
}

func TestEmailRoutingSettingsUpdateSubaddress(t *testing.T) {
	for _, want := range []bool{true, false} {
		t.Run(types.BoolValue(want).String(), func(t *testing.T) {
			requests := 0
			r := subaddressResource(t, func(w http.ResponseWriter, req *http.Request) {
				requests++
				checkSubaddressPatch(t, req, want)
				subaddressResponse(w, want)
			})
			prior := subaddressState(t, types.BoolValue(!want))
			planned := subaddressState(t, types.BoolValue(want))
			response := &resource.UpdateResponse{State: prior}
			r.Update(context.Background(), resource.UpdateRequest{
				State: prior, Plan: tfsdk.Plan{Schema: planned.Schema, Raw: planned.Raw},
			}, response)
			if response.Diagnostics.HasError() {
				t.Fatal(response.Diagnostics)
			}
			if requests != 1 {
				t.Fatalf("PATCH requests = %d; want 1", requests)
			}
			var actual email_routing_settings.EmailRoutingSettingsModel
			if diags := response.State.Get(context.Background(), &actual); diags.HasError() {
				t.Fatal(diags)
			}
			if !actual.SupportSubaddress.Equal(types.BoolValue(want)) || actual.Status.ValueString() != "ready" {
				t.Errorf("unexpected updated state: %#v", actual)
			}
		})
	}
}

func TestEmailRoutingSettingsCreateSubaddress(t *testing.T) {
	for _, value := range []types.Bool{types.BoolValue(true), types.BoolValue(false), types.BoolNull(), types.BoolUnknown()} {
		t.Run(value.String(), func(t *testing.T) {
			requests := []string{}
			r := subaddressResource(t, func(w http.ResponseWriter, req *http.Request) {
				requests = append(requests, req.Method+" "+req.URL.Path)
				if len(requests) == 1 {
					if req.Method != http.MethodPost || req.URL.Path != "/zones/zone-id/email/routing/enable" {
						t.Errorf("unexpected enable request: %s %s", req.Method, req.URL)
					}
					subaddressResponse(w, true)
					return
				}
				checkSubaddressPatch(t, req, value.ValueBool())
				subaddressResponse(w, value.ValueBool())
			})
			planned := subaddressState(t, value)
			response := &resource.CreateResponse{State: tfsdk.State{Schema: planned.Schema}}
			r.Create(context.Background(), resource.CreateRequest{Plan: tfsdk.Plan{Schema: planned.Schema, Raw: planned.Raw}}, response)
			if response.Diagnostics.HasError() {
				t.Fatal(response.Diagnostics)
			}
			wantRequests, wantValue := 2, value.ValueBool()
			if value.IsNull() || value.IsUnknown() {
				wantRequests, wantValue = 1, true
			}
			if len(requests) != wantRequests {
				t.Fatalf("requests = %v; want %d", requests, wantRequests)
			}
			var actual email_routing_settings.EmailRoutingSettingsModel
			if diags := response.State.Get(context.Background(), &actual); diags.HasError() {
				t.Fatal(diags)
			}
			if !actual.SupportSubaddress.Equal(types.BoolValue(wantValue)) {
				t.Errorf("support_subaddress = %s; want %t", actual.SupportSubaddress, wantValue)
			}
		})
	}
}

func TestEmailRoutingSettingsUpdateSubaddressFailure(t *testing.T) {
	r := subaddressResource(t, func(w http.ResponseWriter, req *http.Request) {
		checkSubaddressPatch(t, req, true)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusForbidden)
		_, _ = w.Write([]byte(`{"success":false,"errors":[{"code":10000,"message":"not authorized"}]}`))
	})
	prior := subaddressState(t, types.BoolValue(false))
	planned := subaddressState(t, types.BoolValue(true))
	response := &resource.UpdateResponse{State: prior}
	r.Update(context.Background(), resource.UpdateRequest{
		State: prior, Plan: tfsdk.Plan{Schema: planned.Schema, Raw: planned.Raw},
	}, response)
	if !response.Diagnostics.HasError() {
		t.Fatal("expected API error diagnostic")
	}
	if !response.State.Raw.Equal(prior.Raw) {
		t.Fatal("failed update changed prior state")
	}
}

func TestEmailRoutingSettingsCreateSubaddressFailurePreservesEnabledResource(t *testing.T) {
	requests := 0
	r := subaddressResource(t, func(w http.ResponseWriter, req *http.Request) {
		requests++
		if req.Method == http.MethodPost && req.URL.Path == "/zones/zone-id/email/routing/enable" {
			subaddressResponse(w, true)
			return
		}
		checkSubaddressPatch(t, req, false)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusForbidden)
		_, _ = w.Write([]byte(`{"success":false,"errors":[{"code":10000,"message":"not authorized"}]}`))
	})
	planned := subaddressState(t, types.BoolValue(false))
	response := &resource.CreateResponse{State: tfsdk.State{Schema: planned.Schema}}
	r.Create(context.Background(), resource.CreateRequest{Plan: tfsdk.Plan{Schema: planned.Schema, Raw: planned.Raw}}, response)
	if !response.Diagnostics.HasError() || requests != 2 {
		t.Fatalf("expected PATCH failure after enable; diagnostics=%v, requests=%d", response.Diagnostics, requests)
	}
	var actual email_routing_settings.EmailRoutingSettingsModel
	if diags := response.State.Get(context.Background(), &actual); diags.HasError() {
		t.Fatal(diags)
	}
	if actual.ID.ValueString() != "zone-id" || !actual.Enabled.ValueBool() || !actual.SupportSubaddress.ValueBool() {
		t.Fatalf("failed create did not preserve enabled resource for recovery: %#v", actual)
	}
}

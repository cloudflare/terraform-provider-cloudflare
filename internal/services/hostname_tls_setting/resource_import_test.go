package hostname_tls_setting_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"sync/atomic"
	"testing"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/hostname_tls_setting"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
)

func TestHostnameTLSSettingImport_PreservesIdentityAndRefreshes(t *testing.T) {
	t.Parallel()
	for _, tt := range []struct {
		setting string
		value   any
	}{
		{"min_tls_version", "1.2"},
		{"ciphers", []any{"ECDHE-RSA-AES128-GCM-SHA256", "ECDHE-RSA-AES256-GCM-SHA384"}},
		{"http2", "on"},
	} {
		t.Run(tt.setting, func(t *testing.T) {
			ctx := context.Background()
			const zoneID = "zone-id"
			const hostname = "app.example.com"
			importID := zoneID + "/" + tt.setting + "/" + hostname
			var requests atomic.Int32
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
				requests.Add(1)
				wantPath := "/zones/" + zoneID + "/hostnames/settings/" + tt.setting + "/" + hostname
				if req.Method != http.MethodGet || req.URL.Path != wantPath {
					t.Errorf("unexpected request: %s %s, want GET %s", req.Method, req.URL, wantPath)
					http.Error(w, "unexpected request", http.StatusBadRequest)
					return
				}
				w.Header().Set("Content-Type", "application/json")
				// The API returns neither zone_id nor setting_id. Import must
				// retain them from the import path, not infer them from JSON.
				_ = json.NewEncoder(w).Encode(map[string]any{
					"success": true,
					"result": map[string]any{
						"hostname":   hostname,
						"value":      tt.value,
						"status":     "active",
						"created_at": "2026-09-29T12:00:00Z",
						"updated_at": "2026-09-29T12:00:00Z",
					},
				})
			}))
			defer server.Close()

			r := hostname_tls_setting.NewResource().(*hostname_tls_setting.HostnameTLSSettingResource)
			configure := &resource.ConfigureResponse{}
			r.Configure(ctx, resource.ConfigureRequest{ProviderData: cloudflare.NewClient(
				option.WithBaseURL(server.URL), option.WithAPIToken("test-token"),
			)}, configure)
			if configure.Diagnostics.HasError() {
				t.Fatalf("configure: %v", configure.Diagnostics)
			}

			checkState := func(state tfsdk.State) {
				t.Helper()
				var model hostname_tls_setting.HostnameTLSSettingModel
				if diags := state.Get(ctx, &model); diags.HasError() {
					t.Fatalf("get state: %v", diags)
				}
				if model.ID.ValueString() != tt.setting || model.SettingID.ValueString() != tt.setting {
					t.Fatalf("identity: id = %s, setting_id = %s, want %q for both", model.ID, model.SettingID, tt.setting)
				}
				if model.ZoneID.ValueString() != zoneID || model.Hostname.ValueString() != hostname {
					t.Fatalf("path attributes: zone_id = %s, hostname = %s", model.ZoneID, model.Hostname)
				}
				body, err := model.MarshalJSON()
				if err != nil {
					t.Fatalf("marshal value: %v", err)
				}
				var decoded struct {
					Value any `json:"value"`
				}
				if err := json.Unmarshal(body, &decoded); err != nil {
					t.Fatalf("decode value: %v", err)
				}
				if !reflect.DeepEqual(decoded.Value, tt.value) {
					t.Errorf("value = %#v, want %#v", decoded.Value, tt.value)
				}
			}

			imported := &resource.ImportStateResponse{State: tfsdk.State{Schema: hostname_tls_setting.ResourceSchema(ctx)}}
			r.ImportState(ctx, resource.ImportStateRequest{ID: importID}, imported)
			if imported.Diagnostics.HasError() {
				t.Fatalf("import: %v", imported.Diagnostics)
			}
			checkState(imported.State)

			refreshed := &resource.ReadResponse{State: imported.State}
			r.Read(ctx, resource.ReadRequest{State: imported.State}, refreshed)
			if refreshed.Diagnostics.HasError() {
				t.Fatalf("refresh: %v", refreshed.Diagnostics)
			}
			checkState(refreshed.State)
			if requests.Load() != 2 {
				t.Errorf("got %d requests, want one import GET and one refresh GET", requests.Load())
			}
		})
	}
}

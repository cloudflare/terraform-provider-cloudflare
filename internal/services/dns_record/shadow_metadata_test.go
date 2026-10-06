package dns_record_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/dns_record"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

const shadowMetadataTimestamp = "2026-05-03T19:27:14Z"

func shadowMetadataModel(ctx context.Context, shadow types.Bool) dns_record.DNSRecordModel {
	return dns_record.DNSRecordModel{
		ID:                    types.StringValue("record-id"),
		ZoneID:                types.StringValue("zone-id"),
		IncludeShadowMetadata: shadow,
		Name:                  types.StringValue("example.com"),
		Type:                  types.StringValue("MX"),
		Content:               types.StringValue("route1.mx.cloudflare.net"),
		Priority:              types.Float64Value(40),
		Proxied:               types.BoolValue(false),
		TTL:                   types.Float64Value(300),
		Tags:                  customfield.NewSetMust[types.String](ctx, nil),
		Settings:              customfield.NullObject[dns_record.DNSRecordSettingsModel](ctx),
		CreatedOn:             timetypes.NewRFC3339ValueMust(shadowMetadataTimestamp),
		ModifiedOn:            timetypes.NewRFC3339ValueMust(shadowMetadataTimestamp),
		Proxiable:             types.BoolValue(false),
		Meta:                  jsontypes.NewNormalizedValue("{}"),
	}
}

func shadowMetadataState(t *testing.T, ctx context.Context, model dns_record.DNSRecordModel) tfsdk.State {
	t.Helper()
	state := tfsdk.State{Schema: dns_record.ResourceSchema(ctx)}
	if diags := state.Set(ctx, &model); diags.HasError() {
		t.Fatalf("set state: %v", diags)
	}
	return state
}

func shadowMetadataPlan(t *testing.T, ctx context.Context, model dns_record.DNSRecordModel) tfsdk.Plan {
	t.Helper()
	plan := tfsdk.Plan{Schema: dns_record.ResourceSchema(ctx)}
	if diags := plan.Set(ctx, &model); diags.HasError() {
		t.Fatalf("set plan: %v", diags)
	}
	return plan
}

// Older v5 state (including migrated v4 records) has null for this query-only
// field. The API cannot populate it, so Read must backfill it before planning.
func TestDNSRecordRead_ShadowMetadataRefreshPlansNoUpdate(t *testing.T) {
	t.Parallel()
	for _, tt := range []struct {
		name  string
		prior types.Bool
		want  types.Bool
		query string
	}{
		{"missing_default", types.BoolNull(), types.BoolValue(false), ""},
		{"preserve_false", types.BoolValue(false), types.BoolValue(false), "include_shadow_metadata=false"},
		{"preserve_true", types.BoolValue(true), types.BoolValue(true), "include_shadow_metadata=true"},
	} {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			var requests atomic.Int32
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
				requests.Add(1)
				if req.Method != http.MethodGet || req.URL.Path != "/zones/zone-id/dns_records/record-id" {
					t.Errorf("unexpected request: %s %s", req.Method, req.URL)
					http.Error(w, "Email Routing records cannot be modified", http.StatusBadRequest)
					return
				}
				if req.URL.RawQuery != tt.query {
					t.Errorf("query = %q, want %q", req.URL.RawQuery, tt.query)
				}
				w.Header().Set("Content-Type", "application/json")
				_ = json.NewEncoder(w).Encode(map[string]any{
					"success": true,
					"result": map[string]any{
						"id": "record-id", "name": "example.com", "type": "MX",
						"content": "route1.mx.cloudflare.net", "priority": 40,
						"ttl": 300, "proxied": false, "proxiable": false,
						"created_on": shadowMetadataTimestamp, "modified_on": shadowMetadataTimestamp,
						"tags": []string{}, "meta": map[string]any{},
					},
				})
			}))
			defer server.Close()

			r := dns_record.NewResource().(*dns_record.DNSRecordResource)
			configure := &resource.ConfigureResponse{}
			r.Configure(ctx, resource.ConfigureRequest{ProviderData: cloudflare.NewClient(
				option.WithBaseURL(server.URL), option.WithAPIToken("test-token"),
			)}, configure)
			if configure.Diagnostics.HasError() {
				t.Fatalf("configure: %v", configure.Diagnostics)
			}
			prior := shadowMetadataState(t, ctx, shadowMetadataModel(ctx, tt.prior))
			read := &resource.ReadResponse{State: prior}
			r.Read(ctx, resource.ReadRequest{State: prior}, read)
			if read.Diagnostics.HasError() {
				t.Fatalf("read: %v", read.Diagnostics)
			}
			var refreshed dns_record.DNSRecordModel
			if diags := read.State.Get(ctx, &refreshed); diags.HasError() {
				t.Fatalf("get refreshed state: %v", diags)
			}
			if !refreshed.IncludeShadowMetadata.Equal(tt.want) {
				t.Errorf("refreshed shadow metadata = %s, want %s", refreshed.IncludeShadowMetadata, tt.want)
			}

			// Model the schema default (or explicit true) and framework-computed
			// unknowns, then invoke the real resource's ModifyPlan implementation.
			planned := refreshed
			planned.IncludeShadowMetadata = tt.want
			planned.ModifiedOn = timetypes.NewRFC3339Unknown()
			planned.CreatedOn = timetypes.NewRFC3339Unknown()
			planned.Proxiable = types.BoolUnknown()
			planned.Meta = jsontypes.NewNormalizedUnknown()
			plan := shadowMetadataPlan(t, ctx, planned)
			response := &resource.ModifyPlanResponse{Plan: plan}
			r.ModifyPlan(ctx, resource.ModifyPlanRequest{State: read.State, Plan: plan}, response)
			if response.Diagnostics.HasError() {
				t.Fatalf("modify plan: %v", response.Diagnostics)
			}
			if !response.Plan.Raw.Equal(read.State.Raw) {
				t.Errorf("unchanged record still plans an update:\nstate: %s\nplan: %s", read.State.Raw, response.Plan.Raw)
			}
			if requests.Load() != 1 {
				t.Errorf("got %d API requests, want one GET", requests.Load())
			}
		})
	}
}

func TestDNSRecordModifyPlan_ShadowMetadataChangesInvalidateModifiedOn(t *testing.T) {
	t.Parallel()
	for _, tt := range []struct {
		name          string
		prior         bool
		planned       bool
		contentChange bool
		wantUnknown   bool
	}{
		{"unchanged_false", false, false, false, false},
		{"unchanged_true", true, true, false, false},
		{"enable_shadow_metadata", false, true, false, true},
		{"disable_shadow_metadata", true, false, false, true},
		{"real_dns_change", false, false, true, true},
	} {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			model := shadowMetadataModel(ctx, types.BoolValue(tt.prior))
			state := shadowMetadataState(t, ctx, model)
			model.IncludeShadowMetadata = types.BoolValue(tt.planned)
			model.ModifiedOn = timetypes.NewRFC3339Unknown()
			if tt.contentChange {
				model.Content = types.StringValue("route2.mx.cloudflare.net")
			}
			plan := shadowMetadataPlan(t, ctx, model)
			response := &resource.ModifyPlanResponse{Plan: plan}
			dns_record.NewResource().(*dns_record.DNSRecordResource).ModifyPlan(ctx, resource.ModifyPlanRequest{State: state, Plan: plan}, response)
			if response.Diagnostics.HasError() {
				t.Fatalf("modify plan: %v", response.Diagnostics)
			}
			var result dns_record.DNSRecordModel
			if diags := response.Plan.Get(ctx, &result); diags.HasError() {
				t.Fatalf("get plan: %v", diags)
			}
			if result.ModifiedOn.IsUnknown() != tt.wantUnknown {
				t.Errorf("modified_on = %s, want unknown = %v", result.ModifiedOn, tt.wantUnknown)
			}
			if !tt.wantUnknown && result.ModifiedOn.ValueString() != shadowMetadataTimestamp {
				t.Errorf("unchanged record lost its timestamp: %s", result.ModifiedOn)
			}
		})
	}
}

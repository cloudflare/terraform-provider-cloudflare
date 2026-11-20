package email_auth_dmarc_report_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/email_auth_dmarc_report"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

func TestEmailAuthDMARCReportCreate(t *testing.T) {
	var body map[string]any
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodPatch || req.URL.Path != "/zones/zone-id/email/auth/dmarc-reports" {
			t.Errorf("unexpected request %s %s", req.Method, req.URL.Path)
		}
		if err := json.NewDecoder(req.Body).Decode(&body); err != nil {
			t.Error(err)
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]any{
			"success": true, "errors": []any{}, "messages": []any{},
			"result": map[string]any{
				"zone_id": "zone-id", "enabled": true, "skip_wizard": true,
				"rua_prefix":  "0123456789abcdef0123456789abcdef",
				"created_at":  "2026-11-20T10:00:00Z",
				"modified_at": "2026-11-20T11:00:00Z",
				"records":     map[string]any{"dmarc_records": []any{}},
			},
		})
	}))
	t.Cleanup(server.Close)

	r := email_auth_dmarc_report.NewResource().(*email_auth_dmarc_report.EmailAuthDMARCReportResource)
	configure := &resource.ConfigureResponse{}
	r.Configure(context.Background(), resource.ConfigureRequest{ProviderData: cloudflare.NewClient(
		option.WithBaseURL(server.URL), option.WithAPIToken("token"), option.WithMaxRetries(0),
	)}, configure)
	if configure.Diagnostics.HasError() {
		t.Fatal(configure.Diagnostics)
	}

	schema := email_auth_dmarc_report.ResourceSchema(context.Background())
	plan := tfsdk.Plan{Schema: schema}
	if diags := plan.Set(context.Background(), &email_auth_dmarc_report.EmailAuthDMARCReportModel{
		ID:         types.StringUnknown(),
		ZoneID:     types.StringValue("zone-id"),
		Enabled:    types.BoolValue(true),
		SkipWizard: types.BoolValue(true),
		RuaPrefix:  types.StringUnknown(),
		Status:     types.StringUnknown(),
	}); diags.HasError() {
		t.Fatal(diags)
	}
	resp := &resource.CreateResponse{State: tfsdk.State{Schema: schema, Raw: tftypes.NewValue(schema.Type().TerraformType(context.Background()), nil)}}
	r.Create(context.Background(), resource.CreateRequest{Plan: plan}, resp)
	if resp.Diagnostics.HasError() {
		t.Fatal(resp.Diagnostics)
	}

	if body["enabled"] != true || body["skip_wizard"] != true {
		t.Errorf("unexpected request body %v", body)
	}
	var got email_auth_dmarc_report.EmailAuthDMARCReportModel
	resp.State.Get(context.Background(), &got)
	if got.ID.ValueString() != "zone-id" || got.RuaPrefix.ValueString() != "0123456789abcdef0123456789abcdef" || !got.Enabled.ValueBool() {
		t.Errorf("unexpected state %+v", got)
	}
}

package certificate_pack_test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/certificate_pack"
)

const (
	testZoneID = "023e105f4ecef8ad9ca31a8372d0c353"
	testPackID = "3822ff90-ea29-44df-9e55-21300bb9419b"
)

// certificatePackGetResponse mirrors GET /zones/{zone_id}/ssl/certificate_packs/{id}.
func certificatePackGetResponse(status string) string {
	return fmt.Sprintf(`{
		"success": true,
		"errors": [],
		"messages": [],
		"result": {
			"id": %q,
			"certificate_authority": "lets_encrypt",
			"certificates": [],
			"cloudflare_branding": false,
			"hosts": ["example.com", "*.example.com"],
			"status": %q,
			"type": "advanced",
			"validation_method": "txt",
			"validity_days": 90
		}
	}`, testPackID, status)
}

func newCertificatePackResource(t *testing.T, baseURL string) *certificate_pack.CertificatePackResource {
	t.Helper()
	r := certificate_pack.NewResource().(*certificate_pack.CertificatePackResource)
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIToken("test-token"),
		option.WithMaxRetries(0),
	)
	resp := &resource.ConfigureResponse{}
	r.Configure(context.Background(), resource.ConfigureRequest{ProviderData: client}, resp)
	if resp.Diagnostics.HasError() {
		t.Fatalf("Configure: %v", resp.Diagnostics)
	}
	return r
}

func activeCertificatePackState(t *testing.T, ctx context.Context) tfsdk.State {
	t.Helper()
	model := certificate_pack.CertificatePackModel{
		ID:                   types.StringValue(testPackID),
		ZoneID:               types.StringValue(testZoneID),
		CertificateAuthority: types.StringValue("lets_encrypt"),
		Type:                 types.StringValue("advanced"),
		ValidationMethod:     types.StringValue("txt"),
		ValidityDays:         types.Int64Value(90),
		CloudflareBranding:   types.BoolNull(),
		Hosts:                customfield.NewSetMust[types.String](ctx, []attr.Value{types.StringValue("example.com"), types.StringValue("*.example.com")}),
		PrimaryCertificate:   types.StringNull(),
		Status:               types.StringValue("active"),
		Certificates:         customfield.NullObjectList[certificate_pack.CertificatePackCertificatesModel](ctx),
		DCVDelegationRecords: customfield.NullObjectList[certificate_pack.CertificatePackDCVDelegationRecordsModel](ctx),
		ValidationErrors:     customfield.NullObjectList[certificate_pack.CertificatePackValidationErrorsModel](ctx),
		ValidationRecords:    customfield.NullObjectList[certificate_pack.CertificatePackValidationRecordsModel](ctx),
	}
	state := tfsdk.State{Schema: certificate_pack.ResourceSchema(ctx)}
	if diags := state.Set(ctx, &model); diags.HasError() {
		t.Fatalf("set state: %v", diags)
	}
	return state
}

func TestCertificatePackReadDeletionTimedOut(t *testing.T) {
	ctx := context.Background()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet || r.URL.Path != "/zones/"+testZoneID+"/ssl/certificate_packs/"+testPackID {
			t.Errorf("unexpected request %s %s", r.Method, r.URL.Path)
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, certificatePackGetResponse("deletion_timed_out"))
	}))
	defer srv.Close()

	r := newCertificatePackResource(t, srv.URL)
	state := activeCertificatePackState(t, ctx)
	resp := &resource.ReadResponse{State: state}
	r.Read(ctx, resource.ReadRequest{State: state}, resp)

	if resp.Diagnostics.HasError() {
		t.Fatalf("unexpected error: %v", resp.Diagnostics)
	}
	if !resp.State.Raw.IsNull() {
		t.Fatal("a pack in deletion_timed_out must be removed from state")
	}

	warnings := resp.Diagnostics.Warnings()
	if len(warnings) != 1 {
		t.Fatalf("got %d warnings, want 1: %v", len(warnings), warnings)
	}
	warning := warnings[0]
	if warning.Severity() != diag.SeverityWarning {
		t.Fatalf("severity = %v, want warning", warning.Severity())
	}
	if got, want := warning.Summary(), "Certificate pack deletion timed out"; got != want {
		t.Errorf("summary = %q, want %q", got, want)
	}
	for _, want := range []string{
		"deletion timed out on Cloudflare's side",
		"may still exist",
		"next apply creates a new certificate pack",
		"dashboard",
		"DELETE /zones/" + testZoneID + "/ssl/certificate_packs/" + testPackID,
	} {
		if !strings.Contains(warning.Detail(), want) {
			t.Errorf("detail does not mention %q:\n%s", want, warning.Detail())
		}
	}
}

func TestCertificatePackReadDeletedKeepsGenericWarning(t *testing.T) {
	ctx := context.Background()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, certificatePackGetResponse("deleted"))
	}))
	defer srv.Close()

	r := newCertificatePackResource(t, srv.URL)
	state := activeCertificatePackState(t, ctx)
	resp := &resource.ReadResponse{State: state}
	r.Read(ctx, resource.ReadRequest{State: state}, resp)

	if resp.Diagnostics.HasError() {
		t.Fatalf("unexpected error: %v", resp.Diagnostics)
	}
	if !resp.State.Raw.IsNull() {
		t.Fatal("a deleted pack must be removed from state")
	}
	warnings := resp.Diagnostics.Warnings()
	if len(warnings) != 1 || warnings[0].Summary() != "Resource deleted" {
		t.Fatalf("got warnings %v, want one \"Resource deleted\" warning", warnings)
	}
}

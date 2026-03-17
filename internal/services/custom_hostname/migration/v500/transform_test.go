package v500

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestTransform_NormalizesOptionalFalseyValues(t *testing.T) {
	ctx := context.Background()

	source := SourceCustomHostnameModel{
		ZoneID:                    types.StringValue("zone-id"),
		Hostname:                  types.StringValue("cftftest.example.com"),
		CustomOriginServer:        types.StringValue(""),
		CustomOriginSNI:           types.StringValue(""),
		Status:                    types.StringValue(""),
		CustomMetadata:            types.MapNull(types.StringType),
		OwnershipVerification:     types.MapNull(types.StringType),
		OwnershipVerificationHTTP: types.MapNull(types.StringType),
		SSL: []SourceCustomHostnameSSL{
			{
				BundleMethod:         types.StringValue(""),
				CertificateAuthority: types.StringValue(""),
				CustomCertificate:    types.StringValue(""),
				CustomKey:            types.StringValue(""),
				Method:               types.StringValue(""),
				Type:                 types.StringValue(""),
				Wildcard:             types.BoolValue(false),
				Settings: []SourceCustomHostnameSSLSettings{
					{
						HTTP2:         types.StringValue(""),
						TLS13:         types.StringValue(""),
						MinTLSVersion: types.StringValue(""),
						EarlyHints:    types.StringValue(""),
					},
				},
			},
		},
	}

	target, diags := Transform(ctx, source)
	if diags.HasError() {
		t.Fatalf("unexpected diagnostics: %v", diags)
	}

	if !target.CustomOriginServer.IsNull() {
		t.Fatalf("expected custom_origin_server to be null, got %v", target.CustomOriginServer)
	}
	if !target.CustomOriginSNI.IsNull() {
		t.Fatalf("expected custom_origin_sni to be null, got %v", target.CustomOriginSNI)
	}
	if !target.Status.IsNull() {
		t.Fatalf("expected status to be null, got %v", target.Status)
	}

	if target.SSL == nil {
		t.Fatal("expected ssl to be set")
	}
	if !target.SSL.Wildcard.IsNull() {
		t.Fatalf("expected ssl.wildcard to be null, got %v", target.SSL.Wildcard)
	}
	if !target.SSL.CertificateAuthority.IsNull() {
		t.Fatalf("expected ssl.certificate_authority to be null, got %v", target.SSL.CertificateAuthority)
	}
	if !target.SSL.CustomCertificate.IsNull() {
		t.Fatalf("expected ssl.custom_certificate to be null, got %v", target.SSL.CustomCertificate)
	}
	if !target.SSL.CustomKey.IsNull() {
		t.Fatalf("expected ssl.custom_key to be null, got %v", target.SSL.CustomKey)
	}
	if !target.SSL.Method.IsNull() {
		t.Fatalf("expected ssl.method to be null, got %v", target.SSL.Method)
	}
	if target.SSL.BundleMethod.ValueString() != "ubiquitous" {
		t.Fatalf("expected ssl.bundle_method default ubiquitous, got %q", target.SSL.BundleMethod.ValueString())
	}
	if target.SSL.Type.ValueString() != "dv" {
		t.Fatalf("expected ssl.type default dv, got %q", target.SSL.Type.ValueString())
	}

	if target.SSL.Settings == nil {
		t.Fatal("expected ssl.settings to be set")
	}
	if !target.SSL.Settings.HTTP2.IsNull() {
		t.Fatalf("expected ssl.settings.http2 to be null, got %v", target.SSL.Settings.HTTP2)
	}
	if !target.SSL.Settings.TLS1_3.IsNull() {
		t.Fatalf("expected ssl.settings.tls_1_3 to be null, got %v", target.SSL.Settings.TLS1_3)
	}
	if !target.SSL.Settings.MinTLSVersion.IsNull() {
		t.Fatalf("expected ssl.settings.min_tls_version to be null, got %v", target.SSL.Settings.MinTLSVersion)
	}
	if !target.SSL.Settings.EarlyHints.IsNull() {
		t.Fatalf("expected ssl.settings.early_hints to be null, got %v", target.SSL.Settings.EarlyHints)
	}
}

func TestTransform_PreservesExplicitTrueBool(t *testing.T) {
	ctx := context.Background()

	source := SourceCustomHostnameModel{
		ZoneID:                    types.StringValue("zone-id"),
		Hostname:                  types.StringValue("cftftest.example.com"),
		CustomMetadata:            types.MapNull(types.StringType),
		OwnershipVerification:     types.MapNull(types.StringType),
		OwnershipVerificationHTTP: types.MapNull(types.StringType),
		SSL: []SourceCustomHostnameSSL{
			{
				BundleMethod: types.StringValue("ubiquitous"),
				Type:         types.StringValue("dv"),
				Wildcard:     types.BoolValue(true),
			},
		},
	}

	target, diags := Transform(ctx, source)
	if diags.HasError() {
		t.Fatalf("unexpected diagnostics: %v", diags)
	}

	if target.SSL == nil {
		t.Fatal("expected ssl to be set")
	}
	if target.SSL.Wildcard.IsNull() || target.SSL.Wildcard.IsUnknown() {
		t.Fatalf("expected ssl.wildcard to be known true, got %v", target.SSL.Wildcard)
	}
	if !target.SSL.Wildcard.ValueBool() {
		t.Fatalf("expected ssl.wildcard to remain true, got %v", target.SSL.Wildcard)
	}
}

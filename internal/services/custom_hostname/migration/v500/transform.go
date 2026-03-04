package v500

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/migrations"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func Transform(ctx context.Context, source SourceCustomHostnameModel) (*TargetCustomHostnameModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	if source.ZoneID.IsNull() || source.ZoneID.IsUnknown() {
		diags.AddError("Missing required field", "zone_id is required for custom_hostname migration")
		return nil, diags
	}
	if source.Hostname.IsNull() || source.Hostname.IsUnknown() {
		diags.AddError("Missing required field", "hostname is required for custom_hostname migration")
		return nil, diags
	}

	target := &TargetCustomHostnameModel{
		ID:                        source.ID,
		ZoneID:                    source.ZoneID,
		Hostname:                  source.Hostname,
		CustomOriginServer:        migrations.FalseyStringToNull(source.CustomOriginServer),
		CustomOriginSNI:           migrations.FalseyStringToNull(source.CustomOriginSNI),
		CreatedAt:                 timetypes.NewRFC3339Null(),
		Status:                    migrations.FalseyStringToNull(source.Status),
		VerificationErrors:        customfield.NullList[types.String](ctx),
		OwnershipVerification:     customfield.NullObject[TargetCustomHostnameOwnershipVerification](ctx),
		OwnershipVerificationHTTP: customfield.NullObject[TargetCustomHostnameOwnershipVerificationHTTP](ctx),
	}

	ssl, sslDiags := transformSSL(ctx, source.SSL)
	diags.Append(sslDiags...)
	target.SSL = ssl

	customMetadata, metadataDiags := transformCustomMetadata(ctx, source.CustomMetadata)
	diags.Append(metadataDiags...)
	target.CustomMetadata = customMetadata

	verificationErrors, verificationErrorsDiags := transformVerificationErrors(ctx, source.SSL)
	diags.Append(verificationErrorsDiags...)
	target.VerificationErrors = verificationErrors

	ownershipVerification, ownershipVerificationDiags := transformOwnershipVerification(ctx, source.OwnershipVerification)
	diags.Append(ownershipVerificationDiags...)
	target.OwnershipVerification = ownershipVerification

	ownershipVerificationHTTP, ownershipVerificationHTTPDiags := transformOwnershipVerificationHTTP(ctx, source.OwnershipVerificationHTTP)
	diags.Append(ownershipVerificationHTTPDiags...)
	target.OwnershipVerificationHTTP = ownershipVerificationHTTP

	return target, diags
}

func transformSSL(ctx context.Context, sourceSSL []SourceCustomHostnameSSL) (*TargetCustomHostnameSSLModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	if len(sourceSSL) == 0 {
		return &TargetCustomHostnameSSLModel{
			BundleMethod:       types.StringValue("ubiquitous"),
			Type:               types.StringValue("dv"),
			CloudflareBranding: types.BoolNull(),
			CustomCERTBundle:   nil,
			CustomCertificate:  types.StringNull(),
			CustomKey:          types.StringNull(),
			Method:             types.StringNull(),
			Settings:           nil,
			Wildcard:           types.BoolNull(),
		}, diags
	}

	source := sourceSSL[0]
	target := &TargetCustomHostnameSSLModel{
		BundleMethod:         migrations.FalseyStringToNull(source.BundleMethod),
		CertificateAuthority: migrations.FalseyStringToNull(source.CertificateAuthority),
		CloudflareBranding:   types.BoolNull(),
		CustomCERTBundle:     nil,
		CustomCertificate:    migrations.FalseyStringToNull(source.CustomCertificate),
		CustomKey:            migrations.FalseyStringToNull(source.CustomKey),
		Method:               migrations.FalseyStringToNull(source.Method),
		Type:                 source.Type,
		Wildcard:             migrations.FalseyBoolToNull(source.Wildcard),
	}
	target.Type = migrations.FalseyStringToNull(target.Type)

	if target.BundleMethod.IsNull() || target.BundleMethod.IsUnknown() {
		target.BundleMethod = types.StringValue("ubiquitous")
	}
	if target.Type.IsNull() || target.Type.IsUnknown() {
		target.Type = types.StringValue("dv")
	}

	if len(source.Settings) > 0 {
		settings, settingsDiags := transformSSLSettings(ctx, source.Settings[0])
		diags.Append(settingsDiags...)
		target.Settings = settings
	}

	return target, diags
}

func transformSSLSettings(ctx context.Context, source SourceCustomHostnameSSLSettings) (*TargetCustomHostnameSSLSettingsModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	target := &TargetCustomHostnameSSLSettingsModel{
		EarlyHints:    migrations.FalseyStringToNull(source.EarlyHints),
		HTTP2:         migrations.FalseyStringToNull(source.HTTP2),
		MinTLSVersion: migrations.FalseyStringToNull(source.MinTLSVersion),
		TLS1_3:        migrations.FalseyStringToNull(source.TLS13),
	}

	if !source.Ciphers.IsNull() && !source.Ciphers.IsUnknown() {
		var ciphers []string
		diags.Append(source.Ciphers.ElementsAs(ctx, &ciphers, false)...)
		if !diags.HasError() {
			ciphersValues := make([]types.String, 0, len(ciphers))
			for _, cipher := range ciphers {
				ciphersValues = append(ciphersValues, types.StringValue(cipher))
			}
			target.Ciphers = &ciphersValues
		}
	}

	return target, diags
}

func transformCustomMetadata(ctx context.Context, source types.Map) (*map[string]types.String, diag.Diagnostics) {
	var diags diag.Diagnostics

	if source.IsNull() || source.IsUnknown() {
		return nil, diags
	}

	var metadata map[string]string
	diags.Append(source.ElementsAs(ctx, &metadata, false)...)
	if diags.HasError() {
		return nil, diags
	}

	result := make(map[string]types.String, len(metadata))
	for key, value := range metadata {
		result[key] = types.StringValue(value)
	}

	return &result, diags
}

func transformVerificationErrors(ctx context.Context, sourceSSL []SourceCustomHostnameSSL) (customfield.List[types.String], diag.Diagnostics) {
	var diags diag.Diagnostics

	if len(sourceSSL) == 0 || len(sourceSSL[0].ValidationErrors) == 0 {
		return customfield.NullList[types.String](ctx), diags
	}

	errors := make([]types.String, 0, len(sourceSSL[0].ValidationErrors))
	for _, validationErr := range sourceSSL[0].ValidationErrors {
		if validationErr.Message.IsNull() || validationErr.Message.IsUnknown() {
			continue
		}
		errors = append(errors, validationErr.Message)
	}

	if len(errors) == 0 {
		return customfield.NullList[types.String](ctx), diags
	}

	result, listDiags := customfield.NewList[types.String](ctx, errors)
	diags.Append(listDiags...)
	if diags.HasError() {
		return customfield.NullList[types.String](ctx), diags
	}

	return result, diags
}

func transformOwnershipVerification(ctx context.Context, source types.Map) (customfield.NestedObject[TargetCustomHostnameOwnershipVerification], diag.Diagnostics) {
	var diags diag.Diagnostics

	if source.IsNull() || source.IsUnknown() {
		return customfield.NullObject[TargetCustomHostnameOwnershipVerification](ctx), diags
	}

	var values map[string]string
	diags.Append(source.ElementsAs(ctx, &values, false)...)
	if diags.HasError() {
		return customfield.NullObject[TargetCustomHostnameOwnershipVerification](ctx), diags
	}

	ov := &TargetCustomHostnameOwnershipVerification{
		Name:  types.StringNull(),
		Type:  types.StringNull(),
		Value: types.StringNull(),
	}

	if value, ok := values["name"]; ok {
		ov.Name = types.StringValue(value)
	}
	if value, ok := values["type"]; ok {
		ov.Type = types.StringValue(value)
	}
	if value, ok := values["value"]; ok {
		ov.Value = types.StringValue(value)
	}

	result, objectDiags := customfield.NewObject(ctx, ov)
	diags.Append(objectDiags...)
	if diags.HasError() {
		return customfield.NullObject[TargetCustomHostnameOwnershipVerification](ctx), diags
	}

	return result, diags
}

func transformOwnershipVerificationHTTP(ctx context.Context, source types.Map) (customfield.NestedObject[TargetCustomHostnameOwnershipVerificationHTTP], diag.Diagnostics) {
	var diags diag.Diagnostics

	if source.IsNull() || source.IsUnknown() {
		return customfield.NullObject[TargetCustomHostnameOwnershipVerificationHTTP](ctx), diags
	}

	var values map[string]string
	diags.Append(source.ElementsAs(ctx, &values, false)...)
	if diags.HasError() {
		return customfield.NullObject[TargetCustomHostnameOwnershipVerificationHTTP](ctx), diags
	}

	ov := &TargetCustomHostnameOwnershipVerificationHTTP{
		HTTPBody: types.StringNull(),
		HTTPURL:  types.StringNull(),
	}

	if value, ok := values["http_body"]; ok {
		ov.HTTPBody = types.StringValue(value)
	}
	if value, ok := values["http_url"]; ok {
		ov.HTTPURL = types.StringValue(value)
	}

	result, objectDiags := customfield.NewObject(ctx, ov)
	diags.Append(objectDiags...)
	if diags.HasError() {
		return customfield.NullObject[TargetCustomHostnameOwnershipVerificationHTTP](ctx), diags
	}

	return result, diags
}

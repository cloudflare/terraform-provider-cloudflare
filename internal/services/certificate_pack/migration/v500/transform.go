package v500

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Transform converts source (legacy v4) state to target (current v5) state.
// This function handles the complete transformation from v4 SDKv2 format to v5 Plugin Framework format.
func Transform(ctx context.Context, source SourceCloudflareCertificatePackModel) (*TargetCertificatePackModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	// Step 1: Validate required fields
	if source.ZoneID.IsNull() || source.ZoneID.IsUnknown() {
		diags.AddError(
			"Missing required field",
			"zone_id is required for certificate_pack migration. The source state is missing this field, which indicates corrupted state.",
		)
		return nil, diags
	}
	if source.Type.IsNull() || source.Type.IsUnknown() {
		diags.AddError(
			"Missing required field",
			"type is required for certificate_pack migration. The source state is missing this field, which indicates corrupted state.",
		)
		return nil, diags
	}

	// Step 2: Initialize target with direct field copies
	target := &TargetCertificatePackModel{
		ID:                   source.ID,
		ZoneID:               source.ZoneID,
		Type:                 source.Type,
		ValidationMethod:     source.ValidationMethod,
		ValidityDays:         source.ValidityDays,
		CertificateAuthority: source.CertificateAuthority,
		CloudflareBranding:   source.CloudflareBranding,
	}

	// Step 3: Convert Hosts from types.Set to customfield.Set
	target.Hosts = convertSetToCustomFieldSet(ctx, source.Hosts, &diags)

	// Step 4: Transform validation_records (remove cname_target and cname_name from items)
	target.ValidationRecords = transformValidationRecords(ctx, source.ValidationRecords, &diags)

	// Step 5: Transform validation_errors (structure is compatible, just convert to customfield list)
	target.ValidationErrors = transformValidationErrors(ctx, source.ValidationErrors, &diags)

	// Step 6: Set new computed fields to null (will be refreshed from API)
	target.PrimaryCertificate = types.StringNull()
	target.Status = types.StringNull()
	target.Certificates = customfield.NullObjectList[TargetCertificatesModel](ctx)

	return target, diags
}

// convertSetToCustomFieldSet converts a types.Set to a customfield.Set[types.String].
func convertSetToCustomFieldSet(ctx context.Context, source types.Set, diags *diag.Diagnostics) customfield.Set[types.String] {
	if source.IsNull() {
		return customfield.NullSet[types.String](ctx)
	}
	if source.IsUnknown() {
		return customfield.UnknownSet[types.String](ctx)
	}

	// Extract elements as strings first
	var stringValues []string
	diags.Append(source.ElementsAs(ctx, &stringValues, false)...)
	if diags.HasError() {
		return customfield.NullSet[types.String](ctx)
	}

	// Convert each string to types.StringValue
	elements := make([]attr.Value, len(stringValues))
	for i, s := range stringValues {
		elements[i] = types.StringValue(s)
	}

	return customfield.NewSetMust[types.String](ctx, elements)
}

// transformValidationRecords converts v4 validation_records to v5 format.
// This removes cname_target and cname_name fields from each item.
func transformValidationRecords(ctx context.Context, source []SourceValidationRecordsModel, diags *diag.Diagnostics) customfield.NestedObjectList[TargetValidationRecordsModel] {
	if source == nil || len(source) == 0 {
		// Return empty list (not null) to match tf-migrate behavior
		return customfield.NewObjectListMust(ctx, []TargetValidationRecordsModel{})
	}

	// Transform each item
	targetRecords := make([]TargetValidationRecordsModel, 0, len(source))
	for _, sourceRecord := range source {
		targetRecord := TargetValidationRecordsModel{
			TXTName:  sourceRecord.TXTName,
			TXTValue: sourceRecord.TXTValue,
			HTTPURL:  sourceRecord.HTTPURL,
			HTTPBody: sourceRecord.HTTPBody,
			Emails:   convertListToCustomFieldList(ctx, sourceRecord.Emails, diags),
		}
		targetRecords = append(targetRecords, targetRecord)
	}

	return customfield.NewObjectListMust(ctx, targetRecords)
}

// transformValidationErrors converts v4 validation_errors to v5 format.
// Structure is compatible, just needs conversion to customfield list.
func transformValidationErrors(ctx context.Context, source []SourceValidationErrorsModel, diags *diag.Diagnostics) customfield.NestedObjectList[TargetValidationErrorsModel] {
	if source == nil || len(source) == 0 {
		// Return empty list (not null) to match tf-migrate behavior
		return customfield.NewObjectListMust(ctx, []TargetValidationErrorsModel{})
	}

	// Transform each item (structure is same, just copy)
	targetErrors := make([]TargetValidationErrorsModel, 0, len(source))
	for _, sourceError := range source {
		targetError := TargetValidationErrorsModel{
			Message: sourceError.Message,
		}
		targetErrors = append(targetErrors, targetError)
	}

	return customfield.NewObjectListMust(ctx, targetErrors)
}

// convertListToCustomFieldList converts a types.List to a customfield.List[types.String].
func convertListToCustomFieldList(ctx context.Context, source types.List, diags *diag.Diagnostics) customfield.List[types.String] {
	if source.IsNull() {
		return customfield.NullList[types.String](ctx)
	}
	if source.IsUnknown() {
		return customfield.UnknownList[types.String](ctx)
	}

	// Extract elements as strings first
	var stringValues []string
	diags.Append(source.ElementsAs(ctx, &stringValues, false)...)
	if diags.HasError() {
		return customfield.NullList[types.String](ctx)
	}

	// Convert each string to types.StringValue
	elements := make([]attr.Value, len(stringValues))
	for i, s := range stringValues {
		elements[i] = types.StringValue(s)
	}

	return customfield.NewListMust[types.String](ctx, elements)
}

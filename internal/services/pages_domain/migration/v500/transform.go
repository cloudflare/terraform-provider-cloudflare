package v500

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Transform converts source (legacy v4) state to target (current v5) state.
//
// Transformations performed:
// 1. Field rename: domain → name
// 2. All other fields are direct pass-through
// 3. New computed fields are initialized to Null (will be populated by API)
func Transform(ctx context.Context, source SourcePagesDomainModel) (*TargetPagesDomainModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	// Validate required fields
	if source.AccountID.IsNull() || source.AccountID.IsUnknown() {
		diags.AddError(
			"Missing required field",
			"account_id is required for pages_domain migration. The source state is missing this field, which indicates corrupted state.",
		)
		return nil, diags
	}
	if source.Domain.IsNull() || source.Domain.IsUnknown() {
		diags.AddError(
			"Missing required field",
			"domain is required for pages_domain migration. The source state is missing this field, which indicates corrupted state.",
		)
		return nil, diags
	}
	if source.ProjectName.IsNull() || source.ProjectName.IsUnknown() {
		diags.AddError(
			"Missing required field",
			"project_name is required for pages_domain migration. The source state is missing this field, which indicates corrupted state.",
		)
		return nil, diags
	}

	// Initialize target with transformations
	target := &TargetPagesDomainModel{
		// Direct pass-through fields
		AccountID:   source.AccountID,
		ProjectName: source.ProjectName,
		Status:      source.Status,

		// Field rename: domain → name
		Name: source.Domain,

		// New computed fields in v5 - initialize to Null, will be populated by API
		ID:                   types.StringNull(),
		CertificateAuthority: types.StringNull(),
		CreatedOn:            types.StringNull(),
		DomainID:             types.StringNull(),
		ZoneTag:              types.StringNull(),
		ValidationData:       customfield.NullObject[TargetPagesDomainValidationDataModel](ctx),
		VerificationData:     customfield.NullObject[TargetPagesDomainVerificationDataModel](ctx),
	}

	return target, diags
}

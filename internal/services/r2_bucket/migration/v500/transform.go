package v500

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Transform converts source (v4) state to target (v5) state.
//
// For r2_bucket, this transformation:
// 1. Copies all 4 existing fields unchanged (id, name, account_id, location)
// 2. Sets new Optional+Computed fields to their defaults (jurisdiction="default", storage_class="Standard")
//   - This matches what the API returns for existing buckets
//   - Prevents drift on first plan after migration
//
// 3. Sets computed field to null (creation_date - will be populated by API)
func Transform(ctx context.Context, source SourceR2BucketModel) (*TargetR2BucketModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	// Validate required fields
	if source.Name.IsNull() || source.Name.IsUnknown() {
		diags.AddError(
			"Missing required field",
			"name is required for r2_bucket migration. The source state is missing this field, which indicates corrupted state.",
		)
		return nil, diags
	}
	if source.AccountID.IsNull() || source.AccountID.IsUnknown() {
		diags.AddError(
			"Missing required field",
			"account_id is required for r2_bucket migration. The source state is missing this field, which indicates corrupted state.",
		)
		return nil, diags
	}

	// Initialize target with direct copies from v4
	target := &TargetR2BucketModel{
		ID:        source.ID,
		Name:      source.Name,
		AccountID: source.AccountID,
		Location:  source.Location,
	}

	// Set new v5 Optional+Computed fields to their API defaults
	// This matches what the Cloudflare API returns for existing buckets
	// and prevents drift on first plan after migration
	target.Jurisdiction = types.StringValue("default")
	target.StorageClass = types.StringValue("Standard")

	// creation_date: computed field, leave null - API will populate on first refresh
	target.CreationDate = types.StringNull()

	return target, diags
}

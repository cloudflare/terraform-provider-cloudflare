package v500

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
)

// Transform converts source (legacy v4) state to target (current v5) state.
// This function handles the key → key_name field rename and adds the new metadata field.
func Transform(ctx context.Context, source SourceCloudflareWorkersKVModel) (*TargetWorkersKVModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	// Validate required fields
	if source.AccountID.IsNull() || source.AccountID.IsUnknown() {
		diags.AddError(
			"Missing required field",
			"account_id is required for workers_kv migration. The source state is missing this field, which indicates corrupted state.",
		)
		return nil, diags
	}
	if source.NamespaceID.IsNull() || source.NamespaceID.IsUnknown() {
		diags.AddError(
			"Missing required field",
			"namespace_id is required for workers_kv migration. The source state is missing this field, which indicates corrupted state.",
		)
		return nil, diags
	}
	if source.Key.IsNull() || source.Key.IsUnknown() {
		diags.AddError(
			"Missing required field",
			"key is required for workers_kv migration. The source state is missing this field, which indicates corrupted state.",
		)
		return nil, diags
	}
	if source.Value.IsNull() || source.Value.IsUnknown() {
		diags.AddError(
			"Missing required field",
			"value is required for workers_kv migration. The source state is missing this field, which indicates corrupted state.",
		)
		return nil, diags
	}

	// Create target with field transformations
	target := &TargetWorkersKVModel{
		// Direct copies (same field name and type)
		ID:          source.ID,
		AccountID:   source.AccountID,
		NamespaceID: source.NamespaceID,
		Value:       source.Value,

		// Rename: key → key_name
		KeyName: source.Key,

		// New field in v5: metadata (set to null for migrated resources)
		Metadata: jsontypes.NewNormalizedNull(),
	}

	return target, diags
}

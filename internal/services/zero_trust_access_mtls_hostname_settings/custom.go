package zero_trust_access_mtls_hostname_settings

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/importpath"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// importIDFormat is the import ID format for this resource. Account IDs and zone
// IDs are both 32-character hex strings, so the scope cannot be inferred from the
// ID alone and is carried by an explicit accounts/zones discriminator segment.
const importIDFormat = "<{accounts|zones}/{account_id|zone_id}>"

// resolveImportScope resolves an import ID into exactly one of account or zone
// scope. The returned values are null when the ID is malformed or carries an
// unrecognized discriminator.
func resolveImportScope(importID string) (accountID, zoneID types.String, diags diag.Diagnostics) {
	scope, scopedValue := "", ""
	diags.Append(importpath.ParseImportID(importID, importIDFormat, &scope, &scopedValue)...)
	if diags.HasError() {
		return
	}

	switch scope {
	case "accounts":
		accountID = types.StringValue(scopedValue)
	case "zones":
		zoneID = types.StringValue(scopedValue)
	default:
		diags.AddError(
			"invalid discriminator segment - "+importIDFormat,
			"expected discriminator to be one of {accounts|zones}",
		)
	}
	return
}

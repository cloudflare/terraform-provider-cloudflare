package leaked_credential_check

import "github.com/hashicorp/terraform-plugin-framework/types"

type LeakedCredentialCheckModel struct {
	ZoneID  types.String `tfsdk:"zone_id"`
	Enabled types.Bool   `tfsdk:"enabled"`
}

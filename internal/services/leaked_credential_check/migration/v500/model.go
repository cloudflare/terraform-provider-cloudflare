package v500

import "github.com/hashicorp/terraform-plugin-framework/types"

type SourceLeakedCredentialCheckModel struct {
	ZoneID  types.String `tfsdk:"zone_id"`
	Enabled types.Bool   `tfsdk:"enabled"`
}

type TargetLeakedCredentialCheckModel struct {
	ZoneID  types.String `tfsdk:"zone_id"`
	Enabled types.Bool   `tfsdk:"enabled"`
}

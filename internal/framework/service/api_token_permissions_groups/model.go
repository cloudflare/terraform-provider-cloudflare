package api_token_permissions_groups

import "github.com/hashicorp/terraform-plugin-framework/types"

// APITokenPermissionsGroupModel describes the data source data model for api token permissions.
type APITokenPermissionsGroupModel struct {
	Permissions types.Map `tfsdk:"permissions"`
	Zone        types.Map `tfsdk:"zone"`
	Account     types.Map `tfsdk:"account"`
	User        types.Map `tfsdk:"user"`
	R2          types.Map `tfsdk:"r2"`
}

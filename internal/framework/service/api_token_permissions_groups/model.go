package api_token_permissions_groups

import "github.com/hashicorp/terraform-plugin-framework/types"

// APITokenPermissionsGroupModel describes the data source data model for api token permissions.
type APITokenPermissionsGroupModel struct {
	ID          types.String            `tfsdk:"id"`
	Permissions map[string]types.String `tfsdk:"permissions"`
	Zone        map[string]types.String `tfsdk:"zone"`
	Account     map[string]types.String `tfsdk:"account"`
	User        map[string]types.String `tfsdk:"user"`
	R2          map[string]types.String `tfsdk:"r2"`
}

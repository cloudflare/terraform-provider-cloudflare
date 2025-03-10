// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package api_token_permission_groups

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type APITokenPermissionGroupsListResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[APITokenPermissionGroupsListResultDataSourceModel] `json:"result,computed"`
}

type APITokenPermissionGroupsListDataSourceModel struct {
	MaxItems types.Int64                                                                     `tfsdk:"max_items"`
	Result   customfield.NestedObjectList[APITokenPermissionGroupsListResultDataSourceModel] `tfsdk:"result"`
}

type APITokenPermissionGroupsListResultDataSourceModel struct {
	ID     types.String                   `tfsdk:"id" json:"id,computed"`
	Name   types.String                   `tfsdk:"name" json:"name,computed"`
	Scopes customfield.List[types.String] `tfsdk:"scopes" json:"scopes,computed"`
}

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package oauth_scope

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type OAuthScopesResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[OAuthScopesResultDataSourceModel] `json:"result,computed"`
}

type OAuthScopesDataSourceModel struct {
	MaxItems types.Int64                                                    `tfsdk:"max_items"`
	Result   customfield.NestedObjectList[OAuthScopesResultDataSourceModel] `tfsdk:"result"`
}

type OAuthScopesResultDataSourceModel struct {
	ID       types.String                   `tfsdk:"id" json:"id,computed"`
	Name     types.String                   `tfsdk:"name" json:"name,computed"`
	Category types.String                   `tfsdk:"category" json:"category,computed"`
	Scopes   customfield.List[types.String] `tfsdk:"scopes" json:"scopes,computed"`
}

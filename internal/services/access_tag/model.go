// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package access_tag

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AccessTagResultEnvelope struct {
	Result AccessTagModel `json:"result,computed"`
}

type AccessTagResultDataSourceEnvelope struct {
	Result AccessTagDataSourceModel `json:"result,computed"`
}

type AccessTagsResultDataSourceEnvelope struct {
	Result AccessTagsDataSourceModel `json:"result,computed"`
}

type AccessTagModel struct {
	ID        types.String `tfsdk:"id" json:"-,computed"`
	AccountID types.String `tfsdk:"account_id" path:"account_id"`
	Name      types.String `tfsdk:"name" json:"name"`
	CreatedAt types.String `tfsdk:"created_at" json:"created_at"`
	UpdatedAt types.String `tfsdk:"updated_at" json:"updated_at"`
	AppCount  types.Int64  `tfsdk:"app_count" json:"app_count,computed"`
}

type AccessTagDataSourceModel struct {
}

type AccessTagsDataSourceModel struct {
}

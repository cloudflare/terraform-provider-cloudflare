// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package access_custom_page

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AccessCustomPagesResultListDataSourceEnvelope struct {
	Result *[]*AccessCustomPagesItemsDataSourceModel `json:"result,computed"`
}

type AccessCustomPagesDataSourceModel struct {
	AccountID types.String                              `tfsdk:"account_id" path:"account_id"`
	MaxItems  types.Int64                               `tfsdk:"max_items"`
	Items     *[]*AccessCustomPagesItemsDataSourceModel `tfsdk:"items"`
}

type AccessCustomPagesItemsDataSourceModel struct {
	Name      types.String      `tfsdk:"name" json:"name,computed"`
	Type      types.String      `tfsdk:"type" json:"type,computed"`
	AppCount  types.Int64       `tfsdk:"app_count" json:"app_count"`
	CreatedAt timetypes.RFC3339 `tfsdk:"created_at" json:"created_at"`
	UID       types.String      `tfsdk:"uid" json:"uid"`
	UpdatedAt timetypes.RFC3339 `tfsdk:"updated_at" json:"updated_at"`
}

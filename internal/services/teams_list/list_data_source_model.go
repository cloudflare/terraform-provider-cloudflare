// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package teams_list

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type TeamsListsResultListDataSourceEnvelope struct {
	Result *[]*TeamsListsResultDataSourceModel `json:"result,computed"`
}

type TeamsListsDataSourceModel struct {
	AccountID types.String                        `tfsdk:"account_id" path:"account_id"`
	Type      types.String                        `tfsdk:"type" query:"type"`
	MaxItems  types.Int64                         `tfsdk:"max_items"`
	Result    *[]*TeamsListsResultDataSourceModel `tfsdk:"result"`
}

type TeamsListsResultDataSourceModel struct {
	ID          types.String      `tfsdk:"id" json:"id"`
	ListCount   types.Float64     `tfsdk:"list_count" json:"count,computed"`
	CreatedAt   timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed"`
	Description types.String      `tfsdk:"description" json:"description"`
	Name        types.String      `tfsdk:"name" json:"name"`
	Type        types.String      `tfsdk:"type" json:"type"`
	UpdatedAt   timetypes.RFC3339 `tfsdk:"updated_at" json:"updated_at,computed"`
}

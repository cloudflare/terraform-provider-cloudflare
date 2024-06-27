// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package teams_list

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type TeamsListResultEnvelope struct {
	Result TeamsListModel `json:"result,computed"`
}

type TeamsListModel struct {
	ID          types.String            `tfsdk:"id" json:"id,computed"`
	AccountID   types.String            `tfsdk:"account_id" path:"account_id"`
	Name        types.String            `tfsdk:"name" json:"name"`
	Type        types.String            `tfsdk:"type" json:"type"`
	Description types.String            `tfsdk:"description" json:"description"`
	Items       *[]*TeamsListItemsModel `tfsdk:"items" json:"items"`
	CreatedAt   types.String            `tfsdk:"created_at" json:"created_at,computed"`
	UpdatedAt   types.String            `tfsdk:"updated_at" json:"updated_at,computed"`
	Count       types.Float64           `tfsdk:"count" json:"count,computed"`
}

type TeamsListItemsModel struct {
	CreatedAt   types.String `tfsdk:"created_at" json:"created_at"`
	Description types.String `tfsdk:"description" json:"description"`
	Value       types.String `tfsdk:"value" json:"value"`
}

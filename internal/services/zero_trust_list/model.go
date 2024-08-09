// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_list

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustListResultEnvelope struct {
	Result ZeroTrustListModel `json:"result,computed"`
}

type ZeroTrustListModel struct {
	ID          types.String                `tfsdk:"id" json:"id,computed"`
	AccountID   types.String                `tfsdk:"account_id" path:"account_id"`
	Type        types.String                `tfsdk:"type" json:"type"`
	Items       *[]*ZeroTrustListItemsModel `tfsdk:"items" json:"items"`
	Name        types.String                `tfsdk:"name" json:"name"`
	Description types.String                `tfsdk:"description" json:"description"`
	CreatedAt   timetypes.RFC3339           `tfsdk:"created_at" json:"created_at,computed"`
	ListCount   types.Float64               `tfsdk:"list_count" json:"count,computed"`
	UpdatedAt   timetypes.RFC3339           `tfsdk:"updated_at" json:"updated_at,computed"`
}

type ZeroTrustListItemsModel struct {
	CreatedAt   timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed"`
	Description types.String      `tfsdk:"description" json:"description"`
	Value       types.String      `tfsdk:"value" json:"value"`
}

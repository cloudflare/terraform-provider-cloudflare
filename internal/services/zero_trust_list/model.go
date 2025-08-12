// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_list

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustListResultEnvelope struct {
	Result ZeroTrustListModel `json:"result"`
}

type ZeroTrustListModel struct {
	ID          types.String                `tfsdk:"id" json:"id,computed"`
	AccountID   types.String                `tfsdk:"account_id" path:"account_id,required"`
	Type        types.String                `tfsdk:"type" json:"type,required"`
	Name        types.String                `tfsdk:"name" json:"name,required"`
	Items       *[]*ZeroTrustListItemsModel `tfsdk:"items" json:"items,optional"`
	Description types.String                `tfsdk:"description" json:"description,computed_optional"`
	CreatedAt   timetypes.RFC3339           `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	ListCount   types.Float64               `tfsdk:"list_count" json:"count,computed"`
	UpdatedAt   timetypes.RFC3339           `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
}

func (m ZeroTrustListModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m ZeroTrustListModel) MarshalJSONForUpdate(state ZeroTrustListModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type ZeroTrustListItemsModel struct {
	Description types.String `tfsdk:"description" json:"description,optional"`
	Value       types.String `tfsdk:"value" json:"value,optional"`
}

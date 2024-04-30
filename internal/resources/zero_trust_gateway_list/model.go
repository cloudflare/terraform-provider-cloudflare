// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_gateway_list

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustGatewayListResultEnvelope struct {
	Result ZeroTrustGatewayListModel `json:"result,computed"`
}

type ZeroTrustGatewayListModel struct {
	ID          types.String                       `tfsdk:"id" json:"id,computed"`
	AccountID   types.String                       `tfsdk:"account_id" path:"account_id"`
	Name        types.String                       `tfsdk:"name" json:"name"`
	Type        types.String                       `tfsdk:"type" json:"type"`
	Description types.String                       `tfsdk:"description" json:"description"`
	Items       *[]*ZeroTrustGatewayListItemsModel `tfsdk:"items" json:"items"`
}

type ZeroTrustGatewayListItemsModel struct {
	CreatedAt types.String `tfsdk:"created_at" json:"created_at,computed"`
	Value     types.String `tfsdk:"value" json:"value"`
}

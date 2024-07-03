// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package access_ca_certificate

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AccessCACertificatesResultListDataSourceEnvelope struct {
	Result *[]*AccessCACertificatesItemsDataSourceModel `json:"result,computed"`
}

type AccessCACertificatesDataSourceModel struct {
	AccountID types.String                                 `tfsdk:"account_id" path:"account_id"`
	ZoneID    types.String                                 `tfsdk:"zone_id" path:"zone_id"`
	MaxItems  types.Int64                                  `tfsdk:"max_items"`
	Items     *[]*AccessCACertificatesItemsDataSourceModel `tfsdk:"items"`
}

type AccessCACertificatesItemsDataSourceModel struct {
	ID        types.String `tfsdk:"id" json:"id,computed"`
	AUD       types.String `tfsdk:"aud" json:"aud,computed"`
	PublicKey types.String `tfsdk:"public_key" json:"public_key,computed"`
}

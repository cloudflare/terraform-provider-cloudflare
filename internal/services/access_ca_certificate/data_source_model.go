// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package access_ca_certificate

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AccessCACertificateResultDataSourceEnvelope struct {
	Result AccessCACertificateDataSourceModel `json:"result,computed"`
}

type AccessCACertificateResultListDataSourceEnvelope struct {
	Result *[]*AccessCACertificateDataSourceModel `json:"result,computed"`
}

type AccessCACertificateDataSourceModel struct {
	AppID     types.String `tfsdk:"app_id" path:"app_id"`
	AccountID types.String `tfsdk:"account_id" path:"account_id"`
	ZoneID    types.String `tfsdk:"zone_id" path:"zone_id"`
	ID        types.String `tfsdk:"id" json:"id"`
	AUD       types.String `tfsdk:"aud" json:"aud"`
	PublicKey types.String `tfsdk:"public_key" json:"public_key"`
}

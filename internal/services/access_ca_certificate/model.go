// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package access_ca_certificate

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AccessCACertificateResultEnvelope struct {
	Result AccessCACertificateModel `json:"result,computed"`
}

type AccessCACertificateModel struct {
	ID        types.String `tfsdk:"id" json:"-,computed"`
	AppID     types.String `tfsdk:"app_id" path:"app_id"`
	AccountID types.String `tfsdk:"account_id" path:"account_id"`
	ZoneID    types.String `tfsdk:"zone_id" path:"zone_id"`
}

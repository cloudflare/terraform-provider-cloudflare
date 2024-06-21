// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zone_dnssec

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZoneDNSSECResultEnvelope struct {
	Result ZoneDNSSECModel `json:"result,computed"`
}

type ZoneDNSSECModel struct {
	ZoneID            types.String `tfsdk:"zone_id" path:"zone_id"`
	DNSSECMultiSigner types.Bool   `tfsdk:"dnssec_multi_signer" json:"dnssec_multi_signer"`
	DNSSECPresigned   types.Bool   `tfsdk:"dnssec_presigned" json:"dnssec_presigned"`
	Status            types.String `tfsdk:"status" json:"status"`
}

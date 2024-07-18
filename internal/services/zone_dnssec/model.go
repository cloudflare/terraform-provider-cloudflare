// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zone_dnssec

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZoneDNSSECResultEnvelope struct {
	Result ZoneDNSSECModel `json:"result,computed"`
}

type ZoneDNSSECModel struct {
	ZoneID            types.String  `tfsdk:"zone_id" path:"zone_id"`
	DNSSECMultiSigner types.Bool    `tfsdk:"dnssec_multi_signer" json:"dnssec_multi_signer"`
	DNSSECPresigned   types.Bool    `tfsdk:"dnssec_presigned" json:"dnssec_presigned"`
	Status            types.String  `tfsdk:"status" json:"status"`
	Algorithm         types.String  `tfsdk:"algorithm" json:"algorithm,computed"`
	Digest            types.String  `tfsdk:"digest" json:"digest,computed"`
	DigestAlgorithm   types.String  `tfsdk:"digest_algorithm" json:"digest_algorithm,computed"`
	DigestType        types.String  `tfsdk:"digest_type" json:"digest_type,computed"`
	DS                types.String  `tfsdk:"ds" json:"ds,computed"`
	Flags             types.Float64 `tfsdk:"flags" json:"flags,computed"`
	KeyTag            types.Float64 `tfsdk:"key_tag" json:"key_tag,computed"`
	KeyType           types.String  `tfsdk:"key_type" json:"key_type,computed"`
	ModifiedOn        types.String  `tfsdk:"modified_on" json:"modified_on,computed"`
	PublicKey         types.String  `tfsdk:"public_key" json:"public_key,computed"`
}

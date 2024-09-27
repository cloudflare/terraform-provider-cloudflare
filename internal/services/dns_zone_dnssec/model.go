// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dns_zone_dnssec

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type DNSZoneDNSSECResultEnvelope struct {
	Result DNSZoneDNSSECModel `json:"result"`
}

type DNSZoneDNSSECModel struct {
	ID                types.String      `tfsdk:"id" json:"-,computed"`
	ZoneID            types.String      `tfsdk:"zone_id" path:"zone_id,required"`
	DNSSECMultiSigner types.Bool        `tfsdk:"dnssec_multi_signer" json:"dnssec_multi_signer,optional"`
	DNSSECPresigned   types.Bool        `tfsdk:"dnssec_presigned" json:"dnssec_presigned,optional"`
	Status            types.String      `tfsdk:"status" json:"status,optional"`
	Algorithm         types.String      `tfsdk:"algorithm" json:"algorithm,computed"`
	Digest            types.String      `tfsdk:"digest" json:"digest,computed"`
	DigestAlgorithm   types.String      `tfsdk:"digest_algorithm" json:"digest_algorithm,computed"`
	DigestType        types.String      `tfsdk:"digest_type" json:"digest_type,computed"`
	DS                types.String      `tfsdk:"ds" json:"ds,computed"`
	Flags             types.Float64     `tfsdk:"flags" json:"flags,computed"`
	KeyTag            types.Float64     `tfsdk:"key_tag" json:"key_tag,computed"`
	KeyType           types.String      `tfsdk:"key_type" json:"key_type,computed"`
	ModifiedOn        timetypes.RFC3339 `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
	PublicKey         types.String      `tfsdk:"public_key" json:"public_key,computed"`
}

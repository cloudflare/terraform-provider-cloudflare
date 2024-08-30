// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dns_zone_dnssec

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/dnssec"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type DNSZoneDNSSECResultDataSourceEnvelope struct {
	Result DNSZoneDNSSECDataSourceModel `json:"result,computed"`
}

type DNSZoneDNSSECDataSourceModel struct {
	ZoneID            types.String      `tfsdk:"zone_id" path:"zone_id"`
	Algorithm         types.String      `tfsdk:"algorithm" json:"algorithm"`
	Digest            types.String      `tfsdk:"digest" json:"digest"`
	DigestAlgorithm   types.String      `tfsdk:"digest_algorithm" json:"digest_algorithm"`
	DigestType        types.String      `tfsdk:"digest_type" json:"digest_type"`
	DNSSECMultiSigner types.Bool        `tfsdk:"dnssec_multi_signer" json:"dnssec_multi_signer"`
	DNSSECPresigned   types.Bool        `tfsdk:"dnssec_presigned" json:"dnssec_presigned"`
	DS                types.String      `tfsdk:"ds" json:"ds"`
	Flags             types.Float64     `tfsdk:"flags" json:"flags"`
	KeyTag            types.Float64     `tfsdk:"key_tag" json:"key_tag"`
	KeyType           types.String      `tfsdk:"key_type" json:"key_type"`
	ModifiedOn        timetypes.RFC3339 `tfsdk:"modified_on" json:"modified_on" format:"date-time"`
	PublicKey         types.String      `tfsdk:"public_key" json:"public_key"`
	Status            types.String      `tfsdk:"status" json:"status"`
}

func (m *DNSZoneDNSSECDataSourceModel) toReadParams(_ context.Context) (params dnssec.DNSSECGetParams, diags diag.Diagnostics) {
	params = dnssec.DNSSECGetParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	return
}

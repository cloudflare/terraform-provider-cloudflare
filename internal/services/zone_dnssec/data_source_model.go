// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zone_dnssec

import (
  "context"

  "github.com/cloudflare/cloudflare-go/v4"
  "github.com/cloudflare/cloudflare-go/v4/dns"
  "github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
  "github.com/hashicorp/terraform-plugin-framework/diag"
  "github.com/hashicorp/terraform-plugin-framework/types"
)

type ZoneDNSSECResultDataSourceEnvelope struct {
Result ZoneDNSSECDataSourceModel `json:"result,computed"`
}

type ZoneDNSSECDataSourceModel struct {
ZoneID types.String `tfsdk:"zone_id" path:"zone_id,required"`
Algorithm types.String `tfsdk:"algorithm" json:"algorithm,computed"`
Digest types.String `tfsdk:"digest" json:"digest,computed"`
DigestAlgorithm types.String `tfsdk:"digest_algorithm" json:"digest_algorithm,computed"`
DigestType types.String `tfsdk:"digest_type" json:"digest_type,computed"`
DNSSECMultiSigner types.Bool `tfsdk:"dnssec_multi_signer" json:"dnssec_multi_signer,computed"`
DNSSECPresigned types.Bool `tfsdk:"dnssec_presigned" json:"dnssec_presigned,computed"`
DS types.String `tfsdk:"ds" json:"ds,computed"`
Flags types.Float64 `tfsdk:"flags" json:"flags,computed"`
KeyTag types.Float64 `tfsdk:"key_tag" json:"key_tag,computed"`
KeyType types.String `tfsdk:"key_type" json:"key_type,computed"`
ModifiedOn timetypes.RFC3339 `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
PublicKey types.String `tfsdk:"public_key" json:"public_key,computed"`
Status types.String `tfsdk:"status" json:"status,computed"`
}

func (m *ZoneDNSSECDataSourceModel) toReadParams(_ context.Context) (params dns.DNSSECGetParams, diags diag.Diagnostics) {
  params = dns.DNSSECGetParams{
    ZoneID: cloudflare.F(m.ZoneID.ValueString()),
  }

  return
}

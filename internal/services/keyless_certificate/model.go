// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package keyless_certificate

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type KeylessCertificateResultEnvelope struct {
	Result KeylessCertificateModel `json:"result"`
}

type KeylessCertificateModel struct {
	ID           types.String                   `tfsdk:"id" json:"id,computed"`
	ZoneID       types.String                   `tfsdk:"zone_id" path:"zone_id,required"`
	Certificate  types.String                   `tfsdk:"certificate" json:"certificate,required"`
	BundleMethod types.String                   `tfsdk:"bundle_method" json:"bundle_method,computed_optional"`
	Host         types.String                   `tfsdk:"host" json:"host,required"`
	Enabled      types.Bool                     `tfsdk:"enabled" json:"enabled,optional"`
	Name         types.String                   `tfsdk:"name" json:"name,optional"`
	Tunnel       *KeylessCertificateTunnelModel `tfsdk:"tunnel" json:"tunnel,optional"`
	Port         types.Float64                  `tfsdk:"port" json:"port,computed_optional"`
	CreatedOn    timetypes.RFC3339              `tfsdk:"created_on" json:"created_on,computed" format:"date-time"`
	ModifiedOn   timetypes.RFC3339              `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
	Status       types.String                   `tfsdk:"status" json:"status,computed"`
	Permissions  customfield.List[types.String] `tfsdk:"permissions" json:"permissions,computed"`
}

func (m KeylessCertificateModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m KeylessCertificateModel) MarshalJSONForUpdate(state KeylessCertificateModel) (data []byte, err error) {
	return apijson.MarshalForPatch(m, state)
}

type KeylessCertificateTunnelModel struct {
	PrivateIP types.String `tfsdk:"private_ip" json:"private_ip,required"`
	VnetID    types.String `tfsdk:"vnet_id" json:"vnet_id,required"`
}

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package authenticated_origin_pulls

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AuthenticatedOriginPullsResultEnvelope struct {
	Result AuthenticatedOriginPullsModel `json:"result"`
}

type AuthenticatedOriginPullsModel struct {
	ZoneID         types.String                            `tfsdk:"zone_id" path:"zone_id,required"`
	Hostname       types.String                            `tfsdk:"hostname" path:"hostname,optional"`
	Config         *[]*AuthenticatedOriginPullsConfigModel `tfsdk:"config" json:"config,required,no_refresh"`
	CERTID         types.String                            `tfsdk:"cert_id" json:"cert_id,computed"`
	CERTStatus     types.String                            `tfsdk:"cert_status" json:"cert_status,computed"`
	CERTUpdatedAt  timetypes.RFC3339                       `tfsdk:"cert_updated_at" json:"cert_updated_at,computed" format:"date-time"`
	CERTUploadedOn timetypes.RFC3339                       `tfsdk:"cert_uploaded_on" json:"cert_uploaded_on,computed" format:"date-time"`
	Certificate    types.String                            `tfsdk:"certificate" json:"certificate,computed"`
	CreatedAt      timetypes.RFC3339                       `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Enabled        types.Bool                              `tfsdk:"enabled" json:"enabled,computed"`
	ExpiresOn      timetypes.RFC3339                       `tfsdk:"expires_on" json:"expires_on,computed" format:"date-time"`
	ID             types.String                            `tfsdk:"id" json:"id,computed,no_refresh"`
	Issuer         types.String                            `tfsdk:"issuer" json:"issuer,computed"`
	PrivateKey     types.String                            `tfsdk:"private_key" json:"private_key,computed,no_refresh"`
	SerialNumber   types.String                            `tfsdk:"serial_number" json:"serial_number,computed"`
	Signature      types.String                            `tfsdk:"signature" json:"signature,computed"`
	Status         types.String                            `tfsdk:"status" json:"status,computed"`
	UpdatedAt      timetypes.RFC3339                       `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
}

func (m AuthenticatedOriginPullsModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m AuthenticatedOriginPullsModel) MarshalJSONForUpdate(state AuthenticatedOriginPullsModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type AuthenticatedOriginPullsConfigModel struct {
	CERTID   types.String `tfsdk:"cert_id" json:"cert_id,optional"`
	Enabled  types.Bool   `tfsdk:"enabled" json:"enabled,optional"`
	Hostname types.String `tfsdk:"hostname" json:"hostname,optional"`
}

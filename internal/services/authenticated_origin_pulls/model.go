// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package authenticated_origin_pulls

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AuthenticatedOriginPullsResultEnvelope struct {
	Result AuthenticatedOriginPullsModel `json:"result,computed"`
}

type AuthenticatedOriginPullsModel struct {
	ZoneID         types.String                            `tfsdk:"zone_id" path:"zone_id"`
	Hostname       types.String                            `tfsdk:"hostname" path:"hostname"`
	Config         *[]*AuthenticatedOriginPullsConfigModel `tfsdk:"config" json:"config"`
	CERTID         types.String                            `tfsdk:"cert_id" json:"cert_id,computed"`
	CERTStatus     types.String                            `tfsdk:"cert_status" json:"cert_status,computed"`
	CERTUpdatedAt  types.String                            `tfsdk:"cert_updated_at" json:"cert_updated_at,computed"`
	CERTUploadedOn types.String                            `tfsdk:"cert_uploaded_on" json:"cert_uploaded_on,computed"`
	Certificate    types.String                            `tfsdk:"certificate" json:"certificate,computed"`
	CreatedAt      types.String                            `tfsdk:"created_at" json:"created_at,computed"`
	Enabled        types.Bool                              `tfsdk:"enabled" json:"enabled,computed"`
	ExpiresOn      types.String                            `tfsdk:"expires_on" json:"expires_on,computed"`
	Issuer         types.String                            `tfsdk:"issuer" json:"issuer,computed"`
	SerialNumber   types.String                            `tfsdk:"serial_number" json:"serial_number,computed"`
	Signature      types.String                            `tfsdk:"signature" json:"signature,computed"`
	Status         types.String                            `tfsdk:"status" json:"status,computed"`
	UpdatedAt      types.String                            `tfsdk:"updated_at" json:"updated_at,computed"`
}

type AuthenticatedOriginPullsConfigModel struct {
	CERTID   types.String `tfsdk:"cert_id" json:"cert_id"`
	Enabled  types.Bool   `tfsdk:"enabled" json:"enabled"`
	Hostname types.String `tfsdk:"hostname" json:"hostname"`
}

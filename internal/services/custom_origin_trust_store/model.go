// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package custom_origin_trust_store

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type CustomOriginTrustStoreResultEnvelope struct {
	Result CustomOriginTrustStoreModel `json:"result"`
}

type CustomOriginTrustStoreModel struct {
	ID          types.String      `tfsdk:"id" json:"id,computed"`
	ZoneID      types.String      `tfsdk:"zone_id" path:"zone_id,required"`
	Certificate types.String      `tfsdk:"certificate" json:"certificate,required"`
	ExpiresOn   timetypes.RFC3339 `tfsdk:"expires_on" json:"expires_on,computed" format:"date-time"`
	Issuer      types.String      `tfsdk:"issuer" json:"issuer,computed"`
	Signature   types.String      `tfsdk:"signature" json:"signature,computed"`
	Status      types.String      `tfsdk:"status" json:"status,computed"`
	UpdatedAt   timetypes.RFC3339 `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
	UploadedOn  timetypes.RFC3339 `tfsdk:"uploaded_on" json:"uploaded_on,computed" format:"date-time"`
}

func (m CustomOriginTrustStoreModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m CustomOriginTrustStoreModel) MarshalJSONForUpdate(state CustomOriginTrustStoreModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package custom_origin_trust_store

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/acm"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type CustomOriginTrustStoresResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[CustomOriginTrustStoresResultDataSourceModel] `json:"result,computed"`
}

type CustomOriginTrustStoresDataSourceModel struct {
	ZoneID   types.String                                                               `tfsdk:"zone_id" path:"zone_id,required"`
	Limit    types.Int64                                                                `tfsdk:"limit" query:"limit,optional"`
	Offset   types.Int64                                                                `tfsdk:"offset" query:"offset,optional"`
	MaxItems types.Int64                                                                `tfsdk:"max_items"`
	Result   customfield.NestedObjectList[CustomOriginTrustStoresResultDataSourceModel] `tfsdk:"result"`
}

func (m *CustomOriginTrustStoresDataSourceModel) toListParams(_ context.Context) (params acm.CustomTrustStoreListParams, diags diag.Diagnostics) {
	params = acm.CustomTrustStoreListParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	if !m.Limit.IsNull() {
		params.Limit = cloudflare.F(m.Limit.ValueInt64())
	}
	if !m.Offset.IsNull() {
		params.Offset = cloudflare.F(m.Offset.ValueInt64())
	}

	return
}

type CustomOriginTrustStoresResultDataSourceModel struct {
	ID          types.String      `tfsdk:"id" json:"id,computed"`
	Certificate types.String      `tfsdk:"certificate" json:"certificate,computed"`
	ExpiresOn   timetypes.RFC3339 `tfsdk:"expires_on" json:"expires_on,computed" format:"date-time"`
	Issuer      types.String      `tfsdk:"issuer" json:"issuer,computed"`
	Signature   types.String      `tfsdk:"signature" json:"signature,computed"`
	Status      types.String      `tfsdk:"status" json:"status,computed"`
	UpdatedAt   timetypes.RFC3339 `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
	UploadedOn  timetypes.RFC3339 `tfsdk:"uploaded_on" json:"uploaded_on,computed" format:"date-time"`
}

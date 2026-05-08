// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package custom_origin_trust_store

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/acm"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type CustomOriginTrustStoreResultDataSourceEnvelope struct {
	Result CustomOriginTrustStoreDataSourceModel `json:"result,computed"`
}

type CustomOriginTrustStoreDataSourceModel struct {
	ID                       types.String                                    `tfsdk:"id" path:"custom_origin_trust_store_id,computed"`
	CustomOriginTrustStoreID types.String                                    `tfsdk:"custom_origin_trust_store_id" path:"custom_origin_trust_store_id,optional"`
	ZoneID                   types.String                                    `tfsdk:"zone_id" path:"zone_id,required"`
	Certificate              types.String                                    `tfsdk:"certificate" json:"certificate,computed"`
	ExpiresOn                timetypes.RFC3339                               `tfsdk:"expires_on" json:"expires_on,computed" format:"date-time"`
	Issuer                   types.String                                    `tfsdk:"issuer" json:"issuer,computed"`
	Signature                types.String                                    `tfsdk:"signature" json:"signature,computed"`
	Status                   types.String                                    `tfsdk:"status" json:"status,computed"`
	UpdatedAt                timetypes.RFC3339                               `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
	UploadedOn               timetypes.RFC3339                               `tfsdk:"uploaded_on" json:"uploaded_on,computed" format:"date-time"`
	Filter                   *CustomOriginTrustStoreFindOneByDataSourceModel `tfsdk:"filter"`
}

func (m *CustomOriginTrustStoreDataSourceModel) toReadParams(_ context.Context) (params acm.CustomTrustStoreGetParams, diags diag.Diagnostics) {
	params = acm.CustomTrustStoreGetParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	return
}

func (m *CustomOriginTrustStoreDataSourceModel) toListParams(_ context.Context) (params acm.CustomTrustStoreListParams, diags diag.Diagnostics) {
	params = acm.CustomTrustStoreListParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	if !m.Filter.Limit.IsNull() {
		params.Limit = cloudflare.F(m.Filter.Limit.ValueInt64())
	}
	if !m.Filter.Offset.IsNull() {
		params.Offset = cloudflare.F(m.Filter.Offset.ValueInt64())
	}

	return
}

type CustomOriginTrustStoreFindOneByDataSourceModel struct {
	Limit  types.Int64 `tfsdk:"limit" query:"limit,optional"`
	Offset types.Int64 `tfsdk:"offset" query:"offset,optional"`
}

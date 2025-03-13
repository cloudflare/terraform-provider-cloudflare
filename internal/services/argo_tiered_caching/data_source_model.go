// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package argo_tiered_caching

import (
  "context"

  "github.com/cloudflare/cloudflare-go/v4"
  "github.com/cloudflare/cloudflare-go/v4/argo"
  "github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
  "github.com/hashicorp/terraform-plugin-framework/diag"
  "github.com/hashicorp/terraform-plugin-framework/types"
)

type ArgoTieredCachingResultDataSourceEnvelope struct {
Result ArgoTieredCachingDataSourceModel `json:"result,computed"`
}

type ArgoTieredCachingDataSourceModel struct {
ZoneID types.String `tfsdk:"zone_id" path:"zone_id,required"`
Editable types.Bool `tfsdk:"editable" json:"editable,computed"`
ID types.String `tfsdk:"id" json:"id,computed"`
ModifiedOn timetypes.RFC3339 `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
Value types.String `tfsdk:"value" json:"value,computed"`
}

func (m *ArgoTieredCachingDataSourceModel) toReadParams(_ context.Context) (params argo.TieredCachingGetParams, diags diag.Diagnostics) {
  params = argo.TieredCachingGetParams{
    ZoneID: cloudflare.F(m.ZoneID.ValueString()),
  }

  return
}

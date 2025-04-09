// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package argo_tiered_caching

import (
  "github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
  "github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
  "github.com/hashicorp/terraform-plugin-framework/types"
)

type ArgoTieredCachingResultEnvelope struct {
Result ArgoTieredCachingModel `json:"result"`
}

type ArgoTieredCachingModel struct {
ID types.String `tfsdk:"id" json:"-,computed"`
ZoneID types.String `tfsdk:"zone_id" path:"zone_id,required"`
Value types.String `tfsdk:"value" json:"value,required"`
Editable types.Bool `tfsdk:"editable" json:"editable,computed"`
ModifiedOn timetypes.RFC3339 `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
}

func (m ArgoTieredCachingModel) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(m)
}

func (m ArgoTieredCachingModel) MarshalJSONForUpdate(state ArgoTieredCachingModel) (data []byte, err error) {
  return apijson.MarshalForPatch(m, state)
}

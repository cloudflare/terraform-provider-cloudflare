// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package regional_hostname

import (
  "context"

  "github.com/cloudflare/cloudflare-go/v4"
  "github.com/cloudflare/cloudflare-go/v4/addressing"
  "github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
  "github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
  "github.com/hashicorp/terraform-plugin-framework/diag"
  "github.com/hashicorp/terraform-plugin-framework/types"
)

type RegionalHostnamesResultListDataSourceEnvelope struct {
Result customfield.NestedObjectList[RegionalHostnamesResultDataSourceModel] `json:"result,computed"`
}

type RegionalHostnamesDataSourceModel struct {
ZoneID types.String `tfsdk:"zone_id" path:"zone_id,required"`
MaxItems types.Int64 `tfsdk:"max_items"`
Result customfield.NestedObjectList[RegionalHostnamesResultDataSourceModel] `tfsdk:"result"`
}

func (m *RegionalHostnamesDataSourceModel) toListParams(_ context.Context) (params addressing.RegionalHostnameListParams, diags diag.Diagnostics) {
  params = addressing.RegionalHostnameListParams{
    ZoneID: cloudflare.F(m.ZoneID.ValueString()),
  }

  return
}

type RegionalHostnamesResultDataSourceModel struct {
CreatedOn timetypes.RFC3339 `tfsdk:"created_on" json:"created_on,computed" format:"date-time"`
Hostname types.String `tfsdk:"hostname" json:"hostname,computed"`
RegionKey types.String `tfsdk:"region_key" json:"region_key,computed"`
}

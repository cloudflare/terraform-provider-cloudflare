// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_routing_catch_all

import (
  "context"

  "github.com/cloudflare/cloudflare-go/v4"
  "github.com/cloudflare/cloudflare-go/v4/email_routing"
  "github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
  "github.com/hashicorp/terraform-plugin-framework/diag"
  "github.com/hashicorp/terraform-plugin-framework/types"
)

type EmailRoutingCatchAllResultDataSourceEnvelope struct {
Result EmailRoutingCatchAllDataSourceModel `json:"result,computed"`
}

type EmailRoutingCatchAllDataSourceModel struct {
ZoneID types.String `tfsdk:"zone_id" path:"zone_id,required"`
Enabled types.Bool `tfsdk:"enabled" json:"enabled,computed"`
ID types.String `tfsdk:"id" json:"id,computed"`
Name types.String `tfsdk:"name" json:"name,computed"`
Tag types.String `tfsdk:"tag" json:"tag,computed"`
Actions customfield.NestedObjectList[EmailRoutingCatchAllActionsDataSourceModel] `tfsdk:"actions" json:"actions,computed"`
Matchers customfield.NestedObjectList[EmailRoutingCatchAllMatchersDataSourceModel] `tfsdk:"matchers" json:"matchers,computed"`
}

func (m *EmailRoutingCatchAllDataSourceModel) toReadParams(_ context.Context) (params email_routing.RuleCatchAllGetParams, diags diag.Diagnostics) {
  params = email_routing.RuleCatchAllGetParams{
    ZoneID: cloudflare.F(m.ZoneID.ValueString()),
  }

  return
}

type EmailRoutingCatchAllActionsDataSourceModel struct {
Type types.String `tfsdk:"type" json:"type,computed"`
Value customfield.List[types.String] `tfsdk:"value" json:"value,computed"`
}

type EmailRoutingCatchAllMatchersDataSourceModel struct {
Type types.String `tfsdk:"type" json:"type,computed"`
}

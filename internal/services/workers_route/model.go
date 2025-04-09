// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_route

import (
  "github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
  "github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
  "github.com/hashicorp/terraform-plugin-framework/types"
)

type WorkersRouteResultEnvelope struct {
Result WorkersRouteModel `json:"result"`
}

type WorkersRouteModel struct {
ZoneID types.String `tfsdk:"zone_id" path:"zone_id,required"`
RouteID types.String `tfsdk:"route_id" path:"route_id,optional"`
Pattern types.String `tfsdk:"pattern" json:"pattern,required"`
Script types.String `tfsdk:"script" json:"script,optional"`
ID types.String `tfsdk:"id" json:"id,computed"`
Success types.Bool `tfsdk:"success" json:"success,computed"`
Errors customfield.NestedObjectList[WorkersRouteErrorsModel] `tfsdk:"errors" json:"errors,computed"`
Messages customfield.NestedObjectList[WorkersRouteMessagesModel] `tfsdk:"messages" json:"messages,computed"`
}

func (m WorkersRouteModel) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(m)
}

func (m WorkersRouteModel) MarshalJSONForUpdate(state WorkersRouteModel) (data []byte, err error) {
  return apijson.MarshalForUpdate(m, state)
}

type WorkersRouteErrorsModel struct {
Code types.Int64 `tfsdk:"code" json:"code,computed"`
Message types.String `tfsdk:"message" json:"message,computed"`
DocumentationURL types.String `tfsdk:"documentation_url" json:"documentation_url,computed"`
Source customfield.NestedObject[WorkersRouteErrorsSourceModel] `tfsdk:"source" json:"source,computed"`
}

type WorkersRouteErrorsSourceModel struct {
Pointer types.String `tfsdk:"pointer" json:"pointer,computed"`
}

type WorkersRouteMessagesModel struct {
Code types.Int64 `tfsdk:"code" json:"code,computed"`
Message types.String `tfsdk:"message" json:"message,computed"`
DocumentationURL types.String `tfsdk:"documentation_url" json:"documentation_url,computed"`
Source customfield.NestedObject[WorkersRouteMessagesSourceModel] `tfsdk:"source" json:"source,computed"`
}

type WorkersRouteMessagesSourceModel struct {
Pointer types.String `tfsdk:"pointer" json:"pointer,computed"`
}

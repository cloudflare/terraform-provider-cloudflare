// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_custom_domain

import (
  "github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
  "github.com/hashicorp/terraform-plugin-framework/types"
)

type WorkersCustomDomainResultEnvelope struct {
Result WorkersCustomDomainModel `json:"result"`
}

type WorkersCustomDomainModel struct {
ID types.String `tfsdk:"id" json:"id,computed"`
AccountID types.String `tfsdk:"account_id" path:"account_id,required"`
Environment types.String `tfsdk:"environment" json:"environment,required"`
Hostname types.String `tfsdk:"hostname" json:"hostname,required"`
Service types.String `tfsdk:"service" json:"service,required"`
ZoneID types.String `tfsdk:"zone_id" json:"zone_id,required"`
ZoneName types.String `tfsdk:"zone_name" json:"zone_name,computed"`
}

func (m WorkersCustomDomainModel) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(m)
}

func (m WorkersCustomDomainModel) MarshalJSONForUpdate(state WorkersCustomDomainModel) (data []byte, err error) {
  return apijson.MarshalForUpdate(m, state)
}

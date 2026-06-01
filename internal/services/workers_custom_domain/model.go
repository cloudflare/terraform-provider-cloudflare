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
	ID          types.String `tfsdk:"id" json:"id,computed"`
	AccountID   types.String `tfsdk:"account_id" path:"account_id,required"`
	Hostname    types.String `tfsdk:"hostname" json:"hostname,required"`
	Service     types.String `tfsdk:"service" json:"service,required"`
	Environment types.String `tfsdk:"environment" json:"environment,computed_optional"`
	ZoneID      types.String `tfsdk:"zone_id" json:"zone_id,computed_optional"`
	ZoneName    types.String `tfsdk:"zone_name" json:"zone_name,computed_optional"`
	CERTID      types.String `tfsdk:"cert_id" json:"cert_id,computed"`
}

func (m WorkersCustomDomainModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m WorkersCustomDomainModel) MarshalJSONForUpdate(state WorkersCustomDomainModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

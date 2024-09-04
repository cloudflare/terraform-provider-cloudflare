// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_custom_domain

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type WorkersCustomDomainResultEnvelope struct {
	Result WorkersCustomDomainModel `json:"result"`
}

type WorkersCustomDomainModel struct {
	ID          types.String `tfsdk:"id" json:"-,computed"`
	Hostname    types.String `tfsdk:"hostname" json:"hostname,computed_optional"`
	AccountID   types.String `tfsdk:"account_id" path:"account_id,required"`
	Environment types.String `tfsdk:"environment" json:"environment,computed_optional"`
	Service     types.String `tfsdk:"service" json:"service,computed_optional"`
	ZoneID      types.String `tfsdk:"zone_id" json:"zone_id,computed_optional"`
	ZoneName    types.String `tfsdk:"zone_name" json:"zone_name,computed"`
}

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package worker_domain

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type WorkerDomainResultEnvelope struct {
	Result WorkerDomainModel `json:"result,computed"`
}

type WorkerDomainModel struct {
	ID          types.String `tfsdk:"id" json:"-,computed"`
	Hostname    types.String `tfsdk:"hostname" json:"hostname"`
	AccountID   types.String `tfsdk:"account_id" path:"account_id"`
	Environment types.String `tfsdk:"environment" json:"environment"`
	Service     types.String `tfsdk:"service" json:"service"`
	ZoneID      types.String `tfsdk:"zone_id" json:"zone_id"`
	ZoneName    types.String `tfsdk:"zone_name" json:"zone_name,computed"`
}

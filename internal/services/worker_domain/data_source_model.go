// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package worker_domain

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type WorkerDomainResultDataSourceEnvelope struct {
	Result WorkerDomainDataSourceModel `json:"result,computed"`
}

type WorkerDomainResultListDataSourceEnvelope struct {
	Result *[]*WorkerDomainDataSourceModel `json:"result,computed"`
}

type WorkerDomainDataSourceModel struct {
	AccountID   types.String                          `tfsdk:"account_id" path:"account_id"`
	DomainID    types.String                          `tfsdk:"domain_id" path:"domain_id"`
	ID          types.String                          `tfsdk:"id" json:"id"`
	Environment types.String                          `tfsdk:"environment" json:"environment"`
	Hostname    types.String                          `tfsdk:"hostname" json:"hostname"`
	Service     types.String                          `tfsdk:"service" json:"service"`
	ZoneID      types.String                          `tfsdk:"zone_id" json:"zone_id"`
	ZoneName    types.String                          `tfsdk:"zone_name" json:"zone_name"`
	Filter      *WorkerDomainFindOneByDataSourceModel `tfsdk:"filter"`
}

type WorkerDomainFindOneByDataSourceModel struct {
	AccountID   types.String `tfsdk:"account_id" path:"account_id"`
	Environment types.String `tfsdk:"environment" query:"environment"`
	Hostname    types.String `tfsdk:"hostname" query:"hostname"`
	Service     types.String `tfsdk:"service" query:"service"`
	ZoneID      types.String `tfsdk:"zone_id" query:"zone_id"`
	ZoneName    types.String `tfsdk:"zone_name" query:"zone_name"`
}

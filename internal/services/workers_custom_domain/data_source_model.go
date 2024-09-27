// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_custom_domain

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/workers"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type WorkersCustomDomainResultDataSourceEnvelope struct {
	Result WorkersCustomDomainDataSourceModel `json:"result,computed"`
}

type WorkersCustomDomainResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[WorkersCustomDomainDataSourceModel] `json:"result,computed"`
}

type WorkersCustomDomainDataSourceModel struct {
	AccountID   types.String                                 `tfsdk:"account_id" path:"account_id,optional"`
	DomainID    types.String                                 `tfsdk:"domain_id" path:"domain_id,optional"`
	Environment types.String                                 `tfsdk:"environment" json:"environment,computed"`
	Hostname    types.String                                 `tfsdk:"hostname" json:"hostname,computed"`
	ID          types.String                                 `tfsdk:"id" json:"id,computed"`
	Service     types.String                                 `tfsdk:"service" json:"service,computed"`
	ZoneID      types.String                                 `tfsdk:"zone_id" json:"zone_id,computed"`
	ZoneName    types.String                                 `tfsdk:"zone_name" json:"zone_name,computed"`
	Filter      *WorkersCustomDomainFindOneByDataSourceModel `tfsdk:"filter"`
}

func (m *WorkersCustomDomainDataSourceModel) toReadParams(_ context.Context) (params workers.DomainGetParams, diags diag.Diagnostics) {
	params = workers.DomainGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

func (m *WorkersCustomDomainDataSourceModel) toListParams(_ context.Context) (params workers.DomainListParams, diags diag.Diagnostics) {
	params = workers.DomainListParams{
		AccountID: cloudflare.F(m.Filter.AccountID.ValueString()),
	}

	if !m.Filter.Environment.IsNull() {
		params.Environment = cloudflare.F(m.Filter.Environment.ValueString())
	}
	if !m.Filter.Hostname.IsNull() {
		params.Hostname = cloudflare.F(m.Filter.Hostname.ValueString())
	}
	if !m.Filter.Service.IsNull() {
		params.Service = cloudflare.F(m.Filter.Service.ValueString())
	}
	if !m.Filter.ZoneID.IsNull() {
		params.ZoneID = cloudflare.F(m.Filter.ZoneID.ValueString())
	}
	if !m.Filter.ZoneName.IsNull() {
		params.ZoneName = cloudflare.F(m.Filter.ZoneName.ValueString())
	}

	return
}

type WorkersCustomDomainFindOneByDataSourceModel struct {
	AccountID   types.String `tfsdk:"account_id" path:"account_id,required"`
	Environment types.String `tfsdk:"environment" query:"environment,optional"`
	Hostname    types.String `tfsdk:"hostname" query:"hostname,optional"`
	Service     types.String `tfsdk:"service" query:"service,optional"`
	ZoneID      types.String `tfsdk:"zone_id" query:"zone_id,optional"`
	ZoneName    types.String `tfsdk:"zone_name" query:"zone_name,optional"`
}

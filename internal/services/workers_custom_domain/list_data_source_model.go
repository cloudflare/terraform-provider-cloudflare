// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_custom_domain

import (
	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/workers"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type WorkersCustomDomainsResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[WorkersCustomDomainsResultDataSourceModel] `json:"result,computed"`
}

type WorkersCustomDomainsDataSourceModel struct {
	AccountID   types.String                                                            `tfsdk:"account_id" path:"account_id"`
	Environment types.String                                                            `tfsdk:"environment" query:"environment"`
	Hostname    types.String                                                            `tfsdk:"hostname" query:"hostname"`
	Service     types.String                                                            `tfsdk:"service" query:"service"`
	ZoneID      types.String                                                            `tfsdk:"zone_id" query:"zone_id"`
	ZoneName    types.String                                                            `tfsdk:"zone_name" query:"zone_name"`
	MaxItems    types.Int64                                                             `tfsdk:"max_items"`
	Result      customfield.NestedObjectList[WorkersCustomDomainsResultDataSourceModel] `tfsdk:"result"`
}

func (m *WorkersCustomDomainsDataSourceModel) toListParams() (params workers.DomainListParams, diags diag.Diagnostics) {
	params = workers.DomainListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	if !m.Environment.IsNull() {
		params.Environment = cloudflare.F(m.Environment.ValueString())
	}
	if !m.Hostname.IsNull() {
		params.Hostname = cloudflare.F(m.Hostname.ValueString())
	}
	if !m.Service.IsNull() {
		params.Service = cloudflare.F(m.Service.ValueString())
	}
	if !m.ZoneID.IsNull() {
		params.ZoneID = cloudflare.F(m.ZoneID.ValueString())
	}
	if !m.ZoneName.IsNull() {
		params.ZoneName = cloudflare.F(m.ZoneName.ValueString())
	}

	return
}

type WorkersCustomDomainsResultDataSourceModel struct {
	ID          types.String `tfsdk:"id" json:"id,computed_optional"`
	Environment types.String `tfsdk:"environment" json:"environment,computed_optional"`
	Hostname    types.String `tfsdk:"hostname" json:"hostname,computed_optional"`
	Service     types.String `tfsdk:"service" json:"service,computed_optional"`
	ZoneID      types.String `tfsdk:"zone_id" json:"zone_id,computed_optional"`
	ZoneName    types.String `tfsdk:"zone_name" json:"zone_name,computed_optional"`
}

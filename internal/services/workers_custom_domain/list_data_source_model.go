// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_custom_domain

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type WorkersCustomDomainsResultListDataSourceEnvelope struct {
	Result *[]*WorkersCustomDomainsResultDataSourceModel `json:"result,computed"`
}

type WorkersCustomDomainsDataSourceModel struct {
	AccountID   types.String                                  `tfsdk:"account_id" path:"account_id"`
	Environment types.String                                  `tfsdk:"environment" query:"environment"`
	Hostname    types.String                                  `tfsdk:"hostname" query:"hostname"`
	Service     types.String                                  `tfsdk:"service" query:"service"`
	ZoneID      types.String                                  `tfsdk:"zone_id" query:"zone_id"`
	ZoneName    types.String                                  `tfsdk:"zone_name" query:"zone_name"`
	MaxItems    types.Int64                                   `tfsdk:"max_items"`
	Result      *[]*WorkersCustomDomainsResultDataSourceModel `tfsdk:"result"`
}

type WorkersCustomDomainsResultDataSourceModel struct {
	ID          types.String `tfsdk:"id" json:"id"`
	Environment types.String `tfsdk:"environment" json:"environment"`
	Hostname    types.String `tfsdk:"hostname" json:"hostname"`
	Service     types.String `tfsdk:"service" json:"service"`
	ZoneID      types.String `tfsdk:"zone_id" json:"zone_id"`
	ZoneName    types.String `tfsdk:"zone_name" json:"zone_name"`
}

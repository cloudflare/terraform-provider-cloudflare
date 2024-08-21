// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package hyperdrive_config

import (
	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/hyperdrive"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type HyperdriveConfigsResultListDataSourceEnvelope struct {
	Result *[]*HyperdriveConfigsResultDataSourceModel `json:"result,computed"`
}

type HyperdriveConfigsDataSourceModel struct {
	AccountID types.String                               `tfsdk:"account_id" path:"account_id"`
	MaxItems  types.Int64                                `tfsdk:"max_items"`
	Result    *[]*HyperdriveConfigsResultDataSourceModel `tfsdk:"result"`
}

func (m *HyperdriveConfigsDataSourceModel) toListParams() (params hyperdrive.ConfigListParams, diags diag.Diagnostics) {
	params = hyperdrive.ConfigListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

type HyperdriveConfigsResultDataSourceModel struct {
	Caching *HyperdriveConfigsCachingDataSourceModel `tfsdk:"caching" json:"caching"`
	Name    types.String                             `tfsdk:"name" json:"name"`
	Origin  *HyperdriveConfigsOriginDataSourceModel  `tfsdk:"origin" json:"origin"`
}

type HyperdriveConfigsCachingDataSourceModel struct {
	Disabled             types.Bool  `tfsdk:"disabled" json:"disabled"`
	MaxAge               types.Int64 `tfsdk:"max_age" json:"max_age"`
	StaleWhileRevalidate types.Int64 `tfsdk:"stale_while_revalidate" json:"stale_while_revalidate"`
}

type HyperdriveConfigsOriginDataSourceModel struct {
	Database       types.String `tfsdk:"database" json:"database,computed"`
	Host           types.String `tfsdk:"host" json:"host,computed"`
	Scheme         types.String `tfsdk:"scheme" json:"scheme,computed"`
	User           types.String `tfsdk:"user" json:"user,computed"`
	AccessClientID types.String `tfsdk:"access_client_id" json:"access_client_id"`
	Port           types.Int64  `tfsdk:"port" json:"port"`
}

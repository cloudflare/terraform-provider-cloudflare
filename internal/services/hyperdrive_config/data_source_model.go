// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package hyperdrive_config

import (
	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/hyperdrive"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type HyperdriveConfigResultDataSourceEnvelope struct {
	Result HyperdriveConfigDataSourceModel `json:"result,computed"`
}

type HyperdriveConfigResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[HyperdriveConfigDataSourceModel] `json:"result,computed"`
}

type HyperdriveConfigDataSourceModel struct {
	AccountID    types.String                              `tfsdk:"account_id" path:"account_id"`
	HyperdriveID types.String                              `tfsdk:"hyperdrive_id" path:"hyperdrive_id"`
	Name         types.String                              `tfsdk:"name" json:"name,computed_optional"`
	Caching      *HyperdriveConfigCachingDataSourceModel   `tfsdk:"caching" json:"caching,computed_optional"`
	Origin       *HyperdriveConfigOriginDataSourceModel    `tfsdk:"origin" json:"origin,computed_optional"`
	Filter       *HyperdriveConfigFindOneByDataSourceModel `tfsdk:"filter"`
}

func (m *HyperdriveConfigDataSourceModel) toReadParams() (params hyperdrive.ConfigGetParams, diags diag.Diagnostics) {
	params = hyperdrive.ConfigGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

func (m *HyperdriveConfigDataSourceModel) toListParams() (params hyperdrive.ConfigListParams, diags diag.Diagnostics) {
	params = hyperdrive.ConfigListParams{
		AccountID: cloudflare.F(m.Filter.AccountID.ValueString()),
	}

	return
}

type HyperdriveConfigCachingDataSourceModel struct {
	Disabled             types.Bool  `tfsdk:"disabled" json:"disabled,computed_optional"`
	MaxAge               types.Int64 `tfsdk:"max_age" json:"max_age,computed_optional"`
	StaleWhileRevalidate types.Int64 `tfsdk:"stale_while_revalidate" json:"stale_while_revalidate,computed_optional"`
}

type HyperdriveConfigOriginDataSourceModel struct {
	Database       types.String `tfsdk:"database" json:"database,computed"`
	Host           types.String `tfsdk:"host" json:"host,computed"`
	Scheme         types.String `tfsdk:"scheme" json:"scheme,computed"`
	User           types.String `tfsdk:"user" json:"user,computed"`
	AccessClientID types.String `tfsdk:"access_client_id" json:"access_client_id,computed_optional"`
	Port           types.Int64  `tfsdk:"port" json:"port,computed_optional"`
}

type HyperdriveConfigFindOneByDataSourceModel struct {
	AccountID types.String `tfsdk:"account_id" path:"account_id"`
}

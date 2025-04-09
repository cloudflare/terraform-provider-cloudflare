// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package hyperdrive_config

import (
  "context"

  "github.com/cloudflare/cloudflare-go/v4"
  "github.com/cloudflare/cloudflare-go/v4/hyperdrive"
  "github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
  "github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
  "github.com/hashicorp/terraform-plugin-framework/diag"
  "github.com/hashicorp/terraform-plugin-framework/types"
)

type HyperdriveConfigResultDataSourceEnvelope struct {
Result HyperdriveConfigDataSourceModel `json:"result,computed"`
}

type HyperdriveConfigDataSourceModel struct {
ID types.String `tfsdk:"id" path:"hyperdrive_id,computed"`
HyperdriveID types.String `tfsdk:"hyperdrive_id" path:"hyperdrive_id,optional"`
AccountID types.String `tfsdk:"account_id" path:"account_id,required"`
CreatedOn timetypes.RFC3339 `tfsdk:"created_on" json:"created_on,computed" format:"date-time"`
ModifiedOn timetypes.RFC3339 `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
Name types.String `tfsdk:"name" json:"name,computed"`
Caching customfield.NestedObject[HyperdriveConfigCachingDataSourceModel] `tfsdk:"caching" json:"caching,computed"`
Origin customfield.NestedObject[HyperdriveConfigOriginDataSourceModel] `tfsdk:"origin" json:"origin,computed"`
}

func (m *HyperdriveConfigDataSourceModel) toReadParams(_ context.Context) (params hyperdrive.ConfigGetParams, diags diag.Diagnostics) {
  params = hyperdrive.ConfigGetParams{
    AccountID: cloudflare.F(m.AccountID.ValueString()),
  }

  return
}

type HyperdriveConfigCachingDataSourceModel struct {
Disabled types.Bool `tfsdk:"disabled" json:"disabled,computed"`
MaxAge types.Int64 `tfsdk:"max_age" json:"max_age,computed"`
StaleWhileRevalidate types.Int64 `tfsdk:"stale_while_revalidate" json:"stale_while_revalidate,computed"`
}

type HyperdriveConfigOriginDataSourceModel struct {
Database types.String `tfsdk:"database" json:"database,computed"`
Host types.String `tfsdk:"host" json:"host,computed"`
Password types.String `tfsdk:"password" json:"password,computed"`
Port types.Int64 `tfsdk:"port" json:"port,computed"`
Scheme types.String `tfsdk:"scheme" json:"scheme,computed"`
User types.String `tfsdk:"user" json:"user,computed"`
AccessClientID types.String `tfsdk:"access_client_id" json:"access_client_id,computed"`
AccessClientSecret types.String `tfsdk:"access_client_secret" json:"access_client_secret,computed"`
}

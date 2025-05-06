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

type HyperdriveConfigsResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[HyperdriveConfigsResultDataSourceModel] `json:"result,computed"`
}

type HyperdriveConfigsDataSourceModel struct {
	AccountID types.String                                                         `tfsdk:"account_id" path:"account_id,required"`
	MaxItems  types.Int64                                                          `tfsdk:"max_items"`
	Result    customfield.NestedObjectList[HyperdriveConfigsResultDataSourceModel] `tfsdk:"result"`
}

func (m *HyperdriveConfigsDataSourceModel) toListParams(_ context.Context) (params hyperdrive.ConfigListParams, diags diag.Diagnostics) {
	params = hyperdrive.ConfigListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

type HyperdriveConfigsResultDataSourceModel struct {
	ID         types.String                                                      `tfsdk:"id" json:"id,computed"`
	Name       types.String                                                      `tfsdk:"name" json:"name,computed"`
	Origin     customfield.NestedObject[HyperdriveConfigsOriginDataSourceModel]  `tfsdk:"origin" json:"origin,computed"`
	Caching    customfield.NestedObject[HyperdriveConfigsCachingDataSourceModel] `tfsdk:"caching" json:"caching,computed"`
	CreatedOn  timetypes.RFC3339                                                 `tfsdk:"created_on" json:"created_on,computed" format:"date-time"`
	ModifiedOn timetypes.RFC3339                                                 `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
	MTLS       customfield.NestedObject[HyperdriveConfigsMTLSDataSourceModel]    `tfsdk:"mtls" json:"mtls,computed"`
}

type HyperdriveConfigsOriginDataSourceModel struct {
	Database           types.String `tfsdk:"database" json:"database,computed"`
	Host               types.String `tfsdk:"host" json:"host,computed"`
	Password           types.String `tfsdk:"password" json:"password,computed"`
	Port               types.Int64  `tfsdk:"port" json:"port,computed"`
	Scheme             types.String `tfsdk:"scheme" json:"scheme,computed"`
	User               types.String `tfsdk:"user" json:"user,computed"`
	AccessClientID     types.String `tfsdk:"access_client_id" json:"access_client_id,computed"`
	AccessClientSecret types.String `tfsdk:"access_client_secret" json:"access_client_secret,computed"`
}

type HyperdriveConfigsCachingDataSourceModel struct {
	Disabled             types.Bool  `tfsdk:"disabled" json:"disabled,computed"`
	MaxAge               types.Int64 `tfsdk:"max_age" json:"max_age,computed"`
	StaleWhileRevalidate types.Int64 `tfsdk:"stale_while_revalidate" json:"stale_while_revalidate,computed"`
}

type HyperdriveConfigsMTLSDataSourceModel struct {
	CACertificateID   types.String `tfsdk:"ca_certificate_id" json:"ca_certificate_id,computed"`
	MTLSCertificateID types.String `tfsdk:"mtls_certificate_id" json:"mtls_certificate_id,computed"`
	Sslmode           types.String `tfsdk:"sslmode" json:"sslmode,computed"`
}

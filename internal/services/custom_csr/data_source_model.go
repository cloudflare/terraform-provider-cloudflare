// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package custom_csr

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/custom_csrs"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type CustomCsrResultDataSourceEnvelope struct {
	Result CustomCsrDataSourceModel `json:"result,computed"`
}

type CustomCsrDataSourceModel struct {
	ID                 types.String                   `tfsdk:"id" path:"custom_csr_id,computed"`
	CustomCsrID        types.String                   `tfsdk:"custom_csr_id" path:"custom_csr_id,required"`
	AccountID          types.String                   `tfsdk:"account_id" path:"account_id,optional"`
	ZoneID             types.String                   `tfsdk:"zone_id" path:"zone_id,optional"`
	AccountTag         types.String                   `tfsdk:"account_tag" json:"account_tag,computed"`
	CommonName         types.String                   `tfsdk:"common_name" json:"common_name,computed"`
	Country            types.String                   `tfsdk:"country" json:"country,computed"`
	CreatedAt          timetypes.RFC3339              `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Csr                types.String                   `tfsdk:"csr" json:"csr,computed"`
	Description        types.String                   `tfsdk:"description" json:"description,computed"`
	KeyType            types.String                   `tfsdk:"key_type" json:"key_type,computed"`
	Locality           types.String                   `tfsdk:"locality" json:"locality,computed"`
	Name               types.String                   `tfsdk:"name" json:"name,computed"`
	Organization       types.String                   `tfsdk:"organization" json:"organization,computed"`
	OrganizationalUnit types.String                   `tfsdk:"organizational_unit" json:"organizational_unit,computed"`
	State              types.String                   `tfsdk:"state" json:"state,computed"`
	Sans               customfield.List[types.String] `tfsdk:"sans" json:"sans,computed"`
}

func (m *CustomCsrDataSourceModel) toReadParams(_ context.Context) (params custom_csrs.CustomCsrGetParams, diags diag.Diagnostics) {
	params = custom_csrs.CustomCsrGetParams{}

	if !m.AccountID.IsNull() {
		params.AccountID = cloudflare.F(m.AccountID.ValueString())
	} else {
		params.ZoneID = cloudflare.F(m.ZoneID.ValueString())
	}

	return
}

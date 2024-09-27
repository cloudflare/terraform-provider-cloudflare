// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package registrar_domain

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v3"
	"github.com/cloudflare/cloudflare-go/v3/registrar"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type RegistrarDomainsResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[RegistrarDomainsResultDataSourceModel] `json:"result,computed"`
}

type RegistrarDomainsDataSourceModel struct {
	AccountID types.String                                                        `tfsdk:"account_id" path:"account_id,required"`
	MaxItems  types.Int64                                                         `tfsdk:"max_items"`
	Result    customfield.NestedObjectList[RegistrarDomainsResultDataSourceModel] `tfsdk:"result"`
}

func (m *RegistrarDomainsDataSourceModel) toListParams(_ context.Context) (params registrar.DomainListParams, diags diag.Diagnostics) {
	params = registrar.DomainListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

type RegistrarDomainsResultDataSourceModel struct {
	Errors     customfield.NestedObjectList[RegistrarDomainsErrorsDataSourceModel]   `tfsdk:"errors" json:"errors,computed"`
	Messages   customfield.NestedObjectList[RegistrarDomainsMessagesDataSourceModel] `tfsdk:"messages" json:"messages,computed"`
	Result     jsontypes.Normalized                                                  `tfsdk:"result" json:"result,computed"`
	Success    types.Bool                                                            `tfsdk:"success" json:"success,computed"`
	ResultInfo customfield.NestedObject[RegistrarDomainsResultInfoDataSourceModel]   `tfsdk:"result_info" json:"result_info,computed"`
}

type RegistrarDomainsErrorsDataSourceModel struct {
	Code    types.Int64  `tfsdk:"code" json:"code,computed"`
	Message types.String `tfsdk:"message" json:"message,computed"`
}

type RegistrarDomainsMessagesDataSourceModel struct {
	Code    types.Int64  `tfsdk:"code" json:"code,computed"`
	Message types.String `tfsdk:"message" json:"message,computed"`
}

type RegistrarDomainsResultInfoDataSourceModel struct {
	Count      types.Float64 `tfsdk:"count" json:"count,computed"`
	Page       types.Float64 `tfsdk:"page" json:"page,computed"`
	PerPage    types.Float64 `tfsdk:"per_page" json:"per_page,computed"`
	TotalCount types.Float64 `tfsdk:"total_count" json:"total_count,computed"`
}

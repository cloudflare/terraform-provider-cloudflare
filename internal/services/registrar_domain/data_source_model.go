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

type RegistrarDomainResultDataSourceEnvelope struct {
	Result RegistrarDomainDataSourceModel `json:"result,computed"`
}

type RegistrarDomainResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[RegistrarDomainDataSourceModel] `json:"result,computed"`
}

type RegistrarDomainDataSourceModel struct {
	AccountID  types.String                               `tfsdk:"account_id" path:"account_id,optional"`
	DomainName types.String                               `tfsdk:"domain_name" path:"domain_name,optional"`
	Success    types.Bool                                 `tfsdk:"success" json:"success,optional"`
	Errors     *[]*RegistrarDomainErrorsDataSourceModel   `tfsdk:"errors" json:"errors,optional"`
	Messages   *[]*RegistrarDomainMessagesDataSourceModel `tfsdk:"messages" json:"messages,optional"`
	ResultInfo *RegistrarDomainResultInfoDataSourceModel  `tfsdk:"result_info" json:"result_info,optional"`
	Result     jsontypes.Normalized                       `tfsdk:"result" json:"result,optional"`
	Filter     *RegistrarDomainFindOneByDataSourceModel   `tfsdk:"filter"`
}

func (m *RegistrarDomainDataSourceModel) toReadParams(_ context.Context) (params registrar.DomainGetParams, diags diag.Diagnostics) {
	params = registrar.DomainGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

func (m *RegistrarDomainDataSourceModel) toListParams(_ context.Context) (params registrar.DomainListParams, diags diag.Diagnostics) {
	params = registrar.DomainListParams{
		AccountID: cloudflare.F(m.Filter.AccountID.ValueString()),
	}

	return
}

type RegistrarDomainErrorsDataSourceModel struct {
	Code    types.Int64  `tfsdk:"code" json:"code,computed"`
	Message types.String `tfsdk:"message" json:"message,computed"`
}

type RegistrarDomainMessagesDataSourceModel struct {
	Code    types.Int64  `tfsdk:"code" json:"code,computed"`
	Message types.String `tfsdk:"message" json:"message,computed"`
}

type RegistrarDomainResultInfoDataSourceModel struct {
	Count      types.Float64 `tfsdk:"count" json:"count,computed"`
	Page       types.Float64 `tfsdk:"page" json:"page,computed"`
	PerPage    types.Float64 `tfsdk:"per_page" json:"per_page,computed"`
	TotalCount types.Float64 `tfsdk:"total_count" json:"total_count,computed"`
}

type RegistrarDomainFindOneByDataSourceModel struct {
	AccountID types.String `tfsdk:"account_id" path:"account_id,required"`
}

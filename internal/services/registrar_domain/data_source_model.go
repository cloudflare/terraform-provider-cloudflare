// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package registrar_domain

import (
  "context"

  "github.com/cloudflare/cloudflare-go/v4"
  "github.com/cloudflare/cloudflare-go/v4/registrar"
  "github.com/hashicorp/terraform-plugin-framework/diag"
  "github.com/hashicorp/terraform-plugin-framework/types"
)

type RegistrarDomainResultDataSourceEnvelope struct {
Result RegistrarDomainDataSourceModel `json:"result,computed"`
}

type RegistrarDomainDataSourceModel struct {
AccountID types.String `tfsdk:"account_id" path:"account_id,required"`
DomainName types.String `tfsdk:"domain_name" path:"domain_name,required"`
}

func (m *RegistrarDomainDataSourceModel) toReadParams(_ context.Context) (params registrar.DomainGetParams, diags diag.Diagnostics) {
  params = registrar.DomainGetParams{
    AccountID: cloudflare.F(m.AccountID.ValueString()),
  }

  return
}

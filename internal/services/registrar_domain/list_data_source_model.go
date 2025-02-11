// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package registrar_domain

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/registrar"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
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
}

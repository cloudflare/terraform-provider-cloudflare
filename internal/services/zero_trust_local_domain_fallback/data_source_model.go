// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_local_domain_fallback

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustLocalDomainFallbackResultDataSourceEnvelope struct {
	Result ZeroTrustLocalDomainFallbackDataSourceModel `json:"result,computed"`
}

type ZeroTrustLocalDomainFallbackResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[ZeroTrustLocalDomainFallbackDataSourceModel] `json:"result,computed"`
}

type ZeroTrustLocalDomainFallbackDataSourceModel struct {
	AccountID   types.String                                          `tfsdk:"account_id" path:"account_id,optional"`
	PolicyID    types.String                                          `tfsdk:"policy_id" path:"policy_id,optional"`
	Description types.String                                          `tfsdk:"description" json:"description,optional"`
	Suffix      types.String                                          `tfsdk:"suffix" json:"suffix,optional"`
	DNSServer   *[]types.String                                       `tfsdk:"dns_server" json:"dns_server,optional"`
	Filter      *ZeroTrustLocalDomainFallbackFindOneByDataSourceModel `tfsdk:"filter"`
}

func (m *ZeroTrustLocalDomainFallbackDataSourceModel) toReadParams(_ context.Context) (params zero_trust.DevicePolicyFallbackDomainGetParams, diags diag.Diagnostics) {
	params = zero_trust.DevicePolicyFallbackDomainGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

func (m *ZeroTrustLocalDomainFallbackDataSourceModel) toListParams(_ context.Context) (params zero_trust.DevicePolicyFallbackDomainListParams, diags diag.Diagnostics) {
	params = zero_trust.DevicePolicyFallbackDomainListParams{
		AccountID: cloudflare.F(m.Filter.AccountID.ValueString()),
	}

	return
}

type ZeroTrustLocalDomainFallbackFindOneByDataSourceModel struct {
	AccountID types.String `tfsdk:"account_id" path:"account_id,required"`
}

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package secondary_dns_acl

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v3"
	"github.com/cloudflare/cloudflare-go/v3/secondary_dns"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type SecondaryDNSACLResultDataSourceEnvelope struct {
	Result SecondaryDNSACLDataSourceModel `json:"result,computed"`
}

type SecondaryDNSACLResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[SecondaryDNSACLDataSourceModel] `json:"result,computed"`
}

type SecondaryDNSACLDataSourceModel struct {
	AccountID types.String                             `tfsdk:"account_id" path:"account_id,optional"`
	ACLID     types.String                             `tfsdk:"acl_id" path:"acl_id,optional"`
	ID        types.String                             `tfsdk:"id" json:"id,computed"`
	IPRange   types.String                             `tfsdk:"ip_range" json:"ip_range,computed"`
	Name      types.String                             `tfsdk:"name" json:"name,computed"`
	Filter    *SecondaryDNSACLFindOneByDataSourceModel `tfsdk:"filter"`
}

func (m *SecondaryDNSACLDataSourceModel) toReadParams(_ context.Context) (params secondary_dns.ACLGetParams, diags diag.Diagnostics) {
	params = secondary_dns.ACLGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

func (m *SecondaryDNSACLDataSourceModel) toListParams(_ context.Context) (params secondary_dns.ACLListParams, diags diag.Diagnostics) {
	params = secondary_dns.ACLListParams{
		AccountID: cloudflare.F(m.Filter.AccountID.ValueString()),
	}

	return
}

type SecondaryDNSACLFindOneByDataSourceModel struct {
	AccountID types.String `tfsdk:"account_id" path:"account_id,required"`
}

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_organization_trust_settings

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustOrganizationTrustSettingsResultDataSourceEnvelope struct {
	Result ZeroTrustOrganizationTrustSettingsDataSourceModel `json:"result,computed"`
}

type ZeroTrustOrganizationTrustSettingsDataSourceModel struct {
	ID                           types.String                   `tfsdk:"id" path:"account_id,computed"`
	AccountID                    types.String                   `tfsdk:"account_id" path:"account_id,required"`
	GrantTrustToAllAccountsInOrg types.Bool                     `tfsdk:"grant_trust_to_all_accounts_in_org" json:"grant_trust_to_all_accounts_in_org,computed"`
	TrustingAccounts             customfield.List[types.String] `tfsdk:"trusting_accounts" json:"trusting_accounts,computed"`
}

func (m *ZeroTrustOrganizationTrustSettingsDataSourceModel) toReadParams(_ context.Context) (params zero_trust.OrganizationTrustSettingGetParams, diags diag.Diagnostics) {
	params = zero_trust.OrganizationTrustSettingGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

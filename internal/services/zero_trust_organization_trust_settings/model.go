// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_organization_trust_settings

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustOrganizationTrustSettingsResultEnvelope struct {
	Result ZeroTrustOrganizationTrustSettingsModel `json:"result"`
}

type ZeroTrustOrganizationTrustSettingsModel struct {
	ID                           types.String    `tfsdk:"id" json:"-,computed"`
	AccountID                    types.String    `tfsdk:"account_id" path:"account_id,required"`
	GrantTrustToAllAccountsInOrg types.Bool      `tfsdk:"grant_trust_to_all_accounts_in_org" json:"grant_trust_to_all_accounts_in_org,required"`
	TrustingAccounts             *[]types.String `tfsdk:"trusting_accounts" json:"trusting_accounts,optional"`
}

func (m ZeroTrustOrganizationTrustSettingsModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m ZeroTrustOrganizationTrustSettingsModel) MarshalJSONForUpdate(state ZeroTrustOrganizationTrustSettingsModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

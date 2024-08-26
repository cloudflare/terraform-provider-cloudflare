// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package account

import (
	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/accounts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AccountsResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[AccountsResultDataSourceModel] `json:"result,computed"`
}

type AccountsDataSourceModel struct {
	Direction types.String                                                `tfsdk:"direction" query:"direction"`
	Name      types.String                                                `tfsdk:"name" query:"name"`
	MaxItems  types.Int64                                                 `tfsdk:"max_items"`
	Result    customfield.NestedObjectList[AccountsResultDataSourceModel] `tfsdk:"result"`
}

func (m *AccountsDataSourceModel) toListParams() (params accounts.AccountListParams, diags diag.Diagnostics) {
	params = accounts.AccountListParams{}

	if !m.Direction.IsNull() {
		params.Direction = cloudflare.F(accounts.AccountListParamsDirection(m.Direction.ValueString()))
	}
	if !m.Name.IsNull() {
		params.Name = cloudflare.F(m.Name.ValueString())
	}

	return
}

type AccountsResultDataSourceModel struct {
	ID        types.String                     `tfsdk:"id" json:"id,computed"`
	Name      types.String                     `tfsdk:"name" json:"name,computed"`
	CreatedOn timetypes.RFC3339                `tfsdk:"created_on" json:"created_on,computed"`
	Settings  *AccountsSettingsDataSourceModel `tfsdk:"settings" json:"settings"`
}

type AccountsSettingsDataSourceModel struct {
	AbuseContactEmail           types.String `tfsdk:"abuse_contact_email" json:"abuse_contact_email"`
	DefaultNameservers          types.String `tfsdk:"default_nameservers" json:"default_nameservers,computed"`
	EnforceTwofactor            types.Bool   `tfsdk:"enforce_twofactor" json:"enforce_twofactor,computed"`
	UseAccountCustomNSByDefault types.Bool   `tfsdk:"use_account_custom_ns_by_default" json:"use_account_custom_ns_by_default,computed"`
}

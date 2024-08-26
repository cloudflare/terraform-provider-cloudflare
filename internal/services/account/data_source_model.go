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

type AccountResultDataSourceEnvelope struct {
	Result AccountDataSourceModel `json:"result,computed"`
}

type AccountResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[AccountDataSourceModel] `json:"result,computed"`
}

type AccountDataSourceModel struct {
	AccountID types.String                     `tfsdk:"account_id" path:"account_id"`
	CreatedOn timetypes.RFC3339                `tfsdk:"created_on" json:"created_on,computed"`
	ID        types.String                     `tfsdk:"id" json:"id,computed"`
	Name      types.String                     `tfsdk:"name" json:"name,computed"`
	Settings  *AccountSettingsDataSourceModel  `tfsdk:"settings" json:"settings"`
	Filter    *AccountFindOneByDataSourceModel `tfsdk:"filter"`
}

func (m *AccountDataSourceModel) toReadParams() (params accounts.AccountGetParams, diags diag.Diagnostics) {
	params = accounts.AccountGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

func (m *AccountDataSourceModel) toListParams() (params accounts.AccountListParams, diags diag.Diagnostics) {
	params = accounts.AccountListParams{}

	if !m.Filter.Direction.IsNull() {
		params.Direction = cloudflare.F(accounts.AccountListParamsDirection(m.Filter.Direction.ValueString()))
	}
	if !m.Filter.Name.IsNull() {
		params.Name = cloudflare.F(m.Filter.Name.ValueString())
	}

	return
}

type AccountSettingsDataSourceModel struct {
	AbuseContactEmail           types.String `tfsdk:"abuse_contact_email" json:"abuse_contact_email"`
	DefaultNameservers          types.String `tfsdk:"default_nameservers" json:"default_nameservers,computed"`
	EnforceTwofactor            types.Bool   `tfsdk:"enforce_twofactor" json:"enforce_twofactor,computed"`
	UseAccountCustomNSByDefault types.Bool   `tfsdk:"use_account_custom_ns_by_default" json:"use_account_custom_ns_by_default,computed"`
}

type AccountFindOneByDataSourceModel struct {
	Direction types.String `tfsdk:"direction" query:"direction"`
	Name      types.String `tfsdk:"name" query:"name"`
}

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package account

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AccountResultDataSourceEnvelope struct {
	Result AccountDataSourceModel `json:"result,computed"`
}

type AccountResultListDataSourceEnvelope struct {
	Result *[]*AccountDataSourceModel `json:"result,computed"`
}

type AccountDataSourceModel struct {
	AccountID types.String                     `tfsdk:"account_id" path:"account_id"`
	CreatedOn timetypes.RFC3339                `tfsdk:"created_on" json:"created_on,computed"`
	ID        types.String                     `tfsdk:"id" json:"id,computed"`
	Name      types.String                     `tfsdk:"name" json:"name,computed"`
	Settings  *AccountSettingsDataSourceModel  `tfsdk:"settings" json:"settings"`
	Filter    *AccountFindOneByDataSourceModel `tfsdk:"filter"`
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

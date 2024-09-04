// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package account

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AccountResultEnvelope struct {
	Result AccountModel `json:"result"`
}

type AccountModel struct {
	ID        types.String                                   `tfsdk:"id" json:"id,computed"`
	Type      types.String                                   `tfsdk:"type" json:"type,required"`
	Unit      *AccountUnitModel                              `tfsdk:"unit" json:"unit,optional"`
	Name      types.String                                   `tfsdk:"name" json:"name,computed_optional"`
	Settings  customfield.NestedObject[AccountSettingsModel] `tfsdk:"settings" json:"settings,computed_optional"`
	CreatedOn timetypes.RFC3339                              `tfsdk:"created_on" json:"created_on,computed" format:"date-time"`
}

type AccountUnitModel struct {
	ID types.String `tfsdk:"id" json:"id,computed_optional"`
}

type AccountSettingsModel struct {
	AbuseContactEmail           types.String `tfsdk:"abuse_contact_email" json:"abuse_contact_email,computed_optional"`
	DefaultNameservers          types.String `tfsdk:"default_nameservers" json:"default_nameservers,computed_optional"`
	EnforceTwofactor            types.Bool   `tfsdk:"enforce_twofactor" json:"enforce_twofactor,computed_optional"`
	UseAccountCustomNSByDefault types.Bool   `tfsdk:"use_account_custom_ns_by_default" json:"use_account_custom_ns_by_default,computed_optional"`
}

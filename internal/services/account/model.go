// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package account

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AccountResultEnvelope struct {
	Result AccountModel `json:"result,computed"`
}

type AccountModel struct {
	AccountID types.String          `tfsdk:"account_id" path:"account_id"`
	Name      types.String          `tfsdk:"name" json:"name"`
	Settings  *AccountSettingsModel `tfsdk:"settings" json:"settings"`
}

type AccountSettingsModel struct {
	AbuseContactEmail           types.String `tfsdk:"abuse_contact_email" json:"abuse_contact_email"`
	DefaultNameservers          types.String `tfsdk:"default_nameservers" json:"default_nameservers"`
	EnforceTwofactor            types.Bool   `tfsdk:"enforce_twofactor" json:"enforce_twofactor"`
	UseAccountCustomNSByDefault types.Bool   `tfsdk:"use_account_custom_ns_by_default" json:"use_account_custom_ns_by_default"`
}

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package account

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AccountResultEnvelope struct {
	Result AccountModel `json:"result,computed"`
}

type AccountModel struct {
	ID        types.String          `tfsdk:"id" json:"id,computed"`
	Type      types.String          `tfsdk:"type" json:"type"`
	Unit      *AccountUnitModel     `tfsdk:"unit" json:"unit"`
	Name      types.String          `tfsdk:"name" json:"name"`
	Settings  *AccountSettingsModel `tfsdk:"settings" json:"settings"`
	CreatedOn timetypes.RFC3339     `tfsdk:"created_on" json:"created_on,computed"`
}

type AccountUnitModel struct {
	ID types.String `tfsdk:"id" json:"id"`
}

type AccountSettingsModel struct {
	AbuseContactEmail           types.String `tfsdk:"abuse_contact_email" json:"abuse_contact_email"`
	DefaultNameservers          types.String `tfsdk:"default_nameservers" json:"default_nameservers"`
	EnforceTwofactor            types.Bool   `tfsdk:"enforce_twofactor" json:"enforce_twofactor"`
	UseAccountCustomNSByDefault types.Bool   `tfsdk:"use_account_custom_ns_by_default" json:"use_account_custom_ns_by_default"`
}

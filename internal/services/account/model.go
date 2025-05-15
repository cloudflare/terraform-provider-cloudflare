// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package account

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AccountResultEnvelope struct {
	Result AccountModel `json:"result"`
}

type AccountModel struct {
	ID        types.String                                   `tfsdk:"id" json:"id,computed"`
	Type      types.String                                   `tfsdk:"type" json:"type,required,no_refresh"`
	Unit      *AccountUnitModel                              `tfsdk:"unit" json:"unit,optional,no_refresh"`
	Name      types.String                                   `tfsdk:"name" json:"name,required"`
	Settings  customfield.NestedObject[AccountSettingsModel] `tfsdk:"settings" json:"settings,computed_optional"`
	CreatedOn timetypes.RFC3339                              `tfsdk:"created_on" json:"created_on,computed" format:"date-time"`
}

func (m AccountModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m AccountModel) MarshalJSONForUpdate(state AccountModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type AccountUnitModel struct {
	ID types.String `tfsdk:"id" json:"id,optional"`
}

type AccountSettingsModel struct {
	AbuseContactEmail types.String `tfsdk:"abuse_contact_email" json:"abuse_contact_email,optional"`
	EnforceTwofactor  types.Bool   `tfsdk:"enforce_twofactor" json:"enforce_twofactor,computed_optional"`
}

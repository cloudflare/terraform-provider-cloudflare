// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package sso_connector

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type SSOConnectorResultEnvelope struct {
	Result SSOConnectorModel `json:"result"`
}

type SSOConnectorModel struct {
	ID                types.String                                            `tfsdk:"id" json:"id,computed"`
	AccountID         types.String                                            `tfsdk:"account_id" path:"account_id,required"`
	EmailDomain       types.String                                            `tfsdk:"email_domain" json:"email_domain,required"`
	BeginVerification types.Bool                                              `tfsdk:"begin_verification" json:"begin_verification,computed_optional,no_refresh"`
	Enabled           types.Bool                                              `tfsdk:"enabled" json:"enabled,optional"`
	CreatedOn         timetypes.RFC3339                                       `tfsdk:"created_on" json:"created_on,computed" format:"date-time"`
	UpdatedOn         timetypes.RFC3339                                       `tfsdk:"updated_on" json:"updated_on,computed" format:"date-time"`
	Verification      customfield.NestedObject[SSOConnectorVerificationModel] `tfsdk:"verification" json:"verification,computed"`
}

func (m SSOConnectorModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m SSOConnectorModel) MarshalJSONForUpdate(state SSOConnectorModel) (data []byte, err error) {
	return apijson.MarshalForPatch(m, state)
}

type SSOConnectorVerificationModel struct {
	Code   types.String `tfsdk:"code" json:"code,computed"`
	Status types.String `tfsdk:"status" json:"status,computed"`
}

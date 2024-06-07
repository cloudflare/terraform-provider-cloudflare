// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package access_custom_page

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AccessCustomPageResultEnvelope struct {
	Result AccessCustomPageModel `json:"result,computed"`
}

type AccessCustomPageModel struct {
	UID        types.String `tfsdk:"uid" json:"uid,computed"`
	AccountID  types.String `tfsdk:"account_id" path:"account_id"`
	CustomHTML types.String `tfsdk:"custom_html" json:"custom_html"`
	Name       types.String `tfsdk:"name" json:"name"`
	Type       types.String `tfsdk:"type" json:"type"`
	AppCount   types.Int64  `tfsdk:"app_count" json:"app_count"`
}

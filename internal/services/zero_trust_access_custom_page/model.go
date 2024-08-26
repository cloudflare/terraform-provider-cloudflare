// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_custom_page

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustAccessCustomPageResultEnvelope struct {
	Result ZeroTrustAccessCustomPageModel `json:"result"`
}

type ZeroTrustAccessCustomPageModel struct {
	ID         types.String      `tfsdk:"id" json:"-,computed"`
	UID        types.String      `tfsdk:"uid" json:"uid"`
	AccountID  types.String      `tfsdk:"account_id" path:"account_id"`
	CustomHTML types.String      `tfsdk:"custom_html" json:"custom_html"`
	Name       types.String      `tfsdk:"name" json:"name"`
	Type       types.String      `tfsdk:"type" json:"type"`
	AppCount   types.Int64       `tfsdk:"app_count" json:"app_count"`
	CreatedAt  timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed"`
	UpdatedAt  timetypes.RFC3339 `tfsdk:"updated_at" json:"updated_at,computed"`
}

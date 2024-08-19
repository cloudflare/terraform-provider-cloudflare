// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_custom_page

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustAccessCustomPagesResultListDataSourceEnvelope struct {
	Result *[]*ZeroTrustAccessCustomPagesResultDataSourceModel `json:"result,computed"`
}

type ZeroTrustAccessCustomPagesDataSourceModel struct {
	AccountID types.String                                        `tfsdk:"account_id" path:"account_id"`
	MaxItems  types.Int64                                         `tfsdk:"max_items"`
	Result    *[]*ZeroTrustAccessCustomPagesResultDataSourceModel `tfsdk:"result"`
}

type ZeroTrustAccessCustomPagesResultDataSourceModel struct {
	Name      types.String      `tfsdk:"name" json:"name,computed"`
	Type      types.String      `tfsdk:"type" json:"type,computed"`
	AppCount  types.Int64       `tfsdk:"app_count" json:"app_count"`
	CreatedAt timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed"`
	UID       types.String      `tfsdk:"uid" json:"uid"`
	UpdatedAt timetypes.RFC3339 `tfsdk:"updated_at" json:"updated_at,computed"`
}

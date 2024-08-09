// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_tag

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustAccessTagsResultListDataSourceEnvelope struct {
	Result *[]*ZeroTrustAccessTagsResultDataSourceModel `json:"result,computed"`
}

type ZeroTrustAccessTagsDataSourceModel struct {
	AccountID types.String                                 `tfsdk:"account_id" path:"account_id"`
	MaxItems  types.Int64                                  `tfsdk:"max_items"`
	Result    *[]*ZeroTrustAccessTagsResultDataSourceModel `tfsdk:"result"`
}

type ZeroTrustAccessTagsResultDataSourceModel struct {
	Name      types.String      `tfsdk:"name" json:"name,computed"`
	AppCount  types.Int64       `tfsdk:"app_count" json:"app_count"`
	CreatedAt timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed"`
	UpdatedAt timetypes.RFC3339 `tfsdk:"updated_at" json:"updated_at,computed"`
}

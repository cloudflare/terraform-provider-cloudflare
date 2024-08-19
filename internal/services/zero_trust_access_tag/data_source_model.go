// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_tag

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustAccessTagResultDataSourceEnvelope struct {
	Result ZeroTrustAccessTagDataSourceModel `json:"result,computed"`
}

type ZeroTrustAccessTagResultListDataSourceEnvelope struct {
	Result *[]*ZeroTrustAccessTagDataSourceModel `json:"result,computed"`
}

type ZeroTrustAccessTagDataSourceModel struct {
	AccountID types.String                                `tfsdk:"account_id" path:"account_id"`
	TagName   types.String                                `tfsdk:"tag_name" path:"tag_name"`
	CreatedAt timetypes.RFC3339                           `tfsdk:"created_at" json:"created_at,computed"`
	Name      types.String                                `tfsdk:"name" json:"name,computed"`
	UpdatedAt timetypes.RFC3339                           `tfsdk:"updated_at" json:"updated_at,computed"`
	AppCount  types.Int64                                 `tfsdk:"app_count" json:"app_count"`
	Filter    *ZeroTrustAccessTagFindOneByDataSourceModel `tfsdk:"filter"`
}

type ZeroTrustAccessTagFindOneByDataSourceModel struct {
	AccountID types.String `tfsdk:"account_id" path:"account_id"`
}

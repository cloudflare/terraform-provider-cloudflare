// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package d1_database

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type D1DatabasesResultListDataSourceEnvelope struct {
	Result *[]*D1DatabasesResultDataSourceModel `json:"result,computed"`
}

type D1DatabasesDataSourceModel struct {
	AccountID types.String                         `tfsdk:"account_id" path:"account_id"`
	Name      types.String                         `tfsdk:"name" query:"name"`
	Page      types.Float64                        `tfsdk:"page" query:"page"`
	PerPage   types.Float64                        `tfsdk:"per_page" query:"per_page"`
	MaxItems  types.Int64                          `tfsdk:"max_items"`
	Result    *[]*D1DatabasesResultDataSourceModel `tfsdk:"result"`
}

type D1DatabasesResultDataSourceModel struct {
	CreatedAt timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed"`
	Name      types.String      `tfsdk:"name" json:"name"`
	UUID      types.String      `tfsdk:"uuid" json:"uuid,computed"`
	Version   types.String      `tfsdk:"version" json:"version"`
}

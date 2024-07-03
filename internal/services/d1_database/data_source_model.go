// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package d1_database

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type D1DatabaseResultDataSourceEnvelope struct {
	Result D1DatabaseDataSourceModel `json:"result,computed"`
}

type D1DatabaseResultListDataSourceEnvelope struct {
	Result *[]*D1DatabaseDataSourceModel `json:"result,computed"`
}

type D1DatabaseDataSourceModel struct {
	AccountID  types.String                        `tfsdk:"account_id" path:"account_id"`
	DatabaseID types.String                        `tfsdk:"database_id" path:"database_id"`
	CreatedAt  types.String                        `tfsdk:"created_at" json:"created_at"`
	FileSize   types.Float64                       `tfsdk:"file_size" json:"file_size"`
	Name       types.String                        `tfsdk:"name" json:"name"`
	NumTables  types.Float64                       `tfsdk:"num_tables" json:"num_tables"`
	UUID       types.String                        `tfsdk:"uuid" json:"uuid"`
	Version    types.String                        `tfsdk:"version" json:"version"`
	FindOneBy  *D1DatabaseFindOneByDataSourceModel `tfsdk:"find_one_by"`
}

type D1DatabaseFindOneByDataSourceModel struct {
	AccountID types.String  `tfsdk:"account_id" path:"account_id"`
	Name      types.String  `tfsdk:"name" query:"name"`
	Page      types.Float64 `tfsdk:"page" query:"page"`
	PerPage   types.Float64 `tfsdk:"per_page" query:"per_page"`
}

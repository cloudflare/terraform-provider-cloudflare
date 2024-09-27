// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package d1_database

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/d1"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type D1DatabaseResultDataSourceEnvelope struct {
	Result D1DatabaseDataSourceModel `json:"result,computed"`
}

type D1DatabaseResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[D1DatabaseDataSourceModel] `json:"result,computed"`
}

type D1DatabaseDataSourceModel struct {
	AccountID  types.String                        `tfsdk:"account_id" path:"account_id,optional"`
	DatabaseID types.String                        `tfsdk:"database_id" path:"database_id,optional"`
	FileSize   types.Float64                       `tfsdk:"file_size" json:"file_size,optional"`
	NumTables  types.Float64                       `tfsdk:"num_tables" json:"num_tables,optional"`
	CreatedAt  timetypes.RFC3339                   `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Name       types.String                        `tfsdk:"name" json:"name,computed"`
	UUID       types.String                        `tfsdk:"uuid" json:"uuid,computed"`
	Version    types.String                        `tfsdk:"version" json:"version,computed"`
	Filter     *D1DatabaseFindOneByDataSourceModel `tfsdk:"filter"`
}

func (m *D1DatabaseDataSourceModel) toReadParams(_ context.Context) (params d1.DatabaseGetParams, diags diag.Diagnostics) {
	params = d1.DatabaseGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

func (m *D1DatabaseDataSourceModel) toListParams(_ context.Context) (params d1.DatabaseListParams, diags diag.Diagnostics) {
	params = d1.DatabaseListParams{
		AccountID: cloudflare.F(m.Filter.AccountID.ValueString()),
	}

	if !m.Filter.Name.IsNull() {
		params.Name = cloudflare.F(m.Filter.Name.ValueString())
	}

	return
}

type D1DatabaseFindOneByDataSourceModel struct {
	AccountID types.String `tfsdk:"account_id" path:"account_id,required"`
	Name      types.String `tfsdk:"name" query:"name,optional"`
}

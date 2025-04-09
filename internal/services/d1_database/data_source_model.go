// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package d1_database

import (
  "context"

  "github.com/cloudflare/cloudflare-go/v4"
  "github.com/cloudflare/cloudflare-go/v4/d1"
  "github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
  "github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
  "github.com/hashicorp/terraform-plugin-framework/diag"
  "github.com/hashicorp/terraform-plugin-framework/types"
)

type D1DatabaseResultDataSourceEnvelope struct {
Result D1DatabaseDataSourceModel `json:"result,computed"`
}

type D1DatabaseDataSourceModel struct {
ID types.String `tfsdk:"id" path:"database_id,computed"`
DatabaseID types.String `tfsdk:"database_id" path:"database_id,optional"`
AccountID types.String `tfsdk:"account_id" path:"account_id,required"`
CreatedAt timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
FileSize types.Float64 `tfsdk:"file_size" json:"file_size,computed"`
Name types.String `tfsdk:"name" json:"name,computed"`
NumTables types.Float64 `tfsdk:"num_tables" json:"num_tables,computed"`
UUID types.String `tfsdk:"uuid" json:"uuid,computed"`
Version types.String `tfsdk:"version" json:"version,computed"`
ReadReplication customfield.NestedObject[D1DatabaseReadReplicationDataSourceModel] `tfsdk:"read_replication" json:"read_replication,computed"`
Filter *D1DatabaseFindOneByDataSourceModel `tfsdk:"filter"`
}

func (m *D1DatabaseDataSourceModel) toReadParams(_ context.Context) (params d1.DatabaseGetParams, diags diag.Diagnostics) {
  params = d1.DatabaseGetParams{
    AccountID: cloudflare.F(m.AccountID.ValueString()),
  }

  return
}

func (m *D1DatabaseDataSourceModel) toListParams(_ context.Context) (params d1.DatabaseListParams, diags diag.Diagnostics) {
  params = d1.DatabaseListParams{
    AccountID: cloudflare.F(m.AccountID.ValueString()),
  }

  if !m.Filter.Name.IsNull() {
    params.Name = cloudflare.F(m.Filter.Name.ValueString())
  }

  return
}

type D1DatabaseReadReplicationDataSourceModel struct {
Mode types.String `tfsdk:"mode" json:"mode,computed"`
}

type D1DatabaseFindOneByDataSourceModel struct {
Name types.String `tfsdk:"name" query:"name,optional"`
}

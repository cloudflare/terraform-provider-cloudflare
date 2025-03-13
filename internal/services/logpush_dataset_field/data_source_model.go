// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package logpush_dataset_field

import (
  "context"

  "github.com/cloudflare/cloudflare-go/v4"
  "github.com/cloudflare/cloudflare-go/v4/logpush"
  "github.com/hashicorp/terraform-plugin-framework/diag"
  "github.com/hashicorp/terraform-plugin-framework/types"
)

type LogpushDatasetFieldResultDataSourceEnvelope struct {
Result LogpushDatasetFieldDataSourceModel `json:"result,computed"`
}

type LogpushDatasetFieldDataSourceModel struct {
DatasetID types.String `tfsdk:"dataset_id" path:"dataset_id,required"`
AccountID types.String `tfsdk:"account_id" path:"account_id,optional"`
ZoneID types.String `tfsdk:"zone_id" path:"zone_id,optional"`
}

func (m *LogpushDatasetFieldDataSourceModel) toReadParams(_ context.Context) (params logpush.DatasetFieldGetParams, diags diag.Diagnostics) {
  params = logpush.DatasetFieldGetParams{

  }

  if !m.AccountID.IsNull() {
    params.AccountID = cloudflare.F(m.AccountID.ValueString())
  } else {
    params.ZoneID = cloudflare.F(m.ZoneID.ValueString())
  }

  return
}

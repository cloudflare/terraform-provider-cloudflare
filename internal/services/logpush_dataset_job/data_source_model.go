// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package logpush_dataset_job

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/logpush"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type LogpushDatasetJobResultDataSourceEnvelope struct {
	Result LogpushDatasetJobDataSourceModel `json:"result,computed"`
}

type LogpushDatasetJobDataSourceModel struct {
	DatasetID types.String `tfsdk:"dataset_id" path:"dataset_id,required"`
	AccountID types.String `tfsdk:"account_id" path:"account_id,optional"`
	ZoneID    types.String `tfsdk:"zone_id" path:"zone_id,optional"`
}

func (m *LogpushDatasetJobDataSourceModel) toReadParams(_ context.Context) (params logpush.DatasetJobGetParams, diags diag.Diagnostics) {
	params = logpush.DatasetJobGetParams{}

	if !m.AccountID.IsNull() {
		params.AccountID = cloudflare.F(m.AccountID.ValueString())
	} else {
		params.ZoneID = cloudflare.F(m.ZoneID.ValueString())
	}

	return
}

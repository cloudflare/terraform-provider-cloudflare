// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pipeline

import (
  "context"

  "github.com/cloudflare/cloudflare-go/v4"
  "github.com/cloudflare/cloudflare-go/v4/pipelines"
  "github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
  "github.com/hashicorp/terraform-plugin-framework/diag"
  "github.com/hashicorp/terraform-plugin-framework/types"
)

type PipelineResultDataSourceEnvelope struct {
Result PipelineDataSourceModel `json:"result,computed"`
}

type PipelineDataSourceModel struct {
AccountID types.String `tfsdk:"account_id" path:"account_id,required"`
PipelineName types.String `tfsdk:"pipeline_name" path:"pipeline_name,required"`
Endpoint types.String `tfsdk:"endpoint" json:"endpoint,computed"`
ID types.String `tfsdk:"id" json:"id,computed"`
Name types.String `tfsdk:"name" json:"name,computed"`
Version types.Float64 `tfsdk:"version" json:"version,computed"`
Destination customfield.NestedObject[PipelineDestinationDataSourceModel] `tfsdk:"destination" json:"destination,computed"`
Source customfield.NestedObjectList[PipelineSourceDataSourceModel] `tfsdk:"source" json:"source,computed"`
}

func (m *PipelineDataSourceModel) toReadParams(_ context.Context) (params pipelines.PipelineGetParams, diags diag.Diagnostics) {
  params = pipelines.PipelineGetParams{
    AccountID: cloudflare.F(m.AccountID.ValueString()),
  }

  return
}

type PipelineDestinationDataSourceModel struct {
Batch customfield.NestedObject[PipelineDestinationBatchDataSourceModel] `tfsdk:"batch" json:"batch,computed"`
Compression customfield.NestedObject[PipelineDestinationCompressionDataSourceModel] `tfsdk:"compression" json:"compression,computed"`
Format types.String `tfsdk:"format" json:"format,computed"`
Path customfield.NestedObject[PipelineDestinationPathDataSourceModel] `tfsdk:"path" json:"path,computed"`
Type types.String `tfsdk:"type" json:"type,computed"`
}

type PipelineDestinationBatchDataSourceModel struct {
MaxBytes types.Int64 `tfsdk:"max_bytes" json:"max_bytes,computed"`
MaxDurationS types.Float64 `tfsdk:"max_duration_s" json:"max_duration_s,computed"`
MaxRows types.Int64 `tfsdk:"max_rows" json:"max_rows,computed"`
}

type PipelineDestinationCompressionDataSourceModel struct {
Type types.String `tfsdk:"type" json:"type,computed"`
}

type PipelineDestinationPathDataSourceModel struct {
Bucket types.String `tfsdk:"bucket" json:"bucket,computed"`
Filename types.String `tfsdk:"filename" json:"filename,computed"`
Filepath types.String `tfsdk:"filepath" json:"filepath,computed"`
Prefix types.String `tfsdk:"prefix" json:"prefix,computed"`
}

type PipelineSourceDataSourceModel struct {
Format types.String `tfsdk:"format" json:"format,computed"`
Type types.String `tfsdk:"type" json:"type,computed"`
Authentication types.Bool `tfsdk:"authentication" json:"authentication,computed"`
CORS customfield.NestedObject[PipelineSourceCORSDataSourceModel] `tfsdk:"cors" json:"cors,computed"`
}

type PipelineSourceCORSDataSourceModel struct {
Origins customfield.List[types.String] `tfsdk:"origins" json:"origins,computed"`
}

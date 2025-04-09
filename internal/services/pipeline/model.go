// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pipeline

import (
  "github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
  "github.com/hashicorp/terraform-plugin-framework/types"
)

type PipelineResultEnvelope struct {
Result PipelineModel `json:"result"`
}

type PipelineModel struct {
ID types.String `tfsdk:"id" json:"id,computed"`
AccountID types.String `tfsdk:"account_id" path:"account_id,required"`
Name types.String `tfsdk:"name" json:"name,required"`
Destination *PipelineDestinationModel `tfsdk:"destination" json:"destination,required"`
Source *[]*PipelineSourceModel `tfsdk:"source" json:"source,required"`
Endpoint types.String `tfsdk:"endpoint" json:"endpoint,computed"`
Version types.Float64 `tfsdk:"version" json:"version,computed"`
}

func (m PipelineModel) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(m)
}

func (m PipelineModel) MarshalJSONForUpdate(state PipelineModel) (data []byte, err error) {
  return apijson.MarshalForUpdate(m, state)
}

type PipelineDestinationModel struct {
Batch *PipelineDestinationBatchModel `tfsdk:"batch" json:"batch,required"`
Compression *PipelineDestinationCompressionModel `tfsdk:"compression" json:"compression,required"`
Credentials *PipelineDestinationCredentialsModel `tfsdk:"credentials" json:"credentials,required"`
Format types.String `tfsdk:"format" json:"format,required"`
Path *PipelineDestinationPathModel `tfsdk:"path" json:"path,required"`
Type types.String `tfsdk:"type" json:"type,required"`
}

type PipelineDestinationBatchModel struct {
MaxBytes types.Int64 `tfsdk:"max_bytes" json:"max_bytes,computed_optional"`
MaxDurationS types.Float64 `tfsdk:"max_duration_s" json:"max_duration_s,computed_optional"`
MaxRows types.Int64 `tfsdk:"max_rows" json:"max_rows,computed_optional"`
}

type PipelineDestinationCompressionModel struct {
Type types.String `tfsdk:"type" json:"type,computed_optional"`
}

type PipelineDestinationCredentialsModel struct {
AccessKeyID types.String `tfsdk:"access_key_id" json:"access_key_id,required"`
Endpoint types.String `tfsdk:"endpoint" json:"endpoint,required"`
SecretAccessKey types.String `tfsdk:"secret_access_key" json:"secret_access_key,required"`
}

type PipelineDestinationPathModel struct {
Bucket types.String `tfsdk:"bucket" json:"bucket,required"`
Filename types.String `tfsdk:"filename" json:"filename,optional"`
Filepath types.String `tfsdk:"filepath" json:"filepath,optional"`
Prefix types.String `tfsdk:"prefix" json:"prefix,optional"`
}

type PipelineSourceModel struct {
Format types.String `tfsdk:"format" json:"format,required"`
Type types.String `tfsdk:"type" json:"type,required"`
Authentication types.Bool `tfsdk:"authentication" json:"authentication,optional"`
CORS *PipelineSourceCORSModel `tfsdk:"cors" json:"cors,optional"`
}

type PipelineSourceCORSModel struct {
Origins *[]types.String `tfsdk:"origins" json:"origins,optional"`
}

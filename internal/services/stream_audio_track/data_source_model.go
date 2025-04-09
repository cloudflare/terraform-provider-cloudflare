// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stream_audio_track

import (
  "context"

  "github.com/cloudflare/cloudflare-go/v4"
  "github.com/cloudflare/cloudflare-go/v4/stream"
  "github.com/hashicorp/terraform-plugin-framework/diag"
  "github.com/hashicorp/terraform-plugin-framework/types"
)

type StreamAudioTrackResultDataSourceEnvelope struct {
Result StreamAudioTrackDataSourceModel `json:"result,computed"`
}

type StreamAudioTrackDataSourceModel struct {
AccountID types.String `tfsdk:"account_id" path:"account_id,required"`
Identifier types.String `tfsdk:"identifier" path:"identifier,required"`
Default types.Bool `tfsdk:"default" json:"default,computed"`
Label types.String `tfsdk:"label" json:"label,computed"`
Status types.String `tfsdk:"status" json:"status,computed"`
UID types.String `tfsdk:"uid" json:"uid,computed"`
}

func (m *StreamAudioTrackDataSourceModel) toReadParams(_ context.Context) (params stream.AudioTrackGetParams, diags diag.Diagnostics) {
  params = stream.AudioTrackGetParams{
    AccountID: cloudflare.F(m.AccountID.ValueString()),
  }

  return
}

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stream_audio_track

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type StreamAudioTrackResultEnvelope struct {
	Result StreamAudioTrackModel `json:"result"`
}

type StreamAudioTrackModel struct {
	AccountID       types.String `tfsdk:"account_id" path:"account_id,required"`
	Identifier      types.String `tfsdk:"identifier" path:"identifier,required"`
	AudioIdentifier types.String `tfsdk:"audio_identifier" path:"audio_identifier,optional"`
	Label           types.String `tfsdk:"label" json:"label,optional"`
	Default         types.Bool   `tfsdk:"default" json:"default,computed_optional"`
	Status          types.String `tfsdk:"status" json:"status,computed"`
	UID             types.String `tfsdk:"uid" json:"uid,computed"`
}

func (m StreamAudioTrackModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m StreamAudioTrackModel) MarshalJSONForUpdate(state StreamAudioTrackModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

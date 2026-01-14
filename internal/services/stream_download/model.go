// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stream_download

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type StreamDownloadResultEnvelope struct {
	Result StreamDownloadModel `json:"result"`
}

type StreamDownloadModel struct {
	AccountID       types.String                                         `tfsdk:"account_id" path:"account_id,required"`
	Identifier      types.String                                         `tfsdk:"identifier" path:"identifier,required"`
	PercentComplete types.Float64                                        `tfsdk:"percent_complete" json:"percentComplete,computed,no_refresh"`
	Status          types.String                                         `tfsdk:"status" json:"status,computed,no_refresh"`
	URL             types.String                                         `tfsdk:"url" json:"url,computed,no_refresh"`
	Audio           customfield.NestedObject[StreamDownloadAudioModel]   `tfsdk:"audio" json:"audio,computed"`
	Default         customfield.NestedObject[StreamDownloadDefaultModel] `tfsdk:"default" json:"default,computed"`
}

func (m StreamDownloadModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m StreamDownloadModel) MarshalJSONForUpdate(state StreamDownloadModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type StreamDownloadAudioModel struct {
	PercentComplete types.Float64 `tfsdk:"percent_complete" json:"percentComplete,computed"`
	Status          types.String  `tfsdk:"status" json:"status,computed"`
	URL             types.String  `tfsdk:"url" json:"url,computed"`
}

type StreamDownloadDefaultModel struct {
	PercentComplete types.Float64 `tfsdk:"percent_complete" json:"percentComplete,computed"`
	Status          types.String  `tfsdk:"status" json:"status,computed"`
	URL             types.String  `tfsdk:"url" json:"url,computed"`
}

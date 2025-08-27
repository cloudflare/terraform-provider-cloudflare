// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stream_download

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type StreamDownloadResultEnvelope struct {
	Result StreamDownloadModel `json:"result"`
}

type StreamDownloadModel struct {
	AccountID  types.String `tfsdk:"account_id" path:"account_id,required"`
	Identifier types.String `tfsdk:"identifier" path:"identifier,required"`
}

func (m StreamDownloadModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m StreamDownloadModel) MarshalJSONForUpdate(state StreamDownloadModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

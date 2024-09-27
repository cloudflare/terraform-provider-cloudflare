// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stream_download

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type StreamDownloadResultEnvelope struct {
	Result StreamDownloadModel `json:"result"`
}

type StreamDownloadModel struct {
	AccountID  types.String `tfsdk:"account_id" path:"account_id,required"`
	Identifier types.String `tfsdk:"identifier" path:"identifier,required"`
}

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stream_key

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type StreamKeyResultEnvelope struct {
	Result StreamKeyModel `json:"result"`
}

type StreamKeyModel struct {
	ID        types.String      `tfsdk:"id" json:"id,computed"`
	AccountID types.String      `tfsdk:"account_id" path:"account_id,required"`
	Created   timetypes.RFC3339 `tfsdk:"created" json:"created,computed" format:"date-time"`
	Jwk       types.String      `tfsdk:"jwk" json:"jwk,computed"`
	Pem       types.String      `tfsdk:"pem" json:"pem,computed"`
}

func (m StreamKeyModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m StreamKeyModel) MarshalJSONForUpdate(state StreamKeyModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

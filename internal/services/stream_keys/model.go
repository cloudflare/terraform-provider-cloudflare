// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stream_keys

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type StreamKeysResultEnvelope struct {
	Result StreamKeysModel `json:"result"`
}

type StreamKeysModel struct {
	ID        types.String      `tfsdk:"id" json:"id,computed"`
	AccountID types.String      `tfsdk:"account_id" path:"account_id,required"`
	Created   timetypes.RFC3339 `tfsdk:"created" json:"created,computed" format:"date-time"`
	Jwk       types.String      `tfsdk:"jwk" json:"jwk,computed"`
	Pem       types.String      `tfsdk:"pem" json:"pem,computed"`
}

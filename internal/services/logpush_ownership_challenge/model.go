// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package logpush_ownership_challenge

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type LogpushOwnershipChallengeResultEnvelope struct {
	Result LogpushOwnershipChallengeModel `json:"result"`
}

type LogpushOwnershipChallengeModel struct {
	AccountID       types.String `tfsdk:"account_id" path:"account_id,optional"`
	ZoneID          types.String `tfsdk:"zone_id" path:"zone_id,optional"`
	DestinationConf types.String `tfsdk:"destination_conf" json:"destination_conf,required"`
	Filename        types.String `tfsdk:"filename" json:"filename,computed"`
	Message         types.String `tfsdk:"message" json:"message,computed"`
	Valid           types.Bool   `tfsdk:"valid" json:"valid,computed"`
}

func (m LogpushOwnershipChallengeModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m LogpushOwnershipChallengeModel) MarshalJSONForUpdate(state LogpushOwnershipChallengeModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

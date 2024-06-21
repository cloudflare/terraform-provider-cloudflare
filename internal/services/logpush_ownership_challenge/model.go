// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package logpush_ownership_challenge

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type LogpushOwnershipChallengeResultEnvelope struct {
	Result LogpushOwnershipChallengeModel `json:"result,computed"`
}

type LogpushOwnershipChallengeModel struct {
	AccountID       types.String `tfsdk:"account_id" path:"account_id"`
	ZoneID          types.String `tfsdk:"zone_id" path:"zone_id"`
	DestinationConf types.String `tfsdk:"destination_conf" json:"destination_conf"`
}

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package teams_location

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type TeamsLocationResultEnvelope struct {
	Result TeamsLocationModel `json:"result,computed"`
}

type TeamsLocationModel struct {
	ID            types.String                   `tfsdk:"id" json:"id,computed"`
	AccountID     types.String                   `tfsdk:"account_id" path:"account_id"`
	Name          types.String                   `tfsdk:"name" json:"name"`
	ClientDefault types.Bool                     `tfsdk:"client_default" json:"client_default"`
	EcsSupport    types.Bool                     `tfsdk:"ecs_support" json:"ecs_support"`
	Networks      *[]*TeamsLocationNetworksModel `tfsdk:"networks" json:"networks"`
}

type TeamsLocationNetworksModel struct {
	Network types.String `tfsdk:"network" json:"network"`
}

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package hyperdrive_configs

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type HyperdriveConfigsResultEnvelope struct {
	Result HyperdriveConfigsModel `json:"result,computed"`
}

type HyperdriveConfigsModel struct {
	AccountID    types.String                  `tfsdk:"account_id" path:"account_id"`
	HyperdriveID types.String                  `tfsdk:"hyperdrive_id" path:"hyperdrive_id"`
	Origin       *HyperdriveConfigsOriginModel `tfsdk:"origin" json:"origin"`
	ID           types.String                  `tfsdk:"id" json:"id,computed"`
}

type HyperdriveConfigsOriginModel struct {
	Password types.String `tfsdk:"password" json:"password"`
}

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ct_alerting

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type CTAlertingResultEnvelope struct {
	Result CTAlertingModel `json:"result"`
}

type CTAlertingModel struct {
	ID      types.String    `tfsdk:"id" json:"-,computed"`
	ZoneID  types.String    `tfsdk:"zone_id" path:"zone_id,required"`
	Enabled types.Bool      `tfsdk:"enabled" json:"enabled,required"`
	Emails  *[]types.String `tfsdk:"emails" json:"emails,optional"`
}

func (m CTAlertingModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m CTAlertingModel) MarshalJSONForUpdate(state CTAlertingModel) (data []byte, err error) {
	return apijson.MarshalForPatch(m, state)
}

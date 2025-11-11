// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package logpull_retention

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type LogpullRetentionResultEnvelope struct {
	Result LogpullRetentionModel `json:"result"`
}

type LogpullRetentionModel struct {
	ID     types.String `tfsdk:"id" json:"-,computed"`
	ZoneID types.String `tfsdk:"zone_id" path:"zone_id,required"`
	Flag   types.Bool   `tfsdk:"flag" json:"flag,optional"`
}

func (m LogpullRetentionModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m LogpullRetentionModel) MarshalJSONForUpdate(state LogpullRetentionModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

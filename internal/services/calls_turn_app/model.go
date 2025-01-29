// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package calls_turn_app

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type CallsTURNAppResultEnvelope struct {
	Result CallsTURNAppModel `json:"result"`
}

type CallsTURNAppModel struct {
	AccountID types.String      `tfsdk:"account_id" path:"account_id,required"`
	KeyID     types.String      `tfsdk:"key_id" path:"key_id,optional"`
	Name      types.String      `tfsdk:"name" json:"name,computed_optional"`
	Created   timetypes.RFC3339 `tfsdk:"created" json:"created,computed" format:"date-time"`
	Key       types.String      `tfsdk:"key" json:"key,computed"`
	Modified  timetypes.RFC3339 `tfsdk:"modified" json:"modified,computed" format:"date-time"`
	UID       types.String      `tfsdk:"uid" json:"uid,computed"`
}

func (m CallsTURNAppModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m CallsTURNAppModel) MarshalJSONForUpdate(state CallsTURNAppModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

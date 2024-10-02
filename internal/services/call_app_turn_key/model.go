// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package call_app_turn_key

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type CallAppTURNKeyResultEnvelope struct {
	Result CallAppTURNKeyModel `json:"result"`
}

type CallAppTURNKeyModel struct {
	AccountID types.String      `tfsdk:"account_id" path:"account_id,required"`
	KeyID     types.String      `tfsdk:"key_id" path:"key_id,optional"`
	Name      types.String      `tfsdk:"name" json:"name,computed_optional"`
	Created   timetypes.RFC3339 `tfsdk:"created" json:"created,computed" format:"date-time"`
	Key       types.String      `tfsdk:"key" json:"key,computed"`
	Modified  timetypes.RFC3339 `tfsdk:"modified" json:"modified,computed" format:"date-time"`
	UID       types.String      `tfsdk:"uid" json:"uid,computed"`
}

func (m CallAppTURNKeyModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m CallAppTURNKeyModel) MarshalJSONForUpdate(state CallAppTURNKeyModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package calls_sfu_app

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type CallsSFUAppResultEnvelope struct {
	Result CallsSFUAppModel `json:"result"`
}

type CallsSFUAppModel struct {
	AccountID types.String      `tfsdk:"account_id" path:"account_id,required"`
	AppID     types.String      `tfsdk:"app_id" path:"app_id,optional"`
	Name      types.String      `tfsdk:"name" json:"name,computed_optional"`
	Created   timetypes.RFC3339 `tfsdk:"created" json:"created,computed" format:"date-time"`
	Modified  timetypes.RFC3339 `tfsdk:"modified" json:"modified,computed" format:"date-time"`
	Secret    types.String      `tfsdk:"secret" json:"secret,computed"`
	UID       types.String      `tfsdk:"uid" json:"uid,computed"`
}

func (m CallsSFUAppModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m CallsSFUAppModel) MarshalJSONForUpdate(state CallsSFUAppModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

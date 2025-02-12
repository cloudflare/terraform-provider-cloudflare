// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dns_settings_internal_view

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type DNSSettingsInternalViewResultEnvelope struct {
	Result DNSSettingsInternalViewModel `json:"result"`
}

type DNSSettingsInternalViewModel struct {
	ID           types.String      `tfsdk:"id" json:"id,computed"`
	AccountID    types.String      `tfsdk:"account_id" path:"account_id,required"`
	Name         types.String      `tfsdk:"name" json:"name,optional"`
	Zones        *[]types.String   `tfsdk:"zones" json:"zones,optional"`
	CreatedTime  timetypes.RFC3339 `tfsdk:"created_time" json:"created_time,computed" format:"date-time"`
	ModifiedTime timetypes.RFC3339 `tfsdk:"modified_time" json:"modified_time,computed" format:"date-time"`
}

func (m DNSSettingsInternalViewModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m DNSSettingsInternalViewModel) MarshalJSONForUpdate(state DNSSettingsInternalViewModel) (data []byte, err error) {
	return apijson.MarshalForPatch(m, state)
}

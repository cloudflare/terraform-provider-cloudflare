// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package regional_hostname

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type RegionalHostnameResultEnvelope struct {
	Result RegionalHostnameModel `json:"result"`
}

type RegionalHostnameModel struct {
	ID        types.String                                                `tfsdk:"id" json:"-,computed"`
	Hostname  types.String                                                `tfsdk:"hostname" json:"hostname,required"`
	ZoneID    types.String                                                `tfsdk:"zone_id" path:"zone_id,required"`
	RegionKey types.String                                                `tfsdk:"region_key" json:"region_key,required"`
	CreatedOn timetypes.RFC3339                                           `tfsdk:"created_on" json:"created_on,computed" format:"date-time"`
	Success   types.Bool                                                  `tfsdk:"success" json:"success,computed"`
	Errors    customfield.NestedObjectList[RegionalHostnameErrorsModel]   `tfsdk:"errors" json:"errors,computed"`
	Messages  customfield.NestedObjectList[RegionalHostnameMessagesModel] `tfsdk:"messages" json:"messages,computed"`
}

func (m RegionalHostnameModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m RegionalHostnameModel) MarshalJSONForUpdate(state RegionalHostnameModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type RegionalHostnameErrorsModel struct {
	Code    types.Int64  `tfsdk:"code" json:"code,computed"`
	Message types.String `tfsdk:"message" json:"message,computed"`
}

type RegionalHostnameMessagesModel struct {
	Code    types.Int64  `tfsdk:"code" json:"code,computed"`
	Message types.String `tfsdk:"message" json:"message,computed"`
}

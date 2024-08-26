// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package regional_hostname

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type RegionalHostnameResultEnvelope struct {
	Result RegionalHostnameModel `json:"result"`
}

type RegionalHostnameModel struct {
	ID        types.String                                                `tfsdk:"id" json:"-,computed"`
	Hostname  types.String                                                `tfsdk:"hostname" json:"hostname"`
	ZoneID    types.String                                                `tfsdk:"zone_id" path:"zone_id"`
	RegionKey types.String                                                `tfsdk:"region_key" json:"region_key"`
	CreatedOn timetypes.RFC3339                                           `tfsdk:"created_on" json:"created_on,computed"`
	Success   types.Bool                                                  `tfsdk:"success" json:"success,computed"`
	Errors    customfield.NestedObjectList[RegionalHostnameErrorsModel]   `tfsdk:"errors" json:"errors,computed"`
	Messages  customfield.NestedObjectList[RegionalHostnameMessagesModel] `tfsdk:"messages" json:"messages,computed"`
}

type RegionalHostnameErrorsModel struct {
	Code    types.Int64  `tfsdk:"code" json:"code"`
	Message types.String `tfsdk:"message" json:"message"`
}

type RegionalHostnameMessagesModel struct {
	Code    types.Int64  `tfsdk:"code" json:"code"`
	Message types.String `tfsdk:"message" json:"message"`
}

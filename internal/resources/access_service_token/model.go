// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package access_service_token

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AccessServiceTokenResultEnvelope struct {
	Result AccessServiceTokenModel `json:"result,computed"`
}

type AccessServiceTokenModel struct {
	ID        types.String `tfsdk:"id" json:"id,computed"`
	AccountID types.String `tfsdk:"account_id" path:"account_id"`
	ZoneID    types.String `tfsdk:"zone_id" path:"zone_id"`
	Name      types.String `tfsdk:"name" json:"name"`
	Duration  types.String `tfsdk:"duration" json:"duration"`
}

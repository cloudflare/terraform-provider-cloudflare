// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_routing_dns

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type EmailRoutingDNSResultEnvelope struct {
	Result EmailRoutingDNSModel `json:"result"`
}

type EmailRoutingDNSModel struct {
	ID                types.String      `tfsdk:"id" json:"-,computed"`
	ZoneID            types.String      `tfsdk:"zone_id" path:"zone_id,required"`
	Name              types.String      `tfsdk:"name" json:"name,optional,no_refresh"`
	Created           timetypes.RFC3339 `tfsdk:"created" json:"created,computed,no_refresh" format:"date-time"`
	Enabled           types.Bool        `tfsdk:"enabled" json:"enabled,computed,no_refresh"`
	Modified          timetypes.RFC3339 `tfsdk:"modified" json:"modified,computed,no_refresh" format:"date-time"`
	SkipWizard        types.Bool        `tfsdk:"skip_wizard" json:"skip_wizard,computed,no_refresh"`
	Status            types.String      `tfsdk:"status" json:"status,computed,no_refresh"`
	SupportSubaddress types.Bool        `tfsdk:"support_subaddress" json:"support_subaddress,computed,no_refresh"`
	Tag               types.String      `tfsdk:"tag" json:"tag,computed,no_refresh"`
}

func (m EmailRoutingDNSModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m EmailRoutingDNSModel) MarshalJSONForUpdate(state EmailRoutingDNSModel) (data []byte, err error) {
	return apijson.MarshalForPatch(m, state)
}

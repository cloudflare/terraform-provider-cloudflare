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
	ID         types.String      `tfsdk:"id" json:"-,computed"`
	ZoneID     types.String      `tfsdk:"zone_id" path:"zone_id,required"`
	Name       types.String      `tfsdk:"name" json:"name,required"`
	Created    timetypes.RFC3339 `tfsdk:"created" json:"created,computed" format:"date-time"`
	Enabled    types.Bool        `tfsdk:"enabled" json:"enabled,computed"`
	Modified   timetypes.RFC3339 `tfsdk:"modified" json:"modified,computed" format:"date-time"`
	SkipWizard types.Bool        `tfsdk:"skip_wizard" json:"skip_wizard,computed"`
	Status     types.String      `tfsdk:"status" json:"status,computed"`
	Tag        types.String      `tfsdk:"tag" json:"tag,computed"`
}

func (m EmailRoutingDNSModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m EmailRoutingDNSModel) MarshalJSONForUpdate(state EmailRoutingDNSModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

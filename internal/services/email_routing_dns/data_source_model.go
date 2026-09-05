// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_routing_dns

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/email_routing"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type EmailRoutingDNSResultDataSourceEnvelope struct {
	Result EmailRoutingDNSDataSourceModel `json:"result,computed"`
}

type EmailRoutingDNSDataSourceModel struct {
	ID        types.String                                                    `tfsdk:"id" path:"zone_id,computed"`
	ZoneID    types.String                                                    `tfsdk:"zone_id" path:"zone_id,required"`
	Subdomain types.String                                                    `tfsdk:"subdomain" query:"subdomain,optional"`
	DNS       customfield.NestedObjectList[EmailRoutingDNSDNSDataSourceModel] `tfsdk:"dns" json:"dns,computed"`
}

func (m *EmailRoutingDNSDataSourceModel) toReadParams(_ context.Context) (params email_routing.DNSGetParams, diags diag.Diagnostics) {
	params = email_routing.DNSGetParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	if !m.Subdomain.IsNull() {
		params.Subdomain = cloudflare.F(m.Subdomain.ValueString())
	}

	return
}

type EmailRoutingDNSDNSDataSourceModel struct {
	Content  types.String  `tfsdk:"content" json:"content,computed"`
	Name     types.String  `tfsdk:"name" json:"name,computed"`
	Priority types.Float64 `tfsdk:"priority" json:"priority,computed"`
	TTL      types.Float64 `tfsdk:"ttl" json:"ttl,computed"`
	Type     types.String  `tfsdk:"type" json:"type,computed"`
}

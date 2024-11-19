// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_routing_dns

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v3"
	"github.com/cloudflare/cloudflare-go/v3/email_routing"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type EmailRoutingDNSDataSourceModel struct {
	ZoneID    types.String `tfsdk:"zone_id" path:"zone_id,required"`
	Subdomain types.String `tfsdk:"subdomain" query:"subdomain,optional"`
}

func (m *EmailRoutingDNSDataSourceModel) toReadParams(_ context.Context) (params email_routing.DNSGetParams, diags diag.Diagnostics) {
	params = email_routing.DNSGetParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	return
}

package zone_dnssec

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v5"
	"github.com/cloudflare/cloudflare-go/v5/dns"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/logging"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

func disableDNSSECRecord(ctx context.Context, r *ZoneDNSSECResource, zoneID string, resp *resource.DeleteResponse) {
	var err error
	_, err = r.client.DNS.DNSSEC.Edit(
		ctx,
		dns.DNSSECEditParams{
			ZoneID: cloudflare.F(zoneID),
			Status: cloudflare.F[dns.DNSSECEditParamsStatus](dns.DNSSECEditParamsStatusDisabled),
		},
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to make http request", err.Error())
		return
	}
}

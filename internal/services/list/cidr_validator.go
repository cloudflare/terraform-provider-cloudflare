package list

import (
	"context"
	"fmt"
	"net"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ validator.String = listIPValidator{}

// listIPValidator validates that CIDRs or IPs are normalized.
type listIPValidator struct{}

func (v listIPValidator) Description(ctx context.Context) string {
	return v.MarkdownDescription(ctx)
}

func (v listIPValidator) MarkdownDescription(_ context.Context) string {
	return fmt.Sprintf("CIDRs or IPs must be normalised")
}

func (v listIPValidator) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}

	value := req.ConfigValue.ValueString()

	ip, netmask, err := net.ParseCIDR(value)
	if err != nil {
		ip = net.ParseIP(value)
		if ip == nil {
			return
		}
	}

	if netmask != nil {
		ones, bits := netmask.Mask.Size()
		if bits == 32 {
			// ipv4
			if ones == 32 {
				resp.Diagnostics.AddAttributeError(req.Path,
					"IPv4 /32 CIDRs should have the /32 suffix stripped",
					fmt.Sprintf("CIDR \"%s\" must be represented as \"%s\"", value, ip),
				)
				return
			}
		} else {
			// ipv6
			if ones == 128 {
				resp.Diagnostics.AddAttributeError(req.Path,
					"IPv6 /128 CIDRs should have the /128 suffix stripped",
					fmt.Sprintf("CIDR \"%s\" must be represented as \"%s\"", value, ip),
				)
				return
			}
		}
	}

	normalized := ip.String()
	if value != normalized {
		resp.Diagnostics.AddAttributeError(req.Path,
			"IP address must be normalized",
			fmt.Sprintf("IP address \"%s\" must be normalized: \"%s\"", value, normalized),
		)
	}
}

// ipValidator returns a validator that ensures IPs and CIDRs are normalized.
func ipValidator() validator.String {
	return listIPValidator{}
}

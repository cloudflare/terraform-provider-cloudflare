package access_rule

import (
	"context"
	"fmt"
	"net"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ validator.String = accessRuleIPv6Validator{}

// accessRuleIPv6Validator validates that if a string value is an IPv6 address,
// it must be in long normalized form.
type accessRuleIPv6Validator struct{}

func (v accessRuleIPv6Validator) Description(ctx context.Context) string {
	return v.MarkdownDescription(ctx)
}

func (v accessRuleIPv6Validator) MarkdownDescription(_ context.Context) string {
	return "if the value is an IPv6 address, it must be in long normalized form"
}

func (v accessRuleIPv6Validator) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}

	value := req.ConfigValue.ValueString()

	ip := net.ParseIP(value)
	if ip == nil {
		return
	}

	if ip.To4() != nil {
		return
	}

	normalized := ipv6ToLong(ip)
	if value != normalized {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"IPv6 address must be in long form",
			fmt.Sprintf("IPv6 address \"%s\" must be in long uncompressed form: \"%s\"", value, normalized),
		)
	}
}

// ipv6ToLong normalizes an IPv6 address to expanded form
func ipv6ToLong(ip net.IP) string {
	ip16 := ip.To16()
	return fmt.Sprintf("%02x%02x:%02x%02x:%02x%02x:%02x%02x:%02x%02x:%02x%02x:%02x%02x:%02x%02x",
		ip16[0], ip16[1], ip16[2], ip16[3],
		ip16[4], ip16[5], ip16[6], ip16[7],
		ip16[8], ip16[9], ip16[10], ip16[11],
		ip16[12], ip16[13], ip16[14], ip16[15])
}

// ipv6Validator returns a validator that ensures IPv6 addresses are in normalized form.
func ipv6Validator() validator.String {
	return accessRuleIPv6Validator{}
}

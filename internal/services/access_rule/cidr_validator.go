package access_rule

import (
	"context"
	"fmt"
	"net"
	"slices"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ validator.String = accessRuleCIDRValidator{}

var (
	validIPv4Prefixes = []int{16, 24}
	validIPv6Prefixes = []int{32, 48, 64}
)

// accessRuleCIDRValidator validates that CIDR ranges have allowed prefix lengths.
type accessRuleCIDRValidator struct{}

func (v accessRuleCIDRValidator) Description(ctx context.Context) string {
	return v.MarkdownDescription(ctx)
}

func (v accessRuleCIDRValidator) MarkdownDescription(_ context.Context) string {
	return fmt.Sprintf("CIDR must have a valid prefix length: IPv4 %v, IPv6 %v", validIPv4Prefixes, validIPv6Prefixes)
}

func (v accessRuleCIDRValidator) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}

	value := req.ConfigValue.ValueString()

	_, ipNet, err := net.ParseCIDR(value)
	if err != nil {
		return
	}

	prefixLen, _ := ipNet.Mask.Size()

	if ipNet.IP.To4() != nil {
		if !slices.Contains(validIPv4Prefixes, prefixLen) {
			resp.Diagnostics.AddAttributeError(
				req.Path,
				"Invalid IPv4 CIDR prefix length",
				fmt.Sprintf("IPv4 CIDR \"%s\" has prefix length /%d, but only %v are allowed", value, prefixLen, validIPv4Prefixes),
			)
		}
	} else {
		if !slices.Contains(validIPv6Prefixes, prefixLen) {
			resp.Diagnostics.AddAttributeError(
				req.Path,
				"Invalid IPv6 CIDR prefix length",
				fmt.Sprintf("IPv6 CIDR \"%s\" has prefix length /%d, but only %v are allowed", value, prefixLen, validIPv6Prefixes),
			)
		}
	}
}

// cidrValidator returns a validator that ensures CIDR ranges have valid prefix lengths.
func cidrValidator() validator.String {
	return accessRuleCIDRValidator{}
}

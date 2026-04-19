package list

import (
	"context"
	"fmt"
	"net"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ validator.String = listIPValidator{}

// listIPValidator rejects values that cannot be parsed as an IP or CIDR.
// Canonicalization of parseable values (stripping /32, /128, masking host
// bits, lowercasing IPv6) is handled by the ipNormalizer plan modifier.
type listIPValidator struct{}

func (listIPValidator) Description(_ context.Context) string {
	return "value must be a valid IP address or CIDR"
}

func (v listIPValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (listIPValidator) ValidateString(_ context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}
	value := req.ConfigValue.ValueString()
	if _, _, err := net.ParseCIDR(value); err == nil {
		return
	}
	if net.ParseIP(value) != nil {
		return
	}
	resp.Diagnostics.AddAttributeError(req.Path,
		"Invalid IP address or CIDR",
		fmt.Sprintf("%q is not a valid IP address or CIDR.", value),
	)
}

func ipValidator() validator.String {
	return listIPValidator{}
}

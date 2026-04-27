package list

import (
	"context"
	"net"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// normalizeListIP canonicalizes an IP or CIDR to match the form the Cloudflare
// API stores. Returns (normalized, true) on success, ("", false) if the value
// is not a parseable IP or CIDR.
//
//   - 1.2.3.4/32      → 1.2.3.4
//   - ::1/128         → ::1
//   - 192.168.1.5/24  → 192.168.1.0/24  (host bits masked)
//   - ::0001          → ::1
func normalizeListIP(value string) (string, bool) {
	if _, netmask, err := net.ParseCIDR(value); err == nil {
		ones, bits := netmask.Mask.Size()
		if (bits == 32 && ones == 32) || (bits == 128 && ones == 128) {
			return netmask.IP.String(), true
		}
		return netmask.String(), true
	}
	if ip := net.ParseIP(value); ip != nil {
		return ip.String(), true
	}
	return "", false
}

var _ planmodifier.String = listIPNormalizer{}

// listIPNormalizer rewrites the plan value for list item IPs to the canonical
// form stored by the Cloudflare API, so equivalent config (e.g. "1.2.3.4/32"
// vs "1.2.3.4") does not produce perpetual drift.
type listIPNormalizer struct{}

func (listIPNormalizer) Description(_ context.Context) string {
	return "Normalizes IPs and CIDRs (strips /32 and /128 suffixes, masks CIDR host bits, canonicalizes IPv6)."
}

func (m listIPNormalizer) MarkdownDescription(ctx context.Context) string {
	return m.Description(ctx)
}

func (listIPNormalizer) PlanModifyString(_ context.Context, req planmodifier.StringRequest, resp *planmodifier.StringResponse) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}
	value := req.ConfigValue.ValueString()
	normalized, ok := normalizeListIP(value)
	if !ok {
		return
	}
	if normalized != value {
		resp.PlanValue = types.StringValue(normalized)
	}
}

func ipNormalizer() planmodifier.String {
	return listIPNormalizer{}
}

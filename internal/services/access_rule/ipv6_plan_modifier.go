package access_rule

import (
	"context"
	"net"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
)

// ipv6CanonicalPlanModifier suppresses spurious diffs for the
// `configuration.value` attribute when both the prior state value and the
// configured value parse to the same IPv6 address or network. The Cloudflare
// API normalises IPv6 input (for example expanding `2001:db8::/32` to its
// long form), which would otherwise produce a no-op update on every plan.
type ipv6CanonicalPlanModifier struct{}

func ipv6CanonicalValue() planmodifier.String {
	return ipv6CanonicalPlanModifier{}
}

func (ipv6CanonicalPlanModifier) Description(_ context.Context) string {
	return "Suppresses diffs when prior state and config refer to the same IPv6 address or CIDR network."
}

func (ipv6CanonicalPlanModifier) MarkdownDescription(_ context.Context) string {
	return "Suppresses diffs when prior state and config refer to the same IPv6 address or CIDR network."
}

func (ipv6CanonicalPlanModifier) PlanModifyString(_ context.Context, req planmodifier.StringRequest, resp *planmodifier.StringResponse) {
	// Nothing to compare during create/destroy or when either side is unknown.
	if req.StateValue.IsNull() || req.StateValue.IsUnknown() {
		return
	}
	if req.PlanValue.IsNull() || req.PlanValue.IsUnknown() {
		return
	}

	stateStr := req.StateValue.ValueString()
	planStr := req.PlanValue.ValueString()
	if stateStr == planStr {
		return
	}

	if equalIPv6Value(stateStr, planStr) {
		// Keep the prior state value so Terraform reports no change.
		resp.PlanValue = req.StateValue
	}
}

// equalIPv6Value reports whether two strings refer to the same IPv6 address or
// IPv6 network. Returns false for IPv4 inputs so that the existing diff
// behaviour is preserved.
func equalIPv6Value(a, b string) bool {
	if strings.Contains(a, "/") || strings.Contains(b, "/") {
		_, netA, errA := net.ParseCIDR(a)
		_, netB, errB := net.ParseCIDR(b)
		if errA != nil || errB != nil {
			return false
		}
		if netA.IP.To4() != nil || netB.IP.To4() != nil {
			return false
		}
		sizeA, _ := netA.Mask.Size()
		sizeB, _ := netB.Mask.Size()
		return sizeA == sizeB && netA.IP.Equal(netB.IP)
	}

	ipA := net.ParseIP(a)
	ipB := net.ParseIP(b)
	if ipA == nil || ipB == nil {
		return false
	}
	if ipA.To4() != nil || ipB.To4() != nil {
		return false
	}
	return ipA.Equal(ipB)
}

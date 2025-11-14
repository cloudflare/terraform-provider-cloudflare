package utils

import (
	"net"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// NormalizeIPStringWithCIDR normalizes IP addresses to handle /32 (IPv4) and /128 (IPv6)
// CIDR notation equivalence. If the API value and config value are semantically equal
// (same IP with or without CIDR notation), the config value is preserved to avoid drift.
func NormalizeIPStringWithCIDR(apiValue *basetypes.StringValue, configValue basetypes.StringValue) {
	if apiValue == nil || apiValue.IsNull() || apiValue.IsUnknown() || configValue.IsNull() || configValue.IsUnknown() {
		return
	}

	apiStr := apiValue.ValueString()
	configStr := configValue.ValueString()

	// Normalize both to CIDR notation for comparison
	normalizedAPI := NormalizeIPCIDR(apiStr)
	normalizedConfig := NormalizeIPCIDR(configStr)

	// If they match after normalization, preserve the config value
	if normalizedAPI == normalizedConfig {
		*apiValue = configValue
	}
}

// NormalizeIPCIDR normalizes an IP address or CIDR to always include CIDR notation.
// Single IPv4 addresses get /32, single IPv6 addresses get /128.
// Addresses with existing CIDR notation are returned as-is.
func NormalizeIPCIDR(ipStr string) string {
	ipStr = strings.TrimSpace(ipStr)

	if strings.Contains(ipStr, "/") {
		return ipStr
	}

	ip := net.ParseIP(ipStr)
	if ip == nil {
		return ipStr
	}

	if ip.To4() != nil {
		return ipStr + "/32"
	}
	return ipStr + "/128"
}

package cloudflare

import (
	"fmt"
	"net"
	"strings"
)

// validateRecordType ensures that the cloudflare record type is valid
func validateRecordType(t string, proxied bool) error {
	switch t {
	case "A":
		return nil
	case "AAAA":
		return nil
	case "CNAME":
		return nil
	case "TXT":
		if !proxied {
			return nil
		}
	case "SRV":
		if !proxied {
			return nil
		}
	case "LOC":
		if !proxied {
			return nil
		}
	case "MX":
		if !proxied {
			return nil
		}
	case "NS":
		if !proxied {
			return nil
		}
	case "SPF":
		if !proxied {
			return nil
		}
	case "CAA":
		if !proxied {
			return nil
		}
	default:
		return fmt.Errorf(
			`Invalid type %q. Valid types are "A", "AAAA", "CNAME", "TXT", "SRV", "LOC", "MX", "NS", "SPF" or "CAA"`, t)
	}

	return fmt.Errorf("Type %q cannot be proxied", t)
}

// validateRecordName ensures that based on supplied record type, the name content matches
// Currently only validates A and AAAA types
func validateRecordName(t string, value string) error {
	switch t {
	case "A":
		// Must be ipv4 addr
		addr := net.ParseIP(value)
		if addr == nil || !strings.Contains(value, ".") {
			return fmt.Errorf("A record must be a valid IPv4 address, got: %q", value)
		}
	case "AAAA":
		// Must be ipv6 addr
		addr := net.ParseIP(value)
		if addr == nil || !strings.Contains(value, ":") {
			return fmt.Errorf("AAAA record must be a valid IPv6 address, got: %q", value)
		}
	case "TXT":
		// Must be printable ASCII
		for i := 0; i < len(value); i++ {
			char := value[i]
			if (char < 0x20) || (0x7F < char) {
				return fmt.Errorf("TXT record must contain printable ASCII, found: %q", char)
			}
		}
	}

	return nil
}

func validateStringIP(v interface{}, k string) (warnings []string, errors []error) {
	ip := net.ParseIP(v.(string))
	if ip == nil {
		errors = append(errors, fmt.Errorf("%q is not a valid IP: %q", k, v.(string)))
	}
	return
}

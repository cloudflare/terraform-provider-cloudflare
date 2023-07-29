package sdkv2provider

import (
	"fmt"
	"net"
	"net/url"
	"regexp"
	"strings"
)

var (
	allowedHTTPMethods = []string{"GET", "POST", "PUT", "DELETE", "PATCH", "HEAD", "_ALL_"}
	allowedSchemes     = []string{"HTTP", "HTTPS", "_ALL_"}

	// A typical Zone ID is a 32 characters long alpha-numeric
	// string that closely resembles an MD5 checksum.
	zoneIDRegexp = regexp.MustCompile(`^([0-9a-f]{32}|[0-9A-F]{32})$`)
)

// validateRecordType ensures that the cloudflare record type is valid.
func validateRecordType(t string, proxied bool) error {
	switch t {
	case "A", "AAAA", "CNAME":
		return nil
	case "TXT", "SRV", "LOC", "MX", "NS", "SPF", "CAA", "CERT", "DNSKEY", "DS", "NAPTR", "SMIMEA", "SSHFP", "TLSA", "URI", "PTR", "HTTPS", "SVCB":
		if ![]bool{proxied}[0] {
			return nil
		}
	default:
		return fmt.Errorf(
			`Invalid type %q. Valid types are "A", "AAAA", "CNAME", "TXT", "SRV", "LOC", "MX", "NS", "SPF", "CAA", "CERT", "DNSKEY", "DS", "NAPTR", "SMIMEA", "SSHFP", "TLSA", "URI", "PTR", "HTTPS", "SVCB".`, t)
	}

	return fmt.Errorf("type %q cannot be proxied", t)
}

// validateRecordContent ensures that the record's content is valid for the
// supplied record type. Currently only validates A and AAAA types.
func validateRecordContent(t string, value string) error {
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

// validateURL provides a method to test whether the provided string
// is a valid URL. Relying on `url.ParseRequestURI` isn't the most
// robust solution it will catch majority of the issues we're looking to
// handle here but there _could_ be edge cases.
func validateURL(v interface{}, k string) (s []string, errors []error) {
	if _, err := url.ParseRequestURI(v.(string)); err != nil {
		errors = append(errors, fmt.Errorf("%q: %w", k, err))
	}
	return
}

// validateZoneID ensures that the given Zone ID is valid.
func validateZoneID(value string) error {
	if matched := zoneIDRegexp.MatchString(value); !matched {
		return fmt.Errorf("must be a valid Zone ID, got: %s", value)
	}

	return nil
}

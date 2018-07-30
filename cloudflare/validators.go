package cloudflare

import (
	"fmt"
	"net"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"
)

var allowedHTTPMethods = []string{"GET", "POST", "PUT", "DELETE", "PATCH", "HEAD", "_ALL_"}
var allowedSchemes = []string{"HTTP", "HTTPS", "_ALL_"}

// validateRecordType ensures that the cloudflare record type is valid
func validateRecordType(t string, proxied bool) error {
	switch t {
	case "A", "AAAA", "CNAME":
		return nil
	case "TXT", "SRV", "LOC", "MX", "NS", "SPF", "CAA", "CERT", "DNSKEY", "DS", "NAPTR", "SMIMEA", "SSHFP", "TLSA", "URI":
		if !proxied {
			return nil
		}
	default:
		return fmt.Errorf(
			`Invalid type %q. Valid types are "A", "AAAA", "CNAME", "TXT", "SRV", "LOC", "MX", "NS", "SPF", "CAA", "CERT", "DNSKEY", "DS", "NAPTR", "SMIMEA", "SSHFP", "TLSA" or "URI".`, t)
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

// validateIntInSlice returns a SchemaValidateFunc which tests if the provided value
// is of type int
func validateIntInSlice(valid []int) schema.SchemaValidateFunc {
	return func(i interface{}, k string) ([]string, []error) {
		var es []error
		v, ok := i.(int)
		if !ok {
			es = append(es, fmt.Errorf("expected type of %q to be int", k))
			return nil, es
		}

		for _, str := range valid {
			if v == str {
				return nil, nil
			}
		}

		es = append(es, fmt.Errorf("expected %q to be one of %v, got %d", k, valid, v))
		return nil, es
	}
}

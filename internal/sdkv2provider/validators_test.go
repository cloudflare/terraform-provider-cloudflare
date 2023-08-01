package sdkv2provider

import (
	"testing"

	cloudflare "github.com/cloudflare/cloudflare-go"
)

func TestValidateRecordType(t *testing.T) {
	t.Parallel()

	validTypes := map[string]*bool{
		"A":     cloudflare.BoolPtr(true),
		"AAAA":  cloudflare.BoolPtr(true),
		"CNAME": cloudflare.BoolPtr(true),
		"TXT":   cloudflare.BoolPtr(false),
		"SRV":   cloudflare.BoolPtr(false),
		"LOC":   cloudflare.BoolPtr(false),
		"MX":    cloudflare.BoolPtr(false),
		"NS":    cloudflare.BoolPtr(false),
		"SPF":   cloudflare.BoolPtr(false),
	}
	for k, v := range validTypes {
		err := validateRecordType(k, *v)
		if err != nil {
			t.Fatalf("%s should be a valid record type: %s", k, err)
		}
	}

	invalidTypes := map[string]*bool{
		"a":     cloudflare.BoolPtr(false),
		"cName": cloudflare.BoolPtr(false),
		"txt":   cloudflare.BoolPtr(false),
		"SRv":   cloudflare.BoolPtr(false),
		"foo":   cloudflare.BoolPtr(false),
		"bar":   cloudflare.BoolPtr(false),
		"TXT":   cloudflare.BoolPtr(true),
		"SRV":   cloudflare.BoolPtr(true),
		"SPF":   cloudflare.BoolPtr(true),
	}
	for k, v := range invalidTypes {
		if err := validateRecordType(k, *v); err == nil {
			t.Fatalf("%s should be an invalid record type", k)
		}
	}
}

func TestValidateRecordName(t *testing.T) {
	t.Parallel()

	validNames := map[string]string{
		"A":    "192.168.0.1",
		"AAAA": "2001:0db8:0000:0000:0000:0000:0000:0000",
		"TXT":  " ",
	}

	for k, v := range validNames {
		if err := validateRecordContent(k, v); err != nil {
			t.Fatalf("%q should be valid content for type %q: %v", v, k, err)
		}
	}

	invalidNames := map[string]string{
		"A":    "terraform.io",
		"AAAA": "192.168.0.1",
		"TXT":  "\n",
	}
	for k, v := range invalidNames {
		if err := validateRecordContent(k, v); err == nil {
			t.Fatalf("%q should be invalid content for type %q", v, k)
		}
	}
}

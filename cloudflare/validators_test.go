package cloudflare

import "testing"

func TestValidateRecordType(t *testing.T) {
	validTypes := map[string]*bool{
		"A":     &[]bool{true}[0],
		"AAAA":  &[]bool{true}[0],
		"CNAME": &[]bool{true}[0],
		"TXT":   &[]bool{false}[0],
		"SRV":   &[]bool{false}[0],
		"LOC":   &[]bool{false}[0],
		"MX":    &[]bool{false}[0],
		"NS":    &[]bool{false}[0],
		"SPF":   &[]bool{false}[0],
	}
	for k, v := range validTypes {
		err := validateRecordType(k, *v)
		if err != nil {
			t.Fatalf("%s should be a valid record type: %s", k, err)
		}
	}

	invalidTypes := map[string]*bool{
		"a":     &[]bool{false}[0],
		"cName": &[]bool{false}[0],
		"txt":   &[]bool{false}[0],
		"SRv":   &[]bool{false}[0],
		"foo":   &[]bool{false}[0],
		"bar":   &[]bool{false}[0],
		"TXT":   &[]bool{true}[0],
		"SRV":   &[]bool{true}[0],
		"SPF":   &[]bool{true}[0],
	}
	for k, v := range invalidTypes {
		if err := validateRecordType(k, *v); err == nil {
			t.Fatalf("%s should be an invalid record type", k)
		}
	}
}

func TestValidateRecordName(t *testing.T) {
	validNames := map[string]string{
		"A":    "192.168.0.1",
		"AAAA": "2001:0db8:0000:0042:0000:8a2e:0370:7334",
		"TXT":  " ",
	}

	for k, v := range validNames {
		if err := validateRecordName(k, v); err != nil {
			t.Fatalf("%q should be a valid name for type %q: %v", v, k, err)
		}
	}

	invalidNames := map[string]string{
		"A":    "terraform.io",
		"AAAA": "192.168.0.1",
		"TXT":  "\n",
	}
	for k, v := range invalidNames {
		if err := validateRecordName(k, v); err == nil {
			t.Fatalf("%q should be an invalid name for type %q", v, k)
		}
	}
}

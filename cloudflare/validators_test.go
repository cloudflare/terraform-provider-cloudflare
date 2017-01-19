package cloudflare

import "testing"

func TestValidateRecordType(t *testing.T) {
	validTypes := map[string]bool{
		"A":     true,
		"AAAA":  true,
		"CNAME": true,
		"TXT":   false,
		"SRV":   false,
		"LOC":   false,
		"MX":    false,
		"NS":    false,
		"SPF":   false,
	}
	for k, v := range validTypes {
		err := validateRecordType(k, v)
		if err != nil {
			t.Fatalf("%s should be a valid record type: %s", k, err)
		}
	}

	invalidTypes := map[string]bool{
		"a":     false,
		"cName": false,
		"txt":   false,
		"SRv":   false,
		"foo":   false,
		"bar":   false,
		"TXT":   true,
		"SRV":   true,
		"SPF":   true,
	}
	for k, v := range invalidTypes {
		if err := validateRecordType(k, v); err == nil {
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

func TestValidatePageRuleStatus(t *testing.T) {
	validStatuses := []string{
		"active",
		"paused",
	}
	for _, v := range validStatuses {
		_, errors := validatePageRuleStatus(v, "status")
		if len(errors) != 0 {
			t.Fatalf("%q should be a valid page rule status: %q", v, errors)
		}
	}

	invalidStatuses := []string{
		"on",
		"live",
		"yes",
		"no",
		"true",
		"false",
		"running",
		"stopped",
	}
	for _, v := range invalidStatuses {
		_, errors := validatePageRuleStatus(v, "status")
		if len(errors) == 0 {
			t.Fatalf("%q should be an invalid page rule status: %q", v, errors)
		}
	}
}

func TestValidatePageRuleActionIDs(t *testing.T) {
	validActionIDs := []string{
		"always_online",
		"always_use_https",
		"automatic_https_rewrites",
		"browser_cache_ttl",
		"browser_check",
		"cache_level",
		"disable_apps",
		"disable_performance",
		"disable_railgun",
		"disable_security",
		"edge_cache_ttl",
		"email_obfuscation",
		"forwarding_url",
		"ip_geolocation",
		"opportunistic_encryption",
		"rocket_loader",
		"security_level",
		"server_side_exclude",
		"smart_errors",
		"ssl",
	}
	for _, v := range validActionIDs {
		_, errors := validatePageRuleActionID(v, "action")
		if len(errors) != 0 {
			t.Fatalf("%q should be a valid page rule action: %q", v, errors)
		}
	}

	invalidActionIDs := []string{
		"foo",
		"tls",
		"disable_foobar",
		"hunter2",
	}
	for _, v := range invalidActionIDs {
		_, errors := validatePageRuleActionID(v, "action")
		if len(errors) == 0 {
			t.Fatalf("%q should be an invalid page rule action: %q", v, errors)
		}
	}
}

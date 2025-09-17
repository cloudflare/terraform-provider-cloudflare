
resource "cloudflare_zero_trust_device_custom_profile" "%[1]s" {
	account_id                = "%[2]s"
	allow_mode_switch         = true
	allow_updates             = true
	allowed_to_leave          = true
	auto_connect              = 0
	captive_portal            = 5
	disable_auto_fallback     = true
	enabled                   = true
	match                     = "identity.email == \"foo@example.com\""
	name                      = "%[1]s"
	precedence                = %[6]d
	support_url               = "support_url"
	switch_locked             = true
	exclude_office_ips		  = false
	description               = "%[1]s"
}

resource "cloudflare_zero_trust_device_custom_profile_local_domain_fallback" "%[1]s" {
  account_id = "%[2]s"
  domains = [{
    description = "%[3]s"
    suffix      = "%[4]s"
    dns_server  = ["%[5]s"]
  }]
	policy_id = cloudflare_zero_trust_device_custom_profile.%[1]s.id
}

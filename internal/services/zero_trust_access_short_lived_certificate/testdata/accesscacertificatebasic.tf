
resource "cloudflare_zero_trust_access_application" "%[1]s" {
	name     = "%[1]s"
	%[3]s_id = "%[4]s"
	domain   = "%[1]s.%[2]s"
	type     = "self_hosted"
}

resource "cloudflare_zero_trust_access_short_lived_certificate" "%[1]s" {
  %[3]s_id       = "%[4]s"
  app_id = cloudflare_zero_trust_access_application.%[1]s.id
}

resource "cloudflare_access_application" "%[1]s" {
	name     = "%[1]s"
	%[3]s_id = "%[4]s"
	domain   = "%[1]s.%[2]s"
	type     = "self_hosted"
}

resource "cloudflare_access_ca_certificate" "%[1]s" {
  %[3]s_id       = "%[4]s"
  app_id = cloudflare_access_application.%[1]s.id
}
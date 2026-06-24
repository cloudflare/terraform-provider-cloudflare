resource "cloudflare_zero_trust_access_application" "%[1]s" {
  name    = "%[1]s"
  zone_id = "%[2]s"
  domain  = "%[1]s.%[3]s"
}

resource "cloudflare_zero_trust_access_short_lived_certificate" "%[1]s" {
  zone_id        = "%[2]s"
  application_id = cloudflare_zero_trust_access_application.%[1]s.id
}

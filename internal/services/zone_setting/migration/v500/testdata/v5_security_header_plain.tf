resource "cloudflare_zone_setting" {
  zone_id    = "%[2]s"
  setting_id = "security_header"
  value = {
    strict_transport_security = {
      enabled            = true
      include_subdomains = true
      max_age            = 86400
      nosniff            = false
      preload            = false
    }
  }
}

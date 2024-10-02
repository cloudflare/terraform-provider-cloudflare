resource "cloudflare_zone_setting" "%[1]s" {
  zone_id = "%[2]s"
  setting_id = "security_header"
  value = {
    strict_transport_security = {
      enabled = true
      include_subdomains = false
      max_age = 30
      nosniff = false
      preload = false
    }
  }
}

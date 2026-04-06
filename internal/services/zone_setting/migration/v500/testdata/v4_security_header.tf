resource "cloudflare_zone_settings_override" "%[1]s" {
  zone_id = "%[2]s"
  settings {
    security_header {
      enabled            = true
      max_age            = 86400
      include_subdomains = true
      preload            = false
      nosniff            = false
    }
  }
}

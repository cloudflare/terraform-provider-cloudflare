resource "cloudflare_zone_settings_override" "%[1]s" {
  zone_id = "%[2]s"
  settings {
    http3           = "on"
    min_tls_version = "1.2"
    brotli          = "on"
  }
}

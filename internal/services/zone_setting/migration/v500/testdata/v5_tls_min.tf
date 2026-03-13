resource "cloudflare_zone_setting" "%[1]s_min_tls_version" {
  zone_id    = "%[2]s"
  setting_id = "min_tls_version"
  value      = "1.2"
}

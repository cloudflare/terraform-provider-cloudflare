resource "cloudflare_zone_setting" "%[1]s" {
  zone_id = "%[2]s"
  setting_id = "min_tls_version"
  value = "1.2"
}

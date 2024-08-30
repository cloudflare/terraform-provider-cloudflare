
resource "cloudflare_custom_hostname" "%[2]s" {
  zone_id = "%[1]s"
  hostname = "%[2]s.%[3]s"
  ssl = {
  method = "txt"
}
  wait_for_ssl_pending_validation = true
}

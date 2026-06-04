resource "cloudflare_regional_hostname" "%[1]s" {
  zone_id    = "%[2]s"
  hostname   = "%[3]s"
  region_key = "ca"

  timeouts {
    create = "30s"
    update = "30s"
  }
}

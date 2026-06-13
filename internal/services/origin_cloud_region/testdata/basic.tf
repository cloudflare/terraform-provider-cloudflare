resource "cloudflare_origin_cloud_region" "%[1]s" {
  zone_id   = "%[2]s"
  origin_ip = "%[3]s"
  vendor    = "%[4]s"
  region    = "%[5]s"
}

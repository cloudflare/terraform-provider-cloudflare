resource "cloudflare_zone_setting" "%[1]s" {
  zone_id = "%[2]s"
  setting_id = "ssl_recommender"
  value = {
    enabled = true
  }
}

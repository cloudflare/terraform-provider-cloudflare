resource "cloudflare_workers_custom_domain" "test" {
  account_id  = "%s"
  zone_id     = "%s"
  hostname    = "%s"
  service     = "%s"
  environment = "%s"
}

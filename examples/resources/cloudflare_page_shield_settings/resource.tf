resource "cloudflare_page_shield_settings" "example_page_shield_settings" {
  zone_id                           = "023e105f4ecef8ad9ca31a8372d0c353"
  enabled                           = true
  use_cloudflare_reporting_endpoint = true
  use_connection_url_path           = false
}

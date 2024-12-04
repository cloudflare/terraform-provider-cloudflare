data "cloudflare_web_analytics_sites" "example_web_analytics_sites" {
  account_id = "023e105f4ecef8ad9ca31a8372d0c353"
  order_by = "host"
}

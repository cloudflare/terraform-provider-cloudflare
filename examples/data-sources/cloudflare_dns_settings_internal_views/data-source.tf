data "cloudflare_dns_settings_internal_views" "example_dns_settings_internal_views" {
  account_id = "023e105f4ecef8ad9ca31a8372d0c353"
  name = {
    contains = "view"
    endswith = "ew"
    exact = "my view"
    startswith = "my"
  }
  order = "name"
  zone_id = "ae29bea30e2e427ba9cd8d78b628177b"
  zone_name = "www.example.com"
}

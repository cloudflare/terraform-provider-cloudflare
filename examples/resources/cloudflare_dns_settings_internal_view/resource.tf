resource "cloudflare_dns_settings_internal_view" "example_dns_settings_internal_view" {
  account_id = "023e105f4ecef8ad9ca31a8372d0c353"
  name = "my view"
  zones = ["372e67954025e0ba6aaa6d586b9e0b59"]
}

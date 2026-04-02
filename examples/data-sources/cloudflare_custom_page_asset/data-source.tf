data "cloudflare_custom_page_asset" "example_custom_page_asset" {
  asset_name = "my_custom_error_page"
  account_id = "account_id"
  zone_id = "zone_id"
}

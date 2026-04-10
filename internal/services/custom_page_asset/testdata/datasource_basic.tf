resource "cloudflare_custom_page_asset" "%[1]s" {
  account_id  = "%[2]s"
  name        = "%[1]s"
  description = "Data source test asset"
  url         = "https://example.com/error.html"
}

data "cloudflare_custom_page_asset" "%[1]s" {
  account_id = "%[2]s"
  asset_name = cloudflare_custom_page_asset.%[1]s.name
}
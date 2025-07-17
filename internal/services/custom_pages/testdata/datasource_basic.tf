resource "cloudflare_custom_pages" "%[1]s" {
  account_id = "%[2]s"
  identifier = "basic_challenge"
  state      = "default"
  url        = ""
}

data "cloudflare_custom_pages" "%[1]s" {
  account_id = cloudflare_custom_pages.%[1]s.account_id
  identifier = cloudflare_custom_pages.%[1]s.identifier
}
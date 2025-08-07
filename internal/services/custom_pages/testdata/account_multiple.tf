resource "cloudflare_custom_pages" "%[1]s_challenge" {
  account_id = "%[2]s"
  identifier = "waf_challenge"
  state      = "default"
  url        = ""
}

resource "cloudflare_custom_pages" "%[1]s_block" {
  account_id = "%[2]s"
  identifier = "waf_block"
  state      = "default"
  url        = ""
}
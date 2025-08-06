resource "cloudflare_custom_pages" "%[1]s" {
  account_id = "%[2]s"
  identifier = "basic_challenge"
  state      = "default"
  url        = "http://www.example.com/custom_page"
}
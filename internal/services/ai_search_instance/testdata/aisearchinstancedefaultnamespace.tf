resource "cloudflare_ai_search_instance" "%[1]s" {
  account_id = "%[2]s"
  id         = "%[1]s"
}

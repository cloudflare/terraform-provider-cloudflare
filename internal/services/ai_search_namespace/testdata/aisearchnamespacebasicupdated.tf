resource "cloudflare_ai_search_namespace" "%[1]s" {
  account_id  = "%[2]s"
  name        = "%[1]s"
  description = "tf-acctest namespace updated"
}

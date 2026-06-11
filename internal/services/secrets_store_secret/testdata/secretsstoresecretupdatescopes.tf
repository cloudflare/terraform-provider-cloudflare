
resource "cloudflare_secrets_store_secret" "%[1]s" {
  account_id = "%[2]s"
  store_id   = "%[3]s"
  name       = "%[1]s"
  value      = "updated-secret-value"
  scopes     = ["ai_gateway", "workers"]
  comment    = "updated comment"
}

data "cloudflare_secrets_store_secrets" "example_secrets_store_secrets" {
  account_id = "985e105f4ecef8ad9ca31a8372d0c353"
  store_id = "023e105f4ecef8ad9ca31a8372d0c353"
  scopes = [["workers", "ai_gateway", "dex", "access"]]
  search = "search"
}

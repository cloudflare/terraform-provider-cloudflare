resource "cloudflare_ai_search_token" "example_ai_search_token" {
  account_id = "c3dc5f0b34a14ff8e1b3ec04895e1b22"
  cf_api_id = "a1b2c3d4e5f6"
  cf_api_key = "abc123"
  name = "my-token"
  legacy = true
}

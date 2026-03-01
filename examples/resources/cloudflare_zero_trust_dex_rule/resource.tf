resource "cloudflare_zero_trust_dex_rule" "example_zero_trust_dex_rule" {
  account_id = "01a7362d577a6c3019a474fd6f485823"
  match = "match"
  name = "name"
  description = "description"
}

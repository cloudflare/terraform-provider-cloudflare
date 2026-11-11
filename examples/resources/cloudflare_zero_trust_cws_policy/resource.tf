resource "cloudflare_zero_trust_cws_policy" "example_zero_trust_cws_policy" {
  account_id = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
  entrypoint = "entrypoint"
  filter = "x"
  name = "x"
  rego_source = "x"
  description = "description"
  tags = {
    foo = "string"
  }
}

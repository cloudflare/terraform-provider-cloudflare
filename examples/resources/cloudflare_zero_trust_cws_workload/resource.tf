resource "cloudflare_zero_trust_cws_workload" "example_zero_trust_cws_workload" {
  account_id = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
  filter = "x"
  name = "x"
  description = "description"
  tags = {
    foo = "string"
  }
}

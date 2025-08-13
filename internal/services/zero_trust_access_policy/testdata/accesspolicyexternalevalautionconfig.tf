resource "cloudflare_zero_trust_access_application" "%[1]s" {
  name       = "%[1]s"
  account_id = "%[3]s"
  domain     = "%[1]s.%[2]s"
  type	     = "self_hosted"
}

resource "cloudflare_zero_trust_access_policy" "%[1]s" {
  name           = "%[1]s"
  account_id     = "%[3]s"
  decision       = "allow"
  include = [{
    external_evaluation = {
      evaluate_url = "https://example.com"
      keys_url = "https://example.com/keys"
    }
  }]
}

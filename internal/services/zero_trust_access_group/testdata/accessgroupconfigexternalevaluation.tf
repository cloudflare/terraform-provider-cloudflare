resource "cloudflare_zero_trust_access_group" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[1]s"

  include = [
    {
      external_evaluation = {
        evaluate_url = "https://example.com/evaluate"
        keys_url     = "https://example.com/keys"
      }
    }
  ]
}
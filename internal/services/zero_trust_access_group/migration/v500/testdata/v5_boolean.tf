resource "cloudflare_zero_trust_access_group" "%s" {
  account_id = "%s"
  name       = "%s"

  include = [
    {
      everyone = {}
    },
  ]
}

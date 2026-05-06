resource "cloudflare_zero_trust_access_group" "%s" {
  account_id = "%s"
  name       = "%s"

  include = [
    {
      any_valid_service_token = {}
    },
  ]

  exclude = [
    {
      email = {
        email = "blocked@example.com"
      }
    },
  ]
}

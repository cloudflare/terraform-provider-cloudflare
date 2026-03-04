resource "cloudflare_zero_trust_access_group" "%s" {
  account_id = "%s"
  name       = "%s"

  include = [
    {
      email = {
        email = "user1@example.com"
      }
    },
    {
      email = {
        email = "user2@example.com"
      }
    },
    {
      email = {
        email = "user3@example.com"
      }
    },
  ]
}

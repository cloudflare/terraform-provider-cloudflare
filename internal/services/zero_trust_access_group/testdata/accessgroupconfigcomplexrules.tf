resource "cloudflare_zero_trust_access_group" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[1]s"

  include = [
    {
      email = {
        email = "include@example.com"
      }
    },
    {
      ip = {
        ip = "10.0.0.0/8"
      }
    }
  ]

  exclude = [
    {
      email = {
        email = "exclude@example.com"
      }
    }
  ]

  require = [
    {
      email_domain = {
        domain = "company.com"
      }
    }
  ]
}
resource "cloudflare_zero_trust_access_group" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[1]s"

  include = [
    {
      email = {
        email = "test@example.com"
      }
    },
    {
      ip = {
        ip = "192.0.2.0/24"
      }
    }
  ]

  exclude = [
    {
      ip = {
        ip = "192.0.2.100/32"
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
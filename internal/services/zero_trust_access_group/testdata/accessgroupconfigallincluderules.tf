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
      email_domain = {
        domain = "example.com"
      }
    },
    {
      ip = {
        ip = "192.0.2.1/32"
      }
    },
    {
      geo = {
        country_code = "US"
      }
    },
    {
      everyone = {}
    },
    {
      any_valid_service_token = {}
    },
    {
      certificate = {}
    },
    {
      auth_method = {
        auth_method = "swk"
      }
    }
  ]
}
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
      everyone = {}
    }
  ]

  exclude = [
    {
      geo = {
        country_code = "RU"
      }
    }
  ]

  require = [
    {
      auth_method = {
        auth_method = "swk"
      }
    }
  ]
}
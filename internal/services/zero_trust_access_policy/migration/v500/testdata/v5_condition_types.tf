resource "cloudflare_zero_trust_access_policy" "%[1]s" {
  account_id       = "%[2]s"
  name             = "%[1]s"
  decision         = "allow"
  session_duration = "24h"

  include = [
    { email = { email = "user@example.com" } },
    { email_domain = { domain = "example.com" } },
    { ip = { ip = "192.168.1.0/24" } },
    { everyone = {} },
    { certificate = {} },
    { any_valid_service_token = {} }
  ]

  exclude = [
    { geo = { country_code = "CN" } },
    { geo = { country_code = "RU" } }
  ]

  require = [
    { certificate = {} }
  ]
}

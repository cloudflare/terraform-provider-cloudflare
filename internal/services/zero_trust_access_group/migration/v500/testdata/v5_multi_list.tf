resource "cloudflare_zero_trust_access_group" "%s" {
  account_id = "%s"
  name       = "%s"

  # Order matches the transformation order in transform.go (lines 58-67):
  # email, email_domain, email_list, ip, ip_list, service_token, group, device_posture, login_method, geo
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
      ip = {
        ip = "192.168.1.0/24"
      }
    },
    {
      ip = {
        ip = "10.0.0.0/8"
      }
    },
  ]

  exclude = [
    {
      geo = {
        country_code = "US"
      }
    },
    {
      geo = {
        country_code = "CA"
      }
    },
  ]
}

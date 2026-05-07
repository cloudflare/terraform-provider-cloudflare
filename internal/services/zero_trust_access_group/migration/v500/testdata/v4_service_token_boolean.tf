resource "cloudflare_access_group" "%s" {
  account_id = "%s"
  name       = "%s"

  include {
    any_valid_service_token = true
  }

  exclude {
    email = ["blocked@example.com"]
  }
}

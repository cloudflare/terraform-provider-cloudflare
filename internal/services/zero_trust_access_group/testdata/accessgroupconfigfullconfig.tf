resource "cloudflare_zero_trust_access_group" "%[1]s" {
  account_id = "%[2]s"
  name = "%[1]s"

  include = [{
    email = { email = "%[3]s" }
    email_domain = { domain = "example.com" }
    common_names = ["common", "name"]
  }]

  require = [{
    email = { email = "%[3]s" }
  }]

  exclude = [{
    email = { email = "%[3]s" }
  }]
}

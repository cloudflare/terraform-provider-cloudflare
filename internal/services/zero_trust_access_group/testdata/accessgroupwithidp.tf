resource "cloudflare_zero_trust_access_identity_provider" "%[2]s" {
  account_id = "%[1]s"
  name = "%[2]s"
  type = "github"
  config = {
    client_id = "test"
    client_secret = "secret"
  }
}

resource "cloudflare_zero_trust_access_group" "%[2]s" {
  account_id = "%[1]s"
  name = "%[2]s"

  include = [{
    github_organization = {
      name                 = "%[3]s"
      team                = "%[4]s"
      identity_provider_id = cloudflare_zero_trust_access_identity_provider.%[2]s.id
    }
  }]
}

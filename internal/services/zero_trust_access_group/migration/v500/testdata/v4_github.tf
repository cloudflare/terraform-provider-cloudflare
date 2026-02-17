resource "cloudflare_access_identity_provider" "%[1]s_idp" {
  account_id = "%[2]s"
  name       = "%[3]s-idp"
  type       = "github"
  config {
    client_id     = "test"
    client_secret = "secret"
  }
}

resource "cloudflare_access_group" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[3]s"

  include {
    github {
      name                 = "my-org"
      teams                = ["team1", "team2"]
      identity_provider_id = cloudflare_access_identity_provider.%[1]s_idp.id
    }
  }

  depends_on = [cloudflare_access_identity_provider.%[1]s_idp]
}

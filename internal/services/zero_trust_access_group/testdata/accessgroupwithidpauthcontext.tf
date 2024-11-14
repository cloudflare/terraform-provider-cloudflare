resource "cloudflare_zero_trust_access_identity_provider" "%[2]s" {
  account_id = "%[1]s"
  name = "%[2]s"
  type = "azureAD"
  config = {
    client_id = "test"
    client_secret = "secret"
    directory_id = "foo"
  }
}

resource "cloudflare_zero_trust_access_group" "%[2]s" {
  account_id = "%[1]s"
  name = "%[2]s"

  include = [{
    any_valid_service_token = true
  }]

  require = [{
    auth_context = [{
      id = "%[3]s"
      ac_id = "%[4]s"
      identity_provider_id = cloudflare_zero_trust_access_identity_provider.%[2]s.id
    }
  }
}]

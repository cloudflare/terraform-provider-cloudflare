resource "cloudflare_zero_trust_access_application" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[1]s"
  type       = "saas"
  session_duration = "24h"

  saas_app = {
    auth_type         = "oidc"
    redirect_uris     = ["https://example.com/callback"]
    grant_types       = ["authorization_code", "hybrid"]
    scopes            = ["openid", "profile", "email"]
    app_launcher_url  = "https://example.com/app"
    group_filter_regex = ".*"
    allow_pkce_without_client_secret = false

    custom_claims = [{
      name     = "rank"
      required = true
      scope    = "profile"
      source = { name = "rank" }
    }]

    hybrid_and_implicit_options = {
      return_id_token_from_authorization_endpoint = true
      return_access_token_from_authorization_endpoint = true
    }
  }
  auto_redirect_to_identity = false
}

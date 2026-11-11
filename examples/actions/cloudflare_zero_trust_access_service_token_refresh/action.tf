resource "cloudflare_zero_trust_access_service_token" "ci" {
  account_id = "023e105f4ecef8ad9ca31a8372d0c353"
  name       = "CI/CD token"
  duration   = "8760h"
}

# Extend the token's expiry by its duration without changing the secret:
#   terraform apply -invoke=action.cloudflare_zero_trust_access_service_token_refresh.ci
action "cloudflare_zero_trust_access_service_token_refresh" "ci" {
  config {
    account_id       = cloudflare_zero_trust_access_service_token.ci.account_id
    service_token_id = cloudflare_zero_trust_access_service_token.ci.id
  }
}

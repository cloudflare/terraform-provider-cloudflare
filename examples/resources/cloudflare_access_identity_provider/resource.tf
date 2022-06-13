# one time pin
resource "cloudflare_access_identity_provider" "pin_login" {
  account_id = "1d5fdc9e88c8a8c4518b068cd94331fe"
  name       = "PIN login"
  type       = "onetimepin"
}

# oauth
resource "cloudflare_access_identity_provider" "github_oauth" {
  account_id = "1d5fdc9e88c8a8c4518b068cd94331fe"
  name       = "GitHub OAuth"
  type       = "github"
  config {
    client_id     = "example"
    client_secret = "secret_key"
  }
}

# saml
resource "cloudflare_access_identity_provider" "jumpcloud_saml" {
  account_id = "1d5fdc9e88c8a8c4518b068cd94331fe"
  name       = "JumpCloud SAML"
  type       = "saml"
  config {
    issuer_url      = "jumpcloud"
    sso_target_url  = "https://sso.myexample.jumpcloud.com/saml2/cloudflareaccess"
    attributes      = ["email", "username"]
    sign_request    = false
    idp_public_cert = "MIIDpDCCAoygAwIBAgIGAV2ka+55MA0GCSqGSIb3DQEBCwUAMIGSMQswCQ...GF/Q2/MHadws97cZg\nuTnQyuOqPuHbnN83d/2l1NSYKCbHt24o"
  }
}

# okta
resource "cloudflare_access_identity_provider" "okta" {
  account_id = "1d5fdc9e88c8a8c4518b068cd94331fe"
  name       = "Okta"
  type       = "okta"
  config {
    client_id     = "example"
    client_secret = "secret_key"
    api_token     = "okta_api_token"
  }
}

resource "cloudflare_zero_trust_access_identity_provider" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[1]s"
  type       = "saml"
  config = {
    issuer_url       = "jumpcloud"
    sso_target_url   = "https://sso.myexample.jumpcloud.com/saml2/cloudflareaccess"
    attributes       = ["email", "username"]
    sign_request     = true
    idp_public_certs = ["MIIDpDCCAoygAwIBAgIGAV2ka+55MA0GCSqGSIb3DQEBCwUAMIGSMQswCQYDVQQGEwJVUzETMBEGA1UEC.....GF/Q2/MHadws97cZguTnQyuOqPuHbnN83d/2l1NSYKCbHt24o"]
  }
}

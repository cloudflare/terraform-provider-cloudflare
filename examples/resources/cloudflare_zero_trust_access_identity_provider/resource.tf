resource "cloudflare_zero_trust_access_identity_provider" "example_zero_trust_access_identity_provider" {
  config = {
    claims = ["email_verified", "preferred_username", "custom_claim_name"]
    client_id = "<your client id>"
    client_secret = "<your client secret>"
    conditional_access_enabled = true
    directory_id = "<your azure directory uuid>"
    email_claim_name = "custom_claim_name"
    prompt = "login"
    support_groups = true
  }
  name = "Widget Corps IDP"
  type = {

  }
  zone_id = "zone_id"
  scim_config = {

  }
}

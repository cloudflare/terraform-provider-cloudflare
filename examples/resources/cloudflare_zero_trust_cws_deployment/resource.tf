resource "cloudflare_zero_trust_cws_deployment" "example_zero_trust_cws_deployment" {
  account_id = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
  basic_auth_credential_ids = ["182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"]
  default_mode = "ALLOW"
  name = "x"
  policy_set_ids = ["182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"]
  description = "description"
  domain_acl_id = "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"
  on_error_action = "BLOCK"
  rego_raise_error = true
  tags = {
    foo = "string"
  }
}

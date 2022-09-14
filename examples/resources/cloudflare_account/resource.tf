resource "cloudflare_account" "account" {
  name = "some-enterprise-account"
  type = "enterprise"
  enforce_twofactor = true
}
resource "cloudflare_account" "example" {
  name              = "some-enterprise-account"
  type              = "enterprise"
  enforce_twofactor = true
}

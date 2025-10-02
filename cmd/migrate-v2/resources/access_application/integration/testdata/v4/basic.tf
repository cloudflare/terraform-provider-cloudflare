resource "cloudflare_zero_trust_access_application" "basic_app" {
  account_id = "f037e56e89293a057740de681ac9abbe"
  name       = "Basic Application"
  domain     = "basic.example.com"
  type       = "self_hosted"
}
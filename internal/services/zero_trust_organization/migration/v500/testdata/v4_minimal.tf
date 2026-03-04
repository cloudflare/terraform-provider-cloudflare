# v4 minimal configuration - only required fields
resource "cloudflare_zero_trust_access_organization" "test" {
  account_id  = var.account_id
  name        = "Minimal Organization"
  auth_domain = "minimal.cloudflareaccess.com"
}

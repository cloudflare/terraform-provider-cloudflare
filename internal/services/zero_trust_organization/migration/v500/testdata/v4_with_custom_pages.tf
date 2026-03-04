# v4 configuration with custom_pages block
resource "cloudflare_access_organization" "test" {
  account_id  = var.account_id
  name        = "Test Organization"
  auth_domain = "test.cloudflareaccess.com"

  custom_pages {
    forbidden       = "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
    identity_denied = "yyyyyyyy-yyyy-yyyy-yyyy-yyyyyyyyyyyy"
  }
}

# v5 minimal configuration - boolean defaults will be added by migration
resource "cloudflare_zero_trust_organization" "test" {
  account_id  = var.account_id
  name        = "Minimal Organization"
  auth_domain = "minimal.cloudflareaccess.com"

  # These defaults are added by migration if missing in v4
  allow_authenticate_via_warp = false
  auto_redirect_to_identity   = false
  is_ui_read_only             = false
}

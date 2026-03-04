# v4 configuration using cloudflare_access_organization (deprecated name)
resource "cloudflare_access_organization" "test" {
  account_id                  = var.account_id
  name                        = "Test Organization"
  auth_domain                 = "test.cloudflareaccess.com"
  is_ui_read_only             = false
  auto_redirect_to_identity   = false
  session_duration            = "24h"
  warp_auth_session_duration  = "12h"
  allow_authenticate_via_warp = false
  # Note: user_seat_expiration_inactive_time removed - API requires minimum 730h (1 month)

  login_design {
    background_color = "#000000"
    text_color       = "#FFFFFF"
    logo_path        = "https://example.com/logo.png"
    header_text      = "Welcome"
    footer_text      = "Powered by Cloudflare"
  }
}

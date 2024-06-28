resource "cloudflare_access_organization" "example" {
  account_id                         = "f037e56e89293a057740de681ac9abbe"
  name                               = "example.cloudflareaccess.com"
  auth_domain                        = "example.cloudflareaccess.com"
  is_ui_read_only                    = false
  user_seat_expiration_inactive_time = "720h"
  auto_redirect_to_identity          = false

  login_design {
    background_color = "#ffffff"
    text_color       = "#000000"
    logo_path        = "https://example.com/logo.png"
    header_text      = "My header text"
    footer_text      = "My footer text"
  }
}

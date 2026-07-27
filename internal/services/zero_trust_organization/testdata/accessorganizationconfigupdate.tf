resource "cloudflare_zero_trust_organization" "%[1]s" {
  account_id                         = "%[2]s"
  name                               = "%[3]s"
  auth_domain                        = "%[1]s-%[3]s"
  is_ui_read_only                    = false
  user_seat_expiration_inactive_time = "2190h"
  auto_redirect_to_identity          = false
  session_duration                   = "24h"
  warp_auth_session_duration         = "48h"
  allow_authenticate_via_warp        = false

  login_design = {
    background_color = "#000000"
    text_color       = "#FFFFFF"
    logo_path        = "https://example.com/logo-v2.png"
    header_text      = "%[4]s"
    footer_text      = "Updated footer text"
  }
}

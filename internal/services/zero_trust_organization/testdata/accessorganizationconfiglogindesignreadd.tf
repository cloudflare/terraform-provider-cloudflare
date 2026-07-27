resource "cloudflare_zero_trust_organization" "%[1]s" {
  account_id       = "%[2]s"
  name             = "%[3]s"
  auth_domain      = "%[1]s-%[3]s"
  session_duration = "12h"

  login_design = {
    background_color = "#FF0000"
    text_color       = "#00FF00"
    logo_path        = "https://example.com/new-logo.png"
    header_text      = "Re-added header"
    footer_text      = "Re-added footer"
  }
}

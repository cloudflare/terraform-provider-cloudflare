
		resource "cloudflare_zero_trust_organization" "%[1]s" {
			account_id      = "%[2]s"
			name            = "terraform-cfapi.cloudflareaccess.com"
			auth_domain     = "%[1]s-terraform-cfapi.cloudflareaccess.com"
			is_ui_read_only = false
			user_seat_expiration_inactive_time = "1460h"
			auto_redirect_to_identity = false
			session_duration = "12h"
			warp_auth_session_duration = "36h"
			allow_authenticate_via_warp = false

			login_design = {
  background_color = "#FFFFFF"
				text_color       = "#000000"
				logo_path        = "https://example.com/logo.png"
				header_text      = "%[3]s"
				footer_text      = "My footer text"
}
		}

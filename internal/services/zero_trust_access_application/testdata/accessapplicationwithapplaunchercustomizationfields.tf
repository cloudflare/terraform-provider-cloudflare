
		resource "cloudflare_zero_trust_access_application" "%[1]s" {
			account_id       = "%[2]s"
			type             = "app_launcher"
			session_duration = "24h"
			app_launcher_visible = false
			app_launcher_logo_url = "https://www.cloudflare.com/img/logo-web-badges/cf-logo-on-white-bg.svg"
			bg_color = "#000000"
			header_bg_color = "#000000"

			footer_links =[ {
				name = "footer link"
				url = "https://www.cloudflare.com"
			}]


			landing_page_design =[ {
				title = "title"
				message = "message"
				button_color = "#000000"
				image_url = "https://www.cloudflare.com/img/logo-web-badges/cf-logo-on-white-bg.svg"
				button_text_color = "#000000"
			}]
	}
	
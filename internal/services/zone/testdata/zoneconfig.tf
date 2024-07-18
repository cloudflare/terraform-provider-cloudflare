
				resource "cloudflare_zone" "%[1]s" {
					account_id = "%[5]s"
					zone = "%[2]s"
					paused = %[3]s
					jump_start = %[4]s
				}

				resource "cloudflare_zone" "%[1]s" {
					account_id = "%[6]s"
					zone = "%[2]s"
					paused = %[3]s
					jump_start = %[4]s
					plan = "%[5]s"
					type = "%[7]s"
				}
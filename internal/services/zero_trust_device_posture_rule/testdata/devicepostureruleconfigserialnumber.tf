
resource "cloudflare_zero_trust_device_posture_rule" "%[1]s" {
	account_id                = "%[2]s"
	name                      = "%[1]s"
	type                      = "serial_number"
	description               = "My description"
	schedule                  = "24h"
	expiration                = "24h"
	match =[ {
		platform = "windows"
	}]
	input = {
  id = "asdf-123"
}
}

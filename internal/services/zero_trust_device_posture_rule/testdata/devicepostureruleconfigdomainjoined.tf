
resource "cloudflare_zero_trust_device_posture_rule" "%[1]s" {
	account_id                = "%[2]s"
	name                      = "%[1]s"
	type                      = "domain_joined"
	description               = "My description"
	schedule                  = "24h"
	expiration                = "24h"
	match = [{
		platform = "windows"
	}]
	input = {
  domain = "example.com"
}
}

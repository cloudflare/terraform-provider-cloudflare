resource "cloudflare_zero_trust_device_posture_rule" "%[1]s" {
	account_id                = "%[2]s"
	name                      = "%[1]s-updated"
	type                      = "os_version"
	description               = "Updated description"
	schedule                  = "1h"
	expiration                = "48h"
	match = [{
		platform = "mac"
	}]
	input = {
		version = "11.0.1"
		operator = ">="
		os_version_extra = "(updated)"
	}
}
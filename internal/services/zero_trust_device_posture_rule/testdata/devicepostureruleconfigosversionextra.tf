
resource "cloudflare_zero_trust_device_posture_rule" "%[1]s" {
	account_id                = "%[2]s"
	name                      = "%[1]s"
	type                      = "os_version"
	description               = "My description"
	schedule                  = "24h"
	expiration                = "24h"
	match =[ {
		platform = "mac"
	}]
	input = {
  version = "10.0.1"
		operator = "=="
		os_version_extra = "(a)"
}
}


resource "cloudflare_zero_trust_device_posture_rule" "%[1]s" {
	account_id                = "%[2]s"
	name                      = "%[1]s"
	type                      = "os_version"
	description               = "My description"
	schedule                  = "24h"
	expiration                = "24h"
	match =[ {
		platform = "linux"
	}]
	input = {
  version = "1.0.0"
        operator = "<"
		os_distro_name = "ubuntu"
		os_distro_revision = "1.0.0"
}
}

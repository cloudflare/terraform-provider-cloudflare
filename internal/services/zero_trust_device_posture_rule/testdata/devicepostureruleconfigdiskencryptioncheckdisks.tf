
resource "cloudflare_zero_trust_device_posture_rule" "%[1]s" {
	account_id                = "%[2]s"
	name                      = "%[1]s"
	type                      = "disk_encryption"
	description               = "My description"
	schedule                  = "24h"
	expiration                = "24h"
	match = [{
		platform = "mac"
	}]
	input = {
  require_all = false
		check_disks = ["C", "D"]
}
}

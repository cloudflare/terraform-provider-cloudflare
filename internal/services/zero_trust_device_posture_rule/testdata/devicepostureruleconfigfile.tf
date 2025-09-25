resource "cloudflare_zero_trust_device_posture_rule" "%[1]s" {
	account_id                = "%[2]s"
	name                      = "%[1]s"
	type                      = "file"
	description               = "File posture rule"
	schedule                  = "5m"
	expiration                = "1h"
	match = [{
		platform = "windows"
	}]
	input = {
		path = "C:\\Program Files\\Test\\test.exe"
		exists = true
		sha256 = "abcd1234567890abcdef1234567890abcdef1234567890abcdef1234567890ab"
	}
}
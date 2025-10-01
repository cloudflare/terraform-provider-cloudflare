resource "cloudflare_zero_trust_device_posture_rule" "%[1]s" {
	account_id                = "%[2]s"
	name                      = "%[1]s"
	type                      = "sentinelone_s2s"
	description               = "SentinelOne posture rule"
	schedule                  = "10m"
	expiration                = "2h"
	match = [{
		platform = "windows"
	}]
	input = {
		operating_system = "windows"
		path = "/Applications/SentinelOne Agent.app"
		active_threats = 0
		operator = "=="
		infected = false
		is_active = true
		network_status = "connected"
		operational_state = "na"
	}
}
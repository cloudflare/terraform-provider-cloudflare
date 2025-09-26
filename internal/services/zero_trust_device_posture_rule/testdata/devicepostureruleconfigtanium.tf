resource "cloudflare_zero_trust_device_posture_rule" "%[1]s" {
	account_id                = "%[2]s"
	name                      = "%[1]s"
	type                      = "tanium_s2s"
	description               = "Tanium posture rule"
	match = [{
		platform = "linux"
	}]
	input = {
		connection_id = "12345678-1234-1234-1234-123456789abc"
		eid_last_seen = "2023-01-01T00:00:00Z"
		risk_level = "low"
		score_operator = ">="
		total_score = 85.5
	}
}
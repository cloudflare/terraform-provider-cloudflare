resource "cloudflare_zero_trust_device_posture_rule" "%[1]s" {
	account_id                = "%[2]s"
	name                      = "%[1]s"
	type                      = "application"
	description               = "Application posture rule"
	match = [{
		platform = "mac"
	}]
	input = {
		path = "/Applications/Test.app"
		thumbprint = "abcd1234567890abcdef1234567890abcdef1234"
	}
}
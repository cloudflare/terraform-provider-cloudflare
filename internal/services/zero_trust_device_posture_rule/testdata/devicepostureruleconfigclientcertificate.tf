resource "cloudflare_zero_trust_device_posture_rule" "%[1]s" {
	account_id                = "%[2]s"
	name                      = "%[1]s"
	type                      = "client_certificate_v2"
	description               = "Client certificate posture rule"
	match = [{
		platform = "windows"
	}]
	input = {
		certificate_id = "12345678-1234-1234-1234-123456789abc"
		cn = "example.com"
		check_private_key = true
		extended_key_usage = ["clientAuth", "emailProtection"]
		locations = {
			trust_stores = ["system", "user"]
			paths = ["C:\\Certificates\\client.crt"]
		}
		subject_alternative_names = ["test.example.com", "alt.example.com"]
	}
}
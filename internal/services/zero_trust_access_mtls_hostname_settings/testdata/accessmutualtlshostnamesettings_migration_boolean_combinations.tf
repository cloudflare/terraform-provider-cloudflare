resource "cloudflare_zero_trust_access_mtls_hostname_settings" "%[1]s" {
	%[2]s_id = "%[3]s"
	settings {
		hostname = "%[4]s"
		client_certificate_forwarding = %[5]t
		china_network = %[6]t
	}
}
resource "cloudflare_zero_trust_access_mtls_hostname_settings" "%[1]s" {
	%[2]s_id = "%[3]s"
	settings {
		hostname = "%[4]s"
		china_network = false
		client_certificate_forwarding = false
	}
}
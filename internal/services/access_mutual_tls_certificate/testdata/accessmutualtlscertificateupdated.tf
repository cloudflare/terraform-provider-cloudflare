
resource "cloudflare_access_mutual_tls_certificate" "%[1]s" {
	name                 = "%[1]s"
	%[2]s_id             = "%[3]s"
	associated_hostnames = []
	certificate          = "%[4]s"
}


resource "cloudflare_zero_trust_access_mtls_certificate" "%[1]s" {
	name                 = "%[1]s"
	%[2]s_id             = "%[3]s"
	associated_hostnames = []
	certificate          = <<EOT
%[4]s
EOT
}

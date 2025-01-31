
resource "cloudflare_zero_trust_access_mtls_certificate" "%[1]s" {
	name                 = "%[1]s"
	%[2]s_id             = "%[3]s"
	associated_hostnames = ["%[5]s"]
	certificate          = <<EOT
%[4]s
EOT
}

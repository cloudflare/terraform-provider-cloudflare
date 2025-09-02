resource "cloudflare_access_mutual_tls_certificate" "%[1]s" {
	name                 = "%[1]s"
	%[2]s_id             = "%[3]s"
	certificate          = %[4]s
	associated_hostnames = [
		"%[1]s1.terraform.%[5]s",
		"%[1]s2.terraform.%[5]s",
		"%[1]s3.terraform.%[5]s",
		"%[1]s4.terraform.%[5]s"
	]
}
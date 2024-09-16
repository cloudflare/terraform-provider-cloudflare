
resource "cloudflare_origin_ca_certificate" "%[1]s" {
	csr                = <<EOT
%[3]sEOT
	hostnames          = [ "%[2]s" ]
	request_type       = "origin-rsa"
	requested_validity = 7
}

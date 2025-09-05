resource "cloudflare_access_mutual_tls_certificate" "%[1]s" {
	name        = "%[1]s"
	zone_id     = "%[2]s"
	certificate = %[3]s
}
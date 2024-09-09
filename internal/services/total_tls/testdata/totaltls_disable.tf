resource "cloudflare_total_tls" "%[1]s" {
	zone_id = "%[2]s"
	enabled = false
	certificate_authority = "google"
}

resource "cloudflare_zone_dnssec" "%s" {
	zone_id               = "%s"
	status                = "active"
	dnssec_multi_signer   = true
	dnssec_presigned      = false
	dnssec_use_nsec3      = true
}
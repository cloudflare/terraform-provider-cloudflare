resource "cloudflare_zone_dnssec" "%s" {
	zone_id               = "%s"
	status                = "active"
	dnssec_multi_signer   = true
}
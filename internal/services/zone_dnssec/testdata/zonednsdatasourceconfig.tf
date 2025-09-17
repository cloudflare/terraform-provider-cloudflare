resource "cloudflare_zone_dnssec" "%[1]s" {
  zone_id = "%[2]s"
  status = "active"
}

data "cloudflare_zone_dnssec" "%[1]s" {
  zone_id = "%[2]s"
  depends_on = [ "cloudflare_dns_zone_dnssec.%[1]s" ]
}

resource "cloudflare_dns_zone_transfers_outgoing" "%[1]s" {
  zone_id = "%[2]s"
  name = "%[3]s"
  peers = ["%[4]s"]
}

resource "cloudflare_dns_zone_transfers_peer" "%[1]s" {
  account_id = "%[4]s"
  name = "%[1]s"
}

resource "cloudflare_dns_zone_transfers_outgoing" "%[1]s" {
  zone_id = "%[2]s"
  name = "%[3]s"
  peers = [cloudflare_dns_zone_transfers_peer.%[1]s.id]
}

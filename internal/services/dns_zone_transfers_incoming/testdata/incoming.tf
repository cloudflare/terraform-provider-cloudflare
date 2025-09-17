resource "cloudflare_dns_zone_transfers_peer" "%[1]s" {
  account_id = "%[5]s"
  name = "%[1]s"
}

resource "cloudflare_dns_zone_transfers_incoming" "%[1]s" {
  zone_id = "%[2]s"
  auto_refresh_seconds = "%[3]d"
  name = "%[4]s"
  peers = [cloudflare_dns_zone_transfers_peer.%[1]s.id]
}

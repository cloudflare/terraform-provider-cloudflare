resource "cloudflare_dns_zone_transfers_incoming" "%[1]s" {
  zone_id = "%[2]s"
  auto_refresh_seconds = "%[3]d"
  name = "%[4]s"
  peers = ["%[5]s"]
}

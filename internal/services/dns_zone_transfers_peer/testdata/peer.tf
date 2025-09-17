resource "cloudflare_dns_zone_transfers_peer" "%[1]s" {
  account_id = "%[2]s"
  name = "%[3]s"
  ip = "%[4]s"
  ixfr_enable = "%[5]t"
  port = "%[6]d"
  tsig_id = "%[7]s"
}

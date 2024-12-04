resource "cloudflare_dns_zone_transfers_tsig" "%[1]s" {
  account_id = "%[2]s"
  name = "%[3]s"
  algo = "%[4]s"
  secret = "%[5]s"
}

resource "cloudflare_secondary_dns_peer" "%[1]s" {
  account_id = "%[2]s"
  name = "%[3]s"
  ip = "%[4]s"
  port = "%[5]d"
}

resource "cloudflare_secondary_dns_acl" "%[1]s" {
  account_id = "%[2]s"
  name = "%[3]s"
  ip_range = "%[4]s"
}

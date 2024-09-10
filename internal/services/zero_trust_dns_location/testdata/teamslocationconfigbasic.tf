
resource "cloudflare_zero_trust_dns_location" "%[1]s" {
  name        = "%[1]s"
  account_id  = "%[2]s"
}

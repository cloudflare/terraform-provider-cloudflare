resource "cloudflare_zero_trust_access_application" {
  account_id = "%[2]s"
  name       = "%[1]s"
  type       = "private_ip"
}

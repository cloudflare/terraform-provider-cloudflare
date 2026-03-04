resource "cloudflare_zero_trust_dex_rule" "%[1]s" {
  account_id  = "%[2]s"
  name        = "%[1]s"
  match       = "%[3]s"
  description = "%[4]s"
}

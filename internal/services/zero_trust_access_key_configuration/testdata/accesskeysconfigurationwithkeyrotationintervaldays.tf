
resource "cloudflare_zero_trust_access_key_configuration" "%[1]s" {
  account_id = "%[2]s"
  key_rotation_interval_days = "%[3]d"
}
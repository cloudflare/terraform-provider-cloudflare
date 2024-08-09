
resource "cloudflare_access_keys_configuration" "%[1]s" {
  account_id = "%[2]s"
  key_rotation_interval_days = "%[3]d"
}
resource "cloudflare_dns_settings_internal_view" "%[1]s" {
  account_id = "%[2]s"
  name = "%[3]s"
  zones = ["%[4]s"]
}

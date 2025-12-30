resource "cloudflare_leaked_credential_check" "%[2]s" {
  zone_id = "%[1]s"
  enabled = true
}
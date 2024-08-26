
resource "cloudflare_certificate_pack" "%[3]s" {
  zone_id = "%[1]s"
  type = "%[4]s"
  hosts = [
    "%[3]s.%[2]s",
    "%[2]s"
  ]
  validation_method = "txt"
  validity_days = 90
  certificate_authority = "lets_encrypt"
  cloudflare_branding = false
  wait_for_active_status = true
}
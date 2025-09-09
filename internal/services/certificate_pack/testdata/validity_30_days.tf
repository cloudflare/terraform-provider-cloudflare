resource "cloudflare_certificate_pack" "%[3]s" {
  zone_id = "%[1]s"
  type = "advanced"
  hosts = [
    "*.%[2]s",
    "%[2]s"
  ]
  validation_method = "txt"
  validity_days = 30
  certificate_authority = "google"
  cloudflare_branding = false
}
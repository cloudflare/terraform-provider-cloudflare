resource "cloudflare_certificate_pack" "%s" {
  zone_id               = "%s"
  type                  = "advanced"
  hosts                 = ["%s"]
  validation_method     = "http"
  validity_days         = 90
  certificate_authority = "google"
}

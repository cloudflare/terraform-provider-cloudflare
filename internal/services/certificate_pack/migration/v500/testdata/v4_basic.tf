resource "cloudflare_certificate_pack" "%s" {
  zone_id               = "%s"
  type                  = "advanced"
  hosts                 = ["%s", "*.%s"]
  validation_method     = "txt"
  validity_days         = 90
  certificate_authority = "lets_encrypt"
}

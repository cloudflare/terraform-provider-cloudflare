# Advanced certificate manager for DigiCert
resource "cloudflare_certificate_pack" "example" {
  zone_id               = "1d5fdc9e88c8a8c4518b068cd94331fe"
  type                  = "advanced"
  hosts                 = ["example.com", "sub.example.com"]
  validation_method     = "txt"
  validity_days         = 30
  certificate_authority = "digicert"
  cloudflare_branding   = false
}

# Advanced certificate manager for Let's Encrypt
resource "cloudflare_certificate_pack" "example" {
  zone_id                = "1d5fdc9e88c8a8c4518b068cd94331fe"
  type                   = "advanced"
  hosts                  = ["example.com", "*.example.com"]
  validation_method      = "http"
  validity_days          = 90
  certificate_authority  = "lets_encrypt"
  cloudflare_branding    = false
  wait_for_active_status = true
}

resource "cloudflare_zero_trust_access_application" "%[1]s" {
  zone_id          = "%[2]s"
  name             = "%[1]s"
  domain           = "%[1]s.%[3]s"
  type             = "self_hosted"
  session_duration = "24h"
  cors_headers = {
    allowed_methods = ["GET", "POST", "OPTIONS"]
    allowed_origins = ["https://example.com"]
    allow_credentials = true
    max_age = 10
  }
  auto_redirect_to_identity = false
}

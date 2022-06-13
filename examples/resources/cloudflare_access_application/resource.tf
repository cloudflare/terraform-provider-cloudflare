resource "cloudflare_access_application" "staging_app" {
  zone_id                   = "1d5fdc9e88c8a8c4518b068cd94331fe"
  name                      = "staging application"
  domain                    = "staging.example.com"
  type                      = "self_hosted"
  session_duration          = "24h"
  auto_redirect_to_identity = false
}

# With CORS configuration
resource "cloudflare_access_application" "staging_app" {
  zone_id          = "1d5fdc9e88c8a8c4518b068cd94331fe"
  name             = "staging application"
  domain           = "staging.example.com"
  type             = "self_hosted"
  session_duration = "24h"
  cors_headers {
    allowed_methods   = ["GET", "POST", "OPTIONS"]
    allowed_origins   = ["https://example.com"]
    allow_credentials = true
    max_age           = 10
  }
}

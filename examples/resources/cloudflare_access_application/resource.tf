resource "cloudflare_access_application" "staging_app" {
  zone_id                   = "0da42c8d2132a9ddaf714f9e7c920711"
  name                      = "staging application"
  domain                    = "staging.example.com"
  type                      = "self_hosted"
  session_duration          = "24h"
  auto_redirect_to_identity = false
  policies                  = [
      cloudflare_access_policy.example_1.id,
      cloudflare_access_policy.example_2.id
  ]
}

# With CORS configuration
resource "cloudflare_access_application" "staging_app" {
  zone_id          = "0da42c8d2132a9ddaf714f9e7c920711"
  name             = "staging application"
  domain           = "staging.example.com"
  type             = "self_hosted"
  session_duration = "24h"
  policies         = [
      cloudflare_access_policy.example_1.id,
      cloudflare_access_policy.example_2.id
  ]
  cors_headers {
    allowed_methods   = ["GET", "POST", "OPTIONS"]
    allowed_origins   = ["https://example.com"]
    allow_credentials = true
    max_age           = 10
  }
}

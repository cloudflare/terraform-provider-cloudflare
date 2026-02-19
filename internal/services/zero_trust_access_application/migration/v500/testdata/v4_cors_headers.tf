resource "cloudflare_access_application" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[1]s"
  domain     = "%[1]s.%[3]s"
  type       = "self_hosted"

  cors_headers {
    allowed_methods   = ["GET", "POST", "OPTIONS"]
    allowed_origins   = ["https://example.com"]
    allowed_headers   = ["Authorization", "Content-Type"]
    allow_credentials = true
    max_age           = 600
  }
}

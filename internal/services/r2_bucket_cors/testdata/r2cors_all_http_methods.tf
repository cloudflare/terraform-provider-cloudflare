resource "cloudflare_r2_bucket" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[1]s"
}

resource "cloudflare_r2_bucket_cors" "%[1]s" {
  account_id  = "%[2]s"
  bucket_name = cloudflare_r2_bucket.%[1]s.name

  rules = [{
    allowed = {
      methods = ["GET", "PUT", "POST", "DELETE", "HEAD"]
      origins = ["https://example.com"]
      headers = ["Content-Type", "Authorization"]
    }
    id               = "all-methods-rule"
    expose_headers   = ["ETag", "Content-Length"]
    max_age_seconds  = 3600
  }]
}

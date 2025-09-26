resource "cloudflare_r2_bucket" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[1]s"
}

resource "cloudflare_r2_bucket_cors" "%[1]s" {
  account_id  = "%[2]s"
  bucket_name = cloudflare_r2_bucket.%[1]s.name

  rules = [{
    allowed = {
      methods = ["GET", "POST", "PUT"]
      origins = ["https://full.com", "https://example.com"]
      headers = ["Content-Type", "Authorization", "X-Custom-Header"]
    }
    id               = "full-rule"
    expose_headers   = ["ETag", "Content-Length", "X-Response-Header"]
    max_age_seconds  = 7200
  }, {
    allowed = {
      methods = ["DELETE"]
      origins = ["https://admin.com"]
      headers = ["Authorization"]
    }
    id               = "admin-rule"
    expose_headers   = ["X-Admin-Response"]
    max_age_seconds  = 1800
  }]
}

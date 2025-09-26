resource "cloudflare_r2_bucket" "%[1]s" {
  account_id   = "%[2]s"
  name         = "%[1]s"
  jurisdiction = "fedramp"
}

resource "cloudflare_r2_bucket_cors" "%[1]s" {
  account_id   = "%[2]s"
  bucket_name  = cloudflare_r2_bucket.%[1]s.name
  jurisdiction = "fedramp"

  rules = [{
    allowed = {
      methods = ["GET", "POST"]
      origins = ["https://fedramp.example.com"]
      headers = ["Content-Type"]
    }
    id               = "fedramp-rule"
    expose_headers   = ["ETag"]
    max_age_seconds  = 3600
  }]
}

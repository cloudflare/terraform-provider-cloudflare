resource "cloudflare_r2_bucket" "%[1]s" {
  account_id   = "%[2]s"
  name         = "%[1]s"
  jurisdiction = "eu"
}

resource "cloudflare_r2_bucket_cors" "%[1]s" {
  account_id   = "%[2]s"
  bucket_name  = cloudflare_r2_bucket.%[1]s.name
  jurisdiction = "eu"

  rules = [{
    allowed = {
      methods = ["GET", "POST"]
      origins = ["https://eu.example.com"]
      headers = ["Content-Type"]
    }
    id               = "eu-rule"
    expose_headers   = ["ETag"]
    max_age_seconds  = 3600
  }]
}

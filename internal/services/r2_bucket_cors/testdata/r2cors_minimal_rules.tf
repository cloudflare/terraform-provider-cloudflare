resource "cloudflare_r2_bucket" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[1]s"
}

resource "cloudflare_r2_bucket_cors" "%[1]s" {
  account_id  = "%[2]s"
  bucket_name = cloudflare_r2_bucket.%[1]s.name

  rules = [{
    allowed = {
      methods = ["GET"]
      origins = ["*"]
      # No headers specified - testing minimal allowed block
    }
    # No id, expose_headers, or max_age_seconds - testing minimal rule
  }]
}

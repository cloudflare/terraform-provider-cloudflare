resource "cloudflare_r2_bucket" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[1]s"
}

resource "cloudflare_r2_bucket_lock" "%[1]s" {
  account_id  = "%[2]s"
  bucket_name = cloudflare_r2_bucket.%[1]s.name

  rules = [{
    id      = "date-lock"
    enabled = true
    prefix  = "temp/"
    condition = {
      type = "Date"
      date = "2025-12-31T23:59:59Z"
    }
  }]
}

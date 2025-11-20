resource "cloudflare_r2_bucket" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[1]s"
}

resource "cloudflare_r2_bucket_lock" "%[1]s" {
  account_id  = "%[2]s"
  bucket_name = cloudflare_r2_bucket.%[1]s.name

  rules = [{
    id      = "datasource-test-rule"
    enabled = true
    prefix  = "test-data/"
    condition = {
      type            = "Age"
      max_age_seconds = 604800
    }
  }]
}

data "cloudflare_r2_bucket_lock" "%[1]s" {
  account_id  = "%[2]s"
  bucket_name = cloudflare_r2_bucket_lock.%[1]s.bucket_name
}
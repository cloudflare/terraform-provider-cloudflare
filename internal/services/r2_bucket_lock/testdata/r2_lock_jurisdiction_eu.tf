resource "cloudflare_r2_bucket" "%[1]s" {
  account_id   = "%[2]s"
  name         = "%[1]s"
  jurisdiction = "eu"
}

resource "cloudflare_r2_bucket_lock" "%[1]s" {
  account_id   = "%[2]s"
  bucket_name  = cloudflare_r2_bucket.%[1]s.name
  jurisdiction = "eu"

  rules = [{
    id      = "eu-compliance-lock"
    enabled = true
    prefix  = "gdpr/"
    condition = {
      type            = "Age"
      max_age_seconds = 2592000
    }
  }]
}

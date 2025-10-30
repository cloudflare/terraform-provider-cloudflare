resource "cloudflare_r2_bucket" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[1]s"
}

resource "cloudflare_r2_bucket_lock" "%[1]s" {
  account_id  = "%[2]s"
  bucket_name = cloudflare_r2_bucket.%[1]s.name

  rules = [{
    id      = "archive-lock"
    enabled = true
    prefix  = "archive/"
    condition = {
      type            = "Age"
      max_age_seconds = 604800
    }
  }, {
    id      = "indefinite-lock"
    enabled = false
    prefix  = "permanent/"
    condition = {
      type = "Indefinite"
    }
  }, {
    id      = "retention-rule"
    enabled = true
    prefix  = "documents/"
    condition = {
      type            = "Age"
      max_age_seconds = 172800
    }
  }]
}

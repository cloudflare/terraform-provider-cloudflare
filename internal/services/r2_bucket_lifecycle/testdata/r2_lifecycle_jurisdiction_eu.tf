resource "cloudflare_r2_bucket" "%[1]s" {
  account_id   = "%[2]s"
  name         = "%[1]s"
  jurisdiction = "eu"
}

resource "cloudflare_r2_bucket_lifecycle" "%[1]s" {
  account_id   = "%[2]s"
  bucket_name = cloudflare_r2_bucket.%[1]s.name
  jurisdiction = "eu"

  rules = [{
    id      = "eu-compliance-delete"
    enabled = true
    conditions = {
      prefix = "gdpr/"
    }
    delete_objects_transition = {
      condition = {
        type    = "Age"
        max_age = 7776000 # 90 days
      }
    }
  }]
}

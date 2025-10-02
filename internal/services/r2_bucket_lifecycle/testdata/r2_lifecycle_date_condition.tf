resource "cloudflare_r2_bucket" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[1]s"
}

resource "cloudflare_r2_bucket_lifecycle" "%[1]s" {
  account_id  = "%[2]s"
  bucket_name = cloudflare_r2_bucket.%[1]s.name

  rules = [{
    id      = "delete-by-date"
    enabled = true
    conditions = {
      prefix = "temp/"
    }
    delete_objects_transition = {
      condition = {
        type = "Date"
        date = "2024-12-31T23:59:59Z"
      }
    }
  }]
}

resource "cloudflare_r2_bucket_lock" "example_r2_bucket_lock" {
  account_id = "023e105f4ecef8ad9ca31a8372d0c353"
  bucket_name = "example-bucket"
  rules = [{
    id = "Lock all objects for 24 hours"
    condition = {
      max_age_seconds = 100
      type = "Age"
    }
    enabled = true
    prefix = "prefix"
  }]
}

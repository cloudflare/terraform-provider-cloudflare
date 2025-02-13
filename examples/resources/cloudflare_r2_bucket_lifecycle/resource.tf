resource "cloudflare_r2_bucket_lifecycle" "example_r2_bucket_lifecycle" {
  account_id = "023e105f4ecef8ad9ca31a8372d0c353"
  bucket_name = "example-bucket"
  rules = [{
    id = "Expire all objects older than 24 hours"
    conditions = {
      prefix = "prefix"
    }
    enabled = true
    abort_multipart_uploads_transition = {
      condition = {
        max_age = 0
        type = "Age"
      }
    }
    delete_objects_transition = {
      condition = {
        max_age = 0
        type = "Age"
      }
    }
    storage_class_transitions = [{
      condition = {
        max_age = 0
        type = "Age"
      }
      storage_class = "InfrequentAccess"
    }]
  }]
}

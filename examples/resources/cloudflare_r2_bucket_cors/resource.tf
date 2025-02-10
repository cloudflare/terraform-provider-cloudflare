resource "cloudflare_r2_bucket_cors" "example_r2_bucket_cors" {
  account_id = "023e105f4ecef8ad9ca31a8372d0c353"
  bucket_name = "example-bucket"
  rules = [{
    allowed = {
      methods = ["GET"]
      origins = ["http://localhost:3000"]
      headers = ["x-requested-by"]
    }
    id = "Allow Local Development"
    expose_headers = ["Content-Encoding"]
    max_age_seconds = 3600
  }]
}

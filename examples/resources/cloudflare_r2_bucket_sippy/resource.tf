resource "cloudflare_r2_bucket_sippy" "example_r2_bucket_sippy" {
  account_id = "023e105f4ecef8ad9ca31a8372d0c353"
  bucket_name = "example-bucket"
  destination = {
    access_key_id = "accessKeyId"
    provider = {

    }
    secret_access_key = "secretAccessKey"
  }
  source = {
    access_key_id = "accessKeyId"
    bucket = "bucket"
    provider = "aws"
    region = "region"
    secret_access_key = "secretAccessKey"
  }
}

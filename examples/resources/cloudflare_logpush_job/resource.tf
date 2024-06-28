# Example Usage (Cloudflare R2)
#
# When using Cloudflare R2, no ownership challenge is required.
data "cloudflare_api_token_permission_groups" "all" {}
resource "cloudflare_api_token" "logpush_r2_token" {
  name = "logpush_r2_token"
  policy {
    permission_groups = [
      data.cloudflare_api_token_permission_groups.all.account["Workers R2 Storage Write"],
    ]
    resources = {
      "com.cloudflare.api.account.*" = "*"
    }
  }
}

resource "cloudflare_logpush_job" "http_requests" {
  enabled          = true
  zone_id          = var.zone_id
  name             = "http_requests"
  logpull_options  = "fields=ClientIP,ClientRequestHost,ClientRequestMethod,ClientRequestURI,EdgeEndTimestamp,EdgeResponseBytes,EdgeResponseStatus,EdgeStartTimestamp,RayID&timestamps=rfc3339"
  destination_conf = "r2://cloudflare-logs/http_requests/date={DATE}?account-id=${var.account_id}&access-key-id=${cloudflare_api_token.logpush_r2_token.id}&secret-access-key=${sha256(cloudflare_api_token.logpush_r2_token.value)}"
  dataset          = "http_requests"
}


# Example Usage (with AWS provider)
#
# Please see `cloudflare_logpush_ownership_challenge` for how to use that
# resource and the third party provider documentation if you
# choose to automate the intermediate step of fetching the ownership challenge
# contents.
#
# **Important:** If you're using this approach, the `destination_conf` values
# must match identically in all resources. Otherwise the challenge validation
# will fail.
resource "cloudflare_logpush_ownership_challenge" "ownership_challenge" {
  zone_id          = "0da42c8d2132a9ddaf714f9e7c920711"
  destination_conf = "s3://my-bucket-path?region=us-west-2"
}

data "aws_s3_bucket_object" "challenge_file" {
  bucket = "my-bucket-path"
  key    = cloudflare_logpush_ownership_challenge.ownership_challenge.ownership_challenge_filename
}

resource "cloudflare_logpush_job" "example_job" {
  depends_on          = [cloudflare_logpush_ownership_challenge.ownership_challenge]
  enabled             = true
  zone_id             = "0da42c8d2132a9ddaf714f9e7c920711"
  name                = "My-logpush-job"
  logpull_options     = "fields=RayID,ClientIP,EdgeStartTimestamp&timestamps=rfc3339"
  destination_conf    = "s3://my-bucket-path?region=us-west-2"
  ownership_challenge = data.aws_s3_bucket_object.challenge_file.body
  dataset             = "http_requests"
}

# Example Usage (manual inspection of S3 bucket)
#
# 1. Create `cloudflare_logpush_ownership_challenge` resource

resource "cloudflare_logpush_ownership_challenge" "ownership_challenge" {
  zone_id          = "0da42c8d2132a9ddaf714f9e7c920711"
  destination_conf = "s3://my-bucket-path?region=us-west-2"
}

# 2. Check S3 bucket for your ownership challenge filename and grab the contents.
# 3. Create the `cloudflare_logpush_job` substituting in your manual `ownership_challenge`.
resource "cloudflare_logpush_job" "example_job" {
  enabled             = true
  zone_id             = "0da42c8d2132a9ddaf714f9e7c920711"
  name                = "My-logpush-job"
  logpull_options     = "fields=RayID,ClientIP,EdgeStartTimestamp&timestamps=rfc3339"
  destination_conf    = "s3://my-bucket-path?region=us-west-2"
  ownership_challenge = "0000000000000"
  dataset             = "http_requests"
  frequency           = "high"
}

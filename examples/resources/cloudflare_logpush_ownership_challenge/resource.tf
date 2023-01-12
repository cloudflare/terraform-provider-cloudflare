resource "cloudflare_logpush_ownership_challenge" "example" {
  zone_id          = "d41d8cd98f00b204e9800998ecf8427e"
  destination_conf = "s3://my-bucket-path?region=us-west-2"
}

resource "cloudflare_logpush_ownership_challenge" "example" {
  zone_id          = "0da42c8d2132a9ddaf714f9e7c920711"
  destination_conf = "s3://my-bucket-path?region=us-west-2"
}

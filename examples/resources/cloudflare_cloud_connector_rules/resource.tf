resource "cloudflare_cloud_connector_rules" "example" {
  zone_id = "0da42c8d2132a9ddaf714f9e7c920711"

  rules {
    description = "connect aws bucket"
    enabled     = true
    expression  = "http.uri"
    provider    = "aws_s3"
    parameters {
      host = "mystorage.s3.ams.amazonaws.com"
    }
  }
}

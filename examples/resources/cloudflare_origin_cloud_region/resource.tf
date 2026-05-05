resource "cloudflare_origin_cloud_region" "example_origin_cloud_region" {
  zone_id = "023e105f4ecef8ad9ca31a8372d0c353"
  origin_ip = "192.0.2.1"
  region = "us-east-1"
  vendor = "aws"
}

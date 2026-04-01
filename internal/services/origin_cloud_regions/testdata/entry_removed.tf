resource "cloudflare_origin_cloud_regions" "%[1]s" {
  zone_id = "%[2]s"
  mappings = [
    {
      origin_ip = "192.0.2.1"
      vendor    = "aws"
      region    = "us-east-1"
    },
  ]
}

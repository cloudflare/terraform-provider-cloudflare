resource "cloudflare_origin_cloud_regions" "%[1]s" {
  zone_id = "%[2]s"
  mappings = [
    {
      origin_ip = "192.0.2.1"
      vendor    = "aws"
      region    = "us-west-2"
    },
    {
      origin_ip = "2001:db8::1"
      vendor    = "gcp"
      region    = "us-central1"
    },
  ]
}

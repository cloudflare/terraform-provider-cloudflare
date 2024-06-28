resource "cloudflare_observatory_scheduled_test" "example" {
  zone_id = "0da42c8d2132a9ddaf714f9e7c920711"
  url = "example.com"
  region = "us-central1"
  frequency = "WEEKLY"
}


resource "cloudflare_api_shield" "example" {
  zone_id = "0da42c8d2132a9ddaf714f9e7c920711"
  auth_id_characteristics {
    name = "my-example-header"
    type = "header"
  }
}

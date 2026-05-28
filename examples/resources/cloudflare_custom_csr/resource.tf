resource "cloudflare_custom_csr" "example_custom_csr" {
  common_name = "example.com"
  country = "US"
  locality = "San Francisco"
  organization = "Cloudflare, Inc."
  sans = ["example.com", "www.example.com"]
  state = "California"
  zone_id = "zone_id"
  description = "CSR for example.com wildcard"
  key_type = "rsa2048"
  name = "My Custom CSR"
  organizational_unit = "Engineering"
}

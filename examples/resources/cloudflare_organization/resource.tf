resource "cloudflare_organization" "example_organization" {
  name = "name"
  parent = {
    id = "a7b9c3d2e8f4a1b5c6d0e9f2a3b7c4d8"
  }
  profile = {
    business_address = "business_address"
    business_email = "business_email"
    business_name = "business_name"
    business_phone = "business_phone"
    external_metadata = "external_metadata"
  }
}

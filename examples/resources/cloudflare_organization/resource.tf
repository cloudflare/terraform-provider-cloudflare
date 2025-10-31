resource "cloudflare_organization" "example_organization" {
  name = "name"
  parent = {
    id = "a7b9c3d2e8f4g1h5i6j0k9l2m3n7o4p8"
  }
  profile = {
    business_address = "business_address"
    business_email = "business_email"
    business_name = "business_name"
    business_phone = "business_phone"
    external_metadata = "external_metadata"
  }
}

resource "cloudflare_organization" "example_organization" {
  name = "name"
  parent = {
    id = "id"
  }
  profile = {
    business_address = "business_address"
    business_email = "business_email"
    business_name = "business_name"
    business_phone = "business_phone"
    external_metadata = "external_metadata"
  }
}

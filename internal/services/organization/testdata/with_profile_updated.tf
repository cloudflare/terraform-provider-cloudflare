resource "cloudflare_organization" "%[1]s" {
  name = "%[2]s"
  
  profile = {
    business_name     = "Updated Business"
    business_email    = "updated@example.com"
    business_phone    = "+9876543210"
    business_address  = "456 Updated Ave, New City, NC 54321"
    external_metadata = "{\"key\":\"updated_value\"}"
  }
}

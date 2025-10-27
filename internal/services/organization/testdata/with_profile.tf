resource "cloudflare_organization" "%[1]s" {
  name = "%[2]s"

  profile = {
    business_name     = "Test Business"
    business_email    = "test@example.com"
    business_phone    = "+1234567890"
    business_address  = "123 Test St, Test City, TC 12345"
    external_metadata = "{\"key\":\"value\"}"
  }
}

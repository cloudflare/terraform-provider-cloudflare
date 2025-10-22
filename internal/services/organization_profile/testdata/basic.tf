resource "cloudflare_organization_profile" "%[1]s" {
  organization_id   = "%[2]s"
  business_name     = "%[3]s"
  business_email    = "%[4]s"
  business_phone    = "%[5]s"
  business_address  = "%[6]s"
  external_metadata = "%[7]s"
}

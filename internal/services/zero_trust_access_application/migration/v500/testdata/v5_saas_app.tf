resource "cloudflare_zero_trust_access_application" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[1]s"
  type       = "saas"

  saas_app = {
    consumer_service_url = "https://example.com/sso/saml/consume"
    sp_entity_id         = "example.com"
    name_id_format       = "email"

    custom_attributes = [{
      name        = "email"
      name_format = "urn:oasis:names:tc:SAML:2.0:attrname-format:basic"
      source      = { name = "user_email" }
    }]
  }
}

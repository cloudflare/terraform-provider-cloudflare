resource "cloudflare_zero_trust_access_application" "%[1]s" {
  account_id       = "%[2]s"
  name             = "%[1]s"
  type             = "saas"
  session_duration = "24h"
  saas_app = {
    consumer_service_url             = "https://saas-app.example/sso/saml/consume"
    sp_entity_id                     = "saas-app.example"
    name_id_format                   = "email"
    default_relay_state              = "https://saas-app.example"
    name_id_transform_jsonata        = "$substringBefore(email, '@') & '+sandbox@' & $substringAfter(email, '@')"
    saml_attribute_transform_jsonata = "$ ~>| groups | {'group_name': name} |"
    custom_attributes = [
      {
        name        = "email"
        name_format = "urn:oasis:names:tc:SAML:2.0:attrname-format:basic"
        source = { name = "user_email" }
      },
      {
        name          = "rank"
        required      = true
        friendly_name = "Rank"
        source = { name = "rank" }
      }
    ]
  }
  auto_redirect_to_identity = false
}

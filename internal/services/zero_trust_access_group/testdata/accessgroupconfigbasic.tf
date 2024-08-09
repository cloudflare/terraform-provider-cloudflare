
resource "cloudflare_zero_trust_access_group" "%[1]s" {
  %[3]s_id = "%[4]s"
  name     = "%[1]s"

  include {
    any_valid_service_token = true
    email = ["%[2]s"]
    email_domain = ["example.com"]
    ip = [
      "192.0.2.1/32",
      "192.0.2.2/32"
    ]
    ip_list = [
      "e3a0f205-c525-4e48-a293-ba5d1f00e638",
      "5d54cd30-ce52-46e4-9a46-a47887e1a167"
    ]
    saml = [{
      attribute_name = "Name1"
      attribute_value = "Value1"
    },
    {
      attribute_name = "Name2"
      attribute_value = "Value2"
    }]
    azure = [{
      id = ["group1"]
      identity_provider_id = "1234"
    },
    {
      id = ["group2"]
      identity_provider_id = "5678"
    }]
  }
}

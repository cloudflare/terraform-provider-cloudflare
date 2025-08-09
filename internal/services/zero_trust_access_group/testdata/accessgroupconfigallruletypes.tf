resource "cloudflare_zero_trust_access_group" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[1]s"

  include = [
    {
      email = {
        email = "test@example.com"
      }
    }
  ]
  
  exclude = [
    {
      geo = {
        country_code = "CN"
      }
    },
    {
      device_posture = {
        integration_uid = "test-device-posture-uid"
      }
    },
    {
      external_evaluation = {
        evaluate_url = "https://example.com/evaluate"
        keys_url     = "https://example.com/keys"
      }
    }
  ]

  require = [
    {
      auth_method = {
        auth_method = "hwk"
      }
    },
    {
      certificate = {}
    },
    {
      any_valid_service_token = {}
    }
  ]
}
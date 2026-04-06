resource "cloudflare_zero_trust_device_posture_integration" "%s" {
  account_id = "%s"
  name       = "%s"
  type       = "crowdstrike_s2s"
  interval   = "24h"

  config = {
    api_url       = "%s"
    client_id     = "%s"
    client_secret = "%s"
    customer_id   = "%s"
  }
}

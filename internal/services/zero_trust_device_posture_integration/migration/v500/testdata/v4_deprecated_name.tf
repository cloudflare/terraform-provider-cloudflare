resource "cloudflare_device_posture_integration" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[3]s"
  type       = "crowdstrike_s2s"
  interval   = "24h"

  config {
    api_url       = "%[4]s"
    client_id     = "%[5]s"
    client_secret = "%[6]s"
    customer_id   = "%[7]s"
  }
}

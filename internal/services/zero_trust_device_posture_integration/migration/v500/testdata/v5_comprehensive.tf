# Resource 1: Renamed from deprecated name, identifier removed, unusual interval preserved
resource "cloudflare_zero_trust_device_posture_integration" "%[1]s" {
  account_id = "%[4]s"
  name       = "%[5]s"
  type       = "crowdstrike_s2s"
  interval   = "1h"

  config = {
    api_url       = "%[8]s"
    client_id     = "%[9]s"
    client_secret = "%[10]s"
    customer_id   = "%[11]s"
  }
}

# Resource 2: Current name preserved, non-standard interval preserved
resource "cloudflare_zero_trust_device_posture_integration" "%[2]s" {
  account_id = "%[4]s"
  name       = "%[6]s"
  type       = "crowdstrike_s2s"
  interval   = "6h"

  config = {
    api_url       = "%[8]s"
    client_id     = "%[9]s"
    client_secret = "%[10]s"
    customer_id   = "%[11]s"
  }
}

# Resource 3: Renamed from deprecated name, standard interval preserved
resource "cloudflare_zero_trust_device_posture_integration" "%[3]s" {
  account_id = "%[4]s"
  name       = "%[7]s"
  type       = "crowdstrike_s2s"
  interval   = "24h"

  config = {
    api_url       = "%[8]s"
    client_id     = "%[9]s"
    client_secret = "%[10]s"
    customer_id   = "%[11]s"
  }
}

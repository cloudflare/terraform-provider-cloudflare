# Resource 1: Deprecated name + identifier field + unusual interval
resource "cloudflare_device_posture_integration" "%[1]s" {
  account_id = "%[4]s"
  name       = "%[5]s"
  type       = "crowdstrike_s2s"
  interval   = "1h"
  identifier = "legacy-identifier-to-remove"

  config {
    api_url       = "%[8]s"
    client_id     = "%[9]s"
    client_secret = "%[10]s"
    customer_id   = "%[11]s"
  }
}

# Resource 2: Current name with non-standard interval
resource "cloudflare_zero_trust_device_posture_integration" "%[2]s" {
  account_id = "%[4]s"
  name       = "%[6]s"
  type       = "crowdstrike_s2s"
  interval   = "6h"

  config {
    api_url       = "%[8]s"
    client_id     = "%[9]s"
    client_secret = "%[10]s"
    customer_id   = "%[11]s"
  }
}

# Resource 3: Deprecated name with standard interval
resource "cloudflare_device_posture_integration" "%[3]s" {
  account_id = "%[4]s"
  name       = "%[7]s"
  type       = "crowdstrike_s2s"
  interval   = "24h"

  config {
    api_url       = "%[8]s"
    client_id     = "%[9]s"
    client_secret = "%[10]s"
    customer_id   = "%[11]s"
  }
}

resource "cloudflare_zero_trust_device_posture_rule" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[3]s"
  type     = "firewall"
  schedule = "5m"

  match = [{
    platform = "windows"
    }, {
    platform = "mac"
    }, {
    platform = "linux"
  }]

  input = {
    enabled = true
  }
}

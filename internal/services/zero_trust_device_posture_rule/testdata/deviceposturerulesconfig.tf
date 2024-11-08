resource "cloudflare_zero_trust_device_posture_rule" "%[1]s" {
  account_id = "%[2]s"
  name        = "%[1]s"
  type        = "file"
  description = "check for /dev/random"
  schedule    = "1h"

  match =[ {
    platform = "linux"
  }]

  input = {
  path = "/dev/random"
}
}

data "cloudflare_zero_trust_device_posture_rules" "%[1]s" {
  account_id = "%[2]s"

  depends_on = [cloudflare_zero_trust_device_posture_rule.%[1]s]
}

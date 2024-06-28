resource "cloudflare_device_posture_rule" "eaxmple" {
  account_id  = "f037e56e89293a057740de681ac9abbe"
  name        = "Corporate devices posture rule"
  type        = "os_version"
  description = "Device posture rule for corporate devices."
  schedule    = "24h"
  expiration  = "24h"

  match {
    platform = "linux"
  }

  input {
    id                 = cloudflare_teams_list.corporate_devices.id
    version            = "1.0.0"
    operator           = "<"
    os_distro_name     = "ubuntu"
    os_distro_revision = "1.0.0"
    os_version_extra   = "(a)"
  }
}

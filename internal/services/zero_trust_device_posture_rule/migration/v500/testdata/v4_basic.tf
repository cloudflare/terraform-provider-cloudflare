resource "cloudflare_device_posture_rule" "%[1]s" {
  account_id  = "%[2]s"
  name        = "%[3]s"
  type        = "os_version"
  description = "Device posture rule for corporate devices."
  schedule    = "24h"
  expiration  = "25h"

  match {
    platform = "linux"
  }

  input {
    version            = "1.0.0"
    operator           = "<"
    os_distro_name     = "ubuntu"
    os_distro_revision = "1.0.0"
    os_version_extra   = "(a)"
  }
}

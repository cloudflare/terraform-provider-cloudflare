resource "cloudflare_zero_trust_device_posture_rule" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[3]s"
  type     = "disk_encryption"
  schedule = "5m"

  match = [{
    platform = "windows"
  }]

  input = {
    check_disks = ["C:", "D:"]
    require_all = true
  }
}

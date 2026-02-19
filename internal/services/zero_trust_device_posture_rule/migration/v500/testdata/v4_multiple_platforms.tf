resource "cloudflare_device_posture_rule" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[3]s"
  type       = "firewall"
  schedule   = "5m"

  match {
    platform = "windows"
  }

  match {
    platform = "mac"
  }

  match {
    platform = "linux"
  }

  input {
    enabled = true
  }
}

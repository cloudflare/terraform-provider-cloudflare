resource "cloudflare_device_posture_rule" "%[1]s" {
  account_id = "%[2]s"
  name       = ""
  type       = "os_version"
  schedule   = "5m"

  match {
    platform = "linux"
  }

  input {
    version  = "10.0.0"
    operator = ">="
  }
}

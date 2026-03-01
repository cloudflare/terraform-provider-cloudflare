resource "cloudflare_device_posture_rule" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[3]s"
  type       = "file"
  schedule   = "5m"

  match {
    platform = "windows"
  }

  input {
    path       = "C:\\Program Files\\app.exe"
    exists     = true
    thumbprint = "0123456789abcdef"
    sha256     = "abcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890"
  }
}

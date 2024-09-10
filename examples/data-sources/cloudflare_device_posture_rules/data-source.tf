data "cloudflare_device_posture_rules" "example" {
  account_id = "f037e56e89293a057740de681ac9abbe"
  name       = "check for /dev/random"
  type       = "file"
}

resource "cloudflare_byo_ip_prefix" "example" {
  account_id    = "f037e56e89293a057740de681ac9abbe"
  prefix_id     = "d41d8cd98f00b204e9800998ecf8427e"
  description   = "Example IP Prefix"
  advertisement = "on"
}

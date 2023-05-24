resource "cloudflare_r2_bucket" "example" {
  account_id = "f037e56e89293a057740de681ac9abbe"
  name       = "terraform-bucket"
  location   = "enam"
}

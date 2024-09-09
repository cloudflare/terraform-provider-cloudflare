resource "cloudflare_account_member" "example" {
  account_id    = "f037e56e89293a057740de681ac9abbe"
  email_address = "user@example.com"
  role_ids = [
    "68b329da9893e34099c7d8ad5cb9c940",
    "d784fa8b6d98d27699781bd9a7cf19f0"
  ]
}

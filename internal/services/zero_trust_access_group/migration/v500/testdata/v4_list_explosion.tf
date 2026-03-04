resource "cloudflare_access_group" "%s" {
  account_id = "%s"
  name       = "%s"

  include {
    email = ["user1@example.com", "user2@example.com", "user3@example.com"]
  }
}

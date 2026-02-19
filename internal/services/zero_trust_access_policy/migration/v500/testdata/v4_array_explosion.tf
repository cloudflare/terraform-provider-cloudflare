resource "cloudflare_access_policy" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[1]s"
  decision   = "allow"

  include {
    email = ["user1@example.com", "user2@example.com"]
  }

  exclude {
    email = ["blocked@example.com"]
  }
}

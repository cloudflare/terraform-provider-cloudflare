resource "cloudflare_access_group" "%s" {
  account_id = "%s"
  name       = "%s"

  # Test multiple selector types in same include block (like Pattern 2)
  include {
    email = ["user1@example.com", "user2@example.com"]
    ip    = ["192.168.1.0/24", "10.0.0.0/8"]
  }

  exclude {
    geo = ["US", "CA"]
  }
}

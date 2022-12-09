# Allowing access to `test@example.com` email address only
resource "cloudflare_access_group" "test_group" {
  account_id = "f037e56e89293a057740de681ac9abbe"
  name       = "staging group"

  include {
    email = ["test@example.com"]
  }
}

# Allowing `test@example.com` to access but only when coming from a
# specific IP.
resource "cloudflare_access_group" "test_group" {
  account_id = "f037e56e89293a057740de681ac9abbe"
  name       = "staging group"

  include {
    email = ["test@example.com"]
  }

  require = {
    ip = [var.office_ip]
  }
}

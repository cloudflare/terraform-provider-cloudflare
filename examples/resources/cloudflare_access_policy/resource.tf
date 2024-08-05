# Allowing access to `test@example.com` email address only
resource "cloudflare_access_policy" "test_policy" {
  account_id     = "f037e56e89293a057740de681ac9abbe"
  name           = "staging policy"
  decision       = "allow"

  include {
    email = ["test@example.com"]
  }

  require {
    email = ["test@example.com"]
  }
}

# Allowing `test@example.com` to access but only when coming from a
# specific IP.
resource "cloudflare_access_policy" "test_policy" {
  account_id     = "f037e56e89293a057740de681ac9abbe"
  name           = "staging policy"
  decision       = "allow"

  include {
    email = ["test@example.com"]
  }

  require {
    ip = [var.office_ip]
  }
}

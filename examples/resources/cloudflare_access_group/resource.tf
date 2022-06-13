# Allowing access to `test@example.com` email address only
resource "cloudflare_access_group" "test_group" {
  account_id = "975ecf5a45e3bcb680dba0722a420ad9"
  name       = "staging group"

  include {
    email = ["test@example.com"]
  }
}

# Allowing `test@example.com` to access but only when coming from a
# specific IP.
resource "cloudflare_access_group" "test_group" {
  account_id = "975ecf5a45e3bcb680dba0722a420ad9"
  name       = "staging group"

  include {
    email = ["test@example.com"]
  }

  require = {
    ip = [var.office_ip]
  }
}

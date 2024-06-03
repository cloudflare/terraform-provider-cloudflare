# Allowing access to `test@example.com` email address only
resource "cloudflare_access_policy" "test_policy" {
  account_id     = "0da42c8d2132a9ddaf714f9e7c920711"
  name           = "staging policy"
  precedence     = "1"
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
  account_id     = "0da42c8d2132a9ddaf714f9e7c920711"
  name           = "staging policy"
  precedence     = "1"
  decision       = "allow"

  include {
    email = ["test@example.com"]
  }

  require {
    ip = [var.office_ip]
  }
}

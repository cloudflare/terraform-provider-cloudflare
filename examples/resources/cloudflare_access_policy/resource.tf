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

# Access policy for an infrastructure application
resource "cloudflare_access_policy" "infra-app-example-allow" {
  application_id = cloudflare_zero_trust_access_application.infra-app-example.id
  account_id = "0da42c8d2132a9ddaf714f9e7c920711"
  name       = "infra-app-example-allow"
  decision   = "allow"
  precedence = 1

  include {
    email = ["devuser@gmail.com"]
  }

  connection_rules {
    ssh {
      usernames {
        value = "ec2-user"
      }
    }
  }
}

# Infrastructure application configuration for infra-app-example-allow
resource "cloudflare_zero_trust_access_application" "infra-app-example" {
  account_id = "0da42c8d2132a9ddaf714f9e7c920711"
  name       = "infra-app"
  type       = "infrastructure"
  
  target_criteria {
    port     = 22
    protocol = "SSH"
    target_attributes {
      name = "hostname"
      value {
        value = "tfgo-acc-tests"
      }
    }
  }

  # specify existing access policies by id
  policies = []
}

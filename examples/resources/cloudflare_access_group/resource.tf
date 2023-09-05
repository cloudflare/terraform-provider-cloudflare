# Allowing access to `test@example.com` email address only
resource "cloudflare_access_group" "example" {
  account_id = "f037e56e89293a057740de681ac9abbe"
  name       = "staging group"

  include {
    email = ["test@example.com"]
  }
}

# Allowing `test@example.com` to access but only when coming from a
# specific IP.
resource "cloudflare_access_group" "example" {
  account_id = "f037e56e89293a057740de681ac9abbe"
  name       = "staging group"

  include {
    email = ["test@example.com"]
  }

  require {
    ip = [var.office_ip]
  }
}

# Allow members of an Azure Group. The ID is the group UUID (id) in Azure.
resource "cloudflare_access_group" "example" {
  account_id = "f037e56e89293a057740de681ac9abbe"
  name       = "test_group"

  include {
    azure {
      identity_provider_id = "ca298b82-93b5-41bf-bc2d-10493f09b761"
      id                   = ["86773093-5feb-48dd-814b-7ccd3676ff50"]
    }
  }
}

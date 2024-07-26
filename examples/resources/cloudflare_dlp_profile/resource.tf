# Predefined profile must be imported, cannot be created
resource "cloudflare_dlp_profile" "creds" {
  account_id          = "f037e56e89293a057740de681ac9abbe"
  name                = "Credentials and Secrets"
  type                = "predefined"
  allowed_match_count = 3

  entry {
    enabled = true
    name    = "Amazon AWS Access Key ID"
    id      = "d8fcfc9c-773c-405e-8426-21ecbb67ba93"
  }
  entry {
    enabled = false
    id      = "2c0e33e1-71da-40c8-aad3-32e674ad3d96"
    name    = "Amazon AWS Secret Access Key"
  }
  entry {
    enabled = true
    id      = "4e92c006-3802-4dff-bbe1-8e1513b1c92a"
    name    = "Microsoft Azure Client Secret"
  }
  entry {
    enabled = false
    id      = "5c713294-2375-4904-abcf-e4a15be4d592"
    name    = "SSH Private Key"
  }
  entry {
    enabled = true
    id      = "6c6579e4-d832-42d5-905c-8e53340930f2"
    name    = "Google GCP API Key"
  }
}

# Custom profile
resource "cloudflare_dlp_profile" "example_custom" {
  account_id          = "f037e56e89293a057740de681ac9abbe"
  name                = "Example Custom Profile"
  description         = "A profile with example entries"
  type                = "custom"
  allowed_match_count = 0

  entry {
    name    = "Matches visa credit cards"
    enabled = true
    pattern {
      regex      = "4\\d{3}([-\\. ])?\\d{4}([-\\. ])?\\d{4}([-\\. ])?\\d{4}"
      validation = "luhn"
    }
  }

  entry {
    name    = "Matches diners club card"
    enabled = true
    pattern {
      regex      = "(?:0[0-5]|[68][0-9])[0-9]{11}"
      validation = "luhn"
    }
  }
}

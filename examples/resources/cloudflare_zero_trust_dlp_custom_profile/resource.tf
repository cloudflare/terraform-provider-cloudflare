resource "cloudflare_zero_trust_dlp_custom_profile" "example_zero_trust_dlp_custom_profile" {
	name        = "name"
	account_id  = "account_id"
	description = "Custom profile with entries"

  shared_entries = [
    {
      entry_id = "56a8c060-01bb-4f89-ba1e-3ad42770a342" // amex predefined entry
      entry_type = "predefined"
      enabled = true
    },
  ]
}

// Custom entry that is a part of this new profile
resource "cloudflare_zero_trust_dlp_custom_entry" "example_custom_entry" {
  name = "custom"
	account_id  = "account_id"
  profile_id = cloudflare_zero_trust_dlp_custom_profile.example_zero_trust_dlp_custom_profile.id
  pattern = {
    regex = "customentryregex"
  }

  enabled = true
}

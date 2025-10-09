resource "cloudflare_zero_trust_dlp_custom_entry" "%[1]s" {
  name = "%[1]s"
  account_id = "%[2]s"
  profile_id = cloudflare_zero_trust_dlp_custom_profile.%[1]s.id
  pattern = {
    regex = "customentryregex"
  }

  enabled = %[3]s
}

resource "cloudflare_zero_trust_dlp_custom_entry" "%[1]s-second" {
  name = "%[1]s-second"
  account_id = "%[2]s"
  profile_id = cloudflare_zero_trust_dlp_custom_profile.%[1]s.id
  pattern = {
    regex = "customentryregex"
  }

  enabled = %[3]s
}

resource "cloudflare_zero_trust_dlp_custom_profile" "%[1]s" {
	account_id  = "%[2]s"
	name        = "%[1]s"
	description = "Test with shared entries"

  shared_entries = [
    {
      entry_id = "56a8c060-01bb-4f89-ba1e-3ad42770a342" // amex predefined entry
      entry_type = "predefined"
      enabled = %[3]s
    }
  ]
}

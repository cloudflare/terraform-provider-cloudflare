data "cloudflare_zero_trust_dlp_predefined_entries" "example" {
  account_id = "%[2]s"
}

locals {
  # Find the first custom profile that accepts custom entries
  first_custom_entry = [for entry in data.cloudflare_zero_trust_dlp_predefined_entries.example.result : entry if entry.type == "predefined"][0]
}

resource "cloudflare_zero_trust_dlp_entry" "%[1]s" {
  account_id = "%[2]s"
  profile_id = local.first_custom_entry.profile_id
  name       = "test-dlp-entry-%[1]s"
  enabled    = true
  
  pattern = {
    regex = "[0-9]{4}"
  }
}
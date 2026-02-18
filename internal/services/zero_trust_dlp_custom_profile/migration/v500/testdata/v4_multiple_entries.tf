resource "cloudflare_dlp_profile" "%[1]s" {
  account_id          = "%[2]s"
  name                = "multi-pattern-%[1]s"
  type                = "custom"
  allowed_match_count = 10

  entry {
    name    = "Visa %[1]s"
    enabled = true
    pattern {
      regex      = "4[0-9]{12}(?:[0-9]{3})?"
      validation = "luhn"
    }
  }

  entry {
    name    = "SSN %[1]s"
    enabled = false
    pattern {
      regex = "[0-9]{3}-[0-9]{2}-[0-9]{4}"
    }
  }
}

resource "cloudflare_dlp_profile" "%[1]s" {
  account_id          = "%[2]s"
  name                = "complex-%[1]s"
  description         = "Complex pattern detection"
  type                = "custom"
  allowed_match_count = 3

  entry {
    name    = "Luhn %[1]s"
    enabled = true
    pattern {
      regex      = "3[47][0-9]{13}"
      validation = "luhn"
    }
  }

  entry {
    name    = "NoVal %[1]s"
    enabled = false
    pattern {
      regex = "[A-Z]{2}[0-9]{6}"
    }
  }
}

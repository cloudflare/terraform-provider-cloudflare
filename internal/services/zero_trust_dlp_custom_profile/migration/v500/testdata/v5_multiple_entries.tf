resource "cloudflare_zero_trust_dlp_custom_profile" "%[1]s" {
  account_id          = "%[2]s"
  name                = "multi-pattern-%[1]s"
  allowed_match_count = 10

  entries = [{
    name    = "Visa %[1]s"
    enabled = true
    pattern = {
      regex      = "4[0-9]{12}(?:[0-9]{3})?"
      validation = "luhn"
    }
    }, {
    name    = "SSN %[1]s"
    enabled = false
    pattern = {
      regex = "[0-9]{3}-[0-9]{2}-[0-9]{4}"
    }
  }]
}

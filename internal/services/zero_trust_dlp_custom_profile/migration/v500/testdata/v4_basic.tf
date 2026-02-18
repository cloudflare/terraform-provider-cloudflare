resource "cloudflare_dlp_profile" "%[1]s" {
  account_id          = "%[2]s"
  name                = "test-dlp-%[1]s"
  description         = "Test DLP profile"
  type                = "custom"
  allowed_match_count = 5

  entry {
    name    = "Test CC %[1]s"
    enabled = true
    pattern {
      regex      = "4[0-9]{12}(?:[0-9]{3})?"
      validation = "luhn"
    }
  }
}

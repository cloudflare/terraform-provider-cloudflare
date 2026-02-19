resource "cloudflare_dlp_profile" "%[1]s" {
  account_id          = "%[2]s"
  name                = "minimal-%[1]s"
  type                = "custom"
  allowed_match_count = 1

  entry {
    name    = "Simple %[1]s"
    enabled = true
    pattern {
      regex = "test[0-9]{3,5}"
    }
  }
}

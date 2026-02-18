resource "cloudflare_dlp_profile" "%[1]s" {
  account_id          = "%[2]s"
  name                = "Predefined Multi"
  type                = "predefined"
  allowed_match_count = 5

  entry {
    id      = "%[3]s"
    name    = "Entry One"
    enabled = true
  }

  entry {
    id      = "%[4]s"
    name    = "Entry Two"
    enabled = true
  }

  entry {
    id      = "%[5]s"
    name    = "Entry Three"
    enabled = false
  }
}

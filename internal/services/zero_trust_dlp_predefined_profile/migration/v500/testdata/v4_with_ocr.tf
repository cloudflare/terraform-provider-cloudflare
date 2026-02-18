resource "cloudflare_dlp_profile" "%[1]s" {
  account_id          = "%[2]s"
  name                = "OCR Predefined"
  type                = "predefined"
  allowed_match_count = 0
  ocr_enabled         = true

  entry {
    id      = "%[3]s"
    name    = "OCR Entry"
    enabled = true
  }
}

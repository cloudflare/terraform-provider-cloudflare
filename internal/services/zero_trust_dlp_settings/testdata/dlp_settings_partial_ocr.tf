resource "cloudflare_zero_trust_dlp_settings" "%[1]s" {
  account_id = "%[2]s"
  ocr        = true
}

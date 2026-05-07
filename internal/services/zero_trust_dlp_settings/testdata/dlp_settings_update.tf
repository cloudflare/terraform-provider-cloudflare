resource "cloudflare_zero_trust_dlp_settings" "%[1]s" {
  account_id          = "%[2]s"
  ai_context_analysis = false
  ocr                 = false
  payload_logging = {
    public_key    = "EmpOvSXw8BfbrGCi0fhGiD/3yXk2SiV1Nzg2lru3oj0="
    masking_level = "full"
  }
}

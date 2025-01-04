resource "cloudflare_zone_setting" "%[1]s" {
  zone_id = "%[2]s"
  setting_id = "ciphers"
  value = [
    "ECDHE-ECDSA-AES128-GCM-SHA256",
    "ECDHE-ECDSA-CHACHA20-POLY1305"
  ]
}

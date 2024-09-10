resource "cloudflare_hostname_tls_setting_ciphers" "example" {
  zone_id  = "0da42c8d2132a9ddaf714f9e7c920711"
  hostname = "sub.example.com"
  value    = ["ECDHE-RSA-AES128-GCM-SHA256"]
}

resource "cloudflare_hostname_tls_setting" "example" {
  zone_id  = "0da42c8d2132a9ddaf714f9e7c920711"
  hostname = "sub.example.com"
  setting  = "min_tls_version"
  value    = "1.2"
}

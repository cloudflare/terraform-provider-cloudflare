data "cloudflare_rulesets" "example" {
  zone_id = "0da42c8d2132a9ddaf714f9e7c920711"

  filter {
    name = ".*OWASP.*"
  }
}

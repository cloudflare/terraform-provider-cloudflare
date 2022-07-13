resource "cloudflare_zone" "example" {
  account_id = "d41d8cd98f00b204e9800998ecf8427e"
  zone       = "example.com"
}

resource "cloudflare_worker_domain" "example" {
  account_id = "f037e56e89293a057740de681ac9abbe"
  hostname   = "subdomain.example.com"
  service    = "my-service"
  zone_id    = "0da42c8d2132a9ddaf714f9e7c920711"
}

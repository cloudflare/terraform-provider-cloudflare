resource "cloudflare_pages_domain" "my-domain" {
  account_id   = "f037e56e89293a057740de681ac9abbe"
  project_name = "my-example-project"
  domain       = "example.com"
}

resource "cloudflare_workers_custom_domain" "example_workers_custom_domain" {
  account_id = "023e105f4ecef8ad9ca31a8372d0c353"
  hostname = "app.example.com"
  service = "my-worker"
  environment = "production"
  zone_id = "593c9c94de529bbbfaac7c53ced0447d"
  zone_name = "example.com"
}

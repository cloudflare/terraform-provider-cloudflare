data "cloudflare_workers_custom_domains" "example_workers_custom_domains" {
  account_id = "023e105f4ecef8ad9ca31a8372d0c353"
  environment = "production"
  hostname = "app.example.com"
  service = "my-worker"
  zone_id = "593c9c94de529bbbfaac7c53ced0447d"
  zone_name = "example.com"
}

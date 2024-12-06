data "cloudflare_logpush_dataset_field" "example_logpush_dataset_field" {
  dataset_id = "gateway_dns"
  account_id = "account_id"
  zone_id = "zone_id"
}

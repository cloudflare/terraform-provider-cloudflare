data "cloudflare_logpush_dataset_job" "example_logpush_dataset_job" {
  dataset_id = "gateway_dns"
  account_id = "account_id"
  zone_id = "zone_id"
}

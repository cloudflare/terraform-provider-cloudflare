data "cloudflare_custom_csr" "example_custom_csr" {
  custom_csr_id = "7b163417-1d2b-4c84-a38a-2fb7a0cd7752"
  account_id = "account_id"
  zone_id = "zone_id"
}

data "cloudflare_zero_trust_access_applications" "example_zero_trust_access_applications" {
  account_id = "account_id"
  zone_id = "zone_id"
  aud = "aud"
  domain = "domain"
  exact = true
  name = "name"
  search = "search"
}

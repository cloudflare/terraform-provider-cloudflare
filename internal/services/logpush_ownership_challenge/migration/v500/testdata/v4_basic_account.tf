resource "cloudflare_logpush_ownership_challenge" "%[1]s" {
  account_id       = "%[2]s"
  destination_conf = "gs://cf-terraform-provider-acct-test/ownership_challenges"
}

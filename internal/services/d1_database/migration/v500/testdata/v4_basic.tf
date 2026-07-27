resource "cloudflare_d1_database" "%[1]s" {
  account_id = "%[2]s"
  name       = "tf-acc-test-d1-%[1]s"
}

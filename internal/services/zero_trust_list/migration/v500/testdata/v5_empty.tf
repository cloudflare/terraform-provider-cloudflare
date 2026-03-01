resource "cloudflare_zero_trust_list" "%[1]s" {
  account_id = "%[2]s"
  name       = "tf-acc-test-%[1]s"
  type       = "SERIAL"
}

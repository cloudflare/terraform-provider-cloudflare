resource "cloudflare_zero_trust_list" "%[1]s" {
  account_id = "%[2]s"
  name       = "tf-acc-test-%[1]s"
  type       = "IP"
  items = [
    { value = "192.0.2.1" },
    { value = "192.0.2.2" }
  ]
}

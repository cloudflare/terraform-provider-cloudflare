resource "cloudflare_teams_list" "%[1]s" {
  account_id = "%[2]s"
  name       = "tf-acc-test-%[1]s"
  type       = "IP"
  items      = ["192.0.2.1", "192.0.2.2"]
}

resource "cloudflare_teams_list" "%[1]s" {
  account_id = "%[2]s"
  name       = "tf-acc-test-%[1]s"
  type       = "SERIAL"
  items      = []
}

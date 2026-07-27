resource "cloudflare_d1_database" "%[1]s" {
  account_id = "%[2]s"
  name       = "%[1]s"
  read_replication = {
    mode = "disabled"
  }
}

data "cloudflare_d1_databases" "%[1]s" {
  account_id = "%[2]s"
  name       = cloudflare_d1_database.%[1]s.name
  depends_on = [cloudflare_d1_database.%[1]s]
}

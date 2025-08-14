resource "cloudflare_zone" "%[1]s" {
  zone                = "%[2]s"
  account_id          = "%[3]s"
  paused              = true
  plan                = "enterprise"
  type                = "full"
  jump_start          = false
  vanity_name_servers = ["ns1.%[2]s", "ns2.%[2]s"]
}
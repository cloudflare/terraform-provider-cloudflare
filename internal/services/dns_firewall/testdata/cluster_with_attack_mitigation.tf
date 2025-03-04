resource "cloudflare_dns_firewall" "%[1]s" {
  account_id = "%[2]s"
  name = "%[3]s"
  upstream_ips = ["%[4]s"]
  ratelimit = %[5]s

  attack_mitigation = {
    enabled = true
  }
}

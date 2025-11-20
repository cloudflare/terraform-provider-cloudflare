
resource "cloudflare_dns_record" "%[2]s" {
  zone_id = "%[1]s"
  name = "_xmpp-client._tcp.%[2]s.%[3]s"
  priority = 5
  data = {
    priority = 5
    weight = 0
    port = 5222
    target = "talk.l.google.com"
    proto = "_tcp"
    name = "%[2]s.%[3]s"
}
  type = "SRV"
  ttl = 3600
}

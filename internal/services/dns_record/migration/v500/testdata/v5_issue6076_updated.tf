resource "cloudflare_dns_record" "%[1]s" {
  zone_id = "%[2]s"
  comment = "updated comment for testing"
  name    = "%[3]s"
  type    = "CNAME"
  content = "kay.ns.cloudflare.com"
  ttl     = 1
  proxied = true
  tags    = ["test", "migration"]
}

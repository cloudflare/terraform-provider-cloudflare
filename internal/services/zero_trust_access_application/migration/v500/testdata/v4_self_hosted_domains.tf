resource "cloudflare_access_application" "%[1]s" {
  account_id          = "%[2]s"
  name                = "%[1]s"
  type                = "self_hosted"
  self_hosted_domains = ["%[1]s-a.%[3]s", "%[1]s-b.%[3]s"]
}

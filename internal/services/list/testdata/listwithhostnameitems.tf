resource "cloudflare_list" "%[1]s" {
  account_id = "%[4]s"
  name = "%[2]s"
  description = "%[3]s"
  kind = "hostname"
  items = [
    {
      hostname = {
        url_hostname = "example.com"
      }
    },
    {
      hostname = {
        url_hostname = "*.a.example.com"
        exclude_exact_hostname = true
      }
    },
  ]
}

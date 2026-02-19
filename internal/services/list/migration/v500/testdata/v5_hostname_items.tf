resource "cloudflare_list" "%[1]s" {
  account_id  = "%[2]s"
  name        = "%[3]s"
  kind        = "hostname"
  description = "%[4]s"

  items = [{
    comment = "Test hostname 1"
    hostname = {
      url_hostname = "example.com"
    }
  }, {
    hostname = {
      url_hostname = "test.example.com"
    }
  }]
}

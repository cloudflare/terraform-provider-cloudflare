resource "cloudflare_list" "%[1]s" {
  account_id = "%[4]s"
  name = "%[2]s"
  description = "%[3]s"
  kind = "redirect"
  items = [
    {
      redirect = {
        source_url = "example.com/1"
        target_url = "https://one.example.com"
        status_code = 301
      }
    },
    {
      redirect = {
        source_url = "example.com/2"
        target_url = "https://two.example.com"
        status_code = 301
      }
    },
  ]
}

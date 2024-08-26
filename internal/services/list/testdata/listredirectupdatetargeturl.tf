
  resource "cloudflare_list" "%[1]s" {
    account_id = "%[4]s"
    name = "%[2]s"
    description = "%[3]s"
    kind = "redirect"

    item =[ {
      value =[ {
        redirect = {
          source_url = "cloudflare.com/blog"
          target_url = "https://theblog.cloudflare.com"
        }
      }]
    }]
  }

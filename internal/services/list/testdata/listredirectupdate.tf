
  resource "cloudflare_list" "%[1]s" {
    account_id = "%[4]s"
    name = "%[2]s"
    description = "%[3]s"
    kind = "redirect"

    item =[ {
      value =[ {
        redirect = {
          source_url = "cloudflare.com/blog"
          target_url = "https://blog.cloudflare.com"
        }
      }]
      comment = "one"
    },
    {
    value =[ {
        redirect = {
          source_url = "cloudflare.com/foo"
          target_url = "https://foo.cloudflare.com"
          include_subdomains = "enabled"
          subpath_matching = "enabled"
          status_code = 301
          preserve_query_string = "enabled"
          preserve_path_suffix = "disabled"
		}
      }]
      comment = "two"
    }]

  }

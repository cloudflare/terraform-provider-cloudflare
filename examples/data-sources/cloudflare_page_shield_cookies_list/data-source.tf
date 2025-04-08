data "cloudflare_page_shield_cookies_list" "example_page_shield_cookies_list" {
  zone_id = "023e105f4ecef8ad9ca31a8372d0c353"
  direction = "asc"
  domain = "example.com"
  export = "csv"
  hosts = "blog.cloudflare.com,www.example*,*cloudflare.com"
  http_only = true
  name = "session_id"
  order_by = "first_seen_at"
  page = "2"
  page_url = "example.com/page,*/checkout,example.com/*,*checkout*"
  path = "/"
  per_page = 100
  same_site = "strict"
  secure = true
  type = "first_party"
}

data "cloudflare_page_shield_scripts_list" "example_page_shield_scripts_list" {
  zone_id = "023e105f4ecef8ad9ca31a8372d0c353"
  direction = "asc"
  exclude_urls = "blog.cloudflare.com,www.example"
  export = "csv"
  hosts = "blog.cloudflare.com,www.example*,*cloudflare.com"
  order_by = "first_seen_at"
  page = "2"
  page_url = "example.com/page,*/checkout,example.com/*,*checkout*"
  per_page = 100
  prioritize_malicious = true
  status = "active,inactive"
  urls = "blog.cloudflare.com,www.example"
}

resource "cloudflare_load_balancer_monitor" "%s" {
  account_id = "%s"
  type       = "https"
  method     = "GET"
  path       = "/health"

  header = {
    "Host"       = ["%s"]
    "User-Agent" = ["Cloudflare-Traffic-Manager/1.0"]
  }
}

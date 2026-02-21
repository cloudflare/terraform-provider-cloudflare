resource "cloudflare_load_balancer_monitor" "%s" {
  account_id = "%s"
  type       = "https"
  method     = "GET"
  path       = "/health"

  header {
    header = "Host"
    values = ["%s"]
  }

  header {
    header = "User-Agent"
    values = ["Cloudflare-Traffic-Manager/1.0"]
  }
}

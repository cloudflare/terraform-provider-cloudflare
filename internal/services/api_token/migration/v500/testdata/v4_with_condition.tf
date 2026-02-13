resource "cloudflare_api_token" "%[1]s" {
  name = "%[1]s"

  policy {
    permission_groups = [
      "82e64a83756745bbbb1c9c2701bf816b"
    ]
    resources = {
      "com.cloudflare.api.account.zone.*" = "*"
    }
  }

  condition {
    request_ip {
      in     = ["192.0.2.1/32"]
      not_in = ["198.51.100.0/24"]
    }
  }
}

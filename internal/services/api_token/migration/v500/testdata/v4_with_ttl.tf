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

  not_before = "2027-01-01T00:00:00Z"
  expires_on = "2027-12-31T23:59:59Z"
}

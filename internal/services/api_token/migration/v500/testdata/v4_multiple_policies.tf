resource "cloudflare_api_token" "%[1]s" {
  name = "%[1]s"

  policy {
    permission_groups = [
      "82e64a83756745bbbb1c9c2701bf816b"
    ]
    resources = {
      "com.cloudflare.api.account.%[2]s" = "*"
    }
  }

  policy {
    permission_groups = [
      "e199d584e69344eba202452019deafe3"
    ]
    resources = {
      "com.cloudflare.api.account.zone.*" = "*"
    }
  }
}

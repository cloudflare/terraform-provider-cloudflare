resource "cloudflare_api_token" "%[1]s" {
  name = "%[1]s"

  policies = [
    {
      effect = "allow"
      permission_groups = [
        { id = "82e64a83756745bbbb1c9c2701bf816b" },
        { id = "c8fed203ed3043cba015a93ad1616f1f" },
        { id = "e199d584e69344eba202452019deafe3" }
      ]
      resources = jsonencode({
        "com.cloudflare.api.account.%[2]s" = "*"
      })
    }
  ]
}

resource "cloudflare_account_token" "%[1]s" {
  account_id = "%[2]s"
  name = "%[1]s"

  policies = [{
    effect = "allow"
    permission_groups = [{
      id = "82e64a83756745bbbb1c9c2701bf816b"
    }]
    resources = jsonencode({
      "com.cloudflare.api.account.%[2]s" = "*"
    })
  }]

  not_before = "%[3]s"
  expires_on = "%[4]s"
}

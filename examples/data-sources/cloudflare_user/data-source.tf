data "cloudflare_user" "me" {}
data "cloudflare_api_token_permission_groups" "all" {}

resource "cloudflare_api_token" "example" {
  name = "Terraform Cloud (Terraform)"
  policy {
    permission_groups = [
      data.cloudflare_api_token_permission_groups.all.user["User Details Read"],

    ]
    resources = {
      "com.cloudflare.api.user.${data.cloudflare_user.me.id}" = "*",
    }
  }
}

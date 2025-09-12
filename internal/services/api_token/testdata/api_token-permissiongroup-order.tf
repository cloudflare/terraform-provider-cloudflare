resource "cloudflare_api_token" "%[1]s" {
  name = "%[1]s"
  status = "active"

  policies = [{
    effect = "allow"
    permission_groups = toset([{
      id = "%[2]s"
    },{
      id = "%[3]s"
    }])
    resources = {
      "com.cloudflare.api.account.zone.*" = "*"
    }
  }]
}

data "cloudflare_api_token" "%[1]s" {
  token_id = cloudflare_api_token.%[1]s.id
  depends_on = [cloudflare_api_token.%[1]s]
}

# data "cloudflare_api_token_permission_groups_list" "dns_read" {
#   name  = "DNS Read"
#   scope = "com.cloudflare.api.account.zone"
# }

# data "cloudflare_api_token_permission_groups_list" "disable_esc_read" {
#   name  = "Disable ESC Read"
#   scope = "com.cloudflare.api.account.zone"
# }

# resource "cloudflare_api_token" "test_account_token" {
#   name   = "%[1]s"
#   status = "active"

#   policies = [{
#     effect = "allow"
#     permission_groups = [{
#       id = "%[2]s" != "" ? "%[2]s" : data.cloudflare_api_token_permission_groups_list.dns_read.result[0].id
#       }, {
#       id = "%[3]s" != "" ? "%[3]s" : data.cloudflare_api_token_permission_groups_list.disable_esc_read.result[0].id
#     }]
#     resources = {
#       "com.cloudflare.api.account.zone.*" = "*"
#     }
#   }]
# }

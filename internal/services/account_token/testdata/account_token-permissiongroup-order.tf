resource "cloudflare_account_token" "%[1]s" {
  name = "%[1]s"
  account_id = "%[2]s"

  policies = [{
    effect = "allow"
    permission_groups = [{
      id = "%[3]s"
    },{
      id = "%[4]s"
    }]
    resources = {
      "com.cloudflare.api.account.%[2]s" = "*"
    }
  }]
}

data "cloudflare_account_token" "%[1]s" {
  account_id = "%[2]s"
  token_id = cloudflare_account_token.%[1]s.id
  depends_on = [cloudflare_account_token.%[1]s]
}

# data "cloudflare_account_api_token_permission_groups_list" "dns_read" {
#   account_id = "%[2]s"
#   name       = "DNS Read"
#   scope      = "com.cloudflare.api.account.zone"
# }

# data "cloudflare_account_api_token_permission_groups_list" "disable_esc_read" {
#   account_id = "%[2]s"
#   name       = "Disable ESC Read"
#   scope      = "com.cloudflare.api.account.zone"
# }

# resource "cloudflare_account_token" "test_account_token" {
#   name       = "%[1]s"
#   account_id = "%[2]s"

#   policies = [{
#     effect = "allow"
#     permission_groups = [{
#       id = "%[3]s" != "" ? "%[3]s" : data.cloudflare_account_api_token_permission_groups_list.dns_read.result[0].id
#       }, {
#       id = "%[4]s" != "" ? "%[4]s" : data.cloudflare_account_api_token_permission_groups_list.disable_esc_read.result[0].id
#     }]
#     resources = {
#       "com.cloudflare.api.account.%[2]s" = "*"
#     }
#   }]
# }

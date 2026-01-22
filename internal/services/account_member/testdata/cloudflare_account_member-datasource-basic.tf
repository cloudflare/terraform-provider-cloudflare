data "cloudflare_account_members" "%[1]s" {
  account_id = "%[2]s"
}

data "cloudflare_account_member" "%[1]s" {
  account_id = "%[2]s"
  member_id  = data.cloudflare_account_members.%[1]s.result[0].id
}
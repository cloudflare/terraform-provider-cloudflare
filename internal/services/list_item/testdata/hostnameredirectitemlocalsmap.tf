locals {
  redirect_items = {
    "redirect1" = {
      source_url  = "%[1]s-redirect1.cfapi.net/"
      target_url  = "https://target1.com"
      status_code = 301
    }
    "redirect2" = {
      source_url  = "%[1]s-redirect2.cfapi.net/"
      target_url  = "https://target2.com"
      status_code = 302
    }
  }
}

resource "cloudflare_list" "%[2]s" {
  account_id  = "%[4]s"
  name        = "%[2]s"
  description = "list named %[2]s"
  kind        = "redirect"
}

resource "cloudflare_list_item" "%[1]s_redirect1" {
  account_id = "%[4]s"
  list_id    = cloudflare_list.%[2]s.id
  comment    = "%[3]s-redirect1"
  redirect = {
    source_url  = local.redirect_items["redirect1"].source_url
    target_url  = local.redirect_items["redirect1"].target_url
    status_code = local.redirect_items["redirect1"].status_code
  }
}

resource "cloudflare_list_item" "%[1]s_redirect2" {
  account_id = "%[4]s"
  list_id    = cloudflare_list.%[2]s.id
  comment    = "%[3]s-redirect2"
  redirect = {
    source_url  = local.redirect_items["redirect2"].source_url
    target_url  = local.redirect_items["redirect2"].target_url
    status_code = local.redirect_items["redirect2"].status_code
  }
}

locals {
  redirect_items = {
    "redirect1" = {
      source_url  = "example1.com/"
      target_url  = "https://target1.com"
      status_code = 301
    }
    "redirect2" = {
      source_url  = "example2.com/"
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

resource "cloudflare_list_item" "%[1]s" {
  for_each = local.redirect_items
  account_id = "%[4]s"
  list_id    = cloudflare_list.%[2]s.id
  comment    = "%[3]s-${each.key}"
  redirect = {
    source_url  = each.value.source_url
    target_url  = each.value.target_url
    status_code = each.value.status_code
  }
}

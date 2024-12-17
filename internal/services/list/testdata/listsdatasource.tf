
	resource "cloudflare_list" "%[1]s" {
		account_id = "%[2]s"
		name = "%[1]s"
		description = "example list"
		kind = "ip"
	}

data "cloudflare_lists" "%[1]s" {
  account_id = "%[2]s"
  depends_on = [ cloudflare_list.%[1]s ]
}

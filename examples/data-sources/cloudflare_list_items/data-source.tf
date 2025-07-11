data "cloudflare_list_items" "example_list_items" {
  account_id = "023e105f4ecef8ad9ca31a8372d0c353"
  list_id = "2c0fc9fa937b11eaa1b71c4d701ab86e"
  search = "1.1.1."
}

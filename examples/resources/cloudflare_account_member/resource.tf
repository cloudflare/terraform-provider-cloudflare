resource "cloudflare_account_member" "example_account_member" {
  account_id = "023e105f4ecef8ad9ca31a8372d0c353"
  email = "user@example.com"
  roles = ["3536bcfad5faccb999b47003c79917fb"]
  status = "accepted"
}
